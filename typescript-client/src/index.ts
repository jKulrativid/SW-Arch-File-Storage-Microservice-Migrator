import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import { ProtoGrpcType as FileStorageProtoGrpcType } from "../proto/file_storage";

const packageDefinition = protoLoader.loadSync(
  [
    "../proto/entity.proto",
    "../proto/file_storage.proto",
    "../proto/subject.proto",
  ],
  { longs: String, enums: String, defaults: true, oneofs: true }
);
const fileStorageProto = grpc.loadPackageDefinition(
  packageDefinition
) as unknown as FileStorageProtoGrpcType;

function main() {
  const client = new fileStorageProto.filestorage.FileUploadService(
    "localhost:8080",
    grpc.credentials.createInsecure()
  );

  let uploadCall = client.Upload((error, response) => {
    if (error) {
      console.error(error);
      return;
    }
    console.log(response);
  });
  uploadCall.write({
    fileName: "test_typescript_upload.txt",
    userId: "1",
    subjectId: "1",
    fileContent: "VGhpcyBmaWxlIGlzIHVwbG9hZCBmcm9tIHR5cGVzY3JpcHQgcHJvZ3JhbS4=",
  });
  uploadCall.write({
    fileContent: "VGhpcyBmaWxlIGlzIHVwbG9hZCBmcm9tIHR5cGVzY3JpcHQgcHJvZ3JhbS4=",
  });
  uploadCall.write({
    fileContent: "VGhpcyBmaWxlIGlzIHVwbG9hZCBmcm9tIHR5cGVzY3JpcHQgcHJvZ3JhbS4=",
  });
  uploadCall.end();
}

main();
