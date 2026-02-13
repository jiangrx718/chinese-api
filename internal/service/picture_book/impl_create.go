package picture_book

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/gopkg/utils"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

type RespBookCreateInfo struct {
	BookId     string `json:"book_id"`
	CategoryId string `json:"category_id"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	Status     string `json:"status"`
	Position   int    `json:"position"`
	Type       int    `json:"type"`
}

func (s *Service) BookCreate(ctx context.Context, categoryId, title, icon, status string, position, cType int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	bookId := utils.GenUUID()
	entity := model.SPictureBook{
		BookId:     bookId,
		CategoryId: categoryId,
		Title:      title,
		Icon:       icon,
		Status:     status,
		Position:   position,
		Type:       cType,
	}

	q := g.Use(gorms.GetClient("account"))
	if err := q.SPictureBook.Create(&entity); err != nil {
		logObj.Errorw("SPictureBook Create error", "entity", entity, "error", err)
		return result, err
	}

	result.Data = RespBookCreateInfo{
		BookId:     bookId,
		CategoryId: categoryId,
		Title:      title,
		Icon:       icon,
		Status:     status,
		Position:   position,
		Type:       cType,
	}
	result.SetMessage("操作成功")
	return result, nil
}
