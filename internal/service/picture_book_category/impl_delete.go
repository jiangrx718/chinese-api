package picture_book_category

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
)

func (s *Service) CategoryDelete(ctx context.Context, categoryId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	if _, err := q.SPictureBookCategory.Where(q.SPictureBookCategory.CategoryId.Eq(categoryId)).Delete(); err != nil {
		logObj.Errorw("SPictureBookCategory Delete error", "categoryId", categoryId, "error", err)
		return result, err
	}

	result.SetMessage("操作成功")
	return result, nil
}
