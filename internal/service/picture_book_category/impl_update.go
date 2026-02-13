package picture_book_category

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
)

type RespCategoryUpdateInfo struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Status       string `json:"status"`
	Position     int    `json:"position"`
	CategoryType int    `json:"category_type"`
}

func (s *Service) CategoryUpdate(ctx context.Context, categoryId, categoryName, status string, position, cType int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	entity, err := q.SPictureBookCategory.Where(q.SPictureBookCategory.CategoryId.Eq(categoryId)).Take()
	if err != nil {
		return result, err
	}
	if entity == nil {
		return result, fmt.Errorf("record not found")
	}

	entity.CategoryName = categoryName
	entity.Status = status
	entity.Position = position
	entity.Type = cType

	if _, err := q.SPictureBookCategory.Where(q.SPictureBookCategory.CategoryId.Eq(categoryId)).Updates(entity); err != nil {
		logObj.Errorw("SPictureBookCategory Update error", "entity", entity, "error", err)
		return result, err
	}

	result.Data = RespCategoryUpdateInfo{
		CategoryId:   categoryId,
		CategoryName: categoryName,
		Status:       status,
		Position:     position,
		CategoryType: cType,
	}
	result.SetMessage("操作成功")
	return result, nil
}
