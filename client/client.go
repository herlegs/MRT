package main

import (
	"../pb"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main(){
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHWServiceClient(conn)

	resp, err := client.GetRoute(context.Background(), &pb.RouteRequest{
		From: "from",
		To: "to",
	})
	if err != nil {
		grpclog.Fatalf("could not get response: %v", err)
	}
	log.Printf("response: %s", resp.Message)
}
