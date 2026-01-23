package main

import (
	"belajar_go/config"
	"belajar_go/routes"
	"fmt"
	"log"

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

	routes.UserRoute(r)
	//runnig router
	r.Run()
}
