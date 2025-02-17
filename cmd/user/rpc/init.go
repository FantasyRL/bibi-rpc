package rpc

import (
	"bibi/cmd/user/vector/vector"
)

var (
	pictureClient vector.VectorClient
)

func Init() {
	InitConvertRPC()
}
