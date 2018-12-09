package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/YaoShangTseng/grpc-go-course/sum/sumpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connnect: %v", err)
	}
	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	// doUnary(c)

	doServerStreaming(c)
}

func doUnary(c sumpb.SumServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &sumpb.SumRequest{
		FirstNum:  5,
		SecondNum: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Sum)
}

func doServerStreaming(c sumpb.SumServiceClient) {
	fmt.Println("Starting to do a PrimeDecomposition Server Streaming RPC...")
	req := &sumpb.PrimeNumberDecompositionRequest{
		Number: 14,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Someting happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
