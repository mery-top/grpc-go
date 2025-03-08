package main

import (
	"context"
	"log"
	"net"

	pb "grpc-test/helloworld"

	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedHelloWorldServiceServer
}

func(s * server) SayHello(ctx context.Context, req * pb.HelloRequest) (*pb.HelloResponse, error){
	return &pb.HelloResponse{Message: "Hello" + req.Name }, nil
}

func main(){
	listener, err:= net.Listen("tcp", ":3000");
	if err!=nil{
		log.Fatalf("Faileda to listen %v", err);
	}

	s:=grpc.NewServer();
	pb.RegisterHelloWorldServiceServer(s, &server{});
	log.Println(("gRPC server listening on port 3000"));
	if err:= s.Serve(listener); err!=nil{
		log.Fatalf("Failed to serve: %v", err)
	}

}
