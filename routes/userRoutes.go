package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeevanantham123/insta-golang-api/controller"
)

//HandleUserRoutes for routes
func HandleUserRoutes(router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var output = controller.SayHello("hello")
		json.NewEncoder(w).Encode(output)

	}).Methods("GET")

}
