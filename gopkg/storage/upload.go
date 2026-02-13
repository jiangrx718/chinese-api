package storage

import (
	"bytes"
	"context"
	"crm/gopkg/utils/files"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

func UploadByFileReader(projectName ProjectName, moduleName ModuleName, fileName files.FileName, file io.Reader) (string, error) {
	path := GenRealNameFilePath(projectName, moduleName, fileName)
	if err := UploadByReader(path.String(), file); err != nil {
		return "", err
	}

	return path.String(), nil
}

func UploadByReader(fullpath string, reader io.Reader) error {
	bucket, key, _, err := UriToBucketAndKey(fullpath)
	if err != nil {
		return err
	}

	_, err = s3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(fmt.Sprintf("%s/%s", bucket, key)),
		Body:   reader,
	})
	if err != nil {
		return err
	}

	return nil
}

func UploadByFileBytes(projectName ProjectName, moduleName ModuleName, fileName files.FileName, file []byte) (string, error) {
	path := GenRealNameFilePath(projectName, moduleName, fileName)
	if err := UploadByData(path.String(), file); err != nil {
		return "", err
	}

	return path.String(), nil
}

func UploadByData(fullpath string, data []byte) error {
	if data == nil {
		return fmt.Errorf("data is nil")
	}

	bucket, key, _, err := UriToBucketAndKey(fullpath)
	if err != nil {
		return err
	}

	// 上传
	_, err = s3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(fmt.Sprintf("%s/%s", bucket, key)),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return err
	}
	return nil
}
