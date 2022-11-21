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

type GarageServer struct {
	models.UnimplementedGaragesServer
}

var (
	garageLocal *models.GarageListByUser
)

func init() {
	garageLocal = new(models.GarageListByUser)
	garageLocal.List = make(map[string]*models.GarageList)

	usersLocal = new(models.UserList)
	usersLocal.List = make([]*models.User, 0)
}

func (GarageServer) List(ctx context.Context, param *models.GarageUserId) (*models.GarageList, error) {
	return garageLocal.List[param.UserId], nil
}

func (GarageServer) Add(ctx context.Context, param *models.GarageAndUserId) (*empty.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := garageLocal.List[userId]; !ok {
		garageLocal.List[userId] = new(models.GarageList)
		garageLocal.List[userId].List = make([]*models.Garage, 0)
	}

	garageLocal.List[userId].List = append(garageLocal.List[userId].List, garage)

	log.Println("registering garage", param)

	return new(empty.Empty), nil
}

var (
	usersLocal *models.UserList
)

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
	var server = grpc.NewServer()

	var garageSrv GarageServer
	var userSrv UserServer
	models.RegisterGaragesServer(server, garageSrv)
	models.RegisterUsersServer(server, userSrv)

	log.Println("garage server running at port", config.SERVICE_GARAGE_PORT)

	listen, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("error when listen to port %v with error %v", config.SERVICE_GARAGE_PORT, err.Error())
	}

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("error when serving user services with error %v", err.Error())
	}
}
