package gins

import (
	"crm/gopkg/services"
	"crm/gopkg/utils"
	"crm/internal/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusOK(ctx *gin.Context, result common.ServiceResult) {
	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func BadRequest(ctx *gin.Context, err error) {
	StatusFailed(ctx, http.StatusBadRequest, err)
}

func ServerError(ctx *gin.Context, err error) {
	// 使用 200 OK 作为 HTTP 状态码，业务错误码作为 Response Body 中的 code
	ctx.AbortWithStatusJSON(http.StatusOK, services.NewResult(ctx, 10001, err.Error(), nil))
}

func Unauthorized(ctx *gin.Context) {
	StatusFailed(ctx, http.StatusUnauthorized, nil)
}

func StatusFailed(ctx *gin.Context, code int, err error) {
	if utils.IsProduction() || err == nil {
		// http.StatusText(10001) 会返回空字符串，导致 panic 或空响应
		msg := http.StatusText(code)
		if msg == "" {
			msg = "Unknown Error"
		}
		ctx.AbortWithStatusJSON(code, services.NewResult(ctx, code, msg, nil))
		return
	}

	ctx.AbortWithStatusJSON(code, services.NewResult(ctx, code, err.Error(), nil))
}
