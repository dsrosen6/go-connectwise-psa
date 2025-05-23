package types

type Agreement struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	ChargeFirmFlag bool   `json:"chargeFirmFlag"`
	Info           struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type TaxCode struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type BillingTerms struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type InvoiceTemplate struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type BillingDelivery struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type Currency struct {
	Id                      int    `json:"id"`
	Symbol                  string `json:"symbol"`
	CurrencyCode            string `json:"currencyCode"`
	DecimalSeparator        string `json:"decimalSeparator"`
	NumberOfDecimals        int    `json:"numberOfDecimals"`
	ThousandsSeparator      string `json:"thousandsSeparator"`
	NegativeParenthesesFlag bool   `json:"negativeParenthesesFlag"`
	DisplaySymbolFlag       bool   `json:"displaySymbolFlag"`
	CurrencyIdentifier      string `json:"currencyIdentifier"`
	DisplayIdFlag           bool   `json:"displayIdFlag"`
	RightAlign              bool   `json:"rightAlign"`
	Name                    string `json:"name"`
	Info                    Info   `json:"_info"`
}
