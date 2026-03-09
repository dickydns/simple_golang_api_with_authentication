package main

import (
	"fmt"
	"log"
	"simple_golang_api_with_authentication/config"
	"simple_golang_api_with_authentication/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load file .env
	envload := godotenv.Load()
	if envload != nil {
		log.Fatal("Gagal load .env")
	}

	err := config.ConnectDb()
	if err != nil {
		log.Fatal("Failed Connect to db:", err)
	} else {
		fmt.Println("Connect db")
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)
	routes.UserRoute(r)
	//runnig router
	r.Run(":3100")
}
