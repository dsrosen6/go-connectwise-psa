package connectwise

import "time"

type Board struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	Location struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"location"`
	Department struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"department"`
	InactiveFlag    bool `json:"inactiveFlag,omitempty"`
	SignOffTemplate struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"signOffTemplate,omitempty"`
	SendToContactFlag bool `json:"sendToContactFlag,omitempty"`
	ContactTemplate   struct {
		Id         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Type       string `json:"type,omitempty"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"contactTemplate,omitempty"`
	SendToResourceFlag bool `json:"sendToResourceFlag,omitempty"`
	ResourceTemplate   struct {
		Id         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Type       string `json:"type,omitempty"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"resourceTemplate,omitempty"`
	ProjectFlag          bool `json:"projectFlag,omitempty"`
	ShowDependenciesFlag bool `json:"showDependenciesFlag,omitempty"`
	ShowEstimatesFlag    bool `json:"showEstimatesFlag,omitempty"`
	BoardIcon            struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"boardIcon,omitempty"`
	BillTicketsAfterClosedFlag    bool `json:"billTicketsAfterClosedFlag,omitempty"`
	BillTicketSeparatelyFlag      bool `json:"billTicketSeparatelyFlag,omitempty"`
	BillUnapprovedTimeExpenseFlag bool `json:"billUnapprovedTimeExpenseFlag,omitempty"`
	OverrideBillingSetupFlag      bool `json:"overrideBillingSetupFlag,omitempty"`
	DispatchMember                struct {
		Id            int    `json:"id,omitempty"`
		Identifier    string `json:"identifier,omitempty"`
		Name          string `json:"name,omitempty"`
		DailyCapacity int    `json:"dailyCapacity,omitempty"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"dispatchMember,omitempty"`
	ServiceManagerMember struct {
		Id            int    `json:"id,omitempty"`
		Identifier    string `json:"identifier,omitempty"`
		Name          string `json:"name,omitempty"`
		DailyCapacity int    `json:"dailyCapacity,omitempty"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"serviceManagerMember,omitempty"`
	DutyManagerMember struct {
		Id            int    `json:"id,omitempty"`
		Identifier    string `json:"identifier,omitempty"`
		Name          string `json:"name,omitempty"`
		DailyCapacity int    `json:"dailyCapacity,omitempty"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"dutyManagerMember,omitempty"`
	OncallMember struct {
		Id            int    `json:"id,omitempty"`
		Identifier    string `json:"identifier,omitempty"`
		Name          string `json:"name,omitempty"`
		DailyCapacity int    `json:"dailyCapacity,omitempty"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"oncallMember,omitempty"`
	WorkRole struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"workRole,omitempty"`
	WorkType struct {
		Id              int    `json:"id,omitempty"`
		Name            string `json:"name,omitempty"`
		UtilizationFlag bool   `json:"utilizationFlag,omitempty"`
		Info            struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"workType,omitempty"`
	BillTime        string `json:"billTime,omitempty"`
	BillExpense     string `json:"billExpense,omitempty"`
	BillProduct     string `json:"billProduct,omitempty"`
	AutoCloseStatus struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Sort int    `json:"sort,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"autoCloseStatus,omitempty"`
	AutoAssignNewTicketsFlag            bool   `json:"autoAssignNewTicketsFlag,omitempty"`
	AutoAssignNewECTicketsFlag          bool   `json:"autoAssignNewECTicketsFlag,omitempty"`
	AutoAssignNewPortalTicketsFlag      bool   `json:"autoAssignNewPortalTicketsFlag,omitempty"`
	DiscussionsLockedFlag               bool   `json:"discussionsLockedFlag,omitempty"`
	TimeEntryLockedFlag                 bool   `json:"timeEntryLockedFlag,omitempty"`
	NotifyEmailFrom                     string `json:"notifyEmailFrom,omitempty"`
	NotifyEmailFromName                 string `json:"notifyEmailFromName,omitempty"`
	ClosedLoopDiscussionsFlag           bool   `json:"closedLoopDiscussionsFlag,omitempty"`
	ClosedLoopResolutionFlag            bool   `json:"closedLoopResolutionFlag,omitempty"`
	ClosedLoopInternalAnalysisFlag      bool   `json:"closedLoopInternalAnalysisFlag,omitempty"`
	TimeEntryDiscussionFlag             bool   `json:"timeEntryDiscussionFlag,omitempty"`
	TimeEntryResolutionFlag             bool   `json:"timeEntryResolutionFlag,omitempty"`
	TimeEntryInternalAnalysisFlag       bool   `json:"timeEntryInternalAnalysisFlag,omitempty"`
	ProblemSort                         string `json:"problemSort,omitempty"`
	ResolutionSort                      string `json:"resolutionSort,omitempty"`
	InternalAnalysisSort                string `json:"internalAnalysisSort,omitempty"`
	EmailConnectorAllowReopenClosedFlag bool   `json:"emailConnectorAllowReopenClosedFlag,omitempty"`
	EmailConnectorReopenStatus          struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Sort int    `json:"sort,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"emailConnectorReopenStatus,omitempty"`
	EmailConnectorReopenResourcesFlag         bool   `json:"emailConnectorReopenResourcesFlag,omitempty"`
	EmailConnectorNewTicketNoMatchFlag        bool   `json:"emailConnectorNewTicketNoMatchFlag,omitempty"`
	EmailConnectorNeverReopenByDaysFlag       bool   `json:"emailConnectorNeverReopenByDaysFlag,omitempty"`
	EmailConnectorReopenDaysLimit             int    `json:"emailConnectorReopenDaysLimit,omitempty"`
	EmailConnectorNeverReopenByDaysClosedFlag bool   `json:"emailConnectorNeverReopenByDaysClosedFlag,omitempty"`
	EmailConnectorReopenDaysClosedLimit       int    `json:"emailConnectorReopenDaysClosedLimit,omitempty"`
	UseMemberDisplayNameFlag                  bool   `json:"useMemberDisplayNameFlag,omitempty"`
	SendToCCFlag                              bool   `json:"sendToCCFlag,omitempty"`
	AutoAssignTicketOwnerFlag                 bool   `json:"autoAssignTicketOwnerFlag,omitempty"`
	AutoAssignLimitFlag                       bool   `json:"autoAssignLimitFlag,omitempty"`
	AutoAssignLimitAmount                     int    `json:"autoAssignLimitAmount,omitempty"`
	ClosedLoopAllFlag                         bool   `json:"closedLoopAllFlag,omitempty"`
	PercentageCalculation                     string `json:"percentageCalculation,omitempty"`
	AllSort                                   string `json:"allSort,omitempty"`
	MarkFirstNoteIssueFlag                    bool   `json:"markFirstNoteIssueFlag,omitempty"`
	RestrictBoardByDefaultFlag                bool   `json:"restrictBoardByDefaultFlag,omitempty"`
	SendToBundledFlag                         bool   `json:"sendToBundledFlag,omitempty"`
	Info                                      struct {
		AdditionalProp1 string `json:"additionalProp1,omitempty"`
		AdditionalProp2 string `json:"additionalProp2,omitempty"`
		AdditionalProp3 string `json:"additionalProp3,omitempty"`
	} `json:"_info,omitempty"`
}
type Callback struct {
	Id                   int    `json:"id,omitempty"`
	Description          string `json:"description,omitempty"`
	Url                  string `json:"url"`
	ObjectId             int    `json:"objectId"`
	Type                 string `json:"type"`
	Level                string `json:"level"`
	MemberId             int    `json:"memberId,omitempty"`
	PayloadVersion       string `json:"payloadVersion,omitempty"`
	InactiveFlag         bool   `json:"inactiveFlag,omitempty"`
	IsSoapCallbackFlag   bool   `json:"isSoapCallbackFlag,omitempty"`
	IsSelfSuppressedFlag bool   `json:"isSelfSuppressedFlag,omitempty"`
	ConnectWiseID        string `json:"connectWiseID,omitempty"`
	Info                 struct {
		AdditionalProp1 string `json:"additionalProp1,omitempty"`
		AdditionalProp2 string `json:"additionalProp2,omitempty"`
		AdditionalProp3 string `json:"additionalProp3,omitempty"`
	} `json:"_info,omitempty"`
}

type Company struct {
	Id         int    `json:"id,omitempty"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Status     struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"status,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`
	Country      struct {
		Id         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"country,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	FaxNumber   string `json:"faxNumber,omitempty"`
	Website     string `json:"website,omitempty"`
	Territory   struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"territory,omitempty"`
	Market struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"market,omitempty"`
	AccountNumber  string `json:"accountNumber,omitempty"`
	DefaultContact struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"defaultContact,omitempty"`
	DateAcquired time.Time `json:"dateAcquired,omitempty"`
	SicCode      struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"sicCode,omitempty"`
	ParentCompany struct {
		Id         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"parentCompany,omitempty"`
	AnnualRevenue     int `json:"annualRevenue,omitempty"`
	NumberOfEmployees int `json:"numberOfEmployees,omitempty"`
	YearEstablished   int `json:"yearEstablished,omitempty"`
	RevenueYear       int `json:"revenueYear,omitempty"`
	OwnershipType     struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"ownershipType,omitempty"`
	TimeZoneSetup struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"timeZoneSetup,omitempty"`
	LeadSource      string `json:"leadSource,omitempty"`
	LeadFlag        bool   `json:"leadFlag,omitempty"`
	UnsubscribeFlag bool   `json:"unsubscribeFlag,omitempty"`
	Calendar        struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"calendar,omitempty"`
	UserDefinedField1  string `json:"userDefinedField1,omitempty"`
	UserDefinedField2  string `json:"userDefinedField2,omitempty"`
	UserDefinedField3  string `json:"userDefinedField3,omitempty"`
	UserDefinedField4  string `json:"userDefinedField4,omitempty"`
	UserDefinedField5  string `json:"userDefinedField5,omitempty"`
	UserDefinedField6  string `json:"userDefinedField6,omitempty"`
	UserDefinedField7  string `json:"userDefinedField7,omitempty"`
	UserDefinedField8  string `json:"userDefinedField8,omitempty"`
	UserDefinedField9  string `json:"userDefinedField9,omitempty"`
	UserDefinedField10 string `json:"userDefinedField10,omitempty"`
	VendorIdentifier   string `json:"vendorIdentifier,omitempty"`
	TaxIdentifier      string `json:"taxIdentifier,omitempty"`
	TaxCode            struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"taxCode,omitempty"`
	BillingTerms struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"billingTerms,omitempty"`
	InvoiceTemplate struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"invoiceTemplate,omitempty"`
	PricingSchedule struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"pricingSchedule,omitempty"`
	CompanyEntityType struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"companyEntityType,omitempty"`
	BillToCompany struct {
		Id         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"billToCompany,omitempty"`
	BillingSite struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"billingSite,omitempty"`
	BillingContact struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"billingContact,omitempty"`
	InvoiceDeliveryMethod struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"invoiceDeliveryMethod,omitempty"`
	InvoiceToEmailAddress string    `json:"invoiceToEmailAddress,omitempty"`
	InvoiceCCEmailAddress string    `json:"invoiceCCEmailAddress,omitempty"`
	DeletedFlag           bool      `json:"deletedFlag,omitempty"`
	DateDeleted           time.Time `json:"dateDeleted,omitempty"`
	DeletedBy             string    `json:"deletedBy,omitempty"`
	MobileGuid            string    `json:"mobileGuid,omitempty"`
	FacebookUrl           string    `json:"facebookUrl,omitempty"`
	TwitterUrl            string    `json:"twitterUrl,omitempty"`
	LinkedInUrl           string    `json:"linkedInUrl,omitempty"`
	Currency              struct {
		Id                      int    `json:"id,omitempty"`
		Symbol                  string `json:"symbol,omitempty"`
		CurrencyCode            string `json:"currencyCode,omitempty"`
		DecimalSeparator        string `json:"decimalSeparator,omitempty"`
		NumberOfDecimals        int    `json:"numberOfDecimals,omitempty"`
		ThousandsSeparator      string `json:"thousandsSeparator,omitempty"`
		NegativeParenthesesFlag bool   `json:"negativeParenthesesFlag,omitempty"`
		DisplaySymbolFlag       bool   `json:"displaySymbolFlag,omitempty"`
		CurrencyIdentifier      string `json:"currencyIdentifier,omitempty"`
		DisplayIdFlag           bool   `json:"displayIdFlag,omitempty"`
		RightAlign              bool   `json:"rightAlign,omitempty"`
		Name                    string `json:"name,omitempty"`
		Info                    struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"currency,omitempty"`
	TerritoryManager struct {
		Id            int    `json:"id,omitempty"`
		Identifier    string `json:"identifier,omitempty"`
		Name          string `json:"name,omitempty"`
		DailyCapacity int    `json:"dailyCapacity,omitempty"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"territoryManager,omitempty"`
	ResellerIdentifier string `json:"resellerIdentifier,omitempty"`
	IsVendorFlag       bool   `json:"isVendorFlag,omitempty"`
	Types              []struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"types,omitempty"`
	Site struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"site,omitempty"`
	IntegratorTags []string `json:"integratorTags,omitempty"`
	Info           struct {
		AdditionalProp1 string `json:"additionalProp1,omitempty"`
		AdditionalProp2 string `json:"additionalProp2,omitempty"`
		AdditionalProp3 string `json:"additionalProp3,omitempty"`
	} `json:"_info,omitempty"`
	CustomFields []struct {
		Id               int    `json:"id,omitempty"`
		Caption          string `json:"caption,omitempty"`
		Type             string `json:"type,omitempty"`
		EntryMethod      string `json:"entryMethod,omitempty"`
		NumberOfDecimals int    `json:"numberOfDecimals,omitempty"`
		Value            struct {
		} `json:"value,omitempty"`
		ConnectWiseId string `json:"connectWiseId,omitempty"`
	} `json:"customFields,omitempty"`
}
type Contact struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Company   struct {
		Id         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"company,omitempty"`
	Site struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"site,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`
	Country      struct {
		Id         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"country,omitempty"`
	Relationship struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"relationship,omitempty"`
	RelationshipOverride string `json:"relationshipOverride,omitempty"`
	Department           struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"department,omitempty"`
	InactiveFlag          bool   `json:"inactiveFlag,omitempty"`
	DefaultMergeContactId int    `json:"defaultMergeContactId,omitempty"`
	SecurityIdentifier    string `json:"securityIdentifier,omitempty"`
	ManagerContact        struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"managerContact,omitempty"`
	AssistantContact struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"assistantContact,omitempty"`
	Title                  string `json:"title,omitempty"`
	School                 string `json:"school,omitempty"`
	NickName               string `json:"nickName,omitempty"`
	MarriedFlag            bool   `json:"marriedFlag,omitempty"`
	ChildrenFlag           bool   `json:"childrenFlag,omitempty"`
	Children               string `json:"children,omitempty"`
	SignificantOther       string `json:"significantOther,omitempty"`
	PortalPassword         string `json:"portalPassword,omitempty"`
	PortalSecurityLevel    int    `json:"portalSecurityLevel,omitempty"`
	DisablePortalLoginFlag bool   `json:"disablePortalLoginFlag,omitempty"`
	UnsubscribeFlag        bool   `json:"unsubscribeFlag,omitempty"`
	Gender                 string `json:"gender,omitempty"`
	BirthDay               string `json:"birthDay,omitempty"`
	Anniversary            string `json:"anniversary,omitempty"`
	Presence               string `json:"presence,omitempty"`
	MobileGuid             string `json:"mobileGuid,omitempty"`
	FacebookUrl            string `json:"facebookUrl,omitempty"`
	TwitterUrl             string `json:"twitterUrl,omitempty"`
	LinkedInUrl            string `json:"linkedInUrl,omitempty"`
	DefaultPhoneType       string `json:"defaultPhoneType,omitempty"`
	DefaultPhoneNbr        string `json:"defaultPhoneNbr,omitempty"`
	DefaultPhoneExtension  string `json:"defaultPhoneExtension,omitempty"`
	DefaultBillingFlag     bool   `json:"defaultBillingFlag,omitempty"`
	DefaultFlag            bool   `json:"defaultFlag,omitempty"`
	UserDefinedField1      string `json:"userDefinedField1,omitempty"`
	UserDefinedField2      string `json:"userDefinedField2,omitempty"`
	UserDefinedField3      string `json:"userDefinedField3,omitempty"`
	UserDefinedField4      string `json:"userDefinedField4,omitempty"`
	UserDefinedField5      string `json:"userDefinedField5,omitempty"`
	UserDefinedField6      string `json:"userDefinedField6,omitempty"`
	UserDefinedField7      string `json:"userDefinedField7,omitempty"`
	UserDefinedField8      string `json:"userDefinedField8,omitempty"`
	UserDefinedField9      string `json:"userDefinedField9,omitempty"`
	UserDefinedField10     string `json:"userDefinedField10,omitempty"`
	CompanyLocation        struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"companyLocation,omitempty"`
	CommunicationItems []struct {
		Id   int `json:"id,omitempty"`
		Type struct {
			Id   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
			Info struct {
				AdditionalProp1 string `json:"additionalProp1,omitempty"`
				AdditionalProp2 string `json:"additionalProp2,omitempty"`
				AdditionalProp3 string `json:"additionalProp3,omitempty"`
			} `json:"_info,omitempty"`
		} `json:"type,omitempty"`
		Value             string `json:"value,omitempty"`
		Extension         string `json:"extension,omitempty"`
		DefaultFlag       bool   `json:"defaultFlag,omitempty"`
		Domain            string `json:"domain,omitempty"`
		CommunicationType string `json:"communicationType,omitempty"`
	} `json:"communicationItems,omitempty"`
	Types []struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"types,omitempty"`
	IntegratorTags []string `json:"integratorTags,omitempty"`
	CustomFields   []struct {
		Id               int    `json:"id,omitempty"`
		Caption          string `json:"caption,omitempty"`
		Type             string `json:"type,omitempty"`
		EntryMethod      string `json:"entryMethod,omitempty"`
		NumberOfDecimals int    `json:"numberOfDecimals,omitempty"`
		Value            struct {
		} `json:"value,omitempty"`
		ConnectWiseId string `json:"connectWiseId,omitempty"`
	} `json:"customFields,omitempty"`
	Photo struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"photo,omitempty"`
	IgnoreDuplicates bool `json:"ignoreDuplicates,omitempty"`
	Info             struct {
		AdditionalProp1 string `json:"additionalProp1,omitempty"`
		AdditionalProp2 string `json:"additionalProp2,omitempty"`
		AdditionalProp3 string `json:"additionalProp3,omitempty"`
	} `json:"_info,omitempty"`
	TypeIds []int `json:"typeIds,omitempty"`
}
type Ticket struct {
	Id         int    `json:"id,omitempty"`
	Summary    string `json:"summary"`
	RecordType string `json:"recordType,omitempty"`
	Board      struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"board,omitempty"`
	Status struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Sort int    `json:"sort,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"status,omitempty"`
	WorkRole struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"workRole,omitempty"`
	WorkType struct {
		Id              int    `json:"id,omitempty"`
		Name            string `json:"name,omitempty"`
		UtilizationFlag bool   `json:"utilizationFlag,omitempty"`
		Info            struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"workType,omitempty"`
	Company struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"company"`
	Site struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"site,omitempty"`
	SiteName        string `json:"siteName,omitempty"`
	AddressLine1    string `json:"addressLine1,omitempty"`
	AddressLine2    string `json:"addressLine2,omitempty"`
	City            string `json:"city,omitempty"`
	StateIdentifier string `json:"stateIdentifier,omitempty"`
	Zip             string `json:"zip,omitempty"`
	Country         struct {
		Id         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"country,omitempty"`
	Contact struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"contact,omitempty"`
	ContactName           string `json:"contactName,omitempty"`
	ContactPhoneNumber    string `json:"contactPhoneNumber,omitempty"`
	ContactPhoneExtension string `json:"contactPhoneExtension,omitempty"`
	ContactEmailAddress   string `json:"contactEmailAddress,omitempty"`
	Type                  struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"type,omitempty"`
	SubType struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"subType,omitempty"`
	Item struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"item,omitempty"`
	Team struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"team,omitempty"`
	Owner struct {
		Id            int    `json:"id,omitempty"`
		Identifier    string `json:"identifier,omitempty"`
		Name          string `json:"name,omitempty"`
		DailyCapacity int    `json:"dailyCapacity,omitempty"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"owner,omitempty"`
	Priority struct {
		Id    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Sort  int    `json:"sort,omitempty"`
		Level string `json:"level,omitempty"`
		Info  struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"priority,omitempty"`
	ServiceLocation struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"serviceLocation,omitempty"`
	Source struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"source,omitempty"`
	RequiredDate time.Time `json:"requiredDate,omitempty"`
	BudgetHours  int       `json:"budgetHours,omitempty"`
	Opportunity  struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"opportunity,omitempty"`
	Agreement struct {
		Id             int    `json:"id,omitempty"`
		Name           string `json:"name,omitempty"`
		Type           string `json:"type,omitempty"`
		ChargeFirmFlag bool   `json:"chargeFirmFlag,omitempty"`
		Info           struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"agreement,omitempty"`
	AgreementType              string    `json:"agreementType,omitempty"`
	Severity                   string    `json:"severity,omitempty"`
	Impact                     string    `json:"impact,omitempty"`
	ExternalXRef               string    `json:"externalXRef,omitempty"`
	PoNumber                   string    `json:"poNumber,omitempty"`
	KnowledgeBaseCategoryId    int       `json:"knowledgeBaseCategoryId,omitempty"`
	KnowledgeBaseSubCategoryId int       `json:"knowledgeBaseSubCategoryId,omitempty"`
	AllowAllClientsPortalView  bool      `json:"allowAllClientsPortalView,omitempty"`
	CustomerUpdatedFlag        bool      `json:"customerUpdatedFlag,omitempty"`
	AutomaticEmailContactFlag  bool      `json:"automaticEmailContactFlag,omitempty"`
	AutomaticEmailResourceFlag bool      `json:"automaticEmailResourceFlag,omitempty"`
	AutomaticEmailCcFlag       bool      `json:"automaticEmailCcFlag,omitempty"`
	AutomaticEmailCc           string    `json:"automaticEmailCc,omitempty"`
	InitialDescription         string    `json:"initialDescription,omitempty"`
	InitialInternalAnalysis    string    `json:"initialInternalAnalysis,omitempty"`
	InitialResolution          string    `json:"initialResolution,omitempty"`
	InitialDescriptionFrom     string    `json:"initialDescriptionFrom,omitempty"`
	ContactEmailLookup         string    `json:"contactEmailLookup,omitempty"`
	ProcessNotifications       bool      `json:"processNotifications,omitempty"`
	SkipCallback               bool      `json:"skipCallback,omitempty"`
	ClosedDate                 string    `json:"closedDate,omitempty"`
	ClosedBy                   string    `json:"closedBy,omitempty"`
	ClosedFlag                 bool      `json:"closedFlag,omitempty"`
	ActualHours                int       `json:"actualHours,omitempty"`
	Approved                   bool      `json:"approved,omitempty"`
	EstimatedExpenseCost       int       `json:"estimatedExpenseCost,omitempty"`
	EstimatedExpenseRevenue    int       `json:"estimatedExpenseRevenue,omitempty"`
	EstimatedProductCost       int       `json:"estimatedProductCost,omitempty"`
	EstimatedProductRevenue    int       `json:"estimatedProductRevenue,omitempty"`
	EstimatedTimeCost          int       `json:"estimatedTimeCost,omitempty"`
	EstimatedTimeRevenue       int       `json:"estimatedTimeRevenue,omitempty"`
	BillingMethod              string    `json:"billingMethod,omitempty"`
	BillingAmount              int       `json:"billingAmount,omitempty"`
	HourlyRate                 int       `json:"hourlyRate,omitempty"`
	SubBillingMethod           string    `json:"subBillingMethod,omitempty"`
	SubBillingAmount           int       `json:"subBillingAmount,omitempty"`
	SubDateAccepted            string    `json:"subDateAccepted,omitempty"`
	DateResolved               string    `json:"dateResolved,omitempty"`
	DateResplan                string    `json:"dateResplan,omitempty"`
	DateResponded              string    `json:"dateResponded,omitempty"`
	ResolveMinutes             int       `json:"resolveMinutes,omitempty"`
	ResPlanMinutes             int       `json:"resPlanMinutes,omitempty"`
	RespondMinutes             int       `json:"respondMinutes,omitempty"`
	IsInSla                    bool      `json:"isInSla,omitempty"`
	KnowledgeBaseLinkId        int       `json:"knowledgeBaseLinkId,omitempty"`
	Resources                  string    `json:"resources,omitempty"`
	ParentTicketId             int       `json:"parentTicketId,omitempty"`
	HasChildTicket             bool      `json:"hasChildTicket,omitempty"`
	HasMergedChildTicketFlag   bool      `json:"hasMergedChildTicketFlag,omitempty"`
	KnowledgeBaseLinkType      string    `json:"knowledgeBaseLinkType,omitempty"`
	BillTime                   string    `json:"billTime,omitempty"`
	BillExpenses               string    `json:"billExpenses,omitempty"`
	BillProducts               string    `json:"billProducts,omitempty"`
	PredecessorType            string    `json:"predecessorType,omitempty"`
	PredecessorId              int       `json:"predecessorId,omitempty"`
	PredecessorClosedFlag      bool      `json:"predecessorClosedFlag,omitempty"`
	LagDays                    int       `json:"lagDays,omitempty"`
	LagNonworkingDaysFlag      bool      `json:"lagNonworkingDaysFlag,omitempty"`
	EstimatedStartDate         time.Time `json:"estimatedStartDate,omitempty"`
	Duration                   int       `json:"duration,omitempty"`
	Location                   struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"location,omitempty"`
	Department struct {
		Id         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"department,omitempty"`
	MobileGuid string `json:"mobileGuid,omitempty"`
	Sla        struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"sla,omitempty"`
	SlaStatus            string `json:"slaStatus,omitempty"`
	RequestForChangeFlag bool   `json:"requestForChangeFlag,omitempty"`
	Currency             struct {
		Id                      int    `json:"id,omitempty"`
		Symbol                  string `json:"symbol,omitempty"`
		CurrencyCode            string `json:"currencyCode,omitempty"`
		DecimalSeparator        string `json:"decimalSeparator,omitempty"`
		NumberOfDecimals        int    `json:"numberOfDecimals,omitempty"`
		ThousandsSeparator      string `json:"thousandsSeparator,omitempty"`
		NegativeParenthesesFlag bool   `json:"negativeParenthesesFlag,omitempty"`
		DisplaySymbolFlag       bool   `json:"displaySymbolFlag,omitempty"`
		CurrencyIdentifier      string `json:"currencyIdentifier,omitempty"`
		DisplayIdFlag           bool   `json:"displayIdFlag,omitempty"`
		RightAlign              bool   `json:"rightAlign,omitempty"`
		Name                    string `json:"name,omitempty"`
		Info                    struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"currency,omitempty"`
	MergedParentTicket struct {
		Id      int    `json:"id,omitempty"`
		Summary string `json:"summary,omitempty"`
		Info    struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"mergedParentTicket,omitempty"`
	IntegratorTags []string `json:"integratorTags,omitempty"`
	Info           struct {
		AdditionalProp1 string `json:"additionalProp1,omitempty"`
		AdditionalProp2 string `json:"additionalProp2,omitempty"`
		AdditionalProp3 string `json:"additionalProp3,omitempty"`
	} `json:"_info,omitempty"`
	EscalationStartDateUTC  string `json:"escalationStartDateUTC,omitempty"`
	EscalationLevel         int    `json:"escalationLevel,omitempty"`
	MinutesBeforeWaiting    int    `json:"minutesBeforeWaiting,omitempty"`
	RespondedSkippedMinutes int    `json:"respondedSkippedMinutes,omitempty"`
	ResplanSkippedMinutes   int    `json:"resplanSkippedMinutes,omitempty"`
	RespondedHours          int    `json:"respondedHours,omitempty"`
	RespondedBy             string `json:"respondedBy,omitempty"`
	ResplanHours            int    `json:"resplanHours,omitempty"`
	ResplanBy               string `json:"resplanBy,omitempty"`
	ResolutionHours         int    `json:"resolutionHours,omitempty"`
	ResolvedBy              string `json:"resolvedBy,omitempty"`
	MinutesWaiting          int    `json:"minutesWaiting,omitempty"`
	CustomFields            []struct {
		Id               int    `json:"id,omitempty"`
		Caption          string `json:"caption,omitempty"`
		Type             string `json:"type,omitempty"`
		EntryMethod      string `json:"entryMethod,omitempty"`
		NumberOfDecimals int    `json:"numberOfDecimals,omitempty"`
		Value            struct {
		} `json:"value,omitempty"`
		ConnectWiseId string `json:"connectWiseId,omitempty"`
	} `json:"customFields,omitempty"`
}
type ServiceTicketNote struct {
	Id                    int    `json:"id,omitempty"`
	TicketId              int    `json:"ticketId,omitempty"`
	Text                  string `json:"text,omitempty"`
	DetailDescriptionFlag bool   `json:"detailDescriptionFlag,omitempty"`
	InternalAnalysisFlag  bool   `json:"internalAnalysisFlag,omitempty"`
	ResolutionFlag        bool   `json:"resolutionFlag,omitempty"`
	IssueFlag             bool   `json:"issueFlag,omitempty"`
	Member                struct {
		Id            int    `json:"id,omitempty"`
		Identifier    string `json:"identifier,omitempty"`
		Name          string `json:"name,omitempty"`
		DailyCapacity int    `json:"dailyCapacity,omitempty"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"member,omitempty"`
	Contact struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"contact,omitempty"`
	CustomerUpdatedFlag  bool `json:"customerUpdatedFlag,omitempty"`
	ProcessNotifications bool `json:"processNotifications,omitempty"`
	InternalFlag         bool `json:"internalFlag,omitempty"`
	ExternalFlag         bool `json:"externalFlag,omitempty"`
	Info                 struct {
		AdditionalProp1 string `json:"additionalProp1,omitempty"`
		AdditionalProp2 string `json:"additionalProp2,omitempty"`
		AdditionalProp3 string `json:"additionalProp3,omitempty"`
	} `json:"_info,omitempty"`
}

type ServiceTicketNoteAll struct {
	Id       int    `json:"id,omitempty"`
	NoteType string `json:"noteType,omitempty"`
	Ticket   struct {
		Id      int    `json:"id,omitempty"`
		Summary string `json:"summary,omitempty"`
		Info    struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"ticket,omitempty"`
	Text                  string `json:"text,omitempty"`
	IsMarkdownFlag        bool   `json:"isMarkdownFlag,omitempty"`
	DetailDescriptionFlag bool   `json:"detailDescriptionFlag,omitempty"`
	InternalAnalysisFlag  bool   `json:"internalAnalysisFlag,omitempty"`
	ResolutionFlag        bool   `json:"resolutionFlag,omitempty"`
	TimeStart             string `json:"timeStart,omitempty"`
	TimeEnd               string `json:"timeEnd,omitempty"`
	BundledFlag           bool   `json:"bundledFlag,omitempty"`
	MergedFlag            bool   `json:"mergedFlag,omitempty"`
	IssueFlag             bool   `json:"issueFlag,omitempty"`
	OriginalAuthor        string `json:"originalAuthor,omitempty"`
	CreatedByParentFlag   bool   `json:"createdByParentFlag,omitempty"`
	Member                struct {
		Id            int    `json:"id,omitempty"`
		Identifier    string `json:"identifier,omitempty"`
		Name          string `json:"name,omitempty"`
		DailyCapacity int    `json:"dailyCapacity,omitempty"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"member,omitempty"`
	Contact struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1,omitempty"`
			AdditionalProp2 string `json:"additionalProp2,omitempty"`
			AdditionalProp3 string `json:"additionalProp3,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"contact,omitempty"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1,omitempty"`
		AdditionalProp2 string `json:"additionalProp2,omitempty"`
		AdditionalProp3 string `json:"additionalProp3,omitempty"`
	} `json:"_info,omitempty"`
}
