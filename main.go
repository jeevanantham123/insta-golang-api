package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jeevanantham123/insta-golang-api/controller"
	"github.com/jeevanantham123/insta-golang-api/driver"
	"github.com/jeevanantham123/insta-golang-api/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
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

	userController := controller.UserController{DB: db}
	postController := controller.PostController{DB: db}
	if err == nil {
		fmt.Println("GO server started and running at port" + serverPort)
		router := mux.NewRouter().StrictSlash(true)
		routes.HandleUserRoutes(userController, router)
		routes.HandlePostRoutes(postController, router)
		log.Fatal(http.ListenAndServe(serverPort, cors.AllowAll().Handler(router)))
		fmt.Println("Server stopped")
	}
}
