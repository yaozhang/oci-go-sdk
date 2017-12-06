// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DbNodeActionRequest wrapper for the DbNodeAction operation
type DbNodeActionRequest struct {

	// The database node [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	DbNodeID *string `mandatory:"true" contributesTo:"path" name:"dbNodeId"`

	// The action to perform on the DB Node.
	Action DbNodeActionActionEnum `mandatory:"true" contributesTo:"query" name:"action"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request DbNodeActionRequest) String() string {
	return common.PointerString(request)
}

// DbNodeActionResponse wrapper for the DbNodeAction operation
type DbNodeActionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DbNode instance
	DbNode `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response DbNodeActionResponse) String() string {
	return common.PointerString(response)
}

type DbNodeActionActionEnum string

const (
	DB_NODE_ACTION_ACTION_STOP      DbNodeActionActionEnum = "STOP"
	DB_NODE_ACTION_ACTION_START     DbNodeActionActionEnum = "START"
	DB_NODE_ACTION_ACTION_SOFTRESET DbNodeActionActionEnum = "SOFTRESET"
	DB_NODE_ACTION_ACTION_RESET     DbNodeActionActionEnum = "RESET"
	DB_NODE_ACTION_ACTION_UNKNOWN   DbNodeActionActionEnum = "UNKNOWN"
)

var mapping_Database_action = map[string]DbNodeActionActionEnum{
	"STOP":      DB_NODE_ACTION_ACTION_STOP,
	"START":     DB_NODE_ACTION_ACTION_START,
	"SOFTRESET": DB_NODE_ACTION_ACTION_SOFTRESET,
	"RESET":     DB_NODE_ACTION_ACTION_RESET,
	"UNKNOWN":   DB_NODE_ACTION_ACTION_UNKNOWN,
}

func GetDbNodeActionActionEnumValues() []DbNodeActionActionEnum {
	values := make([]DbNodeActionActionEnum, 0)
	for _, v := range mapping_Database_action {
		if v != DB_NODE_ACTION_ACTION_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}