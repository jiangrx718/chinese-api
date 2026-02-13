package picture_book_item

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"
)

type RespItemUpdateInfo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Pic      string `json:"pic"`
	BPic     string `json:"b_pic"`
	Audio    string `json:"audio"`
	Content  string `json:"content"`
	Status   string `json:"status"`
	Position int    `json:"position"`
}

func (s *Service) ItemUpdate(ctx context.Context, id int, title, pic, bPic, audio, content, status string, position int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	entity, err := q.SPictureBookItem.Where(q.SPictureBookItem.Id.Eq(id)).Take()
	if err != nil {
		return result, err
	}
	if entity == nil {
		return result, fmt.Errorf("record not found")
	}

	entity.Title = title
	entity.Pic = pic
	entity.BPic = bPic
	entity.Audio = audio
	entity.Content = content
	entity.Status = status
	entity.Position = position

	if _, err := q.SPictureBookItem.Where(q.SPictureBookItem.Id.Eq(id)).Updates(entity); err != nil {
		logObj.Errorw("SPictureBookItem Update error", "entity", entity, "error", err)
		return result, err
	}

	result.Data = RespItemUpdateInfo{
		Id:       entity.Id,
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
