package main

import (
	"context"
	"fmt"
	"micro-demo/grpc/demo"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := demo.NewDemoServiceClient(conn)
	reply, err := client.Demo(context.Background(), &demo.String{Value: "123"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(reply.GetValue())
}
