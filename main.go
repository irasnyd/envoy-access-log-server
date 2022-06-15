package main

import (
	"log"
	"net"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/metrics/v3"
	"github.com/irasnyd/envoy-access-log-server/pkg/sink"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	v3.RegisterMetricsServiceServer(grpcServer, sink.New())

	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening on tcp://localhost:9001")
	grpcServer.Serve(listener)
}
