package category

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"

	"gorm.io/gen"
)

func (s *Service) CategoryDelete(ctx context.Context, categoryId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	where := []gen.Condition{
		g.CRMCategory.CategoryId.Eq(categoryId),
	}
	categoryEntity, err := g.CRMCategory.Where(where...).Take()
	if err != nil {
		return result, err
	}
	if categoryEntity == nil {
		return result, fmt.Errorf("category not found")
	}

	// 如果是文章的类型
	if categoryEntity.CategoryType == model.CategoryTypeArticle {
		// 判断是否文章是否使用
		articleWhere := []gen.Condition{
			g.CRMArticle.CategoryId.Eq(categoryId),
		}
		articleEntity, _ := g.CRMArticle.Where(articleWhere...).Take()
		if articleEntity != nil && articleEntity.Id > 0 {
			return result, fmt.Errorf("当前栏目下存在文章数据")
		}
	}

	if _, err := g.CRMCategory.Where(where...).Unscoped().Delete(); err != nil {
		logObj.Errorf("Category Delete Category Delete has error(%v)", err)
		return result, err
	}

	result.Data = map[string]string{
		"category_id": categoryId,
	}

	result.SetMessage("操作成功")

	return result, nil
}
