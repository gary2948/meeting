package social

type Lctb_reply struct {
	Id              int64 `xorm:"pk autoincr"`
	Lc_weiboId      int64
	Lc_replyType    int
	Lc_replyContext string
	Lc_replyStatus  int
}
