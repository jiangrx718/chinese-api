package service

import (
	"context"
	"crm/internal/common"
)

type PictureBookCategoryIFace interface {
	CategoryCreate(ctx context.Context, categoryName, status string, position, categoryType int) (common.ServiceResult, error)
	CategoryUpdate(ctx context.Context, categoryId, categoryName, status string, position, categoryType int) (common.ServiceResult, error)
	CategoryList(ctx context.Context, offset, limit int64, categoryName string, categoryType int) (common.ServiceResult, error)
	CategoryDelete(ctx context.Context, categoryId string) (common.ServiceResult, error)
	CategoryStatus(ctx context.Context, categoryId, status string) (common.ServiceResult, error)
}
