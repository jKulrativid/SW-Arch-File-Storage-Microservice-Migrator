package minio

import (
	"context"
	"io"
	"log"

	"github.com/PongDev/Go-gRPC-Storage/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIO struct {
	*minio.Client
	bucketName string
}

func NewMinIOClient() (*MinIO, error) {
	ctx := context.Background()
	endpoint := config.Env.MINIO_ENDPOINT
	accessKey := config.Env.MINIO_ACCESS_KEY
	secretKey := config.Env.MINIO_SECRET_KEY
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("MinIO Client Error: %s\n", err)
		return nil, err
	}

	bucketName := config.Env.MINIO_BUCKET
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, err := minioClient.BucketExists(ctx, bucketName)
		if err == nil && exists {
			log.Printf("Bucket: %s already exists\n", bucketName)
		} else {
			log.Fatalf("MinIO Check Bucket Exists Error: %s\n", err)
			return nil, err
		}
	} else {
		log.Printf("Successfully created bucket: %s\n", bucketName)
	}
	return &MinIO{Client: minioClient, bucketName: bucketName}, nil
}

func (m *MinIO) UploadFile(objectName string, fileBuffer io.Reader, fileSize int64, contentType string) (minio.UploadInfo, error) {
	info, err := m.PutObject(context.Background(), m.bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return info, nil
}

func (m *MinIO) DownloadFile(objectName string) (*minio.Object, error) {
	obj, err := m.GetObject(context.Background(), m.bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}
