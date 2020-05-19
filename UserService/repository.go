package main

import (
	// "github.com/jinzhu/gorm"
	pb "github.com/vashish1/gRPCtutorial/UserService/proto/user"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user *pb.User) error
	GetByEmail(ctx context.Context, email string) (User, error)
}

type User struct{
	Id string
    Name string
    Company string
     Email string
     Password string
}



type UserRepository struct {
	db *mongo.Collection
}

func (repo *UserRepository) GetAll(ctx context.Context) ([]User, error) {
	var users []User
	users,err := findAll(repo.db,ctx)
	if err!=nil{
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(ctx context.Context, id string) (User, error) {
	user,err:=find(repo.db,id,ctx)
	return user,err
}

func (repo *UserRepository) GetByEmail(ctx context.Context,email string) (User, error) {
	user,err:=GetUserByEmail(repo.db,email,ctx)
	if err!=nil{
		return User{},err
	}
	
	return user,nil
}

func (repo *UserRepository) Create(ctx context.Context,user *pb.User) error {
	Un_user:=unmarshelUser(user)
	err:=createUser(repo.db,Un_user,ctx)
	if err!=nil{
		return err
	}
	return nil
}


func marshalUser(user User)(*pb.User){
	return&pb.User{
		Id: user.Id,
		Name: user.Name,
		Company: user.Company,
		Email: user.Email,
		Password: user.Password,
	}
}

func unmarshelUser(u *pb.User) User{
	return User{
	  Id: u.GetId(),
	  Email: u.GetEmail(),
	  Company: u.GetCompany(),
	  Name: u.GetName(),
	  Password: u.GetPassword(),
	}
}