package service

import (
	"camp/feed/service/feed"
	"camp/feed/service/user"
)

func NewFeed() *feed.Feed {
	return new(feed.Feed)
}

func NewUserService() *user.UserService {
	return &user.UserService{}
}
