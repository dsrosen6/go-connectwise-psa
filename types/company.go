package types

import (
	"time"
)

type Company struct {
	Id                    int             `json:"id,omitempty"`
	Identifier            string          `json:"identifier"`
	Name                  string          `json:"name"`
	Status                CompanyStatus   `json:"status,omitempty"`
	AddressLine1          string          `json:"addressLine1,omitempty"`
	AddressLine2          string          `json:"addressLine2,omitempty"`
	City                  string          `json:"city,omitempty"`
	State                 string          `json:"state,omitempty"`
	Zip                   string          `json:"zip,omitempty"`
	Country               Country         `json:"country,omitempty"`
	PhoneNumber           string          `json:"phoneNumber,omitempty"`
	FaxNumber             string          `json:"faxNumber,omitempty"`
	Website               string          `json:"website,omitempty"`
	Territory             SystemLocation  `json:"territory,omitempty"`
	Market                Market          `json:"market,omitempty"`
	AccountNumber         string          `json:"accountNumber,omitempty"`
	DefaultContact        Contact         `json:"defaultContact,omitempty"`
	DateAcquired          time.Time       `json:"dateAcquired,omitempty"`
	SicCode               SicCode         `json:"sicCode,omitempty"`
	ParentCompany         *Company        `json:"parentCompany,omitempty"`
	AnnualRevenue         int             `json:"annualRevenue,omitempty"`
	NumberOfEmployees     int             `json:"numberOfEmployees,omitempty"`
	YearEstablished       int             `json:"yearEstablished,omitempty"`
	RevenueYear           int             `json:"revenueYear,omitempty"`
	OwnershipType         OwnershipType   `json:"ownershipType,omitempty"`
	TimeZoneSetup         TimeZone        `json:"timeZoneSetup,omitempty"`
	LeadSource            string          `json:"leadSource,omitempty"`
	LeadFlag              bool            `json:"leadFlag,omitempty"`
	UnsubscribeFlag       bool            `json:"unsubscribeFlag,omitempty"`
	Calendar              Calendar        `json:"calendar,omitempty"`
	UserDefinedField1     string          `json:"userDefinedField1,omitempty"`
	UserDefinedField2     string          `json:"userDefinedField2,omitempty"`
	UserDefinedField3     string          `json:"userDefinedField3,omitempty"`
	UserDefinedField4     string          `json:"userDefinedField4,omitempty"`
	UserDefinedField5     string          `json:"userDefinedField5,omitempty"`
	UserDefinedField6     string          `json:"userDefinedField6,omitempty"`
	UserDefinedField7     string          `json:"userDefinedField7,omitempty"`
	UserDefinedField8     string          `json:"userDefinedField8,omitempty"`
	UserDefinedField9     string          `json:"userDefinedField9,omitempty"`
	UserDefinedField10    string          `json:"userDefinedField10,omitempty"`
	VendorIdentifier      string          `json:"vendorIdentifier,omitempty"`
	TaxIdentifier         string          `json:"taxIdentifier,omitempty"`
	TaxCode               TaxCode         `json:"taxCode,omitempty"`
	BillingTerms          BillingTerms    `json:"billingTerms,omitempty"`
	InvoiceTemplate       InvoiceTemplate `json:"invoiceTemplate,omitempty"`
	PricingSchedule       PricingSchedule `json:"pricingSchedule,omitempty"`
	CompanyEntityType     EntityType      `json:"companyEntityType,omitempty"`
	BillToCompany         *Company        `json:"billToCompany,omitempty"`
	BillingSite           Site            `json:"billingSite,omitempty"`
	BillingContact        Contact         `json:"billingContact,omitempty"`
	InvoiceDeliveryMethod BillingDelivery `json:"invoiceDeliveryMethod,omitempty"`
	InvoiceToEmailAddress string          `json:"invoiceToEmailAddress,omitempty"`
	InvoiceCCEmailAddress string          `json:"invoiceCCEmailAddress,omitempty"`
	DeletedFlag           bool            `json:"deletedFlag,omitempty"`
	DateDeleted           time.Time       `json:"dateDeleted,omitempty"`
	DeletedBy             string          `json:"deletedBy,omitempty"`
	MobileGuid            string          `json:"mobileGuid,omitempty"`
	FacebookUrl           string          `json:"facebookUrl,omitempty"`
	TwitterUrl            string          `json:"twitterUrl,omitempty"`
	LinkedInUrl           string          `json:"linkedInUrl,omitempty"`
	Currency              Currency        `json:"currency,omitempty"`
	TerritoryManager      Member          `json:"territoryManager,omitempty"`
	ResellerIdentifier    string          `json:"resellerIdentifier,omitempty"`
	IsVendorFlag          bool            `json:"isVendorFlag,omitempty"`
	Types                 []CompanyType   `json:"types,omitempty"`
	Site                  Site            `json:"site,omitempty"`
	IntegratorTags        []string        `json:"integratorTags,omitempty"`
	Info                  Info            `json:"_info,omitempty"`
	CustomFields          []CustomField   `json:"customFields,omitempty"`
}

type Contact struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type CompanyStatus struct {
	Id                   int    `json:"id,omitempty"`
	Name                 string `json:"name"`
	DefaultFlag          bool   `json:"defaultFlag,omitempty"`
	InactiveFlag         bool   `json:"inactiveFlag,omitempty"`
	NotifyFlag           bool   `json:"notifyFlag,omitempty"`
	DisallowSavingFlag   bool   `json:"disallowSavingFlag,omitempty"`
	NotificationMessage  string `json:"notificationMessage,omitempty"`
	CustomNoteFlag       bool   `json:"customNoteFlag,omitempty"`
	CancelOpenTracksFlag bool   `json:"cancelOpenTracksFlag,omitempty"`
	Track                Track  `json:"track,omitempty"`
	Info                 Info   `json:"_info,omitempty"`
}

type CompanyType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type Track struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type Country struct {
	Id         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Info       Info   `json:"_info"`
}

type Market struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type OwnershipType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type EntityType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type Site struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}
