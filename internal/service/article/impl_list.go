package article

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespArticleService struct {
	Id           int    `json:"id"`
	ArticleId    string `json:"article_id"`
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	ArticleName  string `json:"article_name"`
	ArticleImage string `json:"article_image"`
	Status       string `json:"status"`
	Position     int    `json:"position"`
	CreatedAt    string `json:"created_at"`
}

func (s *Service) ArticleList(ctx context.Context, offset, limit int64, categoryId, articleName string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	articleDataList, count, err := ScanByPage(articleName, categoryId, offset, limit)
	if err != nil {
		logObj.Errorw("ArticleList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	// 批量查询分类名称
	categoryIds := make([]string, 0)
	for _, v := range articleDataList {
		if v.CategoryId != "" {
			categoryIds = append(categoryIds, v.CategoryId)
		}
	}

	categoryMap := make(map[string]string)
	if len(categoryIds) > 0 {
		crmCategory := g.CRMCategory
		categories, err := crmCategory.Where(crmCategory.CategoryId.In(categoryIds...)).Find()
		if err == nil {
			for _, v := range categories {
				categoryMap[v.CategoryId] = v.CategoryName
			}
		} else {
			logObj.Errorw("ArticleList Category Find", "error", err)
		}
	}

	var listArticle []RespArticleService
	for idx, _ := range articleDataList {
		cName := ""
		if val, ok := categoryMap[articleDataList[idx].CategoryId]; ok {
			cName = val
		}

		listArticle = append(listArticle, RespArticleService{
			Id:           int(articleDataList[idx].Id),
			ArticleId:    articleDataList[idx].ArticleId,
			CategoryId:   articleDataList[idx].CategoryId,
			CategoryName: cName,
			ArticleName:  articleDataList[idx].ArticleName,
			ArticleImage: articleDataList[idx].ArticleImage,
			Status:       articleDataList[idx].Status,
			Position:     articleDataList[idx].Position,
			CreatedAt:    articleDataList[idx].CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	if len(listArticle) == 0 {
		listArticle = []RespArticleService{}
	}
	result.Data = map[string]any{"list": listArticle, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(articleName, categoryId string, offset, limit int64) ([]*model.CRMArticle, int64, error) {
	var (
		crmArticle = g.CRMArticle
		response   = make([]*model.CRMArticle, 0)
	)

	q := crmArticle.Debug()
	where := []gen.Condition{}

	// 手机号
	if articleName != "" {
		pattern := "%" + articleName + "%"
		where = append(where, crmArticle.ArticleName.Like(pattern))
	}
	// 启用状态
	if categoryId != "" {
		where = append(where, crmArticle.CategoryId.Eq(categoryId))
	}

	count, err := q.Where(where...).Order(crmArticle.Id.Desc()).ScanByPage(&response, int(offset), int(limit))
	return response, count, err
}
