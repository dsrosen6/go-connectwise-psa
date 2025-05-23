package types

type StatusIndicator struct {
	Id         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Info       Info   `json:"_info"`
}

type SicCode struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type CustomFieldValue struct {
}

type CustomField struct {
	Id               int              `json:"id"`
	Caption          string           `json:"caption"`
	Type             string           `json:"type"`
	EntryMethod      string           `json:"entryMethod"`
	NumberOfDecimals int              `json:"numberOfDecimals"`
	Value            CustomFieldValue `json:"value"`
	ConnectWiseId    string           `json:"connectWiseId"`
}

type Info struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
