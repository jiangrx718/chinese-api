package article

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"

	"gorm.io/gen"
)

func (s *Service) ArticleUpdate(ctx context.Context, articleId, categoryId, articleName, articleImage, status, articleContent string, position int) (common.ServiceResult, error) {
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

	articleEntity.CategoryId = categoryId
	articleEntity.ArticleName = articleName
	articleEntity.ArticleImage = articleImage
	articleEntity.Status = status
	articleEntity.Position = position

	// 更新文章表
	if _, err = g.CRMArticle.Where(
		g.CRMArticle.ArticleId.Eq(articleId),
	).Updates(&articleEntity); err != nil {
		logObj.Errorw("ArticleUpdate article error", "articleEntity", articleEntity, "error", err)
		return result, err
	}

	// 更新文章内容表
	var articleContentEntity model.CRMArticleContent
	articleContentEntity.ArticleContent = articleContent
	if _, err = g.CRMArticleContent.Where(
		g.CRMArticleContent.ArticleId.Eq(articleId),
	).Updates(&articleContentEntity); err != nil {
		logObj.Errorw("ArticleUpdate article error", "articleContentEntity", articleContentEntity, "error", err)
		return result, err
	}

	result.Data = RespArticleInfo{
		ArticleId:    articleId,
		CategoryId:   categoryId,
		ArticleName:  articleName,
		ArticleImage: articleImage,
		Status:       status,
		Position:     position,
	}
	result.SetMessage("操作成功")
	return result, nil
}
