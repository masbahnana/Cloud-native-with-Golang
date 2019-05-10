// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// Amazon EC2 Container Service (Amazon ECS) is a highly scalable, fast, container
// management service that makes it easy to run, stop, and manage Docker containers
// on a cluster of EC2 instances. Amazon ECS lets you launch and stop container-enabled
// applications with simple API calls, allows you to get the state of your cluster
// from a centralized service, and gives you access to many familiar Amazon
// EC2 features like security groups, Amazon EBS volumes, and IAM roles.
//
// You can use Amazon ECS to schedule the placement of containers across your
// cluster based on your resource needs, isolation policies, and availability
// requirements. Amazon EC2 Container Service eliminates the need for you to
// operate your own cluster management and configuration management systems
// or worry about scaling your management infrastructure.
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13
type ECS struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "ecs"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ECS client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ECS client from just a session.
//     svc := ecs.New(mySession)
//
//     // Create a ECS client with additional configuration
//     svc := ecs.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ECS {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ECS {
	svc := &ECS{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-11-13",
				JSONVersion:   "1.1",
				TargetPrefix:  "AmazonEC2ContainerServiceV20141113",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ECS operation and runs any
// custom request initialization.
func (c *ECS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
