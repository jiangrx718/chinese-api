package storage

import (
	"crm/gopkg/viper"
	"crypto/tls"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"net/http"
)

var (
	s3Client *s3.Client
)

func Init() error {
	cfg := aws.Config{
		BaseEndpoint: aws.String(viper.GetAws().Endpoint.URL),
		Region:       viper.GetAws().Region,
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: viper.GetAws().Client.SkipVerify,
				},
			},
		},
	}

	accessKey := viper.GetAws().Secret.ID
	secretKey := viper.GetAws().Secret.Key
	if accessKey != "" && secretKey != "" {
		cfg.Credentials = &credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: secretKey,
			},
		}
	} else {
		cfg.Credentials = aws.AnonymousCredentials{}
	}

	s3Client = s3.NewFromConfig(cfg)
	return nil
}
