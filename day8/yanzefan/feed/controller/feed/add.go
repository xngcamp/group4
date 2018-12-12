package feed

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"

	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"

	"github.com/simplejia/clog/api"
)

// AddReq 定义输入
type AddReq struct {
	Id string `json:"id"`
	Txt string `json:"txt"`
}


// AddResp 定义输出
type AddResp struct {
}

// Add just for demo
// @postfilter("Boss")
func (feed *Feed) Add(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Add"

	var addReq *AddReq
	if err := json.Unmarshal(feed.ReadBody(r), &addReq); err != nil{
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	feedApi := api.NewFeed()
	id,err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	feedApi.Id = id.String()
	feedApi.Txt = addReq.Txt


	if err := service.NewFeed().Add(feedApi); err != nil {
		clog.Error("%s feed.Add err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}


	feed.ReplyOk(w, feed)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.ADD, nil)

	return
}
