package model

import (
	"home-server/utility"
)

type Task struct {
	Model
	Content   string       `json:"content"`
	Address   string       `json:"address"`
	Require   uint         `json:"require"`
	Accept    uint         `json:"accept"`
	TimeStart utility.Time `json:"time_start"`
	TimeEnd   utility.Time `json:"time_end"`
	IssuerId  uint         `json:"issuer_id"`
	Success   bool         `json:"success"`
}
