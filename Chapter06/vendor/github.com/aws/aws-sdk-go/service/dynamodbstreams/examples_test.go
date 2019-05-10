// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodbstreams_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDynamoDBStreams_DescribeStream() {
	sess := session.Must(session.NewSession())

	svc := dynamodbstreams.New(sess)

	params := &dynamodbstreams.DescribeStreamInput{
		StreamArn:             aws.String("StreamArn"), // Required
		ExclusiveStartShardId: aws.String("ShardId"),
		Limit: aws.Int64(1),
	}
	resp, err := svc.DescribeStream(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDBStreams_GetRecords() {
	sess := session.Must(session.NewSession())

	svc := dynamodbstreams.New(sess)

	params := &dynamodbstreams.GetRecordsInput{
		ShardIterator: aws.String("ShardIterator"), // Required
		Limit:         aws.Int64(1),
	}
	resp, err := svc.GetRecords(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDBStreams_GetShardIterator() {
	sess := session.Must(session.NewSession())

	svc := dynamodbstreams.New(sess)

	params := &dynamodbstreams.GetShardIteratorInput{
		ShardId:           aws.String("ShardId"),           // Required
		ShardIteratorType: aws.String("ShardIteratorType"), // Required
		StreamArn:         aws.String("StreamArn"),         // Required
		SequenceNumber:    aws.String("SequenceNumber"),
	}
	resp, err := svc.GetShardIterator(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDBStreams_ListStreams() {
	sess := session.Must(session.NewSession())

	svc := dynamodbstreams.New(sess)

	params := &dynamodbstreams.ListStreamsInput{
		ExclusiveStartStreamArn: aws.String("StreamArn"),
		Limit:     aws.Int64(1),
		TableName: aws.String("TableName"),
	}
	resp, err := svc.ListStreams(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
