build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/go-homework/homework-service \
	  proto/homework.proto
docker:
    docker build -t local/homework-service -f Dockerfile.homework .
kubectl:
    kubectl create -f homework-service.yaml
mock:
    mockgen go-homework/homework-service/proto HomeworkServiceClient > mock_homework/homework_mock.go
