package main

import (
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code
	micro "github.com/micro/go-micro"
	pb "github.com/vashish1/gRPCtutorial/ShippingConsignment/proto/consignment"
	vesselProto "github.com/vashish1/gRPCtutorial/ShippingVessel/proto/vessel"
	"golang.org/x/net/context"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	// Set-up micro instance
	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
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

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.client", srv.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
