package user

import (
	"camp/feed/api"
	"camp/feed/mongo"
	"github.com/globalsign/mgo"
)

type User api.User

func (user *User) Db() (db string) {
	return "feed"
}

func (user *User) Table() (table string)  {
	return "user"
}

func (user *User) GetC() (c *mgo.Collection) {
	db,table := user.Db(),user.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}