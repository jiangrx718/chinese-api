package category

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespCategoryService struct {
	CategoryId string `json:"category_id"`
	Name       string `json:"name"`
}

func (s *Service) CategoryList(ctx context.Context, sType string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	categoryDataList, count, err := ScanByPage(sType)
	if err != nil {
		logObj.Errorw("BookList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	var listCategory []RespCategoryService
	for idx := range categoryDataList {
		listCategory = append(listCategory, RespCategoryService{
			CategoryId: categoryDataList[idx].CategoryId,
			Name:       categoryDataList[idx].CategoryName,
		})
	}

	if len(listCategory) == 0 {
		listCategory = []RespCategoryService{}
	}
	result.Data = map[string]any{"list": listCategory, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(sType string) ([]*model.SPictureCategory, int64, error) {
	var (
		sPictureCategory = g.SPictureCategory
		response         = make([]*model.SPictureCategory, 0)
	)

	q := sPictureCategory.Debug()
	where := []gen.Condition{}

	where = append(where, sPictureCategory.Status.Eq("on"))
	where = append(where, sPictureCategory.Type.Eq(sType))

	count, err := q.Where(where...).Order(sPictureCategory.Position.Desc()).ScanByPage(&response, 0, -1)
	return response, count, err
}
