package user

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	clog "github.com/simplejia/clog/api"
)

// 请求对象
type UpdateReq struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Nick     string        `json:"nick"`
	Sex      int           `json:"sex"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Token    bson.ObjectId `json:"token" bson:"token"`
}

// Regular 用于参数校验
func (updateReq *UpdateReq) Regular() (ok bool) {
	if updateReq == nil {
		return
	}
	if updateReq.Token == "" || updateReq.Data == nil {
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
func (userController *UserController) Update(w http.ResponseWriter, r *http.Request) {
	// 用于日志
	fun := "User.Apiuser.Update"
	var updateReq *UpdateReq
	err := json.Unmarshal(userController.ReadBody(r), &updateReq)

	// 请求验证
	if err != nil || !updateReq.Regular() {
		clog.Error("%v %v", updateReq.Regular(), err != nil)
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}
	// update数据接收有问题
	userAPI := api.NewUser()
	userAPI.Sex = updateReq.Sex
	userAPI.Token = updateReq.Token
	userAPI.Nick = updateReq.Nick
	userAPI.Email = updateReq.Email
	userAPI.Password = updateReq.Password

	// Service业务逻辑，存储
	//err = service.NewBadWordService().Add(wordAPI)
	err = service.NewUserService().Update(userAPI)

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
