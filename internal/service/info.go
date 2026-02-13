package service

import (
	"context"
	"crm/internal/common"
)

type InfoIFace interface {
	InfoList(ctx context.Context, bookId string) (common.ServiceResult, error)
}
