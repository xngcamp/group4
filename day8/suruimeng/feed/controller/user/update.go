package user

import (
	"camp/lib"
	"encoding/json"
	"feed/api"
	"feed/service"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)

// 请求对象
type UpdateReq struct {
	Token bson.ObjectId `json:"token"`
	Data []string `json:"update_info"`
}
// Regular 用于参数校验
func (updateReq *UpdateReq) Regular() (ok bool) {
	if updateReq == nil {
		return
	}
	if  updateReq.Token== "" || updateReq.Data== nil {
		return
	}

	ok = true
	return
}

// 返回给前端的结构体
type UpdateRsp struct {
}

// List just for demo
// @postfilter("Boss")
func (userController *UserControlller) Update(w http.ResponseWriter, r *http.Request) {
	// 用于日志
	fun := "Feed.user.Update"
    clog.Info("%v_____ ",r)
	var updateReq *UpdateReq
	err := json.Unmarshal(userController.ReadBody(r), &updateReq)

	// 请求验证
	if err != nil || !updateReq.Regular() {
		clog.Error("%v %v",updateReq.Regular(),err!=nil)
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}
	// update数据接收有问题
	userAPI :=api.NewAPIUser()
	userAPI.Password=updateReq.Data[3]
	userAPI.Email=updateReq.Data[2]
	userAPI.Token=updateReq.Token
	userAPI.Sex=updateReq.Data[1]
	userAPI.Nick=updateReq.Data[0]

	// Service业务逻辑，存储
	//err = service.NewBadWordService().Add(wordAPI)
	err=service.NewUserService().Update(userAPI)

	if err != nil {
		clog.Error("%s list err: %v, req: %v", fun, err, updateReq)
		userController.ReplyFail(w, lib.CodeSrv)
		return
	}

	// 解决跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 返回OK
	userController.ReplyOk(w, userAPI.Token)
	return
}