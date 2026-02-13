package service

import (
	"context"
	"crm/internal/common"
)

type ArticleIFace interface {
	ArticleCreate(ctx context.Context, categoryId, articleName, articleImage, status, articleContent string, position int) (common.ServiceResult, error)
	ArticleUpdate(ctx context.Context, articleId, categoryId, articleName, articleImage, status, articleContent string, position int) (common.ServiceResult, error)
	ArticleList(ctx context.Context, offset, limit int64, categoryId, articleName string) (common.ServiceResult, error)
	ArticleDelete(ctx context.Context, articleId string) (common.ServiceResult, error)
	ArticleStatus(ctx context.Context, articleId, status string) (common.ServiceResult, error)
	ArticleDetail(ctx context.Context, articleId string) (common.ServiceResult, error)
}
