package service

import (
	"context"
	"crm/internal/common"
)

type PictureBookItemIFace interface {
	ItemCreate(ctx context.Context, bookId, title, pic, bPic, audio, content, status string, position int) (common.ServiceResult, error)
	ItemUpdate(ctx context.Context, id int, title, pic, bPic, audio, content, status string, position int) (common.ServiceResult, error)
	ItemList(ctx context.Context, offset, limit int64, bookId string) (common.ServiceResult, error)
	ItemDelete(ctx context.Context, id int) (common.ServiceResult, error)
	ItemStatus(ctx context.Context, id int, status string) (common.ServiceResult, error)
}
