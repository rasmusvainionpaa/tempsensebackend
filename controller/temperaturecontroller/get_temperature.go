package temperaturecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetTemperature(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"Status": "Ok",
		"Data":   id,
	})
}
