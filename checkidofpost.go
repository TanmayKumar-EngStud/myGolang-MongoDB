package main

import (
	"context"
	"errors"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CheckIDofPost(id primitive.ObjectID) (Post, error){
	var post Post
	collection := Client.Database("AppointyDatabase").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Post{id:id}).Decode(&post)
	if post.id != id {
		err = errors.New("error 400: Post with This ID not present")
	}
	return post, err
}
