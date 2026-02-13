package picture_book

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
)

type RespBookUpdateInfo struct {
	BookId     string `json:"book_id"`
	CategoryId string `json:"category_id"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	Status     string `json:"status"`
	Position   int    `json:"position"`
	Type       int    `json:"type"`
}

func (s *Service) BookUpdate(ctx context.Context, bookId, categoryId, title, icon, status string, position, cType int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	entity, err := q.SPictureBook.Where(q.SPictureBook.BookId.Eq(bookId)).Take()
	if err != nil {
		return result, err
	}
	if entity == nil {
		return result, fmt.Errorf("record not found")
	}

	entity.CategoryId = categoryId
	entity.Title = title
	entity.Icon = icon
	entity.Status = status
	entity.Position = position
	entity.Type = cType

	if _, err := q.SPictureBook.Where(q.SPictureBook.BookId.Eq(bookId)).Updates(entity); err != nil {
		logObj.Errorw("SPictureBook Update error", "entity", entity, "error", err)
		return result, err
	}
	result.Data = RespBookUpdateInfo{
		BookId:     entity.BookId,
		CategoryId: entity.CategoryId,
		Title:      entity.Title,
		Icon:       entity.Icon,
		Status:     entity.Status,
		Position:   entity.Position,
		Type:       entity.Type,
	}
	result.SetMessage("操作成功")
	return result, nil
}
