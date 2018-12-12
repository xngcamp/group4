package user

import (
	"github.com/globalsign/mgo/bson"
)

func (userModel *ModelUser) Update() ( err error)  {
	//获得连接
	c:=userModel.GetC()
	defer c.Database.Session.Close()

	var result *ModelUser
	err=c.Find(bson.M{"email":userModel.Email}).One(&result)
	if err!=nil{
		return
	}
	if result.Token==userModel.Token{
		err1 := c.Update(
			bson.M{"_id":userModel.ID },
			bson.M{"$set": bson.M{ "nick": userModel.Nick, "sex": userModel.Sex,"email":userModel.Email,"password":userModel.Password,"token":userModel.Token }})
		if err1!=nil {
			return
		}
	}
	return
}
