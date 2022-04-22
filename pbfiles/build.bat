protoc --go_out=. --go_opt=paths=source_relative prod.proto
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative prod.proto