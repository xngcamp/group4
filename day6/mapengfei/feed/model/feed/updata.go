package feed

import (
	"github.com/globalsign/mgo/bson"
)

func (feed *Feed) Update() (err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Update(bson.M{"_id": feed.Id}, bson.M{"txt": feed.Txt})
	if err != nil {
		return
	}
	return
}
