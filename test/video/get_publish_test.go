package main

import (
	"bibi/kitex_gen/video"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestListUserVideo(t *testing.T) {
	Convey("list user published videos", t, func() {
		_, _, err := videoService.ListVideo(&video.ListUserVideoRequest{
			UserId:  1,
			PageNum: 1,
		})
		Convey("err should be nil", func() {
			So(err, ShouldEqual, nil)
		})
	})
}
