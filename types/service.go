package types

import (
	"time"
)

type ActivityReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Board struct {
	Info interface{} `json:"_info,omitempty"`
	AllSort string `json:"allSort,omitempty"`
	AutoAssignLimitAmount int `json:"autoAssignLimitAmount,omitempty"`
	AutoAssignLimitFlag bool `json:"autoAssignLimitFlag,omitempty"`
	AutoAssignNewECTicketsFlag bool `json:"autoAssignNewECTicketsFlag,omitempty"`
	AutoAssignNewPortalTicketsFlag bool `json:"autoAssignNewPortalTicketsFlag,omitempty"`
	AutoAssignNewTicketsFlag bool `json:"autoAssignNewTicketsFlag,omitempty"`
	AutoAssignTicketOwnerFlag bool `json:"autoAssignTicketOwnerFlag,omitempty"`
	AutoCloseStatus ServiceStatusReference `json:"autoCloseStatus,omitempty"`
	BillExpense string `json:"billExpense,omitempty"`
	BillProduct string `json:"billProduct,omitempty"`
	BillTicketSeparatelyFlag bool `json:"billTicketSeparatelyFlag,omitempty"`
	BillTicketsAfterClosedFlag bool `json:"billTicketsAfterClosedFlag,omitempty"`
	BillTime string `json:"billTime,omitempty"`
	BillUnapprovedTimeExpenseFlag bool `json:"billUnapprovedTimeExpenseFlag,omitempty"`
	BoardIcon DocumentReference `json:"boardIcon,omitempty"`
	ClosedLoopAllFlag bool `json:"closedLoopAllFlag,omitempty"`
	ClosedLoopDiscussionsFlag bool `json:"closedLoopDiscussionsFlag,omitempty"`
	ClosedLoopInternalAnalysisFlag bool `json:"closedLoopInternalAnalysisFlag,omitempty"`
	ClosedLoopResolutionFlag bool `json:"closedLoopResolutionFlag,omitempty"`
	ContactTemplate ServiceEmailTemplateReference `json:"contactTemplate,omitempty"`
	Department SystemDepartmentReference `json:"department"`
	DiscussionsLockedFlag bool `json:"discussionsLockedFlag,omitempty"`
	DispatchMember MemberReference `json:"dispatchMember,omitempty"`
	DutyManagerMember MemberReference `json:"dutyManagerMember,omitempty"`
	EmailConnectorAllowReopenClosedFlag bool `json:"emailConnectorAllowReopenClosedFlag,omitempty"`
	EmailConnectorNeverReopenByDaysClosedFlag bool `json:"emailConnectorNeverReopenByDaysClosedFlag,omitempty"`
	EmailConnectorNeverReopenByDaysFlag bool `json:"emailConnectorNeverReopenByDaysFlag,omitempty"`
	EmailConnectorNewTicketNoMatchFlag bool `json:"emailConnectorNewTicketNoMatchFlag,omitempty"`
	EmailConnectorReopenDaysClosedLimit int `json:"emailConnectorReopenDaysClosedLimit,omitempty"`
	EmailConnectorReopenDaysLimit int `json:"emailConnectorReopenDaysLimit,omitempty"`
	EmailConnectorReopenResourcesFlag bool `json:"emailConnectorReopenResourcesFlag,omitempty"`
	EmailConnectorReopenStatus ServiceStatusReference `json:"emailConnectorReopenStatus,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	InternalAnalysisSort string `json:"internalAnalysisSort,omitempty"`
	Location SystemLocationReference `json:"location"`
	MarkFirstNoteIssueFlag bool `json:"markFirstNoteIssueFlag,omitempty"`
	Name string `json:"name"`
	NotifyEmailFrom string `json:"notifyEmailFrom,omitempty"`
	NotifyEmailFromName string `json:"notifyEmailFromName,omitempty"`
	OncallMember MemberReference `json:"oncallMember,omitempty"`
	OverrideBillingSetupFlag bool `json:"overrideBillingSetupFlag,omitempty"`
	PercentageCalculation string `json:"percentageCalculation,omitempty"`
	ProblemSort string `json:"problemSort,omitempty"`
	ProjectFlag bool `json:"projectFlag,omitempty"`
	ResolutionSort string `json:"resolutionSort,omitempty"`
	ResourceTemplate ServiceEmailTemplateReference `json:"resourceTemplate,omitempty"`
	RestrictBoardByDefaultFlag bool `json:"restrictBoardByDefaultFlag,omitempty"`
	SendToBundledFlag bool `json:"sendToBundledFlag,omitempty"`
	SendToCCFlag bool `json:"sendToCCFlag,omitempty"`
	SendToContactFlag bool `json:"sendToContactFlag,omitempty"`
	SendToResourceFlag bool `json:"sendToResourceFlag,omitempty"`
	ServiceManagerMember MemberReference `json:"serviceManagerMember,omitempty"`
	ShowDependenciesFlag bool `json:"showDependenciesFlag,omitempty"`
	ShowEstimatesFlag bool `json:"showEstimatesFlag,omitempty"`
	SignOffTemplate ServiceSignoffReference `json:"signOffTemplate,omitempty"`
	TimeEntryDiscussionFlag bool `json:"timeEntryDiscussionFlag,omitempty"`
	TimeEntryInternalAnalysisFlag bool `json:"timeEntryInternalAnalysisFlag,omitempty"`
	TimeEntryLockedFlag bool `json:"timeEntryLockedFlag,omitempty"`
	TimeEntryResolutionFlag bool `json:"timeEntryResolutionFlag,omitempty"`
	UseMemberDisplayNameFlag bool `json:"useMemberDisplayNameFlag,omitempty"`
	WorkRole WorkRoleReference `json:"workRole,omitempty"`
	WorkType WorkTypeReference `json:"workType,omitempty"`
}

type BoardAutoAssignResource struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	NotifyWho NotificationRecipientReference `json:"notifyWho"`
}

type BoardAutoTemplate struct {
	Info interface{} `json:"_info,omitempty"`
	AutoApplyFlag bool `json:"autoApplyFlag,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	BudgetHoursSetting string `json:"budgetHoursSetting,omitempty"`
	DiscussionSetting string `json:"discussionSetting,omitempty"`
	DocumentsSetting string `json:"documentsSetting,omitempty"`
	FinanceInformationSetting string `json:"financeInformationSetting,omitempty"`
	ID int `json:"id,omitempty"`
	ImpactUrgencySetting string `json:"impactUrgencySetting,omitempty"`
	InternalAnalysisSetting string `json:"internalAnalysisSetting,omitempty"`
	Item ServiceItemReference `json:"item"`
	ResolutionSetting string `json:"resolutionSetting,omitempty"`
	ResourcesSetting string `json:"resourcesSetting,omitempty"`
	SendNotesAsEmailSetting string `json:"sendNotesAsEmailSetting,omitempty"`
	ServiceTemplate ServiceTemplateReference `json:"serviceTemplate"`
	Subtype ServiceSubTypeReference `json:"subtype"`
	SummarySetting string `json:"summarySetting,omitempty"`
	TasksSetting string `json:"tasksSetting,omitempty"`
	TemplatePrioritySetting string `json:"templatePrioritySetting,omitempty"`
	Type ServiceTypeReference `json:"type"`
}

type BoardCopy struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type BoardExcludedMember struct {
	Info interface{} `json:"_info,omitempty"`
	BoardId int `json:"boardId,omitempty"`
	ID int `json:"id,omitempty"`
	MemberId int `json:"memberId,omitempty"`
}

type BoardInfo struct {
	Info interface{} `json:"_info,omitempty"`
	AllSort string `json:"allSort,omitempty"`
	BillExpense string `json:"billExpense,omitempty"`
	BillProduct string `json:"billProduct,omitempty"`
	BillTicketsAfterClosedFlag bool `json:"billTicketsAfterClosedFlag,omitempty"`
	BillTime string `json:"billTime,omitempty"`
	BillUnapprovedTimeExpenseFlag bool `json:"billUnapprovedTimeExpenseFlag,omitempty"`
	ClosedLoopAllFlag bool `json:"closedLoopAllFlag,omitempty"`
	ClosedLoopDiscussionsFlag bool `json:"closedLoopDiscussionsFlag,omitempty"`
	ClosedLoopInternalAnalysisFlag bool `json:"closedLoopInternalAnalysisFlag,omitempty"`
	ClosedLoopResolutionFlag bool `json:"closedLoopResolutionFlag,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	InternalAnalysisSort string `json:"internalAnalysisSort,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
	OverrideBillingSetupFlag bool `json:"overrideBillingSetupFlag,omitempty"`
	ProblemSort string `json:"problemSort,omitempty"`
	ProjectFlag bool `json:"projectFlag,omitempty"`
	ResolutionSort string `json:"resolutionSort,omitempty"`
}

type BoardItem struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name"`
}

type BoardItemAssociation struct {
	Info interface{} `json:"_info,omitempty"`
	AddAllSubTypesFlag bool `json:"addAllSubTypesFlag,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	ID int `json:"id"`
	Item ServiceItemReference `json:"item,omitempty"`
	RemoveAllSubTypesFlag bool `json:"removeAllSubTypesFlag,omitempty"`
	SubTypeAssociationIds []int `json:"subTypeAssociationIds,omitempty"`
}

type BoardNotification struct {
	Info interface{} `json:"_info,omitempty"`
	Email string `json:"email,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	NotifyWho NotificationRecipientReference `json:"notifyWho"`
}

type BoardSkillMapping struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	ID int `json:"id,omitempty"`
	Item ServiceItemReference `json:"item,omitempty"`
	Skill SkillReference `json:"skill"`
	SkillCategory SkillCategoryReference `json:"skillCategory"`
	SubType ServiceSubTypeReference `json:"subType,omitempty"`
	Type ServiceTypeReference `json:"type"`
}

type BoardStatus struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	ClosedStatus bool `json:"closedStatus,omitempty"`
	CustomStatusIndicatorName string `json:"customStatusIndicatorName,omitempty"`
	CustomerPortalDescription string `json:"customerPortalDescription,omitempty"`
	CustomerPortalFlag bool `json:"customerPortalFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	DisplayOnBoard bool `json:"displayOnBoard,omitempty"`
	EmailTemplate ServiceEmailTemplateReference `json:"emailTemplate,omitempty"`
	EscalationStatus string `json:"escalationStatus,omitempty"`
	ID int `json:"id,omitempty"`
	Inactive bool `json:"inactive,omitempty"`
	Name string `json:"name"`
	RoundRobinCatchall bool `json:"roundRobinCatchall,omitempty"`
	SaveTimeAsNote bool `json:"saveTimeAsNote,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
	StatusIndicator StatusIndicatorReference `json:"statusIndicator,omitempty"`
	TimeEntryNotAllowed bool `json:"timeEntryNotAllowed,omitempty"`
}

type BoardStatusInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ClosedFlag bool `json:"closedFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type BoardStatusNotification struct {
	Info interface{} `json:"_info,omitempty"`
	Email string `json:"email,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	NotifyWho NotificationRecipientReference `json:"notifyWho"`
	Status ServiceStatusReference `json:"status,omitempty"`
	WorkflowStep int `json:"workflowStep,omitempty"`
}

type BoardSubType struct {
	Info interface{} `json:"_info,omitempty"`
	AddAllTypesFlag bool `json:"addAllTypesFlag,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name"`
	RemoveAllTypesFlag bool `json:"removeAllTypesFlag,omitempty"`
	TypeAssociationIds []int `json:"typeAssociationIds,omitempty"`
}

type BoardSubTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
	TypeAssociationIds []int `json:"typeAssociationIds,omitempty"`
}

type BoardTeam struct {
	Info interface{} `json:"_info,omitempty"`
	BoardId int `json:"boardId,omitempty"`
	BusinessUnitId int `json:"businessUnitId,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	DefaultRoundRobinFlag bool `json:"defaultRoundRobinFlag,omitempty"`
	ID int `json:"id,omitempty"`
	LocationId int `json:"locationId,omitempty"`
	Members []int `json:"members,omitempty"`
	Name string `json:"name"`
	NotifyOnTicketDelete bool `json:"notifyOnTicketDelete,omitempty"`
	RoundRobinFlag bool `json:"roundRobinFlag,omitempty"`
	TeamLeader MemberReference `json:"teamLeader"`
}

type BoardTeamInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BoardType struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	Category string `json:"category,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	IntegrationXref string `json:"integrationXref,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name"`
	RequestForChangeFlag bool `json:"requestForChangeFlag,omitempty"`
	Skill SkillReference `json:"skill,omitempty"`
	SkillCategory SkillCategoryReference `json:"skillCategory,omitempty"`
}

type BoardTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type BoardTypeSubTypeItemAssociation struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	ID int `json:"id,omitempty"`
	Item ServiceItemReference `json:"item,omitempty"`
	SubType ServiceSubTypeReference `json:"subType,omitempty"`
	Type ServiceTypeReference `json:"type,omitempty"`
}

type Code struct {
	Info interface{} `json:"_info,omitempty"`
	BoardId int `json:"boardId,omitempty"`
	BusinessUnitId int `json:"businessUnitId,omitempty"`
	Description string `json:"description"`
	ID int `json:"id,omitempty"`
	LocationId int `json:"locationId,omitempty"`
	Name string `json:"name"`
}

type ConfigurationReference struct {
	Info interface{} `json:"_info,omitempty"`
	DeviceIdentifier string `json:"deviceIdentifier,omitempty"`
	ID int `json:"id,omitempty"`
}

type ConvertToProject struct {
	ID int `json:"id,omitempty"`
	Phase ProjectPhaseReference `json:"phase"`
	Project ProjectReference `json:"project,omitempty"`
	RecordType string `json:"recordType,omitempty"`
	WbsCode string `json:"wbsCode"`
}

type Impact struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	Description string `json:"description"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type KnowledgeBaseArticle struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	BusinessUnitId int `json:"businessUnitId,omitempty"`
	CategoryId int `json:"categoryId,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	DateCreated string `json:"dateCreated,omitempty"`
	ID int `json:"id,omitempty"`
	Issue string `json:"issue"`
	LocationId int `json:"locationId,omitempty"`
	Resolution string `json:"resolution"`
	SubCategoryId int `json:"subCategoryId,omitempty"`
	Title string `json:"title"`
}

type KnowledgeBaseCategory struct {
	Info interface{} `json:"_info,omitempty"`
	Approver MemberReference `json:"approver,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name"`
}

type KnowledgeBaseSettings struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultApprover MemberReference `json:"defaultApprover,omitempty"`
	ID int `json:"id,omitempty"`
	RequireApproval bool `json:"requireApproval,omitempty"`
}

type KnowledgeBaseSubCategory struct {
	Info interface{} `json:"_info,omitempty"`
	Category KBCategoryReference `json:"category"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name"`
}

type Priority struct {
	Info interface{} `json:"_info,omitempty"`
	Color string `json:"color,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	ImageLink string `json:"imageLink,omitempty"`
	Level string `json:"level,omitempty"`
	Name string `json:"name"`
	SortOrder int `json:"sortOrder,omitempty"`
	UrgencySortOrder string `json:"urgencySortOrder,omitempty"`
}

type PriorityInfo struct {
	Info interface{} `json:"_info,omitempty"`
	Color string `json:"color,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	ImageLink string `json:"imageLink,omitempty"`
	Level string `json:"level,omitempty"`
	Name string `json:"name,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
	UrgencySortOrder string `json:"urgencySortOrder,omitempty"`
}

type SLA struct {
	Info interface{} `json:"_info,omitempty"`
	ApplicationOrder int `json:"applicationOrder,omitempty"`
	BasedOn string `json:"basedOn,omitempty"`
	CustomCalendar CalendarReference `json:"customCalendar,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	HiImpactHiUrgency PriorityReference `json:"hiImpactHiUrgency,omitempty"`
	HiImpactLowUrgency PriorityReference `json:"hiImpactLowUrgency,omitempty"`
	HiImpactMedUrgency PriorityReference `json:"hiImpactMedUrgency,omitempty"`
	ID int `json:"id,omitempty"`
	LowImpactHiUrgency PriorityReference `json:"lowImpactHiUrgency,omitempty"`
	LowImpactLowUrgency PriorityReference `json:"lowImpactLowUrgency,omitempty"`
	LowImpactMedUrgency PriorityReference `json:"lowImpactMedUrgency,omitempty"`
	MedImpactHiUrgency PriorityReference `json:"medImpactHiUrgency,omitempty"`
	MedImpactLowUrgency PriorityReference `json:"medImpactLowUrgency,omitempty"`
	MedImpactMedUrgency PriorityReference `json:"medImpactMedUrgency,omitempty"`
	Name string `json:"name"`
	PlanWithin float64 `json:"planWithin,omitempty"`
	PlanWithinPercent int `json:"planWithinPercent,omitempty"`
	ResolutionHours float64 `json:"resolutionHours,omitempty"`
	ResolutionPercent int `json:"resolutionPercent,omitempty"`
	RespondHours float64 `json:"respondHours,omitempty"`
	RespondPercent int `json:"respondPercent,omitempty"`
}

type SLAInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SLAPriority struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	PlanWithin float64 `json:"planWithin,omitempty"`
	PlanWithinPercent int `json:"planWithinPercent,omitempty"`
	Priority PriorityReference `json:"priority"`
	ResolutionHours float64 `json:"resolutionHours,omitempty"`
	ResolutionPercent int `json:"resolutionPercent,omitempty"`
	RespondHours float64 `json:"respondHours,omitempty"`
	RespondPercent int `json:"respondPercent,omitempty"`
	Sla SLAReference `json:"sla,omitempty"`
}

type SchedulingMemberInfo struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultEmail string `json:"defaultEmail,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	FullName string `json:"fullName,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	LastName string `json:"lastName,omitempty"`
	MiddleInitial string `json:"middleInitial,omitempty"`
}

type ServiceEmailTemplate struct {
	Info interface{} `json:"_info,omitempty"`
	Body string `json:"body,omitempty"`
	CopySenderFlag bool `json:"copySenderFlag,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	ExternalContactNotifications bool `json:"externalContactNotifications,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	ID int `json:"id,omitempty"`
	InternalContactNotifications bool `json:"internalContactNotifications,omitempty"`
	LastName string `json:"lastName,omitempty"`
	ResourceRecordsFlag bool `json:"resourceRecordsFlag,omitempty"`
	ServiceBoard BoardReference `json:"serviceBoard,omitempty"`
	ServiceStatus ServiceStatusReference `json:"serviceStatus,omitempty"`
	ServiceSurvey ServiceSurveyReference `json:"serviceSurvey,omitempty"`
	Subject string `json:"subject,omitempty"`
	TasksFlag bool `json:"tasksFlag,omitempty"`
	Type string `json:"type,omitempty"`
	UseSenderFlag bool `json:"useSenderFlag,omitempty"`
}

type ServiceLocation struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Where string `json:"where,omitempty"`
}

type ServiceLocationInfo struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceNote struct {
	Info interface{} `json:"_info,omitempty"`
	Contact ContactReference `json:"contact,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	CustomerUpdatedFlag bool `json:"customerUpdatedFlag,omitempty"`
	DateCreated string `json:"dateCreated,omitempty"`
	DetailDescriptionFlag bool `json:"detailDescriptionFlag,omitempty"`
	ExternalFlag bool `json:"externalFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InternalAnalysisFlag bool `json:"internalAnalysisFlag,omitempty"`
	InternalFlag bool `json:"internalFlag,omitempty"`
	IssueFlag bool `json:"issueFlag,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	ProcessNotifications bool `json:"processNotifications,omitempty"`
	ResolutionFlag bool `json:"resolutionFlag,omitempty"`
	SentimentScore float64 `json:"sentimentScore,omitempty"`
	Text string `json:"text,omitempty"`
	TicketId int `json:"ticketId,omitempty"`
}

type ServiceSignoff struct {
	Info interface{} `json:"_info,omitempty"`
	BillingMethodsText string `json:"billingMethodsText,omitempty"`
	BillingMethodsTextFlag bool `json:"billingMethodsTextFlag,omitempty"`
	BillingTermsFlag bool `json:"billingTermsFlag,omitempty"`
	CompanyInfoFlag bool `json:"companyInfoFlag,omitempty"`
	ConfigurationsFlag bool `json:"configurationsFlag,omitempty"`
	CreditCardFieldsFlag bool `json:"creditCardFieldsFlag,omitempty"`
	CustomerSignoffFieldsFlag bool `json:"customerSignoffFieldsFlag,omitempty"`
	CustomerSignoffText string `json:"customerSignoffText,omitempty"`
	CustomerSignoffTextFlag bool `json:"customerSignoffTextFlag,omitempty"`
	DefaultFFFlag bool `json:"defaultFFFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	DiscussionFlag bool `json:"discussionFlag,omitempty"`
	ExpenseAgreementFlag bool `json:"expenseAgreementFlag,omitempty"`
	ExpenseAmountFlag bool `json:"expenseAmountFlag,omitempty"`
	ExpenseBillFlag bool `json:"expenseBillFlag,omitempty"`
	ExpenseDateFlag bool `json:"expenseDateFlag,omitempty"`
	ExpenseFlag bool `json:"expenseFlag,omitempty"`
	ExpenseManualEntry int `json:"expenseManualEntry,omitempty"`
	ExpenseManualFlag bool `json:"expenseManualFlag,omitempty"`
	ExpenseMemberFlag bool `json:"expenseMemberFlag,omitempty"`
	ExpenseNotesFlag bool `json:"expenseNotesFlag,omitempty"`
	ExpenseTaxFlag bool `json:"expenseTaxFlag,omitempty"`
	ExpenseTypeFlag bool `json:"expenseTypeFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InternalNotesFlag bool `json:"internalNotesFlag,omitempty"`
	Name string `json:"name"`
	ProductAgreementFlag bool `json:"productAgreementFlag,omitempty"`
	ProductBillFlag bool `json:"productBillFlag,omitempty"`
	ProductDescriptionFlag bool `json:"productDescriptionFlag,omitempty"`
	ProductExtendedAmountFlag bool `json:"productExtendedAmountFlag,omitempty"`
	ProductFlag bool `json:"productFlag,omitempty"`
	ProductManualEntry int `json:"productManualEntry,omitempty"`
	ProductManualFlag bool `json:"productManualFlag,omitempty"`
	ProductPriceFlag bool `json:"productPriceFlag,omitempty"`
	ProductQuantityFlag bool `json:"productQuantityFlag,omitempty"`
	ProductTaxFlag bool `json:"productTaxFlag,omitempty"`
	ResolutionFlag bool `json:"resolutionFlag,omitempty"`
	SummaryFlag bool `json:"summaryFlag,omitempty"`
	Task string `json:"task,omitempty"`
	TaskFlag bool `json:"taskFlag,omitempty"`
	TechnicianSignoffFlag bool `json:"technicianSignoffFlag,omitempty"`
	TimeAgreementFlag bool `json:"timeAgreementFlag,omitempty"`
	TimeBillFlag bool `json:"timeBillFlag,omitempty"`
	TimeDateFlag bool `json:"timeDateFlag,omitempty"`
	TimeExtendedAmountFlag bool `json:"timeExtendedAmountFlag,omitempty"`
	TimeFlag bool `json:"timeFlag,omitempty"`
	TimeHoursFlag bool `json:"timeHoursFlag,omitempty"`
	TimeManualEntry int `json:"timeManualEntry,omitempty"`
	TimeManualFlag bool `json:"timeManualFlag,omitempty"`
	TimeMemberFlag bool `json:"timeMemberFlag,omitempty"`
	TimeNotesFlag bool `json:"timeNotesFlag,omitempty"`
	TimeRateFlag bool `json:"timeRateFlag,omitempty"`
	TimeStartEndFlag bool `json:"timeStartEndFlag,omitempty"`
	TimeTaxFlag bool `json:"timeTaxFlag,omitempty"`
	TimeWorkTypeFlag bool `json:"timeWorkTypeFlag,omitempty"`
	VisibleLogoFlag bool `json:"visibleLogoFlag,omitempty"`
}

type ServiceSignoffCustomField struct {
	Info interface{} `json:"_info,omitempty"`
	DisplaySection string `json:"displaySection,omitempty"`
	ID int `json:"id,omitempty"`
	SequenceNumber float64 `json:"sequenceNumber,omitempty"`
	UserDefinedField UserDefinedFieldReference `json:"userDefinedField"`
}

type ServiceSignoffInfo struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceSurvey struct {
	Info interface{} `json:"_info,omitempty"`
	FooterText string `json:"footerText,omitempty"`
	FooterTextVisibleFlag bool `json:"footerTextVisibleFlag,omitempty"`
	HeaderIncludeLogoFlag bool `json:"headerIncludeLogoFlag,omitempty"`
	HeaderText string `json:"headerText,omitempty"`
	HeaderTextVisibleFlag bool `json:"headerTextVisibleFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name"`
	NotifyMember MemberReference `json:"notifyMember,omitempty"`
	NotifyWho GenericIdIdentifierReference `json:"notifyWho,omitempty"`
	NotifyWhoVisibleFlag bool `json:"notifyWhoVisibleFlag,omitempty"`
	ThankYouText string `json:"thankYouText,omitempty"`
}

type ServiceSurveyQuestion struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	IncludeFlag bool `json:"includeFlag,omitempty"`
	NoAnswerPoints int `json:"noAnswerPoints,omitempty"`
	Options []ServiceSurveyQuestionOption `json:"options,omitempty"`
	Question string `json:"question"`
	RequiredFlag bool `json:"requiredFlag,omitempty"`
	SequenceNumber int `json:"sequenceNumber,omitempty"`
	SurveyId int `json:"surveyId,omitempty"`
	Type string `json:"type,omitempty"`
}

type ServiceTask struct {
	Info interface{} `json:"_info,omitempty"`
	ChildScheduleAction string `json:"childScheduleAction,omitempty"`
	ChildTicketId int `json:"childTicketId,omitempty"`
	ClosedFlag bool `json:"closedFlag,omitempty"`
	Code ServiceCodeReference `json:"code,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Notes string `json:"notes,omitempty"`
	Priority int `json:"priority,omitempty"`
	Resolution string `json:"resolution,omitempty"`
	Schedule ScheduleEntryReference `json:"schedule,omitempty"`
	Summary string `json:"summary,omitempty"`
	TicketId int `json:"ticketId,omitempty"`
}

type ServiceTeam struct {
	Info interface{} `json:"_info,omitempty"`
	DeleteNotifyFlag bool `json:"deleteNotifyFlag,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	Leader MemberReference `json:"leader,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceTemplate struct {
	Info interface{} `json:"_info,omitempty"`
	Agreement AgreementReference `json:"agreement,omitempty"`
	AssignedBy MemberReference `json:"assignedBy,omitempty"`
	AssignedNotifyFlag bool `json:"assignedNotifyFlag,omitempty"`
	AttachScheduleToNewServiceFlag bool `json:"attachScheduleToNewServiceFlag,omitempty"`
	BillCompleteFlag bool `json:"billComplete_Flag,omitempty"`
	BillServiceSeparatelyFlag bool `json:"billServiceSeparatelyFlag,omitempty"`
	BillUnapprovedTimeAndExpensesFlag bool `json:"billUnapprovedTimeAndExpensesFlag,omitempty"`
	BillingAmount float64 `json:"billingAmount,omitempty"`
	BillingMethod string `json:"billingMethod,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	Contact ContactReference `json:"contact,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	EmailCC string `json:"emailCC,omitempty"`
	EmailCCFlag bool `json:"emailCCFlag,omitempty"`
	EmailContactFlag bool `json:"emailContactFlag,omitempty"`
	EmailResourceFlag bool `json:"emailResourceFlag,omitempty"`
	ExpenseBillableFlag bool `json:"expenseBillableFlag,omitempty"`
	ExpenseInvoiceFlag bool `json:"expenseInvoiceFlag,omitempty"`
	HoursBudget float64 `json:"hoursBudget,omitempty"`
	ID int `json:"id,omitempty"`
	Impact string `json:"impact,omitempty"`
	InternalAnalysis string `json:"internalAnalysis,omitempty"`
	Item ServiceItemReference `json:"item,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
	OverrideFlag bool `json:"overrideFlag,omitempty"`
	Priority PriorityReference `json:"priority,omitempty"`
	Problem string `json:"problem,omitempty"`
	ProductInvoiceFlag bool `json:"productInvoiceFlag,omitempty"`
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`
	Reference string `json:"reference,omitempty"`
	RestrictDownpaymentFlag bool `json:"restrictDownpaymentFlag,omitempty"`
	ScheduleDaysBefore int `json:"scheduleDaysBefore,omitempty"`
	ServiceDaysBefore int `json:"serviceDaysBefore,omitempty"`
	ServiceLocation ServiceLocationReference `json:"serviceLocation,omitempty"`
	Severity string `json:"severity,omitempty"`
	Site SiteReference `json:"site,omitempty"`
	Source ServiceSourceReference `json:"source,omitempty"`
	Status ServiceStatusReference `json:"status,omitempty"`
	Subtype ServiceSubTypeReference `json:"subtype,omitempty"`
	Summary string `json:"summary,omitempty"`
	Team ServiceTeamReference `json:"team,omitempty"`
	TemplateFlag bool `json:"templateFlag,omitempty"`
	TimeBillableFlag bool `json:"timeBillableFlag,omitempty"`
	TimeInvoiceFlag bool `json:"timeInvoiceFlag,omitempty"`
	Type ServiceTypeReference `json:"type,omitempty"`
}

type ServiceTemplateInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TemplateFlag bool `json:"templateFlag,omitempty"`
}

type ServiceTicketLink struct {
	Info interface{} `json:"_info,omitempty"`
	EnabledFlag bool `json:"enabledFlag,omitempty"`
	ID int `json:"id,omitempty"`
	LinkText string `json:"linkText"`
	Name string `json:"name"`
	URL string `json:"url"`
}

type ServiceTicketLinkInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	LinkText string `json:"linkText,omitempty"`
	Name string `json:"name,omitempty"`
	URL string `json:"url,omitempty"`
}

type ServiceTicketNote struct {
	Info interface{} `json:"_info,omitempty"`
	BundledFlag bool `json:"bundledFlag,omitempty"`
	Contact ContactReference `json:"contact,omitempty"`
	CreatedByParentFlag bool `json:"createdByParentFlag,omitempty"`
	DetailDescriptionFlag bool `json:"detailDescriptionFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InternalAnalysisFlag bool `json:"internalAnalysisFlag,omitempty"`
	IsMarkdownFlag bool `json:"isMarkdownFlag,omitempty"`
	IssueFlag bool `json:"issueFlag,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	MergedFlag bool `json:"mergedFlag,omitempty"`
	NoteType string `json:"noteType,omitempty"`
	OriginalAuthor string `json:"originalAuthor,omitempty"`
	ResolutionFlag bool `json:"resolutionFlag,omitempty"`
	Text string `json:"text,omitempty"`
	Ticket TicketReference `json:"ticket,omitempty"`
	TimeEnd string `json:"timeEnd,omitempty"`
	TimeStart string `json:"timeStart,omitempty"`
}

type Severity struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	Description string `json:"description"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Source struct {
	Info interface{} `json:"_info,omitempty"`
	DateEntered time.Time `json:"dateEntered,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	EnteredBy string `json:"enteredBy,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
}

type SourceInfo struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SurveyOption struct {
	Info interface{} `json:"_info,omitempty"`
	Caption string `json:"caption"`
	ID int `json:"id,omitempty"`
	Points int `json:"points,omitempty"`
	Visibleflag bool `json:"visibleflag,omitempty"`
}

type SurveyResult struct {
	Info interface{} `json:"_info,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	Contact ContactReference `json:"contact,omitempty"`
	ContactMeFlag bool `json:"contactMeFlag,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	FooterResponse string `json:"footerResponse,omitempty"`
	ID int `json:"id,omitempty"`
	Results []SurveyResultDetail `json:"results,omitempty"`
	SurveyId int `json:"surveyId,omitempty"`
	TicketId int `json:"ticketId,omitempty"`
	TotalPoints int `json:"totalPoints,omitempty"`
}

type TeamMember struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member"`
	Team ServiceTeamReference `json:"team"`
	TeamLeaderFlag bool `json:"teamLeaderFlag,omitempty"`
}

type TemplateGeneratedCountsModel struct {
	ScheduleCount int `json:"scheduleCount,omitempty"`
	ServiceCount int `json:"serviceCount,omitempty"`
}

type Ticket struct {
	Info interface{} `json:"_info,omitempty"`
	ActualHours float64 `json:"actualHours,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	Agreement AgreementReference `json:"agreement,omitempty"`
	AgreementType string `json:"agreementType,omitempty"`
	AllowAllClientsPortalView bool `json:"allowAllClientsPortalView,omitempty"`
	Approved bool `json:"approved,omitempty"`
	AutomaticEmailCc string `json:"automaticEmailCc,omitempty"`
	AutomaticEmailCcFlag bool `json:"automaticEmailCcFlag,omitempty"`
	AutomaticEmailContactFlag bool `json:"automaticEmailContactFlag,omitempty"`
	AutomaticEmailResourceFlag bool `json:"automaticEmailResourceFlag,omitempty"`
	BillExpenses string `json:"billExpenses,omitempty"`
	BillProducts string `json:"billProducts,omitempty"`
	BillTime string `json:"billTime,omitempty"`
	BillingAmount float64 `json:"billingAmount,omitempty"`
	BillingMethod string `json:"billingMethod,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	BudgetHours float64 `json:"budgetHours,omitempty"`
	City string `json:"city,omitempty"`
	ClosedBy string `json:"closedBy,omitempty"`
	ClosedDate string `json:"closedDate,omitempty"`
	ClosedFlag bool `json:"closedFlag,omitempty"`
	Company CompanyReference `json:"company"`
	Contact ContactReference `json:"contact,omitempty"`
	ContactEmailAddress string `json:"contactEmailAddress,omitempty"`
	ContactEmailLookup string `json:"contactEmailLookup,omitempty"`
	ContactName string `json:"contactName,omitempty"`
	ContactPhoneExtension string `json:"contactPhoneExtension,omitempty"`
	ContactPhoneNumber string `json:"contactPhoneNumber,omitempty"`
	Country CountryReference `json:"country,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
	CustomerUpdatedFlag bool `json:"customerUpdatedFlag,omitempty"`
	DateResolved string `json:"dateResolved,omitempty"`
	DateResplan string `json:"dateResplan,omitempty"`
	DateResponded string `json:"dateResponded,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	Duration int `json:"duration,omitempty"`
	EscalationLevel int `json:"escalationLevel,omitempty"`
	EscalationStartDateUTC string `json:"escalationStartDateUTC,omitempty"`
	EstimatedExpenseCost float64 `json:"estimatedExpenseCost,omitempty"`
	EstimatedExpenseRevenue float64 `json:"estimatedExpenseRevenue,omitempty"`
	EstimatedProductCost float64 `json:"estimatedProductCost,omitempty"`
	EstimatedProductRevenue float64 `json:"estimatedProductRevenue,omitempty"`
	EstimatedStartDate time.Time `json:"estimatedStartDate,omitempty"`
	EstimatedTimeCost float64 `json:"estimatedTimeCost,omitempty"`
	EstimatedTimeRevenue float64 `json:"estimatedTimeRevenue,omitempty"`
	ExternalXRef string `json:"externalXRef,omitempty"`
	HasChildTicket bool `json:"hasChildTicket,omitempty"`
	HasMergedChildTicketFlag bool `json:"hasMergedChildTicketFlag,omitempty"`
	HourlyRate float64 `json:"hourlyRate,omitempty"`
	ID int `json:"id,omitempty"`
	Impact string `json:"impact,omitempty"`
	InitialDescription string `json:"initialDescription,omitempty"`
	InitialDescriptionFrom string `json:"initialDescriptionFrom,omitempty"`
	InitialInternalAnalysis string `json:"initialInternalAnalysis,omitempty"`
	InitialResolution string `json:"initialResolution,omitempty"`
	IntegratorTags []string `json:"integratorTags,omitempty"`
	IsInSla bool `json:"isInSla,omitempty"`
	Item ServiceItemReference `json:"item,omitempty"`
	KnowledgeBaseCategoryId int `json:"knowledgeBaseCategoryId,omitempty"`
	KnowledgeBaseLinkId int `json:"knowledgeBaseLinkId,omitempty"`
	KnowledgeBaseLinkType string `json:"knowledgeBaseLinkType,omitempty"`
	KnowledgeBaseSubCategoryId int `json:"knowledgeBaseSubCategoryId,omitempty"`
	LagDays int `json:"lagDays,omitempty"`
	LagNonworkingDaysFlag bool `json:"lagNonworkingDaysFlag,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	MergedParentTicket TicketReference `json:"mergedParentTicket,omitempty"`
	MinutesBeforeWaiting int `json:"minutesBeforeWaiting,omitempty"`
	MinutesWaiting int `json:"minutesWaiting,omitempty"`
	MobileGuid string `json:"mobileGuid,omitempty"`
	Opportunity OpportunityReference `json:"opportunity,omitempty"`
	Owner MemberReference `json:"owner,omitempty"`
	ParentTicketId int `json:"parentTicketId,omitempty"`
	PoNumber string `json:"poNumber,omitempty"`
	PredecessorClosedFlag bool `json:"predecessorClosedFlag,omitempty"`
	PredecessorId int `json:"predecessorId,omitempty"`
	PredecessorType string `json:"predecessorType,omitempty"`
	Priority PriorityReference `json:"priority,omitempty"`
	ProcessNotifications bool `json:"processNotifications,omitempty"`
	RecordType string `json:"recordType,omitempty"`
	RequestForChangeFlag bool `json:"requestForChangeFlag,omitempty"`
	RequiredDate time.Time `json:"requiredDate,omitempty"`
	ResPlanMinutes int `json:"resPlanMinutes,omitempty"`
	ResolutionHours float64 `json:"resolutionHours,omitempty"`
	ResolveMinutes int `json:"resolveMinutes,omitempty"`
	ResolvedBy string `json:"resolvedBy,omitempty"`
	Resources string `json:"resources,omitempty"`
	ResplanBy string `json:"resplanBy,omitempty"`
	ResplanHours float64 `json:"resplanHours,omitempty"`
	ResplanSkippedMinutes int `json:"resplanSkippedMinutes,omitempty"`
	RespondMinutes int `json:"respondMinutes,omitempty"`
	RespondedBy string `json:"respondedBy,omitempty"`
	RespondedHours float64 `json:"respondedHours,omitempty"`
	RespondedSkippedMinutes int `json:"respondedSkippedMinutes,omitempty"`
	ServiceLocation ServiceLocationReference `json:"serviceLocation,omitempty"`
	Severity string `json:"severity,omitempty"`
	Site SiteReference `json:"site,omitempty"`
	SiteName string `json:"siteName,omitempty"`
	SkipCallback bool `json:"skipCallback,omitempty"`
	Sla SLAReference `json:"sla,omitempty"`
	SlaStatus string `json:"slaStatus,omitempty"`
	Source ServiceSourceReference `json:"source,omitempty"`
	StateIdentifier string `json:"stateIdentifier,omitempty"`
	Status ServiceStatusReference `json:"status,omitempty"`
	SubBillingAmount float64 `json:"subBillingAmount,omitempty"`
	SubBillingMethod string `json:"subBillingMethod,omitempty"`
	SubDateAccepted string `json:"subDateAccepted,omitempty"`
	SubType ServiceSubTypeReference `json:"subType,omitempty"`
	Summary string `json:"summary"`
	Team ServiceTeamReference `json:"team,omitempty"`
	Type ServiceTypeReference `json:"type,omitempty"`
	WorkRole WorkRoleReference `json:"workRole,omitempty"`
	WorkType WorkTypeReference `json:"workType,omitempty"`
	Zip string `json:"zip,omitempty"`
}

type TicketBundle struct {
	ChildTicketIds []int `json:"childTicketIds,omitempty"`
}

type TicketChangeLog struct {
	Info interface{} `json:"_info,omitempty"`
	Action string `json:"action,omitempty"`
	BoardId int `json:"boardId,omitempty"`
	BoardName string `json:"boardName,omitempty"`
	ClosedFlag bool `json:"closedFlag,omitempty"`
	CompanyIdentifier int `json:"companyIdentifier,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	ContactId int `json:"contactId,omitempty"`
	ContactName string `json:"contactName,omitempty"`
	CustomerUpdatedFlag bool `json:"customerUpdatedFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Impact string `json:"impact,omitempty"`
	MergedParentTicketId int `json:"mergedParentTicketId,omitempty"`
	OwnerIdentifier int `json:"ownerIdentifier,omitempty"`
	ParentTicketId int `json:"parentTicketId,omitempty"`
	PartnerId string `json:"partnerId,omitempty"`
	PriorityId int `json:"priorityId,omitempty"`
	PriorityLevel string `json:"priorityLevel,omitempty"`
	PriorityName string `json:"priorityName,omitempty"`
	PrioritySort int `json:"prioritySort,omitempty"`
	ProcessingStatus string `json:"processingStatus,omitempty"`
	ProductInstanceId string `json:"productInstanceId,omitempty"`
	RecordType string `json:"recordType,omitempty"`
	ResourceList string `json:"resourceList,omitempty"`
	Severity string `json:"severity,omitempty"`
	SlaName string `json:"slaName,omitempty"`
	SlaStatus string `json:"slaStatus,omitempty"`
	Status string `json:"status,omitempty"`
	Summary string `json:"summary,omitempty"`
	TeamName string `json:"teamName,omitempty"`
	TicketNumber int `json:"ticketNumber,omitempty"`
	TicketOwner string `json:"ticketOwner,omitempty"`
}

type TicketInfo struct {
	Info interface{} `json:"_info,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	ID int `json:"id,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type TicketMerge struct {
	MergeTicketIds []int `json:"mergeTicketIds"`
	Status ServiceStatusReference `json:"status"`
}

type TicketSync struct {
	Info interface{} `json:"_info,omitempty"`
	Company CompanyReference `json:"company"`
	ID int `json:"id,omitempty"`
	IntegratorLogin IntegratorLoginReference `json:"integratorLogin"`
	InternalAnalysisFlag bool `json:"internalAnalysisFlag,omitempty"`
	Name string `json:"name"`
	Password string `json:"password,omitempty"`
	ProblemDescriptionFlag bool `json:"problemDescriptionFlag,omitempty"`
	Psg string `json:"psg,omitempty"`
	ResolutionFlag bool `json:"resolutionFlag,omitempty"`
	URL string `json:"url"`
	UserName string `json:"userName,omitempty"`
	VendorType string `json:"vendorType,omitempty"`
}

