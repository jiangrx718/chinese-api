package picture_book_item

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
)

func (s *Service) ItemDelete(ctx context.Context, id int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	if _, err := q.SPictureBookItem.Where(q.SPictureBookItem.Id.Eq(id)).Delete(); err != nil {
		logObj.Errorw("SPictureBookItem Delete error", "id", id, "error", err)
		return result, err
	}

	result.SetMessage("操作成功")
	return result, nil
}
