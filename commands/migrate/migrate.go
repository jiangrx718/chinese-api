package migrate

import (
	"crm/gopkg/gorms"
	"crm/internal/model"
	"fmt"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func doSeed(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 角色-超级管理员
		role := model.CRMRole{
			RoleId:   "c10a9fdd-a7e4-4e51-8633-50368cc765d2",
			RoleName: "超级管理员",
			Status:   model.StatusOn,
			IsInit:   model.IsInitOn,
			IsSuper:  model.IsSuperOn,
		}
		var roleExist model.CRMRole
		_ = tx.Where("role_id = ?", role.RoleId).Take(&roleExist).Error
		if roleExist.Id == 0 {
			if err := tx.Create(&role).Error; err != nil {
				return fmt.Errorf("seed role error: %w", err)
			}
		}

		// 管理员-admin
		admin := model.CRMAdmin{
			AdminId:   "6828a60c-c9e6-48b4-82af-d6c0909a2230",
			UserName:  "admin",
			UserPhone: "18210605516",
			Password:  "a9a43ccea31e4bbce49c86cad83b8604",
			Status:    model.StatusOn,
			RoleId:    role.RoleId,
			IsInit:    model.IsInitOn,
		}
		var adminExist model.CRMAdmin
		_ = tx.Where("admin_id = ?", admin.AdminId).Take(&adminExist).Error
		if adminExist.Id == 0 {
			if err := tx.Create(&admin).Error; err != nil {
				return fmt.Errorf("seed admin error: %w", err)
			}
		}

		// 权限-菜单
		perms := []model.CRMPermission{
			{PermissionId: "3399376f-3729-4cb5-9cdb-067a34edfa13", PermissionName: "首页", PermissionURL: "/home", ParentId: "", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOn},
			{PermissionId: "1aaec075-e246-4323-9af8-ae40ae4c35fa", PermissionName: "管理权限", PermissionURL: "/", ParentId: "", Status: model.StatusOn, PermissionType: 1, Position: 99, IsInit: model.IsInitOn},
			{PermissionId: "a729e8bf-0b28-4e36-8145-39fd835e406b", PermissionName: "管理员列表", PermissionURL: "/admins", ParentId: "1aaec075-e246-4323-9af8-ae40ae4c35fa", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOn},
			{PermissionId: "02fe91da-411d-4d82-9370-d03681be6f85", PermissionName: "角色管理", PermissionURL: "/roles", ParentId: "1aaec075-e246-4323-9af8-ae40ae4c35fa", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOn},
			{PermissionId: "7bf38053-095c-412c-a853-f1df4a09bc0d", PermissionName: "权限设置", PermissionURL: "/permissions", ParentId: "1aaec075-e246-4323-9af8-ae40ae4c35fa", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOn},
			{PermissionId: "dde741e1-0842-4f06-a04b-6ef5ea9e5710", PermissionName: "商品管理", PermissionURL: "/", ParentId: "", Status: model.StatusOn, PermissionType: 1, Position: 98, IsInit: model.IsInitOn},
			{PermissionId: "49dcd7f9-c183-46bb-8d57-ba4fe8125689", PermissionName: "商品分类", PermissionURL: "/product-category", ParentId: "dde741e1-0842-4f06-a04b-6ef5ea9e5710", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOn},
			{PermissionId: "777a715b-9399-42de-a7cb-d6aaaf4c8922", PermissionName: "商品列表", PermissionURL: "/product-list", ParentId: "dde741e1-0842-4f06-a04b-6ef5ea9e5710", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOn},
			{PermissionId: "586c2f0d-4be9-46cf-9433-6d67d8dac3f9", PermissionName: "内容管理", PermissionURL: "/", ParentId: "", Status: model.StatusOn, PermissionType: 1, Position: 97, IsInit: model.IsInitOff},
			{PermissionId: "5b47fd37-5120-4a7f-8d3d-69dc94d2f6d9", PermissionName: "文章列表", PermissionURL: "/article-list", ParentId: "586c2f0d-4be9-46cf-9433-6d67d8dac3f9", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "0607f147-f576-44a2-8ff5-155518ae66d8", PermissionName: "文章分类", PermissionURL: "/article-category", ParentId: "586c2f0d-4be9-46cf-9433-6d67d8dac3f9", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "c3570ddc-9728-4b7d-914f-c65ce6c1c0d3", PermissionName: "订单管理", PermissionURL: "/", ParentId: "", Status: model.StatusOn, PermissionType: 1, Position: 96, IsInit: model.IsInitOff},
			{PermissionId: "130a96ef-fe2b-45bb-83e3-d0fc0a211bcb", PermissionName: "订单列表", PermissionURL: "/order-list", ParentId: "c3570ddc-9728-4b7d-914f-c65ce6c1c0d3", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "0f980aee-bc09-428e-9269-f9100ccf492c", PermissionName: "订单统计", PermissionURL: "/order-statistics", ParentId: "c3570ddc-9728-4b7d-914f-c65ce6c1c0d3", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "ebfccbec-3825-4805-8a47-19a2413494b7", PermissionName: "系统设置", PermissionURL: "/", ParentId: "", Status: model.StatusOn, PermissionType: 1, Position: 95, IsInit: model.IsInitOff},
			// 追加：系统设置下的基础设置与协议设置（菜单）
			{PermissionId: "6c197270-1f67-42f2-affe-ae81032f3f8c", PermissionName: "基础设置", PermissionURL: "/base-settings", ParentId: "ebfccbec-3825-4805-8a47-19a2413494b7", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "16f79612-1168-4ffe-bd69-bc225eea334d", PermissionName: "协议设置", PermissionURL: "/agreement-settings", ParentId: "ebfccbec-3825-4805-8a47-19a2413494b7", Status: model.StatusOn, PermissionType: 1, Position: 100, IsInit: model.IsInitOff},
			// 追加：权限管理接口（接口类型）
			{PermissionId: "c0a177aa-2270-43ca-a73e-e3f0d1238d6a", PermissionName: "列表", PermissionURL: "/api/permission/list", ParentId: "7bf38053-095c-412c-a853-f1df4a09bc0d", Status: model.StatusOn, PermissionType: 3, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "1177c90d-1927-45ff-840e-f368825d335b", PermissionName: "修改", PermissionURL: "/api/permission/update", ParentId: "7bf38053-095c-412c-a853-f1df4a09bc0d", Status: model.StatusOn, PermissionType: 3, Position: 98, IsInit: model.IsInitOff},
			{PermissionId: "b882de06-04c3-402a-b928-0093f7f5d643", PermissionName: "创建", PermissionURL: "/api/permission/create", ParentId: "7bf38053-095c-412c-a853-f1df4a09bc0d", Status: model.StatusOn, PermissionType: 3, Position: 99, IsInit: model.IsInitOff},
			{PermissionId: "e169436d-0675-4a8a-b7e5-38f6fe27c98f", PermissionName: "删除", PermissionURL: "/api/permission/delete", ParentId: "7bf38053-095c-412c-a853-f1df4a09bc0d", Status: model.StatusOn, PermissionType: 3, Position: 96, IsInit: model.IsInitOff},
			{PermissionId: "156128a7-5b48-4726-9593-0fde91b9a6b8", PermissionName: "状态", PermissionURL: "/api/permission/status", ParentId: "7bf38053-095c-412c-a853-f1df4a09bc0d", Status: model.StatusOn, PermissionType: 3, Position: 95, IsInit: model.IsInitOff},
			// 追加：角色管理接口（接口类型）
			{PermissionId: "83c20165-8d46-48a1-b9f1-6b3525d422c1", PermissionName: "列表", PermissionURL: "/api/role/list", ParentId: "02fe91da-411d-4d82-9370-d03681be6f85", Status: model.StatusOn, PermissionType: 3, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "1a91e2d8-4dc0-4876-b66d-392b7fdef95e", PermissionName: "创建", PermissionURL: "/api/role/create", ParentId: "02fe91da-411d-4d82-9370-d03681be6f85", Status: model.StatusOn, PermissionType: 3, Position: 99, IsInit: model.IsInitOff},
			{PermissionId: "f9f2c8e2-e5fd-45cd-9d90-f1ee38657949", PermissionName: "修改", PermissionURL: "/api/role/update", ParentId: "02fe91da-411d-4d82-9370-d03681be6f85", Status: model.StatusOn, PermissionType: 3, Position: 98, IsInit: model.IsInitOff},
			{PermissionId: "dc972189-abf5-477b-9b59-a84715b7ab23", PermissionName: "删除", PermissionURL: "/api/role/delete", ParentId: "02fe91da-411d-4d82-9370-d03681be6f85", Status: model.StatusOn, PermissionType: 3, Position: 96, IsInit: model.IsInitOff},
		}
		for _, p := range perms {
			var exist model.CRMPermission
			_ = tx.Where("permission_id = ?", p.PermissionId).Take(&exist).Error
			if exist.Id == 0 {
				if err := tx.Create(&p).Error; err != nil {
					return fmt.Errorf("seed permission error: %w", err)
				}
			}
		}

		// 追加：管理员接口
		adminPerms := []model.CRMPermission{
			{PermissionId: "4d269ddb-eb6c-41ba-ba92-b67402cd4c68", PermissionName: "列表", PermissionURL: "/api/admin/list", ParentId: "a729e8bf-0b28-4e36-8145-39fd835e406b", Status: model.StatusOn, PermissionType: 3, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "e05cca1f-28b3-47dc-9063-e6b776090263", PermissionName: "创建", PermissionURL: "/api/admin/create", ParentId: "a729e8bf-0b28-4e36-8145-39fd835e406b", Status: model.StatusOn, PermissionType: 3, Position: 99, IsInit: model.IsInitOff},
			{PermissionId: "852c2d86-590e-4867-9e2b-258840219240", PermissionName: "修改", PermissionURL: "/api/admin/update", ParentId: "a729e8bf-0b28-4e36-8145-39fd835e406b", Status: model.StatusOn, PermissionType: 3, Position: 98, IsInit: model.IsInitOff},
			{PermissionId: "b1268a06-7bef-4fa4-816e-5717cb9fc0c7", PermissionName: "删除", PermissionURL: "/api/admin/delete", ParentId: "a729e8bf-0b28-4e36-8145-39fd835e406b", Status: model.StatusOn, PermissionType: 3, Position: 97, IsInit: model.IsInitOff},
		}
		for _, p := range adminPerms {
			var exist model.CRMPermission
			_ = tx.Where("permission_id = ?", p.PermissionId).Take(&exist).Error
			if exist.Id == 0 {
				if err := tx.Create(&p).Error; err != nil {
					return fmt.Errorf("seed permission(admin) error: %w", err)
				}
			}
		}

		// 追加：商品列表接口（在商品列表菜单下）
		productListAPIs := []model.CRMPermission{
			{PermissionId: "7653c4d9-0c79-4735-a8cf-7c51bf366cdd", PermissionName: "列表", PermissionURL: "", ParentId: "777a715b-9399-42de-a7cb-d6aaaf4c8922", Status: model.StatusOn, PermissionType: 3, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "fd7fef22-8491-404f-bbf7-7452634df684", PermissionName: "创建", PermissionURL: "", ParentId: "777a715b-9399-42de-a7cb-d6aaaf4c8922", Status: model.StatusOn, PermissionType: 3, Position: 98, IsInit: model.IsInitOff},
			{PermissionId: "58121cc5-c17a-47bc-8f2f-a000ed5b12d8", PermissionName: "修改", PermissionURL: "", ParentId: "777a715b-9399-42de-a7cb-d6aaaf4c8922", Status: model.StatusOn, PermissionType: 3, Position: 98, IsInit: model.IsInitOff},
			{PermissionId: "6648b050-0569-44f8-a6e5-297e151c77fe", PermissionName: "删除", PermissionURL: "", ParentId: "777a715b-9399-42de-a7cb-d6aaaf4c8922", Status: model.StatusOn, PermissionType: 3, Position: 96, IsInit: model.IsInitOff},
		}
		for _, p := range productListAPIs {
			var exist model.CRMPermission
			_ = tx.Where("permission_id = ?", p.PermissionId).Take(&exist).Error
			if exist.Id == 0 {
				if err := tx.Create(&p).Error; err != nil {
					return fmt.Errorf("seed permission(product-list apis) error: %w", err)
				}
			}
		}

		// 追加：商品分类接口（在商品分类菜单下）
		productCategoryAPIs := []model.CRMPermission{
			{PermissionId: "d7934ece-9734-4ae7-93bb-d7a95610eba2", PermissionName: "列表", PermissionURL: "api/category/product/list", ParentId: "49dcd7f9-c183-46bb-8d57-ba4fe8125689", Status: model.StatusOn, PermissionType: 3, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "1fcff449-22fe-49ab-8ee7-c99326fb33ed", PermissionName: "修改", PermissionURL: "/api/category/product/update", ParentId: "49dcd7f9-c183-46bb-8d57-ba4fe8125689", Status: model.StatusOn, PermissionType: 3, Position: 99, IsInit: model.IsInitOff},
			{PermissionId: "85a8f972-996a-4d30-951b-0df3a82a76e6", PermissionName: "创建", PermissionURL: "/api/category/product/create", ParentId: "49dcd7f9-c183-46bb-8d57-ba4fe8125689", Status: model.StatusOn, PermissionType: 3, Position: 99, IsInit: model.IsInitOff},
			{PermissionId: "727ed34f-1c35-4da8-9881-134b42ef5f39", PermissionName: "删除", PermissionURL: "/api/category/product/delete", ParentId: "49dcd7f9-c183-46bb-8d57-ba4fe8125689", Status: model.StatusOn, PermissionType: 3, Position: 97, IsInit: model.IsInitOff},
			{PermissionId: "61e37028-2452-4a61-8d87-5f1844dfae05", PermissionName: "状态", PermissionURL: "/api/category/product/status", ParentId: "49dcd7f9-c183-46bb-8d57-ba4fe8125689", Status: model.StatusOn, PermissionType: 3, Position: 95, IsInit: model.IsInitOff},
		}
		for _, p := range productCategoryAPIs {
			var exist model.CRMPermission
			_ = tx.Where("permission_id = ?", p.PermissionId).Take(&exist).Error
			if exist.Id == 0 {
				if err := tx.Create(&p).Error; err != nil {
					return fmt.Errorf("seed permission(product-category apis) error: %w", err)
				}
			}
		}

		// 追加：文章分类接口（在文章分类菜单下）
		articleCategoryAPIs := []model.CRMPermission{
			{PermissionId: "0746ea7a-ba36-40ff-a58c-79058cb1dc96", PermissionName: "列表", PermissionURL: "/api/category/article/list", ParentId: "0607f147-f576-44a2-8ff5-155518ae66d8", Status: model.StatusOn, PermissionType: 3, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "bb1293f0-a94e-4cc6-a3bb-a7986fe71481", PermissionName: "创建", PermissionURL: "/api/category/article/create", ParentId: "0607f147-f576-44a2-8ff5-155518ae66d8", Status: model.StatusOn, PermissionType: 3, Position: 99, IsInit: model.IsInitOff},
			{PermissionId: "65ee364d-adfa-449d-b3b1-e6a2b517a9fc", PermissionName: "修改", PermissionURL: "/api/category/article/update", ParentId: "0607f147-f576-44a2-8ff5-155518ae66d8", Status: model.StatusOn, PermissionType: 3, Position: 98, IsInit: model.IsInitOff},
			{PermissionId: "79b94916-16d7-4ea0-8a8f-b8e6b69c9787", PermissionName: "删除", PermissionURL: "/api/category/article/delete", ParentId: "0607f147-f576-44a2-8ff5-155518ae66d8", Status: model.StatusOn, PermissionType: 3, Position: 97, IsInit: model.IsInitOff},
			{PermissionId: "087e4f21-d3e1-483e-a552-e4d00aae6564", PermissionName: "状态", PermissionURL: "/api/category/article/status", ParentId: "0607f147-f576-44a2-8ff5-155518ae66d8", Status: model.StatusOn, PermissionType: 3, Position: 96, IsInit: model.IsInitOff},
		}
		for _, p := range articleCategoryAPIs {
			var exist model.CRMPermission
			_ = tx.Where("permission_id = ?", p.PermissionId).Take(&exist).Error
			if exist.Id == 0 {
				if err := tx.Create(&p).Error; err != nil {
					return fmt.Errorf("seed permission(article-category apis) error: %w", err)
				}
			}
		}

		// 追加：文章接口（在文章列表菜单下）
		articleAPIs := []model.CRMPermission{
			{PermissionId: "ff92c210-f133-4db2-80c0-afc4512deb5d", PermissionName: "列表", PermissionURL: "/api/article/list", ParentId: "5b47fd37-5120-4a7f-8d3d-69dc94d2f6d9", Status: model.StatusOn, PermissionType: 3, Position: 100, IsInit: model.IsInitOff},
			{PermissionId: "638b2dda-a8e0-4559-a6e0-a087456baab3", PermissionName: "创建", PermissionURL: "/api/article/update", ParentId: "5b47fd37-5120-4a7f-8d3d-69dc94d2f6d9", Status: model.StatusOn, PermissionType: 3, Position: 99, IsInit: model.IsInitOff},
			{PermissionId: "38a145aa-4a3d-4f26-877c-328fd2b63510", PermissionName: "修改", PermissionURL: "/api/article/update", ParentId: "5b47fd37-5120-4a7f-8d3d-69dc94d2f6d9", Status: model.StatusOn, PermissionType: 3, Position: 98, IsInit: model.IsInitOff},
			{PermissionId: "76d98b9a-a08d-49d7-90fa-31aff97e5f00", PermissionName: "删除", PermissionURL: "/api/article/delete", ParentId: "5b47fd37-5120-4a7f-8d3d-69dc94d2f6d9", Status: model.StatusOn, PermissionType: 3, Position: 97, IsInit: model.IsInitOff},
			{PermissionId: "b2fc5c06-5209-44f4-bd41-be604943ddfd", PermissionName: "状态", PermissionURL: "/api/article/status", ParentId: "5b47fd37-5120-4a7f-8d3d-69dc94d2f6d9", Status: model.StatusOn, PermissionType: 3, Position: 95, IsInit: model.IsInitOff},
		}
		for _, p := range articleAPIs {
			var exist model.CRMPermission
			_ = tx.Where("permission_id = ?", p.PermissionId).Take(&exist).Error
			if exist.Id == 0 {
				if err := tx.Create(&p).Error; err != nil {
					return fmt.Errorf("seed permission(article apis) error: %w", err)
				}
			}
		}

		// 追加：权限菜单接口（基础）
		permissionMenuAPI := model.CRMPermission{
			PermissionId:   "18c44bef-654c-4a2b-82b6-e06ac0002451",
			PermissionName: "基础",
			PermissionURL:  "/api/permission/menu",
			ParentId:       "",
			Status:         model.StatusOn,
			PermissionType: 3,
			Position:       100,
			IsInit:         model.IsInitOn,
		}
		{
			var exist model.CRMPermission
			_ = tx.Where("permission_id = ?", permissionMenuAPI.PermissionId).Take(&exist).Error
			if exist.Id == 0 {
				if err := tx.Create(&permissionMenuAPI).Error; err != nil {
					return fmt.Errorf("seed permission(permission menu api) error: %w", err)
				}
			}
		}

		return nil
	})
}

func Command() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "数据库迁移",
		Subcommands: []*cli.Command{
			{
				Name:        "up",
				Usage:       "自动迁移数据库",
				Description: "自动迁移数据库",
				Action: func(ctx *cli.Context) error {
					tx := gorms.Client()
					tx.DisableForeignKeyConstraintWhenMigrating = true
					tables := []any{
						model.CRMAdmin{},
						model.CRMPermission{},
						model.CRMRole{},
						model.CRMRolePermission{},
						model.CRMArticle{},
						model.CRMCategory{},
						model.CRMArticleContent{},
					}
					if err := tx.AutoMigrate(tables...); err != nil {
						return err
					}
					return doSeed(tx)
				},
			},
			{
				Name:        "seed",
				Usage:       "初始化基础数据",
				Description: "插入管理员、角色与权限等初始数据",
				Action: func(ctx *cli.Context) error {
					return doSeed(gorms.Client())
				},
			},
		},
	}
}
