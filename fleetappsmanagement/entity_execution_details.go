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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EntityExecutionDetails Activity Resource and execution details including outcome.
type EntityExecutionDetails struct {

	// Resource Identifier associated with the Work Request
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Status of the Job at Resource Level
	Status JobStatusEnum `mandatory:"true" json:"status"`

	// Resource Display Name
	ResourceDisplayName *string `mandatory:"false" json:"resourceDisplayName"`

	// Description of the Work Request
	Description *string `mandatory:"false" json:"description"`

	// The sequence of the Resource
	Sequence *string `mandatory:"false" json:"sequence"`

	// Targets associated.
	Targets []ActivityResourceTarget `mandatory:"false" json:"targets"`

	// The time the task started for the resource. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the task ended for the resource. An RFC3339 formatted datetime string
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m EntityExecutionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EntityExecutionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}