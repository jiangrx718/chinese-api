package admin

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespAdminService struct {
	AdminId   string `json:"admin_id"`
	UserName  string `json:"user_name"`
	UserPhone string `json:"user_phone"`
	Status    string `json:"status"`
	RoleId    string `json:"role_id"`
	IsInit    string `json:"is_init"`
	CreatedAt string `json:"created_at"`
}

func (s *Service) AdminList(ctx context.Context, offset, limit int64, status, userPhone string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	adminDataList, count, err := ScanByPage(userPhone, status, offset, limit)
	if err != nil {
		logObj.Errorw("AdminList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	var listAdmin []RespAdminService
	for idx, _ := range adminDataList {
		listAdmin = append(listAdmin, RespAdminService{
			AdminId:   adminDataList[idx].AdminId,
			UserName:  adminDataList[idx].UserName,
			UserPhone: adminDataList[idx].UserPhone,
			RoleId:    adminDataList[idx].RoleId,
			Status:    adminDataList[idx].Status,
			IsInit:    adminDataList[idx].IsInit,
			CreatedAt: adminDataList[idx].CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	if len(listAdmin) == 0 {
		listAdmin = []RespAdminService{}
	}
	result.Data = map[string]any{"list": listAdmin, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(phone, status string, offset, limit int64) ([]*model.CRMAdmin, int64, error) {
	var (
		crmAdmin = g.CRMAdmin
		response = make([]*model.CRMAdmin, 0)
	)

	q := crmAdmin.Debug()
	where := []gen.Condition{}

	// 手机号
	if phone != "" {
		where = append(where, crmAdmin.UserPhone.Eq(phone))
	}
	// 启用状态
	if status != "" {
		where = append(where, crmAdmin.Status.Eq(status))
	}

	count, err := q.Where(where...).Order(crmAdmin.Id.Asc()).ScanByPage(&response, int(offset), int(limit))
	return response, count, err
}
