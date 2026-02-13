package picture_book_item

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
)

type RespItemStatusInfo struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func (s *Service) ItemStatus(ctx context.Context, id int, status string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	info, err := q.SPictureBookItem.Where(q.SPictureBookItem.Id.Eq(id)).Update(q.SPictureBookItem.Status, status)
	if err != nil {
		logObj.Errorw("ItemStatus Update error", "id", id, "error", err)
		return result, err
	}
	if info.RowsAffected == 0 {
		return result, fmt.Errorf("record not found")
	}

	result.Data = RespItemStatusInfo{Id: id, Status: status}
	result.SetMessage("操作成功")
	return result, nil
}
