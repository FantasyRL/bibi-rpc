package main

import (
	"bibi/cmd/api/biz/rpc"
	"bibi/cmd/video/dal"
	"bibi/cmd/video/dal/cache"
	"bibi/cmd/video/service"
	"bibi/config"
	"context"
	c "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

var videoService *service.VideoService

func TestMain(m *testing.M) {
	videoService = service.NewVideoService(context.TODO())
	config.InitTest()
	dal.Init()
	cache.Init()
	rpc.Init()
	c.SuppressConsoleStatistics()
	result := m.Run()
	c.PrintConsoleStatistics()
	os.Exit(result)

}
