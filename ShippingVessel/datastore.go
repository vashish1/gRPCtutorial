package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"go.mongodb.org/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongodb/mongo-go-driver/mongo/options"
)

// CreateClient -
func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err := conn.Ping(ctx, nil); err != nil {
		if retry >= 3 {
			return nil, err
		}
		retry = retry + 1
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, uri, retry)
	}

	return conn, err
}

func Find(c *mongo.Collection, s Specification, ctx context.Context) (Vessel, error) {
	filter := bson.D{
		primitive.E{Key: "capacity", Value: s.Capacity},
		primitive.E{Key: "maxweight", Value: s.MaxWeight}}
	var result Vessel

	err := c.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return Vessel{}, err
	}
	return result, nil
}

func Insert(c *mongo.Collection, v Vessel, ctx context.Context) (Vessel, error) {

	insertResult, err := c.InsertOne(ctx, v)
	if err != nil {
		fmt.Println("THe error is", err)
		return Vessel{}, err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return v, nil
}
