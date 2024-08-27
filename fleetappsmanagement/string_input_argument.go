// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StringInputArgument The details of the String Input argument.
type StringInputArgument struct {

	// The name of the argument
	Name *string `mandatory:"true" json:"name"`

	// The description of the argument.
	Description *string `mandatory:"false" json:"description"`
}

// GetName returns Name
func (m StringInputArgument) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m StringInputArgument) GetDescription() *string {
	return m.Description
}

func (m StringInputArgument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StringInputArgument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StringInputArgument) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStringInputArgument StringInputArgument
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStringInputArgument
	}{
		"STRING",
		(MarshalTypeStringInputArgument)(m),
	}

	return json.Marshal(&s)
}