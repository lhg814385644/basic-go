package middleware

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lhg814385644/basic-go/webook/internal/web"
	"net/http"
	"strings"
	"time"
)

// TODO: 登录中间件JWT

type LoginJWTMiddlewareBuilder struct {
}

// CheckLogin 检查登录中间件
func (*LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	// 注册一下GOB
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		// 不需要校验的路径
		if ctx.Request.URL.Path == "/users/signIn" || ctx.Request.URL.Path == "/users/signUp" {
			return
		}
		authCode := ctx.GetHeader("Authorization")
		if authCode == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		strs := strings.Split(authCode, " ")
		if len(strs) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := strs[1]
		uc := &web.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, uc, func(token *jwt.Token) (interface{}, error) {
			return web.JWTKey, nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		expireTime := uc.ExpiresAt
		if expireTime.Before(time.Now()) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 每10s刷新一次
		if expireTime.Sub(time.Now()) < 50*time.Second {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			newToken, err := token.SignedString(web.JWTKey)
			if err != nil {
				// TODO: logger
			} else {
				ctx.Header("x-jwt-token", newToken)
			}
		}
		// todo: 断言Claims
		if claims, ok := token.Claims.(*web.UserClaims); ok {
			ctx.Set(web.UserIDCTXKEY, claims.Uid)
		}
	}
}
