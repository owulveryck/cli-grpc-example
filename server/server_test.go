package main

import (
	"bytes"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/owulveryck/cli-grpc-example/myservice"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	c := cli.NewCLI("server", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"hello": func() (cli.Command, error) {
			return &HelloCommand{}, nil
		},
		"goodbye": func() (cli.Command, error) {
			return &GoodbyeCommand{}, nil
		},
	}

	go func() {
		listener, err := net.Listen("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		myservice.RegisterMyServiceServer(grpcServer, &grpcCommands{c.Commands})
		// determine whether to use TLS
		grpcServer.Serve(listener)
	}()
	os.Exit(m.Run())
}

func BenchmarkHello(b *testing.B) {
	for n := 0; n < b.N; n++ {

		hello := strconv.Itoa(n)
		conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
		if err != nil {
			b.Fatal("Cannot reach grpc server", err)
		}
		defer conn.Close()
		client := myservice.NewMyServiceClient(conn)
		output, err := client.Hello(context.Background(), &myservice.Arg{[]string{hello}})
		stdout := bytes.NewBuffer(output.Stdout)
		if stdout.String() != "hello ["+hello+"]\n" {
			b.Fatalf("%v != %v", stdout.String(), "hello ["+hello+"]\n")
		}
	}
}
