package login

import (
	"crm/gopkg/gins"
	service_login "crm/internal/service/login"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// DoLogin 用户登录
func (h *Handler) DoLogin(ctx *gin.Context) {
	var req LoginReq

	if err := ctx.Bind(&req); err != nil {
		gins.BadRequest(ctx, err)
		return
	}

	result, err := h.loginService.Login(ctx, req.UserName, req.Password)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	// 写入登录态 Cookie/Header
	if data, ok := result.GetData().(service_login.RespLogin); ok {
		token := data.Token
		if token != "" {
			expireHour := viper.GetInt("auth.jwt.expire_hour")
			ctx.SetCookie("token", token, expireHour*3600, "/", "", false, true)
			ctx.Header("Authorization", "Bearer "+token)
		}
	}

	gins.StatusOK(ctx, result)
	return
}
