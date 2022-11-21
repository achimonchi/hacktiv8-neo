package main

import (
	"fmt"
	"sesi3/models"
)

func main() {
	var user1 = models.User{
		Id:       "U2",
		Name:     "Bank Neo",
		Password: "bankneo",
		Gender:   models.UserGender_MALE,
	}

	var userList = &models.UserList{
		List: []*models.User{
			&user1,
		},
	}

	var garage1 = &models.Garage{
		Id:   "G1",
		Name: "Garage 1",
		Coordinate: &models.GarageCoordinate{
			Latitude:  23.001,
			Longitude: 53.001,
		},
	}

	var garageList1 = &models.GarageList{
		List: []*models.Garage{
			garage1,
			garage1,
		},
	}

	var garageListByUser = &models.GarageListByUser{
		List: map[string]*models.GarageList{
			user1.Id: garageList1,
		},
	}

	fmt.Printf("User List :%+v\n", userList)
	fmt.Printf("Garage List :%+v\n", garageListByUser)
}
