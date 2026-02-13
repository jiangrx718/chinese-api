package picture_book_item

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"

	"gorm.io/gen"
)

type RespItemService struct {
	Id        int    `json:"id"`
	BookId    string `json:"book_id"`
	Title     string `json:"title"`
	Pic       string `json:"pic"`
	BPic      string `json:"b_pic"`
	Audio     string `json:"audio"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	Position  int    `json:"position"`
	CreatedAt string `json:"created_at"`
}

func (s *Service) ItemList(ctx context.Context, offset, limit int64, bookId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	spbi := q.SPictureBookItem
	query := spbi.Debug()
	var conditions []gen.Condition
	if bookId != "" {
		conditions = append(conditions, spbi.BookId.Eq(bookId))
	}

	list, count, err := query.Where(conditions...).Order(spbi.Position.Desc(), spbi.Id.Desc()).FindByPage(int(offset), int(limit))
	if err != nil {
		logObj.Errorw("ItemList Find error", "error", err)
		return result, err
	}

	var respList []RespItemService
	for _, v := range list {
		respList = append(respList, RespItemService{
			Id:        v.Id,
			BookId:    v.BookId,
			Title:     v.Title,
			Pic:       v.Pic,
			BPic:      v.BPic,
			Audio:     v.Audio,
			Content:   v.Content,
			Status:    v.Status,
			Position:  v.Position,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	result.Data = map[string]any{"list": respList, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}
