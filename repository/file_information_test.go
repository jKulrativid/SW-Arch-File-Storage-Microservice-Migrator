package repository

import (
	"context"
	"errors"
	"sort"
	"testing"

	db "github.com/PongDev/SW-Arch-File-Storage-Microservice/prisma/prisma-client"
)

func TestFileInformationRepository(t *testing.T) {
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

	var userID string
	var fileIDs []string
	var nonExistsFileID string

	fileIDs = []string{"a5245a0c-84b9-4dcd-9e67-791498274cc4", "848ed726-1a73-4f65-90a4-eaa4be449bf2"}
	sort.Strings(fileIDs)

	userID = "6073cd8a-39f3-43d2-99ed-4d5e826229c1"
	nonExistsFileID = "c566acce-8eb5-4893-81d7-e93104b288e4"

	// Test CreateFileInformation
	for _, fileID := range fileIDs {
		resultFileID, err := fileInformationRepo.CreateFileInformation(ctx, fileID, "subject_id", userID, "file_name")
		if err != nil {
			t.Fatalf("CreateFileInformation error: %v\n", err)
		}
		if resultFileID != fileID {
			t.Fatalf("CreateFileInformation error: %v expected %v\n", resultFileID, fileID)
		}
	}

	// Test GetFileInformation
	for _, fileID := range fileIDs {
		resultFileInfo, err := fileInformationRepo.GetFileInformation(ctx, fileID)
		if err != nil {
			t.Fatalf("GetFileInformation error: %v\n", err)
		}
		if resultFileInfo.ID != fileID {
			t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo.ID, fileID)
		}
		if resultFileInfo.SubjectID != "subject_id" {
			t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo.SubjectID, "subject_id")
		}
		if resultFileInfo.OwnerUserID != userID {
			t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo.OwnerUserID, userID)
		}
		if resultFileInfo.FileName != "file_name" {
			t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo.FileName, "file_name")
		}
	}
	resultFileInfo, err := fileInformationRepo.GetFileInformation(ctx, nonExistsFileID)
	if !errors.Is(err, db.ErrNotFound) {
		t.Fatalf("GetFileInformation error: %v expected %v\n", err, db.ErrNotFound)
	}
	if resultFileInfo != nil {
		t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo, nil)
	}

	// Test DeleteFileInformation
	for _, fileID := range fileIDs {
		resultFileID, err := fileInformationRepo.DeleteFileInformation(ctx, fileID)
		if err != nil {
			t.Fatalf("DeleteFileInformation error: %v\n", err)
		}
		if resultFileID != fileID {
			t.Fatalf("DeleteFileInformation error: %v expected %v\n", resultFileID, fileID)
		}
	}
}
