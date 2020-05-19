package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	pb "github.com/vashish1/gRPCtutorial/ShippingConsignment/proto/consignment"
	vesselProto "github.com/vashish1/gRPCtutorial/ShippingVessel/proto/vessel"
	userService "github.com/vashish1/gRPCtutorial/UserService/proto/user"
	"golang.org/x/net/context"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	// Set-up micro instance
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		// Our auth middleware
		micro.WrapHandler(AuthWrapper),
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

	consignmentCollection := client.Database("shipping").Collection("consignments")

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

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userService.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
