package picture_book_category

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"

	"gorm.io/gen"
)

type RespCategoryService struct {
	Id           int    `json:"id"`
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Status       string `json:"status"`
	Position     int    `json:"position"`
	Type         int    `json:"type"`
	CreatedAt    string `json:"created_at"`
}

func (s *Service) CategoryList(ctx context.Context, offset, limit int64, categoryName string, cType int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	spbc := q.SPictureBookCategory
	query := spbc.Debug()
	var conditions []gen.Condition
	if categoryName != "" {
		conditions = append(conditions, spbc.CategoryName.Like("%"+categoryName+"%"))
	}
	if cType != 0 {
		conditions = append(conditions, spbc.Type.Eq(cType))
	}

	list, count, err := query.Where(conditions...).Order(spbc.Position.Desc(), spbc.Id.Desc()).FindByPage(int(offset), int(limit))
	if err != nil {
		logObj.Errorw("CategoryList Find error", "error", err)
		return result, err
	}

	var respList []RespCategoryService
	for _, v := range list {
		respList = append(respList, RespCategoryService{
			Id:           v.Id,
			CategoryId:   v.CategoryId,
			CategoryName: v.CategoryName,
			Status:       v.Status,
			Position:     v.Position,
			Type:         v.Type,
			CreatedAt:    v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	result.Data = map[string]any{"list": respList, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}
