package main

import (
	"bytes"
	"io"
	"os"

	"golang.org/x/net/context"

	"github.com/mitchellh/cli"
	"github.com/owulveryck/cli-grpc-example/terraform-grpc/tfgrpc"
)

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

func (g *grpcCommands) Apply(ctx context.Context, in *tfgrpc.Arg) (*tfgrpc.Output, error) {
	err := os.Chdir(in.WorkingDir)
	if err != nil {
		return &tfgrpc.Output{int32(0), nil, nil}, err
	}

	ret, stdout, stderr, err := wrapper(g.commands["apply"], in.Args)
	return &tfgrpc.Output{ret, stdout, stderr}, err
}
func (g *grpcCommands) Plan(ctx context.Context, in *tfgrpc.Arg) (*tfgrpc.Output, error) {
	err := os.Chdir(in.WorkingDir)
	if err != nil {
		return &tfgrpc.Output{int32(0), nil, nil}, err
	}
	ret, stdout, stderr, err := wrapper(g.commands["plan"], in.Args)
	return &tfgrpc.Output{ret, stdout, stderr}, err
}
