package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"services/greet/greetpb"
)

func main() {
	fmt.Println("Hi I'm in client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Sorry, client can't talk to the server %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	callGreet(c)

}

func callGreet(c greetpb.GreetServiceClient)  {
	fmt.Println("in Call Greet Function... ")
	req := &greetpb.GreetRequest {
		Greeting : &greetpb.Greeting {
			FirstName : "Naveen",
			LastName : "Kumar",
		},
	}
	res, err := c.Greet(context.Background(), req)
	res1, err1 := c.GreetFullName(context.Background(), req)
	if err != nil {
		log.Fatal("Error while called in Greet: %v", err)
	}
	if err1 != nil {
		log.Fatal("Error while called in Greet: %v", err1)
	}

	log.Println("", res.Result)
	log.Println("", res1.Result)
}
