package user

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (user *User) GetUserInfo(email string) (userApi *api.User,err error) {
	userModel := model.NewUser()
	userModel.Email = email
	if userModel, err = userModel.Get(); err != nil {
		return
	}
	userApi = (*api.User)(userModel)
	return
}