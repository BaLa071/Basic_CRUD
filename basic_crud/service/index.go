package service

import (
	"basic_crud/models"
	pb "basic_crud/proto"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	pb.UnimplementedCustomerServiceServer
}

var (
	Collection *mongo.Collection
)

func (s *Server) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateRes, error) {

	customer := &models.CustomerDetails{
		Id:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Salary:   req.Salary,
		Address:  req.Address,
	}
	res, err := Collection.InsertOne(context.Background(), customer)
	if err != nil {
		fmt.Println("Cannot insert data!")
	}
	fmt.Println(res)
	return &pb.CreateRes{Message: "succesfully created"}, nil

}

func (s *Server) Read(ctx context.Context, req *pb.ReadReq) (*pb.ReadRes, error) {
	query := bson.M{"_id": req.Id}
	var customer *models.CustomerDetails
	res := Collection.FindOne(context.Background(), query).Decode(&customer)
	if res != nil {
		fmt.Println(res)
	}
	fmt.Println(customer)
	response := &pb.ReadRes{
		Id:       customer.Id,
		Name:     customer.Name,
		Email:    customer.Email,
		Password: customer.Password,
		Address:  customer.Password,
		Salary:   customer.Salary,
	}
	return response, nil
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateReq) (*pb.UpdateRes, error) {
	filter := bson.M{"_id": req.Id}
	query := bson.D{{"$set", bson.D{{req.Feild, req.Value}}}}
	res, err := Collection.UpdateOne(ctx, filter, query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	return &pb.UpdateRes{
		Message: "Updated!!",
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteReq) (*pb.DeleteRes, error) {
	filter := bson.M{"_id": req.Id}
	res, err := Collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
	return &pb.DeleteRes{
		Message: "Deleted!!",
	}, nil
}
