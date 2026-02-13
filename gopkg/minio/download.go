package minio

import (
	"context"
	"crm/gopkg/viper"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
)

// DownloadFile 从 MinIO 下载文件到本地
func DownloadFile(objectName, filePath string) (err error) {
	ctx := context.Background()

	err = minioClient.FGetObject(ctx, viper.GetMinioCnf().Bucket, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return
	}

	return nil
}

// DownloadFileByte 从 MinIO 下载文件到本地
func DownloadFileByte(objectKey string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 获取对象 - 类似于 S3 的 GetObject
	object, err := minioClient.GetObject(ctx, viper.GetMinioCnf().Bucket, objectKey, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	// 读取对象数据到字节数组
	data, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}

	return data, nil
}
