package temperaturecontroller

import (
	"net/http"
	"os"
	"tempsensebackend/src/model"

	"github.com/gin-gonic/gin"
)

type request struct {
	Key     string  `json:"key"`
	Outside float32 `json:"outside"`
	Inside  float32 `json:"inside"`
}

func (h handler) PostTemperature(c *gin.Context) {
	r := request{}

	if err := c.BindJSON(&r); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if r.Key != os.Getenv("APIKEY") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Missing api key",
		})
		return
	}

	var temp model.Temperature

	temp.Outside = r.Outside
	temp.Inside = r.Inside

	if result := h.DB.Create(&temp); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &temp)
}
