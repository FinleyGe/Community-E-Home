package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Content   string    `json:"content"`
	Address   string    `json:"address"`
	Require   uint      `json:"require"`
	Accept    uint      `json:"accept"`
	TimeStart time.Time `json:"time_start" gorm:"default:2020-1-1 00:00:00"`
	TimeEnd   time.Time `json:"time_end" gorm:"default:2020-1-1 00:00:00"`
	IssuerId  uint      `json:"issuer_id"`
}
