package main

import (
	"fmt"
	"os"

	"github.com/jeevanantham123/insta-golang-api/driver"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Cannot find .env file")
		return
	}

	serverPort := os.Getenv("SERVER_PORT")
	serverPort = ":" + serverPort

	if serverPort == ":" {
		fmt.Println("Unable to find SERVER_PORT from environmental variables")
		return
	}

	db, err := driver.Connect()

	if err == nil {
		fmt.Println("GO server started and running at port 8123")
		fmt.Println(db)
		fmt.Println("Server stopped")
	}
}
