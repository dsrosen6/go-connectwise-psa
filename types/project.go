package types

import (
	"time"
)

type ConvertItem struct {
	ID         int          `json:"id,omitempty"`
	Phase      ProjectPhase `json:"phase,omitempty"`
	Project    Project      `json:"project,omitempty"`
	RecordType string       `json:"recordType,omitempty"`
	WbsCode    string       `json:"wbsCode,omitempty"`
}

type Document struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type PhaseStatus struct {
	Info                      map[string]interface{} `json:"_info,omitempty"`
	BoardAssociationIds       []int                  `json:"boardAssociationIds,omitempty"`
	ClosedFlag                bool                   `json:"closedFlag,omitempty"`
	CollapsedFlag             bool                   `json:"collapsedFlag,omitempty"`
	CustomStatusIndicatorName string                 `json:"customStatusIndicatorName,omitempty"`
	DefaultFlag               bool                   `json:"defaultFlag,omitempty"`
	ID                        int                    `json:"id,omitempty"`
	InactiveFlag              bool                   `json:"inactiveFlag,omitempty"`
	Name                      string                 `json:"name"`
	StatusIndicator           StatusIndicator        `json:"statusIndicator,omitempty"`
}

type PhaseStatusInfo struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	BoardAssociationIds []int                  `json:"boardAssociationIds,omitempty"`
	ClosedFlag          bool                   `json:"closedFlag,omitempty"`
	CollapsedFlag       bool                   `json:"collapsedFlag,omitempty"`
	DefaultFlag         bool                   `json:"defaultFlag,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	InactiveFlag        bool                   `json:"inactiveFlag,omitempty"`
	Name                string                 `json:"name,omitempty"`
}

type Project struct {
	Info                         map[string]interface{} `json:"_info,omitempty"`
	ActualEnd                    time.Time              `json:"actualEnd,omitempty"`
	ActualHours                  float64                `json:"actualHours,omitempty"`
	ActualStart                  time.Time              `json:"actualStart,omitempty"`
	Agreement                    Agreement              `json:"agreement,omitempty"`
	BillExpenses                 string                 `json:"billExpenses,omitempty"`
	BillProducts                 string                 `json:"billProducts,omitempty"`
	BillProjectAfterClosedFlag   bool                   `json:"billProjectAfterClosedFlag,omitempty"`
	BillTime                     string                 `json:"billTime,omitempty"`
	BillToCompany                Company                `json:"billToCompany,omitempty"`
	BillToContact                Contact                `json:"billToContact,omitempty"`
	BillToSite                   Site                   `json:"billToSite,omitempty"`
	BillUnapprovedTimeAndExpense bool                   `json:"billUnapprovedTimeAndExpense,omitempty"`
	BillingAmount                float64                `json:"billingAmount,omitempty"`
	BillingAttention             string                 `json:"billingAttention,omitempty"`
	BillingMethod                string                 `json:"billingMethod,omitempty"`
	BillingRateType              string                 `json:"billingRateType,omitempty"`
	BillingStartDate             time.Time              `json:"billingStartDate,omitempty"`
	BillingTerms                 BillingTerms           `json:"billingTerms,omitempty"`
	Board                        ProjectBoard           `json:"board"`
	BudgetAnalysis               string                 `json:"budgetAnalysis,omitempty"`
	BudgetFlag                   bool                   `json:"budgetFlag,omitempty"`
	BudgetHours                  float64                `json:"budgetHours,omitempty"`
	ClosedFlag                   bool                   `json:"closedFlag,omitempty"`
	Company                      Company                `json:"company"`
	CompanyLocation              SystemLocation         `json:"companyLocation,omitempty"`
	Contact                      Contact                `json:"contact,omitempty"`
	Currency                     Currency               `json:"currency,omitempty"`
	CustomFields                 []CustomFieldValue     `json:"customFields,omitempty"`
	CustomerPO                   string                 `json:"customerPO,omitempty"`
	Department                   SystemDepartment       `json:"department,omitempty"`
	Description                  string                 `json:"description,omitempty"`
	DoNotDisplayInPortalFlag     bool                   `json:"doNotDisplayInPortalFlag,omitempty"`
	Downpayment                  float64                `json:"downpayment,omitempty"`
	EstimatedEnd                 time.Time              `json:"estimatedEnd"`
	EstimatedExpenseCost         float64                `json:"estimatedExpenseCost,omitempty"`
	EstimatedExpenseRevenue      float64                `json:"estimatedExpenseRevenue,omitempty"`
	EstimatedHours               float64                `json:"estimatedHours,omitempty"`
	EstimatedProductCost         float64                `json:"estimatedProductCost,omitempty"`
	EstimatedProductRevenue      float64                `json:"estimatedProductRevenue,omitempty"`
	EstimatedStart               time.Time              `json:"estimatedStart"`
	EstimatedTimeCost            float64                `json:"estimatedTimeCost,omitempty"`
	EstimatedTimeRevenue         float64                `json:"estimatedTimeRevenue,omitempty"`
	ExpenseApprover              Member                 `json:"expenseApprover,omitempty"`
	ID                           int                    `json:"id,omitempty"`
	IncludeDependenciesFlag      bool                   `json:"includeDependenciesFlag,omitempty"`
	IncludeEstimatesFlag         bool                   `json:"includeEstimatesFlag,omitempty"`
	Location                     SystemLocation         `json:"location,omitempty"`
	Manager                      Member                 `json:"manager,omitempty"`
	Name                         string                 `json:"name"`
	Opportunity                  Opportunity            `json:"opportunity,omitempty"`
	PercentComplete              float64                `json:"percentComplete,omitempty"`
	PoAmount                     float64                `json:"poAmount,omitempty"`
	ProjectTemplateId            int                    `json:"projectTemplateId,omitempty"`
	RestrictDownPaymentFlag      bool                   `json:"restrictDownPaymentFlag,omitempty"`
	ScheduledEnd                 time.Time              `json:"scheduledEnd,omitempty"`
	ScheduledHours               float64                `json:"scheduledHours,omitempty"`
	ScheduledStart               time.Time              `json:"scheduledStart,omitempty"`
	ShipToCompany                Company                `json:"shipToCompany,omitempty"`
	ShipToContact                Contact                `json:"shipToContact,omitempty"`
	ShipToSite                   Site                   `json:"shipToSite,omitempty"`
	Site                         Site                   `json:"site,omitempty"`
	Status                       ProjectStatus          `json:"status,omitempty"`
	TaxCode                      TaxCode                `json:"taxCode,omitempty"`
	TimeApprover                 Member                 `json:"timeApprover,omitempty"`
	Type                         ProjectType            `json:"type,omitempty"`
}

type ProjectBillingRate struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	ActivityClassRecId int                    `json:"activityClassRecId,omitempty"`
	HourlyRate         float64                `json:"hourlyRate,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	Member             Member                 `json:"member,omitempty"`
	MemberRecId        int                    `json:"memberRecId,omitempty"`
	ProjectRecId       int                    `json:"projectRecId,omitempty"`
	WorkRole           WorkRole               `json:"workRole,omitempty"`
}

type ProjectBoardKanbanSetting struct {
	Color       string                     `json:"color,omitempty"`
	ID          int                        `json:"id,omitempty"`
	LastUpdated string                     `json:"lastUpdated,omitempty"`
	Name        string                     `json:"name"`
	Order       int                        `json:"order,omitempty"`
	Statuses    []ProjectBoardKanbanStatus `json:"statuses,omitempty"`
	UpdatedBy   string                     `json:"updatedBy,omitempty"`
}

type ProjectBoardTeam struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	Department  SystemDepartment       `json:"department,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Location    SystemLocation         `json:"location,omitempty"`
	Name        string                 `json:"name"`
}

type ProjectBoardTeamInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ProjectBoardTeamMember struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Member      Member                 `json:"member"`
	ProjectRole ProjectRole            `json:"projectRole"`
	WorkRole    WorkRole               `json:"workRole,omitempty"`
}

type ProjectContact struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	Contact   Contact                `json:"contact"`
	ID        int                    `json:"id,omitempty"`
	ProjectId int                    `json:"projectId,omitempty"`
}

type ProjectNote struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	Flagged   bool                   `json:"flagged,omitempty"`
	ID        int                    `json:"id,omitempty"`
	ProjectId int                    `json:"projectId,omitempty"`
	Text      string                 `json:"text"`
	Type      NoteType               `json:"type,omitempty"`
}

type ProjectPhase struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	ActualEnd               string                 `json:"actualEnd,omitempty"`
	ActualHours             float64                `json:"actualHours,omitempty"`
	ActualStart             string                 `json:"actualStart,omitempty"`
	Agreement               Agreement              `json:"agreement,omitempty"`
	BillExpenses            string                 `json:"billExpenses,omitempty"`
	BillPhaseClosedFlag     bool                   `json:"billPhaseClosedFlag,omitempty"`
	BillProducts            string                 `json:"billProducts,omitempty"`
	BillProjectClosedFlag   bool                   `json:"billProjectClosedFlag,omitempty"`
	BillSeparatelyFlag      bool                   `json:"billSeparatelyFlag,omitempty"`
	BillTime                string                 `json:"billTime,omitempty"`
	BillToCompany           Company                `json:"billToCompany,omitempty"`
	BillToContact           Contact                `json:"billToContact,omitempty"`
	BillToSite              Site                   `json:"billToSite,omitempty"`
	BillingMethod           string                 `json:"billingMethod,omitempty"`
	BillingStartDate        time.Time              `json:"billingStartDate,omitempty"`
	BillingTerms            BillingTerms           `json:"billingTerms,omitempty"`
	Board                   ProjectBoard           `json:"board,omitempty"`
	BudgetHours             float64                `json:"budgetHours,omitempty"`
	BusinessUnitId          int                    `json:"businessUnitId,omitempty"`
	Currency                Currency               `json:"currency,omitempty"`
	CustomFields            []CustomFieldValue     `json:"customFields,omitempty"`
	DeadlineDate            time.Time              `json:"deadlineDate,omitempty"`
	Department              BillingUnit            `json:"department,omitempty"`
	Description             string                 `json:"description"`
	Downpayment             float64                `json:"downpayment,omitempty"`
	EndDate                 string                 `json:"endDate,omitempty"`
	EstimatedExpenseCost    float64                `json:"estimatedExpenseCost,omitempty"`
	EstimatedExpenseRevenue float64                `json:"estimatedExpenseRevenue,omitempty"`
	EstimatedProductCost    float64                `json:"estimatedProductCost,omitempty"`
	EstimatedProductRevenue float64                `json:"estimatedProductRevenue,omitempty"`
	EstimatedTimeCost       float64                `json:"estimatedTimeCost,omitempty"`
	EstimatedTimeRevenue    float64                `json:"estimatedTimeRevenue,omitempty"`
	HourlyRate              float64                `json:"hourlyRate,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	LocationId              int                    `json:"locationId,omitempty"`
	MarkAsMilestoneFlag     bool                   `json:"markAsMilestoneFlag,omitempty"`
	Notes                   string                 `json:"notes,omitempty"`
	Opportunity             Opportunity            `json:"opportunity,omitempty"`
	ParentPhase             *ProjectPhase          `json:"parentPhase,omitempty"`
	PoAmount                float64                `json:"poAmount,omitempty"`
	PoNumber                string                 `json:"poNumber,omitempty"`
	ProjectId               int                    `json:"projectId,omitempty"`
	ScheduledEnd            string                 `json:"scheduledEnd,omitempty"`
	ScheduledHours          float64                `json:"scheduledHours,omitempty"`
	ScheduledStart          string                 `json:"scheduledStart,omitempty"`
	ShipToCompany           Company                `json:"shipToCompany,omitempty"`
	ShipToContact           Contact                `json:"shipToContact,omitempty"`
	ShipToSite              Site                   `json:"shipToSite,omitempty"`
	StartDate               string                 `json:"startDate,omitempty"`
	Status                  PhaseStatus            `json:"status,omitempty"`
	TaxCode                 TaxCode                `json:"taxCode,omitempty"`
	WbsCode                 string                 `json:"wbsCode,omitempty"`
}

type ProjectSecurityRole struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	DefaultContactFlag bool                   `json:"defaultContactFlag,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	ManagerRoleFlag    bool                   `json:"managerRoleFlag,omitempty"`
	Name               string                 `json:"name"`
}

type ProjectSecurityRoleInfo struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	DefaultContactFlag bool                   `json:"defaultContactFlag,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	ManagerRoleFlag    bool                   `json:"managerRoleFlag,omitempty"`
	Name               string                 `json:"name,omitempty"`
}

type ProjectSecurityRoleSetting struct {
	Info             map[string]interface{} `json:"_info,omitempty"`
	AddLevel         string                 `json:"addLevel,omitempty"`
	DeleteLevel      string                 `json:"deleteLevel,omitempty"`
	EditLevel        string                 `json:"editLevel,omitempty"`
	ID               int                    `json:"id,omitempty"`
	InquireLevel     string                 `json:"inquireLevel,omitempty"`
	ModuleIdentifier string                 `json:"moduleIdentifier,omitempty"`
	MyFlag           bool                   `json:"myFlag,omitempty"`
}

type ProjectStatus struct {
	Info                      map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag                bool                   `json:"closedFlag,omitempty"`
	CustomStatusIndicatorName string                 `json:"customStatusIndicatorName,omitempty"`
	DefaultFlag               bool                   `json:"defaultFlag,omitempty"`
	ID                        int                    `json:"id,omitempty"`
	InactiveFlag              bool                   `json:"inactiveFlag,omitempty"`
	Name                      string                 `json:"name"`
	NoTimeFlag                bool                   `json:"noTimeFlag,omitempty"`
	StatusIndicator           StatusIndicator        `json:"statusIndicator,omitempty"`
}

type ProjectStatusInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag   bool                   `json:"closedFlag,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
	NoTimeFlag   bool                   `json:"noTimeFlag,omitempty"`
}

type ProjectTeamMember struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	EndDate     time.Time              `json:"endDate,omitempty"`
	Hours       float64                `json:"hours,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Member      Member                 `json:"member"`
	ProjectId   int                    `json:"projectId,omitempty"`
	ProjectRole ProjectRole            `json:"projectRole"`
	StartDate   time.Time              `json:"startDate,omitempty"`
	WorkRole    WorkRole               `json:"workRole,omitempty"`
}

type ProjectTemplate struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	ConnectWiseId string                 `json:"connectWiseId,omitempty"`
	Description   string                 `json:"description,omitempty"`
	ID            int                    `json:"id,omitempty"`
	Name          string                 `json:"name"`
	Type          ProjectType            `json:"type,omitempty"`
}

type ProjectTemplatePhase struct {
	Info                  map[string]interface{} `json:"_info,omitempty"`
	BudgetHours           float64                `json:"budgetHours,omitempty"`
	ConnectWiseId         string                 `json:"connectWiseId,omitempty"`
	Description           string                 `json:"description,omitempty"`
	ID                    int                    `json:"id,omitempty"`
	MarkAsMilestone       bool                   `json:"markAsMilestone,omitempty"`
	Notes                 string                 `json:"notes,omitempty"`
	ParentConnectWiseId   string                 `json:"parentConnectWiseId,omitempty"`
	ParentId              int                    `json:"parentId,omitempty"`
	ParentPhase           int                    `json:"parentPhase,omitempty"`
	PhaseBilledSeparately bool                   `json:"phaseBilledSeparately,omitempty"`
	TemplateRecId         int                    `json:"templateRecId,omitempty"`
	WbsCode               string                 `json:"wbsCode,omitempty"`
}

type ProjectTemplateTask struct {
	Info                     map[string]interface{} `json:"_info,omitempty"`
	Code                     ServiceCode            `json:"code,omitempty"`
	ConnectWiseId            string                 `json:"connectWiseId,omitempty"`
	Description              string                 `json:"description,omitempty"`
	GrandParentConnectWiseId string                 `json:"grandParentConnectWiseId,omitempty"`
	GrandParentId            int                    `json:"grandParentId,omitempty"`
	ID                       int                    `json:"id,omitempty"`
	ParentConnectWiseId      string                 `json:"parentConnectWiseId,omitempty"`
	ParentId                 int                    `json:"parentId,omitempty"`
	Sequence                 int                    `json:"sequence,omitempty"`
	Summary                  string                 `json:"summary,omitempty"`
	TicketId                 int                    `json:"ticketId,omitempty"`
}

type ProjectTemplateTicket struct {
	Info                     map[string]interface{} `json:"_info,omitempty"`
	BillSeparatelyFlag       bool                   `json:"billSeparatelyFlag,omitempty"`
	BudgetHours              float64                `json:"budgetHours,omitempty"`
	ConnectWiseId            string                 `json:"connectWiseId,omitempty"`
	Description              string                 `json:"description"`
	Duration                 int                    `json:"duration,omitempty"`
	ID                       int                    `json:"id,omitempty"`
	InternalAnalysis         string                 `json:"internalAnalysis,omitempty"`
	LagDays                  int                    `json:"lagDays,omitempty"`
	LagNonworkingDaysFlag    bool                   `json:"lagNonworkingDaysFlag,omitempty"`
	LineNumber               float64                `json:"lineNumber,omitempty"`
	MarkAsMilestoneFlag      bool                   `json:"markAsMilestoneFlag,omitempty"`
	Notes                    string                 `json:"notes,omitempty"`
	ParentConnectWiseId      string                 `json:"parentConnectWiseId,omitempty"`
	ParentId                 int                    `json:"parentId,omitempty"`
	PmTmpProjectRecID        int                    `json:"pmTmpProjectRecID,omitempty"`
	PredecessorClosedFlag    bool                   `json:"predecessorClosedFlag,omitempty"`
	PredecessorId            int                    `json:"predecessorId,omitempty"`
	PredecessorType          string                 `json:"predecessorType,omitempty"`
	Priority                 Priority               `json:"priority,omitempty"`
	ProjectTemplateId        int                    `json:"projectTemplateId,omitempty"`
	ProjectTemplatePhaseCwId string                 `json:"projectTemplatePhaseCwId,omitempty"`
	ProjectTemplatePhaseId   int                    `json:"projectTemplatePhaseId,omitempty"`
	RecordType               string                 `json:"recordType,omitempty"`
	Resolution               string                 `json:"resolution,omitempty"`
	Source                   ServiceSource          `json:"source,omitempty"`
	WbsCode                  string                 `json:"wbsCode,omitempty"`
	WorkRole                 WorkRole               `json:"workRole,omitempty"`
	WorkType                 WorkType               `json:"workType,omitempty"`
}

type ProjectTemplateWorkPlan struct {
	Phases     []TemplatePhase `json:"phases,omitempty"`
	TemplateId int             `json:"templateId,omitempty"`
}

type ProjectTicket struct {
	Info                       map[string]interface{} `json:"_info,omitempty"`
	ActualHours                float64                `json:"actualHours,omitempty"`
	AddressLine1               string                 `json:"addressLine1,omitempty"`
	AddressLine2               string                 `json:"addressLine2,omitempty"`
	Agreement                  Agreement              `json:"agreement,omitempty"`
	AgreementType              string                 `json:"agreementType,omitempty"`
	AllowAllClientsPortalView  bool                   `json:"allowAllClientsPortalView,omitempty"`
	Approved                   bool                   `json:"approved,omitempty"`
	AutomaticEmailCc           string                 `json:"automaticEmailCc,omitempty"`
	AutomaticEmailCcFlag       bool                   `json:"automaticEmailCcFlag,omitempty"`
	AutomaticEmailContactFlag  bool                   `json:"automaticEmailContactFlag,omitempty"`
	AutomaticEmailResourceFlag bool                   `json:"automaticEmailResourceFlag,omitempty"`
	BillExpenses               string                 `json:"billExpenses,omitempty"`
	BillProducts               string                 `json:"billProducts,omitempty"`
	BillTime                   string                 `json:"billTime,omitempty"`
	Board                      Board                  `json:"board,omitempty"`
	BudgetHours                float64                `json:"budgetHours,omitempty"`
	City                       string                 `json:"city,omitempty"`
	ClosedBy                   string                 `json:"closedBy,omitempty"`
	ClosedDate                 string                 `json:"closedDate,omitempty"`
	ClosedFlag                 bool                   `json:"closedFlag,omitempty"`
	Company                    Company                `json:"company,omitempty"`
	Contact                    Contact                `json:"contact,omitempty"`
	ContactEmailAddress        string                 `json:"contactEmailAddress,omitempty"`
	ContactEmailLookup         string                 `json:"contactEmailLookup,omitempty"`
	ContactName                string                 `json:"contactName,omitempty"`
	ContactPhoneExtension      string                 `json:"contactPhoneExtension,omitempty"`
	ContactPhoneNumber         string                 `json:"contactPhoneNumber,omitempty"`
	Country                    Country                `json:"country,omitempty"`
	Currency                   Currency               `json:"currency,omitempty"`
	CustomFields               []CustomFieldValue     `json:"customFields,omitempty"`
	CustomerUpdatedFlag        bool                   `json:"customerUpdatedFlag,omitempty"`
	Department                 SystemDepartment       `json:"department,omitempty"`
	Duration                   int                    `json:"duration,omitempty"`
	EstimatedStartDate         time.Time              `json:"estimatedStartDate,omitempty"`
	ID                         int                    `json:"id,omitempty"`
	InitialDescription         string                 `json:"initialDescription,omitempty"`
	InitialInternalAnalysis    string                 `json:"initialInternalAnalysis,omitempty"`
	InitialResolution          string                 `json:"initialResolution,omitempty"`
	IsIssueFlag                bool                   `json:"isIssueFlag,omitempty"`
	Item                       ServiceItem            `json:"item,omitempty"`
	KnowledgeBaseCategoryId    int                    `json:"knowledgeBaseCategoryId,omitempty"`
	KnowledgeBaseLinkId        int                    `json:"knowledgeBaseLinkId,omitempty"`
	KnowledgeBaseLinkType      string                 `json:"knowledgeBaseLinkType,omitempty"`
	KnowledgeBaseSubCategoryId int                    `json:"knowledgeBaseSubCategoryId,omitempty"`
	LagDays                    int                    `json:"lagDays,omitempty"`
	LagNonworkingDaysFlag      bool                   `json:"lagNonworkingDaysFlag,omitempty"`
	Location                   SystemLocation         `json:"location,omitempty"`
	MobileGuid                 string                 `json:"mobileGuid,omitempty"`
	Opportunity                Opportunity            `json:"opportunity,omitempty"`
	Owner                      Member                 `json:"owner,omitempty"`
	Phase                      ProjectPhase           `json:"phase"`
	PredecessorClosedFlag      bool                   `json:"predecessorClosedFlag,omitempty"`
	PredecessorId              int                    `json:"predecessorId,omitempty"`
	PredecessorType            string                 `json:"predecessorType,omitempty"`
	Priority                   Priority               `json:"priority,omitempty"`
	ProcessNotifications       bool                   `json:"processNotifications,omitempty"`
	Project                    Project                `json:"project,omitempty"`
	RequiredDate               time.Time              `json:"requiredDate,omitempty"`
	Resources                  string                 `json:"resources,omitempty"`
	ScheduleEndDate            time.Time              `json:"scheduleEndDate,omitempty"`
	ScheduleStartDate          time.Time              `json:"scheduleStartDate,omitempty"`
	ServiceLocation            ServiceLocation        `json:"serviceLocation,omitempty"`
	Site                       Site                   `json:"site,omitempty"`
	SiteName                   string                 `json:"siteName,omitempty"`
	SkipCallback               bool                   `json:"skipCallback,omitempty"`
	Source                     ServiceSource          `json:"source,omitempty"`
	StateIdentifier            string                 `json:"stateIdentifier,omitempty"`
	Status                     ServiceStatus          `json:"status,omitempty"`
	SubBillingAmount           float64                `json:"subBillingAmount,omitempty"`
	SubBillingMethod           string                 `json:"subBillingMethod,omitempty"`
	SubDateAccepted            string                 `json:"subDateAccepted,omitempty"`
	SubType                    ServiceSubType         `json:"subType,omitempty"`
	Summary                    string                 `json:"summary"`
	Tasks                      []TicketTask           `json:"tasks,omitempty"`
	Type                       ServiceType            `json:"type,omitempty"`
	WbsCode                    string                 `json:"wbsCode,omitempty"`
	WorkRole                   WorkRole               `json:"workRole,omitempty"`
	WorkType                   WorkType               `json:"workType,omitempty"`
	Zip                        string                 `json:"zip,omitempty"`
}

type ProjectTicketNote struct {
	Info                  map[string]interface{} `json:"_info,omitempty"`
	BundledFlag           bool                   `json:"bundledFlag,omitempty"`
	Contact               Contact                `json:"contact,omitempty"`
	DetailDescriptionFlag bool                   `json:"detailDescriptionFlag,omitempty"`
	ID                    int                    `json:"id,omitempty"`
	InternalAnalysisFlag  bool                   `json:"internalAnalysisFlag,omitempty"`
	IssueFlag             bool                   `json:"issueFlag,omitempty"`
	Member                Member                 `json:"member,omitempty"`
	MergedFlag            bool                   `json:"mergedFlag,omitempty"`
	NoteType              string                 `json:"noteType,omitempty"`
	OriginalAuthor        string                 `json:"originalAuthor,omitempty"`
	ResolutionFlag        bool                   `json:"resolutionFlag,omitempty"`
	Text                  string                 `json:"text,omitempty"`
	Ticket                Ticket                 `json:"ticket,omitempty"`
	TimeEnd               time.Time              `json:"timeEnd,omitempty"`
	TimeStart             time.Time              `json:"timeStart,omitempty"`
}

type ProjectType struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag     bool                   `json:"defaultFlag,omitempty"`
	ID              int                    `json:"id,omitempty"`
	InactiveFlag    bool                   `json:"inactiveFlag,omitempty"`
	IntegrationXref string                 `json:"integrationXref,omitempty"`
	Name            string                 `json:"name"`
}

type ProjectTypeInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type ProjectWorkplan struct {
	Phases    []ProjectWorkplanProjectPhase `json:"phases,omitempty"`
	ProjectId int                           `json:"projectId,omitempty"`
}

type StatusIndicator struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Color      string                 `json:"color,omitempty"`
	Icon       string                 `json:"icon,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier,omitempty"`
	Name       string                 `json:"name,omitempty"`
}

type TicketNote struct {
	Info                  map[string]interface{} `json:"_info,omitempty"`
	Contact               Contact                `json:"contact,omitempty"`
	CustomerUpdatedFlag   bool                   `json:"customerUpdatedFlag,omitempty"`
	DetailDescriptionFlag bool                   `json:"detailDescriptionFlag,omitempty"`
	ExternalFlag          bool                   `json:"externalFlag,omitempty"`
	ID                    int                    `json:"id,omitempty"`
	InternalAnalysisFlag  bool                   `json:"internalAnalysisFlag,omitempty"`
	InternalFlag          bool                   `json:"internalFlag,omitempty"`
	IssueFlag             bool                   `json:"issueFlag,omitempty"`
	Member                Member                 `json:"member,omitempty"`
	ProcessNotifications  bool                   `json:"processNotifications,omitempty"`
	ResolutionFlag        bool                   `json:"resolutionFlag,omitempty"`
	Text                  string                 `json:"text,omitempty"`
	TicketId              int                    `json:"ticketId,omitempty"`
}

type TicketTask struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	ChildScheduleAction string                 `json:"childScheduleAction,omitempty"`
	ChildTicketId       int                    `json:"childTicketId,omitempty"`
	ClosedFlag          bool                   `json:"closedFlag,omitempty"`
	Code                ServiceCode            `json:"code,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	Member              Member                 `json:"member,omitempty"`
	Notes               string                 `json:"notes,omitempty"`
	Priority            int                    `json:"priority,omitempty"`
	Resolution          string                 `json:"resolution,omitempty"`
	Schedule            ScheduleEntry          `json:"schedule,omitempty"`
	Summary             string                 `json:"summary,omitempty"`
	TicketId            int                    `json:"ticketId,omitempty"`
}
