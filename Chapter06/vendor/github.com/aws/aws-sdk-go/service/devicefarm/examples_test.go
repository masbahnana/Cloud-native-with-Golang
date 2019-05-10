// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDeviceFarm_CreateDevicePool() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.CreateDevicePoolInput{
		Name:       aws.String("Name"),               // Required
		ProjectArn: aws.String("AmazonResourceName"), // Required
		Rules: []*devicefarm.Rule{ // Required
			{ // Required
				Attribute: aws.String("DeviceAttribute"),
				Operator:  aws.String("RuleOperator"),
				Value:     aws.String("String"),
			},
			// More values...
		},
		Description: aws.String("Message"),
	}
	resp, err := svc.CreateDevicePool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_CreateNetworkProfile() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.CreateNetworkProfileInput{
		Name:                  aws.String("Name"),               // Required
		ProjectArn:            aws.String("AmazonResourceName"), // Required
		Description:           aws.String("Message"),
		DownlinkBandwidthBits: aws.Int64(1),
		DownlinkDelayMs:       aws.Int64(1),
		DownlinkJitterMs:      aws.Int64(1),
		DownlinkLossPercent:   aws.Int64(1),
		Type:                  aws.String("NetworkProfileType"),
		UplinkBandwidthBits:   aws.Int64(1),
		UplinkDelayMs:         aws.Int64(1),
		UplinkJitterMs:        aws.Int64(1),
		UplinkLossPercent:     aws.Int64(1),
	}
	resp, err := svc.CreateNetworkProfile(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_CreateProject() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.CreateProjectInput{
		Name: aws.String("Name"), // Required
		DefaultJobTimeoutMinutes: aws.Int64(1),
	}
	resp, err := svc.CreateProject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_CreateRemoteAccessSession() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.CreateRemoteAccessSessionInput{
		DeviceArn:  aws.String("AmazonResourceName"), // Required
		ProjectArn: aws.String("AmazonResourceName"), // Required
		Configuration: &devicefarm.CreateRemoteAccessSessionConfiguration{
			BillingMethod: aws.String("BillingMethod"),
		},
		Name: aws.String("Name"),
	}
	resp, err := svc.CreateRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_CreateUpload() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.CreateUploadInput{
		Name:        aws.String("Name"),               // Required
		ProjectArn:  aws.String("AmazonResourceName"), // Required
		Type:        aws.String("UploadType"),         // Required
		ContentType: aws.String("ContentType"),
	}
	resp, err := svc.CreateUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteDevicePool() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.DeleteDevicePoolInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteDevicePool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteNetworkProfile() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.DeleteNetworkProfileInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteNetworkProfile(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteProject() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.DeleteProjectInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteProject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteRemoteAccessSession() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.DeleteRemoteAccessSessionInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteRun() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.DeleteRunInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_DeleteUpload() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.DeleteUploadInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.DeleteUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetAccountSettings() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	var params *devicefarm.GetAccountSettingsInput
	resp, err := svc.GetAccountSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetDevice() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetDeviceInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetDevice(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetDevicePool() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetDevicePoolInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetDevicePool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetDevicePoolCompatibility() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetDevicePoolCompatibilityInput{
		DevicePoolArn: aws.String("AmazonResourceName"), // Required
		AppArn:        aws.String("AmazonResourceName"),
		Test: &devicefarm.ScheduleRunTest{
			Type:   aws.String("TestType"), // Required
			Filter: aws.String("Filter"),
			Parameters: map[string]*string{
				"Key": aws.String("String"), // Required
				// More values...
			},
			TestPackageArn: aws.String("AmazonResourceName"),
		},
		TestType: aws.String("TestType"),
	}
	resp, err := svc.GetDevicePoolCompatibility(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetJob() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetJobInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetNetworkProfile() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetNetworkProfileInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetNetworkProfile(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetOfferingStatus() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetOfferingStatusInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.GetOfferingStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetProject() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetProjectInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetProject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetRemoteAccessSession() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetRemoteAccessSessionInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetRun() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetRunInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetSuite() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetSuiteInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetSuite(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetTest() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetTestInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetTest(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_GetUpload() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.GetUploadInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetUpload(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_InstallToRemoteAccessSession() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.InstallToRemoteAccessSessionInput{
		AppArn:                 aws.String("AmazonResourceName"), // Required
		RemoteAccessSessionArn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.InstallToRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListArtifacts() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListArtifactsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		Type:      aws.String("ArtifactCategory"),   // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListArtifacts(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListDevicePools() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListDevicePoolsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
		Type:      aws.String("DevicePoolType"),
	}
	resp, err := svc.ListDevicePools(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListDevices() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListDevicesInput{
		Arn:       aws.String("AmazonResourceName"),
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListDevices(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListJobs() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListJobsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListNetworkProfiles() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListNetworkProfilesInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
		Type:      aws.String("NetworkProfileType"),
	}
	resp, err := svc.ListNetworkProfiles(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListOfferingPromotions() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListOfferingPromotionsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListOfferingPromotions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListOfferingTransactions() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListOfferingTransactionsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListOfferingTransactions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListOfferings() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListOfferingsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListOfferings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListProjects() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListProjectsInput{
		Arn:       aws.String("AmazonResourceName"),
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListProjects(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListRemoteAccessSessions() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListRemoteAccessSessionsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListRemoteAccessSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListRuns() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListRunsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListRuns(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListSamples() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListSamplesInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListSamples(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListSuites() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListSuitesInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListSuites(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListTests() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListTestsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListTests(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListUniqueProblems() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListUniqueProblemsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListUniqueProblems(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ListUploads() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ListUploadsInput{
		Arn:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListUploads(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_PurchaseOffering() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.PurchaseOfferingInput{
		OfferingId:          aws.String("OfferingIdentifier"),
		OfferingPromotionId: aws.String("OfferingPromotionIdentifier"),
		Quantity:            aws.Int64(1),
	}
	resp, err := svc.PurchaseOffering(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_RenewOffering() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.RenewOfferingInput{
		OfferingId: aws.String("OfferingIdentifier"),
		Quantity:   aws.Int64(1),
	}
	resp, err := svc.RenewOffering(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_ScheduleRun() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.ScheduleRunInput{
		DevicePoolArn: aws.String("AmazonResourceName"), // Required
		ProjectArn:    aws.String("AmazonResourceName"), // Required
		Test: &devicefarm.ScheduleRunTest{ // Required
			Type:   aws.String("TestType"), // Required
			Filter: aws.String("Filter"),
			Parameters: map[string]*string{
				"Key": aws.String("String"), // Required
				// More values...
			},
			TestPackageArn: aws.String("AmazonResourceName"),
		},
		AppArn: aws.String("AmazonResourceName"),
		Configuration: &devicefarm.ScheduleRunConfiguration{
			AuxiliaryApps: []*string{
				aws.String("AmazonResourceName"), // Required
				// More values...
			},
			BillingMethod:       aws.String("BillingMethod"),
			ExtraDataPackageArn: aws.String("AmazonResourceName"),
			Locale:              aws.String("String"),
			Location: &devicefarm.Location{
				Latitude:  aws.Float64(1.0), // Required
				Longitude: aws.Float64(1.0), // Required
			},
			NetworkProfileArn: aws.String("AmazonResourceName"),
			Radios: &devicefarm.Radios{
				Bluetooth: aws.Bool(true),
				Gps:       aws.Bool(true),
				Nfc:       aws.Bool(true),
				Wifi:      aws.Bool(true),
			},
		},
		ExecutionConfiguration: &devicefarm.ExecutionConfiguration{
			AccountsCleanup:    aws.Bool(true),
			AppPackagesCleanup: aws.Bool(true),
			JobTimeoutMinutes:  aws.Int64(1),
		},
		Name: aws.String("Name"),
	}
	resp, err := svc.ScheduleRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_StopRemoteAccessSession() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.StopRemoteAccessSessionInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.StopRemoteAccessSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_StopRun() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.StopRunInput{
		Arn: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.StopRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_UpdateDevicePool() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.UpdateDevicePoolInput{
		Arn:         aws.String("AmazonResourceName"), // Required
		Description: aws.String("Message"),
		Name:        aws.String("Name"),
		Rules: []*devicefarm.Rule{
			{ // Required
				Attribute: aws.String("DeviceAttribute"),
				Operator:  aws.String("RuleOperator"),
				Value:     aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateDevicePool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_UpdateNetworkProfile() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.UpdateNetworkProfileInput{
		Arn:                   aws.String("AmazonResourceName"), // Required
		Description:           aws.String("Message"),
		DownlinkBandwidthBits: aws.Int64(1),
		DownlinkDelayMs:       aws.Int64(1),
		DownlinkJitterMs:      aws.Int64(1),
		DownlinkLossPercent:   aws.Int64(1),
		Name:                  aws.String("Name"),
		Type:                  aws.String("NetworkProfileType"),
		UplinkBandwidthBits:   aws.Int64(1),
		UplinkDelayMs:         aws.Int64(1),
		UplinkJitterMs:        aws.Int64(1),
		UplinkLossPercent:     aws.Int64(1),
	}
	resp, err := svc.UpdateNetworkProfile(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDeviceFarm_UpdateProject() {
	sess := session.Must(session.NewSession())

	svc := devicefarm.New(sess)

	params := &devicefarm.UpdateProjectInput{
		Arn: aws.String("AmazonResourceName"), // Required
		DefaultJobTimeoutMinutes: aws.Int64(1),
		Name: aws.String("Name"),
	}
	resp, err := svc.UpdateProject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
