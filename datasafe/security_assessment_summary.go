// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v51/common"
)

// SecurityAssessmentSummary The summary of a security assessment.
type SecurityAssessmentSummary struct {

	// The OCID of the security assessment.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the security assessment.
	LifecycleState SecurityAssessmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the security assessment was created. Conforms to the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time when the security assessment was last updated. Conforms to the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the compartment that contains the security assessment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the security assessment.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Array of database target OCIDs.
	TargetIds []string `mandatory:"true" json:"targetIds"`

	// The type of the security assessment. Possible values are:
	// LATEST: The most up-to-date assessment that is running automatically for a target. It is system generated.
	// SAVED: A saved security assessment. LATEST assessments are always saved in order to maintain the history of runs. A SAVED assessment is also generated by a 'refresh' action (triggered by the user).
	// SAVE_SCHEDULE: The schedule for periodic saves of LATEST assessments.
	// COMPARTMENT: An automatically managed assessment type that stores all details of targets in one compartment.
	// This type keeps an up-to-date assessment of all database risks in one compartment. It is automatically updated when the latest assessment or refresh action is executed. It is also automatically updated when a target is deleted or move to a different compartment.
	Type SecurityAssessmentSummaryTypeEnum `mandatory:"true" json:"type"`

	// The description of the security assessment.
	Description *string `mandatory:"false" json:"description"`

	// Details about the current state of the security assessment.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// List containing maps as values.
	// Example: `{"Operations": [ {"CostCenter": "42"} ] }`
	IgnoredTargetIds []interface{} `mandatory:"false" json:"ignoredTargetIds"`

	// List containing maps as values.
	// Example: `{"Operations": [ {"CostCenter": "42"} ] }`
	IgnoredAssessmentIds []interface{} `mandatory:"false" json:"ignoredAssessmentIds"`

	// Indicates whether or not the assessment is a baseline assessment. This applied to saved security assessments only.
	IsBaseline *bool `mandatory:"false" json:"isBaseline"`

	// Indicates whether or not the security assessment deviates from the baseline.
	IsDeviatedFromBaseline *bool `mandatory:"false" json:"isDeviatedFromBaseline"`

	// The OCID of the baseline against which the latest assessment was compared.
	LastComparedBaselineId *string `mandatory:"false" json:"lastComparedBaselineId"`

	// The OCID of the security assessment that created this scheduled save assessment.
	ScheduleSecurityAssessmentId *string `mandatory:"false" json:"scheduleSecurityAssessmentId"`

	// Schedule of the assessment that runs periodically in the specified format: -
	// <version-string>;<version-specific-schedule>
	// Allowed version strings - "v1"
	// v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month>
	// Each of the above fields potentially introduce constraints. A workrequest is created only
	// when clock time satisfies all the constraints. Constraints introduced:
	// 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59])
	// 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59])
	// 3. hours = <hh> (So, the allowed range for <hh> is [0, 23])
	// <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday))
	// 4. No constraint introduced when it is '*'. When not, day of week must equal the given value
	// <day-of-month> can be either '*' (without quotes or a number between 1 and 28)
	// 5. No constraint introduced when it is '*'. When not, day of month must equal the given value
	Schedule *string `mandatory:"false" json:"schedule"`

	// Indicates whether the security assessment was created by system or by a user.
	TriggeredBy SecurityAssessmentSummaryTriggeredByEnum `mandatory:"false" json:"triggeredBy,omitempty"`

	// The summary of findings for the security assessment.
	Link *string `mandatory:"false" json:"link"`

	Statistics *SecurityAssessmentStatistics `mandatory:"false" json:"statistics"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SecurityAssessmentSummary) String() string {
	return common.PointerString(m)
}

// SecurityAssessmentSummaryTriggeredByEnum Enum with underlying type: string
type SecurityAssessmentSummaryTriggeredByEnum string

// Set of constants representing the allowable values for SecurityAssessmentSummaryTriggeredByEnum
const (
	SecurityAssessmentSummaryTriggeredByUser   SecurityAssessmentSummaryTriggeredByEnum = "USER"
	SecurityAssessmentSummaryTriggeredBySystem SecurityAssessmentSummaryTriggeredByEnum = "SYSTEM"
)

var mappingSecurityAssessmentSummaryTriggeredBy = map[string]SecurityAssessmentSummaryTriggeredByEnum{
	"USER":   SecurityAssessmentSummaryTriggeredByUser,
	"SYSTEM": SecurityAssessmentSummaryTriggeredBySystem,
}

// GetSecurityAssessmentSummaryTriggeredByEnumValues Enumerates the set of values for SecurityAssessmentSummaryTriggeredByEnum
func GetSecurityAssessmentSummaryTriggeredByEnumValues() []SecurityAssessmentSummaryTriggeredByEnum {
	values := make([]SecurityAssessmentSummaryTriggeredByEnum, 0)
	for _, v := range mappingSecurityAssessmentSummaryTriggeredBy {
		values = append(values, v)
	}
	return values
}

// SecurityAssessmentSummaryTypeEnum Enum with underlying type: string
type SecurityAssessmentSummaryTypeEnum string

// Set of constants representing the allowable values for SecurityAssessmentSummaryTypeEnum
const (
	SecurityAssessmentSummaryTypeLatest       SecurityAssessmentSummaryTypeEnum = "LATEST"
	SecurityAssessmentSummaryTypeSaved        SecurityAssessmentSummaryTypeEnum = "SAVED"
	SecurityAssessmentSummaryTypeSaveSchedule SecurityAssessmentSummaryTypeEnum = "SAVE_SCHEDULE"
	SecurityAssessmentSummaryTypeCompartment  SecurityAssessmentSummaryTypeEnum = "COMPARTMENT"
)

var mappingSecurityAssessmentSummaryType = map[string]SecurityAssessmentSummaryTypeEnum{
	"LATEST":        SecurityAssessmentSummaryTypeLatest,
	"SAVED":         SecurityAssessmentSummaryTypeSaved,
	"SAVE_SCHEDULE": SecurityAssessmentSummaryTypeSaveSchedule,
	"COMPARTMENT":   SecurityAssessmentSummaryTypeCompartment,
}

// GetSecurityAssessmentSummaryTypeEnumValues Enumerates the set of values for SecurityAssessmentSummaryTypeEnum
func GetSecurityAssessmentSummaryTypeEnumValues() []SecurityAssessmentSummaryTypeEnum {
	values := make([]SecurityAssessmentSummaryTypeEnum, 0)
	for _, v := range mappingSecurityAssessmentSummaryType {
		values = append(values, v)
	}
	return values
}
