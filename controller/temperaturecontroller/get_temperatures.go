package temperaturecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTemperatures(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"h": "w"})
}
