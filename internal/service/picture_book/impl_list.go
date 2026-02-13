package picture_book

import (
	"context"
	"crm/gopkg/gorms"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"

	"gorm.io/gen"
)

type RespBookService struct {
	BookId     string `json:"book_id"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	CategoryId string `json:"category_id"`
	Status     string `json:"status"`
	Position   int    `json:"position"`
	Type       int    `json:"type"`
	CreatedAt  string `json:"created_at"`
}

func (s *Service) BookList(ctx context.Context, offset, limit int64, title string, cType int, categoryId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	q := g.Use(gorms.GetClient("account"))
	spb := q.SPictureBook
	query := spb.Debug()
	var conditions []gen.Condition
	if title != "" {
		conditions = append(conditions, spb.Title.Like("%"+title+"%"))
	}
	if cType != 0 {
		conditions = append(conditions, spb.Type.Eq(cType))
	}
	if categoryId != "" {
		conditions = append(conditions, spb.CategoryId.Eq(categoryId))
	}

	list, count, err := query.Where(conditions...).Order(spb.Position.Desc(), spb.Id.Desc()).FindByPage(int(offset), int(limit))
	if err != nil {
		logObj.Errorw("BookList Find error", "error", err)
		return result, err
	}

	var respList []RespBookService
	for _, v := range list {
		respList = append(respList, RespBookService{
			BookId:     v.BookId,
			Title:      v.Title,
			Icon:       v.Icon,
			CategoryId: v.CategoryId,
			Status:     v.Status,
			Position:   v.Position,
			Type:       v.Type,
			CreatedAt:  v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	result.Data = map[string]any{"list": respList, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}
