package repository

import (
	"context"

	db "github.com/PongDev/SW-Arch-File-Storage-Microservice/prisma/prisma-client"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
)

type ShareFileRepository interface {
	GetShareFileUsers(ctx context.Context, fileID string) ([]string, error)
	CreateShareFile(ctx context.Context, fileID string, shareWithUserIDs []string) (string, error)
	DeleteShareFile(ctx context.Context, fileID string, shareWithUserIDs []string) (int, error)
	CheckIsFileShareWithUser(ctx context.Context, fileID string, userID string) bool
}

type shareFileRepository struct {
	client *db.PrismaClient
}

func NewShareFileRepository(client *db.PrismaClient) ShareFileRepository {
	return &shareFileRepository{client: client}
}

func (r *shareFileRepository) GetShareFileUsers(ctx context.Context, fileID string) ([]string, error) {
	result, err := r.client.ShareFile.FindMany(db.ShareFile.FileID.Equals(fileID)).Exec(ctx)
	if err != nil {
		return []string{}, err
	}

	userIDs := make([]string, 0)
	for _, shareFile := range result {
		userIDs = append(userIDs, shareFile.UserID)
	}
	return userIDs, nil
}

func (r *shareFileRepository) CreateShareFile(ctx context.Context, fileID string, shareWithUserIDs []string) (string, error) {
	result, err := r.client.FileInformation.FindUnique(
		db.FileInformation.ID.Equals(fileID),
	).Exec(ctx)

	if err != nil {
		return "", err
	}

	transactions := []transaction.Param{}
	for _, shareWithUserID := range shareWithUserIDs {
		transactions = append(transactions, r.client.ShareFile.CreateOne(
			db.ShareFile.UserID.Set(shareWithUserID),
			db.ShareFile.File.Link(db.FileInformation.ID.Equals(result.ID)),
		).Tx())
	}
	if err := r.client.Prisma.Transaction(transactions...).Exec(ctx); err != nil {
		return "", err
	}
	return result.ID, nil
}

func (r *shareFileRepository) DeleteShareFile(ctx context.Context, fileID string, shareWithUserIDs []string) (int, error) {
	result, err := r.client.ShareFile.FindMany(db.ShareFile.FileID.Equals(fileID), db.ShareFile.UserID.In(shareWithUserIDs)).Delete().Exec(ctx)
	if err != nil {
		return 0, err
	}
	return result.Count, nil
}

func (r *shareFileRepository) CheckIsFileShareWithUser(ctx context.Context, fileID string, userID string) bool {
	_, err := r.client.ShareFile.FindUnique(db.ShareFile.UserIDFileID(
		db.ShareFile.UserID.Equals(userID),
		db.ShareFile.FileID.Equals(fileID),
	)).Exec(ctx)

	if err != nil {
		return false
	}
	return true
}
