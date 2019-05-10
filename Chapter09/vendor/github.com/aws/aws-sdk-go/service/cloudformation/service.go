// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/query"
)

// AWS CloudFormation allows you to create and manage AWS infrastructure deployments
// predictably and repeatedly. You can use AWS CloudFormation to leverage AWS
// products, such as Amazon Elastic Compute Cloud, Amazon Elastic Block Store,
// Amazon Simple Notification Service, Elastic Load Balancing, and Auto Scaling
// to build highly-reliable, highly scalable, cost-effective applications without
// creating or configuring the underlying AWS infrastructure.
//
// With AWS CloudFormation, you declare all of your resources and dependencies
// in a template file. The template defines a collection of resources as a single
// unit called a stack. AWS CloudFormation creates and deletes all member resources
// of the stack together and manages all dependencies between the resources
// for you.
//
// For more information about AWS CloudFormation, see the AWS CloudFormation
// Product Page (http://aws.amazon.com/cloudformation/).
//
// Amazon CloudFormation makes use of other AWS products. If you need additional
// technical information about a specific AWS product, you can find the product's
// technical documentation at docs.aws.amazon.com (http://docs.aws.amazon.com/).
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15
type CloudFormation struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "cloudformation" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName      // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the CloudFormation client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CloudFormation client from just a session.
//     svc := cloudformation.New(mySession)
//
//     // Create a CloudFormation client with additional configuration
//     svc := cloudformation.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CloudFormation {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *CloudFormation {
	svc := &CloudFormation{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2010-05-15",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a CloudFormation operation and runs any
// custom request initialization.
func (c *CloudFormation) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
