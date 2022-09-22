package temperaturecontroller

import (
	"net/http"
	"os"
	"tempsensebackend/model"

	"github.com/gin-gonic/gin"
)

type request struct {
	Key                string  `json:"key"`
	OutsideTemperature float32 `json:"OutsideTemperature"`
	OutsideHumidity    float32 `json:"OutsideHumidity"`
	InsideTemperature  float32 `json:"InsideTemperature"`
	InsideHumidity     float32 `json:"InsideHumidity"`
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

	temp.OutsideTemperature = r.OutsideTemperature
	temp.OutsideHumidity = r.OutsideHumidity
	temp.InsideTemperature = r.InsideTemperature
	temp.InsideHumidity = r.InsideHumidity

	if result := h.DB.Create(&temp); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &temp)
}
