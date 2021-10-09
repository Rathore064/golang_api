package controller

import (
	"context"
	"encoding/json"
	"instagram-api/model"
	"instagram-api/secure"
	"log"
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	user_reg     = regexp.MustCompile(`^\/users\/(\w+)$`)
	post_reg     = regexp.MustCompile(`^\/posts\/(\w+)$`)
	userpost_reg = regexp.MustCompile(`\/posts/users\/(\w+)$`)
)

func Connect_database() (*mongo.Collection, *mongo.Collection) {

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	database_1 := client.Database("Fakestagram").Collection("users")
	database_2 := client.Database("POSTS").Collection("posts")

	return database_1, database_2
}

var collection_1, collection_2 = Connect_database()

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	// encrypting the password of user and storing the encrypted password
	key := "123456789012345678901234"
	hashed_password := secure.Encrypt(key, user.Password)
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
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	Path := user_reg.FindStringSubmatch(r.URL.Path)
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
func CreatePost(w http.ResponseWriter, r *http.Request) {
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
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post model.Post

	Path := post_reg.FindStringSubmatch(r.URL.Path)
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

func GetUsersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []model.Post

	Path := userpost_reg.FindStringSubmatch(r.URL.Path)

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
