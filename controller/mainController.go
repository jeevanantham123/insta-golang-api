package controller

import (
	"github.com/jeevanantham123/insta-golang-api/model"
	"github.com/jeevanantham123/insta-golang-api/repo"
	"github.com/jinzhu/gorm"
)

//UserController is a struct for database
type UserController struct {
	DB *gorm.DB
}

//SayHello func to return something
func SayHello(db *gorm.DB, val string) string {
	var out = repo.SayHell(db, val)
	return out
}

//Signup func for new user
func Signup(db *gorm.DB, user model.User) (string, string) {
	var success, err = repo.Signup(db, user)
	return success, err
}
