// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// ChangeCustomProtectionRuleCompartmentDetails The representation of ChangeCustomProtectionRuleCompartmentDetails
type ChangeCustomProtectionRuleCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment into which the resource should be moved. For information about moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeCustomProtectionRuleCompartmentDetails) String() string {
	return common.PointerString(m)
}
