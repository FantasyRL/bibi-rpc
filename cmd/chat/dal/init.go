package dal

import (
	"bibi/cmd/chat/dal/cache"
	"bibi/cmd/chat/dal/db"
)

func Init() {
	cache.Init()
	db.Init()
}
