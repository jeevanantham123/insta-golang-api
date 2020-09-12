package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeevanantham123/insta-golang-api/controller"
	"github.com/jeevanantham123/insta-golang-api/middleware"
)

//Posts structure
type Posts struct {
	PostID     int    `json:"postid"`
	PostURL    string `json:"posturl"`
	PostLiked  bool   `json:"postliked"`
	PostSaved  bool   `json:"postsaved"`
	PosterName string `json:"postername"`
}

//HandlePostRoutes handle routes
func HandlePostRoutes(postController controller.PostController, router *mux.Router) {
	router.Handle("/getpost", middleware.Authmiddleware(http.HandlerFunc(loadpost))).Methods("GET")
}

//Loadpost api call for getting post - predeclared
func loadpost(w http.ResponseWriter, r *http.Request) {
	var apipost = []Posts{
		Posts{PostID: 1, PostURL: "https://wallpaperplay.com/walls/full/c/8/5/128260.jpg", PostLiked: false, PostSaved: false, PosterName: "jeewa"},
		Posts{PostID: 2, PostURL: "https://www.chromethemer.com/google-chrome/backgrounds/download/random-hd-background-for-google-chrome-ct1404.jpg", PostLiked: false, PostSaved: false, PosterName: "dev"},
		Posts{PostID: 3, PostURL: "https://thewallpaper.co//wp-content/uploads/2016/10/adidas-logo-red-ferrari-nature-strike-random-wallpaper-hd-wallpapers-desktop-images-download-free-windows-wallpapers-amazing-colourful-4k-lovely-2560x1600.jpg", PostLiked: false, PostSaved: false, PosterName: "dev"},
		Posts{PostID: 4, PostURL: "https://c1.wallpaperflare.com/preview/800/288/19/cube-play-random-luck.jpg", PostLiked: false, PostSaved: false, PosterName: "hello"},
	}
	json.NewEncoder(w).Encode(apipost)
}
