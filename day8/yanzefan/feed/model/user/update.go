package user

import "github.com/globalsign/mgo/bson"

func (user *User) Update() (err error) {
	c := user.GetC()
	defer c.Database.Session.Close()
	selector := bson.M{"email": user.Email}
	data := bson.M{"$set":bson.M{"nick":user.Nick, "sex":user.Sex, "password":user.Password,}}
	err = c.Update(selector,data)
	if err != nil {
		return
	}

	return
}
