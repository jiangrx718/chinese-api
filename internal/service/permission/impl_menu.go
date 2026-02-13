package permission

import (
	"context"
	"crm/gopkg/cache/redis"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"encoding/json"
	"fmt"
	"time"
)

// PermissionMenu 获取菜单
func (s *Service) PermissionMenu(ctx context.Context, adminId string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 1. 获取用户权限信息 (复用 PermissionFind 逻辑或直接查缓存)
	// var permInfo RespPermissionFindInfo
	cacheKeyMenu := fmt.Sprintf("login_menu:%s", adminId)
	rdb, err := redis.ClientAndErr("crm")
	if err == nil && rdb != nil {
		if val, err := rdb.Get(ctx, cacheKeyMenu).Result(); err == nil && val != "" {
			var menuData []*RespPermissionService
			err := json.Unmarshal([]byte(val), &menuData)
			if err != nil {

			}
			if len(menuData) > 0 {
				result.Data = map[string]any{"list": menuData}
				return result, nil
			}
		}
	}

	// 重新查询用户权限信息 (Copy from PermissionFind logic mostly)
	admin, err := g.CRMAdmin.WithContext(ctx).Where(
		g.CRMAdmin.AdminId.Eq(adminId),
		g.CRMAdmin.Status.Eq(model.StatusOn),
	).Take()
	if err != nil || admin == nil {
		return result, fmt.Errorf("用户不存在或已被禁用")
	}

	roleId := admin.RoleId
	roleInfo, err := g.CRMRole.WithContext(ctx).Where(
		g.CRMRole.RoleId.Eq(roleId),
	).Take()
	if err != nil || roleInfo == nil {
		return result, fmt.Errorf("用户角色不存在")
	}

	// 2. 根据 IsSuper 决定查询逻辑
	var permissionDataList []*model.CRMPermission

	if roleInfo.IsSuper == model.IsSuperOn {
		// 超管：查询所有开启状态的权限
		permissionDataList, _, err = ScanByPage(model.StatusOn)
		if err != nil {
			logObj.Errorw("PermissionMenu Super ScanByPage error", "error", err)
			return result, err
		}
		// 仅保留菜单类型（permission_type = 1），过滤掉接口类型(3)与按钮(2)
		tmp := make([]*model.CRMPermission, 0, len(permissionDataList))
		for _, p := range permissionDataList {
			if p.PermissionType == 1 {
				tmp = append(tmp, p)
			}
		}
		permissionDataList = tmp
	} else {
		// 普通用户：查询角色关联的权限 ID
		rolePermissions, err := g.CRMRolePermission.WithContext(ctx).Where(
			g.CRMRolePermission.RoleId.Eq(roleId),
		).Find()
		if err != nil {
			return result, err
		}

		var permissionIds []string
		for _, rp := range rolePermissions {
			permissionIds = append(permissionIds, rp.PermissionId)
		}

		if len(permissionIds) > 0 {
			// 查询具体的权限详情 (且状态开启)
			permissionDataList, err = g.CRMPermission.WithContext(ctx).Where(
				g.CRMPermission.PermissionId.In(permissionIds...),
				g.CRMPermission.Status.Eq(model.StatusOn),
				g.CRMPermission.PermissionType.Eq(1),
			).Order(g.CRMPermission.Position.Desc(), g.CRMPermission.Id.Desc()).Find()
			if err != nil {
				return result, err
			}
		}
	}

	// 3. 构建树形结构 (复用 PermissionList 的构建逻辑)
	// 将 model 转换为 resp 结构
	idMap := make(map[string]*RespPermissionService, len(permissionDataList))
	for _, p := range permissionDataList {
		idMap[p.PermissionId] = &RespPermissionService{
			PermissionId:   p.PermissionId,
			PermissionName: p.PermissionName,
			PermissionUrl:  p.PermissionURL,
			ParentId:       p.ParentId,
			Status:         p.Status,
			IsInit:         p.IsInit,
			Position:       p.Position,
			CreatedAt:      p.CreatedAt.Format("2006-01-02 15:04:05"),
			ChildList:      []*RespPermissionService{},
		}
	}

	var roots []*RespPermissionService
	for _, p := range permissionDataList {
		node := idMap[p.PermissionId]
		// 如果是根节点 (ParentId 为空)
		if p.ParentId == "" {
			roots = append(roots, node)
			continue
		}
		// 尝试挂载到父节点
		if parent, ok := idMap[p.ParentId]; ok && parent != nil {
			parent.ChildList = append(parent.ChildList, node)
		} else {
		}
	}

	if len(roots) == 0 {
		roots = []*RespPermissionService{}
	}

	if len(roots) > 0 && rdb != nil {
		bytes, _ := json.Marshal(roots)
		if err := rdb.Set(ctx, cacheKeyMenu, string(bytes), 30*time.Minute).Err(); err != nil {
			logObj.Errorw("PermissionMenu Cache Set", "error", err)
		}
	}
	result.Data = map[string]any{"list": roots}
	result.SetMessage("操作成功")
	return result, nil
}
