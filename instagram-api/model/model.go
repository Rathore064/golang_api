package model

import (
	"time"
)

type User struct {
	Id       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

//post structure
type Post struct {
	Id        string    `json:"id" bson:"_id"`
	Caption   string    `json:"caption" bson:"caption"`
	Image_URL string    `json:"image_url" bson:"image_url"`
	TimeStamp time.Time `json:"timestamp" bson:"timestamp"`
	UserID    string    `json:"userid" bson:"_userid"`
}
