package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/vashish1/gRPCtutorial/UserService/proto/user"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())
	UserCollection := client.Database("shipping").Collection("User")
	repo := &UserRepository{UserCollection}
	tokenService := &Token{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
	)

	srv.Init()

	// pubsub := srv.Server().Options().Broker
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
