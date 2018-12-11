package skel

import (
	"camp/feed/api"
	"camp/feed/model"
)

// Add 定义新增操作
func (skel *Skel) Add(skelApi *api.Feed) (err error) {
	skelModel := model.NewSkel()
	skelModel.ID = skelApi.ID
	if err = skelModel.Add(); err != nil {
		return
	}

	return
}
