package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Testimonial struct {
	Sk          string    `json:"sk" dynamodbav:"sk"`
	Pk          string    `json:"pk" dynamodbav:"pk"`
	Author      string    `json:"author" dynamodbav:"author"`
	Testimonial string    `json:"testimonial" dynamodbav:"testimonial"`
	CreatedDate time.Time `json:"createdDate" dynamodbav:"createdDate"`
}
type DynamoDbReadOperationClient interface {
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

func NewDynamoDbClient(ctx context.Context, region string) (*dynamodb.Client, error) {

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %s", err)
	}

	dynamoDbClient := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.Region = region

	})

	return dynamoDbClient, nil

}

func QueryDynamoDB(ctx context.Context, client DynamoDbReadOperationClient) ([]Testimonial, error) {

	var testimonials []Testimonial

	keyEx := expression.Key("pk").Equal(expression.Value("testimonials"))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	response, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String("CdgDynamicContents"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		return nil, err
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &testimonials)
		log.Println()
		if err != nil {
			log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
		}
	}
	return testimonials, nil
}
