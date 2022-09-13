package temperaturecontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func RegisterTemperaturesRoutes(gin *gin.Engine, db *gorm.DB) {
	h := &handler{
		db: db,
	}

	routes := gin.Group("/temperatures")
	routes.GET("/", h.GetTemperatures)
	routes.GET("/:id", h.GetTemperature)
	routes.POST("/", h.PostTemperature)
}
