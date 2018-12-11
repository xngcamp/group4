package feed

import (
	"camp/feed/api"
	"camp/feed/mongo"
	"github.com/globalsign/mgo"
)

type Feed api.Feed

func (feed *Feed) DB() string{
	return "feed"
}

func (feed *Feed) Table() string{
	return "feed"
}

func (feed *Feed) GetC() (c *mgo.Collection){
	db, table := feed.DB(), feed.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}