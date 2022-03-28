package service

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	pb "github.com/assafvayner/grpc-files/fileservice"
	"google.golang.org/genproto/googleapis/rpc/status"
)

// generates error response for push request
func GenPushErrorResponse(code int32, message string) *pb.PushResponse {
	return &pb.PushResponse{
		Status: &status.Status{
			Code:    code,
			Message: message,
		},
	}
}

func (*Server) Push(ctx context.Context, req *pb.PushRequest) (*pb.PushResponse, error) {
	// weakly prevent directory traversal attack
	if strings.Contains(req.GetPath(), "..") {
		return GenPushErrorResponse(1, "Illegal resource name"), nil
	}

	var dir *pb.Directory
	var file *pb.File
	var err error

	file = req.GetFile()
	if file != nil {
		err = AddFile(req.GetPath(), file)
		goto RESPOND
	}
	dir = req.GetDirectory()
	if dir == nil {
		return GenPushErrorResponse(1, "improper resource"), nil
	}
	err = AddDirectory(req.GetPath(), dir)

RESPOND:
	if err != nil {
		return GenPushErrorResponse(1, err.Error()), nil
	}
	res := pb.PushResponse{
		Status: &status.Status{
			Code:    0,
			Message: "push succeeded",
		},
	}

	return &res, nil
}

func AddDirectory(path string, directory *pb.Directory) error {
	log.Println(directory.GetName())
	if strings.Contains(directory.GetName(), "/") || strings.Contains(directory.GetName(), "..") {
		return errors.New("illegal resource name: " + directory.GetName())
	}

	fullpath := "data/" + path + "/" + directory.GetName()
	err := os.MkdirAll(fullpath, 0777)
	if err != nil && !os.IsExist(err) {
		return errors.New("could not create directory")
	}

	for _, file := range directory.GetFiles() {
		err := os.WriteFile(fullpath+"/"+file.GetName(), file.GetData(), 0777)
		if err != nil {
			return errors.New("could not write write file: " + path + "/" + directory.GetName() + "/" + file.GetName())
		}
	}

	for _, dir := range directory.GetDirectories() {
		err := AddDirectory(path+"/"+directory.GetName(), dir)
		if err != nil {
			return err
		}
	}

	return nil
}

func AddFile(path string, file *pb.File) error {
	if strings.Contains(file.GetName(), "/") || strings.Contains(file.GetName(), "..") {
		return errors.New("illegal resource name: " + file.GetName())
	}

	err := os.MkdirAll("data/"+path, 0777)
	if err != nil && !os.IsExist(err) {
		return errors.New("could not create directory")
	}
	err = os.WriteFile("data/"+path+"/"+file.GetName(), file.GetData(), 0777)
	if err != nil {
		return errors.New("could not write write file: " + path + "/" + file.GetName() + "/" + file.GetName())
	}

	return nil
}
