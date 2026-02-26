package service

import (
	"context"
	"crm/internal/common"
)

type CategoryIFace interface {
	CategoryList(ctx context.Context, sType string) (common.ServiceResult, error)
}
