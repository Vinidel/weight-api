package models

import (
	"time"
)

// Weight struct
type Weight struct {
	Kilograms float32  `json:"kilograms"`
	CreatedAt time.Time `json:"createdAt"`
}