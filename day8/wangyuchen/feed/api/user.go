package api

import "github.com/globalsign/mgo/bson"

type User struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Nick string `json:"nick"`
	Sex int `json:"sex"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func NewUser() *User{
	return &User{}
}