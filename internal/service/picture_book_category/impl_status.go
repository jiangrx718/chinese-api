package picture_book_category

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
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

	q := g.Use(gorms.GetClient("account"))
	info, err := q.SPictureBookCategory.Where(q.SPictureBookCategory.CategoryId.Eq(categoryId)).Update(q.SPictureBookCategory.Status, status)
	if err != nil {
		logObj.Errorw("CategoryStatus Update error", "categoryId", categoryId, "error", err)
		return result, err
	}
	if info.RowsAffected == 0 {
		return result, fmt.Errorf("record not found")
	}

	result.Data = RespCategoryStatusInfo{
		CategoryId: categoryId,
		Status:     status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
