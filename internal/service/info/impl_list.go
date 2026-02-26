package info

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespInfoService struct {
	Position int    `json:"position"`
	Pic      string `json:"pic"`
	Mp3      string `json:"mp3"`
}

func (s *Service) InfoList(ctx context.Context, bookId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	infoDataList, count, err := ScanByPage(bookId)
	if err != nil {
		logObj.Errorw("BookList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	var listInfo []RespInfoService
	for idx, _ := range infoDataList {
		listInfo = append(listInfo, RespInfoService{
			Pic:      infoDataList[idx].Pic,
			Position: infoDataList[idx].Position,
			Mp3:      infoDataList[idx].Audio,
		})
	}

	if len(listInfo) == 0 {
		listInfo = []RespInfoService{}
	}
	result.Data = map[string]any{"list": listInfo, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(bookId string) ([]*model.SPictureBookItem, int64, error) {
	var (
		sPictureBookItem = g.SPictureBookItem
		response         = make([]*model.SPictureBookItem, 0)
	)

	q := sPictureBookItem.Debug()
	where := []gen.Condition{}

	where = append(where, sPictureBookItem.Status.Eq("on"))
	where = append(where, sPictureBookItem.BookId.Eq(bookId))

	count, err := q.Where(where...).Order(sPictureBookItem.Position.Asc()).ScanByPage(&response, 0, -1)
	return response, count, err
}
