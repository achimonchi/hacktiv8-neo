package main

import (
	"context"
	"fmt"
	"log"
	"sesi3/config"
	"sesi3/models"
	"time"

	"google.golang.org/grpc"
)

func registerUserServices() models.UsersClient {
	port := config.SERVICE_GARAGE_PORT

	conn := connection(port, grpc.WithInsecure())

	return models.NewUsersClient(conn)
}

func registerGarageServer() models.GaragesClient {
	port := config.SERVICE_GARAGE_PORT

	conn := connection(port, grpc.WithInsecure())

	return models.NewGaragesClient(conn)
}

func connection(port string, opts ...grpc.DialOption) *grpc.ClientConn {
	conn, err := grpc.Dial(port, opts...)
	if err != nil {
		log.Fatalf("could not connect to %v with error %v", port, err.Error())
	}
	return conn
}

func register(ctx context.Context, user models.UsersClient) {
	newCtx, ctxCancel := context.WithTimeout(ctx, 1*time.Second)
	defer ctxCancel()

	user.Register(newCtx, nil)

}

func main() {
	var user1 = models.User{
		Id:       "U2",
		Name:     "Bank Neo",
		Password: "bankneo",
		Gender:   models.UserGender_MALE,
	}

	ctx := context.Background()

	// user := registerUserServices()
	// user.Register(ctx, &user1)

	// list, err := user.List(ctx, new(empty.Empty))
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println(list)

	garage := registerGarageServer()

	var req = &models.GarageAndUserId{
		UserId: user1.Id,
		Garage: &models.Garage{
			Id:   "G1",
			Name: "Garage 1",
			Coordinate: &models.GarageCoordinate{
				Latitude:  32.001,
				Longitude: 94.001,
			},
		},
	}
	garage.Add(ctx, req)

	listGarage, err := garage.List(ctx, &models.GarageUserId{
		UserId: req.UserId,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(listGarage)

}
