package book

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespBookService struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
}

func (s *Service) BookList(ctx context.Context, sType int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	bookDataList, count, err := ScanByPage(sType)
	if err != nil {
		logObj.Errorw("BookList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	var listBook []RespBookService
	for idx, _ := range bookDataList {
		listBook = append(listBook, RespBookService{
			Id:         bookDataList[idx].Id,
			CategoryId: bookDataList[idx].CategoryId,
			Name:       bookDataList[idx].Name,
		})
	}

	if len(listBook) == 0 {
		listBook = []RespBookService{}
	}
	result.Data = map[string]any{"list": listBook, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(sType int) ([]*model.SBookName, int64, error) {
	var (
		sBookName = g.SBookName
		response  = make([]*model.SBookName, 0)
	)

	q := sBookName.Debug()
	where := []gen.Condition{}

	where = append(where, sBookName.Status.Eq(1))
	where = append(where, sBookName.SType.Eq(sType))

	count, err := q.Where(where...).Order(sBookName.Id.Asc()).ScanByPage(&response, 0, -1)
	return response, count, err
}
