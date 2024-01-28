package auth

import "time"

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string
	Email    string
	Phone    string
	Password string
	Created  time.Time `gorm:"column:created_at;autoCreateTime"`
	Updated  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}