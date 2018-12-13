package user

import (
	"github.com/globalsign/mgo/bson"
)

func (users *ModelUser) AddUser() (err error) {
	//获得连接
	c:=users.GetC()
	defer c.Database.Session.Close()
	var result *ModelUser
	err=c.Find(bson.M{"email":users.Email}).One(&result)
	if err!=nil{
		return
	}
	if result != nil{
		return
	}else {
		err1 := c.Insert(users)
		if err1!=nil{
			return
		}
	}



	return
}
