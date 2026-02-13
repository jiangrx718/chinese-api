package minio

import (
	"crm/gopkg/viper"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient *minio.Client
)

func Init() error {
	// 初始化 MinIO 客户端
	client, err := minio.New(viper.GetMinioCnf().Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(viper.GetMinioCnf().SecretId, viper.GetMinioCnf().SecretKey, ""),
		Secure: viper.GetMinioCnf().UseSSL,
	})
	if err != nil {
		return err
	}
	minioClient = client

	return nil
}
