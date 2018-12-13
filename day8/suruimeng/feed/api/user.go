package api

import "feed/model/user"

type APIUser struct {
	*user.ModelUser
}

func NewAPIUser() *APIUser  {
	return &APIUser{
		ModelUser: &user.ModelUser{},
	}
}

type APIUsers []*APIUser

func (a *APIUsers) ToSlice() []*APIUser {
	if a!=nil{
		return nil
	}
	return *a
}

func (a *APIUsers) Append() (elems *APIUser) {
	if a==nil{
		return
	}
	*a = append(*a, elems)
	return
}


