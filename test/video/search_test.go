package main

import (
	"bibi/kitex_gen/video"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSearchVideo(t *testing.T) {
	c.Convey("test search_video", t, func() {

		_, _, err := videoService.SearchVideo(&video.SearchVideoRequest{
			Param:   "test",
			PageNum: 1,
		})
		c.So(err, c.ShouldBeNil)
	})
}
