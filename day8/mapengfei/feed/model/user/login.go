package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (user *UserModel) Login() (err error) {
	c := user.GetC()
	defer c.Database.Session.Close()
	var ret *UserModel
	err = c.Find(bson.M{"email": user.Email}).One(&ret)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
