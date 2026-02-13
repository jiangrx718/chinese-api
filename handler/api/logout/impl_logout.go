package logout

import (
	"crm/gopkg/gins"

	"github.com/gin-gonic/gin"
)

// DoLogout 用户退出
func (h *Handler) DoLogout(ctx *gin.Context) {
	loginUserId, _ := ctx.Get("user_id")
	result, err := h.logoutService.Logout(ctx, loginUserId.(string))
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	// 写入登录态 Cookie/Header
	token := ""
	ctx.SetCookie("token", token, 0, "/", "", false, true)
	ctx.Header("Authorization", "Bearer "+token)

	gins.StatusOK(ctx, result)
	return
}
