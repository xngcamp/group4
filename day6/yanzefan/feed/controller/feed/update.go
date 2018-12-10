package feed

import (
	"encoding/json"
	"net/http"

	"camp/lib"
	"camp/feed/api"
	"camp/feed/service"

	"github.com/simplejia/clog/api"
)

// UpdateReq 定义输入
type UpdateReq struct {
	Id string `json:"id"`
	Txt string `json:"txt"`
}

// UpdateResp 定义输出
type UpdateResp struct {
	Feed *api.Feed `json:"feed,omitempty"`
}


// @postfilter("Boss")
func (feed *Feed) Update(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Update"

	var updateReq *UpdateReq
	if err := json.Unmarshal(feed.ReadBody(r), &updateReq); err != nil{         // 获取前台传的数据
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}


	if err := service.NewFeed().Update(updateReq.Id,updateReq.Txt); err != nil {
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	feedApi,err := service.NewFeed().Get(updateReq.Id)
	if err != nil {
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &UpdateResp{
		Feed: feedApi,
	}
	feed.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.UPDATE, nil)

	return
}
