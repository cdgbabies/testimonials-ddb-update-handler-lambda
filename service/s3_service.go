package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type UploadContentsRequest struct {
	BucketName string
	Key        string
	Contents   string
}
type S3UploadObjectClient interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

func NewS3Client(ctx context.Context, region string) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %s", err)
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = region

	})

	return s3Client, nil

}

func UploadFileContents(ctx context.Context, s3Client S3UploadObjectClient, uploadContentsRequest UploadContentsRequest) error {
	log.Println(uploadContentsRequest)
	putObjectInput := &s3.PutObjectInput{
		Bucket: &uploadContentsRequest.BucketName,
		Key:    &uploadContentsRequest.Key,
		Body:   strings.NewReader(uploadContentsRequest.Contents),
	}
	_, err := s3Client.PutObject(ctx, putObjectInput)
	return err
}
