package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"lambda/internal/handler"
)

func main() {
	lambda.Start(handler.Handler)
}
