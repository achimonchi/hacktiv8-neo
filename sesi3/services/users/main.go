package main

import (
	"context"
	"log"
	"net"
	"sesi3/config"
	"sesi3/models"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	usersLocal *models.UserList
)

func init() {
	usersLocal = new(models.UserList)
	usersLocal.List = make([]*models.User, 0)
}

type UserServer struct {
	models.UnimplementedUsersServer
}

func (u UserServer) Register(ctx context.Context, param *models.User) (*empty.Empty, error) {
	usersLocal.List = append(usersLocal.List, param)
	log.Println("registering user ", param.String())

	return new(empty.Empty), nil
}

func (u UserServer) List(ctx context.Context, empty *empty.Empty) (*models.UserList, error) {
	return usersLocal, nil
}

func main() {
	server := grpc.NewServer()

	var userSrv UserServer
	models.RegisterUsersServer(server, userSrv)

	log.Println("user server running at port", config.SERVICE_USER_PORT)

	listen, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("error when listen to port %v with error %v", config.SERVICE_USER_PORT, err.Error())
	}

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("error when serving user services with error %v", err.Error())
	}
}
