package user

import (
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/simplejia/clog/api"
	"net/http"
	"time"
)

type LoginReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

// @postfilter("boss")
func (user *User) Login(w http.ResponseWriter, r *http.Request){
	fun := "user.User.Login"
	var loginReq *LoginReq

	if err := json.Unmarshal(user.ReadBody(r), &loginReq); err != nil{
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}

	userApi, err := service.NewUser().Get(loginReq.Email)
	if err != nil {
		user.ReplyFail(w, lib.CodeFound)
		return
	}
	if userApi.Password != loginReq.Password{
		http.Error(w, "密码错误", 404)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["email"] = userApi.Email
	claims["id"] = userApi.Id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenSting, err := token.SignedString([]byte("welcome"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		clog.Error("err: %v", err)
		fmt.Fprintln(w, "Error while signing the token")
	}

	resp := &LoginResp{tokenSting}
	user.ReplyOk(w,resp)
}
