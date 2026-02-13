package minio

import (
	"context"
	"crm/gopkg/viper"
	"fmt"
	"time"
)

func GeneratePresignedURL(objectName string, expiry time.Duration) (string, error) {
	// 生成预签名URL，有效期设置为1小时
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), viper.GetMinioCnf().Bucket, objectName, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("生成预签名URL失败: %v", err)
	}

	return presignedURL.String(), nil
}
