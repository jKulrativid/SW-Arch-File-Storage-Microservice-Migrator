package app

import (
	"fmt"
	"net"

	"github.com/PongDev/SW-Arch-File-Storage-Microservice/filestorage"
	db "github.com/PongDev/SW-Arch-File-Storage-Microservice/prisma/prisma-client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	*grpc.Server
	prismaClient *db.PrismaClient
}

func NewServer() (*Server, error) {
	prismaClient := db.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	filestorage.RegisterFileUploadServiceServer(s, filestorage.NewFileStorageService(prismaClient))
	reflection.Register(s)
	return &Server{Server: s, prismaClient: prismaClient}, nil
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

func (s *Server) Cleanup() error {
	if err := s.prismaClient.Prisma.Disconnect(); err != nil {
		return err
	}
	return nil
}
