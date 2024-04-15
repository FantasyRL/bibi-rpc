package dal

import (
	"bibi/rpc/video/dal/cache"
	"bibi/rpc/video/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
