package feed

import (

	"net/http"

	"camp/lib"

	"camp/feed/service"

	"github.com/simplejia/clog/api"
)

// GetReq 定义输入
type GetReq struct {

}



// GetResp 定义输出
type GetResp struct {

}

// Get just for demo
// @postfilter("Boss")
func (feed *Feed) GetAll(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Get"

	var getReq *GetReq

	feedApi, err := service.NewFeed().GetAll()
	if err != nil {
		clog.Error("%s feed.Get err: %v, req: %v", fun, err, getReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	feed.ReplyOk(w, feedApi)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.GET, nil)

	return
}
