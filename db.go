package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"time"
)

var region = os.Getenv("REGION")
var tableName = os.Getenv("TABLE_NAME")
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(region))

func getAlias(alias string) (*aliasEntry, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"alias": {
				S: aws.String(alias),
			},
		},
	}

	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	ae := new(aliasEntry)
	err = dynamodbattribute.UnmarshalMap(result.Item, ae)
	if err != nil {
		return nil, err
	}

	return ae, nil
}

func putAlias(ae *aliasEntry) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"alias": {
				S: aws.String(ae.Alias),
			},
			"ip": {
				S: aws.String(ae.IP),
			},
			"timestamp": {
				S: aws.String(ae.Timestamp.Format(time.RFC3339)),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
