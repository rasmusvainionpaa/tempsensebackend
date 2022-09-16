package main

import (
	"fmt"
	"net/http"
	"tempsensebackend/src/controller/temperaturecontroller"
	"tempsensebackend/src/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	// Connect to database
	db := database.Connect()

	// load env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Error loadin env file")
	}

	// adds Temperature routes
	temperaturecontroller.RegisterTemperaturesRoutes(router, db)

	// gets run when requested route isn't found
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": "404"})
	})

	//starts server
	router.Run(":8000")
}
