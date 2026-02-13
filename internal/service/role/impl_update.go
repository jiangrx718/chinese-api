package role

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"

	"gorm.io/gen"
)

type RespRoleUpdateInfo struct {
	RoleId   string `json:"role_id"`
	RoleName string `json:"role_name"`
	Status   string `json:"status"`
}

func (s *Service) RoleUpdate(ctx context.Context, roleId, roleName, status string, permission []string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	where := []gen.Condition{
		g.CRMRole.RoleId.Eq(roleId),
	}
	roleEntity, err := g.CRMRole.Where(where...).Take()
	if err != nil {
		return result, err
	}
	if roleEntity == nil {
		return result, fmt.Errorf("role not found")
	}

	// 判断是否为初始数据
	if roleEntity.IsInit == model.IsInitOn {
		return result, fmt.Errorf("初始数据禁止修改")
	}

	roleEntity.RoleName = roleName
	roleEntity.Status = status

	if _, err = g.CRMRole.Where(
		g.CRMRole.RoleId.Eq(roleId),
	).Updates(&roleEntity); err != nil {
		logObj.Errorw("RoleUpdate role error", "role", roleEntity, "error", err)
		return result, err
	}

	// 先把原有的权限删除
	if _, err := g.CRMRolePermission.Where([]gen.Condition{
		g.CRMRolePermission.RoleId.Eq(roleId),
	}...).Unscoped().Delete(); err != nil {
		logObj.Errorf("CRMRolePermission Delete role Delete has error(%v)", err)
	}

	// 重新设置对应额权限
	var rolePermission []model.CRMRolePermission
	for _, permissionId := range permission {
		rolePermission = append(rolePermission, model.CRMRolePermission{
			RoleId:       roleId,
			PermissionId: permissionId,
		})
	}
	for _, vItem := range rolePermission {
		_ = g.CRMRolePermission.Create(&vItem)
	}

	result.Data = RespRoleUpdateInfo{
		RoleId:   roleId,
		RoleName: roleName,
		Status:   status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
