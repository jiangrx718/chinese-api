package minio

import (
	"context"
	"crm/gopkg/viper"

	"github.com/minio/minio-go/v7"
)

// DeleteObject 删除 MinIO 中的文件
func DeleteObject(objectName string) error {
	ctx := context.Background()

	err := minioClient.RemoveObject(ctx, viper.GetMinioCnf().Bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
