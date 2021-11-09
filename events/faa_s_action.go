// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v51/common"
)

// FaaSAction An action that delivers to an Oracle Functions endpoint.
type FaaSAction struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the action.
	Id *string `mandatory:"true" json:"id"`

	// A message generated by the Events service about the current state of this action.
	LifecycleMessage *string `mandatory:"true" json:"lifecycleMessage"`

	// Whether or not this action is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Function hosted by Oracle Functions Service.
	FunctionId *string `mandatory:"false" json:"functionId"`

	// The current state of the rule.
	LifecycleState ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m FaaSAction) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m FaaSAction) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m FaaSAction) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

//GetIsEnabled returns IsEnabled
func (m FaaSAction) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m FaaSAction) GetDescription() *string {
	return m.Description
}

func (m FaaSAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FaaSAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFaaSAction FaaSAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeFaaSAction
	}{
		"FAAS",
		(MarshalTypeFaaSAction)(m),
	}

	return json.Marshal(&s)
}
