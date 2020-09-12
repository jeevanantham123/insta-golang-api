package controller

import "github.com/jinzhu/gorm"

//PostController is a struct for database
type PostController struct {
	DB *gorm.DB
}
