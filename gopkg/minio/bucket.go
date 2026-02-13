package minio

import (
	"context"
	"crm/gopkg/viper"

	"github.com/minio/minio-go/v7"
)

// CreateBucket 检查并创建存储桶
func CreateBucket() error {
	ctx := context.Background()

	bucketName := viper.GetMinioCnf().Bucket
	// 检查存储桶是否存在
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if !exists {
		// 创建存储桶
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
