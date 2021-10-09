package handler

import (
	"Rest_api/encrypt"
	"Rest_api/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	getUserRe     = regexp.MustCompile(`^\/users\/(\w+)$`)
	getPostRe     = regexp.MustCompile(`^\/posts\/(\w+)$`)
	getUserPostRe = regexp.MustCompile(`\/posts/users\/(\w+)$`)
)

//function coonects to MongoDB and returns pointers to two collections in the DB
func ConnectDB() (*mongo.Collection, *mongo.Collection) {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB is connected")
	collection_1 := client.Database("appointy").Collection("users")
	collection_2 := client.Database("appointy").Collection("posts")

	return collection_1, collection_2
}

//collection variables
var collection_1, collection_2 = ConnectDB()

//function to create a user in the DB
//POST request
func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	key := "123456789012345678901234"
	hashed_password := encrypt.Encrypt(key, user.Password)
	user.Password = hashed_password

	result, err := collection_1.InsertOne(context.TODO(), &user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(result)

}

//function to get a user by its ID
//GET request
func GetUserByIDEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	Path := getUserRe.FindStringSubmatch(r.URL.Path)
	id := Path[1]

	filter := bson.M{"_id": id}
	err := collection_1.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(user)
}

//function to create posts in DB
//POST request
func CreatePostEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post model.Post
	post.TimeStamp = time.Now()

	_ = json.NewDecoder(r.Body).Decode(&post)
	result, err := collection_2.InsertOne(context.TODO(), &post)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
}

//function to get a post by its ID
//GET Request
func GetPostByIDEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post model.Post

	Path := getPostRe.FindStringSubmatch(r.URL.Path)
	id := Path[1]

	filter := bson.M{"_id": id}
	err := collection_2.FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	json.NewEncoder(w).Encode(post)
}

//function to get all post for a particular user ID
//GET request

func GetUsersPostByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []model.Post

	Path := getUserPostRe.FindStringSubmatch(r.URL.Path)

	id := Path[1]

	cur, err := collection_2.Find(context.TODO(), bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for cur.Next(context.TODO()) {
		var single_post model.Post

		err := cur.Decode(&single_post)
		if err != nil {
			log.Fatal(err)
		}
		if (single_post.UserID) == id {
			posts = append(posts, single_post)
		}
	}

	if err := cur.Err(); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(posts)
}
