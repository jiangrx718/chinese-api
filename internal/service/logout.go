package service

import (
	"context"
	"crm/internal/common"
)

type LogoutIFace interface {
	Logout(ctx context.Context, userId string) (common.ServiceResult, error)
}
