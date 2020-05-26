package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func Find(c *mongo.Collection, ctx context.Context) ([]*Consignment, error) {

	findOptions := options.Find()
	var result []*Consignment
	cur, err := c.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return []*Consignment{}, err
		
		// fmt.Println(err)
	}
	for cur.Next(context.TODO()) {
		var elem *Consignment
		err := cur.Decode(&elem)
		if err != nil {
			return []*Consignment{}, err
			// fmt.Println(err)
		}

		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	cur.Close(context.TODO())
	if err != nil {
		return []*Consignment{}, err
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
	return result, nil
}
