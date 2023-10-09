package app

import (
	"fmt"
	"net"

	"github.com/PongDev/SW-Arch-File-Storage-Microservice/config"
	"github.com/PongDev/SW-Arch-File-Storage-Microservice/filestorage"
	"github.com/PongDev/SW-Arch-File-Storage-Microservice/grpc/subject"
	db "github.com/PongDev/SW-Arch-File-Storage-Microservice/prisma/prisma-client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	*grpc.Server
	prismaClient       *db.PrismaClient
	grpcSubjectService *grpc.ClientConn
}

func NewServer() (*Server, error) {
	prismaClient := db.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		return nil, err
	}

	grpcSubjectService, err := grpc.Dial(
		config.Env.SUBJECT_MICROSERVICE_ENDPOINT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	subjectServiceClient := subject.NewSubjectServiceClient(grpcSubjectService)

	s := grpc.NewServer()
	filestorage.RegisterFileUploadServiceServer(s, filestorage.NewFileStorageService(prismaClient, subjectServiceClient))
	reflection.Register(s)
	return &Server{Server: s, prismaClient: prismaClient, grpcSubjectService: grpcSubjectService}, nil
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
	if err := s.grpcSubjectService.Close(); err != nil {
		return err
	}
	return nil
}
