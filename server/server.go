package server

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/herlegs/MRT/api"
	"github.com/herlegs/MRT/pb"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	GRPCPort = flag.Int("grpc", 10000, "grpc server port")
	HTTPPort = flag.Int("http", 8080, "http server port")
)

func StartRPC() {
	flag.Parse()
	fmt.Println("running grpc at port", *GRPCPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *GRPCPort))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRouteServiceServer(s, api.NewServer())
	if err := s.Serve(lis); err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}

func StartHTTP() {
	flag.Parse()
	fmt.Println("running http at port", *HTTPPort)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterRouteServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", *GRPCPort), opts)
	if err != nil {
		grpclog.Fatalf("failed to register service: %v", err)
	}
	http.ListenAndServe(fmt.Sprintf(":%d", *HTTPPort), mux)
}
