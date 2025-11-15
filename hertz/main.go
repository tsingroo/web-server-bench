package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	hServer := server.Default(server.WithHostPorts(":8080"))

	hServer.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "pong")
	})
	hServer.GET("/hello/:name", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello "+c.Param("name"))
	})

	hServer.Spin()
}
