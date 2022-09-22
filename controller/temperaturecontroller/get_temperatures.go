package temperaturecontroller

import (
	"net/http"
	"tempsensebackend/model"

	"github.com/gin-gonic/gin"
)

func (h handler) GetTemperatures(c *gin.Context) {
	var temps []model.Temperature

	if result := h.DB.Find(&temps); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &temps)
}
