package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type aliasEntry struct {
	Alias     string    `json:"alias"`
	IP        string    `json:"ip"`
	Timestamp time.Time `json:"timestamp"`
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	secret := os.Getenv("SECRET")
	if req.Headers["x-api-key"] != secret { // AWS gateway tranforms into lowercase header names?
		return events.APIGatewayProxyResponse{
			StatusCode: 401,
			Body:       "Unauthorized",
		}, nil
	}

	switch req.HTTPMethod {
	case "GET":
		return get(req)
	case "PUT":
		return put(req)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Not supported.",
		}, nil
	}
}

func put(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	alias := req.PathParameters["alias"]
	if len(alias) < 1 {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Invalid alias",
		}, nil
	}

	ae := aliasEntry{
		Alias:     alias,
		IP:        req.RequestContext.Identity.SourceIP,
		Timestamp: time.Now(),
	}
	err := putAlias(&ae)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func get(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	alias := req.PathParameters["alias"]
	if len(alias) < 1 {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Invalid alias",
		}, nil
	}

	ae, err := getAlias(alias)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	response, err := json.Marshal(ae)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(router)
}
