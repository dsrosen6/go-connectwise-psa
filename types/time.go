package types

import (
	"time"
)

type ActivityStopwatch struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	ActivityId         int                    `json:"activityId,omitempty"`
	ActivityMobileGuid string                 `json:"activityMobileGuid,omitempty"`
	Agreement          Agreement              `json:"agreement,omitempty"`
	BillableOption     string                 `json:"billableOption,omitempty"`
	BusinessUnitId     int                    `json:"businessUnitId,omitempty"`
	DateEntered        time.Time              `json:"dateEntered,omitempty"`
	EndTime            time.Time              `json:"endTime,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	InternalNotes      string                 `json:"internalNotes,omitempty"`
	LocationId         int                    `json:"locationId,omitempty"`
	Member             Member                 `json:"member"`
	MobileGuid         string                 `json:"mobileGuid,omitempty"`
	Notes              string                 `json:"notes,omitempty"`
	StartTime          time.Time              `json:"startTime,omitempty"`
	Status             string                 `json:"status,omitempty"`
	TotalPauseTime     int                    `json:"totalPauseTime,omitempty"`
	WorkRole           WorkRole               `json:"workRole,omitempty"`
	WorkType           WorkType               `json:"workType,omitempty"`
}

type ChargeCode struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	AllowAllExpenseTypeFlag bool                   `json:"allowAllExpenseTypeFlag,omitempty"`
	BillTime                string                 `json:"billTime,omitempty"`
	Company                 Company                `json:"company"`
	Department              SystemDepartment       `json:"department,omitempty"`
	ExpenseEntryFlag        bool                   `json:"expenseEntryFlag,omitempty"`
	ExpenseTypeIds          []int                  `json:"expenseTypeIds,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	IntegrationXref         string                 `json:"integrationXref,omitempty"`
	Location                SystemLocation         `json:"location,omitempty"`
	Name                    string                 `json:"name"`
	TimeEntryFlag           bool                   `json:"timeEntryFlag,omitempty"`
	WorkRole                WorkRole               `json:"workRole,omitempty"`
	WorkType                WorkType               `json:"workType,omitempty"`
}

type ChargeCodeExpenseType struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	ChargeCode ChargeCode             `json:"chargeCode,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Type       ExpenseType            `json:"type"`
}

type ChargeCodeInfo struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	AllowAllExpenseTypeFlag bool                   `json:"allowAllExpenseTypeFlag,omitempty"`
	Department              SystemDepartment       `json:"department,omitempty"`
	ExpenseEntryFlag        bool                   `json:"expenseEntryFlag,omitempty"`
	ExpenseTypeIds          []int                  `json:"expenseTypeIds,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	Location                SystemLocation         `json:"location,omitempty"`
	Name                    string                 `json:"name,omitempty"`
	TimeEntryFlag           bool                   `json:"timeEntryFlag,omitempty"`
	WorkRole                WorkRole               `json:"workRole,omitempty"`
	WorkType                WorkType               `json:"workType,omitempty"`
}

type ScheduleStopwatch struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	Agreement          Agreement              `json:"agreement,omitempty"`
	BillableOption     string                 `json:"billableOption,omitempty"`
	BusinessUnitId     int                    `json:"businessUnitId,omitempty"`
	DateEntered        time.Time              `json:"dateEntered,omitempty"`
	EndTime            time.Time              `json:"endTime,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	InternalNotes      string                 `json:"internalNotes,omitempty"`
	LocationId         int                    `json:"locationId,omitempty"`
	Member             Member                 `json:"member"`
	MobileGuid         string                 `json:"mobileGuid,omitempty"`
	Notes              string                 `json:"notes,omitempty"`
	ScheduleId         int                    `json:"scheduleId,omitempty"`
	ScheduleMobileGuid string                 `json:"scheduleMobileGuid,omitempty"`
	StartTime          time.Time              `json:"startTime,omitempty"`
	Status             string                 `json:"status,omitempty"`
	TotalPauseTime     int                    `json:"totalPauseTime,omitempty"`
	WorkRole           WorkRole               `json:"workRole,omitempty"`
	WorkType           WorkType               `json:"workType,omitempty"`
}

type TicketStopwatch struct {
	Info                      map[string]interface{} `json:"_info,omitempty"`
	Agreement                 Agreement              `json:"agreement,omitempty"`
	BillableOption            string                 `json:"billableOption,omitempty"`
	BusinessUnitId            int                    `json:"businessUnitId,omitempty"`
	DateEntered               time.Time              `json:"dateEntered,omitempty"`
	EmailNotesToContactFlag   bool                   `json:"emailNotesToContactFlag,omitempty"`
	EmailNotesToResourcesFlag bool                   `json:"emailNotesToResourcesFlag,omitempty"`
	EndTime                   time.Time              `json:"endTime,omitempty"`
	ID                        int                    `json:"id,omitempty"`
	InternalNotes             string                 `json:"internalNotes,omitempty"`
	LocationId                int                    `json:"locationId,omitempty"`
	Member                    Member                 `json:"member"`
	MobileGuid                string                 `json:"mobileGuid,omitempty"`
	Notes                     string                 `json:"notes,omitempty"`
	ServiceStatus             ServiceStatus          `json:"serviceStatus,omitempty"`
	ShowNotesInDiscussionFlag bool                   `json:"showNotesInDiscussionFlag,omitempty"`
	ShowNotesInInternalFlag   bool                   `json:"showNotesInInternalFlag,omitempty"`
	ShowNotesInResolutionFlag bool                   `json:"showNotesInResolutionFlag,omitempty"`
	StartTime                 time.Time              `json:"startTime,omitempty"`
	Status                    string                 `json:"status,omitempty"`
	Ticket                    Ticket                 `json:"ticket"`
	TicketMobileGuid          string                 `json:"ticketMobileGuid,omitempty"`
	TotalPauseTime            int                    `json:"totalPauseTime,omitempty"`
	WorkRole                  WorkRole               `json:"workRole,omitempty"`
	WorkType                  WorkType               `json:"workType,omitempty"`
}

type TimeAccrual struct {
	Info                         map[string]interface{} `json:"_info,omitempty"`
	HolidayAvailableType         string                 `json:"holidayAvailableType,omitempty"`
	HolidayCarryoverAllowedFlag  bool                   `json:"holidayCarryoverAllowedFlag,omitempty"`
	HolidayCarryoverLimit        float64                `json:"holidayCarryoverLimit,omitempty"`
	HolidayFlag                  bool                   `json:"holidayFlag,omitempty"`
	ID                           int                    `json:"id,omitempty"`
	Location                     SystemLocation         `json:"location,omitempty"`
	PtoAvailableType             string                 `json:"ptoAvailableType,omitempty"`
	PtoCarryoverAllowedFlag      bool                   `json:"ptoCarryoverAllowedFlag,omitempty"`
	PtoCarryoverLimit            float64                `json:"ptoCarryoverLimit,omitempty"`
	PtoFlag                      bool                   `json:"ptoFlag,omitempty"`
	SickAvailableType            string                 `json:"sickAvailableType,omitempty"`
	SickCarryoverAllowedFlag     bool                   `json:"sickCarryoverAllowedFlag,omitempty"`
	SickCarryoverLimit           float64                `json:"sickCarryoverLimit,omitempty"`
	SickFlag                     bool                   `json:"sickFlag,omitempty"`
	VacationAvailableType        string                 `json:"vacationAvailableType,omitempty"`
	VacationCarryoverAllowedFlag bool                   `json:"vacationCarryoverAllowedFlag,omitempty"`
	VacationCarryoverLimit       float64                `json:"vacationCarryoverLimit,omitempty"`
	VacationFlag                 bool                   `json:"vacationFlag,omitempty"`
}

type TimeAccrualDetail struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	AccrualType string                 `json:"accrualType,omitempty"`
	EndYear     int                    `json:"endYear,omitempty"`
	Hours       float64                `json:"hours,omitempty"`
	ID          int                    `json:"id,omitempty"`
	StartYear   int                    `json:"startYear,omitempty"`
	TimeAccrual TimeAccrual            `json:"timeAccrual,omitempty"`
}

type TimeEntry struct {
	Info                       map[string]interface{} `json:"_info,omitempty"`
	Activity                   Activity               `json:"activity,omitempty"`
	ActualHours                float64                `json:"actualHours,omitempty"`
	AddToDetailDescriptionFlag bool                   `json:"addToDetailDescriptionFlag,omitempty"`
	AddToInternalAnalysisFlag  bool                   `json:"addToInternalAnalysisFlag,omitempty"`
	AddToResolutionFlag        bool                   `json:"addToResolutionFlag,omitempty"`
	Adjustment                 float64                `json:"adjustment,omitempty"`
	Agreement                  Agreement              `json:"agreement,omitempty"`
	AgreementAdjustment        float64                `json:"agreementAdjustment,omitempty"`
	AgreementAmount            float64                `json:"agreementAmount,omitempty"`
	AgreementHours             float64                `json:"agreementHours,omitempty"`
	AgreementType              string                 `json:"agreementType,omitempty"`
	BillableOption             string                 `json:"billableOption,omitempty"`
	BusinessGroupDesc          string                 `json:"businessGroupDesc,omitempty"`
	BusinessUnitId             int                    `json:"businessUnitId,omitempty"`
	ChargeToId                 int                    `json:"chargeToId,omitempty"`
	ChargeToType               string                 `json:"chargeToType,omitempty"`
	Company                    Company                `json:"company,omitempty"`
	CompanyType                string                 `json:"companyType,omitempty"`
	CustomFields               []CustomFieldValue     `json:"customFields,omitempty"`
	DateEntered                time.Time              `json:"dateEntered,omitempty"`
	Department                 BillingUnit            `json:"department,omitempty"`
	EmailCc                    string                 `json:"emailCc,omitempty"`
	EmailCcFlag                bool                   `json:"emailCcFlag,omitempty"`
	EmailContactFlag           bool                   `json:"emailContactFlag,omitempty"`
	EmailResourceFlag          bool                   `json:"emailResourceFlag,omitempty"`
	EnteredBy                  string                 `json:"enteredBy,omitempty"`
	ExtendedInvoiceAmount      float64                `json:"extendedInvoiceAmount,omitempty"`
	HourlyCost                 string                 `json:"hourlyCost,omitempty"`
	HourlyRate                 float64                `json:"hourlyRate,omitempty"`
	HoursBilled                float64                `json:"hoursBilled,omitempty"`
	HoursDeduct                float64                `json:"hoursDeduct,omitempty"`
	ID                         int                    `json:"id,omitempty"`
	InternalNotes              string                 `json:"internalNotes,omitempty"`
	Invoice                    Invoice                `json:"invoice,omitempty"`
	InvoiceFlag                bool                   `json:"invoiceFlag,omitempty"`
	InvoiceHours               float64                `json:"invoiceHours,omitempty"`
	InvoiceReady               int                    `json:"invoiceReady,omitempty"`
	Location                   OwnerLevel             `json:"location,omitempty"`
	LocationId                 int                    `json:"locationId,omitempty"`
	LocationName               string                 `json:"locationName,omitempty"`
	Member                     Member                 `json:"member,omitempty"`
	MobileGuid                 string                 `json:"mobileGuid,omitempty"`
	Notes                      string                 `json:"notes,omitempty"`
	OpportunityRecid           int                    `json:"opportunityRecid,omitempty"`
	OverageRate                float64                `json:"overageRate,omitempty"`
	Phase                      ProjectPhase           `json:"phase,omitempty"`
	Project                    Project                `json:"project,omitempty"`
	ProjectActivity            string                 `json:"projectActivity,omitempty"`
	Status                     string                 `json:"status,omitempty"`
	TaxCode                    TaxCode                `json:"taxCode,omitempty"`
	Territory                  string                 `json:"territory,omitempty"`
	Ticket                     Ticket                 `json:"ticket,omitempty"`
	TicketBoard                string                 `json:"ticketBoard,omitempty"`
	TicketStatus               string                 `json:"ticketStatus,omitempty"`
	TicketSubType              string                 `json:"ticketSubType,omitempty"`
	TicketType                 string                 `json:"ticketType,omitempty"`
	TimeEnd                    time.Time              `json:"timeEnd,omitempty"`
	TimeSheet                  TimeSheet              `json:"timeSheet,omitempty"`
	TimeStart                  time.Time              `json:"timeStart"`
	WorkRole                   WorkRole               `json:"workRole,omitempty"`
	WorkType                   WorkType               `json:"workType,omitempty"`
}

type TimeEntryAudit struct {
	Info     map[string]interface{} `json:"_info,omitempty"`
	ID       int                    `json:"id,omitempty"`
	Member   Member                 `json:"member,omitempty"`
	Message  string                 `json:"message,omitempty"`
	NewValue string                 `json:"newValue,omitempty"`
	OldValue string                 `json:"oldValue,omitempty"`
	Source   string                 `json:"source,omitempty"`
	Type     string                 `json:"type,omitempty"`
	Value    string                 `json:"value,omitempty"`
}

type TimeEntryChangeLog struct {
	Info                     map[string]interface{} `json:"_info,omitempty"`
	Action                   string                 `json:"action,omitempty"`
	Activity                 Activity               `json:"activity,omitempty"`
	ActivitySubject          string                 `json:"activitySubject,omitempty"`
	ActualUtilizedHrs        float64                `json:"actualUtilizedHrs,omitempty"`
	Agreement                Agreement              `json:"agreement,omitempty"`
	AgreementAdjustmentFirm  float64                `json:"agreementAdjustmentFirm,omitempty"`
	AgreementAdjustmentTotal float64                `json:"agreementAdjustmentTotal,omitempty"`
	AgreementAmountCovered   float64                `json:"agreementAmountCovered,omitempty"`
	AgreementHoursCovered    float64                `json:"agreementHoursCovered,omitempty"`
	AgreementType            string                 `json:"agreementType,omitempty"`
	BillableAmount           float64                `json:"billableAmount,omitempty"`
	BillableFlag             bool                   `json:"billableFlag,omitempty"`
	BillableHours            float64                `json:"billableHours,omitempty"`
	BillableOption           string                 `json:"billableOption,omitempty"`
	BillableUtilizedHours    float64                `json:"billableUtilizedHours,omitempty"`
	BillingStatus            string                 `json:"billingStatus,omitempty"`
	BusinessGroup            string                 `json:"businessGroup,omitempty"`
	ChargeCode               string                 `json:"chargeCode,omitempty"`
	ChargeTo                 string                 `json:"chargeTo,omitempty"`
	ChargeToRecId            int                    `json:"chargeToRecId,omitempty"`
	ChargeToType             string                 `json:"chargeToType,omitempty"`
	Company                  Company                `json:"company,omitempty"`
	CompanyAndAgreement      string                 `json:"companyAndAgreement,omitempty"`
	CompanyName              string                 `json:"companyName,omitempty"`
	CompanyType              string                 `json:"companyType,omitempty"`
	DateInvoice              string                 `json:"dateInvoice,omitempty"`
	DateStart                string                 `json:"dateStart,omitempty"`
	FirstName                string                 `json:"firstName,omitempty"`
	HourlyCost               string                 `json:"hourlyCost,omitempty"`
	HourlyCostDecimal        float64                `json:"hourlyCostDecimal,omitempty"`
	HourlyRate               float64                `json:"hourlyRate,omitempty"`
	HoursActual              float64                `json:"hoursActual,omitempty"`
	ID                       int                    `json:"id,omitempty"`
	InternalNote             string                 `json:"internalNote,omitempty"`
	Invoice                  Invoice                `json:"invoice,omitempty"`
	InvoiceAdjustmentFirm    float64                `json:"invoiceAdjustmentFirm,omitempty"`
	InvoiceAdjustmentTotal   float64                `json:"invoiceAdjustmentTotal,omitempty"`
	InvoiceFlag              bool                   `json:"invoiceFlag,omitempty"`
	InvoiceNumber            string                 `json:"invoiceNumber,omitempty"`
	InvoiceReady             bool                   `json:"invoiceReady,omitempty"`
	Invoicedhours            float64                `json:"invoicedhours,omitempty"`
	LastName                 string                 `json:"lastName,omitempty"`
	LocationName             string                 `json:"locationName,omitempty"`
	Member                   Member                 `json:"member,omitempty"`
	MemberDailyCapacity      float64                `json:"memberDailyCapacity,omitempty"`
	MemberId                 string                 `json:"memberId,omitempty"`
	NonBillableAmt           float64                `json:"nonBillableAmt,omitempty"`
	NonBillableHrs           float64                `json:"nonBillableHrs,omitempty"`
	Notes                    string                 `json:"notes,omitempty"`
	OpportunityRecId         int                    `json:"opportunityRecId,omitempty"`
	OptionId                 string                 `json:"optionId,omitempty"`
	PartnerId                string                 `json:"partnerId,omitempty"`
	Phase                    ProjectPhase           `json:"phase,omitempty"`
	ProcessingStatus         string                 `json:"processingStatus,omitempty"`
	ProductInstanceId        string                 `json:"productInstanceId,omitempty"`
	Project                  Project                `json:"project,omitempty"`
	ProjectActivity          string                 `json:"projectActivity,omitempty"`
	ProjectName              string                 `json:"projectName,omitempty"`
	ProjectPhase             string                 `json:"projectPhase,omitempty"`
	ServiceRequestStatus     string                 `json:"serviceRequestStatus,omitempty"`
	ServiceRequestSummary    string                 `json:"serviceRequestSummary,omitempty"`
	Territory                string                 `json:"territory,omitempty"`
	Ticket                   Ticket                 `json:"ticket,omitempty"`
	TicketCurrentBoard       string                 `json:"ticketCurrentBoard,omitempty"`
	TicketSubtype            string                 `json:"ticketSubtype,omitempty"`
	TicketType               string                 `json:"ticketType,omitempty"`
	TimeEnd                  string                 `json:"timeEnd,omitempty"`
	TimeEndUtc               string                 `json:"timeEndUtc,omitempty"`
	TimeRecId                int                    `json:"timeRecId,omitempty"`
	TimeStart                string                 `json:"timeStart,omitempty"`
	TimeStartUtc             string                 `json:"timeStartUtc,omitempty"`
	TimeStatus               string                 `json:"timeStatus,omitempty"`
	UtilizationFlag          bool                   `json:"utilizationFlag,omitempty"`
	WorkRole                 WorkRole               `json:"workRole,omitempty"`
	WorkType                 WorkType               `json:"workType,omitempty"`
}

type TimePeriod struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	DeadlineDate    string                 `json:"deadlineDate,omitempty"`
	EndDate         string                 `json:"endDate,omitempty"`
	ID              int                    `json:"id,omitempty"`
	Period          int                    `json:"period,omitempty"`
	StartDate       string                 `json:"startDate,omitempty"`
	TimePeriodSetup TimePeriodSetup        `json:"timePeriodSetup,omitempty"`
}

type TimePeriodSetup struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	DaysPastEndDate         int                    `json:"daysPastEndDate,omitempty"`
	Description             string                 `json:"description,omitempty"`
	FirstPeriodEndDate      string                 `json:"firstPeriodEndDate"`
	ID                      int                    `json:"id,omitempty"`
	LastDayFlag             bool                   `json:"lastDayFlag,omitempty"`
	MonthlyPeriodEnds       int                    `json:"monthlyPeriodEnds,omitempty"`
	NumberFuturePeriods     int                    `json:"numberFuturePeriods,omitempty"`
	PeriodApplyTo           string                 `json:"periodApplyTo,omitempty"`
	SemiMonthlyFirstPeriod  int                    `json:"semiMonthlyFirstPeriod,omitempty"`
	SemiMonthlyLastDayFlag  bool                   `json:"semiMonthlyLastDayFlag,omitempty"`
	SemiMonthlySecondPeriod int                    `json:"semiMonthlySecondPeriod,omitempty"`
	Type                    string                 `json:"type,omitempty"`
	Year                    int                    `json:"year,omitempty"`
}

type TimePeriodSetupDefaults struct {
}

type TimeSheet struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	DateEnd   string                 `json:"dateEnd,omitempty"`
	DateStart string                 `json:"dateStart,omitempty"`
	Deadline  string                 `json:"deadline,omitempty"`
	Hours     float64                `json:"hours,omitempty"`
	ID        int                    `json:"id,omitempty"`
	Member    Member                 `json:"member,omitempty"`
	Period    int                    `json:"period,omitempty"`
	Status    string                 `json:"status,omitempty"`
	Year      int                    `json:"year,omitempty"`
}

type TimeSheetAudit struct {
	Info     map[string]interface{} `json:"_info,omitempty"`
	ID       int                    `json:"id,omitempty"`
	Member   Member                 `json:"member,omitempty"`
	Message  string                 `json:"message,omitempty"`
	NewValue string                 `json:"newValue,omitempty"`
	OldValue string                 `json:"oldValue,omitempty"`
	Source   string                 `json:"source,omitempty"`
	Type     string                 `json:"type,omitempty"`
	Value    string                 `json:"value,omitempty"`
}

type TimeSheetTierUpdate struct {
	ApprovalType string `json:"approvalType,omitempty"`
	ID           int    `json:"id,omitempty"`
}

type WorkRole struct {
	Info                      map[string]interface{} `json:"_info,omitempty"`
	AddAllAgreementExclusions bool                   `json:"addAllAgreementExclusions,omitempty"`
	AddAllLocations           bool                   `json:"addAllLocations,omitempty"`
	HourlyRate                float64                `json:"hourlyRate,omitempty"`
	ID                        int                    `json:"id,omitempty"`
	InactiveFlag              bool                   `json:"inactiveFlag,omitempty"`
	IntegrationXref           string                 `json:"integrationXref,omitempty"`
	LocationIds               []int                  `json:"locationIds,omitempty"`
	Name                      string                 `json:"name"`
	RemoveAllLocations        bool                   `json:"removeAllLocations,omitempty"`
}

type WorkRoleInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type WorkRoleLocation struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	HourlyRate float64                `json:"hourlyRate,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Location   SystemLocation         `json:"location"`
	WorkRole   WorkRole               `json:"workRole,omitempty"`
}

type WorkType struct {
	Info                      map[string]interface{} `json:"_info,omitempty"`
	AccrualType               string                 `json:"accrualType,omitempty"`
	ActivityDefaultFlag       bool                   `json:"activityDefaultFlag,omitempty"`
	AddAllAgreementExclusions bool                   `json:"addAllAgreementExclusions,omitempty"`
	BillTime                  string                 `json:"billTime,omitempty"`
	CostMultiplier            float64                `json:"costMultiplier,omitempty"`
	HoursMax                  float64                `json:"hoursMax,omitempty"`
	HoursMin                  float64                `json:"hoursMin,omitempty"`
	ID                        int                    `json:"id,omitempty"`
	InactiveFlag              bool                   `json:"inactiveFlag,omitempty"`
	IntegrationXRef           string                 `json:"integrationXRef,omitempty"`
	Name                      string                 `json:"name"`
	OverallDefaultFlag        bool                   `json:"overallDefaultFlag,omitempty"`
	Rate                      float64                `json:"rate,omitempty"`
	RateType                  string                 `json:"rateType,omitempty"`
	RoundBillHoursTo          float64                `json:"roundBillHoursTo,omitempty"`
	UtilizationFlag           bool                   `json:"utilizationFlag,omitempty"`
}

type WorkTypeInfo struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	ActivityDefaultFlag bool                   `json:"activityDefaultFlag,omitempty"`
	DefaultFlag         bool                   `json:"defaultFlag,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	InactiveFlag        bool                   `json:"inactiveFlag,omitempty"`
	Name                string                 `json:"name,omitempty"`
}
