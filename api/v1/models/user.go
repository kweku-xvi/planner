package models

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Username  string    `gorm:"column:username;unique" json:"username"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
