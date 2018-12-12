package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (user *User) Get() (userRet *User, err error){
	c := user.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{"email": user.Email}).One(&userRet)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}
	return
}
