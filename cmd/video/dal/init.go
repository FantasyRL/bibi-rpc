package dal

import (
	"bibi/cmd/video/dal/cache"
	"bibi/cmd/video/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
