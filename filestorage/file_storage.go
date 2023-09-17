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
