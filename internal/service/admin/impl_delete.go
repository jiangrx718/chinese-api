package admin

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"

	"gorm.io/gen"
)

func (s *Service) AdminDelete(ctx context.Context, adminId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	where := []gen.Condition{
		g.CRMAdmin.AdminId.Eq(adminId),
	}
	adminEntity, err := g.CRMAdmin.Where(where...).Take()
	if err != nil {
		return result, err
	}
	if adminEntity == nil {
		return result, fmt.Errorf("admin not found")
	}

	// 判断是否为初始数据
	if adminEntity.IsInit == model.IsInitOn {
		return result, fmt.Errorf("初始数据禁止删除")
	}

	if _, err := g.CRMAdmin.Where(where...).Unscoped().Delete(); err != nil {
		logObj.Errorf("CRMAdmin Delete admin Delete has error(%v)", err)
		return result, err
	}

	result.Data = map[string]string{
		"admin_id": adminId,
	}

	result.SetMessage("操作成功")

	return result, nil
}
