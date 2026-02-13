package minio

import (
	"context"
	"crm/gopkg/viper"
	"fmt"

	"github.com/minio/minio-go/v7"
)

// ListObjects 列出存储桶中的所有对象
func ListObjects() error {
	ctx := context.Background()

	bucketName := viper.GetMinioCnf().Bucket
	fmt.Printf("\n📁 存储桶 '%s' 中的文件列表:\n", bucketName)

	// 创建对象通道
	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})

	// 遍历对象
	count := 0
	for object := range objectCh {
		if object.Err != nil {
			return fmt.Errorf("列出对象时出错: %v", object.Err)
		}
		fmt.Printf("   - %s (大小: %d bytes, 最后修改: %s)\n",
			object.Key, object.Size, object.LastModified.Format("2006-01-02 15:04:05"))
		count++
	}

	if count == 0 {
		fmt.Println("   存储桶为空")
	} else {
		fmt.Printf("   共 %d 个文件\n", count)
	}

	return nil
}
