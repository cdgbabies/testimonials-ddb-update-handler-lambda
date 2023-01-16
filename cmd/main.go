package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cdgbabies/testimonials-ddb-update-handler-lambda.git/service"
)

type handler struct {
	dynamoDbClient service.DynamoDbReadOperationClient
	s3Client       service.S3UploadObjectClient
}

func (h *handler) handleRequest(ctx context.Context, event events.DynamoDBEvent) error {

	eventJson, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJson)
	log.Println("Inside Handle Requet")
	testimonials, err := service.QueryDynamoDB(ctx, h.dynamoDbClient)
	if err != nil {
		return err
	}
	body, err := json.Marshal(testimonials)
	if err != nil {
		return err
	}
	err = service.UploadFileContents(ctx, h.s3Client, service.UploadContentsRequest{
		BucketName: os.Getenv("bucket_name"),
		Key:        os.Getenv("key_name"),
		Contents:   string(body),
	})
	return err
}
func main() {
	region_name := os.Getenv("region_name")

	dynamoDbClient, err := service.NewDynamoDbClient(context.Background(), region_name)
	if err != nil {
		panic(err)
	}
	s3Client, err := service.NewS3Client(context.Background(), region_name)
	if err != nil {
		panic(err)
	}

	h := handler{
		dynamoDbClient: dynamoDbClient,
		s3Client:       s3Client,
	}
	lambda.Start(h.handleRequest)
}
