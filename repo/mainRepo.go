package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// "github.com/jeevanantham123/insta-golang-api/model"
)

//SayHell func to return something
func SayHell(db *gorm.DB, val string) string {
	fmt.Println(db)
	return val
}
