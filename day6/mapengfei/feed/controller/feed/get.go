package feed

import (
	"net/http"

	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"

	"github.com/simplejia/clog/api"
)

// GetReq 定义输入
type GetReq struct {
	ID int64 `json:"id"`
}

// Regular 用于参数校验
func (getReq *GetReq) Regular() (ok bool) {
	if getReq == nil {
		return
	}

	if getReq.ID <= 0 {
		return
	}

	ok = true
	return
}

// GetResp 定义输出
type GetResp struct {
	Skel *api.Skel `json:"skel,omitempty"`
}

// Get just for demo
// @postfilter("Boss")
func (skel *Skel) GetAll(w http.ResponseWriter, r *http.Request) {
	fun := "skel.Skel.Get"

	feedApi, err := service.NewFeed().Get()
	if err != nil {
		clog.Error("%s skel.Get err: %v, req: %v", fun, err, getReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	feed.ReplyOk(w, feedApi)

	// 进行一些异步处理的工作
	//	go lib.Updates(skelApi, lib.GET, nil)

	return
}
