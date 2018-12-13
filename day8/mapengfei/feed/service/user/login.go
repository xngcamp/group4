package user

import (
	"camp/feed/api"
	"errors"
)

func (user *UserService) Login(userAPI *api.ApiUser) (err error) {
	if userAPI == nil {
		err = errors.New("输入内容为空")
		return
	}
	err = userAPI.Login()
	if err != nil {
		return
	}
	return
}
