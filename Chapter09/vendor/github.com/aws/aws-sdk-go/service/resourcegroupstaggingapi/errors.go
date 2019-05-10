// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroupstaggingapi

const (

	// ErrCodeInternalServiceException for service response error code
	// "InternalServiceException".
	//
	// The request processing failed because of an unknown error, exception, or
	// failure. You can retry the request.
	ErrCodeInternalServiceException = "InternalServiceException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// A parameter is missing or a malformed string or invalid or out-of-range value
	// was supplied for the request parameter.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodePaginationTokenExpiredException for service response error code
	// "PaginationTokenExpiredException".
	//
	// A PaginationToken is valid for a maximum of 15 minutes. Your request was
	// denied because the specified PaginationToken has expired.
	ErrCodePaginationTokenExpiredException = "PaginationTokenExpiredException"

	// ErrCodeThrottledException for service response error code
	// "ThrottledException".
	//
	// The request was denied to limit the frequency of submitted requests.
	ErrCodeThrottledException = "ThrottledException"
)
