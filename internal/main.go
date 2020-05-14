package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/bubble-hook/lightnet-assignment/shared"

	"github.com/bubble-hook/lightnet-assignment/internal/services"

	pb "github.com/bubble-hook/lightnet-assignment/calculator_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type calculatorServer struct {
	calculatorService services.Calculator
}

func (c *calculatorServer) Sum(ctx context.Context, request *pb.CalculateRequest) (*pb.CalculateResponnse, error) {
	return &pb.CalculateResponnse{
		A:      request.GetA(),
		B:      request.GetB(),
		Result: c.calculatorService.Sum(shared.CalculateRequest{A: request.GetA(), B: request.GetB()}),
	}, nil
}
func (c *calculatorServer) Sub(ctx context.Context, request *pb.CalculateRequest) (*pb.CalculateResponnse, error) {
	return &pb.CalculateResponnse{
		A:      request.GetA(),
		B:      request.GetB(),
		Result: c.calculatorService.Sub(shared.CalculateRequest{A: request.GetA(), B: request.GetB()}),
	}, nil
}
func (c *calculatorServer) Mul(ctx context.Context, request *pb.CalculateRequest) (*pb.CalculateResponnse, error) {
	return &pb.CalculateResponnse{
		A:      request.GetA(),
		B:      request.GetB(),
		Result: c.calculatorService.Mul(shared.CalculateRequest{A: request.GetA(), B: request.GetB()}),
	}, nil
}
func (c *calculatorServer) Div(ctx context.Context, request *pb.CalculateRequest) (*pb.CalculateResponnse, error) {
	return &pb.CalculateResponnse{
		A:      request.GetA(),
		B:      request.GetB(),
		Result: c.calculatorService.Div(shared.CalculateRequest{A: request.GetA(), B: request.GetB()}),
	}, nil
}

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "localhost", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := calculatorServer{
		calculatorService: services.Calculator{},
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculateServer(grpcServer, &server)

	reflection.Register(grpcServer)
	fmt.Println("Internal Server Start @ :" + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
