package feed

func (feed *Feed) Add() (err error){
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(feed)
	if err != nil {
		return
	}
	return
}
