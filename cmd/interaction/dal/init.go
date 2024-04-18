package dal

import (
	"bibi/cmd/interaction/dal/cache"
	"bibi/cmd/interaction/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
