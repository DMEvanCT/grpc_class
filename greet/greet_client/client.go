package main

import (
	"context"
	"fmt"
	"github.com/DMEvanCT/grpcclass/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err )
	}

	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doClientStreaming(c)

}
func doServiceStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting greet many times client request")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName:            "Evan",
			LastName:             "Haston",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},
	}
	resrStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalln("Error calling RPC...")
	}
	for {
		msg, err := resrStream.Recv()
		if err == io.EOF {
			// WE REACHED THE END OF THE STREAM
			break
		}

		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
		}
		log.Printf("Result from stream: %v", msg.GetResult())
	}



}

func doClientStreaming(c greetpb.GreetServiceClient)  {
	fmt.Println("Starting client streaming")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName:            "Evan",

			},

		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName:            "John",

			},

		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName:            "Steven",

			},

		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName:            "Lucy",

			},

		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName:            "Piper",

			},

		},

	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Issue calling long greet: %v", err)
	}

	// itterate over slice and send each message individual
	for _, req := range requests {
		fmt.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}


	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reciving responce: %v", err)
	}

	fmt.Printf("Long greet responces: %v\n", res)


}