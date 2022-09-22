package model

import (
	"gorm.io/gorm"
)

type Temperature struct {
	gorm.Model                 // creates a gorm model that adds id and creation time
	OutsideTemperature float32 `json:"OutsideTemperature"`
	OutsideHumidity    float32 `json:"OutsideHumidity"`
	InsideTemperature  float32 `json:"InsideTemperature"`
	InsideHumidity     float32 `json:"InsideHumidity"`
}
