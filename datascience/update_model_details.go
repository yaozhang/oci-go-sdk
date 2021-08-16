// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// UpdateModelDetails Details for updating a model.
type UpdateModelDetails struct {

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	//  Example: `My Model`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the model.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An array of custom metadata details for the model.
	CustomMetadataList []Metadata `mandatory:"false" json:"customMetadataList"`

	// An array of defined metadata details for the model.
	DefinedMetadataList []Metadata `mandatory:"false" json:"definedMetadataList"`
}

func (m UpdateModelDetails) String() string {
	return common.PointerString(m)
}
