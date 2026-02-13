package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadFile(fullpath string) ([]byte, error) {
	bucket, key, _, err := UriToBucketAndKey(fullpath)
	if err != nil {
		return nil, err
	}
	output, err := s3Client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(fmt.Sprintf("%s/%s", bucket, key)),
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = output.Body.Close()
	}()

	bytes, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
