package feed

import (
	"camp/feed/model"
)

// Update 定义更新操作
func (feed *Feed) Update(id string,txt string) (err error) {

	feedModel := model.NewFeed()
	feedModel.Id = id
	feedModel.Txt = txt
	if err = feedModel.Update(); err != nil {
		return
	}



	return
}
