// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

const (

	// ErrCodeAccessDenied for service response error code
	// "AccessDenied".
	//
	// Access denied.
	ErrCodeAccessDenied = "AccessDenied"

	// ErrCodeBatchTooLarge for service response error code
	// "BatchTooLarge".
	ErrCodeBatchTooLarge = "BatchTooLarge"

	// ErrCodeCNAMEAlreadyExists for service response error code
	// "CNAMEAlreadyExists".
	ErrCodeCNAMEAlreadyExists = "CNAMEAlreadyExists"

	// ErrCodeDistributionAlreadyExists for service response error code
	// "DistributionAlreadyExists".
	//
	// The caller reference you attempted to create the distribution with is associated
	// with another distribution.
	ErrCodeDistributionAlreadyExists = "DistributionAlreadyExists"

	// ErrCodeDistributionNotDisabled for service response error code
	// "DistributionNotDisabled".
	ErrCodeDistributionNotDisabled = "DistributionNotDisabled"

	// ErrCodeIllegalUpdate for service response error code
	// "IllegalUpdate".
	//
	// Origin and CallerReference cannot be updated.
	ErrCodeIllegalUpdate = "IllegalUpdate"

	// ErrCodeInconsistentQuantities for service response error code
	// "InconsistentQuantities".
	//
	// The value of Quantity and the size of Items do not match.
	ErrCodeInconsistentQuantities = "InconsistentQuantities"

	// ErrCodeInvalidArgument for service response error code
	// "InvalidArgument".
	//
	// The argument is invalid.
	ErrCodeInvalidArgument = "InvalidArgument"

	// ErrCodeInvalidDefaultRootObject for service response error code
	// "InvalidDefaultRootObject".
	//
	// The default root object file name is too big or contains an invalid character.
	ErrCodeInvalidDefaultRootObject = "InvalidDefaultRootObject"

	// ErrCodeInvalidErrorCode for service response error code
	// "InvalidErrorCode".
	ErrCodeInvalidErrorCode = "InvalidErrorCode"

	// ErrCodeInvalidForwardCookies for service response error code
	// "InvalidForwardCookies".
	//
	// Your request contains forward cookies option which doesn't match with the
	// expectation for the whitelisted list of cookie names. Either list of cookie
	// names has been specified when not allowed or list of cookie names is missing
	// when expected.
	ErrCodeInvalidForwardCookies = "InvalidForwardCookies"

	// ErrCodeInvalidGeoRestrictionParameter for service response error code
	// "InvalidGeoRestrictionParameter".
	ErrCodeInvalidGeoRestrictionParameter = "InvalidGeoRestrictionParameter"

	// ErrCodeInvalidHeadersForS3Origin for service response error code
	// "InvalidHeadersForS3Origin".
	ErrCodeInvalidHeadersForS3Origin = "InvalidHeadersForS3Origin"

	// ErrCodeInvalidIfMatchVersion for service response error code
	// "InvalidIfMatchVersion".
	//
	// The If-Match version is missing or not valid for the distribution.
	ErrCodeInvalidIfMatchVersion = "InvalidIfMatchVersion"

	// ErrCodeInvalidLambdaFunctionAssociation for service response error code
	// "InvalidLambdaFunctionAssociation".
	//
	// The specified Lambda function association is invalid.
	ErrCodeInvalidLambdaFunctionAssociation = "InvalidLambdaFunctionAssociation"

	// ErrCodeInvalidLocationCode for service response error code
	// "InvalidLocationCode".
	ErrCodeInvalidLocationCode = "InvalidLocationCode"

	// ErrCodeInvalidMinimumProtocolVersion for service response error code
	// "InvalidMinimumProtocolVersion".
	ErrCodeInvalidMinimumProtocolVersion = "InvalidMinimumProtocolVersion"

	// ErrCodeInvalidOrigin for service response error code
	// "InvalidOrigin".
	//
	// The Amazon S3 origin server specified does not refer to a valid Amazon S3
	// bucket.
	ErrCodeInvalidOrigin = "InvalidOrigin"

	// ErrCodeInvalidOriginAccessIdentity for service response error code
	// "InvalidOriginAccessIdentity".
	//
	// The origin access identity is not valid or doesn't exist.
	ErrCodeInvalidOriginAccessIdentity = "InvalidOriginAccessIdentity"

	// ErrCodeInvalidOriginKeepaliveTimeout for service response error code
	// "InvalidOriginKeepaliveTimeout".
	ErrCodeInvalidOriginKeepaliveTimeout = "InvalidOriginKeepaliveTimeout"

	// ErrCodeInvalidOriginReadTimeout for service response error code
	// "InvalidOriginReadTimeout".
	ErrCodeInvalidOriginReadTimeout = "InvalidOriginReadTimeout"

	// ErrCodeInvalidProtocolSettings for service response error code
	// "InvalidProtocolSettings".
	//
	// You cannot specify SSLv3 as the minimum protocol version if you only want
	// to support only clients that support Server Name Indication (SNI).
	ErrCodeInvalidProtocolSettings = "InvalidProtocolSettings"

	// ErrCodeInvalidQueryStringParameters for service response error code
	// "InvalidQueryStringParameters".
	ErrCodeInvalidQueryStringParameters = "InvalidQueryStringParameters"

	// ErrCodeInvalidRelativePath for service response error code
	// "InvalidRelativePath".
	//
	// The relative path is too big, is not URL-encoded, or does not begin with
	// a slash (/).
	ErrCodeInvalidRelativePath = "InvalidRelativePath"

	// ErrCodeInvalidRequiredProtocol for service response error code
	// "InvalidRequiredProtocol".
	//
	// This operation requires the HTTPS protocol. Ensure that you specify the HTTPS
	// protocol in your request, or omit the RequiredProtocols element from your
	// distribution configuration.
	ErrCodeInvalidRequiredProtocol = "InvalidRequiredProtocol"

	// ErrCodeInvalidResponseCode for service response error code
	// "InvalidResponseCode".
	ErrCodeInvalidResponseCode = "InvalidResponseCode"

	// ErrCodeInvalidTTLOrder for service response error code
	// "InvalidTTLOrder".
	ErrCodeInvalidTTLOrder = "InvalidTTLOrder"

	// ErrCodeInvalidTagging for service response error code
	// "InvalidTagging".
	ErrCodeInvalidTagging = "InvalidTagging"

	// ErrCodeInvalidViewerCertificate for service response error code
	// "InvalidViewerCertificate".
	ErrCodeInvalidViewerCertificate = "InvalidViewerCertificate"

	// ErrCodeInvalidWebACLId for service response error code
	// "InvalidWebACLId".
	ErrCodeInvalidWebACLId = "InvalidWebACLId"

	// ErrCodeMissingBody for service response error code
	// "MissingBody".
	//
	// This operation requires a body. Ensure that the body is present and the Content-Type
	// header is set.
	ErrCodeMissingBody = "MissingBody"

	// ErrCodeNoSuchCloudFrontOriginAccessIdentity for service response error code
	// "NoSuchCloudFrontOriginAccessIdentity".
	//
	// The specified origin access identity does not exist.
	ErrCodeNoSuchCloudFrontOriginAccessIdentity = "NoSuchCloudFrontOriginAccessIdentity"

	// ErrCodeNoSuchDistribution for service response error code
	// "NoSuchDistribution".
	//
	// The specified distribution does not exist.
	ErrCodeNoSuchDistribution = "NoSuchDistribution"

	// ErrCodeNoSuchInvalidation for service response error code
	// "NoSuchInvalidation".
	//
	// The specified invalidation does not exist.
	ErrCodeNoSuchInvalidation = "NoSuchInvalidation"

	// ErrCodeNoSuchOrigin for service response error code
	// "NoSuchOrigin".
	//
	// No origin exists with the specified Origin Id.
	ErrCodeNoSuchOrigin = "NoSuchOrigin"

	// ErrCodeNoSuchResource for service response error code
	// "NoSuchResource".
	ErrCodeNoSuchResource = "NoSuchResource"

	// ErrCodeNoSuchStreamingDistribution for service response error code
	// "NoSuchStreamingDistribution".
	//
	// The specified streaming distribution does not exist.
	ErrCodeNoSuchStreamingDistribution = "NoSuchStreamingDistribution"

	// ErrCodeOriginAccessIdentityAlreadyExists for service response error code
	// "OriginAccessIdentityAlreadyExists".
	//
	// If the CallerReference is a value you already sent in a previous request
	// to create an identity but the content of the CloudFrontOriginAccessIdentityConfig
	// is different from the original request, CloudFront returns a CloudFrontOriginAccessIdentityAlreadyExists
	// error.
	ErrCodeOriginAccessIdentityAlreadyExists = "OriginAccessIdentityAlreadyExists"

	// ErrCodeOriginAccessIdentityInUse for service response error code
	// "OriginAccessIdentityInUse".
	ErrCodeOriginAccessIdentityInUse = "OriginAccessIdentityInUse"

	// ErrCodePreconditionFailed for service response error code
	// "PreconditionFailed".
	//
	// The precondition given in one or more of the request-header fields evaluated
	// to false.
	ErrCodePreconditionFailed = "PreconditionFailed"

	// ErrCodeStreamingDistributionAlreadyExists for service response error code
	// "StreamingDistributionAlreadyExists".
	ErrCodeStreamingDistributionAlreadyExists = "StreamingDistributionAlreadyExists"

	// ErrCodeStreamingDistributionNotDisabled for service response error code
	// "StreamingDistributionNotDisabled".
	ErrCodeStreamingDistributionNotDisabled = "StreamingDistributionNotDisabled"

	// ErrCodeTooManyCacheBehaviors for service response error code
	// "TooManyCacheBehaviors".
	//
	// You cannot create more cache behaviors for the distribution.
	ErrCodeTooManyCacheBehaviors = "TooManyCacheBehaviors"

	// ErrCodeTooManyCertificates for service response error code
	// "TooManyCertificates".
	//
	// You cannot create anymore custom SSL/TLS certificates.
	ErrCodeTooManyCertificates = "TooManyCertificates"

	// ErrCodeTooManyCloudFrontOriginAccessIdentities for service response error code
	// "TooManyCloudFrontOriginAccessIdentities".
	//
	// Processing your request would cause you to exceed the maximum number of origin
	// access identities allowed.
	ErrCodeTooManyCloudFrontOriginAccessIdentities = "TooManyCloudFrontOriginAccessIdentities"

	// ErrCodeTooManyCookieNamesInWhiteList for service response error code
	// "TooManyCookieNamesInWhiteList".
	//
	// Your request contains more cookie names in the whitelist than are allowed
	// per cache behavior.
	ErrCodeTooManyCookieNamesInWhiteList = "TooManyCookieNamesInWhiteList"

	// ErrCodeTooManyDistributionCNAMEs for service response error code
	// "TooManyDistributionCNAMEs".
	//
	// Your request contains more CNAMEs than are allowed per distribution.
	ErrCodeTooManyDistributionCNAMEs = "TooManyDistributionCNAMEs"

	// ErrCodeTooManyDistributions for service response error code
	// "TooManyDistributions".
	//
	// Processing your request would cause you to exceed the maximum number of distributions
	// allowed.
	ErrCodeTooManyDistributions = "TooManyDistributions"

	// ErrCodeTooManyDistributionsWithLambdaAssociations for service response error code
	// "TooManyDistributionsWithLambdaAssociations".
	//
	// Processing your request would cause the maximum number of distributions with
	// Lambda function associations per owner to be exceeded.
	ErrCodeTooManyDistributionsWithLambdaAssociations = "TooManyDistributionsWithLambdaAssociations"

	// ErrCodeTooManyHeadersInForwardedValues for service response error code
	// "TooManyHeadersInForwardedValues".
	ErrCodeTooManyHeadersInForwardedValues = "TooManyHeadersInForwardedValues"

	// ErrCodeTooManyInvalidationsInProgress for service response error code
	// "TooManyInvalidationsInProgress".
	//
	// You have exceeded the maximum number of allowable InProgress invalidation
	// batch requests, or invalidation objects.
	ErrCodeTooManyInvalidationsInProgress = "TooManyInvalidationsInProgress"

	// ErrCodeTooManyLambdaFunctionAssociations for service response error code
	// "TooManyLambdaFunctionAssociations".
	//
	// Your request contains more Lambda function associations than are allowed
	// per distribution.
	ErrCodeTooManyLambdaFunctionAssociations = "TooManyLambdaFunctionAssociations"

	// ErrCodeTooManyOriginCustomHeaders for service response error code
	// "TooManyOriginCustomHeaders".
	ErrCodeTooManyOriginCustomHeaders = "TooManyOriginCustomHeaders"

	// ErrCodeTooManyOrigins for service response error code
	// "TooManyOrigins".
	//
	// You cannot create more origins for the distribution.
	ErrCodeTooManyOrigins = "TooManyOrigins"

	// ErrCodeTooManyQueryStringParameters for service response error code
	// "TooManyQueryStringParameters".
	ErrCodeTooManyQueryStringParameters = "TooManyQueryStringParameters"

	// ErrCodeTooManyStreamingDistributionCNAMEs for service response error code
	// "TooManyStreamingDistributionCNAMEs".
	ErrCodeTooManyStreamingDistributionCNAMEs = "TooManyStreamingDistributionCNAMEs"

	// ErrCodeTooManyStreamingDistributions for service response error code
	// "TooManyStreamingDistributions".
	//
	// Processing your request would cause you to exceed the maximum number of streaming
	// distributions allowed.
	ErrCodeTooManyStreamingDistributions = "TooManyStreamingDistributions"

	// ErrCodeTooManyTrustedSigners for service response error code
	// "TooManyTrustedSigners".
	//
	// Your request contains more trusted signers than are allowed per distribution.
	ErrCodeTooManyTrustedSigners = "TooManyTrustedSigners"

	// ErrCodeTrustedSignerDoesNotExist for service response error code
	// "TrustedSignerDoesNotExist".
	//
	// One or more of your trusted signers do not exist.
	ErrCodeTrustedSignerDoesNotExist = "TrustedSignerDoesNotExist"
)
