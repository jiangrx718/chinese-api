package service

import (
	"crm/internal/common"

	"golang.org/x/net/context"
)

type CategoryIFace interface {
	CategoryCreate(ctx context.Context, parentId, categoryName, categoryImage, status string, categoryType, position int) (common.ServiceResult, error)
	CategoryUpdate(ctx context.Context, categoryId, parentId, categoryName, categoryImage, status string, categoryType, position int) (common.ServiceResult, error)
	CategoryList(ctx context.Context, offset, limit int64, categoryType int) (common.ServiceResult, error)
	CategoryStatus(ctx context.Context, categoryId, status string) (common.ServiceResult, error)
	CategoryDelete(ctx context.Context, categoryId string) (common.ServiceResult, error)
}
