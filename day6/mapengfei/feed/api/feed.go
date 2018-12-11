package api

// Skel 用于示例

type Feed struct {
	Id  int64  `json:"id" bson:"_id"`
	Txt string `json:"txt" bson:"_txt"`
}

// NewSkel 生成skel对象
func NewFeed() *Feed {
	return &Feed{}
}
