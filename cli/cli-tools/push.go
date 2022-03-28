package cli

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/assafvayner/grpc-files/fileservice"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Push(resource, destination string) {
	stat, err := os.Stat(resource)
	if err != nil {
		log.Fatalf("Resource %s is not available\n", resource)
	}
	var dir *pb.Directory
	var file *pb.File

	if stat.IsDir() {
		dir, err = ParseDirectory(resource)
	} else {
		file, err = ParseFile(resource)
	}
	if err != nil {
		log.Fatal(err)
	}

	req := pb.PushRequest{
		Path: destination,
	}

	if dir != nil {
		req.Resource = &pb.PushRequest_Directory{
			Directory: dir,
		}
	} else if file != nil {
		req.Resource = &pb.PushRequest_File{
			File: file,
		}
	} else {
		log.Fatal("Could not parse anything to push to server")
	}

	conn, err := grpc.Dial("localhost:8808", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error on grpc dial", err)
	}
	defer conn.Close()

	client := pb.NewRetreiverClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Push(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	if res.GetStatus().GetCode() != 0 {
		log.Fatal("Error from server:\n\t" + res.GetStatus().GetMessage())
	}
	log.Printf("Successfully pushed: %s\n", resource)
	os.Exit(0)
}

func ParseDirectory(path string) (*pb.Directory, error) {
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
		data, err := os.ReadFile(fullpath)
		if err != nil {
			return nil, errors.New("could not read file " + fullpath[5:])
		}
		file.Data = data
	}

	for i := range dirs {
		dirpath := path + "/" + dirs[i].Name
		dir, err := ParseDirectory(dirpath)
		if err != nil {
			return nil, err
		}
		dirs[i] = dir
	}

	name := NameFromPath(path)

	res := pb.Directory{
		Name:        name,
		Files:       files,
		Directories: dirs,
	}

	return &res, nil
}

func ParseFile(path string) (*pb.File, error) {
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

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("could not read file " + path[5:])
	}
	file.Data = data

	return &file, nil
}

func NameFromPath(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx == len(path)-1 {
		idx = strings.LastIndex(path[:len(path)-1], "/")
	}
	if idx == -1 {
		return path
	}
	return path[idx+1:]
}
