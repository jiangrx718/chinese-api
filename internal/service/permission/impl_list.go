package permission

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespPermissionService struct {
	PermissionId   string                   `json:"permission_id"`
	PermissionName string                   `json:"permission_name"`
	PermissionUrl  string                   `json:"permission_url"`
	ParentId       string                   `json:"parent_id"`
	Status         string                   `json:"status"`
	IsInit         string                   `json:"is_init"`
	Position       int                      `json:"position"`
	PermissionType int                      `json:"permission_type"`
	CreatedAt      string                   `json:"created_at"`
	ChildList      []*RespPermissionService `json:"child_list"`
}

func (s *Service) PermissionList(ctx context.Context, status string) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	permissionDataList, count, err := ScanByPage(status)
	if err != nil {
		logObj.Errorw("PermissionList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	// 构建 id -> 节点 映射
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
			PermissionType: p.PermissionType,
			CreatedAt:      p.CreatedAt.Format("2006-01-02 15:04:05"),
			ChildList:      []*RespPermissionService{},
		}
	}

	// 组装树结构：parent_id 与 permission_id 对应
	var roots []*RespPermissionService
	for _, p := range permissionDataList {
		node := idMap[p.PermissionId]
		if p.ParentId == "" {
			roots = append(roots, node)
			continue
		}
		if parent, ok := idMap[p.ParentId]; ok && parent != nil {
			parent.ChildList = append(parent.ChildList, node)
		} else {
			// 如果找不到父节点，则作为根节点返回，避免数据丢失
			//roots = append(roots, node)
		}
	}

	if len(roots) == 0 {
		roots = []*RespPermissionService{}
	}
	result.Data = map[string]any{"list": roots, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(status string) ([]*model.CRMPermission, int64, error) {
	var (
		crmPermission = g.CRMPermission
		response      = make([]*model.CRMPermission, 0)
	)

	q := crmPermission.Debug()
	where := []gen.Condition{}
	if status != "" {
		where = append(where, crmPermission.Status.Eq(status))
	}

	count, err := q.Where(where...).Order(crmPermission.Position.Desc(), crmPermission.Id.Desc()).ScanByPage(&response, int(0), int(100))
	return response, count, err
}
