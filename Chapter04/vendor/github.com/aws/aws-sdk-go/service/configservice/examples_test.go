// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleConfigService_DeleteConfigRule() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DeleteConfigRuleInput{
		ConfigRuleName: aws.String("StringWithCharLimit64"), // Required
	}
	resp, err := svc.DeleteConfigRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DeleteConfigurationRecorder() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DeleteConfigurationRecorderInput{
		ConfigurationRecorderName: aws.String("RecorderName"), // Required
	}
	resp, err := svc.DeleteConfigurationRecorder(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DeleteDeliveryChannel() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DeleteDeliveryChannelInput{
		DeliveryChannelName: aws.String("ChannelName"), // Required
	}
	resp, err := svc.DeleteDeliveryChannel(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DeleteEvaluationResults() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DeleteEvaluationResultsInput{
		ConfigRuleName: aws.String("StringWithCharLimit64"), // Required
	}
	resp, err := svc.DeleteEvaluationResults(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DeliverConfigSnapshot() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DeliverConfigSnapshotInput{
		DeliveryChannelName: aws.String("ChannelName"), // Required
	}
	resp, err := svc.DeliverConfigSnapshot(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DescribeComplianceByConfigRule() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DescribeComplianceByConfigRuleInput{
		ComplianceTypes: []*string{
			aws.String("ComplianceType"), // Required
			// More values...
		},
		ConfigRuleNames: []*string{
			aws.String("StringWithCharLimit64"), // Required
			// More values...
		},
		NextToken: aws.String("String"),
	}
	resp, err := svc.DescribeComplianceByConfigRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DescribeComplianceByResource() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DescribeComplianceByResourceInput{
		ComplianceTypes: []*string{
			aws.String("ComplianceType"), // Required
			// More values...
		},
		Limit:        aws.Int64(1),
		NextToken:    aws.String("NextToken"),
		ResourceId:   aws.String("StringWithCharLimit256"),
		ResourceType: aws.String("StringWithCharLimit256"),
	}
	resp, err := svc.DescribeComplianceByResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DescribeConfigRuleEvaluationStatus() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DescribeConfigRuleEvaluationStatusInput{
		ConfigRuleNames: []*string{
			aws.String("StringWithCharLimit64"), // Required
			// More values...
		},
		Limit:     aws.Int64(1),
		NextToken: aws.String("String"),
	}
	resp, err := svc.DescribeConfigRuleEvaluationStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DescribeConfigRules() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DescribeConfigRulesInput{
		ConfigRuleNames: []*string{
			aws.String("StringWithCharLimit64"), // Required
			// More values...
		},
		NextToken: aws.String("String"),
	}
	resp, err := svc.DescribeConfigRules(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DescribeConfigurationRecorderStatus() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DescribeConfigurationRecorderStatusInput{
		ConfigurationRecorderNames: []*string{
			aws.String("RecorderName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeConfigurationRecorderStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DescribeConfigurationRecorders() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DescribeConfigurationRecordersInput{
		ConfigurationRecorderNames: []*string{
			aws.String("RecorderName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeConfigurationRecorders(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DescribeDeliveryChannelStatus() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DescribeDeliveryChannelStatusInput{
		DeliveryChannelNames: []*string{
			aws.String("ChannelName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeDeliveryChannelStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_DescribeDeliveryChannels() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.DescribeDeliveryChannelsInput{
		DeliveryChannelNames: []*string{
			aws.String("ChannelName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeDeliveryChannels(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_GetComplianceDetailsByConfigRule() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.GetComplianceDetailsByConfigRuleInput{
		ConfigRuleName: aws.String("StringWithCharLimit64"), // Required
		ComplianceTypes: []*string{
			aws.String("ComplianceType"), // Required
			// More values...
		},
		Limit:     aws.Int64(1),
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.GetComplianceDetailsByConfigRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_GetComplianceDetailsByResource() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.GetComplianceDetailsByResourceInput{
		ResourceId:   aws.String("StringWithCharLimit256"), // Required
		ResourceType: aws.String("StringWithCharLimit256"), // Required
		ComplianceTypes: []*string{
			aws.String("ComplianceType"), // Required
			// More values...
		},
		NextToken: aws.String("String"),
	}
	resp, err := svc.GetComplianceDetailsByResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_GetComplianceSummaryByConfigRule() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	var params *configservice.GetComplianceSummaryByConfigRuleInput
	resp, err := svc.GetComplianceSummaryByConfigRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_GetComplianceSummaryByResourceType() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.GetComplianceSummaryByResourceTypeInput{
		ResourceTypes: []*string{
			aws.String("StringWithCharLimit256"), // Required
			// More values...
		},
	}
	resp, err := svc.GetComplianceSummaryByResourceType(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_GetResourceConfigHistory() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.GetResourceConfigHistoryInput{
		ResourceId:         aws.String("ResourceId"),   // Required
		ResourceType:       aws.String("ResourceType"), // Required
		ChronologicalOrder: aws.String("ChronologicalOrder"),
		EarlierTime:        aws.Time(time.Now()),
		LaterTime:          aws.Time(time.Now()),
		Limit:              aws.Int64(1),
		NextToken:          aws.String("NextToken"),
	}
	resp, err := svc.GetResourceConfigHistory(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_ListDiscoveredResources() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.ListDiscoveredResourcesInput{
		ResourceType:            aws.String("ResourceType"), // Required
		IncludeDeletedResources: aws.Bool(true),
		Limit:     aws.Int64(1),
		NextToken: aws.String("NextToken"),
		ResourceIds: []*string{
			aws.String("ResourceId"), // Required
			// More values...
		},
		ResourceName: aws.String("ResourceName"),
	}
	resp, err := svc.ListDiscoveredResources(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_PutConfigRule() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.PutConfigRuleInput{
		ConfigRule: &configservice.ConfigRule{ // Required
			Source: &configservice.Source{ // Required
				Owner:            aws.String("Owner"),                  // Required
				SourceIdentifier: aws.String("StringWithCharLimit256"), // Required
				SourceDetails: []*configservice.SourceDetail{
					{ // Required
						EventSource:               aws.String("EventSource"),
						MaximumExecutionFrequency: aws.String("MaximumExecutionFrequency"),
						MessageType:               aws.String("MessageType"),
					},
					// More values...
				},
			},
			ConfigRuleArn:             aws.String("String"),
			ConfigRuleId:              aws.String("String"),
			ConfigRuleName:            aws.String("StringWithCharLimit64"),
			ConfigRuleState:           aws.String("ConfigRuleState"),
			Description:               aws.String("EmptiableStringWithCharLimit256"),
			InputParameters:           aws.String("StringWithCharLimit1024"),
			MaximumExecutionFrequency: aws.String("MaximumExecutionFrequency"),
			Scope: &configservice.Scope{
				ComplianceResourceId: aws.String("StringWithCharLimit256"),
				ComplianceResourceTypes: []*string{
					aws.String("StringWithCharLimit256"), // Required
					// More values...
				},
				TagKey:   aws.String("StringWithCharLimit128"),
				TagValue: aws.String("StringWithCharLimit256"),
			},
		},
	}
	resp, err := svc.PutConfigRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_PutConfigurationRecorder() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.PutConfigurationRecorderInput{
		ConfigurationRecorder: &configservice.ConfigurationRecorder{ // Required
			Name: aws.String("RecorderName"),
			RecordingGroup: &configservice.RecordingGroup{
				AllSupported:               aws.Bool(true),
				IncludeGlobalResourceTypes: aws.Bool(true),
				ResourceTypes: []*string{
					aws.String("ResourceType"), // Required
					// More values...
				},
			},
			RoleARN: aws.String("String"),
		},
	}
	resp, err := svc.PutConfigurationRecorder(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_PutDeliveryChannel() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.PutDeliveryChannelInput{
		DeliveryChannel: &configservice.DeliveryChannel{ // Required
			ConfigSnapshotDeliveryProperties: &configservice.ConfigSnapshotDeliveryProperties{
				DeliveryFrequency: aws.String("MaximumExecutionFrequency"),
			},
			Name:         aws.String("ChannelName"),
			S3BucketName: aws.String("String"),
			S3KeyPrefix:  aws.String("String"),
			SnsTopicARN:  aws.String("String"),
		},
	}
	resp, err := svc.PutDeliveryChannel(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_PutEvaluations() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.PutEvaluationsInput{
		ResultToken: aws.String("String"), // Required
		Evaluations: []*configservice.Evaluation{
			{ // Required
				ComplianceResourceId:   aws.String("StringWithCharLimit256"), // Required
				ComplianceResourceType: aws.String("StringWithCharLimit256"), // Required
				ComplianceType:         aws.String("ComplianceType"),         // Required
				OrderingTimestamp:      aws.Time(time.Now()),                 // Required
				Annotation:             aws.String("StringWithCharLimit256"),
			},
			// More values...
		},
		TestMode: aws.Bool(true),
	}
	resp, err := svc.PutEvaluations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_StartConfigRulesEvaluation() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.StartConfigRulesEvaluationInput{
		ConfigRuleNames: []*string{
			aws.String("StringWithCharLimit64"), // Required
			// More values...
		},
	}
	resp, err := svc.StartConfigRulesEvaluation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_StartConfigurationRecorder() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.StartConfigurationRecorderInput{
		ConfigurationRecorderName: aws.String("RecorderName"), // Required
	}
	resp, err := svc.StartConfigurationRecorder(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleConfigService_StopConfigurationRecorder() {
	sess := session.Must(session.NewSession())

	svc := configservice.New(sess)

	params := &configservice.StopConfigurationRecorderInput{
		ConfigurationRecorderName: aws.String("RecorderName"), // Required
	}
	resp, err := svc.StopConfigurationRecorder(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
