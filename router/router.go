package router

import (
	"context"
	"fmt"
	"github.com/go-to/egp_backend/usecase"
	"github.com/go-to/egp_backend/usecase/input"
	"github.com/go-to/egp_protobuf/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	pb.UnimplementedEgpServiceServer
	Usecase Usecase
}

type Usecase struct {
	Shop  usecase.IShopUsecase
	Stamp usecase.IStampUsecase
}

func New(port int, u Usecase) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterEgpServiceServer(s, NewServer(u))

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

func NewServer(u Usecase) *Server {
	return &Server{Usecase: u}
}

func (s *Server) GetShops(ctx context.Context, req *pb.ShopsRequest) (*pb.ShopsResponse, error) {
	in := input.ShopsInput{
		ShopsRequest: req,
	}

	out, err := s.Usecase.Shop.GetShops(&in)
	if err != nil {
		return nil, err
	}

	return &out.ShopsResponse, nil
}

func (s *Server) GetShop(ctx context.Context, req *pb.ShopRequest) (*pb.ShopResponse, error) {
	in := input.ShopInput{
		ShopRequest: req,
	}

	out, err := s.Usecase.Shop.GetShop(&in)
	if err != nil {
		return nil, err
	}

	return &out.ShopResponse, nil
}

func (s *Server) AddStamp(ctx context.Context, req *pb.StampRequest) (*pb.StampResponse, error) {
	in := input.StampInput{
		StampRequest: req,
	}

	out, err := s.Usecase.Stamp.AddStamp(&in)
	if err != nil {
		return nil, err
	}

	return &out.StampResponse, nil
}

func (s *Server) DeleteStamp(ctx context.Context, req *pb.StampRequest) (*pb.StampResponse, error) {
	in := input.StampInput{
		StampRequest: req,
	}

	out, err := s.Usecase.Stamp.DeleteStamp(&in)
	if err != nil {
		return nil, err
	}

	return &out.StampResponse, nil
}
