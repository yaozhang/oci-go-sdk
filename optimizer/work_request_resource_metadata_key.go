// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"strings"
)

// WorkRequestResourceMetadataKeyEnum Enum with underlying type: string
type WorkRequestResourceMetadataKeyEnum string

// Set of constants representing the allowable values for WorkRequestResourceMetadataKeyEnum
const (
	WorkRequestResourceMetadataKeyOperationName WorkRequestResourceMetadataKeyEnum = "OPERATION_NAME"
)

var mappingWorkRequestResourceMetadataKeyEnum = map[string]WorkRequestResourceMetadataKeyEnum{
	"OPERATION_NAME": WorkRequestResourceMetadataKeyOperationName,
}

var mappingWorkRequestResourceMetadataKeyEnumLowerCase = map[string]WorkRequestResourceMetadataKeyEnum{
	"operation_name": WorkRequestResourceMetadataKeyOperationName,
}

// GetWorkRequestResourceMetadataKeyEnumValues Enumerates the set of values for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumValues() []WorkRequestResourceMetadataKeyEnum {
	values := make([]WorkRequestResourceMetadataKeyEnum, 0)
	for _, v := range mappingWorkRequestResourceMetadataKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceMetadataKeyEnumStringValues Enumerates the set of values in String for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumStringValues() []string {
	return []string{
		"OPERATION_NAME",
	}
}

// GetMappingWorkRequestResourceMetadataKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceMetadataKeyEnum(val string) (WorkRequestResourceMetadataKeyEnum, bool) {
	enum, ok := mappingWorkRequestResourceMetadataKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
