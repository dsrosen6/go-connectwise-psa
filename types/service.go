package types

type Board struct {
	Id                                        int                  `json:"id,omitempty"`
	Name                                      string               `json:"name"`
	Location                                  SystemLocation       `json:"location"`
	Department                                Department           `json:"department"`
	InactiveFlag                              bool                 `json:"inactiveFlag,omitempty"`
	SignOffTemplate                           ServiceSignoff       `json:"signOffTemplate,omitempty"`
	SendToContactFlag                         bool                 `json:"sendToContactFlag,omitempty"`
	ContactTemplate                           ServiceEmailTemplate `json:"contactTemplate,omitempty"`
	SendToResourceFlag                        bool                 `json:"sendToResourceFlag,omitempty"`
	ResourceTemplate                          ServiceEmailTemplate `json:"resourceTemplate,omitempty"`
	ProjectFlag                               bool                 `json:"projectFlag,omitempty"`
	ShowDependenciesFlag                      bool                 `json:"showDependenciesFlag,omitempty"`
	ShowEstimatesFlag                         bool                 `json:"showEstimatesFlag,omitempty"`
	BoardIcon                                 Document             `json:"boardIcon,omitempty"`
	BillTicketsAfterClosedFlag                bool                 `json:"billTicketsAfterClosedFlag,omitempty"`
	BillTicketSeparatelyFlag                  bool                 `json:"billTicketSeparatelyFlag,omitempty"`
	BillUnapprovedTimeExpenseFlag             bool                 `json:"billUnapprovedTimeExpenseFlag,omitempty"`
	OverrideBillingSetupFlag                  bool                 `json:"overrideBillingSetupFlag,omitempty"`
	DispatchMember                            Member               `json:"dispatchMember,omitempty"`
	ServiceManagerMember                      Member               `json:"serviceManagerMember,omitempty"`
	DutyManagerMember                         Member               `json:"dutyManagerMember,omitempty"`
	OncallMember                              Member               `json:"oncallMember,omitempty"`
	WorkRole                                  WorkRole             `json:"workRole,omitempty"`
	WorkType                                  WorkType             `json:"workType,omitempty"`
	BillTime                                  string               `json:"billTime,omitempty"`
	BillExpense                               string               `json:"billExpense,omitempty"`
	BillProduct                               string               `json:"billProduct,omitempty"`
	AutoCloseStatus                           ServiceStatus        `json:"autoCloseStatus,omitempty"`
	AutoAssignNewTicketsFlag                  bool                 `json:"autoAssignNewTicketsFlag,omitempty"`
	AutoAssignNewECTicketsFlag                bool                 `json:"autoAssignNewECTicketsFlag,omitempty"`
	AutoAssignNewPortalTicketsFlag            bool                 `json:"autoAssignNewPortalTicketsFlag,omitempty"`
	DiscussionsLockedFlag                     bool                 `json:"discussionsLockedFlag,omitempty"`
	TimeEntryLockedFlag                       bool                 `json:"timeEntryLockedFlag,omitempty"`
	NotifyEmailFrom                           string               `json:"notifyEmailFrom,omitempty"`
	NotifyEmailFromName                       string               `json:"notifyEmailFromName,omitempty"`
	ClosedLoopDiscussionsFlag                 bool                 `json:"closedLoopDiscussionsFlag,omitempty"`
	ClosedLoopResolutionFlag                  bool                 `json:"closedLoopResolutionFlag,omitempty"`
	ClosedLoopInternalAnalysisFlag            bool                 `json:"closedLoopInternalAnalysisFlag,omitempty"`
	TimeEntryDiscussionFlag                   bool                 `json:"timeEntryDiscussionFlag,omitempty"`
	TimeEntryResolutionFlag                   bool                 `json:"timeEntryResolutionFlag,omitempty"`
	TimeEntryInternalAnalysisFlag             bool                 `json:"timeEntryInternalAnalysisFlag,omitempty"`
	ProblemSort                               string               `json:"problemSort,omitempty"`
	ResolutionSort                            string               `json:"resolutionSort,omitempty"`
	InternalAnalysisSort                      string               `json:"internalAnalysisSort,omitempty"`
	EmailConnectorAllowReopenClosedFlag       bool                 `json:"emailConnectorAllowReopenClosedFlag,omitempty"`
	EmailConnectorReopenStatus                ServiceStatus        `json:"emailConnectorReopenStatus,omitempty"`
	EmailConnectorReopenResourcesFlag         bool                 `json:"emailConnectorReopenResourcesFlag,omitempty"`
	EmailConnectorNewTicketNoMatchFlag        bool                 `json:"emailConnectorNewTicketNoMatchFlag,omitempty"`
	EmailConnectorNeverReopenByDaysFlag       bool                 `json:"emailConnectorNeverReopenByDaysFlag,omitempty"`
	EmailConnectorReopenDaysLimit             int                  `json:"emailConnectorReopenDaysLimit,omitempty"`
	EmailConnectorNeverReopenByDaysClosedFlag bool                 `json:"emailConnectorNeverReopenByDaysClosedFlag,omitempty"`
	EmailConnectorReopenDaysClosedLimit       int                  `json:"emailConnectorReopenDaysClosedLimit,omitempty"`
	UseMemberDisplayNameFlag                  bool                 `json:"useMemberDisplayNameFlag,omitempty"`
	SendToCCFlag                              bool                 `json:"sendToCCFlag,omitempty"`
	AutoAssignTicketOwnerFlag                 bool                 `json:"autoAssignTicketOwnerFlag,omitempty"`
	AutoAssignLimitFlag                       bool                 `json:"autoAssignLimitFlag,omitempty"`
	AutoAssignLimitAmount                     int                  `json:"autoAssignLimitAmount,omitempty"`
	ClosedLoopAllFlag                         bool                 `json:"closedLoopAllFlag,omitempty"`
	PercentageCalculation                     string               `json:"percentageCalculation,omitempty"`
	AllSort                                   string               `json:"allSort,omitempty"`
	MarkFirstNoteIssueFlag                    bool                 `json:"markFirstNoteIssueFlag,omitempty"`
	RestrictBoardByDefaultFlag                bool                 `json:"restrictBoardByDefaultFlag,omitempty"`
	SendToBundledFlag                         bool                 `json:"sendToBundledFlag,omitempty"`
	Info                                      Info                 `json:"_info,omitempty"`
}

type BoardStatus struct {
	Id                        int                  `json:"id,omitempty"`
	Name                      string               `json:"name"`
	Board                     Board                `json:"board,omitempty"`
	SortOrder                 int                  `json:"sortOrder,omitempty"`
	DisplayOnBoard            bool                 `json:"displayOnBoard,omitempty"`
	Inactive                  bool                 `json:"inactive,omitempty"`
	ClosedStatus              bool                 `json:"closedStatus,omitempty"`
	TimeEntryNotAllowed       bool                 `json:"timeEntryNotAllowed,omitempty"`
	RoundRobinCatchall        bool                 `json:"roundRobinCatchall,omitempty"`
	DefaultFlag               bool                 `json:"defaultFlag,omitempty"`
	EscalationStatus          string               `json:"escalationStatus,omitempty"`
	CustomerPortalDescription string               `json:"customerPortalDescription,omitempty"`
	CustomerPortalFlag        bool                 `json:"customerPortalFlag,omitempty"`
	EmailTemplate             ServiceEmailTemplate `json:"emailTemplate,omitempty"`
	StatusIndicator           StatusIndicator      `json:"statusIndicator,omitempty"`
	CustomStatusIndicatorName string               `json:"customStatusIndicatorName,omitempty"`
	Info                      Info                 `json:"_info,omitempty"`
	SaveTimeAsNote            bool                 `json:"saveTimeAsNote,omitempty"`
}

type ServiceType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type ServiceSubType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type ServiceItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
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

type ServiceTeam struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type ServiceLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type ServiceSource struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type Sla struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type ServiceSignoff struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info Info   `json:"_info"`
}

type ServiceEmailTemplate struct {
	Id         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Info       Info   `json:"_info"`
}

type ServiceStatus struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
	Info Info   `json:"_info"`
}
