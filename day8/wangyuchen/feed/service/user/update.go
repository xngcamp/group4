package user

import (
	"camp/feed/model"
)

func (user *User) Update(nick string, sex int, email string, password string) (err error){
	userModel := model.NewUser()
	userModel.Nick = nick
	userModel.Sex = sex
	userModel.Email = email
	userModel.Password = password

	err = userModel.Update()
	return
}
