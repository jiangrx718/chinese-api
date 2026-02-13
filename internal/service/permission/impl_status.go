package permission

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"fmt"

	"gorm.io/gen"
)

type RespPermissionStatusInfo struct {
	PermissionId string `json:"permission_id"`
	Status       string `json:"status"`
}

func (s *Service) PermissionStatus(ctx context.Context, permissionId, status string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	where := []gen.Condition{
		g.CRMPermission.PermissionId.Eq(permissionId),
	}
	permissionEntity, err := g.CRMPermission.Where(where...).Take()
	if err != nil {
		return result, err
	}
	if permissionEntity == nil {
		return result, fmt.Errorf("permission not found")
	}

	permissionEntity.Status = status

	if _, err = g.CRMPermission.Where(
		g.CRMPermission.PermissionId.Eq(permissionId),
	).Updates(&permissionEntity); err != nil {
		logObj.Errorw("PermissionStatus permission error", "permission", permissionEntity, "error", err)
		return result, err
	}

	result.Data = RespPermissionStatusInfo{
		PermissionId: permissionId,
		Status:       status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
