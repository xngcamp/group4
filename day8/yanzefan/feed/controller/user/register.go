package user

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/simplejia/clog/api"
	"net/http"
)

type AddReq struct {
	Id string `json:"id"`
	Nick string `json:"nick"`
	Sex int `json:"sex"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type AddResp struct {
}

// @postfilter("Boss")
func (user *User)AddUser(w http.ResponseWriter,r *http.Request)  {
	fun := "user.User.Add"

	var addReq *AddReq     // 定义一个请求结构体
	if err := json.Unmarshal(user.ReadBody(r), &addReq) ; err != nil{   // 把前台json转换为AddReq结构体
		clog.Error("%s param err: %v, req: %v",fun,err,addReq)
		user.ReplyFail(w,lib.CodePara)
		return
	}

	if _,err := service.NewUser().GetUserInfo(addReq.Email); err != nil {
		clog.Error("用户已存在")
		user.ReplyFail(w,lib.CodeExist)
		return
	}
	id,err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}

	userApi := api.NewUser()
	userApi.Id = id.String()
	userApi.Email = addReq.Email
	userApi.Nick = addReq.Nick
	userApi.Password = addReq.Password
	userApi.Sex = addReq.Sex

	if err := service.NewUser().AddUser(userApi); err != nil {
		clog.Error("%s user.Add err: %v,req: %v",fun,err,addReq)
		user.ReplyFail(w,lib.CodePara)
		return
	}

	user.ReplyOk(w,userApi)

}

