package user

import (
	"camp/feed/api"
	"errors"
)

func (user *UserService) Add(userAPI *api.ApiUser) (err error) {
	//userModle := model.NewUser()
	if userAPI == nil {
		err = errors.New("Parameter userAPI is nil")
		return
	}
	//是否已经存在
	if EmailManager.Exists(userAPI.Email) {
		err = errors.New("The email has registered")
		return
	}
	////////
	err = userAPI.Add()
	// userModle.Id = userAPI.Id
	// userModle.Nick = userAPI.Nick
	// userModle.Sex = userAPI.Sex
	// userModle.Email = userAPI.Email
	// userModle.Password = userAPI.Password

	//	err = userModle.Add()
	//err = userAPI.
	if err != nil {
		return
	}
	//添加到注册记录中
	EmailManager.AddEmail(userAPI.Email)

	return

}
