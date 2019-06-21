package main

import (
	"./server"
)

//go:generate protoc -I/usr/local/include -I./pb -I$PROTO_PATH/include -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src --go_out=plugins=grpc:./pb --grpc-gateway_out=logtostderr=true:./pb --swagger_out=logtostderr=true:./pb ./pb/route.proto

func main() {
	go server.StartRPC()
	server.StartHTTP()
}
