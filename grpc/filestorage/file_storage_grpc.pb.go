// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: file_storage.proto

package filestorage

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileUploadServiceClient is the client API for FileUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileUploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileUploadService_UploadClient, error)
	Download(ctx context.Context, in *FileDownloadRequest, opts ...grpc.CallOption) (FileUploadService_DownloadClient, error)
	Delete(ctx context.Context, in *FileDeleteRequest, opts ...grpc.CallOption) (*FileDeleteResponse, error)
	ShareFile(ctx context.Context, in *ShareFileRequest, opts ...grpc.CallOption) (*ShareFileResponse, error)
	CreateBookmarkFile(ctx context.Context, in *CreateBookmarkFileRequest, opts ...grpc.CallOption) (*CreateBookmarkFileResponse, error)
	DeleteBookmarkFile(ctx context.Context, in *DeleteBookmarkFileRequest, opts ...grpc.CallOption) (*DeleteBookmarkFileResponse, error)
	GetBookmarkFiles(ctx context.Context, in *GetBookmarkFilesRequest, opts ...grpc.CallOption) (*GetBookmarkFilesResponse, error)
	SearchFile(ctx context.Context, in *SearchFileRequest, opts ...grpc.CallOption) (*SearchFileResponse, error)
}

type fileUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileUploadServiceClient(cc grpc.ClientConnInterface) FileUploadServiceClient {
	return &fileUploadServiceClient{cc}
}

func (c *fileUploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileUploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileUploadService_ServiceDesc.Streams[0], "/filestorage.FileUploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileUploadServiceUploadClient{stream}
	return x, nil
}

type FileUploadService_UploadClient interface {
	Send(*FileUploadRequest) error
	CloseAndRecv() (*FileUploadResponse, error)
	grpc.ClientStream
}

type fileUploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *fileUploadServiceUploadClient) Send(m *FileUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileUploadServiceUploadClient) CloseAndRecv() (*FileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileUploadServiceClient) Download(ctx context.Context, in *FileDownloadRequest, opts ...grpc.CallOption) (FileUploadService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileUploadService_ServiceDesc.Streams[1], "/filestorage.FileUploadService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileUploadServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileUploadService_DownloadClient interface {
	Recv() (*FileDownloadResponse, error)
	grpc.ClientStream
}

type fileUploadServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *fileUploadServiceDownloadClient) Recv() (*FileDownloadResponse, error) {
	m := new(FileDownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileUploadServiceClient) Delete(ctx context.Context, in *FileDeleteRequest, opts ...grpc.CallOption) (*FileDeleteResponse, error) {
	out := new(FileDeleteResponse)
	err := c.cc.Invoke(ctx, "/filestorage.FileUploadService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileUploadServiceClient) ShareFile(ctx context.Context, in *ShareFileRequest, opts ...grpc.CallOption) (*ShareFileResponse, error) {
	out := new(ShareFileResponse)
	err := c.cc.Invoke(ctx, "/filestorage.FileUploadService/ShareFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileUploadServiceClient) CreateBookmarkFile(ctx context.Context, in *CreateBookmarkFileRequest, opts ...grpc.CallOption) (*CreateBookmarkFileResponse, error) {
	out := new(CreateBookmarkFileResponse)
	err := c.cc.Invoke(ctx, "/filestorage.FileUploadService/CreateBookmarkFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileUploadServiceClient) DeleteBookmarkFile(ctx context.Context, in *DeleteBookmarkFileRequest, opts ...grpc.CallOption) (*DeleteBookmarkFileResponse, error) {
	out := new(DeleteBookmarkFileResponse)
	err := c.cc.Invoke(ctx, "/filestorage.FileUploadService/DeleteBookmarkFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileUploadServiceClient) GetBookmarkFiles(ctx context.Context, in *GetBookmarkFilesRequest, opts ...grpc.CallOption) (*GetBookmarkFilesResponse, error) {
	out := new(GetBookmarkFilesResponse)
	err := c.cc.Invoke(ctx, "/filestorage.FileUploadService/GetBookmarkFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileUploadServiceClient) SearchFile(ctx context.Context, in *SearchFileRequest, opts ...grpc.CallOption) (*SearchFileResponse, error) {
	out := new(SearchFileResponse)
	err := c.cc.Invoke(ctx, "/filestorage.FileUploadService/SearchFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileUploadServiceServer is the server API for FileUploadService service.
// All implementations must embed UnimplementedFileUploadServiceServer
// for forward compatibility
type FileUploadServiceServer interface {
	Upload(FileUploadService_UploadServer) error
	Download(*FileDownloadRequest, FileUploadService_DownloadServer) error
	Delete(context.Context, *FileDeleteRequest) (*FileDeleteResponse, error)
	ShareFile(context.Context, *ShareFileRequest) (*ShareFileResponse, error)
	CreateBookmarkFile(context.Context, *CreateBookmarkFileRequest) (*CreateBookmarkFileResponse, error)
	DeleteBookmarkFile(context.Context, *DeleteBookmarkFileRequest) (*DeleteBookmarkFileResponse, error)
	GetBookmarkFiles(context.Context, *GetBookmarkFilesRequest) (*GetBookmarkFilesResponse, error)
	SearchFile(context.Context, *SearchFileRequest) (*SearchFileResponse, error)
	mustEmbedUnimplementedFileUploadServiceServer()
}

// UnimplementedFileUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileUploadServiceServer struct {
}

func (UnimplementedFileUploadServiceServer) Upload(FileUploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFileUploadServiceServer) Download(*FileDownloadRequest, FileUploadService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedFileUploadServiceServer) Delete(context.Context, *FileDeleteRequest) (*FileDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFileUploadServiceServer) ShareFile(context.Context, *ShareFileRequest) (*ShareFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareFile not implemented")
}
func (UnimplementedFileUploadServiceServer) CreateBookmarkFile(context.Context, *CreateBookmarkFileRequest) (*CreateBookmarkFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookmarkFile not implemented")
}
func (UnimplementedFileUploadServiceServer) DeleteBookmarkFile(context.Context, *DeleteBookmarkFileRequest) (*DeleteBookmarkFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookmarkFile not implemented")
}
func (UnimplementedFileUploadServiceServer) GetBookmarkFiles(context.Context, *GetBookmarkFilesRequest) (*GetBookmarkFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookmarkFiles not implemented")
}
func (UnimplementedFileUploadServiceServer) SearchFile(context.Context, *SearchFileRequest) (*SearchFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFile not implemented")
}
func (UnimplementedFileUploadServiceServer) mustEmbedUnimplementedFileUploadServiceServer() {}

// UnsafeFileUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileUploadServiceServer will
// result in compilation errors.
type UnsafeFileUploadServiceServer interface {
	mustEmbedUnimplementedFileUploadServiceServer()
}

func RegisterFileUploadServiceServer(s grpc.ServiceRegistrar, srv FileUploadServiceServer) {
	s.RegisterService(&FileUploadService_ServiceDesc, srv)
}

func _FileUploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileUploadServiceServer).Upload(&fileUploadServiceUploadServer{stream})
}

type FileUploadService_UploadServer interface {
	SendAndClose(*FileUploadResponse) error
	Recv() (*FileUploadRequest, error)
	grpc.ServerStream
}

type fileUploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *fileUploadServiceUploadServer) SendAndClose(m *FileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileUploadServiceUploadServer) Recv() (*FileUploadRequest, error) {
	m := new(FileUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileUploadService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileDownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileUploadServiceServer).Download(m, &fileUploadServiceDownloadServer{stream})
}

type FileUploadService_DownloadServer interface {
	Send(*FileDownloadResponse) error
	grpc.ServerStream
}

type fileUploadServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *fileUploadServiceDownloadServer) Send(m *FileDownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileUploadService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileUploadServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filestorage.FileUploadService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileUploadServiceServer).Delete(ctx, req.(*FileDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileUploadService_ShareFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileUploadServiceServer).ShareFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filestorage.FileUploadService/ShareFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileUploadServiceServer).ShareFile(ctx, req.(*ShareFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileUploadService_CreateBookmarkFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookmarkFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileUploadServiceServer).CreateBookmarkFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filestorage.FileUploadService/CreateBookmarkFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileUploadServiceServer).CreateBookmarkFile(ctx, req.(*CreateBookmarkFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileUploadService_DeleteBookmarkFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookmarkFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileUploadServiceServer).DeleteBookmarkFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filestorage.FileUploadService/DeleteBookmarkFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileUploadServiceServer).DeleteBookmarkFile(ctx, req.(*DeleteBookmarkFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileUploadService_GetBookmarkFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookmarkFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileUploadServiceServer).GetBookmarkFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filestorage.FileUploadService/GetBookmarkFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileUploadServiceServer).GetBookmarkFiles(ctx, req.(*GetBookmarkFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileUploadService_SearchFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileUploadServiceServer).SearchFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filestorage.FileUploadService/SearchFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileUploadServiceServer).SearchFile(ctx, req.(*SearchFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileUploadService_ServiceDesc is the grpc.ServiceDesc for FileUploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileUploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filestorage.FileUploadService",
	HandlerType: (*FileUploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _FileUploadService_Delete_Handler,
		},
		{
			MethodName: "ShareFile",
			Handler:    _FileUploadService_ShareFile_Handler,
		},
		{
			MethodName: "CreateBookmarkFile",
			Handler:    _FileUploadService_CreateBookmarkFile_Handler,
		},
		{
			MethodName: "DeleteBookmarkFile",
			Handler:    _FileUploadService_DeleteBookmarkFile_Handler,
		},
		{
			MethodName: "GetBookmarkFiles",
			Handler:    _FileUploadService_GetBookmarkFiles_Handler,
		},
		{
			MethodName: "SearchFile",
			Handler:    _FileUploadService_SearchFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileUploadService_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _FileUploadService_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file_storage.proto",
}
