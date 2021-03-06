package main

import (
	"log"
	"net"
	"sort"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "go-homework/homework-service/proto"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) GetSum(ctx context.Context, in *pb.Numbers) (*pb.Sum, error) {
	return &pb.Sum{Sum: in.Nb1 + in.Nb2}, nil
}

func (s *server) OrderNumbers(ctx context.Context, in *pb.NumbersToOrder) (*pb.OrderedNumbers, error) {
	sort.Slice(in.Array, func(i, j int) bool { return in.Array[i] < in.Array[j] })
	return &pb.OrderedNumbers{OrderedArray: in.Array}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHomeworkServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
