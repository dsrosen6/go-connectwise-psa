package connectwise

import "time"

type Board struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"location"`
	Department struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"department"`
	InactiveFlag    bool `json:"inactiveFlag"`
	SignOffTemplate struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"signOffTemplate"`
	SendToContactFlag bool `json:"sendToContactFlag"`
	ContactTemplate   struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"contactTemplate"`
	SendToResourceFlag bool `json:"sendToResourceFlag"`
	ResourceTemplate   struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"resourceTemplate"`
	ProjectFlag          bool `json:"projectFlag"`
	ShowDependenciesFlag bool `json:"showDependenciesFlag"`
	ShowEstimatesFlag    bool `json:"showEstimatesFlag"`
	BoardIcon            struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"boardIcon"`
	BillTicketsAfterClosedFlag    bool `json:"billTicketsAfterClosedFlag"`
	BillTicketSeparatelyFlag      bool `json:"billTicketSeparatelyFlag"`
	BillUnapprovedTimeExpenseFlag bool `json:"billUnapprovedTimeExpenseFlag"`
	OverrideBillingSetupFlag      bool `json:"overrideBillingSetupFlag"`
	DispatchMember                struct {
		Id            int    `json:"id"`
		Identifier    string `json:"identifier"`
		Name          string `json:"name"`
		DailyCapacity int    `json:"dailyCapacity"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"dispatchMember"`
	ServiceManagerMember struct {
		Id            int    `json:"id"`
		Identifier    string `json:"identifier"`
		Name          string `json:"name"`
		DailyCapacity int    `json:"dailyCapacity"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"serviceManagerMember"`
	DutyManagerMember struct {
		Id            int    `json:"id"`
		Identifier    string `json:"identifier"`
		Name          string `json:"name"`
		DailyCapacity int    `json:"dailyCapacity"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"dutyManagerMember"`
	OncallMember struct {
		Id            int    `json:"id"`
		Identifier    string `json:"identifier"`
		Name          string `json:"name"`
		DailyCapacity int    `json:"dailyCapacity"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"oncallMember"`
	WorkRole struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"workRole"`
	WorkType struct {
		Id              int    `json:"id"`
		Name            string `json:"name"`
		UtilizationFlag bool   `json:"utilizationFlag"`
		Info            struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"workType"`
	BillTime        string `json:"billTime"`
	BillExpense     string `json:"billExpense"`
	BillProduct     string `json:"billProduct"`
	AutoCloseStatus struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Sort int    `json:"sort"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"autoCloseStatus"`
	AutoAssignNewTicketsFlag            bool   `json:"autoAssignNewTicketsFlag"`
	AutoAssignNewECTicketsFlag          bool   `json:"autoAssignNewECTicketsFlag"`
	AutoAssignNewPortalTicketsFlag      bool   `json:"autoAssignNewPortalTicketsFlag"`
	DiscussionsLockedFlag               bool   `json:"discussionsLockedFlag"`
	TimeEntryLockedFlag                 bool   `json:"timeEntryLockedFlag"`
	NotifyEmailFrom                     string `json:"notifyEmailFrom"`
	NotifyEmailFromName                 string `json:"notifyEmailFromName"`
	ClosedLoopDiscussionsFlag           bool   `json:"closedLoopDiscussionsFlag"`
	ClosedLoopResolutionFlag            bool   `json:"closedLoopResolutionFlag"`
	ClosedLoopInternalAnalysisFlag      bool   `json:"closedLoopInternalAnalysisFlag"`
	TimeEntryDiscussionFlag             bool   `json:"timeEntryDiscussionFlag"`
	TimeEntryResolutionFlag             bool   `json:"timeEntryResolutionFlag"`
	TimeEntryInternalAnalysisFlag       bool   `json:"timeEntryInternalAnalysisFlag"`
	ProblemSort                         string `json:"problemSort"`
	ResolutionSort                      string `json:"resolutionSort"`
	InternalAnalysisSort                string `json:"internalAnalysisSort"`
	EmailConnectorAllowReopenClosedFlag bool   `json:"emailConnectorAllowReopenClosedFlag"`
	EmailConnectorReopenStatus          struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Sort int    `json:"sort"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"emailConnectorReopenStatus"`
	EmailConnectorReopenResourcesFlag         bool   `json:"emailConnectorReopenResourcesFlag"`
	EmailConnectorNewTicketNoMatchFlag        bool   `json:"emailConnectorNewTicketNoMatchFlag"`
	EmailConnectorNeverReopenByDaysFlag       bool   `json:"emailConnectorNeverReopenByDaysFlag"`
	EmailConnectorReopenDaysLimit             int    `json:"emailConnectorReopenDaysLimit"`
	EmailConnectorNeverReopenByDaysClosedFlag bool   `json:"emailConnectorNeverReopenByDaysClosedFlag"`
	EmailConnectorReopenDaysClosedLimit       int    `json:"emailConnectorReopenDaysClosedLimit"`
	UseMemberDisplayNameFlag                  bool   `json:"useMemberDisplayNameFlag"`
	SendToCCFlag                              bool   `json:"sendToCCFlag"`
	AutoAssignTicketOwnerFlag                 bool   `json:"autoAssignTicketOwnerFlag"`
	AutoAssignLimitFlag                       bool   `json:"autoAssignLimitFlag"`
	AutoAssignLimitAmount                     int    `json:"autoAssignLimitAmount"`
	ClosedLoopAllFlag                         bool   `json:"closedLoopAllFlag"`
	PercentageCalculation                     string `json:"percentageCalculation"`
	AllSort                                   string `json:"allSort"`
	MarkFirstNoteIssueFlag                    bool   `json:"markFirstNoteIssueFlag"`
	RestrictBoardByDefaultFlag                bool   `json:"restrictBoardByDefaultFlag"`
	SendToBundledFlag                         bool   `json:"sendToBundledFlag"`
	Info                                      struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Ticket struct {
	Id         int    `json:"id"`
	Summary    string `json:"summary"`
	RecordType string `json:"recordType"`
	Board      struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"board"`
	Status struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Sort int    `json:"sort"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"status"`
	WorkRole struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"workRole"`
	WorkType struct {
		Id              int    `json:"id"`
		Name            string `json:"name"`
		UtilizationFlag bool   `json:"utilizationFlag"`
		Info            struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"workType"`
	Company struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"company"`
	Site struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"site"`
	SiteName        string `json:"siteName"`
	AddressLine1    string `json:"addressLine1"`
	AddressLine2    string `json:"addressLine2"`
	City            string `json:"city"`
	StateIdentifier string `json:"stateIdentifier"`
	Zip             string `json:"zip"`
	Country         struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"country"`
	Contact struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"contact"`
	ContactName           string `json:"contactName"`
	ContactPhoneNumber    string `json:"contactPhoneNumber"`
	ContactPhoneExtension string `json:"contactPhoneExtension"`
	ContactEmailAddress   string `json:"contactEmailAddress"`
	Type                  struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"type"`
	SubType struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"subType"`
	Item struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"item"`
	Team struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"team"`
	Owner struct {
		Id            int    `json:"id"`
		Identifier    string `json:"identifier"`
		Name          string `json:"name"`
		DailyCapacity int    `json:"dailyCapacity"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"owner"`
	Priority struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Sort  int    `json:"sort"`
		Level string `json:"level"`
		Info  struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"priority"`
	ServiceLocation struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"serviceLocation"`
	Source struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"source"`
	RequiredDate time.Time `json:"requiredDate"`
	BudgetHours  int       `json:"budgetHours"`
	Opportunity  struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"opportunity"`
	Agreement struct {
		Id             int    `json:"id"`
		Name           string `json:"name"`
		Type           string `json:"type"`
		ChargeFirmFlag bool   `json:"chargeFirmFlag"`
		Info           struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"agreement"`
	AgreementType              string    `json:"agreementType"`
	Severity                   string    `json:"severity"`
	Impact                     string    `json:"impact"`
	ExternalXRef               string    `json:"externalXRef"`
	PoNumber                   string    `json:"poNumber"`
	KnowledgeBaseCategoryId    int       `json:"knowledgeBaseCategoryId"`
	KnowledgeBaseSubCategoryId int       `json:"knowledgeBaseSubCategoryId"`
	AllowAllClientsPortalView  bool      `json:"allowAllClientsPortalView"`
	CustomerUpdatedFlag        bool      `json:"customerUpdatedFlag"`
	AutomaticEmailContactFlag  bool      `json:"automaticEmailContactFlag"`
	AutomaticEmailResourceFlag bool      `json:"automaticEmailResourceFlag"`
	AutomaticEmailCcFlag       bool      `json:"automaticEmailCcFlag"`
	AutomaticEmailCc           string    `json:"automaticEmailCc"`
	InitialDescription         string    `json:"initialDescription"`
	InitialInternalAnalysis    string    `json:"initialInternalAnalysis"`
	InitialResolution          string    `json:"initialResolution"`
	InitialDescriptionFrom     string    `json:"initialDescriptionFrom"`
	ContactEmailLookup         string    `json:"contactEmailLookup"`
	ProcessNotifications       bool      `json:"processNotifications"`
	SkipCallback               bool      `json:"skipCallback"`
	ClosedDate                 string    `json:"closedDate"`
	ClosedBy                   string    `json:"closedBy"`
	ClosedFlag                 bool      `json:"closedFlag"`
	ActualHours                int       `json:"actualHours"`
	Approved                   bool      `json:"approved"`
	EstimatedExpenseCost       int       `json:"estimatedExpenseCost"`
	EstimatedExpenseRevenue    int       `json:"estimatedExpenseRevenue"`
	EstimatedProductCost       int       `json:"estimatedProductCost"`
	EstimatedProductRevenue    int       `json:"estimatedProductRevenue"`
	EstimatedTimeCost          int       `json:"estimatedTimeCost"`
	EstimatedTimeRevenue       int       `json:"estimatedTimeRevenue"`
	BillingMethod              string    `json:"billingMethod"`
	BillingAmount              int       `json:"billingAmount"`
	HourlyRate                 int       `json:"hourlyRate"`
	SubBillingMethod           string    `json:"subBillingMethod"`
	SubBillingAmount           int       `json:"subBillingAmount"`
	SubDateAccepted            string    `json:"subDateAccepted"`
	DateResolved               string    `json:"dateResolved"`
	DateResplan                string    `json:"dateResplan"`
	DateResponded              string    `json:"dateResponded"`
	ResolveMinutes             int       `json:"resolveMinutes"`
	ResPlanMinutes             int       `json:"resPlanMinutes"`
	RespondMinutes             int       `json:"respondMinutes"`
	IsInSla                    bool      `json:"isInSla"`
	KnowledgeBaseLinkId        int       `json:"knowledgeBaseLinkId"`
	Resources                  string    `json:"resources"`
	ParentTicketId             int       `json:"parentTicketId"`
	HasChildTicket             bool      `json:"hasChildTicket"`
	HasMergedChildTicketFlag   bool      `json:"hasMergedChildTicketFlag"`
	KnowledgeBaseLinkType      string    `json:"knowledgeBaseLinkType"`
	BillTime                   string    `json:"billTime"`
	BillExpenses               string    `json:"billExpenses"`
	BillProducts               string    `json:"billProducts"`
	PredecessorType            string    `json:"predecessorType"`
	PredecessorId              int       `json:"predecessorId"`
	PredecessorClosedFlag      bool      `json:"predecessorClosedFlag"`
	LagDays                    int       `json:"lagDays"`
	LagNonworkingDaysFlag      bool      `json:"lagNonworkingDaysFlag"`
	EstimatedStartDate         time.Time `json:"estimatedStartDate"`
	Duration                   int       `json:"duration"`
	Location                   struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"location"`
	Department struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"department"`
	MobileGuid string `json:"mobileGuid"`
	Sla        struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"sla"`
	SlaStatus            string `json:"slaStatus"`
	RequestForChangeFlag bool   `json:"requestForChangeFlag"`
	Currency             struct {
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
	} `json:"currency"`
	MergedParentTicket struct {
		Id      int    `json:"id"`
		Summary string `json:"summary"`
		Info    struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"mergedParentTicket"`
	IntegratorTags []string `json:"integratorTags"`
	Info           struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
	EscalationStartDateUTC  string `json:"escalationStartDateUTC"`
	EscalationLevel         int    `json:"escalationLevel"`
	MinutesBeforeWaiting    int    `json:"minutesBeforeWaiting"`
	RespondedSkippedMinutes int    `json:"respondedSkippedMinutes"`
	ResplanSkippedMinutes   int    `json:"resplanSkippedMinutes"`
	RespondedHours          int    `json:"respondedHours"`
	RespondedBy             string `json:"respondedBy"`
	ResplanHours            int    `json:"resplanHours"`
	ResplanBy               string `json:"resplanBy"`
	ResolutionHours         int    `json:"resolutionHours"`
	ResolvedBy              string `json:"resolvedBy"`
	MinutesWaiting          int    `json:"minutesWaiting"`
	CustomFields            []struct {
		Id               int    `json:"id"`
		Caption          string `json:"caption"`
		Type             string `json:"type"`
		EntryMethod      string `json:"entryMethod"`
		NumberOfDecimals int    `json:"numberOfDecimals"`
		Value            struct {
		} `json:"value"`
		ConnectWiseId string `json:"connectWiseId"`
	} `json:"customFields"`
}

type ServiceTicketNote struct {
	Id                    int    `json:"id"`
	TicketId              int    `json:"ticketId"`
	Text                  string `json:"text"`
	DetailDescriptionFlag bool   `json:"detailDescriptionFlag"`
	InternalAnalysisFlag  bool   `json:"internalAnalysisFlag"`
	ResolutionFlag        bool   `json:"resolutionFlag"`
	IssueFlag             bool   `json:"issueFlag"`
	Member                struct {
		Id            int    `json:"id"`
		Identifier    string `json:"identifier"`
		Name          string `json:"name"`
		DailyCapacity int    `json:"dailyCapacity"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"member"`
	Contact struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"contact"`
	CustomerUpdatedFlag  bool `json:"customerUpdatedFlag"`
	ProcessNotifications bool `json:"processNotifications"`
	InternalFlag         bool `json:"internalFlag"`
	ExternalFlag         bool `json:"externalFlag"`
	Info                 struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type ServiceTicketNoteAll struct {
	Id       int    `json:"id"`
	NoteType string `json:"noteType"`
	Ticket   struct {
		Id      int    `json:"id"`
		Summary string `json:"summary"`
		Info    struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"ticket"`
	Text                  string `json:"text"`
	IsMarkdownFlag        bool   `json:"isMarkdownFlag"`
	DetailDescriptionFlag bool   `json:"detailDescriptionFlag"`
	InternalAnalysisFlag  bool   `json:"internalAnalysisFlag"`
	ResolutionFlag        bool   `json:"resolutionFlag"`
	TimeStart             string `json:"timeStart"`
	TimeEnd               string `json:"timeEnd"`
	BundledFlag           bool   `json:"bundledFlag"`
	MergedFlag            bool   `json:"mergedFlag"`
	IssueFlag             bool   `json:"issueFlag"`
	OriginalAuthor        string `json:"originalAuthor"`
	CreatedByParentFlag   bool   `json:"createdByParentFlag"`
	Member                struct {
		Id            int    `json:"id"`
		Identifier    string `json:"identifier"`
		Name          string `json:"name"`
		DailyCapacity int    `json:"dailyCapacity"`
		Info          struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"member"`
	Contact struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"_info"`
	} `json:"contact"`
	Info struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}
