package model

import (
	"database/sql"
	"home-server/utility"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt utility.Time
	UpdatedAt utility.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
