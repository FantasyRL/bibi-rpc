package dal

import (
	"bibi/cmd/user/dal/cache"
	"bibi/cmd/user/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
