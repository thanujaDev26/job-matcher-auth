package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Mobile      string
	City        string
	Province    string
	Email       string `gorm:"unique"`
	Password    string
	Linkedin    string
	Portfolio   string
	ResetToken  string
}
