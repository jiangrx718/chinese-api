package service

import (
	"crm/internal/common"

	"golang.org/x/net/context"
)

type RoleIFace interface {
	RoleCreate(ctx context.Context, roleName, status string, permission []string) (common.ServiceResult, error)
	RoleUpdate(ctx context.Context, roleId, roleName, status string, permission []string) (common.ServiceResult, error)
	RoleList(ctx context.Context, offset, limit int64) (common.ServiceResult, error)
	RoleDelete(ctx context.Context, roleId string) (common.ServiceResult, error)
}
