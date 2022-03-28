package main

import (
	"fmt"
	"net"
	"os"

	pb "github.com/assafvayner/grpc-files/fileservice"
	"github.com/assafvayner/grpc-files/server/service"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("start of program")

	lis, err := net.Listen("tcp", "localhost:8808")

	if err != nil {
		os.Exit(1)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	retrieverServer := service.Server{}

	pb.RegisterRetreiverServer(grpcServer, &retrieverServer)
	fmt.Println("calling serve")
	grpcServer.Serve(lis)

	// lis.Close()

	// time.Sleep(5)

	fmt.Println("end of program")
}
