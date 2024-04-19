package main

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/route"
	"testing"
)

func newTestEngine() *route.Engine {
	opt := config.NewOptions([]config.Option{})
	return route.NewEngine(opt)
}
func TestPerformRequest(t *testing.T) {
	router := newTestEngine()
	router.PUT("/hey/:user", func(ctx context.Context, c *app.RequestContext) {
		user := c.Param("user")
		if string(c.Request.Body()) == "1" {
			assert.DeepEqual(t, "close", c.Request.Header.Get("Connection"))
			c.Response.SetConnectionClose()
			c.JSON(consts.StatusCreated, map[string]string{"hi": user})
		} else if string(c.Request.Body()) == "" {
			c.AbortWithMsg("unauthorized", consts.StatusUnauthorized)
		} else {
			assert.DeepEqual(t, "PUT /hey/dy HTTP/1.1\r\nContent-Type: application/x-www-form-urlencoded\r\nTransfer-Encoding: chunked\r\n\r\n", string(c.Request.Header.Header()))
			c.String(consts.StatusAccepted, "body:%v", string(c.Request.Body()))
		}
	})
	router.GET("/her/header", func(ctx context.Context, c *app.RequestContext) {
		assert.DeepEqual(t, "application/json", string(c.GetHeader("Content-Type")))
		assert.DeepEqual(t, 1, c.Request.Header.ContentLength())
		assert.DeepEqual(t, "a", c.Request.Header.Get("dummy"))
	})

	w := ut.PerformRequest(router, "GET", "/hey/hertz", &ut.Body{bytes.NewBufferString("1"), 1},
		ut.Header{"Connection", "close"})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "{\"hi\":\"hertz\"}", string(resp.Body()))
}
