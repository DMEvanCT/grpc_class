package main

import (
	"context"
	"fmt"
	"github.com/DMEvanCT/grpcclass/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"net"
	"log"
)

type server struct {
}

func (s *server) SquareRoot(ctx context.Context, req *calcpb.SquareRootReq) (*calcpb.SquareRootResp, error) {
	number := req.GetNumber()
	if (number < 0) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Recived negative number: %v", number))

	}
	return &calcpb.SquareRootResp{
		NumberRoot:           math.Sqrt(float64(number)),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}


func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed the listen %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}




