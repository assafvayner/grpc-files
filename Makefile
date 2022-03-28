# Makefile for this grpc project
# used to compile protos as well as compile and run server

# requires proto file to be in protos directory
# protoc_single=protoc -I=protos/ -I=./ --go_out=. ./protos/$(1).proto

protoc_single=protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./$(1)/$(1).proto

.PHONY: FORCE

proto_service: fileservice/fileservice.proto
	$(call protoc_single,fileservice)

protos: FORCE
	make proto_service

proto_clean: FORCE
	rm -rf protos/*.pb.go

server: FORCE
	go run server/server.go

client: FORCE
	go run client/client.go
