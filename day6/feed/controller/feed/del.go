package feed

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"net/http"

	"camp/lib"
	"github.com/simplejia/clog/api"
)

// DelReq 定义输入
type DelReq struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
}

// Regular 用于参数校验
func (delReq *DelReq) Regular() (ok bool) {
	if delReq == nil {
		return
	}
	ok = true
	return
}

// DelResp 定义输出
type DelResp struct {
}

// filter("Boss")
func (feed *Feed) Del(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Del"

	var delReq *DelReq
	if err := json.Unmarshal(feed.ReadBody(r), &delReq); err != nil || !delReq.Regular() {
		clog.Error("%s err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

}
