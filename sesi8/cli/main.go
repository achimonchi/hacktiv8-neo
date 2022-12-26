package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("App", "Simple App")

	add             = app.Command("add", "Add User")
	addFlagOverride = add.Flag("override", "Override user").Short('o').Default("false").String()
	addArgUser      = add.Arg("user", "username").Required().String()

	update           = app.Command("update", "Update User")
	updateArgOldUser = update.Arg("old", "old username").Required().String()
	updateArgNewUser = update.Arg("new", "new username").Required().String()

	delete          = app.Command("delete", "Delete User")
	deleteFlagForce = delete.Flag("force", "Force delete user").Default("false").Short('f').String()
	deleteArgUser   = delete.Arg("user", "username").Required().String()

	get         = app.Command("get", "Get user")
	getFlagAll  = get.Flag("all", "get all users").Short('a').String()
	getFlagUser = get.Flag("username", "get by username").Short('u').String()
)

var Users []string = []string{}

var filePath = "./cli/data.json"

func main() {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var datas = map[string][]string{}
	err = json.Unmarshal(data, &datas)
	if err != nil {
		panic(err)
	}

	Users = datas["users"]

	updateFunc := func(datas map[string][]string, newData []string, key string) error {
		datas[key] = newData
		dataJson, err := json.Marshal(datas)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(filePath, dataJson, 0644)
	}

	add.Action(func(pc *kingpin.ParseContext) error {
		user := *addArgUser
		override := *addFlagOverride

		overrideBool, err := strconv.ParseBool(override)
		if err != nil {
			return err
		}

		fmt.Printf("adding user %s, override %t \n", user, overrideBool)

		if overrideBool {
			index, isExist := findUser(user)
			if isExist {
				Users[index] = user
				return updateFunc(datas, Users, "users")
			} else {
				fmt.Printf("user %s not exists. so will continue to add it \n", user)
			}
		}

		Users = append(Users, user)
		return updateFunc(datas, Users, "users")
	})
	update.Action(func(pc *kingpin.ParseContext) error {
		old := *updateArgOldUser
		new := *updateArgNewUser

		index, isExist := findUser(old)
		if !isExist {
			return fmt.Errorf("username %s not exist", old)
		}

		Users[index] = new

		fmt.Printf("updating %s to %s\n", old, new)
		return updateFunc(datas, Users, "users")
	})
	delete.Action(func(pc *kingpin.ParseContext) error {
		user := *deleteArgUser
		force := *deleteFlagForce

		forceBool, err := strconv.ParseBool(force)
		if err != nil {
			return err
		}

		index, isExist := findUser(user)
		if !isExist {
			return fmt.Errorf("username %s not exist", user)
		}

		fmt.Printf("delete user with username %s with force is %t\n", user, forceBool)
		Users = append(Users[:index], Users[index+1:]...)

		return updateFunc(datas, Users, "users")
	})
	get.Action(func(pc *kingpin.ParseContext) error {
		all := *getFlagAll
		username := *getFlagUser

		if all == "all" {
			fmt.Printf("%s GET ALL USERS %s\n", strings.Repeat("=", 10), strings.Repeat("=", 10))
			for _, user := range Users {
				fmt.Println("User :", user)
			}
			fmt.Printf("%s\n", strings.Repeat("=", 35))
		} else {
			index, isExist := findUser(username)
			if !isExist {
				return fmt.Errorf("username %s not exist", username)
			}
			fmt.Println("User :", Users[index])
		}
		return nil
	})

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func findUser(username string) (int, bool) {
	for i, user := range Users {
		if user == username {
			return i, true
		}
	}

	return 0, false
}
