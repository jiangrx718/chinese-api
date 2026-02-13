package service

import (
	"crm/internal/common"

	"golang.org/x/net/context"
)

type PermissionIFace interface {
	PermissionCreate(ctx context.Context, permissionName, permissionUrl, parentId, status string, position, permissionType int) (common.ServiceResult, error)
	PermissionUpdate(ctx context.Context, permissionId, permissionName, permissionUrl, parentId, status string, position int) (common.ServiceResult, error)
	PermissionStatus(ctx context.Context, permissionId, status string) (common.ServiceResult, error)
	PermissionList(ctx context.Context, status string) (common.ServiceResult, error)
	PermissionDelete(ctx context.Context, permissionId string) (common.ServiceResult, error)
	PermissionFind(ctx context.Context, adminId string) (common.ServiceResult, error)
	PermissionMenu(ctx context.Context, adminId string) (common.ServiceResult, error)
}
