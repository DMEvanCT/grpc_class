package main

import (
	"context"
	"fmt"
	"github.com/DMEvanCT/grpcclass/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err )
	}

	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doServiceStreaming(c)

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