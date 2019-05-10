// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleKinesisAnalytics_AddApplicationInput() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.AddApplicationInputInput{
		ApplicationName:             aws.String("ApplicationName"), // Required
		CurrentApplicationVersionId: aws.Int64(1),                  // Required
		Input: &kinesisanalytics.Input{ // Required
			InputSchema: &kinesisanalytics.SourceSchema{ // Required
				RecordColumns: []*kinesisanalytics.RecordColumn{ // Required
					{ // Required
						Name:    aws.String("RecordColumnName"),    // Required
						SqlType: aws.String("RecordColumnSqlType"), // Required
						Mapping: aws.String("RecordColumnMapping"),
					},
					// More values...
				},
				RecordFormat: &kinesisanalytics.RecordFormat{ // Required
					RecordFormatType: aws.String("RecordFormatType"), // Required
					MappingParameters: &kinesisanalytics.MappingParameters{
						CSVMappingParameters: &kinesisanalytics.CSVMappingParameters{
							RecordColumnDelimiter: aws.String("RecordColumnDelimiter"), // Required
							RecordRowDelimiter:    aws.String("RecordRowDelimiter"),    // Required
						},
						JSONMappingParameters: &kinesisanalytics.JSONMappingParameters{
							RecordRowPath: aws.String("RecordRowPath"), // Required
						},
					},
				},
				RecordEncoding: aws.String("RecordEncoding"),
			},
			NamePrefix: aws.String("InAppStreamName"), // Required
			InputParallelism: &kinesisanalytics.InputParallelism{
				Count: aws.Int64(1),
			},
			KinesisFirehoseInput: &kinesisanalytics.KinesisFirehoseInput{
				ResourceARN: aws.String("ResourceARN"), // Required
				RoleARN:     aws.String("RoleARN"),     // Required
			},
			KinesisStreamsInput: &kinesisanalytics.KinesisStreamsInput{
				ResourceARN: aws.String("ResourceARN"), // Required
				RoleARN:     aws.String("RoleARN"),     // Required
			},
		},
	}
	resp, err := svc.AddApplicationInput(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_AddApplicationOutput() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.AddApplicationOutputInput{
		ApplicationName:             aws.String("ApplicationName"), // Required
		CurrentApplicationVersionId: aws.Int64(1),                  // Required
		Output: &kinesisanalytics.Output{ // Required
			DestinationSchema: &kinesisanalytics.DestinationSchema{ // Required
				RecordFormatType: aws.String("RecordFormatType"),
			},
			Name: aws.String("InAppStreamName"), // Required
			KinesisFirehoseOutput: &kinesisanalytics.KinesisFirehoseOutput{
				ResourceARN: aws.String("ResourceARN"), // Required
				RoleARN:     aws.String("RoleARN"),     // Required
			},
			KinesisStreamsOutput: &kinesisanalytics.KinesisStreamsOutput{
				ResourceARN: aws.String("ResourceARN"), // Required
				RoleARN:     aws.String("RoleARN"),     // Required
			},
		},
	}
	resp, err := svc.AddApplicationOutput(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_AddApplicationReferenceDataSource() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.AddApplicationReferenceDataSourceInput{
		ApplicationName:             aws.String("ApplicationName"), // Required
		CurrentApplicationVersionId: aws.Int64(1),                  // Required
		ReferenceDataSource: &kinesisanalytics.ReferenceDataSource{ // Required
			ReferenceSchema: &kinesisanalytics.SourceSchema{ // Required
				RecordColumns: []*kinesisanalytics.RecordColumn{ // Required
					{ // Required
						Name:    aws.String("RecordColumnName"),    // Required
						SqlType: aws.String("RecordColumnSqlType"), // Required
						Mapping: aws.String("RecordColumnMapping"),
					},
					// More values...
				},
				RecordFormat: &kinesisanalytics.RecordFormat{ // Required
					RecordFormatType: aws.String("RecordFormatType"), // Required
					MappingParameters: &kinesisanalytics.MappingParameters{
						CSVMappingParameters: &kinesisanalytics.CSVMappingParameters{
							RecordColumnDelimiter: aws.String("RecordColumnDelimiter"), // Required
							RecordRowDelimiter:    aws.String("RecordRowDelimiter"),    // Required
						},
						JSONMappingParameters: &kinesisanalytics.JSONMappingParameters{
							RecordRowPath: aws.String("RecordRowPath"), // Required
						},
					},
				},
				RecordEncoding: aws.String("RecordEncoding"),
			},
			TableName: aws.String("InAppTableName"), // Required
			S3ReferenceDataSource: &kinesisanalytics.S3ReferenceDataSource{
				BucketARN:        aws.String("BucketARN"), // Required
				FileKey:          aws.String("FileKey"),   // Required
				ReferenceRoleARN: aws.String("RoleARN"),   // Required
			},
		},
	}
	resp, err := svc.AddApplicationReferenceDataSource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_CreateApplication() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.CreateApplicationInput{
		ApplicationName:        aws.String("ApplicationName"), // Required
		ApplicationCode:        aws.String("ApplicationCode"),
		ApplicationDescription: aws.String("ApplicationDescription"),
		Inputs: []*kinesisanalytics.Input{
			{ // Required
				InputSchema: &kinesisanalytics.SourceSchema{ // Required
					RecordColumns: []*kinesisanalytics.RecordColumn{ // Required
						{ // Required
							Name:    aws.String("RecordColumnName"),    // Required
							SqlType: aws.String("RecordColumnSqlType"), // Required
							Mapping: aws.String("RecordColumnMapping"),
						},
						// More values...
					},
					RecordFormat: &kinesisanalytics.RecordFormat{ // Required
						RecordFormatType: aws.String("RecordFormatType"), // Required
						MappingParameters: &kinesisanalytics.MappingParameters{
							CSVMappingParameters: &kinesisanalytics.CSVMappingParameters{
								RecordColumnDelimiter: aws.String("RecordColumnDelimiter"), // Required
								RecordRowDelimiter:    aws.String("RecordRowDelimiter"),    // Required
							},
							JSONMappingParameters: &kinesisanalytics.JSONMappingParameters{
								RecordRowPath: aws.String("RecordRowPath"), // Required
							},
						},
					},
					RecordEncoding: aws.String("RecordEncoding"),
				},
				NamePrefix: aws.String("InAppStreamName"), // Required
				InputParallelism: &kinesisanalytics.InputParallelism{
					Count: aws.Int64(1),
				},
				KinesisFirehoseInput: &kinesisanalytics.KinesisFirehoseInput{
					ResourceARN: aws.String("ResourceARN"), // Required
					RoleARN:     aws.String("RoleARN"),     // Required
				},
				KinesisStreamsInput: &kinesisanalytics.KinesisStreamsInput{
					ResourceARN: aws.String("ResourceARN"), // Required
					RoleARN:     aws.String("RoleARN"),     // Required
				},
			},
			// More values...
		},
		Outputs: []*kinesisanalytics.Output{
			{ // Required
				DestinationSchema: &kinesisanalytics.DestinationSchema{ // Required
					RecordFormatType: aws.String("RecordFormatType"),
				},
				Name: aws.String("InAppStreamName"), // Required
				KinesisFirehoseOutput: &kinesisanalytics.KinesisFirehoseOutput{
					ResourceARN: aws.String("ResourceARN"), // Required
					RoleARN:     aws.String("RoleARN"),     // Required
				},
				KinesisStreamsOutput: &kinesisanalytics.KinesisStreamsOutput{
					ResourceARN: aws.String("ResourceARN"), // Required
					RoleARN:     aws.String("RoleARN"),     // Required
				},
			},
			// More values...
		},
	}
	resp, err := svc.CreateApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_DeleteApplication() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.DeleteApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		CreateTimestamp: aws.Time(time.Now()),          // Required
	}
	resp, err := svc.DeleteApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_DeleteApplicationOutput() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.DeleteApplicationOutputInput{
		ApplicationName:             aws.String("ApplicationName"), // Required
		CurrentApplicationVersionId: aws.Int64(1),                  // Required
		OutputId:                    aws.String("Id"),              // Required
	}
	resp, err := svc.DeleteApplicationOutput(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_DeleteApplicationReferenceDataSource() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.DeleteApplicationReferenceDataSourceInput{
		ApplicationName:             aws.String("ApplicationName"), // Required
		CurrentApplicationVersionId: aws.Int64(1),                  // Required
		ReferenceId:                 aws.String("Id"),              // Required
	}
	resp, err := svc.DeleteApplicationReferenceDataSource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_DescribeApplication() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.DescribeApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
	}
	resp, err := svc.DescribeApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_DiscoverInputSchema() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.DiscoverInputSchemaInput{
		InputStartingPositionConfiguration: &kinesisanalytics.InputStartingPositionConfiguration{ // Required
			InputStartingPosition: aws.String("InputStartingPosition"),
		},
		ResourceARN: aws.String("ResourceARN"), // Required
		RoleARN:     aws.String("RoleARN"),     // Required
	}
	resp, err := svc.DiscoverInputSchema(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_ListApplications() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.ListApplicationsInput{
		ExclusiveStartApplicationName: aws.String("ApplicationName"),
		Limit: aws.Int64(1),
	}
	resp, err := svc.ListApplications(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_StartApplication() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.StartApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		InputConfigurations: []*kinesisanalytics.InputConfiguration{ // Required
			{ // Required
				Id: aws.String("Id"), // Required
				InputStartingPositionConfiguration: &kinesisanalytics.InputStartingPositionConfiguration{ // Required
					InputStartingPosition: aws.String("InputStartingPosition"),
				},
			},
			// More values...
		},
	}
	resp, err := svc.StartApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_StopApplication() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.StopApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
	}
	resp, err := svc.StopApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKinesisAnalytics_UpdateApplication() {
	sess := session.Must(session.NewSession())

	svc := kinesisanalytics.New(sess)

	params := &kinesisanalytics.UpdateApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		ApplicationUpdate: &kinesisanalytics.ApplicationUpdate{ // Required
			ApplicationCodeUpdate: aws.String("ApplicationCode"),
			InputUpdates: []*kinesisanalytics.InputUpdate{
				{ // Required
					InputId: aws.String("Id"), // Required
					InputParallelismUpdate: &kinesisanalytics.InputParallelismUpdate{
						CountUpdate: aws.Int64(1),
					},
					InputSchemaUpdate: &kinesisanalytics.InputSchemaUpdate{
						RecordColumnUpdates: []*kinesisanalytics.RecordColumn{
							{ // Required
								Name:    aws.String("RecordColumnName"),    // Required
								SqlType: aws.String("RecordColumnSqlType"), // Required
								Mapping: aws.String("RecordColumnMapping"),
							},
							// More values...
						},
						RecordEncodingUpdate: aws.String("RecordEncoding"),
						RecordFormatUpdate: &kinesisanalytics.RecordFormat{
							RecordFormatType: aws.String("RecordFormatType"), // Required
							MappingParameters: &kinesisanalytics.MappingParameters{
								CSVMappingParameters: &kinesisanalytics.CSVMappingParameters{
									RecordColumnDelimiter: aws.String("RecordColumnDelimiter"), // Required
									RecordRowDelimiter:    aws.String("RecordRowDelimiter"),    // Required
								},
								JSONMappingParameters: &kinesisanalytics.JSONMappingParameters{
									RecordRowPath: aws.String("RecordRowPath"), // Required
								},
							},
						},
					},
					KinesisFirehoseInputUpdate: &kinesisanalytics.KinesisFirehoseInputUpdate{
						ResourceARNUpdate: aws.String("ResourceARN"),
						RoleARNUpdate:     aws.String("RoleARN"),
					},
					KinesisStreamsInputUpdate: &kinesisanalytics.KinesisStreamsInputUpdate{
						ResourceARNUpdate: aws.String("ResourceARN"),
						RoleARNUpdate:     aws.String("RoleARN"),
					},
					NamePrefixUpdate: aws.String("InAppStreamName"),
				},
				// More values...
			},
			OutputUpdates: []*kinesisanalytics.OutputUpdate{
				{ // Required
					OutputId: aws.String("Id"), // Required
					DestinationSchemaUpdate: &kinesisanalytics.DestinationSchema{
						RecordFormatType: aws.String("RecordFormatType"),
					},
					KinesisFirehoseOutputUpdate: &kinesisanalytics.KinesisFirehoseOutputUpdate{
						ResourceARNUpdate: aws.String("ResourceARN"),
						RoleARNUpdate:     aws.String("RoleARN"),
					},
					KinesisStreamsOutputUpdate: &kinesisanalytics.KinesisStreamsOutputUpdate{
						ResourceARNUpdate: aws.String("ResourceARN"),
						RoleARNUpdate:     aws.String("RoleARN"),
					},
					NameUpdate: aws.String("InAppStreamName"),
				},
				// More values...
			},
			ReferenceDataSourceUpdates: []*kinesisanalytics.ReferenceDataSourceUpdate{
				{ // Required
					ReferenceId: aws.String("Id"), // Required
					ReferenceSchemaUpdate: &kinesisanalytics.SourceSchema{
						RecordColumns: []*kinesisanalytics.RecordColumn{ // Required
							{ // Required
								Name:    aws.String("RecordColumnName"),    // Required
								SqlType: aws.String("RecordColumnSqlType"), // Required
								Mapping: aws.String("RecordColumnMapping"),
							},
							// More values...
						},
						RecordFormat: &kinesisanalytics.RecordFormat{ // Required
							RecordFormatType: aws.String("RecordFormatType"), // Required
							MappingParameters: &kinesisanalytics.MappingParameters{
								CSVMappingParameters: &kinesisanalytics.CSVMappingParameters{
									RecordColumnDelimiter: aws.String("RecordColumnDelimiter"), // Required
									RecordRowDelimiter:    aws.String("RecordRowDelimiter"),    // Required
								},
								JSONMappingParameters: &kinesisanalytics.JSONMappingParameters{
									RecordRowPath: aws.String("RecordRowPath"), // Required
								},
							},
						},
						RecordEncoding: aws.String("RecordEncoding"),
					},
					S3ReferenceDataSourceUpdate: &kinesisanalytics.S3ReferenceDataSourceUpdate{
						BucketARNUpdate:        aws.String("BucketARN"),
						FileKeyUpdate:          aws.String("FileKey"),
						ReferenceRoleARNUpdate: aws.String("RoleARN"),
					},
					TableNameUpdate: aws.String("InAppTableName"),
				},
				// More values...
			},
		},
		CurrentApplicationVersionId: aws.Int64(1), // Required
	}
	resp, err := svc.UpdateApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
