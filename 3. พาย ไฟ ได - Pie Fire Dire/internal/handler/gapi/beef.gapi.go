package gapi

import (
	pb "beef/internal/pd"
	"beef/internal/service"
	"context"
)

type BeefHandlerGrpc struct {
	pb.UnimplementedBeefServiceServer
	beef service.BeefService
}

func NewBeefHandlerGrpcHandler(beef service.BeefService) *BeefHandlerGrpc {
	beefServer := BeefHandlerGrpc{
		beef: beef,
	}

	return &beefServer
}

func (b *BeefHandlerGrpc) FindBeef(ctx context.Context, req *pb.FindBeefRequest) (*pb.FindBeefResponse, error) {

	beefRes, err := b.beef.FindBeef()

	if err != nil {
		return nil, err
	}

	res := &pb.FindBeefResponse{
		Beef: beefRes.Beef,
	}

	return res, nil
}
