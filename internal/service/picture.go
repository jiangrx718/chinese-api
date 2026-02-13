package service

import (
	"context"
	"crm/internal/common"
)

type PictureIFace interface {
	PictureList(ctx context.Context, sType int, offset, limit int64) (common.ServiceResult, error)
}
