package service

import (
	pb "github.com/assafvayner/grpc-files/fileservice"
)

type Server struct {
	pb.UnimplementedRetreiverServer
}
