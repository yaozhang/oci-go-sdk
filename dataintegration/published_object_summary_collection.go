// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// PublishedObjectSummaryCollection This is the collection of published object summaries, it may be a collection of lightweight details or full definitions.
type PublishedObjectSummaryCollection struct {

	// The array of published object summaries.
	Items []PublishedObjectSummary `mandatory:"true" json:"items"`
}

func (m PublishedObjectSummaryCollection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *PublishedObjectSummaryCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []publishedobjectsummary `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]PublishedObjectSummary, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(PublishedObjectSummary)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
