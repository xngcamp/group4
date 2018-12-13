package feed

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
)
//请求对象
type AddReq struct {
	Txt string `json:"txt"`
}
//参数验证
func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}

	if addReq.Txt == "" {
		return
	}

	return true
}
//返回请求
type AddResp struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Txt string `json:"txt"`
}

// @prefilter("Cors")
func (feed *Feed) Add(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Add"  //用于日志

	var addReq *AddReq  //请求验证及参数校验
	if err := json.Unmarshal(feed.ReadBody(r), &addReq); err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}
	//生成API层对象  存储
	feedApi := api.NewFeed()
	feedApi.Txt = addReq.Txt
	feedApi.Id = bson.NewObjectId()
	//业务逻辑
	err := service.NewFeed().Add(feedApi)
	if err != nil {
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := AddResp{feedApi.Id, feedApi.Txt}
	//解决跨域问题


	feed.ReplyOk(w, &resp)
}