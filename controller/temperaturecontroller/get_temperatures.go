package temperaturecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetTemperatures(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"h": "w"})
}
