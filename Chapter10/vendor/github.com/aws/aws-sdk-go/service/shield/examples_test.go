// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/shield"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleShield_CreateProtection() {
	sess := session.Must(session.NewSession())

	svc := shield.New(sess)

	params := &shield.CreateProtectionInput{
		Name:        aws.String("ProtectionName"), // Required
		ResourceArn: aws.String("ResourceArn"),    // Required
	}
	resp, err := svc.CreateProtection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleShield_CreateSubscription() {
	sess := session.Must(session.NewSession())

	svc := shield.New(sess)

	var params *shield.CreateSubscriptionInput
	resp, err := svc.CreateSubscription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleShield_DeleteProtection() {
	sess := session.Must(session.NewSession())

	svc := shield.New(sess)

	params := &shield.DeleteProtectionInput{
		ProtectionId: aws.String("ProtectionId"), // Required
	}
	resp, err := svc.DeleteProtection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleShield_DeleteSubscription() {
	sess := session.Must(session.NewSession())

	svc := shield.New(sess)

	var params *shield.DeleteSubscriptionInput
	resp, err := svc.DeleteSubscription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleShield_DescribeAttack() {
	sess := session.Must(session.NewSession())

	svc := shield.New(sess)

	params := &shield.DescribeAttackInput{
		AttackId: aws.String("AttackId"), // Required
	}
	resp, err := svc.DescribeAttack(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleShield_DescribeProtection() {
	sess := session.Must(session.NewSession())

	svc := shield.New(sess)

	params := &shield.DescribeProtectionInput{
		ProtectionId: aws.String("ProtectionId"), // Required
	}
	resp, err := svc.DescribeProtection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleShield_DescribeSubscription() {
	sess := session.Must(session.NewSession())

	svc := shield.New(sess)

	var params *shield.DescribeSubscriptionInput
	resp, err := svc.DescribeSubscription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleShield_ListAttacks() {
	sess := session.Must(session.NewSession())

	svc := shield.New(sess)

	params := &shield.ListAttacksInput{
		EndTime: &shield.TimeRange{
			FromInclusive: aws.Time(time.Now()),
			ToExclusive:   aws.Time(time.Now()),
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("Token"),
		ResourceArns: []*string{
			aws.String("ResourceArn"), // Required
			// More values...
		},
		StartTime: &shield.TimeRange{
			FromInclusive: aws.Time(time.Now()),
			ToExclusive:   aws.Time(time.Now()),
		},
	}
	resp, err := svc.ListAttacks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleShield_ListProtections() {
	sess := session.Must(session.NewSession())

	svc := shield.New(sess)

	params := &shield.ListProtectionsInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("Token"),
	}
	resp, err := svc.ListProtections(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
