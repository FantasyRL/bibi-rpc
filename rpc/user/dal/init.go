package dal

import (
	"bibi/rpc/user/dal/cache"
	"bibi/rpc/user/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
