PROTOC_GEN_GO := protoc --proto_path=greet/proto --go_out=. --go_opt=module=github.com/rcafalchio/go-with-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/rcafalchio/go-with-grpc

all: generate

generate:
	$(PROTOC_GEN_GO) greet/proto/greet.proto
	go build -o bin/greet/server ./greet/server
	go build -o bin/greet/client ./greet/client

clean:
	rm -f greet/proto/*.pb.go

.PHONY: all generate