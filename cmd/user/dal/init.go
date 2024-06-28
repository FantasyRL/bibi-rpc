package dal

import (
	"bibi/cmd/user/dal/cache"
	"bibi/cmd/user/dal/db"
)

func Init() {
	db.InitMySQL()
	cache.Init()
}
