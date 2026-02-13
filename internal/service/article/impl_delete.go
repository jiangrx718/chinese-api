package article

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

func (s *Service) ArticleDelete(ctx context.Context, articleId string) (common.ServiceResult, error) {
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

	if _, err := g.CRMArticle.Where(where...).Unscoped().Delete(); err != nil {
		logObj.Errorf("CRMArticle Delete article Delete has error(%v)", err)
		return result, err
	}

	contentWhere := []gen.Condition{
		g.CRMArticleContent.ArticleId.Eq(articleId),
	}
	_, _ = g.CRMArticleContent.Where(contentWhere...).Unscoped().Delete()

	result.Data = map[string]string{
		"article_id": articleId,
	}

	result.SetMessage("操作成功")
	return result, nil
}
