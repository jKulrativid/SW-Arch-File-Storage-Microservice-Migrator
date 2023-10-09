package repository

import (
	"context"
	"errors"
	"reflect"
	"sort"
	"testing"

	db "github.com/PongDev/SW-Arch-File-Storage-Microservice/prisma/prisma-client"
)

func TestBookmarkFileRepository(t *testing.T) {
	prismaClient := db.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		t.Fatalf("Prisma connected error: %v\n", err)
		return
	}
	defer func() {
		if err := prismaClient.Prisma.Disconnect(); err != nil {
			t.Fatalf("Prisma disconnected error: %v\n", err)
			return
		}
	}()

	ctx := context.Background()

	fileInformationRepo := NewFileInformationRepository(prismaClient)
	bookmarkRepo := NewBookmarkFileRepository(prismaClient)

	var userID string
	var fileID string
	var fileIDs []string
	var nonExistsFileID string

	fileIDs = []string{"a5245a0c-84b9-4dcd-9e67-791498274cc4", "848ed726-1a73-4f65-90a4-eaa4be449bf2"}
	nonUserOwnFileIDs := []string{"f344e208-594d-4859-bbcf-23443ab10c68"}
	sort.Strings(fileIDs)

	userID = "6073cd8a-39f3-43d2-99ed-4d5e826229c1"
	nonExistsFileID = "c566acce-8eb5-4893-81d7-e93104b288e4"

	for _, fileID := range fileIDs {
		fileInformationRepo.CreateFileInformation(ctx, fileID, "subject_id", userID, "file_name")
	}
	for _, fileID := range nonUserOwnFileIDs {
		fileInformationRepo.CreateFileInformation(ctx, fileID, "subject_id", "non_user_id", "file_name")
	}

	defer func() {
		for _, fileID := range fileIDs {
			fileInformationRepo.DeleteFileInformation(ctx, fileID)
		}
		for _, fileID := range nonUserOwnFileIDs {
			fileInformationRepo.DeleteFileInformation(ctx, fileID)
		}
	}()

	// Test CreateBookmark 1
	fileID = fileIDs[0]
	resultFileID, err := bookmarkRepo.CreateBookmark(ctx, userID, fileID)
	if err != nil {
		t.Fatalf("CreateBookmark error: %v\n", err)
	}
	if resultFileID != fileID {
		t.Fatalf("CreateBookmark not equal: %v expected %v\n", resultFileID, fileID)
	}

	// Test GetBookmark 1
	resultFileIDs, err := bookmarkRepo.GetBookmark(ctx, userID)
	if err != nil {
		t.Fatalf("GetBookmark error: %v\n", err)
	}

	if !reflect.DeepEqual(resultFileIDs, []string{fileID}) {
		t.Fatalf("GetBookmark not equal: %v expected %v\n", resultFileIDs, []string{fileID})
	}

	// Test CreateBookmark 2
	fileID = fileIDs[1]
	resultFileID, err = bookmarkRepo.CreateBookmark(ctx, userID, fileID)
	if err != nil {
		t.Fatalf("CreateBookmark error: %v\n", err)
	}
	if resultFileID != fileID {
		t.Fatalf("CreateBookmark not equal: %v expected %v\n", resultFileID, fileID)
	}

	// Test GetBookmark 2
	resultFileIDs, err = bookmarkRepo.GetBookmark(ctx, userID)
	if err != nil {
		t.Fatalf("GetBookmark error: %v\n", err)
	}

	sort.Strings(resultFileIDs)
	if !reflect.DeepEqual(resultFileIDs, fileIDs) {
		t.Fatalf("GetBookmark not equal: %v expected %v\n", resultFileIDs, []string{fileID})
	}

	// Test DeleteBookmark
	for _, v := range fileIDs {
		resultFileID, err = bookmarkRepo.DeleteBookmark(ctx, userID, v)
		if err != nil {
			t.Fatalf("DeleteBookmark error: %v\n", err)
		}

		if resultFileID != v {
			t.Fatalf("DeleteBookmark not equal: %v expected %v\n", resultFileID, v)
		}
	}

	// Test DeleteBookmark: Non Exists Bookmark
	resultFileID, err = bookmarkRepo.DeleteBookmark(ctx, userID, nonExistsFileID)
	if !errors.Is(err, db.ErrNotFound) {
		t.Fatalf("DeleteBookmark not equal: %v expected %v\n", err, db.ErrNotFound)
	}

	// Test CreateBookmark: Non Exist File
	fileID = nonExistsFileID
	resultFileID, err = bookmarkRepo.CreateBookmark(ctx, userID, fileID)
	if err == nil {
		t.Fatalf("CreateBookmark should error\n")
	}
}
