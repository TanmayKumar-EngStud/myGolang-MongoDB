package main

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckIDofUser(id primitive.ObjectID)(User, error){
	var user User
	collection := Client.Database("AppointyDatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, User{id:id}).Decode(&user)
	if user.id != id {
		err = 
		errors.New("error 400: This User ID not present")
	}
	return user, err
}