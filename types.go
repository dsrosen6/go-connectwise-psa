package connectwise

import "time"

type Board struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"location"`
	Department struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"department"`
	InactiveFlag    bool `json:"inactiveFlag"`
	SignOffTemplate struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SignOffTemplateHref string `json:"signOffTemplate_href"`
		} `json:"_info"`
	} `json:"signOffTemplate"`
	SendToContactFlag bool `json:"sendToContactFlag"`
	ContactTemplate   struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Type       string `json:"type"`
		Info       struct {
			EmailTemplateHref string `json:"emailTemplate_href"`
		} `json:"_info"`
	} `json:"contactTemplate,omitempty"`
	SendToResourceFlag bool `json:"sendToResourceFlag"`
	ProjectFlag        bool `json:"projectFlag"`
	BoardIcon          struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			Filename             string `json:"filename"`
			DocumentHref         string `json:"document_href"`
			DocumentDownloadHref string `json:"documentDownload_href"`
		} `json:"_info"`
	} `json:"boardIcon,omitempty"`
	BillTicketsAfterClosedFlag    bool `json:"billTicketsAfterClosedFlag"`
	BillTicketSeparatelyFlag      bool `json:"billTicketSeparatelyFlag"`
	BillUnapprovedTimeExpenseFlag bool `json:"billUnapprovedTimeExpenseFlag"`
	OverrideBillingSetupFlag      bool `json:"overrideBillingSetupFlag"`
	OncallMember                  struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"oncallMember,omitempty"`
	WorkRole struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			WorkRoleHref string `json:"workRole_href"`
		} `json:"_info"`
	} `json:"workRole,omitempty"`
	WorkType struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			WorkTypeHref string `json:"workType_href"`
		} `json:"_info"`
	} `json:"workType,omitempty"`
	BillTime                            string `json:"billTime"`
	BillExpense                         string `json:"billExpense"`
	BillProduct                         string `json:"billProduct"`
	AutoAssignNewTicketsFlag            bool   `json:"autoAssignNewTicketsFlag"`
	AutoAssignNewECTicketsFlag          bool   `json:"autoAssignNewECTicketsFlag"`
	AutoAssignNewPortalTicketsFlag      bool   `json:"autoAssignNewPortalTicketsFlag"`
	DiscussionsLockedFlag               bool   `json:"discussionsLockedFlag"`
	TimeEntryLockedFlag                 bool   `json:"timeEntryLockedFlag"`
	NotifyEmailFrom                     string `json:"notifyEmailFrom,omitempty"`
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
	} `json:"emailConnectorReopenStatus,omitempty"`
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
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
	} `json:"_info"`
	ShowDependenciesFlag bool `json:"showDependenciesFlag,omitempty"`
	ShowEstimatesFlag    bool `json:"showEstimatesFlag,omitempty"`
	DispatchMember       struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"dispatchMember,omitempty"`
	ServiceManagerMember struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"serviceManagerMember,omitempty"`
	ResourceTemplate struct {
		Id         int    `json:"id"`
		Identifier string `json:"identifier"`
		Type       string `json:"type"`
		Info       struct {
			EmailTemplateHref string `json:"emailTemplate_href"`
		} `json:"_info"`
	} `json:"resourceTemplate,omitempty"`
}
