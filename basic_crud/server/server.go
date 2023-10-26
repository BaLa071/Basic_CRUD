package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "basic_crud/proto"
	"basic_crud/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	service.Collection= client.Database("testing").Collection("customer")
	lis, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Listening")
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCustomerServiceServer(s, &service.Server{})
	if err2 := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v", err2)
	}
}
