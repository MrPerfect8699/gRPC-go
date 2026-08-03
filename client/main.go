package main

import (
	"context"
	"fmt"
	"grpc-employee/generated/grpc-employee/generated"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		500*time.Second,
	)
	defer cancel()

	md := metadata.New(map[string]string{
		"authorization":  "Bearer my-token-123",
		"request-id":     "req-1001",
		"client-version": "1.0",
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := generated.NewEmployeeServiceClient(conn)

	reqId := &generated.EmployeeRequest{
		Id: 1,
	}

	res, err := client.GetEmployee(ctx, reqId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response is \n", res)

	// stream, err := client.GetAllEmployees(ctx, reqId)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("streaming started...")
	// for {
	// 	res, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(res)
	// }

}
