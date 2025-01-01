package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/go-to/egp-protobuf/pb"
)

type EgpServer struct {
	pb.UnimplementedEgpServiceServer
}

func NewEgpServer() *EgpServer {
	return &EgpServer{}
}

func (s *EgpServer) GetShops(ctx context.Context, req *pb.ShopsRequest) (*pb.ShopsResponse, error) {
	fmt.Println("GetShops")

	var shops []*pb.Shop
	for i := 1; i <= 10; i++ {
		shops = append(shops, &pb.Shop{
			ID: int64(i),
		})
	}

	return &pb.ShopsResponse{
		Shops: shops,
	}, nil
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterEgpServiceServer(s, NewEgpServer())

	reflection.Register(s)

	go func() {
		log.Printf("grpc server listening on port %d", port)
		err := s.Serve(listener)
		if err != nil {
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutting down server...")
	s.GracefulStop()
}
