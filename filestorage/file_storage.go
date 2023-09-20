package filestorage

import (
	"bytes"
	"io"

	filestorage_grpc "github.com/PongDev/Go-gRPC-Storage/filestorage/grpc"
	"github.com/PongDev/Go-gRPC-Storage/minio"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type FileStorageService struct {
	filestorage_grpc.UnimplementedFileUploadServiceServer
	minio *minio.MinIO
}

func RegisterFileUploadServiceServer(s grpc.ServiceRegistrar, srv filestorage_grpc.FileUploadServiceServer) {
	s.RegisterService(&filestorage_grpc.FileUploadService_ServiceDesc, srv)
}

func NewFileStorageService() *FileStorageService {
	minio, err := minio.NewMinIOClient()
	if err != nil {
		panic(err)
	}

	return &FileStorageService{
		minio: minio,
	}
}

func (f *FileStorageService) Upload(stream filestorage_grpc.FileUploadService_UploadServer) error {
	fileBuffer := &bytes.Buffer{}
	fileName := ""

	for {
		buffer, err := stream.Recv()
		if err == io.EOF {
			fileInfo, err := f.minio.UploadFile(fileName, io.Reader(fileBuffer), int64(fileBuffer.Len()), "")
			if err != nil {
				return err
			}
			return stream.SendAndClose(&filestorage_grpc.FileUploadResponse{
				FileName: fileInfo.Key,
				Size:     uint32(fileInfo.Size),
			})
		} else if err != nil {
			return err
		}
		fileBuffer.Write(buffer.FileContent)
		if buffer.FileName != nil {
			fileName = uuid.NewString() + "_" + *buffer.FileName
		}
	}
}

func (f *FileStorageService) Download(req *filestorage_grpc.FileDownloadRequest, stream filestorage_grpc.FileUploadService_DownloadServer) error {
	obj, err := f.minio.DownloadFile(req.FileName)
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
