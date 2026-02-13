package category

import (
	"context"
	"crm/gopkg/log"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"

	"gorm.io/gen"
)

type RespCategoryService struct {
	Id            int                    `json:"id"`
	CategoryId    string                 `json:"category_id"`
	CategoryName  string                 `json:"category_name"`
	CategoryImage string                 `json:"category_image"`
	ParentId      string                 `json:"parent_id"`
	Status        string                 `json:"status"`
	Position      int                    `json:"position"`
	CreatedAt     string                 `json:"created_at"`
	ChildList     []*RespCategoryService `json:"child_list"`
}

func (s *Service) CategoryList(ctx context.Context, offset, limit int64, categoryType int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	categoryDataList, count, err := ScanByPage(offset, limit, categoryType)
	if err != nil {
		logObj.Errorw("CategoryList Find", "error", err)
		result.SetError(&common.ServiceError{Code: -1, Message: "failed"})
		result.SetMessage("操作失败")
		return result, nil
	}

	// 构建 id -> 节点 映射
	idMap := make(map[string]*RespCategoryService, len(categoryDataList))
	for _, p := range categoryDataList {
		idMap[p.CategoryId] = &RespCategoryService{
			Id:            p.Id,
			CategoryId:    p.CategoryId,
			CategoryName:  p.CategoryName,
			CategoryImage: p.CategoryImage,
			ParentId:      p.ParentId,
			Status:        p.Status,
			Position:      p.Position,
			CreatedAt:     p.CreatedAt.Format("2006-01-02 15:04:05"),
			ChildList:     []*RespCategoryService{},
		}
	}

	// 组装树结构：parent_id 与 permission_id 对应
	var roots []*RespCategoryService
	for _, p := range categoryDataList {
		node := idMap[p.CategoryId]
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
		roots = []*RespCategoryService{}
	}

	result.Data = map[string]any{"list": roots, "count": count}
	result.SetMessage("操作成功")
	return result, nil
}

func ScanByPage(offset, limit int64, categoryType int) ([]*model.CRMCategory, int64, error) {
	var (
		crmCategory = g.CRMCategory
		response    = make([]*model.CRMCategory, 0)
	)

	q := crmCategory.Debug()
	where := []gen.Condition{
		g.CRMCategory.CategoryType.Eq(categoryType),
	}

	count, err := q.Where(where...).Order(crmCategory.Id.Asc()).ScanByPage(&response, int(offset), int(limit))
	return response, count, err
}
