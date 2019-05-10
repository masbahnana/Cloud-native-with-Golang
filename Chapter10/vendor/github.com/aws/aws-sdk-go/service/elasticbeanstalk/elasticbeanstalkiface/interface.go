// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elasticbeanstalkiface provides an interface to enable mocking the AWS Elastic Beanstalk service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elasticbeanstalkiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

// ElasticBeanstalkAPI provides an interface to enable mocking the
// elasticbeanstalk.ElasticBeanstalk service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Elastic Beanstalk.
//    func myFunc(svc elasticbeanstalkiface.ElasticBeanstalkAPI) bool {
//        // Make svc.AbortEnvironmentUpdate request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := elasticbeanstalk.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockElasticBeanstalkClient struct {
//        elasticbeanstalkiface.ElasticBeanstalkAPI
//    }
//    func (m *mockElasticBeanstalkClient) AbortEnvironmentUpdate(input *elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockElasticBeanstalkClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ElasticBeanstalkAPI interface {
	AbortEnvironmentUpdate(*elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error)
	AbortEnvironmentUpdateWithContext(aws.Context, *elasticbeanstalk.AbortEnvironmentUpdateInput, ...request.Option) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error)
	AbortEnvironmentUpdateRequest(*elasticbeanstalk.AbortEnvironmentUpdateInput) (*request.Request, *elasticbeanstalk.AbortEnvironmentUpdateOutput)

	ApplyEnvironmentManagedAction(*elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error)
	ApplyEnvironmentManagedActionWithContext(aws.Context, *elasticbeanstalk.ApplyEnvironmentManagedActionInput, ...request.Option) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error)
	ApplyEnvironmentManagedActionRequest(*elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*request.Request, *elasticbeanstalk.ApplyEnvironmentManagedActionOutput)

	CheckDNSAvailability(*elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error)
	CheckDNSAvailabilityWithContext(aws.Context, *elasticbeanstalk.CheckDNSAvailabilityInput, ...request.Option) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error)
	CheckDNSAvailabilityRequest(*elasticbeanstalk.CheckDNSAvailabilityInput) (*request.Request, *elasticbeanstalk.CheckDNSAvailabilityOutput)

	ComposeEnvironments(*elasticbeanstalk.ComposeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)
	ComposeEnvironmentsWithContext(aws.Context, *elasticbeanstalk.ComposeEnvironmentsInput, ...request.Option) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)
	ComposeEnvironmentsRequest(*elasticbeanstalk.ComposeEnvironmentsInput) (*request.Request, *elasticbeanstalk.EnvironmentDescriptionsMessage)

	CreateApplication(*elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)
	CreateApplicationWithContext(aws.Context, *elasticbeanstalk.CreateApplicationInput, ...request.Option) (*elasticbeanstalk.ApplicationDescriptionMessage, error)
	CreateApplicationRequest(*elasticbeanstalk.CreateApplicationInput) (*request.Request, *elasticbeanstalk.ApplicationDescriptionMessage)

	CreateApplicationVersion(*elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)
	CreateApplicationVersionWithContext(aws.Context, *elasticbeanstalk.CreateApplicationVersionInput, ...request.Option) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)
	CreateApplicationVersionRequest(*elasticbeanstalk.CreateApplicationVersionInput) (*request.Request, *elasticbeanstalk.ApplicationVersionDescriptionMessage)

	CreateConfigurationTemplate(*elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)
	CreateConfigurationTemplateWithContext(aws.Context, *elasticbeanstalk.CreateConfigurationTemplateInput, ...request.Option) (*elasticbeanstalk.ConfigurationSettingsDescription, error)
	CreateConfigurationTemplateRequest(*elasticbeanstalk.CreateConfigurationTemplateInput) (*request.Request, *elasticbeanstalk.ConfigurationSettingsDescription)

	CreateEnvironment(*elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)
	CreateEnvironmentWithContext(aws.Context, *elasticbeanstalk.CreateEnvironmentInput, ...request.Option) (*elasticbeanstalk.EnvironmentDescription, error)
	CreateEnvironmentRequest(*elasticbeanstalk.CreateEnvironmentInput) (*request.Request, *elasticbeanstalk.EnvironmentDescription)

	CreatePlatformVersion(*elasticbeanstalk.CreatePlatformVersionInput) (*elasticbeanstalk.CreatePlatformVersionOutput, error)
	CreatePlatformVersionWithContext(aws.Context, *elasticbeanstalk.CreatePlatformVersionInput, ...request.Option) (*elasticbeanstalk.CreatePlatformVersionOutput, error)
	CreatePlatformVersionRequest(*elasticbeanstalk.CreatePlatformVersionInput) (*request.Request, *elasticbeanstalk.CreatePlatformVersionOutput)

	CreateStorageLocation(*elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error)
	CreateStorageLocationWithContext(aws.Context, *elasticbeanstalk.CreateStorageLocationInput, ...request.Option) (*elasticbeanstalk.CreateStorageLocationOutput, error)
	CreateStorageLocationRequest(*elasticbeanstalk.CreateStorageLocationInput) (*request.Request, *elasticbeanstalk.CreateStorageLocationOutput)

	DeleteApplication(*elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *elasticbeanstalk.DeleteApplicationInput, ...request.Option) (*elasticbeanstalk.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*elasticbeanstalk.DeleteApplicationInput) (*request.Request, *elasticbeanstalk.DeleteApplicationOutput)

	DeleteApplicationVersion(*elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error)
	DeleteApplicationVersionWithContext(aws.Context, *elasticbeanstalk.DeleteApplicationVersionInput, ...request.Option) (*elasticbeanstalk.DeleteApplicationVersionOutput, error)
	DeleteApplicationVersionRequest(*elasticbeanstalk.DeleteApplicationVersionInput) (*request.Request, *elasticbeanstalk.DeleteApplicationVersionOutput)

	DeleteConfigurationTemplate(*elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error)
	DeleteConfigurationTemplateWithContext(aws.Context, *elasticbeanstalk.DeleteConfigurationTemplateInput, ...request.Option) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error)
	DeleteConfigurationTemplateRequest(*elasticbeanstalk.DeleteConfigurationTemplateInput) (*request.Request, *elasticbeanstalk.DeleteConfigurationTemplateOutput)

	DeleteEnvironmentConfiguration(*elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error)
	DeleteEnvironmentConfigurationWithContext(aws.Context, *elasticbeanstalk.DeleteEnvironmentConfigurationInput, ...request.Option) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error)
	DeleteEnvironmentConfigurationRequest(*elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*request.Request, *elasticbeanstalk.DeleteEnvironmentConfigurationOutput)

	DeletePlatformVersion(*elasticbeanstalk.DeletePlatformVersionInput) (*elasticbeanstalk.DeletePlatformVersionOutput, error)
	DeletePlatformVersionWithContext(aws.Context, *elasticbeanstalk.DeletePlatformVersionInput, ...request.Option) (*elasticbeanstalk.DeletePlatformVersionOutput, error)
	DeletePlatformVersionRequest(*elasticbeanstalk.DeletePlatformVersionInput) (*request.Request, *elasticbeanstalk.DeletePlatformVersionOutput)

	DescribeApplicationVersions(*elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error)
	DescribeApplicationVersionsWithContext(aws.Context, *elasticbeanstalk.DescribeApplicationVersionsInput, ...request.Option) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error)
	DescribeApplicationVersionsRequest(*elasticbeanstalk.DescribeApplicationVersionsInput) (*request.Request, *elasticbeanstalk.DescribeApplicationVersionsOutput)

	DescribeApplications(*elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error)
	DescribeApplicationsWithContext(aws.Context, *elasticbeanstalk.DescribeApplicationsInput, ...request.Option) (*elasticbeanstalk.DescribeApplicationsOutput, error)
	DescribeApplicationsRequest(*elasticbeanstalk.DescribeApplicationsInput) (*request.Request, *elasticbeanstalk.DescribeApplicationsOutput)

	DescribeConfigurationOptions(*elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error)
	DescribeConfigurationOptionsWithContext(aws.Context, *elasticbeanstalk.DescribeConfigurationOptionsInput, ...request.Option) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error)
	DescribeConfigurationOptionsRequest(*elasticbeanstalk.DescribeConfigurationOptionsInput) (*request.Request, *elasticbeanstalk.DescribeConfigurationOptionsOutput)

	DescribeConfigurationSettings(*elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error)
	DescribeConfigurationSettingsWithContext(aws.Context, *elasticbeanstalk.DescribeConfigurationSettingsInput, ...request.Option) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error)
	DescribeConfigurationSettingsRequest(*elasticbeanstalk.DescribeConfigurationSettingsInput) (*request.Request, *elasticbeanstalk.DescribeConfigurationSettingsOutput)

	DescribeEnvironmentHealth(*elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error)
	DescribeEnvironmentHealthWithContext(aws.Context, *elasticbeanstalk.DescribeEnvironmentHealthInput, ...request.Option) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error)
	DescribeEnvironmentHealthRequest(*elasticbeanstalk.DescribeEnvironmentHealthInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentHealthOutput)

	DescribeEnvironmentManagedActionHistory(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error)
	DescribeEnvironmentManagedActionHistoryWithContext(aws.Context, *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput, ...request.Option) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error)
	DescribeEnvironmentManagedActionHistoryRequest(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput)

	DescribeEnvironmentManagedActions(*elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error)
	DescribeEnvironmentManagedActionsWithContext(aws.Context, *elasticbeanstalk.DescribeEnvironmentManagedActionsInput, ...request.Option) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error)
	DescribeEnvironmentManagedActionsRequest(*elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentManagedActionsOutput)

	DescribeEnvironmentResources(*elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error)
	DescribeEnvironmentResourcesWithContext(aws.Context, *elasticbeanstalk.DescribeEnvironmentResourcesInput, ...request.Option) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error)
	DescribeEnvironmentResourcesRequest(*elasticbeanstalk.DescribeEnvironmentResourcesInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentResourcesOutput)

	DescribeEnvironments(*elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)
	DescribeEnvironmentsWithContext(aws.Context, *elasticbeanstalk.DescribeEnvironmentsInput, ...request.Option) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)
	DescribeEnvironmentsRequest(*elasticbeanstalk.DescribeEnvironmentsInput) (*request.Request, *elasticbeanstalk.EnvironmentDescriptionsMessage)

	DescribeEvents(*elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error)
	DescribeEventsWithContext(aws.Context, *elasticbeanstalk.DescribeEventsInput, ...request.Option) (*elasticbeanstalk.DescribeEventsOutput, error)
	DescribeEventsRequest(*elasticbeanstalk.DescribeEventsInput) (*request.Request, *elasticbeanstalk.DescribeEventsOutput)

	DescribeEventsPages(*elasticbeanstalk.DescribeEventsInput, func(*elasticbeanstalk.DescribeEventsOutput, bool) bool) error
	DescribeEventsPagesWithContext(aws.Context, *elasticbeanstalk.DescribeEventsInput, func(*elasticbeanstalk.DescribeEventsOutput, bool) bool, ...request.Option) error

	DescribeInstancesHealth(*elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error)
	DescribeInstancesHealthWithContext(aws.Context, *elasticbeanstalk.DescribeInstancesHealthInput, ...request.Option) (*elasticbeanstalk.DescribeInstancesHealthOutput, error)
	DescribeInstancesHealthRequest(*elasticbeanstalk.DescribeInstancesHealthInput) (*request.Request, *elasticbeanstalk.DescribeInstancesHealthOutput)

	DescribePlatformVersion(*elasticbeanstalk.DescribePlatformVersionInput) (*elasticbeanstalk.DescribePlatformVersionOutput, error)
	DescribePlatformVersionWithContext(aws.Context, *elasticbeanstalk.DescribePlatformVersionInput, ...request.Option) (*elasticbeanstalk.DescribePlatformVersionOutput, error)
	DescribePlatformVersionRequest(*elasticbeanstalk.DescribePlatformVersionInput) (*request.Request, *elasticbeanstalk.DescribePlatformVersionOutput)

	ListAvailableSolutionStacks(*elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error)
	ListAvailableSolutionStacksWithContext(aws.Context, *elasticbeanstalk.ListAvailableSolutionStacksInput, ...request.Option) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error)
	ListAvailableSolutionStacksRequest(*elasticbeanstalk.ListAvailableSolutionStacksInput) (*request.Request, *elasticbeanstalk.ListAvailableSolutionStacksOutput)

	ListPlatformVersions(*elasticbeanstalk.ListPlatformVersionsInput) (*elasticbeanstalk.ListPlatformVersionsOutput, error)
	ListPlatformVersionsWithContext(aws.Context, *elasticbeanstalk.ListPlatformVersionsInput, ...request.Option) (*elasticbeanstalk.ListPlatformVersionsOutput, error)
	ListPlatformVersionsRequest(*elasticbeanstalk.ListPlatformVersionsInput) (*request.Request, *elasticbeanstalk.ListPlatformVersionsOutput)

	RebuildEnvironment(*elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error)
	RebuildEnvironmentWithContext(aws.Context, *elasticbeanstalk.RebuildEnvironmentInput, ...request.Option) (*elasticbeanstalk.RebuildEnvironmentOutput, error)
	RebuildEnvironmentRequest(*elasticbeanstalk.RebuildEnvironmentInput) (*request.Request, *elasticbeanstalk.RebuildEnvironmentOutput)

	RequestEnvironmentInfo(*elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error)
	RequestEnvironmentInfoWithContext(aws.Context, *elasticbeanstalk.RequestEnvironmentInfoInput, ...request.Option) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error)
	RequestEnvironmentInfoRequest(*elasticbeanstalk.RequestEnvironmentInfoInput) (*request.Request, *elasticbeanstalk.RequestEnvironmentInfoOutput)

	RestartAppServer(*elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error)
	RestartAppServerWithContext(aws.Context, *elasticbeanstalk.RestartAppServerInput, ...request.Option) (*elasticbeanstalk.RestartAppServerOutput, error)
	RestartAppServerRequest(*elasticbeanstalk.RestartAppServerInput) (*request.Request, *elasticbeanstalk.RestartAppServerOutput)

	RetrieveEnvironmentInfo(*elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error)
	RetrieveEnvironmentInfoWithContext(aws.Context, *elasticbeanstalk.RetrieveEnvironmentInfoInput, ...request.Option) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error)
	RetrieveEnvironmentInfoRequest(*elasticbeanstalk.RetrieveEnvironmentInfoInput) (*request.Request, *elasticbeanstalk.RetrieveEnvironmentInfoOutput)

	SwapEnvironmentCNAMEs(*elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error)
	SwapEnvironmentCNAMEsWithContext(aws.Context, *elasticbeanstalk.SwapEnvironmentCNAMEsInput, ...request.Option) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error)
	SwapEnvironmentCNAMEsRequest(*elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*request.Request, *elasticbeanstalk.SwapEnvironmentCNAMEsOutput)

	TerminateEnvironment(*elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)
	TerminateEnvironmentWithContext(aws.Context, *elasticbeanstalk.TerminateEnvironmentInput, ...request.Option) (*elasticbeanstalk.EnvironmentDescription, error)
	TerminateEnvironmentRequest(*elasticbeanstalk.TerminateEnvironmentInput) (*request.Request, *elasticbeanstalk.EnvironmentDescription)

	UpdateApplication(*elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)
	UpdateApplicationWithContext(aws.Context, *elasticbeanstalk.UpdateApplicationInput, ...request.Option) (*elasticbeanstalk.ApplicationDescriptionMessage, error)
	UpdateApplicationRequest(*elasticbeanstalk.UpdateApplicationInput) (*request.Request, *elasticbeanstalk.ApplicationDescriptionMessage)

	UpdateApplicationResourceLifecycle(*elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error)
	UpdateApplicationResourceLifecycleWithContext(aws.Context, *elasticbeanstalk.UpdateApplicationResourceLifecycleInput, ...request.Option) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error)
	UpdateApplicationResourceLifecycleRequest(*elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*request.Request, *elasticbeanstalk.UpdateApplicationResourceLifecycleOutput)

	UpdateApplicationVersion(*elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)
	UpdateApplicationVersionWithContext(aws.Context, *elasticbeanstalk.UpdateApplicationVersionInput, ...request.Option) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)
	UpdateApplicationVersionRequest(*elasticbeanstalk.UpdateApplicationVersionInput) (*request.Request, *elasticbeanstalk.ApplicationVersionDescriptionMessage)

	UpdateConfigurationTemplate(*elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)
	UpdateConfigurationTemplateWithContext(aws.Context, *elasticbeanstalk.UpdateConfigurationTemplateInput, ...request.Option) (*elasticbeanstalk.ConfigurationSettingsDescription, error)
	UpdateConfigurationTemplateRequest(*elasticbeanstalk.UpdateConfigurationTemplateInput) (*request.Request, *elasticbeanstalk.ConfigurationSettingsDescription)

	UpdateEnvironment(*elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)
	UpdateEnvironmentWithContext(aws.Context, *elasticbeanstalk.UpdateEnvironmentInput, ...request.Option) (*elasticbeanstalk.EnvironmentDescription, error)
	UpdateEnvironmentRequest(*elasticbeanstalk.UpdateEnvironmentInput) (*request.Request, *elasticbeanstalk.EnvironmentDescription)

	ValidateConfigurationSettings(*elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error)
	ValidateConfigurationSettingsWithContext(aws.Context, *elasticbeanstalk.ValidateConfigurationSettingsInput, ...request.Option) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error)
	ValidateConfigurationSettingsRequest(*elasticbeanstalk.ValidateConfigurationSettingsInput) (*request.Request, *elasticbeanstalk.ValidateConfigurationSettingsOutput)
}

var _ ElasticBeanstalkAPI = (*elasticbeanstalk.ElasticBeanstalk)(nil)
