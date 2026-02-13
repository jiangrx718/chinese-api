package service

import (
	"context"
	"crm/internal/common"
)

type AdminIFace interface {
	AdminCreate(ctx context.Context, userName, userPhone, password, status, roleId string) (common.ServiceResult, error)
	AdminUpdate(ctx context.Context, adminId, password, status, roleId string) (common.ServiceResult, error)
	AdminList(ctx context.Context, offset, limit int64, status, userPhone string) (common.ServiceResult, error)
	AdminDelete(ctx context.Context, adminId string) (common.ServiceResult, error)
}
