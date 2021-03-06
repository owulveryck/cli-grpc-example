package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/mitchellh/cli"
	"github.com/owulveryck/cli-grpc-example/myservice"
)

// GoodbyeCommand ...
type GoodbyeCommand struct{}

// Help ...
func (t *GoodbyeCommand) Help() string {
	return "help"
}

// Run ...
func (t *GoodbyeCommand) Run(args []string) int {
	log.Println("goodbye", args)
	fmt.Fprintf(os.Stderr, "goodbye %v", args)
	return 0
}

// Synopsis ...
func (t *GoodbyeCommand) Synopsis() string {
	return "synopsis..."
}

// HelloCommand ...
type HelloCommand struct{}

// Help ...
func (t *HelloCommand) Help() string {
	return "hello [arg0] [arg1] ... says hello to everyone"
}

// Run ...
func (t *HelloCommand) Run(args []string) int {
	fmt.Println("hello", args)
	return 0
}

// Synopsis ...
func (t *HelloCommand) Synopsis() string {
	return "A sample command that says hello on stdout"
}

func wrapper(cf cli.CommandFactory, args []string) (int32, []byte, []byte, error) {
	var ret int32
	oldStdout := os.Stdout // keep backup of the real stdout
	oldStderr := os.Stderr

	// Backup the stdout
	r, w, err := os.Pipe()
	if err != nil {
		return ret, nil, nil, err
	}
	re, we, err := os.Pipe()
	if err != nil {
		return ret, nil, nil, err
	}
	os.Stdout = w
	os.Stderr = we

	runner, err := cf()
	if err != nil {
		return ret, nil, nil, err
	}
	ret = int32(runner.Run(args))

	outC := make(chan []byte)
	errC := make(chan []byte)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.Bytes()
	}()
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, re)
		errC <- buf.Bytes()
	}()

	// back to normal state
	w.Close()
	we.Close()
	os.Stdout = oldStdout // restoring the real stdout
	os.Stderr = oldStderr
	stdout := <-outC
	stderr := <-errC
	return ret, stdout, stderr, nil
}

type grpcCommands struct {
	commands map[string]cli.CommandFactory
}

func (g *grpcCommands) Hello(ctx context.Context, in *myservice.Arg) (*myservice.Output, error) {
	ret, stdout, stderr, err := wrapper(g.commands["hello"], in.Args)
	return &myservice.Output{ret, stdout, stderr}, err
}
func (g *grpcCommands) Goodbye(ctx context.Context, in *myservice.Arg) (*myservice.Output, error) {
	ret, stdout, stderr, err := wrapper(g.commands["goodbye"], in.Args)
	return &myservice.Output{ret, stdout, stderr}, err
}

func main() {
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

	if len(c.Args) == 0 {
		log.Println("Listening on 127.0.0.1:1234")
		listener, err := net.Listen("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		myservice.RegisterMyServiceServer(grpcServer, &grpcCommands{c.Commands})
		// determine whether to use TLS
		grpcServer.Serve(listener)

	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
