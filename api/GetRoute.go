package api

import (
	"../pb"
	"golang.org/x/net/context"
)

func (s *Server) GetRoute(ctx context.Context, req *pb.RouteRequest) (*pb.RouteResponse, error){
	return &pb.RouteResponse{
		Message: "hello ya " + req.From + req.To,
	}, nil
}