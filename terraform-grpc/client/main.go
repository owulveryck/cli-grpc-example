package main

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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
	case "push":

		// Create a buffer to write our archive to.
		buf := new(bytes.Buffer)

		// Create a new zip archive.
		archive := zip.NewWriter(buf)

		source := "."
		info, err := os.Stat(source)
		if err != nil {
			log.Fatal(err)
		}

		var baseDir string
		if info.IsDir() {
			baseDir = filepath.Base(source)
		}

		filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if info.IsDir() {
				header.Name += "/"
			} else {
				ok, err := regexp.MatchString(".tf$", info.Name())
				if err != nil {
					return err
				}
				if !ok {
					return nil
				}

				header.Method = zip.Deflate
			}

			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			return err
		})
		// Make sure to check the error on Close.
		err = archive.Close()
		if err != nil {
			log.Fatal("Cannot close zip writer", err)
		}
		pushClient, err := client.Push(context.Background(), grpc.MaxCallRecvMsgSize(65536))
		if err != nil {
			log.Fatal("Cannot create grpc push client", err)
		}
		err = pushClient.Send(&tfgrpc.Body{
			Zipfile: buf.Bytes(),
		})
		if err != nil {
			log.Fatal("Send error", err)
		}
		id, err := pushClient.CloseAndRecv()
		if err != nil {
			log.Fatal("Received returned an error", err)
		}
		fmt.Println(id.Tmpdir)

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
