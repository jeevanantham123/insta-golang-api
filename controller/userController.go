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

//Login func
func Login(db *gorm.DB, username string, password string) (*gorm.DB, error) {
	var success, err = repo.Login(db, username, password)
	return success, err
}

//Friends func
func Friends(db *gorm.DB, username string) ([]string, error) {
	var output, err = repo.Friends(db, username)
	return output, err
}

//About func
func About(db *gorm.DB, username string) (string, error) {
	var output, err = repo.About(db, username)
	return output, err
}

//Profile func
func Profile(db *gorm.DB, username string) (string, error) {
	var output, err = repo.Profile(db, username)
	return output, err
}

//SuggestionTable func
func SuggestionTable(db *gorm.DB, username string) ([]model.SuggestionTab, error) {
	var output, err = repo.SuggestionTable(db, username)
	return output, err
}

//Requesting func
func Requesting(db *gorm.DB, username string, friendname string) ([]string, error) {
	var output, err = repo.Requesting(db, username, friendname)
	return output, err
}

//Accepting func
func Accepting(db *gorm.DB, username string, friendname string) ([]string, error) {
	var output, err = repo.Accepting(db, username, friendname)
	return output, err
}
