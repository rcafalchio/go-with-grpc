PROTOC_GEN_GO := protoc --proto_path=greet/proto --go_out=. --go_opt=module=github.com/rcafalchio/grpc-go --go-grpc_out=. --go-grpc_opt=module=github.com/rcafalchio/grpc-go

all: generate

generate:
	$(PROTOC_GEN_GO) greet/proto/dummy.proto

clean:
	rm -f greet/proto/*.pb.go

.PHONY: all generate