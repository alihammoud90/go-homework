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
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHomeworkServiceClient(conn)

	fmt.Println("Choose Task: ")
	fmt.Println("1) Sum ")
	fmt.Println("2) Order Numbers ")
	var task int
	fmt.Scanln(&task)
	switch task {
	case 1:
		callSum(client)
	case 2:
		callOrder(client)
	default:
		fmt.Println("Invalid Task")
	}
}

func callSum(client pb.HomeworkServiceClient) {
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

func callOrder(client pb.HomeworkServiceClient) {
	fmt.Print("Enter Array size: ");
	var size int;
	fmt.Scanln(&size);
	array := make([]int64, size);
	for i := 0; i < size; i++ {
		fmt.Printf("Number %d: ", i + 1);
		var nbx int64;
		fmt.Scanln(&nbx);
		array[i] = nbx;
	}

	result, err := client.OrderNumbers(context.Background(), &pb.NumbersToOrder{Array: array});
	if err != nil {
		log.Fatalf("Could not get sum: %v", err)
	}
	log.Println("ordered: ", result.OrderedArray);
}
