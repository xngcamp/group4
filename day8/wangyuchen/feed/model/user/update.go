package user

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
)

func (user *User) Update() (err error){
	c := user.GetC()
	defer c.Database.Session.Close()

	fmt.Printf("%v\n", user)
	err = c.Update(bson.M{"_id": user.Id}, bson.M{
		"nick": user.Nick,
		"sex": user.Sex,
		"email": user.Email,
		"password": user.Password,
	})
	return
}

