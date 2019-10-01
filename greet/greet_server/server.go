package main

import (
	"context"
	"fmt"
	"github.com/DMEvanCT/grpcclass/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	first_name := req.GetGreeting().GetFirstName()
	result := "Hello " + first_name
	resp := &greetpb.GreetResponse{
		Result:               result,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	return resp, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	first_name := req.GetGreeting().GetFirstName()
	for i := 0; i < 10000; i++  {
		result := "Hello " + first_name + " number " + strconv.Itoa(i)
		res := &greetpb.GreateManyTimesResp{
			Result:               result,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
		stream.Send(res)
		//time.Sleep(1000 * time.Millisecond)

	}
	return nil
}

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed the listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
