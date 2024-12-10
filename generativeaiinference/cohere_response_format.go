// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to /EN/generative-ai-inference/latest/ChatResult/Chat, /EN/generative-ai-inference/latest/GenerateTextResult/GenerateText, /EN/generative-ai-inference/latest/SummarizeTextResult/SummarizeText, and /EN/generative-ai-inference/latest/EmbedTextResult/EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the /EN/generative-ai/latest/ to /EN/generative-ai/latest/Model/ by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a /EN/generative-ai/latest/DedicatedAiCluster/. Then, create a /EN/generative-ai/latest/DedicatedAiCluster/ with an Endpoint to host your custom model. For resource management in the Generative AI service, use the /EN/generative-ai/latest/.
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CohereResponseFormat Specify the format the model output is guaranteed to be of
type CohereResponseFormat interface {
}

type cohereresponseformat struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *cohereresponseformat) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercohereresponseformat cohereresponseformat
	s := struct {
		Model Unmarshalercohereresponseformat
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *cohereresponseformat) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TEXT":
		mm := CohereResponseTextFormat{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JSON_OBJECT":
		mm := CohereResponseJsonFormat{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CohereResponseFormat: %s.", m.Type)
		return *m, nil
	}
}

func (m cohereresponseformat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m cohereresponseformat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CohereResponseFormatTypeEnum Enum with underlying type: string
type CohereResponseFormatTypeEnum string

// Set of constants representing the allowable values for CohereResponseFormatTypeEnum
const (
	CohereResponseFormatTypeJsonObject CohereResponseFormatTypeEnum = "JSON_OBJECT"
	CohereResponseFormatTypeText       CohereResponseFormatTypeEnum = "TEXT"
)

var mappingCohereResponseFormatTypeEnum = map[string]CohereResponseFormatTypeEnum{
	"JSON_OBJECT": CohereResponseFormatTypeJsonObject,
	"TEXT":        CohereResponseFormatTypeText,
}

var mappingCohereResponseFormatTypeEnumLowerCase = map[string]CohereResponseFormatTypeEnum{
	"json_object": CohereResponseFormatTypeJsonObject,
	"text":        CohereResponseFormatTypeText,
}

// GetCohereResponseFormatTypeEnumValues Enumerates the set of values for CohereResponseFormatTypeEnum
func GetCohereResponseFormatTypeEnumValues() []CohereResponseFormatTypeEnum {
	values := make([]CohereResponseFormatTypeEnum, 0)
	for _, v := range mappingCohereResponseFormatTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereResponseFormatTypeEnumStringValues Enumerates the set of values in String for CohereResponseFormatTypeEnum
func GetCohereResponseFormatTypeEnumStringValues() []string {
	return []string{
		"JSON_OBJECT",
		"TEXT",
	}
}

// GetMappingCohereResponseFormatTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereResponseFormatTypeEnum(val string) (CohereResponseFormatTypeEnum, bool) {
	enum, ok := mappingCohereResponseFormatTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}