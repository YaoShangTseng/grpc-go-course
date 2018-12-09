package main

import (
	"context"
	"fmt"
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

	doUnary(c)
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
