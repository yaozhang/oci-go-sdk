// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeleteBackendRequest wrapper for the DeleteBackend operation
type DeleteBackendRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend set and server.
	LoadBalancerID *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The name of the backend set associated with the backend server.
	// Example: `My_backend_set`
	BackendSetName *string `mandatory:"true" contributesTo:"path" name:"backendSetName"`

	// The IP address and port of the backend server to remove.
	// Example: `1.1.1.7:42`
	BackendName *string `mandatory:"true" contributesTo:"path" name:"backendName"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request DeleteBackendRequest) String() string {
	return common.PointerString(request)
}

// DeleteBackendResponse wrapper for the DeleteBackend operation
type DeleteBackendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the work request.
	OpcWorkRequestID *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response DeleteBackendResponse) String() string {
	return common.PointerString(response)
}