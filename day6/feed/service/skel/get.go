package skel

import (
	"camp/feed/api"
	"camp/feed/model"
)

// Get 定义获取操作
func (skel *Skel) Get(id int64) (skelApi *api.Feed, err error) {
	skelModel := model.NewSkel()
	skelModel.ID = id
	if skelModel, err = skelModel.Get(); err != nil {
		return
	}

	if skelModel == nil {
		return
	}

	skelApi = (*api.Feed)(skelModel)

	return
}
