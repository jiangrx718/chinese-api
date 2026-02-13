package storage

import (
	"strings"
)

func IsS3Path(path string) bool {
	return strings.HasPrefix(path, "s3://")
}
func IsWebPath(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}

func ParseS3Path(path string) bool {
	return strings.HasPrefix(path, "s3://")
}

func GetS3Dir(input string) string {
	lastSlash := strings.LastIndex(input, "/")
	if lastSlash == -1 {
		return ""
	}

	// 拆分路径
	dir := input[:lastSlash]
	return dir
}
