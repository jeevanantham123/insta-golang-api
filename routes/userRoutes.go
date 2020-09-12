package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jeevanantham123/insta-golang-api/controller"
	"github.com/jeevanantham123/insta-golang-api/middleware"
	"github.com/jeevanantham123/insta-golang-api/model"
	// "github.com/jeevanantham123/insta-golang-api/middleware"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

//HandleUserRoutes for routes
func HandleUserRoutes(userController controller.UserController, router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var output = controller.SayHello(userController.DB, "hello")
		json.NewEncoder(w).Encode(output)

	}).Methods("GET")

	//Signup router
	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != io.EOF {
			var success, err = controller.Signup(userController.DB, user)

			if err == "" {
				expirationTime := time.Now().Add(24 * time.Minute)
				// Create the JWT claims, which includes the username and expiry time
				claims := &model.Claims{
					Username: user.UserName,
					StandardClaims: jwt.StandardClaims{
						// In JWT, the expiry time is expressed as unix milliseconds
						ExpiresAt: expirationTime.Unix(),
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				// Create the JWT string
				tokenString, er := token.SignedString(jwtKey)
				if er == nil {
					http.SetCookie(w, &http.Cookie{
						Name:    "token",
						Value:   tokenString,
						Expires: expirationTime,
						Path:    "/",
					})
					json.NewEncoder(w).Encode(model.JwtToken{Token: tokenString, Success: success})
				} else {
					json.NewEncoder(w).Encode(er)
				}
			} else {
				json.NewEncoder(w).Encode(err)
			}
		}
	}).Methods("GET")

	//Login Route
	router.HandleFunc("/login/{username}/{password}", func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		password := mux.Vars(r)["password"]

		var _, err = controller.Login(userController.DB, username, password)

		if err == nil {
			expirationTime := time.Now().Add(24 * time.Minute)
			// Create the JWT claims, which includes the username and expiry time
			claims := &model.Claims{
				Username: username,
				StandardClaims: jwt.StandardClaims{
					// In JWT, the expiry time is expressed as unix milliseconds
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// Create the JWT string
			tokenString, er := token.SignedString(jwtKey)
			if er == nil {
				http.SetCookie(w, &http.Cookie{
					Name:    "token",
					Value:   tokenString,
					Expires: expirationTime,
					Path:    "/",
				})
				json.NewEncoder(w).Encode(model.JwtToken{Token: tokenString, Success: "200"})
			} else {
				json.NewEncoder(w).Encode(er)
			}
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
		}
	}).Methods("GET")

	//Friends fetching
	router.Handle("/getfriends/{username}", middleware.Authmiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		var output, err = controller.Friends(userController.DB, username)
		if err == nil {
			json.NewEncoder(w).Encode(output)
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
		}
	}))).Methods("GET")

	//About
	router.Handle("/getabout/{username}", middleware.Authmiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		var output, err = controller.About(userController.DB, username)
		if err == nil {
			json.NewEncoder(w).Encode(output)
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
		}
	}))).Methods("GET")

	//Profile
	router.Handle("/getprofile/{username}", middleware.Authmiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		var output, err = controller.Profile(userController.DB, username)
		if err == nil {
			json.NewEncoder(w).Encode(output)
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
		}
	}))).Methods("GET")

	//Suggestion
	router.Handle("/suggestiontable/{username}", middleware.Authmiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		var output, err = controller.SuggestionTable(userController.DB, username)
		if err == nil {
			json.NewEncoder(w).Encode(output)
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
		}
	}))).Methods("GET")

	//Requesting
	router.Handle("/requesting/{username}/{friendname}", middleware.Authmiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		friendname := mux.Vars(r)["friendname"]
		var output, err = controller.Requesting(userController.DB, username, friendname)
		if err == nil {
			json.NewEncoder(w).Encode(output)
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
		}
	}))).Methods("GET")

	//Accepting
	router.Handle("/accepting/{username}/{friendname}", middleware.Authmiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		friendname := mux.Vars(r)["friendname"]
		var output, err = controller.Accepting(userController.DB, username, friendname)
		if err == nil {
			json.NewEncoder(w).Encode(output)
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
		}
	}))).Methods("GET")

}
