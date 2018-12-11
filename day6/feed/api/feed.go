package api

import "github.com/globalsign/mgo/bson"

// Feed 用于示例
type Feed struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Txt string `json:"txt"`
}

// NewFeed 生成feed对象
func NewFeed() *Feed {
	return &Feed{}
}
