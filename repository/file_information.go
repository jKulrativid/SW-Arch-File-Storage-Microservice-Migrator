package repository

import (
	"context"

	db "github.com/PongDev/SW-Arch-File-Storage-Microservice/prisma/prisma-client"
)

type FileInformationRepository interface {
	GetFileInformation(ctx context.Context, id string) (*db.FileInformationModel, error)
	CreateFileInformation(ctx context.Context, id, subjectID, ownerUserID, fileName string) (string, error)
	DeleteFileInformation(ctx context.Context, id string) (string, error)
}

type fileInformationRepository struct {
	client *db.PrismaClient
}

func NewFileInformationRepository(client *db.PrismaClient) FileInformationRepository {
	return &fileInformationRepository{client: client}
}

func (r *fileInformationRepository) GetFileInformation(ctx context.Context, id string) (*db.FileInformationModel, error) {
	result, err := r.client.FileInformation.FindUnique(
		db.FileInformation.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *fileInformationRepository) CreateFileInformation(ctx context.Context, id, subjectID, ownerUserID, fileName string) (string, error) {
	result, err := r.client.FileInformation.CreateOne(
		db.FileInformation.SubjectID.Set(subjectID),
		db.FileInformation.OwnerUserID.Set(ownerUserID),
		db.FileInformation.FileName.Set(fileName),
		db.FileInformation.ID.Set(id),
	).Exec(ctx)

	if err != nil {
		return "", err
	}
	return result.ID, nil
}

func (r *fileInformationRepository) DeleteFileInformation(ctx context.Context, id string) (string, error) {
	result, err := r.client.FileInformation.FindUnique(
		db.FileInformation.ID.Equals(id),
	).Delete().Exec(ctx)

	if err != nil {
		return "", err
	}
	return result.ID, nil
}
