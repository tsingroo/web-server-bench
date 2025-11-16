package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.Writer.WriteString("pong")
	})
	engine.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.Writer.WriteString("my name is " + name)
	})

	engine.Run(":8080")
}
