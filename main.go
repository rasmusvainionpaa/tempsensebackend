package main

import (
	"net/http"
	"tempsensebackend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// adds Temperature routes
	routes.RegisterTemperaturesRoutes(router)

	// gets run when requested route isn't found
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	//starts server
	router.Run(":8000")
}
