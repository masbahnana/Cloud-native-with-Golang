// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationautoscaling_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleApplicationAutoScaling_DeleteScalingPolicy() {
	sess := session.Must(session.NewSession())

	svc := applicationautoscaling.New(sess)

	params := &applicationautoscaling.DeleteScalingPolicyInput{
		PolicyName:        aws.String("ResourceIdMaxLen1600"), // Required
		ResourceId:        aws.String("ResourceIdMaxLen1600"), // Required
		ScalableDimension: aws.String("ScalableDimension"),    // Required
		ServiceNamespace:  aws.String("ServiceNamespace"),     // Required
	}
	resp, err := svc.DeleteScalingPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationAutoScaling_DeregisterScalableTarget() {
	sess := session.Must(session.NewSession())

	svc := applicationautoscaling.New(sess)

	params := &applicationautoscaling.DeregisterScalableTargetInput{
		ResourceId:        aws.String("ResourceIdMaxLen1600"), // Required
		ScalableDimension: aws.String("ScalableDimension"),    // Required
		ServiceNamespace:  aws.String("ServiceNamespace"),     // Required
	}
	resp, err := svc.DeregisterScalableTarget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationAutoScaling_DescribeScalableTargets() {
	sess := session.Must(session.NewSession())

	svc := applicationautoscaling.New(sess)

	params := &applicationautoscaling.DescribeScalableTargetsInput{
		ServiceNamespace: aws.String("ServiceNamespace"), // Required
		MaxResults:       aws.Int64(1),
		NextToken:        aws.String("XmlString"),
		ResourceIds: []*string{
			aws.String("ResourceIdMaxLen1600"), // Required
			// More values...
		},
		ScalableDimension: aws.String("ScalableDimension"),
	}
	resp, err := svc.DescribeScalableTargets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationAutoScaling_DescribeScalingActivities() {
	sess := session.Must(session.NewSession())

	svc := applicationautoscaling.New(sess)

	params := &applicationautoscaling.DescribeScalingActivitiesInput{
		ServiceNamespace:  aws.String("ServiceNamespace"), // Required
		MaxResults:        aws.Int64(1),
		NextToken:         aws.String("XmlString"),
		ResourceId:        aws.String("ResourceIdMaxLen1600"),
		ScalableDimension: aws.String("ScalableDimension"),
	}
	resp, err := svc.DescribeScalingActivities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationAutoScaling_DescribeScalingPolicies() {
	sess := session.Must(session.NewSession())

	svc := applicationautoscaling.New(sess)

	params := &applicationautoscaling.DescribeScalingPoliciesInput{
		ServiceNamespace: aws.String("ServiceNamespace"), // Required
		MaxResults:       aws.Int64(1),
		NextToken:        aws.String("XmlString"),
		PolicyNames: []*string{
			aws.String("ResourceIdMaxLen1600"), // Required
			// More values...
		},
		ResourceId:        aws.String("ResourceIdMaxLen1600"),
		ScalableDimension: aws.String("ScalableDimension"),
	}
	resp, err := svc.DescribeScalingPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationAutoScaling_PutScalingPolicy() {
	sess := session.Must(session.NewSession())

	svc := applicationautoscaling.New(sess)

	params := &applicationautoscaling.PutScalingPolicyInput{
		PolicyName:        aws.String("PolicyName"),           // Required
		ResourceId:        aws.String("ResourceIdMaxLen1600"), // Required
		ScalableDimension: aws.String("ScalableDimension"),    // Required
		ServiceNamespace:  aws.String("ServiceNamespace"),     // Required
		PolicyType:        aws.String("PolicyType"),
		StepScalingPolicyConfiguration: &applicationautoscaling.StepScalingPolicyConfiguration{
			AdjustmentType:         aws.String("AdjustmentType"),
			Cooldown:               aws.Int64(1),
			MetricAggregationType:  aws.String("MetricAggregationType"),
			MinAdjustmentMagnitude: aws.Int64(1),
			StepAdjustments: []*applicationautoscaling.StepAdjustment{
				{ // Required
					ScalingAdjustment:        aws.Int64(1), // Required
					MetricIntervalLowerBound: aws.Float64(1.0),
					MetricIntervalUpperBound: aws.Float64(1.0),
				},
				// More values...
			},
		},
	}
	resp, err := svc.PutScalingPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleApplicationAutoScaling_RegisterScalableTarget() {
	sess := session.Must(session.NewSession())

	svc := applicationautoscaling.New(sess)

	params := &applicationautoscaling.RegisterScalableTargetInput{
		ResourceId:        aws.String("ResourceIdMaxLen1600"), // Required
		ScalableDimension: aws.String("ScalableDimension"),    // Required
		ServiceNamespace:  aws.String("ServiceNamespace"),     // Required
		MaxCapacity:       aws.Int64(1),
		MinCapacity:       aws.Int64(1),
		RoleARN:           aws.String("ResourceIdMaxLen1600"),
	}
	resp, err := svc.RegisterScalableTarget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
