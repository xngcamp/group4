package user

import (
	"camp/feed/model"
)

//func (user *User) Update(temp *api.User) (err error) {
//
//	userModel := model.NewUser()
//	userModel.Nick = temp.Nick
//	userModel.Sex = temp.Sex
//	userModel.Password = temp.Password
//	if err = userModel.Update(); err != nil {
//		return
//	}
//	return
//}

func (user *User) Update(email string,nick string,password string,sex int) (err error) {

	userModel := model.NewUser()
	userModel.Email = email
	userModel.Nick = nick
	userModel.Password = password
	userModel.Sex = sex

	if err = userModel.Update(); err != nil {
		return
	}



	return
}