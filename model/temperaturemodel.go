package model

import (
	"gorm.io/gorm"
)

type temperature struct {
	gorm.Model         // creates a gorm model that adds id and creation time
	Outside    float32 `json:"outside"`
	Inside     float32 `json:"inside"`
}
