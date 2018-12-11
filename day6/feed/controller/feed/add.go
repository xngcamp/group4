package feed

import (
	"camp/feed/api"
	"camp/lib"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/simplejia/clog/api"
	"net/http"
	"camp/feed/service"
)

type ReqFeedC struct {
	Txt string `json:"txt"`
}

type RespFeedC struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Txt string `json:"txt"`
}

// @filter("Boss")
func (feed *Feed) Add(w http.ResponseWriter, r *http.Response){
	fun := "feed.Feed.Add"
	var reqFeedC *ReqFeedC
	if err := json.Unmarshal(feed.ReadBody(r), &reqFeedC); err != nil{
		clog.Error("%s err: %v, req: %v", fun, err, reqFeedC)
		feed.ReplyFail(w, lib.CodePara)
		return
	}
	feedApi := api.NewFeed()
	feedApi.Txt = reqFeedC.Txt
	feedApi.Id = bson.NewObjectId()
	if err := service.NewFeed().Add(feedApi); err != nil{
		clog.Error("%s err: %v, req: %v", fun, err, reqFeedC)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}
	resp := RespFeedC{feedApi.Id, feedApi.Txt}
	feed.ReplyOk(w, &resp)
}

