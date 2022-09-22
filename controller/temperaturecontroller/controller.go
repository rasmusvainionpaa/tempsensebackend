package temperaturecontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterTemperaturesRoutes(gin *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := gin.Group("/temperatures")
	routes.GET("/", h.GetTemperatures)
	routes.POST("/", h.PostTemperature)
}
