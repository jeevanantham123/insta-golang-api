package driver

import (
	"errors"
	"fmt"
	"os"

	"github.com/jeevanantham123/insta-golang-api/model"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	//it is the db connector for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func initialMigration() {
	db.AutoMigrate(&model.User{}, &model.FriendSuggestion{}, &model.UserFriend{})
	db.Model(&model.FriendSuggestion{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&model.UserFriend{}).AddForeignKey("user_name", "users(user_name)", "RESTRICT", "RESTRICT")
}

//Connect function returns database pointer
func Connect() (*gorm.DB, error) {

	err := godotenv.Load(".env")

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	if username == "" || password == "" || port == "" || dbName == "" {
		err := errors.New("Environmental variables not found to connect database")
		fmt.Println(err)
		return nil, err
	}

	dbConfig := "port=" + port + " user=" + username + " dbname=" + dbName + " password=" + password + " host=" + host + " sslmode=disable"
	// fmt.Println(dbConfig)
	db, err = gorm.Open("postgres", dbConfig)
	if err != nil {
		fmt.Println("Failed to connect database", err)
	} else {
		fmt.Println("Database connection Successfull")
		initialMigration()
	}

	return db, err

}
