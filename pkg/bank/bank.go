package bank

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"time"
)

var (
	ErrorFieldDatabase string = "database column error"
)

type BankRate struct {
	BankName  string    `json:"BankName"`
	RateType  string    `json:"RateType"`
	Rate      string    `json:"Rate"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func GetBankMortgage(bankName string, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*BankRate, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"BankName": {
				S: aws.String(bankName),
			},
		},
		TableName: aws.String(tableName),
	}
	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFieldDatabase)
	}

	item := new(BankRate)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFieldDatabase)
	}
	return item, nil
}
