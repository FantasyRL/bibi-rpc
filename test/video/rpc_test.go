package main

import (
	"bibi/kitex_gen/interaction/interactionhandler"
	"bibi/kitex_gen/user/userhandler"
	"bibi/pkg/constants"

	"github.com/cloudwego/kitex/client"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRPC(t *testing.T) {
	c.Convey("rpc user", t, func() {
		_, err := userhandler.NewClient(constants.UserServiceName,
			client.WithMuxConnection(constants.MuxConnection),
			client.WithHostPorts("0.0.0.0:114514"))

		c.Convey("err should be nil", func() {
			c.So(err, c.ShouldEqual, nil)
		})
	})
	c.Convey("rpc interaction", t, func() {
		_, err := interactionhandler.NewClient(constants.InteractionServiceName,
			client.WithMuxConnection(constants.MuxConnection),
			client.WithHostPorts("0.0.0.0:114515"))

		c.Convey("err should be nil", func() {
			c.So(err, c.ShouldEqual, nil)
		})
	})
}
