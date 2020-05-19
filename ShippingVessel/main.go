package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/vashish1/gRPCtutorial/ShippingVessel/proto/vessel"
		
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shipping").Collection("vessel")

	repository := &MongoRepository{vesselCollection}

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repository})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
