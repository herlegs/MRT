package main

import (
	"context"
	"flag"
	"log"

	"github.com/herlegs/MRT/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteServiceClient(conn)

	resp, err := client.GetRoute(context.Background(), &pb.RouteRequest{
		From: "hollandvillage",
		To:   "bugis",
	})
	if err != nil {
		grpclog.Fatalf("could not get response: %v", err)
	}
	log.Printf("response: %#v", resp)
}
