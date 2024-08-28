package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"lambda/pkg/bank"
	"net/http"
)

type ErrorBody struct {
	error string
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

	bankName := req.QueryStringParameters["BankName"]
	if len(bankName) > 0 {
		result, err := bank.GetBankMortgage(bankName, tableName, dynaClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{error: "Error Object"})
		}
		return apiResponse(http.StatusOK, result)
	}

	return nil, nil
}

func UnhandleMethods() (*events.APIGatewayProxyResponse, error) {

	return apiResponse(http.StatusOK, "result")
}
