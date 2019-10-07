package main

import (
	"context"
	"fmt"
	calcpb "github.com/DMEvanCT/grpcclass/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	fmt.Printf("Calculator client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("There was an error dialing grpc server %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalculatorServiceClient(cc)

	doErrorUnary(c)

}

func doErrorUnary(c calcpb.CalculatorServiceClient) {
	// correct call
	doErrorCall(c, 10)


	//error call
	doErrorCall(c, -2 )
}

func doErrorCall(c calcpb.CalculatorServiceClient, n int32) {
	fmt.Println("Starting square root unary")
	res, err := c.SquareRoot(context.Background(), &calcpb.SquareRootReq{Number: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual user error
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number")
				return
			}

		} else {
			log.Fatalf("Big error calling square root: %v", err)
			return
		}
	}
	fmt.Printf("Square root of %v = %v", n, res.GetNumberRoot())
}