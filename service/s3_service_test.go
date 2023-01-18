package service

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type mockS3Client struct {
	err error
}

func (m *mockS3Client) PutObject(ctx context.Context, input *s3.PutObjectInput, options ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return nil, m.err
}

func TestUploadFileContents(t *testing.T) {
	// Test case 1: Successful upload
	s3Client := &mockS3Client{}
	uploadContentsRequest := UploadContentsRequest{
		BucketName: "test-bucket",
		Key:        "test-key",
		Contents:   "test contents",
	}
	err := UploadFileContents(context.Background(), s3Client, uploadContentsRequest)
	if err != nil {
		t.Errorf("Unable to upload file contents")
	}
}
