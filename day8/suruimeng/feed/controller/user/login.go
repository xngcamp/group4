package user

import (
	"camp/lib"
	"encoding/json"
	"feed/api"
	"feed/service"
	"github.com/simplejia/clog/api"
	"net/http"
)
// 请求对象
type LoginReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
// Regular 用于参数校验
func (loginReq *LoginReq) Regular() (ok bool) {
	if loginReq == nil {
		return
	}
	if  loginReq.Email== "" || loginReq.Password== "" {

		return
	}

	ok = true
	return
}

// 返回给前端的结构体
type LoginRsp struct {
	Token string `json:"token"`
}

// List just for demo
// @postfilter("Boss")
func (userController *UserControlller) Login(w http.ResponseWriter, r *http.Request) {
	// 用于日志
	fun := "Feed.user.AddUser"

	var loginReq *LoginReq
	err := json.Unmarshal(userController.ReadBody(r), &loginReq)
	// 请求验证
	if err != nil || !loginReq.Regular() {
		clog.Error("%v %v",loginReq.Regular(),err!=nil)
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}
	// 要添加的关键词，生成API层的对象
	userAPI :=api.NewAPIUser()
	userAPI.Password=loginReq.Password
	userAPI.Email=loginReq.Email

	// Service业务逻辑，存储
	//err = service.NewBadWordService().Add(wordAPI)
	err=service.NewUserService().Login(userAPI)

	if err != nil {
		clog.Error("%s list err: %v, req: %v", fun, err, loginReq)
		userController.ReplyFail(w, lib.CodeSrv)
		return
	}

	// 解决跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 返回OK
	userController.ReplyOk(w,userAPI.Token )
	return
}
