package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/YaoShangTseng/grpc-go-course/sum/sumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	firstNum := req.GetFirstNum()
	secondNum := req.GetSecondNum()
	result := firstNum + secondNum
	res := &sumpb.SumResponse{
		Sum: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
