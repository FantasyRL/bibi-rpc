package dal

import (
	"bibi/cmd/follow/dal/cache"
	"bibi/cmd/follow/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
