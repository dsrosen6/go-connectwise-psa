package types

type ConfigurationTypeQuestionInfo struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	ConfigurationType ConfigurationType      `json:"configurationType,omitempty"`
	EntryType         string                 `json:"entryType,omitempty"`
	FieldType         string                 `json:"fieldType,omitempty"`
	ID                int                    `json:"id,omitempty"`
	InactiveFlag      bool                   `json:"inactiveFlag,omitempty"`
	NumberOfDecimals  int                    `json:"numberOfDecimals,omitempty"`
	Question          string                 `json:"question,omitempty"`
	RequiredFlag      bool                   `json:"requiredFlag,omitempty"`
	SequenceNumber    float64                `json:"sequenceNumber,omitempty"`
}

type ConfigurationTypeQuestionValueInfo struct {
	Info              map[string]interface{}    `json:"_info,omitempty"`
	ConfigurationType ConfigurationType         `json:"configurationType,omitempty"`
	DefaultFlag       bool                      `json:"defaultFlag,omitempty"`
	ID                int                       `json:"id,omitempty"`
	InactiveFlag      bool                      `json:"inactiveFlag,omitempty"`
	Question          ConfigurationTypeQuestion `json:"question,omitempty"`
	Value             string                    `json:"value,omitempty"`
}
