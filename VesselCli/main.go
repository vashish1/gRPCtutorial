package main

import (
	"context"
	"log"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	vesselProto "github.com/vashish1/gRPCtutorial/ShippingVessel/proto/vessel"
)

func main() {
	cmd.Init()
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", microclient.DefaultClient)
	vessel:=*&vesselProto.Vessel{
		MaxWeight: 34,
		Capacity: 45,
		Name: "testingvessel",
	}
	resp,err:=vesselClient.Create(context.Background(),&vessel)
	if err!=nil{
		log.Print("error is",err)
	}
	log.Print(resp)
}
