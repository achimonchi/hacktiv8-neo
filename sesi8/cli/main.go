package main

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("App", "Simple App")

	add             = app.Command("add", "Add User")
	addFlagOverride = add.Flag("override", "Override user").Short('o').String()
	addArgUser      = add.Arg("user", "username").Required().String()

	update           = app.Command("update", "Update User")
	updateArgOldUser = update.Arg("old", "old username").Required().String()
	updateArgNewUser = update.Arg("new", "new username").Required().String()

	delete          = app.Command("delete", "Delete User")
	deleteFlagForce = delete.Flag("force", "Force delete user").Short('f').String()
	deleteArgUser   = delete.Arg("user", "username").Required().String()

	get         = app.Command("get", "Get user")
	getFlagAll  = get.Flag("all", "get all users").Short('a').String()
	getFlagUser = get.Flag("username", "get by username").Short('u').String()
)

var Users []string = []string{}

func main() {
	add.Action(func(pc *kingpin.ParseContext) error {
		user := *addArgUser
		override := *addFlagOverride

		overrideBool, err := strconv.ParseBool(override)
		if err != nil {
			return err
		}

		fmt.Printf("adding user %s, override %t \n", user, overrideBool)

		return nil
	})
	update.Action(func(pc *kingpin.ParseContext) error {
		old := *updateArgOldUser
		new := *updateArgNewUser

		fmt.Printf("updating %s to %s\n", old, new)
		return nil
	})
	delete.Action(func(pc *kingpin.ParseContext) error {
		user := *deleteArgUser
		force := *deleteFlagForce

		forceBool, err := strconv.ParseBool(force)
		if err != nil {
			return err
		}

		fmt.Printf("delete user with username %s with force is %t\n", user, forceBool)
		return nil
	})

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
