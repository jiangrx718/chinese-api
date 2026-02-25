package picture

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespPictureService struct {
	BookId string `json:"book_id"`
	Title  string `json:"title"`
	Icon   string `json:"icon"`
}

func (s *Service) PictureList(ctx context.Context, sType int, offset, limit int64) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	pictureDataList, count, err := ScanByPage(sType, offset, limit)
	if err != nil {
		logObj.Errorw("AdminList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	var listPicture []RespPictureService
	for idx := range pictureDataList {
		listPicture = append(listPicture, RespPictureService{
			BookId: pictureDataList[idx].BookId,
			Title:  pictureDataList[idx].Title,
			Icon:   pictureDataList[idx].Icon,
		})
	}

	if len(listPicture) == 0 {
		listPicture = []RespPictureService{}
	}
	result.Data = map[string]any{"list": listPicture, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(sType int, offset, limit int64) ([]*model.SChinesePicture, int64, error) {
	var (
		sChinesePicture = g.SChinesePicture
		response        = make([]*model.SChinesePicture, 0)
	)

	q := sChinesePicture.Debug()
	where := []gen.Condition{}

	where = append(where, sChinesePicture.Status.Eq(1))
	if sType > 0 {
		where = append(where, sChinesePicture.Type.Eq(sType))
	}

	// 启用状态
	count, err := q.Where(where...).Order(sChinesePicture.Position.Asc()).ScanByPage(&response, int(offset), int(limit))
	return response, count, err
}
