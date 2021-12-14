// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// MacsManagedExternalHostInsightSummary Summary of a MACS-managed external host insight resource.
type MacsManagedExternalHostInsightSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	// The user-friendly name for the host. The name does not have to be unique.
	HostDisplayName *string `mandatory:"false" json:"hostDisplayName"`

	// Operations Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
	HostType *string `mandatory:"false" json:"hostType"`

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	ProcessorCount *int `mandatory:"false" json:"processorCount"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time the the host insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the host insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Platform type.
	PlatformType MacsManagedExternalHostInsightSummaryPlatformTypeEnum `mandatory:"false" json:"platformType,omitempty"`

	// Indicates the status of a host insight in Operations Insights
	Status ResourceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The current state of the host.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m MacsManagedExternalHostInsightSummary) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m MacsManagedExternalHostInsightSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetHostName returns HostName
func (m MacsManagedExternalHostInsightSummary) GetHostName() *string {
	return m.HostName
}

//GetHostDisplayName returns HostDisplayName
func (m MacsManagedExternalHostInsightSummary) GetHostDisplayName() *string {
	return m.HostDisplayName
}

//GetHostType returns HostType
func (m MacsManagedExternalHostInsightSummary) GetHostType() *string {
	return m.HostType
}

//GetProcessorCount returns ProcessorCount
func (m MacsManagedExternalHostInsightSummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

//GetFreeformTags returns FreeformTags
func (m MacsManagedExternalHostInsightSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m MacsManagedExternalHostInsightSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m MacsManagedExternalHostInsightSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetStatus returns Status
func (m MacsManagedExternalHostInsightSummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

//GetTimeCreated returns TimeCreated
func (m MacsManagedExternalHostInsightSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m MacsManagedExternalHostInsightSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m MacsManagedExternalHostInsightSummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m MacsManagedExternalHostInsightSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m MacsManagedExternalHostInsightSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m MacsManagedExternalHostInsightSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMacsManagedExternalHostInsightSummary MacsManagedExternalHostInsightSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeMacsManagedExternalHostInsightSummary
	}{
		"MACS_MANAGED_EXTERNAL_HOST",
		(MarshalTypeMacsManagedExternalHostInsightSummary)(m),
	}

	return json.Marshal(&s)
}

// MacsManagedExternalHostInsightSummaryPlatformTypeEnum Enum with underlying type: string
type MacsManagedExternalHostInsightSummaryPlatformTypeEnum string

// Set of constants representing the allowable values for MacsManagedExternalHostInsightSummaryPlatformTypeEnum
const (
	MacsManagedExternalHostInsightSummaryPlatformTypeLinux MacsManagedExternalHostInsightSummaryPlatformTypeEnum = "LINUX"
)

var mappingMacsManagedExternalHostInsightSummaryPlatformType = map[string]MacsManagedExternalHostInsightSummaryPlatformTypeEnum{
	"LINUX": MacsManagedExternalHostInsightSummaryPlatformTypeLinux,
}

// GetMacsManagedExternalHostInsightSummaryPlatformTypeEnumValues Enumerates the set of values for MacsManagedExternalHostInsightSummaryPlatformTypeEnum
func GetMacsManagedExternalHostInsightSummaryPlatformTypeEnumValues() []MacsManagedExternalHostInsightSummaryPlatformTypeEnum {
	values := make([]MacsManagedExternalHostInsightSummaryPlatformTypeEnum, 0)
	for _, v := range mappingMacsManagedExternalHostInsightSummaryPlatformType {
		values = append(values, v)
	}
	return values
}
