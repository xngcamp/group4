package user

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (user *User) Add(userApi *api.User) (err error) {
	userModel := model.NewUser()
	userModel.Id = userApi.Id
	userModel.Nick = userApi.Nick
	userModel.Sex = userApi.Sex
	userModel.Email = userApi.Email
	userModel.Password = userApi.Password

	if err = userModel.Add(); err != nil {
		return
	}
	return
}
