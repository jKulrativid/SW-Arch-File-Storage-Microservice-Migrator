package app

import (
	"fmt"
	"net"

	"github.com/PongDev/Go-gRPC-Storage/filestorage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	*grpc.Server
}

func NewServer() *Server {
	s := grpc.NewServer()
	filestorage.RegisterFileUploadServiceServer(s, filestorage.NewFileStorageService())
	reflection.Register(s)
	return &Server{s}
}

func (s *Server) Start(port uint) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
