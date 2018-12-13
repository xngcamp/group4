package user

import (
	"feed/mongo"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type ModelUser struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	Nick string `json:"nick" bson:"nick"`
	Sex string `json:"sex" bson:"sex"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Token bson.ObjectId `json:"token" bson:"token"`
}

type ModelUsers []*ModelUser


func (users *ModelUser) Db() (db string) {
	return fmt.Sprint("feed")
}

func (users *ModelUser) Table() (table string) {
	return fmt.Sprint("users")
}

// GetC 返回db col
func (users *ModelUser) GetC() (c *mgo.Collection) {
	db, table := users.Db(), users.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}

// util
func (users *ModelUsers) ToSlice() []*ModelUser {
	if users == nil {
		return nil
	}

	return *users
}

func (users *ModelUsers) Append(elems ...*ModelUser) {
	if users == nil {
		return
	}

	*users = append(*users, elems...)
	return
}
