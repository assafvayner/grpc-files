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

func Fetch(resource, destination string) {
	conn, err := grpc.Dial("localhost:8808", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error on grpc dial", err)
	}
	defer conn.Close()

	client := pb.NewRetreiverClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := pb.FetchRequest{
		Path: resource,
	}

	res, err := client.Fetch(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	if res.GetStatus().GetCode() != 0 {
		log.Fatal("Error from server:\n\t" + res.GetStatus().GetMessage())
	}

	err = os.MkdirAll(destination, 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal("could not create destination directory")
	}

	file := res.GetFile()
	if file != nil {
		
		os.WriteFile(destination+"/"+file.GetName(), file.GetData(), 0777)
		log.Printf(destination + "/" + file.GetName())
		os.Exit(0)
	}

	dir := res.GetDirectory()
	if dir == nil {
		log.Fatalf("Failed to retrieve %s\n", resource)
	}

	AddDirectory(destination, dir)

	os.Exit(0)
}

func AddDirectory(path string, dir *pb.Directory) {
	fullpath := path + "/" + dir.GetName()
	os.MkdirAll(fullpath, 0777)
	log.Println(fullpath + "/")
	for _, file := range dir.GetFiles() {
		os.WriteFile(fullpath+"/"+file.GetName(), file.GetData(), 0777)
		log.Printf("\t" + fullpath + "/" + file.GetName())
	}
	for _, subdir := range dir.GetDirectories() {
		AddDirectory(fullpath, subdir)
	}
}
