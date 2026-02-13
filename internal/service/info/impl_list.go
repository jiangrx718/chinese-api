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
			Mp3:      infoDataList[idx].Mp3,
		})
	}

	if len(listInfo) == 0 {
		listInfo = []RespInfoService{}
	}
	result.Data = map[string]any{"list": listInfo, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(bookId string) ([]*model.SChinesePictureInfo, int64, error) {
	var (
		sChinesePictureInfo = g.SChinesePictureInfo
		response            = make([]*model.SChinesePictureInfo, 0)
	)

	q := sChinesePictureInfo.Debug()
	where := []gen.Condition{}

	where = append(where, sChinesePictureInfo.Status.Eq(1))
	where = append(where, sChinesePictureInfo.BookId.Eq(bookId))

	count, err := q.Where(where...).Order(sChinesePictureInfo.Position.Asc()).ScanByPage(&response, 0, -1)
	return response, count, err
}
