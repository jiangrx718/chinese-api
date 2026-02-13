package role

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespRoleService struct {
	RoleId     string   `json:"role_id"`
	RoleName   string   `json:"role_name"`
	Status     string   `json:"status"`
	IsInit     string   `json:"is_init"`
	CreatedAt  string   `json:"created_at"`
	Permission []string `json:"permission"`
}

func (s *Service) RoleList(ctx context.Context, offset, limit int64) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	roleDataList, count, err := ScanByPage(offset, limit)
	if err != nil {
		logObj.Errorw("AdminList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	var listRole []RespRoleService
	if len(roleDataList) == 0 {
		listRole = []RespRoleService{}
		result.Data = map[string]any{"list": listRole, "count": count}
		result.SetMessage("操作成功")
		return result, nil
	}

	// 收集 RoleID 列表
	var roleIds []string
	for _, v := range roleDataList {
		roleIds = append(roleIds, v.RoleId)
	}

	// 查询角色关联的权限
	var rolePermissions []*model.CRMRolePermission
	if err := g.CRMRolePermission.Where(g.CRMRolePermission.RoleId.In(roleIds...)).Scan(&rolePermissions); err != nil {
		logObj.Errorw("RoleList Find Permissions", "error", err)
		// 权限查询失败不影响主列表，只是权限为空
	}

	// 组织权限数据 map[role_id][]permission_id
	permissionMap := make(map[string][]string)
	for _, v := range rolePermissions {
		permissionMap[v.RoleId] = append(permissionMap[v.RoleId], v.PermissionId)
	}

	for idx, _ := range roleDataList {
		roleId := roleDataList[idx].RoleId
		perms := permissionMap[roleId]
		if perms == nil {
			perms = []string{}
		}
		listRole = append(listRole, RespRoleService{
			RoleId:     roleId,
			RoleName:   roleDataList[idx].RoleName,
			Status:     roleDataList[idx].Status,
			IsInit:     roleDataList[idx].IsInit,
			CreatedAt:  roleDataList[idx].CreatedAt.Format("2006-01-02 15:04:05"),
			Permission: perms,
		})
	}

	result.Data = map[string]any{"list": listRole, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(offset, limit int64) ([]*model.CRMRole, int64, error) {
	var (
		crmRole  = g.CRMRole
		response = make([]*model.CRMRole, 0)
	)

	q := crmRole.Debug()
	where := []gen.Condition{}

	count, err := q.Where(where...).Order(crmRole.Id.Asc()).ScanByPage(&response, int(offset), int(limit))
	return response, count, err
}
