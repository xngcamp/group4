package user

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"github.com/simplejia/clog/api"
	"net/http"
)

type UpdateReq struct {
	Token string `json:"token"`
	UpdateInfo UpdateInfo
}

// UpdateResp 定义输出
type UpdateResp struct {
	User *api.User `json:"feed,omitempty"`
}

type UpdateInfo struct {
	Email string `json:"email"`
	Nick string `json:"nick"`
	Sex int `json:"sex"`
	Password string `json:"password"`
}

// @postfilter("Boss")
func (user *User) Update(w http.ResponseWriter, r *http.Request) {

	var updateReq *UpdateReq

	if err := json.Unmarshal(user.ReadBody(r), &updateReq); err != nil{         // 获取前台传的数据
		clog.Error("err: %v, req: %v", err, updateReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}

	//token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
	//	func(token *jwt.Token) (interface{}, error) {
	//		return []byte("welcome"), nil
	//	})
	//
	//if err == nil {
	//	if token.Valid {
	//		if err := service.NewUser().Update(updateReq.UpdateInfo.Email,updateReq.UpdateInfo.Nick,updateReq.UpdateInfo.Password,updateReq.UpdateInfo.Sex); err != nil {
	//			clog.Error("err: %v, req: %v", err, updateReq)
	//			user.ReplyFail(w, lib.CodeSrv)
	//			return
	//		}
	//		resp := UpdateResp{}
	//		user.ReplyOk(w, resp)
	//
	//	} else {
	//		w.WriteHeader(http.StatusUnauthorized)
	//		fmt.Fprint(w, "Token is not valid")
	//	}
	//} else {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	fmt.Fprint(w, "Unauthorized access to this resource")
	//}



	if updateReq.Token == updateReq.UpdateInfo.Email && updateReq.Token != ""{
		if err := service.NewUser().Update(updateReq.UpdateInfo.Email,updateReq.UpdateInfo.Nick,updateReq.UpdateInfo.Password,updateReq.UpdateInfo.Sex); err != nil {
			clog.Error("err: %v, req: %v", err, updateReq)
			user.ReplyFail(w, lib.CodeSrv)
			return
		}

		userApi,err := service.NewUser().GetUserInfo(updateReq.UpdateInfo.Email)
		if err != nil {
			clog.Error("err: %v, req: %v", err, updateReq)
			user.ReplyFail(w, lib.CodeSrv)
			return
		}

		resp := &UpdateResp{
			User:userApi,
		}
		user.ReplyOk(w, resp)
	}


}
