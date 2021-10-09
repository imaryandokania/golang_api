package main

import (
	"Rest_api/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users", handler.CreateUserEndpoint)
	http.HandleFunc("/users/", handler.GetUserByIDEndpoint)
	http.HandleFunc("/posts", handler.CreatePostEndpoint)
	http.HandleFunc("/posts/", handler.GetPostByIDEndpoint)
	http.HandleFunc("/posts/users/", handler.GetUsersPostByIdEndpoint)
	fmt.Println("API has started...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
