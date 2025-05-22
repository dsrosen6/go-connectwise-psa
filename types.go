package connectwise

import "time"

type Board struct {
	Id                                        int                   `json:"id,omitempty"`
	Name                                      string                `json:"name"`
	Location                                  SystemLocation        `json:"location"`
	Department                                SystemDepartment      `json:"department"`
	InactiveFlag                              bool                  `json:"inactiveFlag,omitempty"`
	SignOffTemplate                           SignOffTemplate       `json:"signOffTemplate,omitempty"`
	SendToContactFlag                         bool                  `json:"sendToContactFlag,omitempty"`
	ContactTemplate                           EmailTemplate         `json:"contactTemplate,omitempty"`
	SendToResourceFlag                        bool                  `json:"sendToResourceFlag,omitempty"`
	ResourceTemplate                          EmailTemplate         `json:"resourceTemplate,omitempty"`
	ProjectFlag                               bool                  `json:"projectFlag,omitempty"`
	ShowDependenciesFlag                      bool                  `json:"showDependenciesFlag,omitempty"`
	ShowEstimatesFlag                         bool                  `json:"showEstimatesFlag,omitempty"`
	BoardIcon                                 Document              `json:"boardIcon,omitempty"`
	BillTicketsAfterClosedFlag                bool                  `json:"billTicketsAfterClosedFlag,omitempty"`
	BillTicketSeparatelyFlag                  bool                  `json:"billTicketSeparatelyFlag,omitempty"`
	BillUnapprovedTimeExpenseFlag             bool                  `json:"billUnapprovedTimeExpenseFlag,omitempty"`
	OverrideBillingSetupFlag                  bool                  `json:"overrideBillingSetupFlag,omitempty"`
	DispatchMember                            Member                `json:"dispatchMember,omitempty"`
	ServiceManagerMember                      Member                `json:"serviceManagerMember,omitempty"`
	DutyManagerMember                         Member                `json:"dutyManagerMember,omitempty"`
	OncallMember                              Member                `json:"oncallMember,omitempty"`
	WorkRole                                  WorkRole              `json:"workRole,omitempty"`
	WorkType                                  WorkType              `json:"workType,omitempty"`
	BillTime                                  string                `json:"billTime,omitempty"`
	BillExpense                               string                `json:"billExpense,omitempty"`
	BillProduct                               string                `json:"billProduct,omitempty"`
	AutoCloseStatus                           Status                `json:"autoCloseStatus,omitempty"`
	AutoAssignNewTicketsFlag                  bool                  `json:"autoAssignNewTicketsFlag,omitempty"`
	AutoAssignNewECTicketsFlag                bool                  `json:"autoAssignNewECTicketsFlag,omitempty"`
	AutoAssignNewPortalTicketsFlag            bool                  `json:"autoAssignNewPortalTicketsFlag,omitempty"`
	DiscussionsLockedFlag                     bool                  `json:"discussionsLockedFlag,omitempty"`
	TimeEntryLockedFlag                       bool                  `json:"timeEntryLockedFlag,omitempty"`
	NotifyEmailFrom                           string                `json:"notifyEmailFrom,omitempty"`
	NotifyEmailFromName                       string                `json:"notifyEmailFromName,omitempty"`
	ClosedLoopDiscussionsFlag                 bool                  `json:"closedLoopDiscussionsFlag,omitempty"`
	ClosedLoopResolutionFlag                  bool                  `json:"closedLoopResolutionFlag,omitempty"`
	ClosedLoopInternalAnalysisFlag            bool                  `json:"closedLoopInternalAnalysisFlag,omitempty"`
	TimeEntryDiscussionFlag                   bool                  `json:"timeEntryDiscussionFlag,omitempty"`
	TimeEntryResolutionFlag                   bool                  `json:"timeEntryResolutionFlag,omitempty"`
	TimeEntryInternalAnalysisFlag             bool                  `json:"timeEntryInternalAnalysisFlag,omitempty"`
	ProblemSort                               string                `json:"problemSort,omitempty"`
	ResolutionSort                            string                `json:"resolutionSort,omitempty"`
	InternalAnalysisSort                      string                `json:"internalAnalysisSort,omitempty"`
	EmailConnectorAllowReopenClosedFlag       bool                  `json:"emailConnectorAllowReopenClosedFlag,omitempty"`
	EmailConnectorReopenStatus                ConnectorReopenStatus `json:"emailConnectorReopenStatus,omitempty"`
	EmailConnectorReopenResourcesFlag         bool                  `json:"emailConnectorReopenResourcesFlag,omitempty"`
	EmailConnectorNewTicketNoMatchFlag        bool                  `json:"emailConnectorNewTicketNoMatchFlag,omitempty"`
	EmailConnectorNeverReopenByDaysFlag       bool                  `json:"emailConnectorNeverReopenByDaysFlag,omitempty"`
	EmailConnectorReopenDaysLimit             int                   `json:"emailConnectorReopenDaysLimit,omitempty"`
	EmailConnectorNeverReopenByDaysClosedFlag bool                  `json:"emailConnectorNeverReopenByDaysClosedFlag,omitempty"`
	EmailConnectorReopenDaysClosedLimit       int                   `json:"emailConnectorReopenDaysClosedLimit,omitempty"`
	UseMemberDisplayNameFlag                  bool                  `json:"useMemberDisplayNameFlag,omitempty"`
	SendToCCFlag                              bool                  `json:"sendToCCFlag,omitempty"`
	AutoAssignTicketOwnerFlag                 bool                  `json:"autoAssignTicketOwnerFlag,omitempty"`
	AutoAssignLimitFlag                       bool                  `json:"autoAssignLimitFlag,omitempty"`
	AutoAssignLimitAmount                     int                   `json:"autoAssignLimitAmount,omitempty"`
	ClosedLoopAllFlag                         bool                  `json:"closedLoopAllFlag,omitempty"`
	PercentageCalculation                     string                `json:"percentageCalculation,omitempty"`
	AllSort                                   string                `json:"allSort,omitempty"`
	MarkFirstNoteIssueFlag                    bool                  `json:"markFirstNoteIssueFlag,omitempty"`
	RestrictBoardByDefaultFlag                bool                  `json:"restrictBoardByDefaultFlag,omitempty"`
	SendToBundledFlag                         bool                  `json:"sendToBundledFlag,omitempty"`
	Info                                      Info                  `json:"_info,omitempty"`
}

type Company struct {
	Id                    int             `json:"id,omitempty"`
	Identifier            string          `json:"identifier"`
	Name                  string          `json:"name"`
	Status                Status          `json:"status,omitempty"`
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

type Ticket struct {
	Id                         int              `json:"id,omitempty"`
	Summary                    string           `json:"summary"`
	RecordType                 string           `json:"recordType,omitempty"`
	Board                      Board            `json:"board,omitempty"`
	Status                     Status           `json:"status,omitempty"`
	WorkRole                   WorkRole         `json:"workRole,omitempty"`
	WorkType                   WorkType         `json:"workType,omitempty"`
	Company                    Company          `json:"company"`
	Site                       Site             `json:"site,omitempty"`
	SiteName                   string           `json:"siteName,omitempty"`
	AddressLine1               string           `json:"addressLine1,omitempty"`
	AddressLine2               string           `json:"addressLine2,omitempty"`
	City                       string           `json:"city,omitempty"`
	StateIdentifier            string           `json:"stateIdentifier,omitempty"`
	Zip                        string           `json:"zip,omitempty"`
	Country                    Country          `json:"country,omitempty"`
	Contact                    Contact          `json:"contact,omitempty"`
	ContactName                string           `json:"contactName,omitempty"`
	ContactPhoneNumber         string           `json:"contactPhoneNumber,omitempty"`
	ContactPhoneExtension      string           `json:"contactPhoneExtension,omitempty"`
	ContactEmailAddress        string           `json:"contactEmailAddress,omitempty"`
	Type                       ServiceType      `json:"type,omitempty"`
	SubType                    ServiceSubType   `json:"subType,omitempty"`
	Item                       ServiceItem      `json:"item,omitempty"`
	Team                       ServiceTeam      `json:"team,omitempty"`
	Owner                      Member           `json:"owner,omitempty"`
	Priority                   Priority         `json:"priority,omitempty"`
	ServiceLocation            ServiceLocation  `json:"serviceLocation,omitempty"`
	Source                     ServiceSource    `json:"source,omitempty"`
	RequiredDate               time.Time        `json:"requiredDate,omitempty"`
	BudgetHours                int              `json:"budgetHours,omitempty"`
	Opportunity                Opportunity      `json:"opportunity,omitempty"`
	Agreement                  Agreement        `json:"agreement,omitempty"`
	AgreementType              string           `json:"agreementType,omitempty"`
	Severity                   string           `json:"severity,omitempty"`
	Impact                     string           `json:"impact,omitempty"`
	ExternalXRef               string           `json:"externalXRef,omitempty"`
	PoNumber                   string           `json:"poNumber,omitempty"`
	KnowledgeBaseCategoryId    int              `json:"knowledgeBaseCategoryId,omitempty"`
	KnowledgeBaseSubCategoryId int              `json:"knowledgeBaseSubCategoryId,omitempty"`
	AllowAllClientsPortalView  bool             `json:"allowAllClientsPortalView,omitempty"`
	CustomerUpdatedFlag        bool             `json:"customerUpdatedFlag,omitempty"`
	AutomaticEmailContactFlag  bool             `json:"automaticEmailContactFlag,omitempty"`
	AutomaticEmailResourceFlag bool             `json:"automaticEmailResourceFlag,omitempty"`
	AutomaticEmailCcFlag       bool             `json:"automaticEmailCcFlag,omitempty"`
	AutomaticEmailCc           string           `json:"automaticEmailCc,omitempty"`
	InitialDescription         string           `json:"initialDescription,omitempty"`
	InitialInternalAnalysis    string           `json:"initialInternalAnalysis,omitempty"`
	InitialResolution          string           `json:"initialResolution,omitempty"`
	InitialDescriptionFrom     string           `json:"initialDescriptionFrom,omitempty"`
	ContactEmailLookup         string           `json:"contactEmailLookup,omitempty"`
	ProcessNotifications       bool             `json:"processNotifications,omitempty"`
	SkipCallback               bool             `json:"skipCallback,omitempty"`
	ClosedDate                 string           `json:"closedDate,omitempty"`
	ClosedBy                   string           `json:"closedBy,omitempty"`
	ClosedFlag                 bool             `json:"closedFlag,omitempty"`
	ActualHours                int              `json:"actualHours,omitempty"`
	Approved                   bool             `json:"approved,omitempty"`
	EstimatedExpenseCost       int              `json:"estimatedExpenseCost,omitempty"`
	EstimatedExpenseRevenue    int              `json:"estimatedExpenseRevenue,omitempty"`
	EstimatedProductCost       int              `json:"estimatedProductCost,omitempty"`
	EstimatedProductRevenue    int              `json:"estimatedProductRevenue,omitempty"`
	EstimatedTimeCost          int              `json:"estimatedTimeCost,omitempty"`
	EstimatedTimeRevenue       int              `json:"estimatedTimeRevenue,omitempty"`
	BillingMethod              string           `json:"billingMethod,omitempty"`
	BillingAmount              int              `json:"billingAmount,omitempty"`
	HourlyRate                 int              `json:"hourlyRate,omitempty"`
	SubBillingMethod           string           `json:"subBillingMethod,omitempty"`
	SubBillingAmount           int              `json:"subBillingAmount,omitempty"`
	SubDateAccepted            string           `json:"subDateAccepted,omitempty"`
	DateResolved               string           `json:"dateResolved,omitempty"`
	DateResplan                string           `json:"dateResplan,omitempty"`
	DateResponded              string           `json:"dateResponded,omitempty"`
	ResolveMinutes             int              `json:"resolveMinutes,omitempty"`
	ResPlanMinutes             int              `json:"resPlanMinutes,omitempty"`
	RespondMinutes             int              `json:"respondMinutes,omitempty"`
	IsInSla                    bool             `json:"isInSla,omitempty"`
	KnowledgeBaseLinkId        int              `json:"knowledgeBaseLinkId,omitempty"`
	Resources                  string           `json:"resources,omitempty"`
	ParentTicketId             int              `json:"parentTicketId,omitempty"`
	HasChildTicket             bool             `json:"hasChildTicket,omitempty"`
	HasMergedChildTicketFlag   bool             `json:"hasMergedChildTicketFlag,omitempty"`
	KnowledgeBaseLinkType      string           `json:"knowledgeBaseLinkType,omitempty"`
	BillTime                   string           `json:"billTime,omitempty"`
	BillExpenses               string           `json:"billExpenses,omitempty"`
	BillProducts               string           `json:"billProducts,omitempty"`
	PredecessorType            string           `json:"predecessorType,omitempty"`
	PredecessorId              int              `json:"predecessorId,omitempty"`
	PredecessorClosedFlag      bool             `json:"predecessorClosedFlag,omitempty"`
	LagDays                    int              `json:"lagDays,omitempty"`
	LagNonworkingDaysFlag      bool             `json:"lagNonworkingDaysFlag,omitempty"`
	EstimatedStartDate         time.Time        `json:"estimatedStartDate,omitempty"`
	Duration                   int              `json:"duration,omitempty"`
	Location                   SystemLocation   `json:"location,omitempty"`
	Department                 SystemDepartment `json:"department,omitempty"`
	MobileGuid                 string           `json:"mobileGuid,omitempty"`
	Sla                        Sla              `json:"sla,omitempty"`
	SlaStatus                  string           `json:"slaStatus,omitempty"`
	RequestForChangeFlag       bool             `json:"requestForChangeFlag,omitempty"`
	Currency                   Currency         `json:"currency,omitempty"`
	MergedParentTicket         *Ticket          `json:"mergedParentTicket,omitempty"`
	IntegratorTags             []string         `json:"integratorTags,omitempty"`
	Info                       Info             `json:"_info,omitempty"`
	EscalationStartDateUTC     string           `json:"escalationStartDateUTC,omitempty"`
	EscalationLevel            int              `json:"escalationLevel,omitempty"`
	MinutesBeforeWaiting       int              `json:"minutesBeforeWaiting,omitempty"`
	RespondedSkippedMinutes    int              `json:"respondedSkippedMinutes,omitempty"`
	ResplanSkippedMinutes      int              `json:"resplanSkippedMinutes,omitempty"`
	RespondedHours             int              `json:"respondedHours,omitempty"`
	RespondedBy                string           `json:"respondedBy,omitempty"`
	ResplanHours               int              `json:"resplanHours,omitempty"`
	ResplanBy                  string           `json:"resplanBy,omitempty"`
	ResolutionHours            int              `json:"resolutionHours,omitempty"`
	ResolvedBy                 string           `json:"resolvedBy,omitempty"`
	MinutesWaiting             int              `json:"minutesWaiting,omitempty"`
	CustomFields               []CustomField    `json:"customFields,omitempty"`
}

type ServiceType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type ServiceSubType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type ServiceItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type ServiceTeam struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Priority struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Sort  int    `json:"sort"`
	Level string `json:"level"`
	Info  struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type ServiceLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type ServiceSource struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Opportunity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

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

type SystemLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Sla struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Country struct {
	Id         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Info       struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Market struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Contact struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type SicCode struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type OwnershipType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type TimeZone struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Calendar struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type TaxCode struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type BillingTerms struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type InvoiceTemplate struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type PricingSchedule struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type EntityType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Site struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type BillingDelivery struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
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
	Info                    struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type CompanyType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type CustomField struct {
	Id               int    `json:"id"`
	Caption          string `json:"caption"`
	Type             string `json:"type"`
	EntryMethod      string `json:"entryMethod"`
	NumberOfDecimals int    `json:"numberOfDecimals"`
	Value            struct {
	} `json:"value"`
	ConnectWiseId string `json:"connectWiseId"`
}

type SystemDepartment struct {
	Id         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Info       struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type SignOffTemplate struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type EmailTemplate struct {
	Id         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Info       struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Document struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Member struct {
	Id            int    `json:"id"`
	Identifier    string `json:"identifier"`
	Name          string `json:"name"`
	DailyCapacity int    `json:"dailyCapacity"`
	Info          struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type WorkRole struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type WorkType struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	UtilizationFlag bool   `json:"utilizationFlag"`
	Info            struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type ConnectorReopenStatus struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Info struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
