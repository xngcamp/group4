package feed

import (
	"encoding/json"
	"net/http"

	"camp/feed/service"
	"camp/lib"

	"github.com/simplejia/clog/api"
)

// DelReq 定义输入
type DelReq struct {
	Id string `json:"id"`
}


// DelResp 定义输出
type DelResp struct {
}

// Del just for demo
// @postfilter("Boss")
func (feed *Feed) Del(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Del"

	var delReq *DelReq
	if err := json.Unmarshal(feed.ReadBody(r), &delReq); err != nil {
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	if err := service.NewFeed().Del(delReq.Id); err != nil {
		clog.Error("%s feed.Del err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &DelResp{}
	feed.ReplyOk(w, resp)


	return
}
