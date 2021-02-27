genG:
	  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculatorpb/*.proto
clean:
	rm calculatorpb/*.go
runs:
	go run greet_server/server.go
runc:
	go run greet_client/client.go