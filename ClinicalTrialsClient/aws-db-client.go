package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func createDb() {
    cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
        o.Region = "us-west-1"
        return nil
    })
    if err != nil {
        panic(err)
    }

	svc := dynamodb.NewFromConfig(cfg)
    out, err := svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
        AttributeDefinitions: []types.AttributeDefinition{
            {
                AttributeName: aws.String("id"),
                AttributeType: types.ScalarAttributeTypeS,
            },
        },
        KeySchema: []types.KeySchemaElement{
            {
                AttributeName: aws.String("id"),
                KeyType:       types.KeyTypeHash,
            },
        },
        TableName:   aws.String("my-table"),
        BillingMode: types.BillingModePayPerRequest,
    })
    if err != nil {
        panic(err)
    }

    fmt.Println(out)
}