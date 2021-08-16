// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// DomainGovernance The model for a domain governance entity.
type DomainGovernance struct {

	// The OCID of the domain governance entity.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy that owns this domain governance entity.
	OwnerId *string `mandatory:"true" json:"ownerId"`

	// The OCID of the domain associated with this domain governance entity.
	DomainId *string `mandatory:"true" json:"domainId"`

	// Lifecycle state of the domain governance entity.
	LifecycleState DomainGovernanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The ONS topic associated with this domain governance entity.
	OnsTopicId *string `mandatory:"true" json:"onsTopicId"`

	// The ONS subscription associated with this domain governance entity.
	OnsSubscriptionId *string `mandatory:"true" json:"onsSubscriptionId"`

	// Indicates whether governance is enabled for this domain.
	IsGovernanceEnabled *bool `mandatory:"false" json:"isGovernanceEnabled"`

	// The email to notify the user, and that the ONS subscription will be created with.
	SubscriptionEmail *string `mandatory:"false" json:"subscriptionEmail"`

	// Date-time when this domain governance was created. An RFC3339-formatted date and time string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date-time when this domain governance was last updated. An RFC3339-formatted date and time string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DomainGovernance) String() string {
	return common.PointerString(m)
}

// DomainGovernanceLifecycleStateEnum Enum with underlying type: string
type DomainGovernanceLifecycleStateEnum string

// Set of constants representing the allowable values for DomainGovernanceLifecycleStateEnum
const (
	DomainGovernanceLifecycleStateActive   DomainGovernanceLifecycleStateEnum = "ACTIVE"
	DomainGovernanceLifecycleStateInactive DomainGovernanceLifecycleStateEnum = "INACTIVE"
)

var mappingDomainGovernanceLifecycleState = map[string]DomainGovernanceLifecycleStateEnum{
	"ACTIVE":   DomainGovernanceLifecycleStateActive,
	"INACTIVE": DomainGovernanceLifecycleStateInactive,
}

// GetDomainGovernanceLifecycleStateEnumValues Enumerates the set of values for DomainGovernanceLifecycleStateEnum
func GetDomainGovernanceLifecycleStateEnumValues() []DomainGovernanceLifecycleStateEnum {
	values := make([]DomainGovernanceLifecycleStateEnum, 0)
	for _, v := range mappingDomainGovernanceLifecycleState {
		values = append(values, v)
	}
	return values
}
