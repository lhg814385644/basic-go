package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lhg814385644/basic-go/webook/internal/repository"
	"github.com/lhg814385644/basic-go/webook/internal/repository/dao"
	"github.com/lhg814385644/basic-go/webook/internal/service"
	"github.com/lhg814385644/basic-go/webook/internal/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer()
	initUser(server, db)
	server.Run(":8080")
}

// 初始化web服务
func initWebServer() *gin.Engine {
	server := gin.Default()
	// 跨域处理中间件
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
	// Session中间件
	store := cookie.NewStore([]byte("secret"))
	// cookie的名字ssid
	server.Use(sessions.Sessions("ssid", store))
	// 登录校验
	login := &SignInMiddlewareBuilder{}
	server.Use(login.SignInCheck())
	return server
}

// 初始化数据库
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:bb123456@tcp(localhost:3306)/webook"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := dao.InitTable(db); err != nil {
		panic(err)
	}
	return db
}

func initUser(server *gin.Engine, db *gorm.DB) {
	userDao := dao.NewUserDao(db)
	userRepo := repository.NewUserRepo(userDao)
	userSvc := service.NewUserService(userRepo)
	hdl := web.NewUserHandler(userSvc)
	hdl.RegisterRoutes(server)
}

// 日志中间件
func logMid(ctx *gin.Context) {
	fmt.Printf("LogMid: %s\n", ctx.Request.Method)
}

type SignInMiddlewareBuilder struct {
}

// SignInCheck 登录校验中间件
func (*SignInMiddlewareBuilder) SignInCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 不需要校验的路径
		if ctx.Request.URL.Path == "/users/signIn" || ctx.Request.URL.Path == "/users/signUp" {
			return
		}
		session := sessions.Default(ctx)
		// 验证
		userID := session.Get("userID")
		if userID == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// TODO: 如果能取到，将该userID放到上下文context中向下传播
		ctx.Set(web.UserIDCTXKEY, userID)
	}
}
