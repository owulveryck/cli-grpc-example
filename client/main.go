package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/owulveryck/cli-grpc-example/myservice"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot reach grpc server", err)
	}
	defer conn.Close()
	client := myservice.NewMyServiceClient(conn)
	output, err := client.Hello(context.Background(), &myservice.Arg{os.Args[1]})
	fmt.Println(output)
	output, err = client.Goodbye(context.Background(), &myservice.Arg{os.Args[1]})
	fmt.Println(output)
}
