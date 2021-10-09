package main

// All POST function are going to be made here,
// Validation of GET function are made in separate go files.
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct{
	id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"Name,omitempty" bson:"Name,omitempty"`
	Email string `json:"Email,omitempty" bson:"Email,omitempty"`
	Password string `json:"Password,omitempty" bson:"Password,omitempty"`
}

type Post struct{
	id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption string `json:"Caption,omitempty" bson:"Caption,omitempty"`
	ImgURL string `json:"ImgURL,omitempty" bson:"ImgURL,omitempty"`
	user User `json:"User,omitempty" bson:"User,omitempty"`
	startTime string `json:"startTime,omitempty" bson:"startTime,omitempty"`
	endTime string `json:"endTime,omitempty" bson:"endTime,omitempty"`
}

func aCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json") 
	// create user
	fmt.Println("Creating user")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	collection := Client.Database("AppointyDatabase").Collection("users")
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, user)
	user.id = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(user)
	fmt.Println("User Id is: ", user)
}


func bGETUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("context-type","application/json")
	fmt.Println("The Requested User id is:- ",path.Base(r.URL.Path))
	CheckID := path.Base(r.URL.Path)
	checkID, _ := primitive.ObjectIDFromHex(CheckID)
	user, _ := CheckIDofUser(checkID)
	json.NewEncoder(w).Encode(user)
}

func cCreatePost(r http.ResponseWriter, rq *http.Request) {
	r.Header().Set("content-type", "application/json") 
	// create post
	fmt.Println("Creating post")
	var post Post
	_ = json.NewDecoder(rq.Body).Decode(&post)

	collection := Client.Database("AppointyDatabase").Collection("posts")
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, post)
	post.id = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(r).Encode(post)
	fmt.Println("Post Id is: ", post)
}

func dGETPostByID(r http.ResponseWriter, rq *http.Request) {
	r.Header().Set("context-type","application/json")
	fmt.Println("The Requested post id is:- ",path.Base(rq.URL.Path))
	CheckID := path.Base(rq.URL.Path)
	checkID, _ := primitive.ObjectIDFromHex(CheckID)
	Posts, _ := CheckIDofPost(checkID)
	json.NewEncoder(r).Encode(Posts)
}
func eGETUserPosts(r http.ResponseWriter, rq *http.Request) {
	r.Header().Add("content-type", "application/json") 
	// Getting all posts
	fmt.Println("Getting all posts")
	var posts []Post
	collection := Client.Database("AppointyDatabase").Collection("posts")
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post Post
		err := cursor.Decode(&post)
		if err != nil {
			fmt.Println(err)
		}
		posts = append(posts, post)
	}
	json.NewEncoder(r).Encode(posts)
}


func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		aCreateUser(w, r)
	case "GET":
		bGETUserByID(w, r)
	}
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cCreatePost(w, r)
	case "GET":
		dGETPostByID(w, r)
	}
}

func UserPostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		eGETUserPosts(w, r)
	}
}	

var Client *mongo.Client
func main() {
	// create client
	fmt.Println("Starting the application")
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// check connection
	err := Client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to MongoDB!")
	
	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/users/", UserHandler)
	http.HandleFunc("/Post", PostHandler)
	http.HandleFunc("/Posts/", PostHandler)
	http.HandleFunc("/users/Post", UserPostsHandler) 
	http.ListenAndServe(":8080",nil)
}