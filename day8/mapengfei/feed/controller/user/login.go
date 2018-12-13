package user

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	clog "clog/api"
	"encoding/json"
	"net/http"
)

// 请求对象
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Regular 用于参数校验
func (loginReq *LoginReq) Regular() (ok bool) {
	if loginReq == nil {
		return
	}
	if loginReq.Email == "" || loginReq.Password == "" {
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
func (userController *UserController) Login(w http.ResponseWriter, r *http.Request) {
	// 用于日志
	fun := "User.Apiuser.Login"

	var loginReq *LoginReq
	err := json.Unmarshal(userController.ReadBody(r), &loginReq)
	// 请求验证
	if err != nil || !loginReq.Regular() {
		clog.Error("%v %v", loginReq.Regular(), err != nil)
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}
	// 要添加的关键词，生成API层的对象
	userAPI := api.NewUser()
	userAPI.Password = loginReq.Password
	userAPI.Email = loginReq.Email

	// Service业务逻辑，存储
	//err = service.NewBadWordService().Add(wordAPI)
	err = service.NewUserService().Login(userAPI)

	if err != nil {
		clog.Error("%s list err: %v, req: %v", fun, err, loginReq)
		userController.ReplyFail(w, lib.CodeSrv)
		return
	}
	//resp := userAPI
	// 解决跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 返回OK
	userController.ReplyOk(w, err)
	return
}
