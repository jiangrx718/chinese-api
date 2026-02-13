package article

import (
	"context"
	"crm/gopkg/log"
	"crm/gopkg/utils"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

type RespArticleInfo struct {
	ArticleId    string `json:"article_id"`
	CategoryId   string `json:"category_id"`
	ArticleName  string `json:"article_name"`
	ArticleImage string `json:"article_image"`
	Status       string `json:"status"`
	Position     int    `json:"position"`
}

func (s *Service) ArticleCreate(ctx context.Context, categoryId, articleName, articleImage, status, articleContent string, position int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	articleId := utils.GenUUID()
	crmArticle := model.CRMArticle{
		ArticleId:    articleId,
		CategoryId:   categoryId,
		ArticleName:  articleName,
		ArticleImage: articleImage,
		Status:       status,
		Position:     position,
	}

	// 添加到文章表
	if createErr := g.CRMArticle.Create(&crmArticle); createErr != nil {
		logObj.Errorw("CRMArticle Create crmArticle error", "crmArticle", crmArticle, "error", createErr)
		return result, createErr
	}

	// 添加到文章内容表
	crmArticleContent := model.CRMArticleContent{
		ContentId:      utils.GenUUID(),
		ArticleId:      articleId,
		ArticleContent: articleContent,
	}
	if createErr := g.CRMArticleContent.Create(&crmArticleContent); createErr != nil {
		logObj.Errorw("CRMArticleContent Create crmArticleContent error", "crmArticleContent", crmArticleContent, "error", createErr)
		return result, createErr
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
