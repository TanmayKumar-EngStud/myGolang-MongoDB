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

// A. Create a user
// B. Get a user by ID
// C. Create a post
// D. Get a post by ID
// E. Get all posts by user ID
func aCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json") 
	// create user
	fmt.Println("Creating user")
	var Newuser User
	_ = json.NewDecoder(r.Body).Decode(&Newuser)
	var user User;    // This extra step is added, in case some of the elements are not entered in the new user
	user= FillUser(Newuser) 
	user.Password = HashPassword(user.Password) // Securing the password using SHA512
	collection := Client.Database("AppointyDatabase").Collection("users")
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, user)
	user.id = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(user)
	fmt.Println("User Id is: ", user.id)
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
	var Newpost Post
	_ = json.NewDecoder(rq.Body).Decode(&Newpost)
	var post Post;    // This extra step is added, in case some of the elements are not entered in the new post
	post= FillPost(Newpost)
	collection := Client.Database("AppointyDatabase").Collection("posts")
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, post)
	post.id = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(r).Encode(post)
	fmt.Println("Post Id is: ", post.id)
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
// Handlers Sepratiing on the basis of HTTP Method
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

// Client is global
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
	
	http.HandleFunc("/users", UserHandler)
	http.HandleFunc("/users/", UserHandler)
	http.HandleFunc("/Posts", PostHandler)
	http.HandleFunc("/Posts/", PostHandler)
	http.HandleFunc("/users/Post", UserPostsHandler) 
	http.ListenAndServe(":8080",nil)
}