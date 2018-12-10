package api

// Skel 用于示例
type Feed struct {
	Id string `json:"id" bson:"_id"`
	Txt string `json:"txt"`            // 内容
}

// NewSkel 生成skel对象
func NewFeed() *Feed {
	return &Feed{}
}
