package picture_book

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
)

type RespBookDetail struct {
	BookId     string `json:"book_id"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	CategoryId string `json:"category_id"`
	Status     string `json:"status"`
	Position   int    `json:"position"`
	Type       int    `json:"type"`
	CreatedAt  string `json:"created_at"`
}

func (s *Service) BookDetail(ctx context.Context, bookId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	entity, err := q.SPictureBook.Where(q.SPictureBook.BookId.Eq(bookId)).Take()
	if err != nil {
		logObj.Errorw("BookDetail Find error", "bookId", bookId, "error", err)
		return result, err
	}
	if entity == nil {
		return result, fmt.Errorf("record not found")
	}

	result.Data = RespBookDetail{
		BookId:     entity.BookId,
		Title:      entity.Title,
		Icon:       entity.Icon,
		CategoryId: entity.CategoryId,
		Status:     entity.Status,
		Position:   entity.Position,
		Type:       entity.Type,
		CreatedAt:  entity.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	result.SetMessage("操作成功")
	return result, nil
}
