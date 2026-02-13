package service

import (
	"context"
	"crm/internal/common"
)

type BookIFace interface {
	BookList(ctx context.Context, sType int) (common.ServiceResult, error)
}
