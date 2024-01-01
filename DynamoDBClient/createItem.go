package main

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"log"
)

type Item struct {
	Year 	int
	Title	string
	Plot	string
	Rating	float64
}

// snippet-end:[dynamodb.go.create_item.struct]

func createItem() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDb Client
	svc := dynamodb.New(sess)

	// snippet-start:[dynamodb.go.create_item.assign_struct]
	item := Item{
		Year: 2015,
		Title: "The Big New Movie",
		Plot: "Noting happens at all.",
		Rating: 0.0,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	tableName := "Movies"

	input := &dynamodb.PutItemInput{
		Item: 		av,
		TableName: 	aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	year := strconv.Itoa(item.Year)

	fmt.Println("Successfully added '" + item.Title + "' ("+ year +") to table " + tableName)
	// snippet-end:[dynamodb.go.create_item.call]
}