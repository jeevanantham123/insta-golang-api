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
	"github.com/jeevanantham123/insta-golang-api/model"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

//JwtToken struct
type JwtToken struct {
	Token   string `json:"token"`
	Success string `json:"success"`
}

//Exception struct
type Exception struct {
	Message string `json:"message"`
}

//Claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//HandleUserRoutes for routes
func HandleUserRoutes(userController controller.UserController, router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var output = controller.SayHello(userController.DB, "hello")
		json.NewEncoder(w).Encode(output)

	}).Methods("GET")

	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != io.EOF {
			var success, err = controller.Signup(userController.DB, user)

			if err == "" {
				expirationTime := time.Now().Add(24 * time.Minute)
				// Create the JWT claims, which includes the username and expiry time
				claims := &Claims{
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
					json.NewEncoder(w).Encode(JwtToken{Token: tokenString, Success: success})
				} else {
					json.NewEncoder(w).Encode(er)
				}
			} else {
				json.NewEncoder(w).Encode(err)
			}
		}
	}).Methods("GET")

}
