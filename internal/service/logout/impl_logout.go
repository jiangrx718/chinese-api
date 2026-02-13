package logout

import (
	"context"
	"crm/gopkg/cache/redis"
	"crm/internal/common"
	"fmt"
)

func (s *Service) Logout(ctx context.Context, userId string) (common.ServiceResult, error) {
	var (
		result = common.NewCRMServiceResult()
	)
	cacheKeyMenu := fmt.Sprintf("login_menu:%s", userId)
	rdb, err := redis.ClientAndErr("crm")
	if err == nil && rdb != nil {
		rdb.Del(ctx, cacheKeyMenu)
	}
	result.SetMessage("退出成功")
	return result, nil
}
