package feed

func (feed *Feed) Add() (err error) {
	ad := feed.GetC()
	defer ad.Database.Session.Close()

	err = ad.Insert(feed)
	return
}
