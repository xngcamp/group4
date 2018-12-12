package user

import (
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
)

func (userModel *ModelUser) Login() ( err error)  {
	//获得连接
	c:=userModel.GetC()
	defer c.Database.Session.Close()
	clog.Info("email:%v",userModel.Email)
	var result *ModelUser
	err=c.Find(bson.M{"email":userModel.Email}).One(&result)
	clog.Info("email:%v",result)
	if err!=nil{
		return
	}
	if result.Password==userModel.Password{
		userModel.Token=result.Token
	}
	return
}
