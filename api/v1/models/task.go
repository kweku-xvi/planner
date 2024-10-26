package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `gorm:"column:title;not null" json:"title"`
	Description string    `gorm:"column:description" json:"description,omitempty"`
	Priority    string    `gorm:"column:priority" json:"priority,omitempty"`
	Deadline    time.Time `gorm:"column:deadline" json:"deadline,omitempty"`
	Status      string    `gorm:"column:status" json:"status,omitempty"`
}
