package main

import (
	"context"
	"fmt"
	"github.com/DMEvanCT/grpcclass/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type server struct {
}



func (s *server) GreetEveryone( stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("Starting greet everyone stream")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err )
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		sendErr := stream.Send(&greetpb.GreetEveryoneResp{
			Result:               result,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending the stream to client: %v", err )
			return err

		}
	}
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

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//Finished reading stream
			return stream.SendAndClose(&greetpb.LongGreetResp{
				Result: result,
			})

		}
		if err != nil {
			log.Fatalf("Error reading stream %v",err )
		}

		firstName := req.GetGreeting().GetFirstName()
		result += firstName + "!"
	}
	return nil
}

func (*server) GreatEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("Invoking Greet Everyone")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil

		}
		if err != nil {
			log.Fatalf("ERROR READING STREAM: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! \n"
		Senderr := stream.Send(&greetpb.GreetEveryoneResp{
			Result: result,
		})
		if Senderr != nil {
			log.Printf("ERROR SENDING RESULT: %V", err)
			return  err
		}
	}
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

