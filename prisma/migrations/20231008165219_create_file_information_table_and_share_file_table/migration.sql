-- CreateTable
CREATE TABLE "FileInformation" (
    "id" TEXT NOT NULL,
    "subjectID" TEXT NOT NULL,
    "ownerUserID" TEXT NOT NULL,
    "fileName" TEXT NOT NULL,

    CONSTRAINT "FileInformation_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "ShareFile" (
    "userID" TEXT NOT NULL,
    "fileID" TEXT NOT NULL,

    CONSTRAINT "ShareFile_pkey" PRIMARY KEY ("userID","fileID")
);

-- AddForeignKey
ALTER TABLE "ShareFile" ADD CONSTRAINT "ShareFile_fileID_fkey" FOREIGN KEY ("fileID") REFERENCES "FileInformation"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
