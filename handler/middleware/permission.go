package middleware

import (
	"crm/gopkg/gins"
	"crm/internal/g"
	"crm/internal/model"
	"crm/internal/service/permission"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v, ok := ctx.Get("user_id")
		if !ok {
			gins.Unauthorized(ctx)
			ctx.Abort()
			return
		}
		userID, _ := v.(string)
		if userID == "" {
			gins.Unauthorized(ctx)
			ctx.Abort()
			return
		}

		svc := permission.NewService()
		res, err := svc.PermissionFind(ctx, userID)
		if err != nil {
			gins.ServerError(ctx, err)
			ctx.Abort()
			return
		}

		info, ok := res.GetData().(permission.RespPermissionFindInfo)
		if !ok {
			gins.ServerError(ctx, nil)
			ctx.Abort()
			return
		}
		if info.IsSuper == model.IsSuperOn {
			ctx.Next()
			return
		}

		current := ctx.FullPath()
		p, _ := g.CRMPermission.WithContext(ctx).Where(
			g.CRMPermission.PermissionURL.Eq(current),
			g.CRMPermission.Status.Eq(model.StatusOn),
		).Take()
		if p == nil {
			gins.StatusFailed(ctx, http.StatusForbidden, nil)
			ctx.Abort()
			return
		}

		allowed := make(map[string]struct{}, len(info.PermissionIds))
		for _, id := range info.PermissionIds {
			allowed[id] = struct{}{}
		}
		if _, ok := allowed[p.PermissionId]; !ok {
			gins.StatusFailed(ctx, http.StatusForbidden, nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
