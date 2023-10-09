package repository

import (
	"context"
	"reflect"
	"sort"
	"testing"

	db "github.com/PongDev/SW-Arch-File-Storage-Microservice/prisma/prisma-client"
)

func TestShareFileRepository(t *testing.T) {
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
	shareFileRepo := NewShareFileRepository(prismaClient)

	userID := "6073cd8a-39f3-43d2-99ed-4d5e826229c1"
	fileID := "a5245a0c-84b9-4dcd-9e67-791498274cc4"
	nonExistsFileID := "c566acce-8eb5-4893-81d7-e93104b288e4"

	fileInformationRepo.CreateFileInformation(ctx, fileID, "subject_id", userID, "file_name")

	defer func() {
		fileInformationRepo.DeleteFileInformation(ctx, fileID)
	}()

	// Test CreateShareFile 1
	resultFileID, err := shareFileRepo.CreateShareFile(ctx, fileID, []string{"*"})
	if err != nil {
		t.Fatalf("CreateShareFile error: %v\n", err)
	}
	if resultFileID != fileID {
		t.Fatalf("CreateShareFile error: %v expected %v\n", resultFileID, fileID)
	}

	// Test GetShareFile 1
	resultUserIDs, err := shareFileRepo.GetShareFileUsers(ctx, fileID)
	if err != nil {
		t.Fatalf("GetShareFileUsers error: %v\n", err)
	}
	expectedUserIDs := []string{"*"}
	if !reflect.DeepEqual(resultUserIDs, expectedUserIDs) {
		t.Fatalf("GetShareFileUsers error: %v expected %v\n", resultUserIDs, expectedUserIDs)
	}

	// Test CreateShareFile 2
	resultFileID, err = shareFileRepo.CreateShareFile(ctx, fileID, []string{"user1", "user2"})
	if err != nil {
		t.Fatalf("CreateShareFile error: %v\n", err)
	}
	if resultFileID != fileID {
		t.Fatalf("CreateShareFile error: %v expected %v\n", resultFileID, fileID)
	}

	// Test CreateShare Not Exists File
	resultFileID, err = shareFileRepo.CreateShareFile(ctx, nonExistsFileID, []string{"user1", "user2"})
	if err == nil {
		t.Fatalf("CreateShareFile should error\n")
	}

	// Test GetShareFile 2
	resultUserIDs, err = shareFileRepo.GetShareFileUsers(ctx, fileID)
	if err != nil {
		t.Fatalf("GetShareFileUsers error: %v\n", err)
	}
	expectedUserIDs = []string{"*", "user1", "user2"}
	sort.Strings(resultUserIDs)
	sort.Strings(expectedUserIDs)
	if !reflect.DeepEqual(resultUserIDs, expectedUserIDs) {
		t.Fatalf("GetShareFileUsers error: %v expected %v\n", resultUserIDs, expectedUserIDs)
	}

	// Test CheckIsFileShareWithUser 1
	for _, userID := range []string{"*", "user1", "user2"} {
		if !shareFileRepo.CheckIsFileShareWithUser(ctx, fileID, userID) {
			t.Fatalf("CheckIsFileShareWithUser error: fileID %v should share with userID %v\n", fileID, userID)
		}
	}

	// Test DeleteShareFile 1
	shareFileRepo.DeleteShareFile(ctx, fileID, []string{"*"})

	// Test GetShareFile 3
	resultUserIDs, err = shareFileRepo.GetShareFileUsers(ctx, fileID)
	if err != nil {
		t.Fatalf("GetShareFileUsers error: %v\n", err)
	}
	expectedUserIDs = []string{"user1", "user2"}
	sort.Strings(resultUserIDs)
	sort.Strings(expectedUserIDs)
	if !reflect.DeepEqual(resultUserIDs, expectedUserIDs) {
		t.Fatalf("GetShareFileUsers error: %v expected %v\n", resultUserIDs, expectedUserIDs)
	}

	// Test CheckIsFileShareWithUser 2
	for _, userID := range []string{"user1", "user2"} {
		if !shareFileRepo.CheckIsFileShareWithUser(ctx, fileID, userID) {
			t.Fatalf("CheckIsFileShareWithUser error: fileID %v should share with userID %v\n", fileID, userID)
		}
	}
	for _, userID := range []string{"*"} {
		if shareFileRepo.CheckIsFileShareWithUser(ctx, fileID, userID) {
			t.Fatalf("CheckIsFileShareWithUser error: fileID %v shouldn't share with userID %v\n", fileID, userID)
		}
	}

	// Test DeleteShareFile 1
	resultUserIDs, err = shareFileRepo.GetShareFileUsers(ctx, fileID)
	if err != nil {
		t.Fatalf("GetShareFileUsers error: %v\n", err)
	}
	resultCount, err := shareFileRepo.DeleteShareFile(ctx, fileID, resultUserIDs)
	if err != nil {
		t.Fatalf("DeleteShareFile error: %v\n", err)
	}
	if resultCount != 2 {
		t.Fatalf("DeleteShareFile error: %v expected %v\n", resultCount, 2)
	}
}
