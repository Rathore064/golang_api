package main

import (
	"fmt"
	"instagram-api/controller"
	"log"
	"net/http"
)

func main() {

	//function to handle api endpoints
	http.HandleFunc("/users", controller.CreatePost)
	http.HandleFunc("/users/", controller.GetPost)
	http.HandleFunc("/posts", controller.CreatePost)
	http.HandleFunc("/posts/", controller.GetPost)
	http.HandleFunc("/posts/users/", controller.GetUsersPost)

	//start the server on port 8000
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
