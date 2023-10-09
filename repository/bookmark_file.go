package repository

import (
	"context"

	db "github.com/PongDev/SW-Arch-File-Storage-Microservice/prisma/prisma-client"
)

type BookmarkFileRepository interface {
	CreateBookmark(ctx context.Context, userID string, fileID string) (string, error)
	DeleteBookmark(ctx context.Context, userID string, fileID string) (string, error)
	GetBookmark(ctx context.Context, userID string) ([]string, error)
}

type bookmarkFileRepository struct {
	client *db.PrismaClient
}

func NewBookmarkFileRepository(client *db.PrismaClient) BookmarkFileRepository {
	return &bookmarkFileRepository{client: client}
}

func (r *bookmarkFileRepository) CreateBookmark(ctx context.Context, userID string, fileID string) (string, error) {
	result, err := r.client.BookmarkFile.CreateOne(
		db.BookmarkFile.UserID.Set(userID),
		db.BookmarkFile.File.Link(db.FileInformation.ID.Equals(fileID)),
	).Exec(ctx)

	if err != nil {
		return "", err
	}
	return result.FileID, nil
}

func (r *bookmarkFileRepository) DeleteBookmark(ctx context.Context, userID string, fileID string) (string, error) {
	result, err := r.client.BookmarkFile.FindUnique(
		db.BookmarkFile.UserIDFileID(db.BookmarkFile.UserID.Equals(userID), db.BookmarkFile.FileID.Equals(fileID)),
	).Delete().Exec(ctx)

	if err != nil {
		return "", err
	}
	return result.FileID, nil
}

func (r *bookmarkFileRepository) GetBookmark(ctx context.Context, userID string) ([]string, error) {
	result, err := r.client.BookmarkFile.FindMany(
		db.BookmarkFile.UserID.Equals((userID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	fileIDs := make([]string, len(result))
	for i, bookmark := range result {
		fileIDs[i] = bookmark.FileID
	}

	return fileIDs, nil
}
