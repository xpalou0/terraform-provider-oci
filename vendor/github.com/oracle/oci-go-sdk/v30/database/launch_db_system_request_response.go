// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"net/http"
)

// LaunchDbSystemRequest wrapper for the LaunchDbSystem operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/LaunchDbSystem.go.html to see an example of how to use LaunchDbSystemRequest.
type LaunchDbSystemRequest struct {

	// Request to launch a DB system.
	LaunchDbSystemDetails LaunchDbSystemBase `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request LaunchDbSystemRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request LaunchDbSystemRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request LaunchDbSystemRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// LaunchDbSystemResponse wrapper for the LaunchDbSystem operation
type LaunchDbSystemResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DbSystem instance
	DbSystem `presentIn:"body"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request. Multiple OCID values are returned in a comma-separated list. Use GetWorkRequest with a work request OCID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response LaunchDbSystemResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response LaunchDbSystemResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
