package user

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (user *User) AddUser(userApi *api.User) (err error) {
	userModel := model.NewUser()
	userModel.Id = userApi.Id
	userModel.Sex = userApi.Sex
	userModel.Nick = userApi.Nick
	userModel.Email = userApi.Email
	userModel.Password = userApi.Password
	if err = userModel.Add(); err != nil {
		return
	}
	return
}