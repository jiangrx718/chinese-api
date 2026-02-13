package picture_book_item

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

type RespItemCreateInfo struct {
	BookId   string `json:"book_id"`
	Title    string `json:"title"`
	Pic      string `json:"pic"`
	BPic     string `json:"b_pic"`
	Audio    string `json:"audio"`
	Content  string `json:"content"`
	Status   string `json:"status"`
	Position int    `json:"position"`
}

func (s *Service) ItemCreate(ctx context.Context, bookId, title, pic, bPic, audio, content, status string, position int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	entity := model.SPictureBookItem{
		BookId:   bookId,
		Title:    title,
		Pic:      pic,
		BPic:     bPic,
		Audio:    audio,
		Content:  content,
		Status:   status,
		Position: position,
	}

	q := g.Use(gorms.GetClient("account"))
	if err := q.SPictureBookItem.Create(&entity); err != nil {
		logObj.Errorw("SPictureBookItem Create error", "entity", entity, "error", err)
		return result, err
	}

	result.Data = RespItemCreateInfo{
		BookId:   bookId,
		Title:    title,
		Pic:      pic,
		BPic:     bPic,
		Audio:    audio,
		Content:  content,
		Status:   status,
		Position: position,
	}
	result.SetMessage("操作成功")
	return result, nil
}
