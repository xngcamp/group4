package feed

import (
	"camp/feed/service"
	"camp/lib"
	"net/http"
)

// @prefilter("Cors")
func (feed *Feed) GetAll(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.GetAll"

	feedApi, err := service.NewFeed().GetAll()
	if err != nil {
		clog.Error("%s feed.Add err: %v", fun, err)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	feed.ReplyOk(w, feedApi)
}
