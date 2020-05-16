package main

import (
	"context"

	"github.com/vashish1/gRPCtutorial/ShippingVessel/proto/vessel"
	pb "github.com/vashish1/gRPCtutorial/ShippingVessel/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	FindVessel(context.Context, *pb.Specification) (*pb.Vessel, error)
	CreateVessel(context.Context, *pb.Vessel)(*pb.Response,error)
}

type Vessel struct{
	Id string
	Capacity int32
	MaxWeight int32
	Name string
	Available bool
	OwnerId string
}

type Specification struct{
	Capacity int32
	MaxWeight int32
}

type Response struct{
	Vessel Vessel
	Vessels []Vessel
}
type MongoRepository struct{
	collection *mongo.Collection
}

func (repo MongoRepository)FindVessel(ctx context.Context,spec *pb.Specification ) (*pb.Vessel, error){
	  specs:=unmarshalSpecs(spec)
	  vessel,err:=Find(repo.collection,specs,ctx)
	  if err!=nil{
		  return &pb.Vessel{},nil
	  }
	  return marshalVessel(vessel)
}

func (repo MongoRepository)CreateVessel(ctx context.Context,vessel *pb.Vessel)(*pb.Response,error){
	vessels:=unmarshalVessel(vessel)
	vessel,err:=Insert(repo.collection,vessels,ctx)
	if err!=nil{
		return &pb.Response{},err
	}
	return marshalResponse(vessel)
}


func unmarshalSpecs(spec *pb.Specification)(Specification){
	return Specification{
		Capacity: spec.GetCapacity(),
		MaxWeight: spec.GetMaxWeight(),
	}
}

func unmarshalVessel(vessel *pb.Vessel)(Vessel){
	return Vessel{
		Id: vessel.GetId(),
		OwnerId: vessel.GetOwnerId(),
		Capacity: vessel.GetCapacity(),
		Name: vessel.GetName(),
		Available: vessel.GetAvailable(),
		MaxWeight: vessel.GetMaxWeight(),
	}
}

func marshalVessel(v Vessel)(*pb.Vessel,error){
	return &pb.Vessel{
		Id: v.Id,
		Capacity: v.Capacity,
		MaxWeight: v.MaxWeight,
		OwnerId: v.OwnerId,
		Available: v.Available,
		Name: v.Name,
	},nil
}

func marshalResponse(v Vessel)(*pb.Response,error){
	ves,_:=marshalVessel(v)
	return &pb.Response{
		Vessel: ves,
	},nil
}