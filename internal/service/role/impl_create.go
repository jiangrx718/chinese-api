package role

import (
	"context"
	"crm/gopkg/log"
	"crm/gopkg/utils"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

type RespRoleCreateInfo struct {
	RoleId   string `json:"role_id"`
	RoleName string `json:"role_name"`
	Status   string `json:"status"`
}

func (s *Service) RoleCreate(ctx context.Context, roleName, status string, permission []string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	roleId := utils.GenUUID()
	crmRole := model.CRMRole{
		RoleId:   roleId,
		RoleName: roleName,
		Status:   status,
	}
	if createErr := g.CRMRole.Create(&crmRole); createErr != nil {
		logObj.Errorw("CRMRole Create crmRole error", "crmRole", crmRole, "error", createErr)
		return result, createErr
	}

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

	result.Data = RespRoleCreateInfo{
		RoleId:   roleId,
		RoleName: roleName,
		Status:   status,
	}
	result.SetMessage("操作成功")
	return result, nil
}
