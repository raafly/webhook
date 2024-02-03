package auth

import (
	"time"
)

type User struct {
	ID 				  string `gorm:"primaryKey"`
	UUID              string `gorm:"column:uuid"`
	Email             string `gorm:"unique"`
	Name              string
	Token             string
	Password          string
	EmailVerifiedAt   time.Time `gorm:"column:email_verified_at"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdateAt          time.Time `gorm:"column:update_at;autoCreateTime;autoUpdateTime"`
}

/*
type resetPassword struct {
	ID        string
	Email     string
	Token     string
	CreatedAt time.Time	`gorm:"column:created_at;autoCreateTime;<-:create"`
}

func (t *resetPassword) TableName() string {
	return "reset_password"
}
*/