package service

import (
	"context"
	"crm/internal/common"
)

type PictureIFace interface {
	PictureList(ctx context.Context, categoryId string, offset, limit int64) (common.ServiceResult, error)
}
