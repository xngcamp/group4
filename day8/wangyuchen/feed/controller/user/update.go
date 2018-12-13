package user

import (
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"github.com/simplejia/clog/api"
	"net/http"
)

type UpdateReq struct {
	Token string `json:"token"`
	UpdateInfo
}

type UpdateInfo struct {
	Nick     string `json:"nick"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateResp struct {
}

// @postFilter("boss")
func (user *User) Update(w http.ResponseWriter, r *http.Request){
	fun := "user.User.Update"
	var updateReq *UpdateReq

	if err := json.Unmarshal(user.ReadBody(r), &updateReq); err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}

	if updateReq.Token == updateReq.UpdateInfo.Email {
		if err := service.NewUser().Update(updateReq.UpdateInfo.Nick, updateReq.UpdateInfo.Sex, updateReq.UpdateInfo.Email, updateReq. UpdateInfo.Password); err != nil {
			clog.Error("err: %v, req: %v", err, updateReq)
			user.ReplyFail(w, lib.CodeSrv)
			return
		}

		_,err := service.NewUser().Get(updateReq.UpdateInfo.Email)
		if err != nil {
			clog.Error("err: %v, req: %v", err, updateReq)
			user.ReplyFail(w, lib.CodeSrv)
			return
		}

		resp := &UpdateResp{}
		user.ReplyOk(w, resp)
	}


}
