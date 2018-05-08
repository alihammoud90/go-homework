package main

import (
	"fmt"
	"log"

	pb "go-homework/homework-service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHomeworkServiceClient(conn)

	fmt.Print("Enter First Number: ");
	var nb1 int64;
	fmt.Scanln(&nb1);
	fmt.Print("Enter Second Number: ");
	var nb2 int64;
	fmt.Scanln(&nb2);

	sum, err := client.GetSum(context.Background(), &pb.Numbers{Nb1: nb1, Nb2: nb2});
	if err != nil {
		log.Fatalf("Could not get sum: %v", err)
	}
	log.Printf("Sum = %d", sum.Sum);
}
