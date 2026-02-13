package minio

import (
	"context"
	"crm/gopkg/viper"

	"github.com/minio/minio-go/v7"
)

// UploadFile 上传本地文件到 MinIO
func UploadFile(objectName, filePath string) (info minio.UploadInfo, err error) {
	ctx := context.Background()

	// 上传文件
	info, err = minioClient.FPutObject(ctx, viper.GetMinioCnf().Bucket, objectName, filePath, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return
	}

	return
}
