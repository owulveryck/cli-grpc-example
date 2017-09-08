package main

import (
	"log"
	"net"

	"github.com/hashicorp/terraform/command"
	"github.com/owulveryck/cli-grpc-example/terraform-grpc/tfgrpc"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Listening on 127.0.0.1:1234")
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(16500545))
	// PluginOverrides are paths that override discovered plugins, set from
	// the config file.
	var PluginOverrides command.PluginOverrides

	meta := command.Meta{
		Color:            false,
		GlobalPluginDirs: globalPluginDirs(),
		PluginOverrides:  &PluginOverrides,
		Ui:               &grpcUI{},
	}

	tfgrpc.RegisterTerraformServer(grpcServer, &grpcCommands{meta: meta})
	// determine whether to use TLS
	grpcServer.Serve(listener)

}
