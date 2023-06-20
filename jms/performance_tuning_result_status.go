// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// PerformanceTuningResultStatusEnum Enum with underlying type: string
type PerformanceTuningResultStatusEnum string

// Set of constants representing the allowable values for PerformanceTuningResultStatusEnum
const (
	PerformanceTuningResultStatusActionRecommended PerformanceTuningResultStatusEnum = "ACTION_RECOMMENDED"
	PerformanceTuningResultStatusNoWarnings        PerformanceTuningResultStatusEnum = "NO_WARNINGS"
)

var mappingPerformanceTuningResultStatusEnum = map[string]PerformanceTuningResultStatusEnum{
	"ACTION_RECOMMENDED": PerformanceTuningResultStatusActionRecommended,
	"NO_WARNINGS":        PerformanceTuningResultStatusNoWarnings,
}

var mappingPerformanceTuningResultStatusEnumLowerCase = map[string]PerformanceTuningResultStatusEnum{
	"action_recommended": PerformanceTuningResultStatusActionRecommended,
	"no_warnings":        PerformanceTuningResultStatusNoWarnings,
}

// GetPerformanceTuningResultStatusEnumValues Enumerates the set of values for PerformanceTuningResultStatusEnum
func GetPerformanceTuningResultStatusEnumValues() []PerformanceTuningResultStatusEnum {
	values := make([]PerformanceTuningResultStatusEnum, 0)
	for _, v := range mappingPerformanceTuningResultStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPerformanceTuningResultStatusEnumStringValues Enumerates the set of values in String for PerformanceTuningResultStatusEnum
func GetPerformanceTuningResultStatusEnumStringValues() []string {
	return []string{
		"ACTION_RECOMMENDED",
		"NO_WARNINGS",
	}
}

// GetMappingPerformanceTuningResultStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPerformanceTuningResultStatusEnum(val string) (PerformanceTuningResultStatusEnum, bool) {
	enum, ok := mappingPerformanceTuningResultStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}