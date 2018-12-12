package api

type User struct {
	Id string `json:"id" bson:"_id"`
	Email string `json:"email"`
	Nick string `json:"nick"`
	Sex int `json:"sex"`
	Password string `json:"password"`
}

func NewUser() *User {
	return &User{}
}
