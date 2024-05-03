package ws

//go:generate msgp -tests=false -o=reply_msgp.go -io=false
type ReplyMsg struct {
	Code    int64  `json:"code" msg:"target"`
	From    int64  `json:"from" msg:"from"`
	Content string `json:"content" msg:"content"`
}
