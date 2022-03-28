package service

import (
	"context"
	"os"
	"strings"

	pb "github.com/assafvayner/grpc-files/fileservice"
	"google.golang.org/genproto/googleapis/rpc/status"
)

// generates error response for remove request
func GenRemoveErrorResponse(code int32, message string) *pb.RemoveResponse {
	return &pb.RemoveResponse{
		Status: &status.Status{
			Code:    code,
			Message: message,
		},
	}
}

func (*Server) Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	// weakly prevent directory traversal attack
	if strings.Contains(req.GetPath(), "..") {
		return GenRemoveErrorResponse(1, "Illegal resource name"), nil
	}

	// compose resource path (data directory)
	fullpath := "data/" + req.GetPath()

	err := os.RemoveAll(fullpath)
	if err != nil {
		return GenRemoveErrorResponse(1, "could not file/directory, it may not exist"), nil
	}

	res := pb.RemoveResponse{
		Status: &status.Status{ // generates error response for fetch request
			Code:    0,
			Message: "remove file succeeded",
		},
	}

	return &res, nil
}
