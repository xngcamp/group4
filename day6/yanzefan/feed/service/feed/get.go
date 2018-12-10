package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

// Get 定义获取操作
func (feed *Feed) Get(id string) (feedApi *api.Feed, err error) {
	feedModel := model.NewFeed()
	feedModel.Id = id
	if feedModel, err = feedModel.Get(); err != nil {
		return
	}

	if feedModel == nil {
		return
	}

	feedApi = (*api.Feed)(feedModel)

	return
}

func (feed *Feed) GetAll() (feedApi []*api.Feed, err error) {
	feedModel := model.NewFeed()


	if feedApi, err = feedModel.GetAll(); err != nil {
		return
	}

	return
}
