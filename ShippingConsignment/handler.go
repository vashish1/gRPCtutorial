package main

import (
	"errors"
	"fmt"

	pb "github.com/vashish1/gRPCtutorial/ShippingConsignment/proto/consignment"

	vesselProto "github.com/vashish1/gRPCtutorial/ShippingVessel/proto/vessel"
	// "golang.org/x/net/context"
	"golang.org/x/net/context"
)

type handler struct {
	repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	v := &vesselProto.Vessel{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
		Name:      "Number 1.",
	}
	if vesselResponse == nil {
		fmt.Println("11")
		resp, err := s.vesselClient.Create(ctx, v)
		fmt.Println("Vessel created when not found",resp)
		if err!=nil{
			fmt.Print(err)
			return errors.New("error fetching vessel and creating vessel, returned nil")
		}
	}
    fmt.Println("12")
	if err != nil{
		return err
	}
    
	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	if err = s.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments -
func (s *handler) GetConsignment(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
