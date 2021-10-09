package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func testMain(t *testing.T) { // Creating the base in both the tables
	var user User
	var post Post
	user.Name = "Tanmay Kumar"
	user.Email = "tanmay44a@gmail.com"
	user.Password = "Random"
	post.Caption = "Caption"
	post.ImgURL = "https://picsum.photos/200/300"
	post.StartTime = "12:05:12"
	post.EndTime = "12:12:40"
	bytesRepresentation, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8080/Users", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("NO response")
	}
}
func testTrail2(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/Users/6161a79f9d1ed5032a900ed5")
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("NO response")
	}
}

func testTrail3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:8080/Users/6161c222115192de123b3789")
	}
}

func testTrail4(b *testing.B) {
	var user User
	var post Post
	user.Name = "Tanmay Kumar"
	user.Email = "tanmay44a@gmail.com"
	user.Password = "RandomPass"
	post.Caption = "Caption"
	post.ImgURL = "https://picsum.photos/200/300"
	post.StartTime = "12:05:12"
	post.EndTime = "12:12:40"
}

func testTrail5(b *testing.B) {
	var user User
	var post Post
	user.Name = "Tanmay Kumar"
	user.Email = "tanmay44a@gmail.com"
	user.Password = "RandomPass"
	post.Caption = "Caption"
	post.ImgURL = "https://picsum.photos/200/300"
	post.StartTime = "12:05:12"
	post.EndTime = "12:12:40"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:8080/")
	Client, _ = mongo.Connect(ctx, clientOptions)
}