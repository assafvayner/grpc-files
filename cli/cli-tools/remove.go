package cli

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/assafvayner/grpc-files/fileservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Remove(resource string) {
	conn, err := grpc.Dial("localhost:8808", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error on grpc dial", err)
	}
	defer conn.Close()

	client := pb.NewRetreiverClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := pb.RemoveRequest{
		Path: resource,
	}

	res, err := client.Remove(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	if res.GetStatus().GetCode() != 0 {
		log.Fatal("Error from server:\n\t" + res.GetStatus().GetMessage())
	}
	log.Printf("Successfully removed: %s\n", resource)
	os.Exit(0)
}
