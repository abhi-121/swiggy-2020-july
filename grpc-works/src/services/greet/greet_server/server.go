package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"services/greet/greetpb"
	"strings"
)

type server struct {
}

func (*server) GreetFullName( ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetFullNameResponse, error){
	fmt.Println("Function called")

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Hello : " + strings.ToUpper(firstName) + ", " + strings.ToUpper(lastName)
	res := &greetpb.GreetFullNameResponse{
		Result : result,
	}
	return res, nil
}

func (*server) Greet( ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Println("Function called")

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Hello : " + firstName + ", " + lastName
	res := &greetpb.GreetResponse{
		Result : result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello from server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Sorry failed to load server %v: ",err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if s.Serve(lis); err != nil {
		log.Fatal("failed to serve %v: ",err)

	}

}
