package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lhg814385644/basic-go/webook/internal/web"
	"strings"
	"time"
)

func main() {
	hdl := web.NewUserHandler()
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// 自定义中间件
	server.Use(logMid)
	hdl.RegisterRoutes(server)

	server.Run(":8080")
}

// 日志中间件
func logMid(ctx *gin.Context) {
	fmt.Printf("LogMid: %s\n", ctx.Request.Method)
}
