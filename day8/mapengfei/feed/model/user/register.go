package user

func (user *UserModel) Add() (err error) {
	c := user.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(user)
	if err != nil {
		return
	}

	return
}
