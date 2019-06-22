package api

import (
	"flag"

	"github.com/herlegs/MRT/pb"

	"github.com/herlegs/MRT/core"
)

type Server struct {
	// TODO use LRU if memory is a concern
	cache map[string]*pb.RouteResponse
}

var (
	MRTFile     = flag.String("mrt", "stationmap.csv", "mrt station map file")
	TrafficFile = flag.String("traffic", "traffic.json", "traffic config for different time")
)

func NewServer() *Server {
	core.LoadGraph(*MRTFile, *TrafficFile)
	return &Server{
		cache: make(map[string]*pb.RouteResponse),
	}
}
