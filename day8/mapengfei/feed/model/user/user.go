package user

import (
	"camp/feed/mongo"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type UserModel struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Nick     string        `json:"nick"`
	Sex      int           `json:"sex"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Token    bson.ObjectId `json:"token" bson:"token"`
}

//type UsersModel []*UserModel

func (user *UserModel) Db() (db string) {
	return "skel"
}

func (user *UserModel) Table() (table string) {
	return "user"
}

func (user *UserModel) GetC() (c *mgo.Collection) { //模板方法
	db, table := user.Db(), user.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
