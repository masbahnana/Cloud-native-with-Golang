// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// Overview
//
// AWS Service Catalog (https://aws.amazon.com/servicecatalog/) allows organizations
// to create and manage catalogs of IT services that are approved for use on
// AWS. This documentation provides reference material for the AWS Service Catalog
// end user API. To get the most out of this documentation, you need to be familiar
// with the terminology discussed in AWS Service Catalog Concepts (http://docs.aws.amazon.com/servicecatalog/latest/userguide/what-is_concepts.html).
//
// Additional Resources
//
//    * AWS Service Catalog Administrator Guide (http://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html)
//
//    * AWS Service Catalog User Guide (http://docs.aws.amazon.com/servicecatalog/latest/userguide/introduction.html)
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10
type ServiceCatalog struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "servicecatalog" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName      // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ServiceCatalog client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ServiceCatalog client from just a session.
//     svc := servicecatalog.New(mySession)
//
//     // Create a ServiceCatalog client with additional configuration
//     svc := servicecatalog.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ServiceCatalog {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ServiceCatalog {
	svc := &ServiceCatalog{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-12-10",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWS242ServiceCatalogService",
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

// newRequest creates a new request for a ServiceCatalog operation and runs any
// custom request initialization.
func (c *ServiceCatalog) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
