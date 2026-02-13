package category

import (
	"context"
	"crm/gopkg/log"
	"crm/gopkg/utils"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
)

type RespCategoryInfo struct {
	CategoryId    string `json:"category_id"`
	CategoryName  string `json:"category_name"`
	CategoryImage string `json:"category_image"`
	ParentId      string `json:"parent_id"`
	Status        string `json:"status"`
	CategoryType  int    `json:"category_type"`
	Position      int    `json:"position"`
}

func (s *Service) CategoryCreate(ctx context.Context, parentId, categoryName, categoryImage, status string, categoryType, position int) (common.ServiceResult, error) {
	var (
		logObj = log.SugarContext(ctx)
		result = common.NewCRMServiceResult()
	)

	// 检查数据是否存在
	categoryEntity, err := g.CRMCategory.Where(
		g.CRMCategory.CategoryName.Eq(categoryName),
		g.CRMCategory.CategoryType.Eq(categoryType)).Take()

	if err != nil && err.Error() != "record not found" {
		logObj.Errorw("CategoryCreate Check Exist Error", "error", err)
		return result, err
	}
	if categoryEntity != nil {
		result.SetCode(10001) // 业务错误码
		result.SetMessage("栏目已存在")
		return result, nil // 返回 nil error，让 controller 处理 result
	}
	categoryId := utils.GenUUID()
	crmCategory := model.CRMCategory{
		CategoryType:  categoryType,
		CategoryId:    categoryId,
		CategoryName:  categoryName,
		CategoryImage: categoryImage,
		ParentId:      parentId,
		Position:      position,
		Status:        status,
	}

	if createErr := g.CRMCategory.Create(&crmCategory); createErr != nil {
		logObj.Errorw("CRMCategory Create crmCategory error", "crmCategory", crmCategory, "error", createErr)
		return result, createErr
	}

	result.Data = RespCategoryInfo{
		CategoryId:    categoryId,
		CategoryName:  categoryName,
		CategoryImage: categoryImage,
		ParentId:      parentId,
		Position:      position,
		Status:        status,
		CategoryType:  categoryType,
	}
	result.SetMessage("操作成功")
	return result, nil
}
