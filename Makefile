all: server/server client/client

myservice/myservice.pb.go: myservice/myservice.proto
	protoc --go_out=plugins=grpc:. myservice/myservice.proto	

server/server: server/*.go myservice/myservice.pb.go
	cd server && go build

client/client: client/*.go myservice/myservice.pb.go
	cd client && go build

