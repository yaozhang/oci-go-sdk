// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssociateMonitoredResourcesDetails The information required to create new monitored resource association.
type AssociateMonitoredResourcesDetails struct {

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Association type to be created between source and destination resources
	AssociationType *string `mandatory:"true" json:"associationType"`

	// Source Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	SourceResourceId *string `mandatory:"true" json:"sourceResourceId"`

	// Destination Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	DestinationResourceId *string `mandatory:"true" json:"destinationResourceId"`
}

func (m AssociateMonitoredResourcesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociateMonitoredResourcesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
