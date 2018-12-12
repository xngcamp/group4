package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

// Add 定义新增操作
func (feed *Feed) Add(feedApi *api.Feed) (err error) {
	feedModel := model.NewFeed()

	feedModel.Txt = feedApi.Txt
	feedModel.Id = feedApi.Id
	if err = feedModel.Add(); err != nil {
		return
	}
	return
}
