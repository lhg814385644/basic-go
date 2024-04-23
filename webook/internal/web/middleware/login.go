package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lhg814385644/basic-go/webook/internal/web"
	"net/http"
	"time"
)

// TODO: 登录中间件

// 更新时间
const updateTimeKey = "update_time"

type SignInMiddlewareBuilder struct {
}

// CheckLogin 检查登录中间件
func (s *SignInMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	// 注册一下GOB
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		// 不需要校验的路径
		if ctx.Request.URL.Path == "/users/signIn" || ctx.Request.URL.Path == "/users/signUp" {
			return
		}
		sess := sessions.Default(ctx)
		userID := sess.Get("userID")
		if userID == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// TODO: 刷新时间（间隔一段时间刷新）
		now := time.Now()
		// 尝试拿出上一次刷新时间
		val := sess.Get(updateTimeKey)
		lastUpdateTime, ok := val.(time.Time)
		if val == nil || !ok || now.Sub(lastUpdateTime) > time.Second*10 {
			sess.Set(updateTimeKey, now)
			sess.Set("userID", userID)
			sess.Options(sessions.Options{MaxAge: 60}) // 重新设置session过期时间（否则会永不过期）
			if err := sess.Save(); err != nil {
				// 日志收集
				fmt.Printf("session save err:%v\n", err)
			}
		}
		// 验证
		// TODO: 如果能取到，将该userID放到上下文context中向下传播
		ctx.Set(web.UserIDCTXKEY, userID)
	}
}
