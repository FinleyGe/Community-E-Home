package model

import (
	"home-server/utility"
)

type Model struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt utility.Time
	UpdatedAt utility.Time
	DeletedAt utility.Time
}
