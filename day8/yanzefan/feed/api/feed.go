package api

type Feed struct {
	Id string `json:"id" bson:"_id"`
	Txt string `json:"txt"`            // 内容
}

func NewFeed() *Feed {
	return &Feed{}
}
