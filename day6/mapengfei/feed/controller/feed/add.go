package feed

import (
	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
)

// AddReq 定义输入
type AddReq struct {
	Id  int64  `json:"id"`
	Txt string `json:"txt"`
}

// AddResp 定义输出
type AddResp struct {
}

// Add just for demo
// @postfilter("Boss")
func (feed *Feed) Add(w http.ResponseWriter, r *http.Request) {
	fun := "skel.Skel.Add"

	var addReq *AddReq
	if err := json.Unmarshal(feed.ReadBody(r), &addReq); err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		skel.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi := api.NewFeed()
	feedApi.Txt = addReq.Txt
	feedApi.Id = bson.NewObjectId()

	if err := service.NewSkel().Add(skelAPI); err != nil {
		clog.Error("%s skel.Add err: %v, req: %v", fun, err, addReq)
		skel.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &AddResp{}

	// 进行一些异步处理的工作

	return
}
