package model

import "github.com/jinzhu/gorm"

//User structure
type User struct {
	gorm.Model

	Firstname string

	LastName string

	Email string

	UserName string `gorm:"unique"`

	Password string

	ProfileURL string

	About string
}

//UserOutput model
type UserOutput struct {
	ID    uint
	Name  string
	Email string
}
