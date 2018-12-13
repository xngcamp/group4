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
type AddReq struct {
	Nick string `json:"nick"`
	Sex string `json:"sex"`
	Email string `json:"email"`
	Password string `json:"password"`
}

// Regular 用于参数校验
func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}
	if addReq.Nick== "" || addReq.Email== "" || addReq.Password== "" ||addReq.Sex==""{

		return
	}

	ok = true
	return
}

// 返回给前端的结构体
type AddRsp struct {
}

// List just for demo
// @postfilter("Boss")
func (userController *UserControlller) AddUser(w http.ResponseWriter, r *http.Request) {
	// 用于日志
	fun := "Feed.user.AddUser"

	var addReq *AddReq
	err := json.Unmarshal(userController.ReadBody(r), &addReq)
	// 请求验证
	if err != nil || !addReq.Regular() {
		clog.Error("%v %v",addReq.Regular(),err!=nil)
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}
	//clog.Info("ssssssssssssss")
	// 要添加的关键词，生成API层的对象
	userAPI :=api.NewAPIUser()
	userAPI.ID = bson.NewObjectId()
	userAPI.Sex = addReq.Sex
	userAPI.Password=addReq.Password
	userAPI.Email=addReq.Email
	userAPI.Nick=addReq.Nick
	userAPI.Token= userAPI.ID
	//
	// Service业务逻辑，存储
	//err = service.NewBadWordService().Add(wordAPI)
	err=service.NewUserService().Add(userAPI)

	if err != nil {
		clog.Error("%s list err: %v, req: %v", fun, err, addReq)
		userController.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &AddRsp{}

	// 解决跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 返回OK
	userController.ReplyOk(w, resp)

	return
}