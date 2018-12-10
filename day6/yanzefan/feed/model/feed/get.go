package feed

import (
	"camp/feed/api"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Get 定义获取操作
func (feed *Feed) Get() (feedRet *Feed, err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Find(feed).One(&feedRet)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}

func (feed *Feed) GetAll() (feedRet []*api.Feed, err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{}).All(&feedRet)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
