package user

import (
	"errors"
	"feed/api"
)

// 添加用户信息
func (users *UserService) Add(userAPI *api.APIUser) (err error) {
	// 是否无效调用
	if userAPI == nil {
		err = errors.New("Parameter userAPI is nil")
		return
	}
	// 添加到数据库中
	//err = wordAPI.BadWordModel.Add()
	err = userAPI.ModelUser.AddUser()
	if err != nil {
		return
	}
	return
}
