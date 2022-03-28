package main

import (
	"log"
	"os"

	cli "github.com/assafvayner/grpc-files/cli/cli-tools"
)

func main() {
	if len(os.Args) < 2 {
		Usage("help")
	}

	cmd := os.Args[1]

	switch cmd {

	case "fetch":
		if len(os.Args) != 4 {
			Usage("fetch")
		}
		cli.Fetch(os.Args[2], os.Args[3])
	case "push":
		if len(os.Args) != 4 {
			Usage("push")
		}
		cli.Push(os.Args[2], os.Args[3])
	case "remove":
		if len(os.Args) != 3 {
			Usage("remove")
		}
		cli.Remove(os.Args[2])
	default:
		Usage("help")
	}
}

func Usage(tool string) {
	switch tool {
	case "fetch":
		log.Fatal("./cli fetch <remote resource path> <local destination path>")
	case "push":
		log.Fatal("./cli push <local resource path> <remote destination path>")
	case "remove":
		log.Fatal("./cli remove <remote resource path>")
	default:
		PrintHelpMessage()
		os.Exit(1)
	}
}

func PrintHelpMessage() {
	log.Println("unimplemented")
}
