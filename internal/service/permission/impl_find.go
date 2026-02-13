package permission

import (
	"context"
	"crm/gopkg/cache/redis"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"
	"time"
)

type RespPermissionFindInfo struct {
	IsSuper       string   `json:"is_super"`
	PermissionIds []string `json:"permission_ids"`
}

// PermissionFind 查询权限
func (s *Service) PermissionFind(ctx context.Context, adminId string) (common.ServiceResult, error) {
	var (
		result = common.NewCRMServiceResult()
	)

	// 尝试从缓存获取
	cacheKey := fmt.Sprintf("login_auth:%s", adminId)
	rdb, err := redis.ClientAndErr("crm")
	if err == nil && rdb != nil {
		if val, err := rdb.Get(ctx, cacheKey).Result(); err == nil {
			// 缓存命中，直接返回成功
			result.Data = val
			result.SetMessage("操作成功")
			return result, nil
		}
	}

	// 查询数据库验证用户状态
	admin, err := g.CRMAdmin.WithContext(ctx).Where(
		g.CRMAdmin.AdminId.Eq(adminId),
		g.CRMAdmin.Status.Eq(model.StatusOn),
	).Take()

	if err != nil || admin == nil {
		return result, fmt.Errorf("用户不存在或已被禁用")
	}

	// 查询用户所处的角色
	roleId := admin.RoleId
	if roleId == "" {
		return result, fmt.Errorf("用户角色不存在")
	}

	// 查询对应的角色
	roleInfo, err := g.CRMRole.WithContext(ctx).Where(
		g.CRMRole.RoleId.Eq(roleId),
	).Take()
	if err != nil || roleInfo == nil {
		return result, fmt.Errorf("用户角色不存在")
	}

	// 查询角色所授予的权限
	rolePermission, err := g.CRMRolePermission.WithContext(ctx).Where(
		g.CRMRolePermission.RoleId.Eq(roleId),
	).Find()

	// 如果错误或者没有权限，并且不是超管角色
	if (err != nil || len(rolePermission) < 1) && roleInfo.IsSuper != model.IsSuperOn {
		return result, fmt.Errorf("用户角色对应的权限不存在")
	}

	var permissionIds []string
	for _, permission := range rolePermission {
		permissionIds = append(permissionIds, permission.PermissionId)
	}

	var respPermissionFind RespPermissionFindInfo
	respPermissionFind.IsSuper = roleInfo.IsSuper
	respPermissionFind.PermissionIds = permissionIds

	// 写入缓存，有效期 30 分钟
	if rdb != nil {
		rdb.Set(ctx, cacheKey, respPermissionFind, 30*time.Minute)
	}

	result.Data = respPermissionFind
	result.SetMessage("操作成功")
	return result, nil
}
