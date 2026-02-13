package middleware

import (
	capture "crm/gopkg/gins/capture"
	"crm/gopkg/utils"

	"github.com/gin-gonic/gin"
)

func RequestCapture() gin.HandlerFunc {
	return capture.RequestCapture(capture.Options{
		FilterPaths: []string{
			"/api/demo/put",
		},
	}, func(ctx *gin.Context, request *capture.Request) {
		ctx.Set(utils.ClientIPKey, request.ClientIP)
	}, func(ctx *gin.Context, capture *capture.Capture) {
		//log.Sugar().Info(ctx, "request", zap.Any("capture", capture))
	})
}
