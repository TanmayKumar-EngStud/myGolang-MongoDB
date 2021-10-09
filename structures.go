package main

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	StartTime string `json:"StartTime,omitempty" bson:"StartTime,omitempty"`
	EndTime string `json:"EndTime,omitempty" bson:"EndTime,omitempty"`
}

// Setting Constructors for Default value of User and Post
func NewUser() User{
	return User{
		Name: "Username",
		Email: "Email",
		Password: "Password",
	}
}

func NewPost() Post{
	return Post{
		Caption: "Caption",
		ImgURL: "#Blank",
		StartTime: "--:--:--",
		EndTime: "--:--:--",
	}
} 

