package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

/*
	This go file returns the tables
	that are stored in the database
*/

func getTablesFromDynamoDb() {
	fmt.Println("✅ Starting task getTablesFromDynamoDb....")
	ctx := context.TODO()
	//creating a configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-south-1"))

	if err != nil {
		fmt.Println("❌ Encountered a problem: ", err)
	}
	//creating a DynamoDb client
	awsClientDynamoDB := dynamodb.NewFromConfig(cfg)
	output, err := awsClientDynamoDB.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("❌ Unable to read tables in DB: ",err)
	}
	for _, object := range output.TableNames {
		fmt.Println(object)
	}
}
