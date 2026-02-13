package category

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

type RespCategoryStatusInfo struct {
	CategoryId string `json:"category_id"`
	Status     string `json:"status"`
}

func (s *Service) CategoryStatus(ctx context.Context, categoryId, status string) (common.ServiceResult, error) {
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
		return result, fmt.Errorf("categoryEntity not found")
	}

	categoryEntity.Status = status

	if _, err = g.CRMCategory.Where(
		g.CRMCategory.CategoryId.Eq(categoryId),
	).Updates(&categoryEntity); err != nil {
		logObj.Errorw("CategoryStatus CRMCategory error", "CRMCategory", categoryEntity, "error", err)
		return result, err
	}

	result.Data = RespCategoryStatusInfo{
		CategoryId: categoryId,
		Status:     status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
