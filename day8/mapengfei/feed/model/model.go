package model

import (
	"camp/feed/model/feed"
	"camp/feed/model/user"
)

func NewFeed() *feed.Feed {
	return new(feed.Feed)
}

func NewUser() *user.UserModel {
	return &user.UserModel{}
}
