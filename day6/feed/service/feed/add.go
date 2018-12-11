package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (feed *Feed) Add(feedApi *api.Feed) (err error) {
	feedModel := model.NewFeed()
	feedModel.Id = feedApi.Id
	feedModel.Txt = feedApi.Txt
	if err = feedModel.Add(); err != nil {
		return
	}
	return
}
