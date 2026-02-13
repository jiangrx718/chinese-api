package article

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

func (s *Service) ArticleStatus(ctx context.Context, articleId, status string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	where := []gen.Condition{
		g.CRMArticle.ArticleId.Eq(articleId),
	}
	articleEntity, err := g.CRMArticle.Where(where...).Take()
	if err != nil {
		return result, err
	}
	if articleEntity == nil {
		return result, fmt.Errorf("article not found")
	}

	articleEntity.Status = status

	// 更新文章表
	if _, err = g.CRMArticle.Where(
		g.CRMArticle.ArticleId.Eq(articleId),
	).Updates(&articleEntity); err != nil {
		logObj.Errorw("ArticleUpdate article error", "articleEntity", articleEntity, "error", err)
		return result, err
	}

	result.Data = RespArticleInfo{
		ArticleId: articleId,
		Status:    status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
