package connectwise

type Department struct {
	Id         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Info       struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"_info"`
}

type Location struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info struct {
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

type Board struct {
	Id                                        int                   `json:"id"`
	Name                                      string                `json:"name"`
	Location                                  Location              `json:"location"`
	Department                                Department            `json:"department"`
	InactiveFlag                              bool                  `json:"inactiveFlag"`
	SignOffTemplate                           SignOffTemplate       `json:"signOffTemplate"`
	SendToContactFlag                         bool                  `json:"sendToContactFlag"`
	ContactTemplate                           EmailTemplate         `json:"contactTemplate"`
	SendToResourceFlag                        bool                  `json:"sendToResourceFlag"`
	ResourceTemplate                          EmailTemplate         `json:"resourceTemplate"`
	ProjectFlag                               bool                  `json:"projectFlag"`
	ShowDependenciesFlag                      bool                  `json:"showDependenciesFlag"`
	ShowEstimatesFlag                         bool                  `json:"showEstimatesFlag"`
	BoardIcon                                 Document              `json:"boardIcon"`
	BillTicketsAfterClosedFlag                bool                  `json:"billTicketsAfterClosedFlag"`
	BillTicketSeparatelyFlag                  bool                  `json:"billTicketSeparatelyFlag"`
	BillUnapprovedTimeExpenseFlag             bool                  `json:"billUnapprovedTimeExpenseFlag"`
	OverrideBillingSetupFlag                  bool                  `json:"overrideBillingSetupFlag"`
	DispatchMember                            Member                `json:"dispatchMember"`
	ServiceManagerMember                      Member                `json:"serviceManagerMember"`
	DutyManagerMember                         Member                `json:"dutyManagerMember"`
	OncallMember                              Member                `json:"oncallMember"`
	WorkRole                                  WorkRole              `json:"workRole"`
	WorkType                                  WorkType              `json:"workType"`
	BillTime                                  string                `json:"billTime"`
	BillExpense                               string                `json:"billExpense"`
	BillProduct                               string                `json:"billProduct"`
	AutoCloseStatus                           Status                `json:"autoCloseStatus"`
	AutoAssignNewTicketsFlag                  bool                  `json:"autoAssignNewTicketsFlag"`
	AutoAssignNewECTicketsFlag                bool                  `json:"autoAssignNewECTicketsFlag"`
	AutoAssignNewPortalTicketsFlag            bool                  `json:"autoAssignNewPortalTicketsFlag"`
	DiscussionsLockedFlag                     bool                  `json:"discussionsLockedFlag"`
	TimeEntryLockedFlag                       bool                  `json:"timeEntryLockedFlag"`
	NotifyEmailFrom                           string                `json:"notifyEmailFrom"`
	NotifyEmailFromName                       string                `json:"notifyEmailFromName"`
	ClosedLoopDiscussionsFlag                 bool                  `json:"closedLoopDiscussionsFlag"`
	ClosedLoopResolutionFlag                  bool                  `json:"closedLoopResolutionFlag"`
	ClosedLoopInternalAnalysisFlag            bool                  `json:"closedLoopInternalAnalysisFlag"`
	TimeEntryDiscussionFlag                   bool                  `json:"timeEntryDiscussionFlag"`
	TimeEntryResolutionFlag                   bool                  `json:"timeEntryResolutionFlag"`
	TimeEntryInternalAnalysisFlag             bool                  `json:"timeEntryInternalAnalysisFlag"`
	ProblemSort                               string                `json:"problemSort"`
	ResolutionSort                            string                `json:"resolutionSort"`
	InternalAnalysisSort                      string                `json:"internalAnalysisSort"`
	EmailConnectorAllowReopenClosedFlag       bool                  `json:"emailConnectorAllowReopenClosedFlag"`
	EmailConnectorReopenStatus                ConnectorReopenStatus `json:"emailConnectorReopenStatus"`
	EmailConnectorReopenResourcesFlag         bool                  `json:"emailConnectorReopenResourcesFlag"`
	EmailConnectorNewTicketNoMatchFlag        bool                  `json:"emailConnectorNewTicketNoMatchFlag"`
	EmailConnectorNeverReopenByDaysFlag       bool                  `json:"emailConnectorNeverReopenByDaysFlag"`
	EmailConnectorReopenDaysLimit             int                   `json:"emailConnectorReopenDaysLimit"`
	EmailConnectorNeverReopenByDaysClosedFlag bool                  `json:"emailConnectorNeverReopenByDaysClosedFlag"`
	EmailConnectorReopenDaysClosedLimit       int                   `json:"emailConnectorReopenDaysClosedLimit"`
	UseMemberDisplayNameFlag                  bool                  `json:"useMemberDisplayNameFlag"`
	SendToCCFlag                              bool                  `json:"sendToCCFlag"`
	AutoAssignTicketOwnerFlag                 bool                  `json:"autoAssignTicketOwnerFlag"`
	AutoAssignLimitFlag                       bool                  `json:"autoAssignLimitFlag"`
	AutoAssignLimitAmount                     int                   `json:"autoAssignLimitAmount"`
	ClosedLoopAllFlag                         bool                  `json:"closedLoopAllFlag"`
	PercentageCalculation                     string                `json:"percentageCalculation"`
	AllSort                                   string                `json:"allSort"`
	MarkFirstNoteIssueFlag                    bool                  `json:"markFirstNoteIssueFlag"`
	RestrictBoardByDefaultFlag                bool                  `json:"restrictBoardByDefaultFlag"`
	SendToBundledFlag                         bool                  `json:"sendToBundledFlag"`
	Info                                      Info                  `json:"_info"`
}
