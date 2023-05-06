package server

import (
	"context"

	pb "github.com/martzing/7solutions-assignment/3-pie-fire-dire/api"
	"github.com/martzing/7solutions-assignment/3-pie-fire-dire/service"
)

type GRPCServer struct {
	pb.UnimplementedBeefServiceServer
}

func (s *GRPCServer) BeefSummary(ctx context.Context, req *pb.BeefSummaryRequest) (*pb.BeefSummaryResponse, error) {
	beef, err := service.GetBeefSummary()
	if err != nil {
		return nil, err
	}
	response := &pb.BeefSummaryResponse{
		Beef: beef,
	}
	return response, nil
}
