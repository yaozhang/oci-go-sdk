// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Cpe. An object you create when setting up an IPSec VPN between your on-premises network
// and VCN. The `Cpe` is a virtual representation of your Customer-Premises Equipment,
// which is the actual router on-premises at your site at your end of the IPSec VPN connection.
// For more information,
// see [Overview of the Networking Service]({{DOC_SERVER_URL}}/Content/Network/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Cpe struct {

	// The OCID of the compartment containing the CPE.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The CPE's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The public IP address of the on-premises router.
	IpAddress *string `mandatory:"true" json:"ipAddress,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The date and time the CPE was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model Cpe) String() string {
	return common.PointerString(model)
}