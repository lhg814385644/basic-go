package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lhg814385644/basic-go/webook/internal/domain"
	"github.com/lhg814385644/basic-go/webook/internal/service"
	"net/http"
)

type UserHandler struct {
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
	userSvc          *service.UserService
}

func NewUserHandler(userSvc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRegexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		userSvc:          userSvc,
	}
}

const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	userIDKey            = "userID"
)

// RegisterRoutes 注册路由
func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	// REST 风格
	// server.POST("/user",h.SignUp)
	// server.PUT("/user",h.Edit)
	// server.GET("/user/:id",h.Profile)

	ug := server.Group("/users")

	ug.POST("/signUp", h.SignUp)
	ug.POST("/signIn", h.SignIn)
	ug.POST("/edit", h.Edit)
	ug.GET("/profile", h.Profile)
}

// SignUp 注册
func (h *UserHandler) SignUp(ctx *gin.Context) {
	// 内部结构体
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	req := &SignUpReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.String(http.StatusBadRequest, "bind error")
	}
	// 验证数据有效性
	isEmail, err := h.emailRegexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱格式错误")
		return
	}
	isPassword, err := h.passwordRegexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密码格式错误")
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "密码不一致")
		return
	}
	err = h.userSvc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "重复邮箱,请换一个邮箱")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "服务器异常,注册失败")
		return
	}
	ctx.String(http.StatusOK, "sign up ok")
}

// SignIn 登录
func (h *UserHandler) SignIn(ctx *gin.Context) {
	type SignInReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	req := &SignInReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.String(http.StatusBadRequest, "bind error")
	}
	user, err := h.userSvc.SignIn(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "无效的邮箱或密码")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "服务器异常")
		return
	}
	// TODO: set session
	session := sessions.Default(ctx)
	session.Options(sessions.Options{
		MaxAge: 120, // 过期时间(S)
	})
	session.Set(userIDKey, user.ID)
	err = session.Save()
	if err != nil {
		ctx.String(http.StatusOK, "服务器异常")
		return
	}
	ctx.String(http.StatusOK, "sign ok")
}

// Profile 查询用户信息
func (h *UserHandler) Profile(ctx *gin.Context) {
	type profile struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}
	sess := sessions.Default(ctx)
	if userID, ok := sess.Get(userIDKey).(int); ok {
		user, err := h.userSvc.Profile(ctx, userID)
		if err == service.ErrUserNotFound {
			ctx.String(http.StatusOK, "用户不存在")
			return
		}
		if err != nil {
			ctx.String(http.StatusOK, "服务器异常")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": profile{
				ID:    user.ID,
				Email: user.Email,
			},
		})
		return
	}
	ctx.String(http.StatusUnauthorized, "login failed")
}

// Edit 修改信息
func (h *UserHandler) Edit(ctx *gin.Context) {

}
