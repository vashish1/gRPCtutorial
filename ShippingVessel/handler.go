package main

import (
	"context"

	pb "github.com/vashish1/gRPCtutorial/ShippingVessel/proto/vessel"
)

type service struct {
	repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	vessel, err := s.repository.FindVessel(ctx, req)
	if err != nil {
		return err
	}
	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	response, err := s.repository.CreateVessel(ctx, req)
	if err != nil {

		return err
	}
	res=response
	return nil
}

// func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

// 	// Find the next available vessel
// 	vessel, err := s.repo.FindAvailable(req)
// 	if err != nil {
// 		return err
// 	}

// 	// Set the vessel as part of the response message type
// 	res.Vessel = vessel
// 	return nil
// }
