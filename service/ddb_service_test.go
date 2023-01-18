package service

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type mockDynamoDBClient struct {
	response dynamodb.QueryOutput
	err      error
}

func (m *mockDynamoDBClient) Query(ctx context.Context, input *dynamodb.QueryInput, options ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	return &m.response, m.err
}

func TestQueryDynamoDB(t *testing.T) {
	// Test case 1: Successful query

	response := dynamodb.QueryOutput{
		Items: []map[string]types.AttributeValue{
			{

				"sk":          &types.AttributeValueMemberS{Value: "123"},
				"pk":          &types.AttributeValueMemberS{Value: "testimonials"},
				"author":      &types.AttributeValueMemberS{Value: "John Doe"},
				"testimonial": &types.AttributeValueMemberS{Value: "This is a great service!"},
				"createdDate": &types.AttributeValueMemberS{Value: "2022-01-01T00:00:00Z"},
			},
		},
	}
	client := &mockDynamoDBClient{response: response}
	testimonials, err := QueryDynamoDB(context.Background(), client)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(testimonials) != 1 {
		t.Errorf("Expected 1 testimonial, got %d", len(testimonials))
	}
	if testimonials[0].Author != "John Doe" {
		t.Errorf("Expected author to be John Doe, got %s", testimonials[0].Author)
	}
	if testimonials[0].CreatedDate.Format(time.RFC3339) != "2022-01-01T00:00:00Z" {
		t.Errorf("Expected created date to be 2022-01-01T00:00:00Z, got %s", testimonials[0].CreatedDate.Format(time.RFC3339))
	}
}
