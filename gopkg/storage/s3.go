package storage

import (
	"context"
	"crm/gopkg/utils"
	"crm/gopkg/utils/files"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

const (
	uriPrefix = `s3://`
	separator = `/`
)

func S3() *s3.Client {
	return s3Client
}

func GetDatasetDocumentPath(key string) string {
	bucket := viper.GetString("aws.bucket")
	pathList := []string{
		bucket,
		key,
	}
	return uriPrefix + strings.Join(pathList, "/")
}

func UriToBucketAndKey(uri string) (bucket string, key string, isPrefix bool, err error) {

	if len(uri) == 0 {
		return "", "", false, errors.New("s3 loader source uri is empty")
	}

	if !strings.HasPrefix(uri, uriPrefix) {
		return "", "", false, fmt.Errorf("uri is not s3 uri, uri: %s", uri)
	}

	bucketAndKey := strings.TrimPrefix(uri, uriPrefix)
	bucketEnd := strings.Index(bucketAndKey, separator)
	if bucketEnd == -1 {
		return "", "", false, fmt.Errorf("s3 uri incomplete: %s", uri)
	}

	bucket = bucketAndKey[:bucketEnd]
	key = bucketAndKey[bucketEnd+1:]

	if strings.HasSuffix(key, separator) {
		return bucket, key, true, nil
	}

	return bucket, key, false, nil
}

func GeneratePresignPutPath(key string) (path string, err error) {
	presignClient := s3.NewPresignClient(s3Client)
	params := &s3.PutObjectInput{
		Bucket:      aws.String(viper.GetString("aws.bucket")),
		Key:         aws.String(key),
		ContentType: aws.String("application/octet-stream"),
	}
	object, err := presignClient.PresignPutObject(context.Background(), params)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`^https?://`)
	endpoint := re.ReplaceAllString(viper.GetString("aws.endpoint.url"), "")
	proxy := re.ReplaceAllString(viper.GetString("aws.proxy"), "")
	replacer := strings.NewReplacer(endpoint, proxy)
	res := replacer.Replace(object.URL)
	return res, nil
}

type ProjectName string

func (p ProjectName) String() string {
	return string(p)
}

type ModuleName string

func (m ModuleName) String() string {
	return string(m)
}

type Path string

func (p Path) String() string {
	return string(p)
}

func baseGenFilePath(projectName ProjectName, moduleName ModuleName, fileName files.FileName, rename bool) Path {
	if rename {
		fileName = fileName.GenSnowflakeFileName()
	}
	return Path(GetDatasetDocumentPath(fmt.Sprintf("%s/%s/%s/%s/%s/%s.%s",
		projectName.String(),
		moduleName.String(),
		time.Now().Format("2006"),
		time.Now().Format("0102"),
		folderSharding(utils.SnowflakeGenIntUUID()),
		fileName.Name(),
		fileName.Ext(),
	)))
}

func GenRenameFilePath(projectName ProjectName, moduleName ModuleName, fileName files.FileName) Path {
	return baseGenFilePath(projectName, moduleName, fileName, true)
}

func GenRealNameFilePath(projectName ProjectName, moduleName ModuleName, fileName files.FileName) Path {
	return baseGenFilePath(projectName, moduleName, fileName, false)
}

func folderSharding(uuid int64) string {
	return strconv.FormatInt(uuid%1000, 10)
}
