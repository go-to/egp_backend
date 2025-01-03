package router

import (
	"context"
	"fmt"
	"github.com/go-to/egp-backend/usecase"
	"github.com/go-to/egp-backend/usecase/input"
	"github.com/go-to/egp-protobuf/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	pb.UnimplementedEgpServiceServer
}

func New(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterEgpServiceServer(s, NewServer())

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

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetShops(ctx context.Context, req *pb.ShopsRequest) (*pb.ShopsResponse, error) {
	fmt.Println("router.GetShops")

	in := input.ShopsInput{
		ShopsRequest: req,
	}

	out, err := usecase.GetShops(&in)
	if err != nil {
		return nil, err
	}
	fmt.Println(&out)

	return &out.ShopsResponse, err
}
