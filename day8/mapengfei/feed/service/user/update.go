package user

import (
	"camp/feed/api"
	"errors"
)

func (users *UserService) Update(userAPI *api.ApiUser) (err error) {
	// 是否无效调用
	if userAPI == nil {
		err = errors.New("Parameter userAPI is nil")
		return
	}

	// 添加到数据库中
	//err = wordAPI.BadWordModel.Add()
	err = userAPI.Update()
	if err != nil {
		return
	}
	return
}
