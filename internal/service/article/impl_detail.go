package article

import (
	"context"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

type RespArticleDetail struct {
	Id             int    `json:"id"`
	ArticleId      string `json:"article_id"`
	CategoryId     string `json:"category_id"`
	CategoryName   string `json:"category_name"`
	ArticleName    string `json:"article_name"`
	ArticleImage   string `json:"article_image"`
	Status         string `json:"status"`
	Position       int    `json:"position"`
	ArticleContent string `json:"article_content"`
	CreatedAt      string `json:"created_at"`
}

func (s *Service) ArticleDetail(ctx context.Context, articleId string) (common.ServiceResult, error) {
	var (
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

	// 更新文章内容表
	articleContentWhere := []gen.Condition{
		g.CRMArticleContent.ArticleId.Eq(articleId),
	}
	articleContentEntity, err := g.CRMArticleContent.Where(articleContentWhere...).Take()
	if err != nil {
		return result, err
	}
	if articleContentEntity == nil {
		return result, fmt.Errorf("article content not found")
	}
	
	result.Data = RespArticleDetail{
		Id:             articleEntity.Id,
		ArticleId:      articleEntity.ArticleId,
		CategoryId:     articleEntity.CategoryId,
		ArticleName:    articleEntity.ArticleName,
		ArticleImage:   articleEntity.ArticleImage,
		Status:         articleEntity.Status,
		Position:       articleEntity.Position,
		ArticleContent: articleContentEntity.ArticleContent,
		CreatedAt:      articleEntity.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	result.SetMessage("操作成功")
	return result, nil
}
