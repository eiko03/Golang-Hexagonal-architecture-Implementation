package grpc

import (
	"google.golang.org/grpc"
	"hexArch/internal/adapters/framework/left/grpc/pb"
	"hexArch/internal/ports"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc server %v", err)
	}
}
