package user

import (
	"github.com/globalsign/mgo/bson"
)

func (user *UserModel) Update() (err error) {
	//获得连接
	c := user.GetC()
	defer c.Database.Session.Close()

	var ret *UserModel
	err = c.Find(bson.M{"email": user.Email}).One(&ret)
	if err != nil {
		return
	}
	if ret.Token == user.Token {
		err1 := c.Update(
			bson.M{"_id": user.ID},
			bson.M{"$set": bson.M{"nick": user.Nick, "sex": user.Sex, "email": user.Email, "password": user.Password, "token": user.Token}})
		if err1 != nil {
			return
		}
	}
	return
}
