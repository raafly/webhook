package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID 			string	`gorm:"primaryKey"`
	Username	string
	Email		string
	Phone 		string
	Password 	string
}