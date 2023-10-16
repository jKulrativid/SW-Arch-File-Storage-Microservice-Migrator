package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
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
	var subjectIDs []string
	var fileNames []string
	var nonExistsFileID string
	var fileInfos []db.FileInformationModel

	fileIDs = []string{"a5245a0c-84b9-4dcd-9e67-791498274cc4", "848ed726-1a73-4f65-90a4-eaa4be449bf2"}
	sort.Strings(fileIDs)
	subjectIDs = []string{"subject_id_1", "subject_id_2"}
	fileNames = []string{"file_1", "file_2"}
	fileInfos = []db.FileInformationModel{}

	userID = "6073cd8a-39f3-43d2-99ed-4d5e826229c1"
	nonExistsFileID = "c566acce-8eb5-4893-81d7-e93104b288e4"

	for idx, fileID := range fileIDs {
		fileInfos = append(fileInfos, db.FileInformationModel{
			InnerFileInformation: db.InnerFileInformation{
				ID:          fileID,
				SubjectID:   subjectIDs[idx],
				OwnerUserID: userID,
				FileName:    fileNames[idx],
			},
		})
	}

	// Test CreateFileInformation
	for idx, fileID := range fileIDs {
		fmt.Println(fileID)
		resultFileID, err := fileInformationRepo.CreateFileInformation(ctx, fileID, subjectIDs[idx], userID, fileNames[idx])
		if err != nil {
			t.Fatalf("CreateFileInformation error: %v\n", err)
		}
		if resultFileID != fileID {
			t.Fatalf("CreateFileInformation error: %v expected %v\n", resultFileID, fileID)
		}
	}

	defer func() {
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
	}()

	// Test GetFileInformation
	for idx, fileID := range fileIDs {
		resultFileInfo, err := fileInformationRepo.GetFileInformation(ctx, fileID)
		log.Printf("%s | %s\n", fileID, resultFileInfo.ID)
		if err != nil {
			t.Fatalf("GetFileInformation error: %v\n", err)
		}
		if resultFileInfo.ID != fileID {
			t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo.ID, fileID)
		}
		if resultFileInfo.SubjectID != subjectIDs[idx] {
			t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo.SubjectID, subjectIDs[idx])
		}
		if resultFileInfo.OwnerUserID != userID {
			t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo.OwnerUserID, userID)
		}
		if resultFileInfo.FileName != fileNames[idx] {
			t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo.FileName, fileNames[idx])
		}
	}
	resultFileInfo, err := fileInformationRepo.GetFileInformation(ctx, nonExistsFileID)
	if !errors.Is(err, db.ErrNotFound) {
		t.Fatalf("GetFileInformation error: %v expected %v\n", err, db.ErrNotFound)
	}
	if resultFileInfo != nil {
		t.Fatalf("GetFileInformation error: %v expected %v\n", resultFileInfo, nil)
	}

	// Test SearchFileInformation
	resultFileInfos, err := fileInformationRepo.SearchFileInformation(ctx, "subject_id", userID, "file_nameX")
	if err != nil || !reflect.DeepEqual(resultFileInfos, []db.FileInformationModel{}) {
		t.Fatalf("SearchFileInformation error: %v expected %v\n", resultFileInfo, []db.FileInformationModel{})
	}
	resultFileInfos, err = fileInformationRepo.SearchFileInformation(ctx, "", userID, "")
	sort.Slice(resultFileInfos, func(i, j int) bool {
		return resultFileInfos[i].ID < resultFileInfos[j].ID
	})
	if !reflect.DeepEqual(resultFileInfos, fileInfos) {
		t.Fatalf("SearchFileInformation error: %v expected %v\n", resultFileInfos, fileInfos)
	}
	resultFileInfos, err = fileInformationRepo.SearchFileInformation(ctx, "", "", "")
	sort.Slice(resultFileInfos, func(i, j int) bool {
		return resultFileInfos[i].ID < resultFileInfos[j].ID
	})
	if !reflect.DeepEqual(resultFileInfos, fileInfos) {
		t.Fatalf("SearchFileInformation error: %v expected %v\n", resultFileInfos, fileInfos)
	}
	resultFileInfos, err = fileInformationRepo.SearchFileInformation(ctx, subjectIDs[0], "", "")
	if err != nil || len(resultFileInfos) != 1 || !reflect.DeepEqual(resultFileInfos[0], fileInfos[0]) {
		t.Fatalf("SearchFileInformation error: %v expected %v\n", resultFileInfos, []db.FileInformationModel{fileInfos[0]})
	}
	resultFileInfos, err = fileInformationRepo.SearchFileInformation(ctx, "", "", fileNames[1])
	if err != nil || len(resultFileInfos) != 1 || !reflect.DeepEqual(resultFileInfos[0], fileInfos[1]) {
		t.Fatalf("SearchFileInformation error: %v expected %v\n", resultFileInfos, []db.FileInformationModel{fileInfos[1]})
	}
	resultFileInfos, err = fileInformationRepo.SearchFileInformation(ctx, subjectIDs[0], "", "file_")
	if err != nil || len(resultFileInfos) != 1 || !reflect.DeepEqual(resultFileInfos[0], fileInfos[0]) {
		t.Fatalf("SearchFileInformation error: %v expected %v\n", resultFileInfos, []db.FileInformationModel{fileInfos[0]})
	}
	resultFileInfos, err = fileInformationRepo.SearchFileInformation(ctx, "", "", "file_")
	sort.Slice(resultFileInfos, func(i, j int) bool {
		return resultFileInfos[i].ID < resultFileInfos[j].ID
	})
	if !reflect.DeepEqual(resultFileInfos, fileInfos) {
		t.Fatalf("SearchFileInformation error: %v expected %v\n", resultFileInfos, fileInfos)
	}
	resultFileInfos, err = fileInformationRepo.SearchFileInformation(ctx, "", userID, "file_")
	sort.Slice(resultFileInfos, func(i, j int) bool {
		return resultFileInfos[i].ID < resultFileInfos[j].ID
	})
	if !reflect.DeepEqual(resultFileInfos, fileInfos) {
		t.Fatalf("SearchFileInformation error: %v expected %v\n", resultFileInfos, fileInfos)
	}
}
