package main

// All POST function are going to be made here,
// Validation of GET function are made in separate go files
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

func (user *User) cons() {
	if user.Password == "" {
		user.Password = "N/A"
	}
	if user.Email == "" {
		user.Email = "JohnDoe@gmail.com"
	}
	if user.Name == "" {
		user.Name = "John Doe"
	}
}

func (post *Post) def() {
	if post.Caption == "" {
		post.Caption = "Untitled"
	}
	if post.startTime == "" {
		post.startTime = string(time.Now().Format(time.RFC3339))
	}
	if post.endTime == "" {
		post.endTime = string(time.Now().Local().Add(time.Hour * time.Duration(1)).Format(time.RFC3339))
	}
	if &post.user.id == nil {
		post.user.id = primitive.NewObjectID()
		post.user.cons()
	}
}


func aCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json") 
	// create user
	fmt.Println("Creating user")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.cons()
	collection := Client.Database("AppointyDatabase").Collection("users")
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(result)
	fmt.Println(result)
	fmt.Println("User Id is: ", user.id)
}
func bGETUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("context-type","application/json")
	fmt.Println("The Requested User id is:- ",(r.URL.Query()["_id"][0]))
	CheckID := r.URL.Query()["_id"][0]
	checkID, _ := primitive.ObjectIDFromHex(CheckID)
	user, _ := CheckIDofUser(checkID)
	json.NewEncoder(w).Encode(user)
}

func cCreatePost(r http.ResponseWriter, rq *http.Request) {
	r.Header().Add("content-type", "application/json") 
	// create post
	fmt.Println("Creating post")
	var post Post
	_ = json.NewDecoder(rq.Body).Decode(&post)
	post.def()
	collection := Client.Database("AppointyDatabase").Collection("posts")
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(r).Encode(result)
	fmt.Println(result)
	fmt.Println("Post Id is: ", post.id)
}

func dGETPostByID(r http.ResponseWriter, rq *http.Request) {
	r.Header().Set("context-type","application/json")
	fmt.Println("The Requested post id is:- ",(rq.URL.Query()["_id"][0]))
	CheckID := rq.URL.Query()["_id"][0]
	checkID, _ := primitive.ObjectIDFromHex(CheckID)
	Posts, _ := CheckIDofPost(checkID)
	json.NewEncoder(r).Encode(Posts)
}
func eGETAllPosts(r http.ResponseWriter, rq *http.Request) {
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
		eGETAllPosts(w, r)
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
	
	http.HandleFunc("/users", UserHandler)
	http.HandleFunc("/Post", PostHandler)  
	http.HandleFunc("/users/Post", UserPostsHandler) 
	http.ListenAndServe(":8080",nil)
}