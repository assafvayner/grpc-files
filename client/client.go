package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/assafvayner/grpc-files/fileservice"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8808", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error on grpc dial", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewRetreiverClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := pb.FetchRequest{
		Path: "temp",
	}

	res, err := client.Fetch(ctx, &req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("status:\n\tcode: %d\n\tmessage: %s\n", res.GetStatus().GetCode(), res.GetStatus().GetMessage())

	fmt.Println(string(res.GetFile().GetData()))

	fmt.Println(res.GetFile())

	req = pb.FetchRequest{
		Path:         "wordle-suggester/main.cc",
		MetadataOnly: true,
	}

	res, err = client.Fetch(ctx, &req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("status:\n\tcode: %d\n\tmessage: %s\n", res.GetStatus().GetCode(), res.GetStatus().GetMessage())

	pushreq := pb.PushRequest{
		Path: "",
		Resource: &pb.PushRequest_File{
			File: &pb.File{
				Name: "pushed.txt",
				Data: []byte("the contents of pushed.txt"),
			},
		},
	}

	pushres, err := client.Push(ctx, &pushreq)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("status:\n\tcode: %d\n\tmessage: %s\n", pushres.GetStatus().GetCode(), pushres.GetStatus().GetMessage())

	removereq := pb.RemoveRequest{
		Path: "pushed.txt",
	}

	removeres, err := client.Remove(ctx, &removereq)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("status:\n\tcode: %d\n\tmessage: %s\n", removeres.GetStatus().GetCode(), removeres.GetStatus().GetMessage())

	dirreq := pb.FetchRequest{
		Path:         "wordle-suggester",
		MetadataOnly: true,
	}

	dirres, err := client.Fetch(ctx, &dirreq)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("status:\n\tcode: %d\n\tmessage: %s\n", dirres.GetStatus().GetCode(), dirres.GetStatus().GetMessage())

}
