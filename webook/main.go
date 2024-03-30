package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lhg814385644/basic-go/webook/internal/repository"
	"github.com/lhg814385644/basic-go/webook/internal/repository/dao"
	"github.com/lhg814385644/basic-go/webook/internal/service"
	"github.com/lhg814385644/basic-go/webook/internal/web"
	"github.com/lhg814385644/basic-go/webook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	// Session中间件基于cookie（也就是说你的userID是存在cookie里面的）
	//store := cookie.NewStore([]byte("secret"))
	// TODO:采用基于内存实现,第一个参数是authentication key(用于认证)，第二个是 encryption key(用于加密) 最好都是是32位或64位
	// store := memstore.NewStore([]byte("authentication key"), []byte("encryption key"))
	// TODO:基于redis存储
	store, err := redis.NewStore(3, "tcp", "localhost:6379", "",
		[]byte("vTTI7yzD0O3H7zYx4vKqda0IBKrKN5a8"),
		[]byte("vTTI7yzD0O3H7zYx4vKqda0IBKrKN5a9"))
	if err != nil {
		panic(err)
	}
	// cookie的名字ssid
	server.Use(sessions.Sessions("ssid", store))
	// 登录校验
	login := &middleware.SignInMiddlewareBuilder{}
	server.Use(login.CheckLogin())
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
