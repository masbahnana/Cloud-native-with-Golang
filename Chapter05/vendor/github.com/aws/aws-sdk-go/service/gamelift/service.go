// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// Amazon GameLift is a managed service for developers who need a scalable,
// dedicated server solution for their multiplayer games. Amazon GameLift provides
// tools to acquire computing resources and deploy game servers, scale game
// server capacity to meet player demand, and track in-depth metrics on player
// usage and server performance.
//
// The Amazon GameLift service API includes important features:
//
//    * Find game sessions and match players to games – Retrieve information
//    on available game sessions; create new game sessions; send player requests
//    to join a game session.
//
//    * Configure and manage game server resources – Manage builds, fleets,
//    queues, and aliases; set autoscaling policies; retrieve logs and metrics.
//
// This reference guide describes the low-level service API for Amazon GameLift.
// We recommend using either the Amazon Web Services software development kit
// (AWS SDK (http://aws.amazon.com/tools/#sdk)), available in multiple languages,
// or the AWS command-line interface (http://aws.amazon.com/cli/) (CLI) tool.
// Both of these align with the low-level service API. In addition, you can
// use the AWS Management Console (https://console.aws.amazon.com/gamelift/home)
// for Amazon GameLift for many administrative actions.
//
// You can use some API actions with Amazon GameLift Local, a testing tool that
// lets you test your game integration locally before deploying on Amazon GameLift.
// You can call these APIs from the AWS CLI or programmatically; API calls to
// Amazon GameLift Local servers perform exactly as they do when calling Amazon
// GameLift web servers. For more information on using Amazon GameLift Local,
// see Testing an Integration (http://docs.aws.amazon.com/gamelift/latest/developerguide/integration-testing-local.html).
//
// MORE RESOURCES
//
//    * Amazon GameLift Developer Guide (http://docs.aws.amazon.com/gamelift/latest/developerguide/)
//    – Learn more about Amazon GameLift features and how to use them.
//
//    * Lumberyard and Amazon GameLift Tutorials (https://gamedev.amazon.com/forums/tutorials)
//    – Get started fast with walkthroughs and sample projects.
//
//    * GameDev Blog (http://aws.amazon.com/blogs/gamedev/) – Stay up to date
//    with new features and techniques.
//
//    * GameDev Forums (https://gamedev.amazon.com/forums/spaces/123/gamelift-discussion.html)
//    – Connect with the GameDev community.
//
//    * Amazon GameLift Document History (http://docs.aws.amazon.com/gamelift/latest/developerguide/doc-history.html)
//    – See changes to the Amazon GameLift service, SDKs, and documentation,
//    as well as links to release notes.
//
// API SUMMARY
//
// This list offers a functional overview of the Amazon GameLift service API.
//
// Finding Games and Joining Players
//
// You can enable players to connect to game servers on Amazon GameLift from
// a game client or through a game service (such as a matchmaking service).
// You can use these operations to discover actively running game or start new
// games. You can also match players to games, either singly or as a group.
//
//    * Discover existing game sessions
//
// SearchGameSessions – Get all available game sessions or search for game sessions
//    that match a set of criteria. Available in Amazon GameLift Local.
//
//    * Start a new game session
//
// Game session placement – Use a queue to process new game session requests
//    and create game sessions on fleets designated for the queue.
//
// StartGameSessionPlacement – Request a new game session placement and add
//    one or more players to it.
//
// DescribeGameSessionPlacement – Get details on a placement request, including
//    status.
//
// StopGameSessionPlacement – Cancel a placement request.
//
// CreateGameSession – Start a new game session on a specific fleet. Available
//    in Amazon GameLift Local.
//
//    * Manage game session objects
//
// DescribeGameSessions – Retrieve metadata for one or more game sessions, including
//    length of time active and current player count. Available in Amazon GameLift
//    Local.
//
// DescribeGameSessionDetails – Retrieve metadata and the game session protection
//    setting for one or more game sessions.
//
// UpdateGameSession – Change game session settings, such as maximum player
//    count and join policy.
//
// GetGameSessionLogUrl – Get the location of saved logs for a game session.
//
//    * Manage player sessions objects
//
// CreatePlayerSession – Send a request for a player to join a game session.
//    Available in Amazon GameLift Local.
//
// CreatePlayerSessions – Send a request for multiple players to join a game
//    session. Available in Amazon GameLift Local.
//
// DescribePlayerSessions – Get details on player activity, including status,
//    playing time, and player data. Available in Amazon GameLift Local.
//
// Setting Up and Managing Game Servers
//
// When setting up Amazon GameLift, first create a game build and upload the
// files to Amazon GameLift. Then use these operations to set up a fleet of
// resources to run your game servers. Manage games to scale capacity, adjust
// configuration settings, access raw utilization data, and more.
//
//    * Manage game builds
//
// CreateBuild – Create a new build by uploading files stored in an Amazon S3
//    bucket. (To create a build stored at a local file location, use the AWS
//    CLI command upload-build.)
//
// ListBuilds – Get a list of all builds uploaded to a Amazon GameLift region.
//
// DescribeBuild – Retrieve information associated with a build.
//
// UpdateBuild – Change build metadata, including build name and version.
//
// DeleteBuild – Remove a build from Amazon GameLift.
//
//    * Manage fleets
//
// CreateFleet – Configure and activate a new fleet to run a build's game servers.
//
// DeleteFleet – Terminate a fleet that is no longer running game servers or
//    hosting players.
//
// View / update fleet configurations.
//
// ListFleets – Get a list of all fleet IDs in a Amazon GameLift region (all
//    statuses).
//
// DescribeFleetAttributes / UpdateFleetAttributes – View or change a fleet's
//    metadata and settings for game session protection and resource creation
//    limits.
//
// DescribeFleetPortSettings / UpdateFleetPortSettings – View or change the
//    inbound permissions (IP address and port setting ranges) allowed for a
//    fleet.
//
// DescribeRuntimeConfiguration / UpdateRuntimeConfiguration – View or change
//    what server processes (and how many) to run on each instance in a fleet.
//
// DescribeInstances – Get information on each instance in a fleet, including
//    instance ID, IP address, and status.
//
//    * Control fleet capacity
//
// DescribeEC2InstanceLimits – Retrieve maximum number of instances allowed
//    for the current AWS account and the current usage level.
//
// DescribeFleetCapacity / UpdateFleetCapacity – Retrieve the capacity settings
//    and the current number of instances in a fleet; adjust fleet capacity
//    settings to scale up or down.
//
// Autoscale – Manage autoscaling rules and apply them to a fleet.
//
// PutScalingPolicy – Create a new autoscaling policy, or update an existing
//    one.
//
// DescribeScalingPolicies – Retrieve an existing autoscaling policy.
//
// DeleteScalingPolicy – Delete an autoscaling policy and stop it from affecting
//    a fleet's capacity.
//
//    * Access fleet activity statistics
//
// DescribeFleetUtilization – Get current data on the number of server processes,
//    game sessions, and players currently active on a fleet.
//
// DescribeFleetEvents – Get a fleet's logged events for a specified time span.
//
// DescribeGameSessions – Retrieve metadata associated with one or more game
//    sessions, including length of time active and current player count.
//
//    * Remotely access an instance
//
// GetInstanceAccess – Request access credentials needed to remotely connect
//    to a specified instance in a fleet.
//
//    * Manage fleet aliases
//
// CreateAlias – Define a new alias and optionally assign it to a fleet.
//
// ListAliases – Get all fleet aliases defined in a Amazon GameLift region.
//
// DescribeAlias – Retrieve information on an existing alias.
//
// UpdateAlias – Change settings for a alias, such as redirecting it from one
//    fleet to another.
//
// DeleteAlias – Remove an alias from the region.
//
// ResolveAlias – Get the fleet ID that a specified alias points to.
//
//    * Manage game session queues
//
// CreateGameSessionQueue – Create a queue for processing requests for new game
//    sessions.
//
// DescribeGameSessionQueues – Get data on all game session queues defined in
//    a Amazon GameLift region.
//
// UpdateGameSessionQueue – Change the configuration of a game session queue.
//
// DeleteGameSessionQueue – Remove a game session queue from the region.
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01
type GameLift struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "gamelift"  // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the GameLift client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a GameLift client from just a session.
//     svc := gamelift.New(mySession)
//
//     // Create a GameLift client with additional configuration
//     svc := gamelift.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *GameLift {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *GameLift {
	svc := &GameLift{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-10-01",
				JSONVersion:   "1.1",
				TargetPrefix:  "GameLift",
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

// newRequest creates a new request for a GameLift operation and runs any
// custom request initialization.
func (c *GameLift) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
