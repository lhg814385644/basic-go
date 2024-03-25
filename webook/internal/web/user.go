package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		emailRegexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
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
	ctx.String(http.StatusOK, "sign up ok")
}

// SignIn 登录
func (h *UserHandler) SignIn(ctx *gin.Context) {

}

// Profile 查询用户信息
func (h *UserHandler) Profile(ctx *gin.Context) {

}

// Edit 修改信息
func (h *UserHandler) Edit(ctx *gin.Context) {

}
