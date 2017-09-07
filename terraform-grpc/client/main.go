package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"

	"github.com/owulveryck/cli-grpc-example/terraform-grpc/tfgrpc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot reach grpc server", err)
	}
	defer conn.Close()
	if len(os.Args) < 2 {
		log.Fatal("Wrong numbers of arguments")
	}
	client := tfgrpc.NewTerraformClient(conn)
	output := &tfgrpc.Output{}
	switch os.Args[1] {
	case "plan":
		output, err = client.Plan(context.Background(), &tfgrpc.Arg{
			os.Args[2],
			os.Args[3:],
		})
	case "apply":
		output, err = client.Apply(context.Background(), &tfgrpc.Arg{
			os.Args[2],
			os.Args[3:],
		})
	default:
		log.Fatal("Unknown command")
	}
	stdout := bytes.NewBuffer(output.Stdout)
	stderr := bytes.NewBuffer(output.Stderr)
	io.Copy(os.Stdout, stdout)
	io.Copy(os.Stderr, stderr)
}
