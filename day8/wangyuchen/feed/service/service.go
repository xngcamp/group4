/*
Package service 用于定义服务层代码。
只允许在这里添加对外暴露的接口
*/
package service

import (
	"camp/feed/service/user"
	"camp/feed/service/feed"
)

func NewFeed() *feed.Feed {
	return &feed.Feed{}
}

func NewUser() *user.User  {
	return &user.User{}
}
