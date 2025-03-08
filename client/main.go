package main

import (
	"context"
	"log"
	"time"

	pb "grpc-test/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main(){
	conn,err:= grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("Failed to connect:%v", err)
	}
	defer conn.Close()
	client:= pb.NewHelloWorldServiceClient(conn)
	ctx, cancel:= context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err:=client.SayHello(ctx, $pb.He&pb.HelloRequest{Name: "Alice"})
	if err!=nil{
		log.Fatalf("Error Calling SayHello: %v", err)
		log.Printf("Server Response: %s", res.Message)
	}
}