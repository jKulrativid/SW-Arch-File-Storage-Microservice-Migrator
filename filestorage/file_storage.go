package filestorage

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strconv"

	filestorage_grpc "github.com/PongDev/SW-Arch-File-Storage-Microservice/grpc/filestorage"
	"github.com/PongDev/SW-Arch-File-Storage-Microservice/grpc/subject"
	"github.com/PongDev/SW-Arch-File-Storage-Microservice/minio"
	db "github.com/PongDev/SW-Arch-File-Storage-Microservice/prisma/prisma-client"
	"github.com/PongDev/SW-Arch-File-Storage-Microservice/repository"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type FileStorageService struct {
	filestorage_grpc.UnimplementedFileUploadServiceServer
	minio                minio.MinIO
	subjectServiceClient subject.SubjectServiceClient
	shareFileRepo        repository.ShareFileRepository
	bookmarkFileRepo     repository.BookmarkFileRepository
	fileInformationRepo  repository.FileInformationRepository
}

func RegisterFileUploadServiceServer(s grpc.ServiceRegistrar, srv filestorage_grpc.FileUploadServiceServer) {
	s.RegisterService(&filestorage_grpc.FileUploadService_ServiceDesc, srv)
}

func NewFileStorageService(prismaClient *db.PrismaClient, subjectServiceClient subject.SubjectServiceClient) *FileStorageService {
	minio, err := minio.NewMinIOClient()
	if err != nil {
		panic(err)
	}

	return &FileStorageService{
		minio:                minio,
		subjectServiceClient: subjectServiceClient,
		shareFileRepo:        repository.NewShareFileRepository(prismaClient),
		bookmarkFileRepo:     repository.NewBookmarkFileRepository(prismaClient),
		fileInformationRepo:  repository.NewFileInformationRepository(prismaClient),
	}
}

func (f *FileStorageService) IsUserFileOwner(ctx context.Context, userID, fileID string) bool {
	fileInfo, err := f.fileInformationRepo.GetFileInformation(ctx, fileID)
	if err != nil {
		return false
	}
	return fileInfo.OwnerUserID == userID
}

func (f *FileStorageService) IsUserCanAccessFile(ctx context.Context, userID, fileID string) bool {
	return f.IsUserFileOwner(ctx, userID, fileID) ||
		f.shareFileRepo.CheckIsFileShareWithUser(ctx, fileID, userID) ||
		f.shareFileRepo.CheckIsFileShareWithUser(ctx, fileID, "*")
}

func (f *FileStorageService) validateSubject(ctx context.Context, subjectId string) bool {
	subjectIdInteger, err := strconv.ParseInt(subjectId, 10, 64)
	if err != nil {
		return false
	}
	result, err := f.subjectServiceClient.ValidateSubjectId(ctx, &subject.ValidateSubjectIdRequest{
		Id: subjectIdInteger,
	})
	if err != nil {
		return false
	}

	if !result.Valid {
		return false
	}
	return true
}

func (f *FileStorageService) Upload(stream filestorage_grpc.FileUploadService_UploadServer) error {
	fileBuffer := &bytes.Buffer{}
	fileName := ""
	userID := ""
	subjectId := ""

	for {
		buffer, err := stream.Recv()
		if err == io.EOF {
			if userID == "" || subjectId == "" {
				return errors.New("Empty user_id or subject_id")
			}

			if !f.validateSubject(stream.Context(), subjectId) {
				return errors.New("Invalid subject_id")
			}

			fileID, err := f.fileInformationRepo.CreateFileInformation(context.Background(), uuid.NewString(), subjectId, userID, fileName)
			if err != nil {
				return err
			}

			fileInfo, err := f.minio.UploadFile(fileID, io.Reader(fileBuffer), int64(fileBuffer.Len()), "")
			if err != nil {
				return err
			}
			return stream.SendAndClose(&filestorage_grpc.FileUploadResponse{
				FileId: fileID,
				Size:   uint32(fileInfo.Size),
			})
		} else if err != nil {
			return err
		}
		fileBuffer.Write(buffer.FileContent)
		if buffer.FileName != nil {
			fileName = *buffer.FileName
		}
		if buffer.UserId != nil {
			userID = *buffer.UserId
		}
		if buffer.SubjectId != nil {
			subjectId = *buffer.SubjectId
		}
	}
}

func (f *FileStorageService) Download(req *filestorage_grpc.FileDownloadRequest, stream filestorage_grpc.FileUploadService_DownloadServer) error {
	if !f.IsUserCanAccessFile(stream.Context(), req.UserId, req.FileId) {
		return errors.New("User can't access this file")
	}
	obj, err := f.minio.DownloadFile(req.FileId)
	if err != nil {
		return err
	}
	defer obj.Close()

	info, err := obj.Stat()
	if err != nil {
		return err
	}

	sendMetaData := false
	fileSize := uint32(info.Size)
	fileName := info.Key

	buffer := make([]byte, 1024)
	isEOF := false
	for {
		n, err := obj.Read(buffer)
		if err == io.EOF {
			if n == 0 {
				break
			} else {
				isEOF = true
			}
		} else if err != nil {
			return err
		}
		res := &filestorage_grpc.FileDownloadResponse{
			FileContent: buffer[:n],
		}
		if !sendMetaData {
			res.Size = &fileSize
			res.FileName = &fileName
			sendMetaData = true
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		if isEOF {
			break
		}
	}
	return nil
}

func (f *FileStorageService) Delete(ctx context.Context, req *filestorage_grpc.FileDeleteRequest) (*filestorage_grpc.FileDeleteResponse, error) {
	if !f.IsUserFileOwner(ctx, req.UserId, req.FileId) {
		return nil, errors.New("User can't delete this file")
	}
	err := f.minio.DeleteFile(req.FileId)
	if err != nil {
		return nil, err
	}

	return &filestorage_grpc.FileDeleteResponse{
		FileId: req.FileId,
	}, nil
}

func (f *FileStorageService) ShareFile(ctx context.Context, req *filestorage_grpc.ShareFileRequest) (*filestorage_grpc.ShareFileResponse, error) {
	if !f.IsUserFileOwner(ctx, req.UserId, req.FileId) {
		return nil, errors.New("User can't share this file")
	}
	fileID, err := f.shareFileRepo.CreateShareFile(ctx, req.FileId, req.ShareWithUserIds)
	if err != nil {
		return nil, err
	}
	return &filestorage_grpc.ShareFileResponse{
		FileId: fileID,
	}, nil
}

func (f *FileStorageService) CreateBookmarkFile(ctx context.Context, req *filestorage_grpc.CreateBookmarkFileRequest) (*filestorage_grpc.CreateBookmarkFileResponse, error) {
	if !f.IsUserCanAccessFile(ctx, req.UserId, req.FileId) {
		return nil, errors.New("User can't bookmark this file")
	}
	fileID, err := f.bookmarkFileRepo.CreateBookmark(ctx, req.UserId, req.FileId)
	if err != nil {
		return nil, err
	}
	return &filestorage_grpc.CreateBookmarkFileResponse{
		FileId: fileID,
	}, nil
}

func (f *FileStorageService) DeleteBookmarkFile(ctx context.Context, req *filestorage_grpc.DeleteBookmarkFileRequest) (*filestorage_grpc.DeleteBookmarkFileResponse, error) {
	fileID, err := f.bookmarkFileRepo.DeleteBookmark(ctx, req.UserId, req.FileId)
	if err != nil {
		return nil, err
	}
	return &filestorage_grpc.DeleteBookmarkFileResponse{
		FileId: fileID,
	}, nil
}

func (f *FileStorageService) GetBookmarkFiles(ctx context.Context, req *filestorage_grpc.GetBookmarkFilesRequest) (*filestorage_grpc.GetBookmarkFilesResponse, error) {
	fileIDs, err := f.bookmarkFileRepo.GetBookmark(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &filestorage_grpc.GetBookmarkFilesResponse{
		FileIds: fileIDs,
	}, nil
}
