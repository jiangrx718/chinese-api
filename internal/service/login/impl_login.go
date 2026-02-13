package login

import (
	"context"
	"crm/gopkg/auth"
	"crm/gopkg/utils/str"
	"crm/internal/common"
	"crm/internal/g"
	"crm/internal/model"
	"fmt"

	"gorm.io/gen"
)

type RespLogin struct {
	Token    string `json:"token"`
	AdminId  string `json:"admin_id"`
	UserName string `json:"user_name"`
}

func (s *Service) Login(ctx context.Context, userName, password string) (common.ServiceResult, error) {
	var (
		result = common.NewCRMServiceResult()
	)

	// 查询用户并校验密码与状态
	where := []gen.Condition{
		g.CRMAdmin.Password.Eq(str.MD5String(fmt.Sprintf("%s%s", password, model.SaltValue))),
		g.CRMAdmin.Status.Eq(model.StatusOn),
	}
	isMobile := str.IsPhoneNumber(userName)
	if isMobile {
		where = append(where, g.CRMAdmin.UserPhone.Eq(userName))
	} else {
		where = append(where, g.CRMAdmin.UserName.Eq(userName))
	}
	admin, err := g.CRMAdmin.Where(where...).Take()
	if err != nil {
		return result, fmt.Errorf("用户名或密码错误")
	}
	if admin == nil {
		result.SetCode(401)
		result.SetMessage("用户名或密码错误")
		return result, nil
	}

	// 生成 JWT
	token, err := auth.GenerateToken(admin.AdminId)
	if err != nil {
		return result, err
	}

	// 返回 token 与基础用户信息
	result.Data = RespLogin{
		Token:    token,
		AdminId:  admin.AdminId,
		UserName: userName,
	}
	result.SetMessage("登录成功")
	return result, nil
}
