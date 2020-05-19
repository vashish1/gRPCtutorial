package main

import (
	"context"
	"fmt"

	// "fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/mongodb/mongo-go-driver/bson"/
	// "github.com/mongodb/mongo-go-driver/bson/primitive"
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

func find(c *mongo.Collection, id string, ctx context.Context) (User, error) {
	filter := bson.D{
		primitive.E{Key: "id", Value: id},
	}
	var result User

	err := c.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	return result, nil
}

func GetUserByEmail(c *mongo.Collection,email string,ctx context.Context)(User,error){
	filter:=bson.D{
		{"email",email},
	}
	var result User
	err:=c.FindOne(ctx,filter).Decode(&result)
	if err!=nil{
		fmt.Println(err)
		return User{},err
	}
	return result,nil
}

func findAll(c *mongo.Collection,ctx context.Context)([]User,error){
	
	findOptions := options.Find()
	var result []User
	cur, err := c.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return []User{},err
		fmt.Println(err)
	}
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			return []User{},err
			fmt.Println(err)
		}

		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	cur.Close(context.TODO())
   if err!=nil{
	   return []User{},err
   }
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
	return result,nil
}

func createUser(c *mongo.Collection,user User,ctx context.Context) error{

	insertResult, err := c.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("THe error is", err)
		return err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return  nil
}
