package service

import (
	"context"
	"errors"
	"log"

	"os"
	"strings"

	pb "github.com/assafvayner/grpc-files/fileservice"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/genproto/googleapis/rpc/status"
)

// generates error response for fetch request
func GenFetchErrorResponse(code int32, message string) *pb.FetchResponse {
	return &pb.FetchResponse{
		Status: &status.Status{
			Code:    code,
			Message: message,
		},
	}
}

func (*Server) Fetch(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	// weakly prevent directory traversal attack
	if strings.Contains(req.GetPath(), "..") {
		return GenFetchErrorResponse(1, "Illegal resource name"), nil
	}

	var file *pb.File = nil
	var dir *pb.Directory = nil

	// compose resource path (data directory)
	fullpath := "data/" + req.GetPath()

	// check that resource exists, and is a file
	stat, err := os.Stat(fullpath)
	if err != nil {
		return GenFetchErrorResponse(1, "Requested resource does not exist"), nil
	} else if stat.Mode().IsDir() {
		dir, err = ParseDirectory(fullpath, req.GetMetadataOnly())
	} else if stat.Mode().IsRegular() {
		file, err = ParseFile(fullpath, req.GetMetadataOnly())
	} else {
		log.Println("tried to fetch irregular file: " + fullpath)
		return GenFetchErrorResponse(1, "requested resource is not available"), nil
	}

	if err != nil {
		return GenFetchErrorResponse(1, err.Error()), nil
	}

	res := pb.FetchResponse{
		Status: &status.Status{
			Code:    0,
			Message: "Fetch succeeded",
		},
	}

	if file != nil {
		res.Resource = &pb.FetchResponse_File{
			File: file,
		}
	} else if dir != nil {
		res.Resource = &pb.FetchResponse_Directory{
			Directory: dir,
		}
	} else {
		return GenFetchErrorResponse(1, "requested resource is not available"), nil
	}

	return &res, nil
}

func ParseDirectory(path string, metadataOnly bool) (*pb.Directory, error) {
	direntries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	files := []*pb.File{}
	dirs := []*pb.Directory{}

	for _, dirent := range direntries {
		if dirent.IsDir() {
			// directory contents added later
			dirs = append(dirs, &pb.Directory{
				Name: dirent.Name(),
			})
			continue
		}
		// file contents added later
		files = append(files, &pb.File{
			Name: dirent.Name(),
		})
	}

	for _, file := range files {
		fullpath := path + "/" + file.Name
		stat, err := os.Stat(fullpath)
		if err != nil {
			return nil, errors.New("could not stat file")
		}
		file.Metadata = &pb.FileMetadata{
			Path: fullpath,
			Size: stat.Size(),
			LastModified: &timestamp.Timestamp{
				Seconds: stat.ModTime().Unix(),
				Nanos:   0,
			},
		}
		if metadataOnly {
			continue
		}
		data, err := os.ReadFile(fullpath)
		if err != nil {
			return nil, errors.New("could not read file " + fullpath[5:])
		}
		file.Data = data
	}

	for i := range dirs {
		dirpath := path + "/" + dirs[i].Name
		dir, err := ParseDirectory(dirpath, metadataOnly)
		if err != nil {
			return nil, err
		}
		dirs[i] = dir
	}

	res := pb.Directory{
		Files:       files,
		Directories: dirs,
	}

	return &res, nil
}

func ParseFile(path string, metadataOnly bool) (*pb.File, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, errors.New("could not stat file")
	}
	file := pb.File{
		Metadata: &pb.FileMetadata{
			Path: path,
			Size: stat.Size(),
			LastModified: &timestamp.Timestamp{
				Seconds: stat.ModTime().Unix(),
				Nanos:   0,
			},
		},
		Name: stat.Name(),
	}

	if metadataOnly {
		return &file, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("could not read file " + path[5:])
	}
	file.Data = data

	return &file, nil
}
