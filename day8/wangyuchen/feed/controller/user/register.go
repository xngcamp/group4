package user

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

type AddReq struct {
	Nick string `json:"nick"`
	Sex int `json:"sex"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type AddResp struct {
}

// @postfilter("Boss")
func (user *User) Register(w http.ResponseWriter, r *http.Request){
	fun := "user.User.Register"
	var addReq *AddReq

	if err := json.Unmarshal(user.ReadBody(r), &addReq); err != nil{
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}

	if _, err := service.NewUser().Get(addReq.Email); err != nil{
		clog.Error("该邮箱已被注册")
		user.ReplyFail(w, lib.CodeExist)
		return
	}

	userApi := api.NewUser()
	userApi.Id = bson.NewObjectId()
	userApi.Nick = addReq.Nick
	userApi.Sex = addReq.Sex
	userApi.Email = addReq.Email
	userApi.Password = addReq.Password

	if err := service.NewUser().Add(userApi); err != nil{
		clog.Error("%s user.Register err: %v, req: %v", fun, err, addReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}
	resp := AddResp{}
	user.ReplyOk(w, &resp)

}
