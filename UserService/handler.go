package main

import (
	"fmt"
	"log"

	pb "github.com/vashish1/gRPCtutorial/UserService/proto/user"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type service struct {
	repo Repository
	tokenService Authable
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(ctx, req.Id)
	if err != nil {
		return err
	}
    u:=marshalUser(user)
	res.User = u
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll(ctx)
	if err != nil {
		return err
	}
	collection := make([]*pb.User, 0)
	for _,u:=range users{
		 pbUser:=marshalUser(u)
		 collection=append(collection,pbUser)
	}
	res.Users = *&collection
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)
	r:=unmarshelUser(req)
	user, err := srv.repo.GetByEmail(ctx,r.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(req)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
		// Generates a hashed version of our password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Print("1")
			return err
		}
		req.Password = string(hashedPass) 
		if err := srv.repo.Create(ctx,req); err != nil {
			fmt.Print("2")
			return err
		}
		res.User = req
		return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
