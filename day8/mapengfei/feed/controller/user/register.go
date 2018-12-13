package user

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	clog "clog/api"
	"encoding/json"
	"net/http"
	"strings"
)

type AddReq struct {
	//	ID       bson.ObjectId `json:"id" bson:"_id"`
	Nick     string `json:"nick"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Regular 用于参数校验
func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}
	// if addReq.Email == "" || addReq.Nick == "" || addReq.Password == "" {
	// 	return
	// }
	// if addReq.Sex != 1 || addReq.Sex != 2 {
	// 	return
	// }
	ok = true
	return
}

// 返回给前端的结构体
type AddRsp struct {
	Nick     string `json:"nick"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// List just for demo
// @postfilter("Boss")
func (userController *UserController) Add(w http.ResponseWriter, r *http.Request) {
	// 用于日志
	fun := "user.Apiuser.Add"

	var addReq *AddReq
	// 请求验证

	err := json.Unmarshal(userController.ReadBody(r), &addReq)
	if err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		userController.ReplyFail(w, lib.CodePara)
		return
	}
	// 要添加的关键词，生成API层的对象
	userAPI := api.NewUser()
	//	userAPI.ID = bson.NewObjectId()
	userAPI.Nick = strings.Trim(addReq.Nick, " ")
	userAPI.Sex = addReq.Sex
	userAPI.Email = strings.Trim(addReq.Email, " ")
	userAPI.Password = strings.Trim(addReq.Password, " ")
	// Service业务逻辑，存储
	err = service.NewUserService().Add(userAPI)

	if err != nil {
		clog.Error("%s list err: %v, req: %v", fun, err, addReq)
		userController.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := userAPI

	// 解决跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 返回OK
	userController.ReplyOk(w, resp)

	return
}
