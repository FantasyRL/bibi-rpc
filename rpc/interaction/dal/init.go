package dal

import (
	"bibi/rpc/interaction/dal/cache"
	"bibi/rpc/interaction/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
