package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	// 静态路由
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	// 参数路由
	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "这是你传过来的名字 %s", name)
	})

	// 通配符路由
	server.GET("views/*.html", func(ctx *gin.Context) {
		path := ctx.Param(".html")
		ctx.String(http.StatusOK, "匹配上的值是 %s", path)
	})
	// Query获取查询参数 ex: /order?id=123
	server.GET("/order", func(ctx *gin.Context) {
		orderID := ctx.Query("id")
		ctx.String(http.StatusOK, "这是你传过来的订单号 %s", orderID)
	})
	server.Run(":8080")
}
