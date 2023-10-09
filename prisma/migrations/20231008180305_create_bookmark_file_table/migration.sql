-- CreateTable
CREATE TABLE "BookmarkFile" (
    "userID" TEXT NOT NULL,
    "fileID" TEXT NOT NULL,

    CONSTRAINT "BookmarkFile_pkey" PRIMARY KEY ("userID","fileID")
);

-- AddForeignKey
ALTER TABLE "BookmarkFile" ADD CONSTRAINT "BookmarkFile_fileID_fkey" FOREIGN KEY ("fileID") REFERENCES "FileInformation"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
