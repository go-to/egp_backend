package input

import "github.com/go-to/egp_protobuf/pb"

type AddStampInput struct {
	AddStampRequest *pb.AddStampRequest
}

type DeleteStampInput struct {
	DeleteStampRequest *pb.DeleteStampRequest
}
