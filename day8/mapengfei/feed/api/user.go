package api

import "camp/feed/model/user"

//"camp/feed/api"

// type User struct {
// 	ID       bson.ObjectId `json:"id" bson:"_id"`
// 	Nick     string        `json:"nick" bson:"nick"`
// 	Sex      int           `json:"sex" bson:"sex"`
// 	Email    string        `json:"email" bson:"email"`
// 	Password string        `json:"password" bson:"password"`
// }

type ApiUser struct {
	*user.UserModel
}

func NewUser() *ApiUser {
	return &ApiUser{
		UserModel: &user.UserModel{},
	}
}

//type ApiUsers []*ApiUser
