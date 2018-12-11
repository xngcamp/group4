package feed

import (
	"encoding/json"
	"net/http"

	"camp/feed/service"
	"camp/lib"
)

// DelReq 定义输入
type DelReq struct {
	ID int64 `json:"id"`
}

// Regular 用于参数校验
func (delReq *DelReq) Regular() (ok bool) {
	if delReq == nil {
		return
	}

	if delReq.ID <= 0 {
		return
	}

	ok = true
	return
}

// DelResp 定义输出
type DelResp struct {
}

// Del just for demo
// @postfilter("Boss")
func (feed *Feed) Del(w http.ResponseWriter, r *http.Request) {
	fun := "skel.Skel.Del"

	var delReq *DelReq
	if err := json.Unmarshal(feed.ReadBody(r), &delReq); err != nil || !delReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi, err := service.NewFeed().Get(delReq.Id)
	if err != nil {
		clog.Error("%s skel.Get err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	if feedApi == nil {
		detail := "skel not found"
		feed.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	if err := service.NewFeed().Del(delReq.ID); err != nil {
		clog.Error("%s skel.Del err: %v, req: %v", fun, err, delReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &DelResp{}
	skel.ReplyOk(w, resp)

	// 进行一些异步处理的工作

	return
}
