package str

import (
	"regexp"
	"strings"
)

// IsPhoneNumber 判断是否为手机号的基础正则
func IsPhoneNumber(phone string) bool {
	// 去除空格和特殊字符
	phone = strings.TrimSpace(phone)

	// 手机号正则（中国大陆）
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}
