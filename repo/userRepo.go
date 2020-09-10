package repo

import (
	"github.com/jeevanantham123/insta-golang-api/model"
	"github.com/jinzhu/gorm"
)

//Signup func for new User
func Signup(db *gorm.DB, user model.User) (string, string) {

	var res model.User
	data := db.Where("Email = ?", user.Email).First(&res).RecordNotFound()

	if data == true {
		if err := db.Create(&user).Error; err != nil {
			return "", err.Error()
		}
		return "User added successfully", ""
	}

	return "", "Email already exists - Error"
}
