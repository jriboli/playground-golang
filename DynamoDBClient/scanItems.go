package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"fmt"
	"log"
)

func scanItems() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDb Client
	svc := dynamodb.New(sess)

	tableName := "Movies"
	minRating := 4.0
	year := 2013

	// Get all moves in that year, we'll pull out those with a higher rating later
	filt := expression.Name("Year").Equal(expression.Value(year))

	// Get back the title, year, and rating
	proj := expression.NamesList(expression.Name("Title"), expression.Name("Yeary"), expression.Name("Rating"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames: 	expr.Names(),
		ExpressionAttributeValues: 	expr.Values(),
		FilterExpression: 			expr.Filter(),
		ProjectionExpression: 		expr.Projection(),
		TableName: 					aws.String(tableName),
	}

	// Make the DynamoDb Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	numItems := 0

	for _, i := range result.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}

		if item.Rating > minRating {
			numItems ++

			fmt.Println("Title: ", item.Title)
			fmt.Println("Rating: ", item.Rating)
			fmt.Println()
		}
	}

	fmt.Println("Found", numItems, "movie(s) with a rating above", minRating, "in", year)
}