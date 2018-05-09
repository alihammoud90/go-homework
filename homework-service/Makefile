build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/go-homework/homework-service \
	  proto/homework.proto
docker:
    docker build -t local/homework-service -f Dockerfile.homework .
