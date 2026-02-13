package picture_book

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
)

type RespBookStatusInfo struct {
	BookId string `json:"book_id"`
	Status string `json:"status"`
}

func (s *Service) BookStatus(ctx context.Context, bookId, status string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	info, err := q.SPictureBook.Where(q.SPictureBook.BookId.Eq(bookId)).Update(q.SPictureBook.Status, status)
	if err != nil {
		logObj.Errorw("BookStatus Update error", "bookId", bookId, "error", err)
		return result, err
	}
	if info.RowsAffected == 0 {
		return result, fmt.Errorf("record not found")
	}

	result.Data = RespBookStatusInfo{
		BookId: bookId,
		Status: status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
