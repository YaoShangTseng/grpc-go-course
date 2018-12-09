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

func (*server) PrimeNumberDecomposition(req *sumpb.PrimeNumberDecompositionRequest, stream sumpb.SumService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition function was invoked with %v\n", req)
	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&sumpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
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
