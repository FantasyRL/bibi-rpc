package cache

import (
	"bibi/config"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var rMessage *redis.Client

func Init() {
	rMessage = redis.NewClient(&redis.Options{
		Addr:       config.Redis.Addr,
		ClientName: "Message",
		DB:         3,
	})
}
func i64ToStr(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}
