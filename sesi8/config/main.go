package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	viper "github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file changed in time", in.Name)
	})
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		GetKey()
		fmt.Println("PORT", viper.GetString("app.port"))
		fmt.Println("DB_PASS", os.Getenv("DB_PASS"))
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "ok",
		})
	})

	port := fmt.Sprintf(":%s", viper.GetString("app.port"))
	e.Start(port)

}

func GetKey() {
	godotenv.Overload("./config/.fly.env")
}
