package service

import (
	"context"
	"crm/internal/common"
)

type LoginIFace interface {
	Login(ctx context.Context, userName, password string) (common.ServiceResult, error)
}
