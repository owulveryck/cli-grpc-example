package main

import (
	"log"
	"net"

	"github.com/owulveryck/cli-grpc-example/terraform-grpc/tfgrpc"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Listening on 127.0.0.1:1234")
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	tfgrpc.RegisterTerraformServer(grpcServer, &grpcCommands{Commands})
	// determine whether to use TLS
	grpcServer.Serve(listener)

}
