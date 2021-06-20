package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

/*
	A simple demo of the AWS SDK for Golang,
	Script will list the contents of an S3 bucket
*/

func main() {
	fmt.Println("✅ Starting S3 Demo.")
	// creating a context, don't know why this is used
	ctx := context.TODO()
	// Loading AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("❌ Encountered a problem!")
		os.Exit(1)
	}
	// creating an S3 client; the access credentials give access only to S3
	awsClientS3 := s3.NewFromConfig(cfg)
	// getting bucket results using ListObjectsV2
	output, err := awsClientS3.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("DEFAULT_BUCKET_NAME")),
	})
	if err != nil {
		fmt.Println("❌ Encountered a problem!")
	}
	if output.KeyCount == 0 {
		fmt.Println("❌ Bucket empty!")
	}
	for _, object := range output.Contents {
		fmt.Printf("key=%s\nsize=%d\n", aws.ToString(object.Key), object.Size)
	}
}
