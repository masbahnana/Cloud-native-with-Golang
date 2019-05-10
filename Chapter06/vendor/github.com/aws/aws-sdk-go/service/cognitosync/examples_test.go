// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitosync"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCognitoSync_BulkPublish() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.BulkPublishInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.BulkPublish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_DeleteDataset() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.DeleteDatasetInput{
		DatasetName:    aws.String("DatasetName"),    // Required
		IdentityId:     aws.String("IdentityId"),     // Required
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DeleteDataset(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_DescribeDataset() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.DescribeDatasetInput{
		DatasetName:    aws.String("DatasetName"),    // Required
		IdentityId:     aws.String("IdentityId"),     // Required
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DescribeDataset(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_DescribeIdentityPoolUsage() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.DescribeIdentityPoolUsageInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DescribeIdentityPoolUsage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_DescribeIdentityUsage() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.DescribeIdentityUsageInput{
		IdentityId:     aws.String("IdentityId"),     // Required
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DescribeIdentityUsage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_GetBulkPublishDetails() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.GetBulkPublishDetailsInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.GetBulkPublishDetails(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_GetCognitoEvents() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.GetCognitoEventsInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.GetCognitoEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_GetIdentityPoolConfiguration() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.GetIdentityPoolConfigurationInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.GetIdentityPoolConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_ListDatasets() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.ListDatasetsInput{
		IdentityId:     aws.String("IdentityId"),     // Required
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		MaxResults:     aws.Int64(1),
		NextToken:      aws.String("String"),
	}
	resp, err := svc.ListDatasets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_ListIdentityPoolUsage() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.ListIdentityPoolUsageInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("String"),
	}
	resp, err := svc.ListIdentityPoolUsage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_ListRecords() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.ListRecordsInput{
		DatasetName:      aws.String("DatasetName"),    // Required
		IdentityId:       aws.String("IdentityId"),     // Required
		IdentityPoolId:   aws.String("IdentityPoolId"), // Required
		LastSyncCount:    aws.Int64(1),
		MaxResults:       aws.Int64(1),
		NextToken:        aws.String("String"),
		SyncSessionToken: aws.String("SyncSessionToken"),
	}
	resp, err := svc.ListRecords(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_RegisterDevice() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.RegisterDeviceInput{
		IdentityId:     aws.String("IdentityId"),     // Required
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		Platform:       aws.String("Platform"),       // Required
		Token:          aws.String("PushToken"),      // Required
	}
	resp, err := svc.RegisterDevice(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_SetCognitoEvents() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.SetCognitoEventsInput{
		Events: map[string]*string{ // Required
			"Key": aws.String("LambdaFunctionArn"), // Required
			// More values...
		},
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.SetCognitoEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_SetIdentityPoolConfiguration() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.SetIdentityPoolConfigurationInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		CognitoStreams: &cognitosync.CognitoStreams{
			RoleArn:         aws.String("AssumeRoleArn"),
			StreamName:      aws.String("StreamName"),
			StreamingStatus: aws.String("StreamingStatus"),
		},
		PushSync: &cognitosync.PushSync{
			ApplicationArns: []*string{
				aws.String("ApplicationArn"), // Required
				// More values...
			},
			RoleArn: aws.String("AssumeRoleArn"),
		},
	}
	resp, err := svc.SetIdentityPoolConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_SubscribeToDataset() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.SubscribeToDatasetInput{
		DatasetName:    aws.String("DatasetName"),    // Required
		DeviceId:       aws.String("DeviceId"),       // Required
		IdentityId:     aws.String("IdentityId"),     // Required
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.SubscribeToDataset(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_UnsubscribeFromDataset() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.UnsubscribeFromDatasetInput{
		DatasetName:    aws.String("DatasetName"),    // Required
		DeviceId:       aws.String("DeviceId"),       // Required
		IdentityId:     aws.String("IdentityId"),     // Required
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.UnsubscribeFromDataset(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoSync_UpdateRecords() {
	sess := session.Must(session.NewSession())

	svc := cognitosync.New(sess)

	params := &cognitosync.UpdateRecordsInput{
		DatasetName:      aws.String("DatasetName"),      // Required
		IdentityId:       aws.String("IdentityId"),       // Required
		IdentityPoolId:   aws.String("IdentityPoolId"),   // Required
		SyncSessionToken: aws.String("SyncSessionToken"), // Required
		ClientContext:    aws.String("ClientContext"),
		DeviceId:         aws.String("DeviceId"),
		RecordPatches: []*cognitosync.RecordPatch{
			{ // Required
				Key:                    aws.String("RecordKey"), // Required
				Op:                     aws.String("Operation"), // Required
				SyncCount:              aws.Int64(1),            // Required
				DeviceLastModifiedDate: aws.Time(time.Now()),
				Value: aws.String("RecordValue"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateRecords(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
