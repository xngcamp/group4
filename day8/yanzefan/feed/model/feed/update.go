package feed

import "github.com/globalsign/mgo/bson"

// Update 根据ID更新数据
func (feed *Feed) Update() (err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	txt := bson.M{"$set":bson.M{"txt":feed.Txt}}          //
	err = c.UpdateId(feed.Id,txt)
	if err != nil {
		return
	}

	return
}
