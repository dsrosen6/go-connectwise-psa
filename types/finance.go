package types

import (
	"time"
)

type AccountingBatch struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	BatchIdentifier    string                 `json:"batchIdentifier,omitempty"`
	ClosedFlag         bool                   `json:"closedFlag,omitempty"`
	ExportExpensesFlag bool                   `json:"exportExpensesFlag,omitempty"`
	ExportInvoicesFlag bool                   `json:"exportInvoicesFlag,omitempty"`
	ExportProductsFlag bool                   `json:"exportProductsFlag,omitempty"`
	ID                 int                    `json:"id,omitempty"`
}

type AccountingPackage struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier,omitempty"`
	Name       string                 `json:"name,omitempty"`
}

type AccountingPackageSetup struct {
	Info                       map[string]interface{} `json:"_info,omitempty"`
	AccountingPackage          AccountingPackage      `json:"accountingPackage"`
	DirectTransferFlag         bool                   `json:"directTransferFlag,omitempty"`
	EnableTaxGroupsFlag        bool                   `json:"enableTaxGroupsFlag,omitempty"`
	ExpenseFormat              string                 `json:"expenseFormat,omitempty"`
	ID                         int                    `json:"id,omitempty"`
	IncludeCogsDropShipFlag    bool                   `json:"includeCogsDropShipFlag,omitempty"`
	IncludeCogsEntriesFlag     bool                   `json:"includeCogsEntriesFlag,omitempty"`
	IncludeExpensesFlag        bool                   `json:"includeExpensesFlag,omitempty"`
	IncludeInvoicesFlag        bool                   `json:"includeInvoicesFlag,omitempty"`
	IncludeItemsFlag           bool                   `json:"includeItemsFlag,omitempty"`
	IncludeSalesTaxFlag        bool                   `json:"includeSalesTaxFlag,omitempty"`
	InventorySOHFlag           bool                   `json:"inventorySOHFlag,omitempty"`
	InvoiceFormat              string                 `json:"invoiceFormat,omitempty"`
	SendComponentAmountFlag    bool                   `json:"sendComponentAmountFlag,omitempty"`
	SendUomFlag                bool                   `json:"sendUomFlag,omitempty"`
	SuppressMemoFlag           bool                   `json:"suppressMemoFlag,omitempty"`
	SyncPaymentInfoFlag        bool                   `json:"syncPaymentInfoFlag,omitempty"`
	SyncWisePayPaymentInfoFlag bool                   `json:"syncWisePayPaymentInfoFlag,omitempty"`
	TransferExpensesAsBillFlag bool                   `json:"transferExpensesAsBillFlag,omitempty"`
	ZeroDollarTaxAmountsFlag   bool                   `json:"zeroDollarTaxAmountsFlag,omitempty"`
}

type Addition struct {
	Info                     map[string]interface{} `json:"_info,omitempty"`
	AgreementId              int                    `json:"agreementId,omitempty"`
	AgreementStatus          string                 `json:"agreementStatus,omitempty"`
	BillCustomer             string                 `json:"billCustomer,omitempty"`
	BilledQuantity           float64                `json:"billedQuantity,omitempty"`
	CancelledDate            time.Time              `json:"cancelledDate,omitempty"`
	CustomFields             []CustomFieldValue     `json:"customFields,omitempty"`
	Description              string                 `json:"description,omitempty"`
	EffectiveDate            time.Time              `json:"effectiveDate,omitempty"`
	ExtCost                  float64                `json:"extCost,omitempty"`
	ExtPrice                 float64                `json:"extPrice,omitempty"`
	ExtendedProrateCost      float64                `json:"extendedProrateCost,omitempty"`
	ExtendedProratePrice     float64                `json:"extendedProratePrice,omitempty"`
	ID                       int                    `json:"id,omitempty"`
	InvoiceDescription       string                 `json:"invoiceDescription,omitempty"`
	InvoiceGrouping          InvoiceGrouping        `json:"invoiceGrouping,omitempty"`
	LessIncluded             float64                `json:"lessIncluded,omitempty"`
	Margin                   float64                `json:"margin,omitempty"`
	Opportunity              Opportunity            `json:"opportunity,omitempty"`
	Product                  IvItem                 `json:"product"`
	ProrateCost              float64                `json:"prorateCost,omitempty"`
	ProrateCurrentPeriodFlag bool                   `json:"prorateCurrentPeriodFlag,omitempty"`
	ProratePrice             float64                `json:"proratePrice,omitempty"`
	PurchaseItemFlag         bool                   `json:"purchaseItemFlag,omitempty"`
	Quantity                 float64                `json:"quantity,omitempty"`
	SequenceNumber           float64                `json:"sequenceNumber,omitempty"`
	SerialNumber             string                 `json:"serialNumber,omitempty"`
	SpecialOrderFlag         bool                   `json:"specialOrderFlag,omitempty"`
	TaxableFlag              bool                   `json:"taxableFlag,omitempty"`
	UnitCost                 float64                `json:"unitCost,omitempty"`
	UnitPrice                float64                `json:"unitPrice,omitempty"`
	Uom                      string                 `json:"uom,omitempty"`
}

type Agreement struct {
	Info                         map[string]interface{} `json:"_info,omitempty"`
	AgreementStatus              string                 `json:"agreementStatus,omitempty"`
	AllowOverruns                bool                   `json:"allowOverruns,omitempty"`
	ApplicationCycle             string                 `json:"applicationCycle,omitempty"`
	ApplicationLimit             float64                `json:"applicationLimit,omitempty"`
	ApplicationUnits             string                 `json:"applicationUnits,omitempty"`
	ApplicationUnlimitedFlag     bool                   `json:"applicationUnlimitedFlag,omitempty"`
	AutoInvoiceFlag              bool                   `json:"autoInvoiceFlag,omitempty"`
	BillAmount                   float64                `json:"billAmount,omitempty"`
	BillExpenses                 string                 `json:"billExpenses,omitempty"`
	BillOneTimeFlag              bool                   `json:"billOneTimeFlag,omitempty"`
	BillProducts                 string                 `json:"billProducts,omitempty"`
	BillStartDate                time.Time              `json:"billStartDate,omitempty"`
	BillTime                     string                 `json:"billTime,omitempty"`
	BillToCompany                Company                `json:"billToCompany,omitempty"`
	BillToContact                Contact                `json:"billToContact,omitempty"`
	BillToSite                   Site                   `json:"billToSite,omitempty"`
	BillableExpenseInvoice       bool                   `json:"billableExpenseInvoice,omitempty"`
	BillableProductInvoice       bool                   `json:"billableProductInvoice,omitempty"`
	BillableTimeInvoice          bool                   `json:"billableTimeInvoice,omitempty"`
	BillingCycle                 BillingCycle           `json:"billingCycle,omitempty"`
	BillingTerms                 BillingTerms           `json:"billingTerms,omitempty"`
	BottomComment                bool                   `json:"bottomComment,omitempty"`
	CancelledFlag                bool                   `json:"cancelledFlag,omitempty"`
	CarryOverUnused              bool                   `json:"carryOverUnused,omitempty"`
	ChargeToFirm                 bool                   `json:"chargeToFirm,omitempty"`
	CompHourlyRate               float64                `json:"compHourlyRate,omitempty"`
	CompLimitAmount              float64                `json:"compLimitAmount,omitempty"`
	Company                      Company                `json:"company"`
	CompanyLocation              SystemLocation         `json:"companyLocation,omitempty"`
	Contact                      Contact                `json:"contact"`
	CoverAgreementExpense        bool                   `json:"coverAgreementExpense,omitempty"`
	CoverAgreementProduct        bool                   `json:"coverAgreementProduct,omitempty"`
	CoverAgreementTime           bool                   `json:"coverAgreementTime,omitempty"`
	CoverSalesTax                bool                   `json:"coverSalesTax,omitempty"`
	Currency                     Currency               `json:"currency,omitempty"`
	CustomFields                 []CustomFieldValue     `json:"customFields,omitempty"`
	CustomerPO                   string                 `json:"customerPO,omitempty"`
	DateCancelled                time.Time              `json:"dateCancelled,omitempty"`
	Department                   SystemDepartment       `json:"department,omitempty"`
	EmployeeCompNotExceed        string                 `json:"employeeCompNotExceed,omitempty"`
	EmployeeCompRate             string                 `json:"employeeCompRate,omitempty"`
	EndDate                      time.Time              `json:"endDate,omitempty"`
	ExpireWhenZero               bool                   `json:"expireWhenZero,omitempty"`
	ExpiredDays                  int                    `json:"expiredDays,omitempty"`
	ID                           int                    `json:"id,omitempty"`
	InternalNotes                string                 `json:"internalNotes,omitempty"`
	InvoiceDescription           string                 `json:"invoiceDescription,omitempty"`
	InvoiceProratedAdditionsFlag bool                   `json:"invoiceProratedAdditionsFlag,omitempty"`
	InvoiceTemplate              InvoiceTemplate        `json:"invoiceTemplate,omitempty"`
	InvoicingCycle               string                 `json:"invoicingCycle,omitempty"`
	Limit                        int                    `json:"limit,omitempty"`
	Location                     SystemLocation         `json:"location,omitempty"`
	Name                         string                 `json:"name"`
	NextInvoiceDate              string                 `json:"nextInvoiceDate,omitempty"`
	NoEndingDateFlag             bool                   `json:"noEndingDateFlag,omitempty"`
	OneTimeFlag                  bool                   `json:"oneTimeFlag,omitempty"`
	Opportunity                  Opportunity            `json:"opportunity,omitempty"`
	ParentAgreement              *Agreement             `json:"parentAgreement,omitempty"`
	PeriodType                   string                 `json:"periodType,omitempty"`
	ProjectType                  ProjectType            `json:"projectType,omitempty"`
	ProrateFirstBill             float64                `json:"prorateFirstBill,omitempty"`
	ProrateFlag                  bool                   `json:"prorateFlag,omitempty"`
	ReasonCancelled              string                 `json:"reasonCancelled,omitempty"`
	RestrictDepartmentFlag       bool                   `json:"restrictDepartmentFlag,omitempty"`
	RestrictDownPayment          bool                   `json:"restrictDownPayment,omitempty"`
	RestrictLocationFlag         bool                   `json:"restrictLocationFlag,omitempty"`
	ShipToCompany                Company                `json:"shipToCompany,omitempty"`
	ShipToContact                Contact                `json:"shipToContact,omitempty"`
	ShipToSite                   Site                   `json:"shipToSite,omitempty"`
	Site                         Site                   `json:"site,omitempty"`
	Sla                          SLA                    `json:"sla,omitempty"`
	StartDate                    time.Time              `json:"startDate,omitempty"`
	SubContractCompany           Company                `json:"subContractCompany,omitempty"`
	SubContractContact           Contact                `json:"subContractContact,omitempty"`
	TaxCode                      TaxCode                `json:"taxCode,omitempty"`
	Taxable                      bool                   `json:"taxable,omitempty"`
	TopComment                   bool                   `json:"topComment,omitempty"`
	Type                         AgreementType          `json:"type"`
	WorkOrder                    string                 `json:"workOrder,omitempty"`
	WorkRole                     WorkRole               `json:"workRole,omitempty"`
	WorkType                     WorkType               `json:"workType,omitempty"`
}

type AgreementAdjustment struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	AgreementId   int                    `json:"agreementId,omitempty"`
	Amount        float64                `json:"amount,omitempty"`
	CustomFields  []CustomFieldValue     `json:"customFields,omitempty"`
	Description   string                 `json:"description,omitempty"`
	EffectiveDate time.Time              `json:"effectiveDate,omitempty"`
	ID            int                    `json:"id,omitempty"`
}

type AgreementApplicationParameters struct {
	AgrBillingCycle        AgreementApplicationBillingCycle `json:"agrBillingCycle,omitempty"`
	AgreementExpiresFlag   bool                             `json:"agreementExpiresFlag,omitempty"`
	AllowOverrunsFlag      bool                             `json:"allowOverrunsFlag,omitempty"`
	ApplicationLimit       AgreementApplicationLimit        `json:"applicationLimit,omitempty"`
	ApplicationLimitAmount float64                          `json:"applicationLimitAmount,omitempty"`
	ApplicationUnit        AgreementApplicationUnit         `json:"applicationUnit,omitempty"`
	AvailablePer           AgreementApplicationAviablePer   `json:"availablePer,omitempty"`
	CarryOverDays          int                              `json:"carryOverDays,omitempty"`
	CarryoverUnusedFlag    bool                             `json:"carryoverUnusedFlag,omitempty"`
	ChargeAdjustmentsFlag  bool                             `json:"chargeAdjustmentsFlag,omitempty"`
	CoversExpensesFlag     bool                             `json:"coversExpensesFlag,omitempty"`
	CoversProductsFlag     bool                             `json:"coversProductsFlag,omitempty"`
	CoversTaxFlag          bool                             `json:"coversTaxFlag,omitempty"`
	CoversTimeFlag         bool                             `json:"coversTimeFlag,omitempty"`
	OverrunLimit           int                              `json:"overrunLimit,omitempty"`
	PrepayFlag             bool                             `json:"prepayFlag,omitempty"`
	UserDefinedFieldValues []UserDefinedFieldValueModel     `json:"userDefinedFieldValues,omitempty"`
}

type AgreementBatchSetup struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	DaysInAdvance int                    `json:"daysInAdvance,omitempty"`
	ID            int                    `json:"id,omitempty"`
	NextRunDate   time.Time              `json:"nextRunDate"`
}

type AgreementRecap struct {
	AdjustmentAmount      float64 `json:"adjustmentAmount,omitempty"`
	AgreementStatus       string  `json:"agreementStatus,omitempty"`
	AvailableAmount       float64 `json:"availableAmount,omitempty"`
	CompanyName           string  `json:"companyName,omitempty"`
	ID                    int     `json:"id,omitempty"`
	IsUnlimited           string  `json:"isUnlimited,omitempty"`
	LastInvoiceAmount     string  `json:"lastInvoiceAmount,omitempty"`
	LastInvoiceDate       string  `json:"lastInvoiceDate,omitempty"`
	LastInvoiceNumber     string  `json:"lastInvoiceNumber,omitempty"`
	Name                  string  `json:"name,omitempty"`
	NextInvoiceAmount     float64 `json:"nextInvoiceAmount,omitempty"`
	NextInvoiceDate       string  `json:"nextInvoiceDate,omitempty"`
	OverrunAmount         float64 `json:"overrunAmount,omitempty"`
	RemainingAmount       float64 `json:"remainingAmount,omitempty"`
	StartingAmount        float64 `json:"startingAmount,omitempty"`
	UnbilledOverageAmount float64 `json:"unbilledOverageAmount,omitempty"`
	UnbilledPeriods       int     `json:"unbilledPeriods,omitempty"`
	UsedAmount            float64 `json:"usedAmount,omitempty"`
}

type AgreementRecurringParameters struct {
	AGRAmount                    float64                      `json:"aGRAmount,omitempty"`
	AGRProrate                   float64                      `json:"aGRProrate,omitempty"`
	AdditionsAmount              float64                      `json:"additionsAmount,omitempty"`
	AutoInvoiceFlag              bool                         `json:"autoInvoiceFlag,omitempty"`
	BillStartDate                string                       `json:"billStartDate,omitempty"`
	BillingCycle                 GenericNameIdDTO             `json:"billingCycle,omitempty"`
	ChildrenAmount               float64                      `json:"childrenAmount,omitempty"`
	Currency                     GenericNameIdDTO             `json:"currency,omitempty"`
	CycleBase                    GenericNameIdDTO             `json:"cycleBase,omitempty"`
	InvoiceProratedAdditionsFlag bool                         `json:"invoiceProratedAdditionsFlag,omitempty"`
	ProrateFlag                  bool                         `json:"prorateFlag,omitempty"`
	RestrictDownpayment          bool                         `json:"restrictDownpayment,omitempty"`
	TaxCode                      GenericNameIdDTO             `json:"taxCode,omitempty"`
	Taxable                      bool                         `json:"taxable,omitempty"`
	Terms                        GenericNameIdDTO             `json:"terms,omitempty"`
	TotalAmount                  float64                      `json:"totalAmount,omitempty"`
	UserDefinedFieldValues       []UserDefinedFieldValueModel `json:"userDefinedFieldValues,omitempty"`
}

type AgreementSite struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	AgreementId  int                    `json:"agreementId,omitempty"`
	Company      Company                `json:"company"`
	CustomFields []CustomFieldValue     `json:"customFields,omitempty"`
	ID           int                    `json:"id,omitempty"`
	Site         Site                   `json:"site,omitempty"`
}

type AgreementTabsCount struct {
}

type AgreementType struct {
	Info                         map[string]interface{} `json:"_info,omitempty"`
	AddAllWorkRoleExclusions     bool                   `json:"addAllWorkRoleExclusions,omitempty"`
	AddAllWorkTypeExclusions     bool                   `json:"addAllWorkTypeExclusions,omitempty"`
	AllowOverrunsFlag            bool                   `json:"allowOverrunsFlag,omitempty"`
	ApplicationCycle             string                 `json:"applicationCycle,omitempty"`
	ApplicationLimit             float64                `json:"applicationLimit,omitempty"`
	ApplicationUnits             string                 `json:"applicationUnits,omitempty"`
	ApplicationUnlimitedFlag     bool                   `json:"applicationUnlimitedFlag,omitempty"`
	AutoInvoiceFlag              bool                   `json:"autoInvoiceFlag,omitempty"`
	BillAmount                   float64                `json:"billAmount,omitempty"`
	BillExpenses                 string                 `json:"billExpenses,omitempty"`
	BillOneTimeFlag              bool                   `json:"billOneTimeFlag,omitempty"`
	BillProducts                 string                 `json:"billProducts,omitempty"`
	BillTime                     string                 `json:"billTime,omitempty"`
	BillableExpenseInvoiceFlag   bool                   `json:"billableExpenseInvoiceFlag,omitempty"`
	BillableProductInvoiceFlag   bool                   `json:"billableProductInvoiceFlag,omitempty"`
	BillableTimeInvoiceFlag      bool                   `json:"billableTimeInvoiceFlag,omitempty"`
	BillingCycle                 BillingCycle           `json:"billingCycle,omitempty"`
	BillingTerms                 BillingTerms           `json:"billingTerms,omitempty"`
	BottomCommentFlag            bool                   `json:"bottomCommentFlag,omitempty"`
	CarryOverUnusedFlag          bool                   `json:"carryOverUnusedFlag,omitempty"`
	ChargeToFirmFlag             bool                   `json:"chargeToFirmFlag,omitempty"`
	CompHourlyRate               float64                `json:"compHourlyRate,omitempty"`
	CompLimitAmount              float64                `json:"compLimitAmount,omitempty"`
	CopyWorkRolesFlag            bool                   `json:"copyWorkRolesFlag,omitempty"`
	CopyWorkTypesFlag            bool                   `json:"copyWorkTypesFlag,omitempty"`
	CoverAgreementExpenseFlag    bool                   `json:"coverAgreementExpenseFlag,omitempty"`
	CoverAgreementProductFlag    bool                   `json:"coverAgreementProductFlag,omitempty"`
	CoverAgreementTimeFlag       bool                   `json:"coverAgreementTimeFlag,omitempty"`
	CoverSalesTaxFlag            bool                   `json:"coverSalesTaxFlag,omitempty"`
	DefaultFlag                  bool                   `json:"defaultFlag,omitempty"`
	Department                   SystemDepartment       `json:"department,omitempty"`
	EmailTemplate                EmailTemplate          `json:"emailTemplate,omitempty"`
	EmployeeCompNotExceed        string                 `json:"employeeCompNotExceed,omitempty"`
	EmployeeCompRate             string                 `json:"employeeCompRate,omitempty"`
	ExclusionWorkRoleIds         []int                  `json:"exclusionWorkRoleIds,omitempty"`
	ExclusionWorkTypeIds         []int                  `json:"exclusionWorkTypeIds,omitempty"`
	ExpireWhenZero               bool                   `json:"expireWhenZero,omitempty"`
	ExpiredDays                  int                    `json:"expiredDays,omitempty"`
	ID                           int                    `json:"id,omitempty"`
	InactiveFlag                 bool                   `json:"inactiveFlag,omitempty"`
	IntegrationXRef              string                 `json:"integrationXRef,omitempty"`
	InvoiceDescription           string                 `json:"invoiceDescription,omitempty"`
	InvoicePreSuffix             string                 `json:"invoicePreSuffix,omitempty"`
	InvoiceProratedAdditionsFlag bool                   `json:"invoiceProratedAdditionsFlag,omitempty"`
	InvoiceTemplate              InvoiceTemplate        `json:"invoiceTemplate,omitempty"`
	InvoicingCycle               string                 `json:"invoicingCycle,omitempty"`
	Limit                        int                    `json:"limit,omitempty"`
	Location                     SystemLocation         `json:"location,omitempty"`
	Name                         string                 `json:"name"`
	OneTimeFlag                  bool                   `json:"oneTimeFlag,omitempty"`
	PrePaymentFlag               bool                   `json:"prePaymentFlag,omitempty"`
	PrefixSuffixOption           string                 `json:"prefixSuffixOption,omitempty"`
	ProjectType                  ProjectType            `json:"projectType,omitempty"`
	ProrateFlag                  bool                   `json:"prorateFlag,omitempty"`
	RemoveAllWorkRoleExclusions  bool                   `json:"removeAllWorkRoleExclusions,omitempty"`
	RemoveAllWorkTypeExclusions  bool                   `json:"removeAllWorkTypeExclusions,omitempty"`
	RestrictDepartmentFlag       bool                   `json:"restrictDepartmentFlag,omitempty"`
	RestrictDownPaymentFlag      bool                   `json:"restrictDownPaymentFlag,omitempty"`
	RestrictLocationFlag         bool                   `json:"restrictLocationFlag,omitempty"`
	Sla                          SLA                    `json:"sla,omitempty"`
	TaxableFlag                  bool                   `json:"taxableFlag,omitempty"`
	TopCommentFlag               bool                   `json:"topCommentFlag,omitempty"`
	WorkRole                     WorkRole               `json:"workRole,omitempty"`
	WorkType                     WorkType               `json:"workType,omitempty"`
}

type AgreementTypeBoardDefault struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	Board       Board                  `json:"board,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	Department  SystemDepartment       `json:"department,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Location    SystemLocation         `json:"location"`
	ServiceType ServiceType            `json:"serviceType,omitempty"`
	Type        AgreementType          `json:"type,omitempty"`
}

type AgreementTypeInfo struct {
	Info             map[string]interface{} `json:"_info,omitempty"`
	ApplicationUnits string                 `json:"applicationUnits,omitempty"`
	BillingTerms     BillingTerms           `json:"billingTerms,omitempty"`
	ID               int                    `json:"id,omitempty"`
	InactiveFlag     bool                   `json:"inactiveFlag,omitempty"`
	Name             string                 `json:"name,omitempty"`
}

type AgreementTypeWorkRole struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	EffectiveDate time.Time              `json:"effectiveDate,omitempty"`
	EndingDate    time.Time              `json:"endingDate,omitempty"`
	ID            int                    `json:"id,omitempty"`
	LimitTo       float64                `json:"limitTo,omitempty"`
	Rate          float64                `json:"rate,omitempty"`
	RateType      string                 `json:"rateType,omitempty"`
	Type          AgreementType          `json:"type,omitempty"`
	WorkRole      WorkRole               `json:"workRole,omitempty"`
}

type AgreementTypeWorkRoleExclusion struct {
	Info     map[string]interface{} `json:"_info,omitempty"`
	ID       int                    `json:"id,omitempty"`
	Type     AgreementType          `json:"type,omitempty"`
	WorkRole WorkRole               `json:"workRole"`
}

type AgreementTypeWorkRoleInfo struct {
	Info     map[string]interface{} `json:"_info,omitempty"`
	ID       int                    `json:"id,omitempty"`
	Type     AgreementType          `json:"type,omitempty"`
	WorkRole WorkRole               `json:"workRole,omitempty"`
}

type AgreementTypeWorkType struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	BillTime        string                 `json:"billTime,omitempty"`
	EffectiveDate   time.Time              `json:"effectiveDate,omitempty"`
	EndingDate      time.Time              `json:"endingDate,omitempty"`
	HoursMax        float64                `json:"hoursMax,omitempty"`
	HoursMin        float64                `json:"hoursMin,omitempty"`
	ID              int                    `json:"id,omitempty"`
	LimitTo         float64                `json:"limitTo,omitempty"`
	OverageRate     float64                `json:"overageRate,omitempty"`
	OverageRateType string                 `json:"overageRateType,omitempty"`
	Rate            float64                `json:"rate,omitempty"`
	RateType        string                 `json:"rateType,omitempty"`
	RoundBillHours  float64                `json:"roundBillHours,omitempty"`
	Type            AgreementType          `json:"type,omitempty"`
	WorkType        WorkType               `json:"workType,omitempty"`
}

type AgreementTypeWorkTypeExclusion struct {
	Info     map[string]interface{} `json:"_info,omitempty"`
	ID       int                    `json:"id,omitempty"`
	Type     AgreementType          `json:"type,omitempty"`
	WorkType WorkType               `json:"workType"`
}

type AgreementWorkRole struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	Agreement     Agreement              `json:"agreement,omitempty"`
	AgreementId   int                    `json:"agreementId,omitempty"`
	EffectiveDate time.Time              `json:"effectiveDate,omitempty"`
	EndingDate    time.Time              `json:"endingDate,omitempty"`
	ID            int                    `json:"id,omitempty"`
	LimitTo       float64                `json:"limitTo,omitempty"`
	Location      OwnerLevel             `json:"location,omitempty"`
	LocationId    int                    `json:"locationId,omitempty"`
	Rate          float64                `json:"rate,omitempty"`
	RateType      string                 `json:"rateType,omitempty"`
	WorkRole      WorkRole               `json:"workRole,omitempty"`
}

type AgreementWorkRoleExclusion struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	AgreementId int                    `json:"agreementId,omitempty"`
	ID          int                    `json:"id,omitempty"`
	WorkRole    WorkRole               `json:"workRole"`
}

type AgreementWorkType struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	Agreement       Agreement              `json:"agreement,omitempty"`
	AgreementId     int                    `json:"agreementId,omitempty"`
	AgreementLimit  float64                `json:"agreementLimit,omitempty"`
	BillTime        string                 `json:"billTime,omitempty"`
	Company         Company                `json:"company,omitempty"`
	EffectiveDate   time.Time              `json:"effectiveDate,omitempty"`
	EndingDate      time.Time              `json:"endingDate,omitempty"`
	HoursMax        float64                `json:"hoursMax,omitempty"`
	HoursMin        float64                `json:"hoursMin,omitempty"`
	ID              int                    `json:"id,omitempty"`
	Location        OwnerLevel             `json:"location,omitempty"`
	LocationId      int                    `json:"locationId,omitempty"`
	OverageRate     float64                `json:"overageRate,omitempty"`
	OverageRateType string                 `json:"overageRateType,omitempty"`
	Rate            float64                `json:"rate,omitempty"`
	RateType        string                 `json:"rateType,omitempty"`
	RoundBillHours  float64                `json:"roundBillHours,omitempty"`
	Site            Site                   `json:"site,omitempty"`
	WorkType        WorkType               `json:"workType,omitempty"`
}

type AgreementWorkTypeExclusion struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	AgreementId int                    `json:"agreementId,omitempty"`
	ID          int                    `json:"id,omitempty"`
	WorkType    WorkType               `json:"workType"`
}

type BatchEntry struct {
	Info                         map[string]interface{} `json:"_info,omitempty"`
	AccountNumber                string                 `json:"accountNumber,omitempty"`
	AccountType                  string                 `json:"accountType,omitempty"`
	AdjustmentDetail             AdjustmentDetail       `json:"adjustmentDetail,omitempty"`
	Cost                         float64                `json:"cost,omitempty"`
	CostOfGoodsSoldAccountNumber string                 `json:"costOfGoodsSoldAccountNumber,omitempty"`
	Credit                       float64                `json:"credit,omitempty"`
	Debit                        float64                `json:"debit,omitempty"`
	Expense                      ExpenseDetail          `json:"expense,omitempty"`
	ID                           int                    `json:"id,omitempty"`
	Invoice                      Invoice                `json:"invoice,omitempty"`
	Item                         string                 `json:"item,omitempty"`
	LineItem                     PurchaseOrderLineItem  `json:"lineItem,omitempty"`
	Name                         string                 `json:"name,omitempty"`
	PurchaseOrder                PurchaseOrder          `json:"purchaseOrder,omitempty"`
	SalesCode                    string                 `json:"salesCode,omitempty"`
	Transfer                     string                 `json:"transfer,omitempty"`
}

type BillingCycle struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	BillingOptions string                 `json:"billingOptions,omitempty"`
	DefaultFlag    bool                   `json:"defaultFlag,omitempty"`
	ID             int                    `json:"id,omitempty"`
	Identifier     string                 `json:"identifier"`
	Name           string                 `json:"name"`
}

type BillingCycleInfo struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
}

type BillingSetup struct {
	Info                             map[string]interface{} `json:"_info,omitempty"`
	AddressOne                       string                 `json:"addressOne,omitempty"`
	AddressTwo                       string                 `json:"addressTwo,omitempty"`
	AgreementInvoice                 InvoiceTemplate        `json:"agreementInvoice,omitempty"`
	AllowRestrictedDeptOnRoutingFlag bool                   `json:"allowRestrictedDeptOnRoutingFlag,omitempty"`
	AttachXmlInvoiceFlag             bool                   `json:"attachXmlInvoiceFlag,omitempty"`
	BillProductAfterShipFlag         bool                   `json:"billProductAfterShipFlag,omitempty"`
	BillProjectCompleteFlag          bool                   `json:"billProjectCompleteFlag,omitempty"`
	BillProjectUnapprovedFlag        bool                   `json:"billProjectUnapprovedFlag,omitempty"`
	BillSalesOrderCompleteFlag       bool                   `json:"billSalesOrderCompleteFlag,omitempty"`
	BillTicketCompleteFlag           bool                   `json:"billTicketCompleteFlag,omitempty"`
	BillTicketSeparatelyFlag         bool                   `json:"billTicketSeparatelyFlag,omitempty"`
	BillTicketUnapprovedFlag         bool                   `json:"billTicketUnapprovedFlag,omitempty"`
	BusinessNumber                   string                 `json:"businessNumber,omitempty"`
	ChargeAdjToFirmFlag              bool                   `json:"chargeAdjToFirmFlag,omitempty"`
	City                             string                 `json:"city,omitempty"`
	CompanyCode                      string                 `json:"companyCode,omitempty"`
	CopyAgreementProductsFlag        bool                   `json:"copyAgreementProductsFlag,omitempty"`
	CopyNonServiceProductsFlag       bool                   `json:"copyNonServiceProductsFlag,omitempty"`
	CopyServiceProductsFlag          bool                   `json:"copyServiceProductsFlag,omitempty"`
	Country                          Country                `json:"country,omitempty"`
	CreditMemoInvoice                InvoiceTemplate        `json:"creditMemoInvoice,omitempty"`
	Currency                         Currency               `json:"currency,omitempty"`
	CustomLabel                      string                 `json:"customLabel,omitempty"`
	CustomText                       string                 `json:"customText,omitempty"`
	DeliveryReceiptFlag              bool                   `json:"deliveryReceiptFlag,omitempty"`
	DisableRoutingEmailFlag          bool                   `json:"disableRoutingEmailFlag,omitempty"`
	DisplayTaxFlag                   bool                   `json:"displayTaxFlag,omitempty"`
	DownPaymentInvoice               InvoiceTemplate        `json:"downPaymentInvoice,omitempty"`
	EmailTemplate                    EmailTemplate          `json:"emailTemplate"`
	ExcludeAvalaraFlag               bool                   `json:"excludeAvalaraFlag,omitempty"`
	ExcludeDoNotBillExpenseFlag      bool                   `json:"excludeDoNotBillExpenseFlag,omitempty"`
	ExcludeDoNotBillProductFlag      bool                   `json:"excludeDoNotBillProductFlag,omitempty"`
	ExcludeDoNotBillTimeFlag         bool                   `json:"excludeDoNotBillTimeFlag,omitempty"`
	ID                               int                    `json:"id,omitempty"`
	InvoiceFooter                    string                 `json:"invoiceFooter,omitempty"`
	InvoiceTitle                     string                 `json:"invoiceTitle"`
	LocalizedCountry                 Country                `json:"localizedCountry,omitempty"`
	Location                         SystemLocation         `json:"location"`
	MiscInvoice                      InvoiceTemplate        `json:"miscInvoice,omitempty"`
	NoWatermarkFlag                  bool                   `json:"noWatermarkFlag,omitempty"`
	OverallInvoiceDefault            InvoiceTemplate        `json:"overallInvoiceDefault"`
	PayableName                      string                 `json:"payableName"`
	Phone                            string                 `json:"phone,omitempty"`
	PrefixSuffixFlag                 string                 `json:"prefixSuffixFlag,omitempty"`
	PrefixSuffixText                 string                 `json:"prefixSuffixText,omitempty"`
	PrintLogoFlag                    bool                   `json:"printLogoFlag,omitempty"`
	ProgressInvoice                  InvoiceTemplate        `json:"progressInvoice,omitempty"`
	ProgressTimeFlag                 bool                   `json:"progressTimeFlag,omitempty"`
	QuoteFooter                      string                 `json:"quoteFooter,omitempty"`
	ReadReceiptFlag                  bool                   `json:"readReceiptFlag,omitempty"`
	RemitName                        string                 `json:"remitName"`
	RestrictDownpaymentFlag          bool                   `json:"restrictDownpaymentFlag,omitempty"`
	RestrictProjectDownpaymentFlag   bool                   `json:"restrictProjectDownpaymentFlag,omitempty"`
	SalesOrderInvoice                InvoiceTemplate        `json:"salesOrderInvoice,omitempty"`
	StandardInvoiceActual            InvoiceTemplate        `json:"standardInvoiceActual,omitempty"`
	StandardInvoiceFixed             InvoiceTemplate        `json:"standardInvoiceFixed,omitempty"`
	State                            State                  `json:"state,omitempty"`
	Topcomment                       string                 `json:"topcomment,omitempty"`
	Zip                              string                 `json:"zip,omitempty"`
}

type BillingSetupInfo struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	Currency  Currency               `json:"currency,omitempty"`
	ID        int                    `json:"id,omitempty"`
	Location  SystemLocation         `json:"location,omitempty"`
	RemitName string                 `json:"remitName,omitempty"`
}

type BillingSetupRouting struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	ID             int                    `json:"id,omitempty"`
	InvoiceRule    string                 `json:"invoiceRule,omitempty"`
	Member         Member                 `json:"member,omitempty"`
	RoutingRule    string                 `json:"routingRule,omitempty"`
	SequenceNumber int                    `json:"sequenceNumber,omitempty"`
}

type BillingStatus struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag   bool                   `json:"closedFlag,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name"`
	SentFlag     bool                   `json:"sentFlag,omitempty"`
	SortOrder    int                    `json:"sortOrder,omitempty"`
}

type BillingStatusInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
	SortOrder    int                    `json:"sortOrder,omitempty"`
}

type BillingTerm struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	DueDays     int                    `json:"dueDays,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name"`
	TermsXref   string                 `json:"termsXref,omitempty"`
}

type BillingTermInfo struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
}

type BoardDefault struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	AgreementId int                    `json:"agreementId,omitempty"`
	Board       Board                  `json:"board"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	ServiceType ServiceType            `json:"serviceType,omitempty"`
}

type ClosedInvoice struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	ID            int                    `json:"id,omitempty"`
	InternalNotes string                 `json:"internalNotes,omitempty"`
	Status        BillingStatus          `json:"status,omitempty"`
}

type CompanyFinance struct {
	Info                 map[string]interface{} `json:"_info,omitempty"`
	BillCompletePmFlag   bool                   `json:"billCompletePmFlag,omitempty"`
	BillCompleteSrFlag   bool                   `json:"billCompleteSrFlag,omitempty"`
	BillOverrideFlag     bool                   `json:"billOverrideFlag,omitempty"`
	BillRestrictPmFlag   bool                   `json:"billRestrictPmFlag,omitempty"`
	BillSrFlag           bool                   `json:"billSrFlag,omitempty"`
	BillUnapprovedPmFlag bool                   `json:"billUnapprovedPmFlag,omitempty"`
	BillUnapprovedSrFlag bool                   `json:"billUnapprovedSrFlag,omitempty"`
	Company              Company                `json:"company,omitempty"`
	CustomFields         []CustomFieldValue     `json:"customFields,omitempty"`
	EmailTemplate        EmailTemplate          `json:"emailTemplate,omitempty"`
	ID                   int                    `json:"id,omitempty"`
}

type Configuration struct {
	Info             map[string]interface{} `json:"_info,omitempty"`
	DeviceIdentifier string                 `json:"deviceIdentifier,omitempty"`
	ID               int                    `json:"id,omitempty"`
}

type CreateAccountingBatchRequest struct {
	BatchIdentifier       string `json:"batchIdentifier,omitempty"`
	ExportExpensesFlag    bool   `json:"exportExpensesFlag,omitempty"`
	ExportInvoicesFlag    bool   `json:"exportInvoicesFlag,omitempty"`
	ExportProductsFlag    bool   `json:"exportProductsFlag,omitempty"`
	GlInterfaceIdentifier string `json:"glInterfaceIdentifier,omitempty"`
	ID                    int    `json:"id,omitempty"`
	ProcessedRecordIds    []int  `json:"processedRecordIds"`
	SummarizeExpenses     bool   `json:"summarizeExpenses,omitempty"`
}

type CurrencyCode struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type CurrencyInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type DeliveryMethod struct {
	Info                  map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag           bool                   `json:"defaultFlag,omitempty"`
	EmailFlag             bool                   `json:"emailFlag,omitempty"`
	ID                    int                    `json:"id,omitempty"`
	IntegrationActiveFlag bool                   `json:"integrationActiveFlag,omitempty"`
	IntegrationEmailFlag  bool                   `json:"integrationEmailFlag,omitempty"`
	IntegrationPrintFlag  bool                   `json:"integrationPrintFlag,omitempty"`
	Name                  string                 `json:"name"`
}

type ExpenseTypeExemption struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	ExpenseType   ExpenseType            `json:"expenseType"`
	ID            int                    `json:"id,omitempty"`
	TaxableLevels []int                  `json:"taxableLevels,omitempty"`
}

type ExportAccountingBatchRequest struct {
	BatchIdentifier       string    `json:"batchIdentifier,omitempty"`
	ExcludedExpenseIds    []int     `json:"excludedExpenseIds,omitempty"`
	ExcludedInvoiceIds    []int     `json:"excludedInvoiceIds,omitempty"`
	ExcludedProductIds    []string  `json:"excludedProductIds,omitempty"`
	ExportExpensesFlag    bool      `json:"exportExpensesFlag,omitempty"`
	ExportInvoicesFlag    bool      `json:"exportInvoicesFlag,omitempty"`
	ExportPaymentsFlag    bool      `json:"exportPaymentsFlag,omitempty"`
	ExportProductsFlag    bool      `json:"exportProductsFlag,omitempty"`
	GlInterfaceIdentifier string    `json:"glInterfaceIdentifier,omitempty"`
	IncludedExpenseIds    []int     `json:"includedExpenseIds,omitempty"`
	IncludedInvoiceIds    []int     `json:"includedInvoiceIds,omitempty"`
	IncludedPaymentIds    []int     `json:"includedPaymentIds,omitempty"`
	IncludedProductIds    []string  `json:"includedProductIds,omitempty"`
	LocationId            int       `json:"locationId,omitempty"`
	SummarizeInvoices     string    `json:"summarizeInvoices,omitempty"`
	ThruDate              time.Time `json:"thruDate,omitempty"`
}

type FinanceCurrency struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	CurrencyCode            CurrencyCode           `json:"currencyCode,omitempty"`
	CurrencyIdentifier      string                 `json:"currencyIdentifier"`
	DecimalSeparator        string                 `json:"decimalSeparator,omitempty"`
	DisplayIdFlag           bool                   `json:"displayIdFlag,omitempty"`
	DisplaySymbolFlag       bool                   `json:"displaySymbolFlag,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	Name                    string                 `json:"name"`
	NegativeParenthesesFlag bool                   `json:"negativeParenthesesFlag,omitempty"`
	NumberOfDecimals        int                    `json:"numberOfDecimals,omitempty"`
	ReportFormat            string                 `json:"reportFormat,omitempty"`
	RightAlign              bool                   `json:"rightAlign,omitempty"`
	Symbol                  string                 `json:"symbol,omitempty"`
	ThousandsSeparator      string                 `json:"thousandsSeparator,omitempty"`
}

type GLAccount struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	Cogs1        string                 `json:"cogs1,omitempty"`
	Cogs10       string                 `json:"cogs10,omitempty"`
	Cogs2        string                 `json:"cogs2,omitempty"`
	Cogs3        string                 `json:"cogs3,omitempty"`
	Cogs4        string                 `json:"cogs4,omitempty"`
	Cogs5        string                 `json:"cogs5,omitempty"`
	Cogs6        string                 `json:"cogs6,omitempty"`
	Cogs7        string                 `json:"cogs7,omitempty"`
	Cogs8        string                 `json:"cogs8,omitempty"`
	Cogs9        string                 `json:"cogs9,omitempty"`
	GlType       string                 `json:"glType,omitempty"`
	ID           int                    `json:"id,omitempty"`
	Inventory    string                 `json:"inventory,omitempty"`
	MappedRecord MappedRecord           `json:"mappedRecord"`
	MappedType   MappedType             `json:"mappedType"`
	ProductId    string                 `json:"productId,omitempty"`
	SalesCode    string                 `json:"salesCode,omitempty"`
	Segment1     string                 `json:"segment1,omitempty"`
	Segment10    string                 `json:"segment10,omitempty"`
	Segment2     string                 `json:"segment2,omitempty"`
	Segment3     string                 `json:"segment3,omitempty"`
	Segment4     string                 `json:"segment4,omitempty"`
	Segment5     string                 `json:"segment5,omitempty"`
	Segment6     string                 `json:"segment6,omitempty"`
	Segment7     string                 `json:"segment7,omitempty"`
	Segment8     string                 `json:"segment8,omitempty"`
	Segment9     string                 `json:"segment9,omitempty"`
}

type GLCaption struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	Cogs1         string                 `json:"cogs1,omitempty"`
	Cogs10        string                 `json:"cogs10,omitempty"`
	Cogs2         string                 `json:"cogs2,omitempty"`
	Cogs3         string                 `json:"cogs3,omitempty"`
	Cogs4         string                 `json:"cogs4,omitempty"`
	Cogs5         string                 `json:"cogs5,omitempty"`
	Cogs6         string                 `json:"cogs6,omitempty"`
	Cogs7         string                 `json:"cogs7,omitempty"`
	Cogs8         string                 `json:"cogs8,omitempty"`
	Cogs9         string                 `json:"cogs9,omitempty"`
	ID            int                    `json:"id,omitempty"`
	Segment1      string                 `json:"segment1,omitempty"`
	Segment10     string                 `json:"segment10,omitempty"`
	Segment10type string                 `json:"segment10type,omitempty"`
	Segment1type  string                 `json:"segment1type,omitempty"`
	Segment2      string                 `json:"segment2,omitempty"`
	Segment2type  string                 `json:"segment2type,omitempty"`
	Segment3      string                 `json:"segment3,omitempty"`
	Segment3type  string                 `json:"segment3type,omitempty"`
	Segment4      string                 `json:"segment4,omitempty"`
	Segment4type  string                 `json:"segment4type,omitempty"`
	Segment5      string                 `json:"segment5,omitempty"`
	Segment5type  string                 `json:"segment5type,omitempty"`
	Segment6      string                 `json:"segment6,omitempty"`
	Segment6type  string                 `json:"segment6type,omitempty"`
	Segment7      string                 `json:"segment7,omitempty"`
	Segment7type  string                 `json:"segment7type,omitempty"`
	Segment8      string                 `json:"segment8,omitempty"`
	Segment8type  string                 `json:"segment8type,omitempty"`
	Segment9      string                 `json:"segment9,omitempty"`
	Segment9type  string                 `json:"segment9type,omitempty"`
}

type GLEntry struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	Amount      float64                `json:"amount,omitempty"`
	Cogs1       string                 `json:"cogs1,omitempty"`
	Cogs10      string                 `json:"cogs10,omitempty"`
	Cogs2       string                 `json:"cogs2,omitempty"`
	Cogs3       string                 `json:"cogs3,omitempty"`
	Cogs4       string                 `json:"cogs4,omitempty"`
	Cogs5       string                 `json:"cogs5,omitempty"`
	Cogs6       string                 `json:"cogs6,omitempty"`
	Cogs7       string                 `json:"cogs7,omitempty"`
	Cogs8       string                 `json:"cogs8,omitempty"`
	Cogs9       string                 `json:"cogs9,omitempty"`
	Cost        float64                `json:"cost,omitempty"`
	Description string                 `json:"description,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Inventory   string                 `json:"inventory,omitempty"`
	IsBatched   bool                   `json:"isBatched,omitempty"`
	ProductId   string                 `json:"productId,omitempty"`
	SalesCode   string                 `json:"salesCode,omitempty"`
	Segment1    string                 `json:"segment1,omitempty"`
	Segment10   string                 `json:"segment10,omitempty"`
	Segment2    string                 `json:"segment2,omitempty"`
	Segment3    string                 `json:"segment3,omitempty"`
	Segment4    string                 `json:"segment4,omitempty"`
	Segment5    string                 `json:"segment5,omitempty"`
	Segment6    string                 `json:"segment6,omitempty"`
	Segment7    string                 `json:"segment7,omitempty"`
	Segment8    string                 `json:"segment8,omitempty"`
	Segment9    string                 `json:"segment9,omitempty"`
	Type        string                 `json:"type,omitempty"`
}

type GLExport struct {
	AdjustmentTransactions []GLExportAdjustmentTransaction `json:"adjustmentTransactions,omitempty"`
	Customers              []GLExportCustomer              `json:"customers,omitempty"`
	ExpenseBills           []GLExportExpenseBill           `json:"expenseBills,omitempty"`
	Expenses               []GLExportExpense               `json:"expenses,omitempty"`
	ExportSettings         GLExportSettings                `json:"exportSettings,omitempty"`
	InventoryTransfers     []GLExportInventoryTransfer     `json:"inventoryTransfers,omitempty"`
	PurchaseTransactions   []GLExportPurchaseTransaction   `json:"purchaseTransactions,omitempty"`
	Transactions           []GLExportTransaction           `json:"transactions,omitempty"`
	Vendors                []GLExportVendor                `json:"vendors,omitempty"`
}

type GLPath struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	DatabaseName      string                 `json:"databaseName,omitempty"`
	ID                int                    `json:"id,omitempty"`
	LastPaymentSync   time.Time              `json:"lastPaymentSync,omitempty"`
	LastPaymentSyncBy Member                 `json:"lastPaymentSyncBy,omitempty"`
	Location          SystemLocation         `json:"location,omitempty"`
	Path              string                 `json:"path,omitempty"`
	SqlServerName     string                 `json:"sqlServerName,omitempty"`
}

type Invoice struct {
	Info                           map[string]interface{} `json:"_info,omitempty"`
	AccountNumber                  string                 `json:"accountNumber,omitempty"`
	AddToBatchEmailList            bool                   `json:"addToBatchEmailList,omitempty"`
	AdjustedBy                     string                 `json:"adjustedBy,omitempty"`
	AdjustmentReason               string                 `json:"adjustmentReason,omitempty"`
	Agreement                      Agreement              `json:"agreement,omitempty"`
	AgreementAmount                float64                `json:"agreementAmount,omitempty"`
	ApplyToId                      int                    `json:"applyToId,omitempty"`
	ApplyToType                    string                 `json:"applyToType,omitempty"`
	Attention                      string                 `json:"attention,omitempty"`
	Balance                        float64                `json:"balance,omitempty"`
	BillToCompany                  Company                `json:"billToCompany,omitempty"`
	BillingSetup                   BillingSetup           `json:"billingSetup,omitempty"`
	BillingSite                    Site                   `json:"billingSite,omitempty"`
	BillingSiteAddressLine1        string                 `json:"billingSiteAddressLine1,omitempty"`
	BillingSiteAddressLine2        string                 `json:"billingSiteAddressLine2,omitempty"`
	BillingSiteCity                string                 `json:"billingSiteCity,omitempty"`
	BillingSiteCountry             string                 `json:"billingSiteCountry,omitempty"`
	BillingSiteState               string                 `json:"billingSiteState,omitempty"`
	BillingSiteZip                 string                 `json:"billingSiteZip,omitempty"`
	BillingTerms                   BillingTerms           `json:"billingTerms,omitempty"`
	BottomComment                  string                 `json:"bottomComment,omitempty"`
	ClosedBy                       string                 `json:"closedBy,omitempty"`
	Company                        Company                `json:"company"`
	Credits                        float64                `json:"credits,omitempty"`
	Currency                       Currency               `json:"currency,omitempty"`
	CustomFields                   []CustomFieldValue     `json:"customFields,omitempty"`
	CustomerPO                     string                 `json:"customerPO,omitempty"`
	Date                           time.Time              `json:"date,omitempty"`
	Department                     BillingUnit            `json:"department,omitempty"`
	DepartmentId                   int                    `json:"departmentId,omitempty"`
	DownpaymentApplied             float64                `json:"downpaymentApplied,omitempty"`
	DownpaymentPreviouslyTaxedFlag bool                   `json:"downpaymentPreviouslyTaxedFlag,omitempty"`
	DueDate                        time.Time              `json:"dueDate,omitempty"`
	EmailTemplateId                int                    `json:"emailTemplateId,omitempty"`
	ExpenseTotal                   float64                `json:"expenseTotal,omitempty"`
	GlBatch                        Batch                  `json:"glBatch,omitempty"`
	ID                             int                    `json:"id,omitempty"`
	InternalNotes                  string                 `json:"internalNotes,omitempty"`
	InvoiceNumber                  string                 `json:"invoiceNumber,omitempty"`
	InvoiceTemplate                InvoiceTemplateDetail  `json:"invoiceTemplate,omitempty"`
	Location                       SystemLocation         `json:"location,omitempty"`
	LocationId                     int                    `json:"locationId,omitempty"`
	OverrideDownPaymentAmountFlag  bool                   `json:"overrideDownPaymentAmountFlag,omitempty"`
	Payments                       float64                `json:"payments,omitempty"`
	Phase                          ProjectPhase           `json:"phase,omitempty"`
	PreviousProgressApplied        float64                `json:"previousProgressApplied,omitempty"`
	ProductTotal                   float64                `json:"productTotal,omitempty"`
	Project                        Project                `json:"project,omitempty"`
	string                         `json:",omitempty"`
	RemainingDownpayment           float64        `json:"remainingDownpayment,omitempty"`
	RestrictDownpaymentFlag        bool           `json:"restrictDownpaymentFlag,omitempty"`
	SalesOrder                     SalesOrder     `json:"salesOrder,omitempty"`
	SalesTax                       float64        `json:"salesTax,omitempty"`
	ServiceAdjustmentAmount        float64        `json:"serviceAdjustmentAmount,omitempty"`
	ServiceTotal                   float64        `json:"serviceTotal,omitempty"`
	ShipToAttention                string         `json:"shipToAttention,omitempty"`
	ShipToCompany                  Company        `json:"shipToCompany,omitempty"`
	ShippingSite                   Site           `json:"shippingSite,omitempty"`
	ShippingSiteAddressLine1       string         `json:"shippingSiteAddressLine1,omitempty"`
	ShippingSiteAddressLine2       string         `json:"shippingSiteAddressLine2,omitempty"`
	ShippingSiteCity               string         `json:"shippingSiteCity,omitempty"`
	ShippingSiteCountry            string         `json:"shippingSiteCountry,omitempty"`
	ShippingSiteState              string         `json:"shippingSiteState,omitempty"`
	ShippingSiteZip                string         `json:"shippingSiteZip,omitempty"`
	SpecialInvoiceFlag             bool           `json:"specialInvoiceFlag,omitempty"`
	Status                         BillingStatus  `json:"status,omitempty"`
	Subtotal                       float64        `json:"subtotal,omitempty"`
	TaxCode                        TaxCode        `json:"taxCode,omitempty"`
	TaxableFlag                    bool           `json:"taxableFlag,omitempty"`
	TemplateSetupId                int            `json:"templateSetupId,omitempty"`
	Territory                      SystemLocation `json:"territory,omitempty"`
	TerritoryId                    int            `json:"territoryId,omitempty"`
	Ticket                         Ticket         `json:"ticket,omitempty"`
	TopComment                     string         `json:"topComment,omitempty"`
	Total                          float64        `json:"total,omitempty"`
	Type                           string         `json:"type,omitempty"`
	UnbatchedBatch                 Batch          `json:"unbatchedBatch,omitempty"`
}

type InvoiceCommission struct {
	Info             map[string]interface{} `json:"_info,omitempty"`
	Activity         Activity               `json:"activity,omitempty"`
	AdjustedBy       string                 `json:"adjustedBy,omitempty"`
	AdjustedDate     string                 `json:"adjustedDate,omitempty"`
	Adjustment       float64                `json:"adjustment,omitempty"`
	AdjustmentReason string                 `json:"adjustmentReason,omitempty"`
	Agreement        Agreement              `json:"agreement,omitempty"`
	Amount           float64                `json:"amount,omitempty"`
	ID               int                    `json:"id,omitempty"`
	Invoice          Invoice                `json:"invoice,omitempty"`
	Member           Member                 `json:"member,omitempty"`
	NetAmount        float64                `json:"netAmount,omitempty"`
	Opportunity      Opportunity            `json:"opportunity,omitempty"`
	Percent          float64                `json:"percent,omitempty"`
	Project          Project                `json:"project,omitempty"`
	SalesOrder       SalesOrder             `json:"salesOrder,omitempty"`
	SplitPercent     float64                `json:"splitPercent,omitempty"`
	Ticket           Ticket                 `json:"ticket,omitempty"`
}

type InvoiceEmailTemplate struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	AttachInvoiceFlag bool                   `json:"attachInvoiceFlag,omitempty"`
	Body              string                 `json:"body,omitempty"`
	CopySenderFlag    bool                   `json:"copySenderFlag,omitempty"`
	EmailAddress      string                 `json:"emailAddress,omitempty"`
	FirstName         string                 `json:"firstName,omitempty"`
	ID                int                    `json:"id,omitempty"`
	InvoiceStatus     BillingStatus          `json:"invoiceStatus,omitempty"`
	LastName          string                 `json:"lastName,omitempty"`
	Name              string                 `json:"name"`
	ServiceSurvey     ServiceSurvey          `json:"serviceSurvey,omitempty"`
	Subject           string                 `json:"subject"`
	UseSenderFlag     bool                   `json:"useSenderFlag,omitempty"`
}

type InvoiceEmailTemplateInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type InvoiceInfo struct {
	Info                  map[string]interface{} `json:"_info,omitempty"`
	AgreementBillingInfo  []AgreementBillingInfo `json:"agreementBillingInfo,omitempty"`
	BillingSetup          BillingSetup           `json:"billingSetup,omitempty"`
	BundledComponentsInfo []ProductComponent     `json:"bundledComponentsInfo,omitempty"`
	Expenses              []ExpenseEntry         `json:"expenses,omitempty"`
	ID                    int                    `json:"id,omitempty"`
	Invoice               Invoice                `json:"invoice,omitempty"`
	InvoiceTemplate       InvoiceTemplate        `json:"invoiceTemplate,omitempty"`
	Logo                  DocumentInfo           `json:"logo,omitempty"`
	Products              []ProductItem          `json:"products,omitempty"`
	TimeEntries           []TimeEntry            `json:"timeEntries,omitempty"`
}

type InvoicePayment struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	ARPaymentAccount  string                 `json:"aRPaymentAccount,omitempty"`
	Amount            float64                `json:"amount,omitempty"`
	AppliedBy         string                 `json:"appliedBy,omitempty"`
	Balance           float64                `json:"balance,omitempty"`
	Credit            Invoice                `json:"credit,omitempty"`
	GlBatchID         string                 `json:"glBatchID,omitempty"`
	ID                int                    `json:"id,omitempty"`
	Invoice           Invoice                `json:"invoice,omitempty"`
	InvoiceBalance    float64                `json:"invoiceBalance,omitempty"`
	InvoiceTotal      float64                `json:"invoiceTotal,omitempty"`
	PaymentAccount    string                 `json:"paymentAccount,omitempty"`
	PaymentDate       time.Time              `json:"paymentDate,omitempty"`
	PaymentSyncDate   time.Time              `json:"paymentSyncDate,omitempty"`
	PaymentSyncStatus string                 `json:"paymentSyncStatus,omitempty"`
	Source            string                 `json:"source,omitempty"`
	Type              string                 `json:"type,omitempty"`
	WisePayPayment    WisePayPayment         `json:"wisePayPayment,omitempty"`
}

type InvoiceRouting struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	DateReviewedUTC time.Time              `json:"dateReviewedUTC,omitempty"`
	ID              int                    `json:"id,omitempty"`
	Invoice         Invoice                `json:"invoice,omitempty"`
	Member          Member                 `json:"member,omitempty"`
	ReviewedFlag    bool                   `json:"reviewedFlag,omitempty"`
	SequenceNumber  int                    `json:"sequenceNumber,omitempty"`
}

type InvoiceTemplate struct {
	Info                                            map[string]interface{} `json:"_info,omitempty"`
	AdjustmentAgrTypeVisibleFlag                    bool                   `json:"adjustmentAgrTypeVisibleFlag,omitempty"`
	AdjustmentAmountCaption                         string                 `json:"adjustmentAmountCaption,omitempty"`
	AdjustmentAmountVisibleFlag                     bool                   `json:"adjustmentAmountVisibleFlag,omitempty"`
	AdjustmentDescriptionCaption                    string                 `json:"adjustmentDescriptionCaption,omitempty"`
	AdjustmentDescriptionVisibleFlag                bool                   `json:"adjustmentDescriptionVisibleFlag,omitempty"`
	AdjustmentPriceCaption                          string                 `json:"adjustmentPriceCaption,omitempty"`
	AdjustmentPriceVisibleFlag                      bool                   `json:"adjustmentPriceVisibleFlag,omitempty"`
	AdjustmentQuantityCaption                       string                 `json:"adjustmentQuantityCaption,omitempty"`
	AdjustmentQuantityVisibleFlag                   bool                   `json:"adjustmentQuantityVisibleFlag,omitempty"`
	AdjustmentTotalVisibleFlag                      bool                   `json:"adjustmentTotalVisibleFlag,omitempty"`
	CreditCreditAmountCaption                       string                 `json:"creditCreditAmountCaption,omitempty"`
	CreditCreditAmountVisibleFlag                   bool                   `json:"creditCreditAmountVisibleFlag,omitempty"`
	CreditRemainingAmountCaption                    string                 `json:"creditRemainingAmountCaption,omitempty"`
	CreditRemainingAmountVisibleFlag                bool                   `json:"creditRemainingAmountVisibleFlag,omitempty"`
	CurrencyIdVisibleFlag                           bool                   `json:"currencyIdVisibleFlag,omitempty"`
	CurrencySymbolVisibleFlag                       bool                   `json:"currencySymbolVisibleFlag,omitempty"`
	ExpenseDetailAgreementExtAmountVisibleFlag      bool                   `json:"expenseDetailAgreementExtAmountVisibleFlag,omitempty"`
	ExpenseDetailAgreementVisibleFlag               bool                   `json:"expenseDetailAgreementVisibleFlag,omitempty"`
	ExpenseDetailBillableVisibleFlag                bool                   `json:"expenseDetailBillableVisibleFlag,omitempty"`
	ExpenseDetailContactsVisibleFlag                bool                   `json:"expenseDetailContactsVisibleFlag,omitempty"`
	ExpenseDetailExtAmountVisibleFlag               bool                   `json:"expenseDetailExtAmountVisibleFlag,omitempty"`
	ExpenseDetailMembersVisibleFlag                 bool                   `json:"expenseDetailMembersVisibleFlag,omitempty"`
	ExpenseDetailNonbillableCaption                 string                 `json:"expenseDetailNonbillableCaption,omitempty"`
	ExpenseDetailNotesVisibleFlag                   bool                   `json:"expenseDetailNotesVisibleFlag,omitempty"`
	ExpenseDetailPmPhaseVisibleFlag                 bool                   `json:"expenseDetailPmPhaseVisibleFlag,omitempty"`
	ExpenseDetailPmSummaryVisibleFlag               bool                   `json:"expenseDetailPmSummaryVisibleFlag,omitempty"`
	ExpenseDetailPrimarySortDirection               string                 `json:"expenseDetailPrimarySortDirection,omitempty"`
	ExpenseDetailPrimarySortField                   string                 `json:"expenseDetailPrimarySortField,omitempty"`
	ExpenseDetailSecondarySortDirection             string                 `json:"expenseDetailSecondarySortDirection,omitempty"`
	ExpenseDetailSecondarySortField                 string                 `json:"expenseDetailSecondarySortField,omitempty"`
	ExpenseDetailSrAddressVisibleFlag               bool                   `json:"expenseDetailSrAddressVisibleFlag,omitempty"`
	ExpenseDetailSrContactVisibleFlag               bool                   `json:"expenseDetailSrContactVisibleFlag,omitempty"`
	ExpenseDetailSrTicketSummaryVisibleFlag         bool                   `json:"expenseDetailSrTicketSummaryVisibleFlag,omitempty"`
	ExpenseDetailSubtotalVisibleFlag                bool                   `json:"expenseDetailSubtotalVisibleFlag,omitempty"`
	ExpenseDetailTicketNumberVisibleFlag            bool                   `json:"expenseDetailTicketNumberVisibleFlag,omitempty"`
	ExpenseDetailVisibleFlag                        bool                   `json:"expenseDetailVisibleFlag,omitempty"`
	ExpensesAmountCaption                           string                 `json:"expensesAmountCaption,omitempty"`
	ExpensesAmountVisibleFlag                       bool                   `json:"expensesAmountVisibleFlag,omitempty"`
	ExpensesCollapsedFlag                           bool                   `json:"expensesCollapsedFlag,omitempty"`
	ExpensesStaffCaption                            string                 `json:"expensesStaffCaption,omitempty"`
	ExpensesStaffVisibleFlag                        bool                   `json:"expensesStaffVisibleFlag,omitempty"`
	ExpensesTotalVisibleFlag                        bool                   `json:"expensesTotalVisibleFlag,omitempty"`
	ExpensesTypeCaption                             string                 `json:"expensesTypeCaption,omitempty"`
	ExpensesTypeVisibleFlag                         bool                   `json:"expensesTypeVisibleFlag,omitempty"`
	HeaderAccountCaption                            string                 `json:"headerAccountCaption,omitempty"`
	HeaderAccountVisibleFlag                        bool                   `json:"headerAccountVisibleFlag,omitempty"`
	HeaderAddressPosition                           string                 `json:"headerAddressPosition,omitempty"`
	HeaderDueDateCaption                            string                 `json:"headerDueDateCaption,omitempty"`
	HeaderDueDateVisibleFlag                        bool                   `json:"headerDueDateVisibleFlag,omitempty"`
	HeaderHoursBasedExtendedAmountVisibleFlag       bool                   `json:"headerHoursBasedExtendedAmountVisibleFlag,omitempty"`
	HeaderLogoPosition                              string                 `json:"headerLogoPosition,omitempty"`
	HeaderPoNumberCaption                           string                 `json:"headerPoNumberCaption,omitempty"`
	HeaderPoNumberVisibleFlag                       bool                   `json:"headerPoNumberVisibleFlag,omitempty"`
	HeaderCaption                                   string                 `json:"headerCaption,omitempty"`
	HeaderVisibleFlag                               bool                   `json:"headerVisibleFlag,omitempty"`
	HeaderShipToCaption                             string                 `json:"headerShipToCaption,omitempty"`
	HeaderShipToVisibleFlag                         bool                   `json:"headerShipToVisibleFlag,omitempty"`
	HeaderTaxIdCaption                              string                 `json:"headerTaxIdCaption,omitempty"`
	HeaderTaxIdVisibleFlag                          bool                   `json:"headerTaxIdVisibleFlag,omitempty"`
	HeaderTermsCaption                              string                 `json:"headerTermsCaption,omitempty"`
	HeaderTermsVisibleFlag                          bool                   `json:"headerTermsVisibleFlag,omitempty"`
	HeaderTitleCaption                              string                 `json:"headerTitleCaption,omitempty"`
	HeaderTitleFont                                 string                 `json:"headerTitleFont,omitempty"`
	HeaderTitlePosition                             string                 `json:"headerTitlePosition,omitempty"`
	HeaderTitleVisibleFlag                          bool                   `json:"headerTitleVisibleFlag,omitempty"`
	ID                                              int                    `json:"id,omitempty"`
	InvoiceBalanceDueCaption                        string                 `json:"invoiceBalanceDueCaption,omitempty"`
	InvoiceBalanceDueVisibleFlag                    bool                   `json:"invoiceBalanceDueVisibleFlag,omitempty"`
	InvoiceCreditAmountCaption                      string                 `json:"invoiceCreditAmountCaption,omitempty"`
	InvoiceCreditAmountVisibleFlag                  bool                   `json:"invoiceCreditAmountVisibleFlag,omitempty"`
	InvoicePaymentAmountCaption                     string                 `json:"invoicePaymentAmountCaption,omitempty"`
	InvoicePaymentAmountVisibleFlag                 bool                   `json:"invoicePaymentAmountVisibleFlag,omitempty"`
	LogoVisibleFlag                                 bool                   `json:"logoVisibleFlag,omitempty"`
	MarginBottom                                    float64                `json:"marginBottom,omitempty"`
	MarginLeft                                      float64                `json:"marginLeft,omitempty"`
	MarginRight                                     float64                `json:"marginRight,omitempty"`
	MarginTop                                       float64                `json:"marginTop,omitempty"`
	Name                                            string                 `json:"name"`
	OtherChargesAmountCaption                       string                 `json:"otherChargesAmountCaption,omitempty"`
	OtherChargesAmountVisibleFlag                   bool                   `json:"otherChargesAmountVisibleFlag,omitempty"`
	OtherChargesCollapsedFlag                       bool                   `json:"otherChargesCollapsedFlag,omitempty"`
	OtherChargesDescriptionCaption                  string                 `json:"otherChargesDescriptionCaption,omitempty"`
	OtherChargesDescriptionVisibleFlag              bool                   `json:"otherChargesDescriptionVisibleFlag,omitempty"`
	OtherChargesDisplaySixDecimals                  bool                   `json:"otherChargesDisplaySixDecimals,omitempty"`
	OtherChargesItemIdVisibleFlag                   bool                   `json:"otherChargesItemIdVisibleFlag,omitempty"`
	OtherChargesPriceCaption                        string                 `json:"otherChargesPriceCaption,omitempty"`
	OtherChargesPriceVisibleFlag                    bool                   `json:"otherChargesPriceVisibleFlag,omitempty"`
	OtherChargesQuantityCaption                     string                 `json:"otherChargesQuantityCaption,omitempty"`
	OtherChargesQuantityVisibleFlag                 bool                   `json:"otherChargesQuantityVisibleFlag,omitempty"`
	OtherChargesSerialNumberVisibleFlag             bool                   `json:"otherChargesSerialNumberVisibleFlag,omitempty"`
	OtherChargesTotalVisibleFlag                    bool                   `json:"otherChargesTotalVisibleFlag,omitempty"`
	PayableCaption                                  string                 `json:"payableCaption,omitempty"`
	PortalFlag                                      bool                   `json:"portalFlag,omitempty"`
	ProjectHeaderAmountCaption                      string                 `json:"projectHeaderAmountCaption,omitempty"`
	ProjectHeaderAmountVisibleFlag                  bool                   `json:"projectHeaderAmountVisibleFlag,omitempty"`
	ProjectHeaderBillingMethodCaption               string                 `json:"projectHeaderBillingMethodCaption,omitempty"`
	ProjectHeaderBillingMethodVisibleFlag           bool                   `json:"projectHeaderBillingMethodVisibleFlag,omitempty"`
	ProjectHeaderBillingTypeCaption                 string                 `json:"projectHeaderBillingTypeCaption,omitempty"`
	ProjectHeaderBillingTypeVisibleFlag             bool                   `json:"projectHeaderBillingTypeVisibleFlag,omitempty"`
	ProjectHeaderCompanyNameCaption                 string                 `json:"projectHeaderCompanyNameCaption,omitempty"`
	ProjectHeaderCompanyNameVisibleFlag             bool                   `json:"projectHeaderCompanyNameVisibleFlag,omitempty"`
	ProjectHeaderContactNameCaption                 string                 `json:"projectHeaderContactNameCaption,omitempty"`
	ProjectHeaderContactNameVisibleFlag             bool                   `json:"projectHeaderContactNameVisibleFlag,omitempty"`
	ProjectHeaderOriginalDownpaymentCaption         string                 `json:"projectHeaderOriginalDownpaymentCaption,omitempty"`
	ProjectHeaderOriginalDownpaymentVisibleFlag     bool                   `json:"projectHeaderOriginalDownpaymentVisibleFlag,omitempty"`
	ProjectHeaderProjectNameCaption                 string                 `json:"projectHeaderProjectNameCaption,omitempty"`
	ProjectHeaderProjectNameVisibleFlag             bool                   `json:"projectHeaderProjectNameVisibleFlag,omitempty"`
	RemitToVisibleFlag                              bool                   `json:"remitToVisibleFlag,omitempty"`
	ServiceHeaderAmountCaption                      string                 `json:"serviceHeaderAmountCaption,omitempty"`
	ServiceHeaderAmountVisibleFlag                  bool                   `json:"serviceHeaderAmountVisibleFlag,omitempty"`
	ServiceHeaderBillingMethodCaption               string                 `json:"serviceHeaderBillingMethodCaption,omitempty"`
	ServiceHeaderBillingMethodVisibleFlag           bool                   `json:"serviceHeaderBillingMethodVisibleFlag,omitempty"`
	ServiceHeaderBundledTicketsVisibleFlag          bool                   `json:"serviceHeaderBundledTicketsVisibleFlag,omitempty"`
	ServiceHeaderClosedTasksVisibleFlag             bool                   `json:"serviceHeaderClosedTasksVisibleFlag,omitempty"`
	ServiceHeaderCompanyNameCaption                 string                 `json:"serviceHeaderCompanyNameCaption,omitempty"`
	ServiceHeaderCompanyNameVisibleFlag             bool                   `json:"serviceHeaderCompanyNameVisibleFlag,omitempty"`
	ServiceHeaderContactNameCaption                 string                 `json:"serviceHeaderContactNameCaption,omitempty"`
	ServiceHeaderContactNameVisibleFlag             bool                   `json:"serviceHeaderContactNameVisibleFlag,omitempty"`
	ServiceHeaderDetailDescriptionCaption           string                 `json:"serviceHeaderDetailDescriptionCaption,omitempty"`
	ServiceHeaderDetailDescriptionVisibleFlag       bool                   `json:"serviceHeaderDetailDescriptionVisibleFlag,omitempty"`
	ServiceHeaderOpenTasksVisibleFlag               bool                   `json:"serviceHeaderOpenTasksVisibleFlag,omitempty"`
	ServiceHeaderResolutionCaption                  string                 `json:"serviceHeaderResolutionCaption,omitempty"`
	ServiceHeaderResolutionVisibleFlag              bool                   `json:"serviceHeaderResolutionVisibleFlag,omitempty"`
	ServiceHeaderSummaryCaption                     string                 `json:"serviceHeaderSummaryCaption,omitempty"`
	ServiceHeaderSummaryVisibleFlag                 bool                   `json:"serviceHeaderSummaryVisibleFlag,omitempty"`
	ServiceHeaderTicketNumberCaption                string                 `json:"serviceHeaderTicketNumberCaption,omitempty"`
	ServiceHeaderTicketNumberVisibleFlag            bool                   `json:"serviceHeaderTicketNumberVisibleFlag,omitempty"`
	ServicesAmountCaption                           string                 `json:"servicesAmountCaption,omitempty"`
	ServicesAmountVisibleFlag                       bool                   `json:"servicesAmountVisibleFlag,omitempty"`
	ServicesCollapsedFlag                           bool                   `json:"servicesCollapsedFlag,omitempty"`
	ServicesHoursCaption                            string                 `json:"servicesHoursCaption,omitempty"`
	ServicesHoursVisibleFlag                        bool                   `json:"servicesHoursVisibleFlag,omitempty"`
	ServicesMemberNameCaption                       string                 `json:"servicesMemberNameCaption,omitempty"`
	ServicesMemberNameVisibleFlag                   bool                   `json:"servicesMemberNameVisibleFlag,omitempty"`
	ServicesRateCaption                             string                 `json:"servicesRateCaption,omitempty"`
	ServicesRateVisibleFlag                         bool                   `json:"servicesRateVisibleFlag,omitempty"`
	ServicesStaffCaption                            string                 `json:"servicesStaffCaption,omitempty"`
	ServicesStaffVisibleFlag                        bool                   `json:"servicesStaffVisibleFlag,omitempty"`
	ServicesTotalVisibleFlag                        bool                   `json:"servicesTotalVisibleFlag,omitempty"`
	ServicesWorkRoleCaption                         string                 `json:"servicesWorkRoleCaption,omitempty"`
	ServicesWorkRoleVisibleFlag                     bool                   `json:"servicesWorkRoleVisibleFlag,omitempty"`
	ServicesWorkTypeCaption                         string                 `json:"servicesWorkTypeCaption,omitempty"`
	ServicesWorkTypeVisibleFlag                     bool                   `json:"servicesWorkTypeVisibleFlag,omitempty"`
	TimeDetailAgreementVisibleFlag                  bool                   `json:"timeDetailAgreementVisibleFlag,omitempty"`
	TimeDetailAmountBasedExtAmountVisibleFlag       bool                   `json:"timeDetailAmountBasedExtAmountVisibleFlag,omitempty"`
	TimeDetailAmountBasedHourlyRateVisibleFlag      bool                   `json:"timeDetailAmountBasedHourlyRateVisibleFlag,omitempty"`
	TimeDetailAmountBasedHoursVisibleFlag           bool                   `json:"timeDetailAmountBasedHoursVisibleFlag,omitempty"`
	TimeDetailBillableVisibleFlag                   bool                   `json:"timeDetailBillableVisibleFlag,omitempty"`
	TimeDetailContactsVisibleFlag                   bool                   `json:"timeDetailContactsVisibleFlag,omitempty"`
	TimeDetailDatesVisibleFlag                      bool                   `json:"timeDetailDatesVisibleFlag,omitempty"`
	TimeDetailDollarAmountsOnHourseBasedVisibleFlag bool                   `json:"timeDetailDollarAmountsOnHourseBasedVisibleFlag,omitempty"`
	TimeDetailExtendedAmountVisibleFlag             bool                   `json:"timeDetailExtendedAmountVisibleFlag,omitempty"`
	TimeDetailHourlyRateVisibleFlag                 bool                   `json:"timeDetailHourlyRateVisibleFlag,omitempty"`
	TimeDetailHoursBasedExtAmountVisibleFlag        bool                   `json:"timeDetailHoursBasedExtAmountVisibleFlag,omitempty"`
	TimeDetailHoursBasedHoursVisibleFlag            bool                   `json:"timeDetailHoursBasedHoursVisibleFlag,omitempty"`
	TimeDetailHoursVisibleFlag                      bool                   `json:"timeDetailHoursVisibleFlag,omitempty"`
	TimeDetailHoursbasedHourlyRateVisibleFlag       bool                   `json:"timeDetailHoursbasedHourlyRateVisibleFlag,omitempty"`
	TimeDetailMembersVisibleFlag                    bool                   `json:"timeDetailMembersVisibleFlag,omitempty"`
	TimeDetailNonBillableCaption                    string                 `json:"timeDetailNonBillableCaption,omitempty"`
	TimeDetailNotesVisibleFlag                      bool                   `json:"timeDetailNotesVisibleFlag,omitempty"`
	TimeDetailPmPhaseVisibleFlag                    bool                   `json:"timeDetailPmPhaseVisibleFlag,omitempty"`
	TimeDetailPmSummaryVisibleFlag                  bool                   `json:"timeDetailPmSummaryVisibleFlag,omitempty"`
	TimeDetailPrimarySortDirection                  string                 `json:"timeDetailPrimarySortDirection,omitempty"`
	TimeDetailPrimarySortField                      string                 `json:"timeDetailPrimarySortField,omitempty"`
	TimeDetailSRAddressVisibleFlag                  bool                   `json:"timeDetailSRAddressVisibleFlag,omitempty"`
	TimeDetailSRContactVisibleFlag                  bool                   `json:"timeDetailSRContactVisibleFlag,omitempty"`
	TimeDetailSRTicketSummaryVisibleFlag            bool                   `json:"timeDetailSRTicketSummaryVisibleFlag,omitempty"`
	TimeDetailSecondarySortDirection                string                 `json:"timeDetailSecondarySortDirection,omitempty"`
	TimeDetailSecondarySortField                    string                 `json:"timeDetailSecondarySortField,omitempty"`
	TimeDetailStartEndTimeVisibleFlag               bool                   `json:"timeDetailStartEndTimeVisibleFlag,omitempty"`
	TimeDetailSubtotalVisibleFlag                   bool                   `json:"timeDetailSubtotalVisibleFlag,omitempty"`
	TimeDetailTicketNumberVisibleFlag               bool                   `json:"timeDetailTicketNumberVisibleFlag,omitempty"`
	TimeDetailVisibleFlag                           bool                   `json:"timeDetailVisibleFlag,omitempty"`
}

type InvoiceTemplateSetup struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	CustomFlag bool                   `json:"customFlag,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Name       string                 `json:"name,omitempty"`
}

type MappedType struct {
	GlType     string `json:"glType,omitempty"`
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	RecIdField string `json:"recIdField,omitempty"`
	SortOrder  int    `json:"sortOrder,omitempty"`
	Table      string `json:"table,omitempty"`
}

type ProductTypeExemption struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	ID            int                    `json:"id,omitempty"`
	ProductType   ProductType            `json:"productType"`
	TaxableLevels []int                  `json:"taxableLevels,omitempty"`
}

type TaxCode struct {
	Info                          map[string]interface{} `json:"_info,omitempty"`
	AddAllExpenseTypes            bool                   `json:"addAllExpenseTypes,omitempty"`
	AddAllProductTypes            bool                   `json:"addAllProductTypes,omitempty"`
	AddAllWorkRoles               bool                   `json:"addAllWorkRoles,omitempty"`
	CanadaCalculateGSTFlag        bool                   `json:"canadaCalculateGSTFlag,omitempty"`
	CancelDate                    time.Time              `json:"cancelDate,omitempty"`
	Country                       Country                `json:"country,omitempty"`
	DefaultFlag                   bool                   `json:"defaultFlag,omitempty"`
	Description                   string                 `json:"description"`
	DisplayOnInvoiceFlag          bool                   `json:"displayOnInvoiceFlag,omitempty"`
	EffectiveDate                 time.Time              `json:"effectiveDate"`
	ExpenseTypeIds                []int                  `json:"expenseTypeIds,omitempty"`
	ID                            int                    `json:"id,omitempty"`
	Identifier                    string                 `json:"identifier"`
	InvoiceCaption                string                 `json:"invoiceCaption"`
	LevelFiveAgencyXref           string                 `json:"levelFiveAgencyXref,omitempty"`
	LevelFiveApplySingleUnitFlag  bool                   `json:"levelFiveApplySingleUnitFlag,omitempty"`
	LevelFiveApplySingleUnitMax   float64                `json:"levelFiveApplySingleUnitMax,omitempty"`
	LevelFiveApplySingleUnitMin   float64                `json:"levelFiveApplySingleUnitMin,omitempty"`
	LevelFiveCaption              string                 `json:"levelFiveCaption,omitempty"`
	LevelFiveExpensesFlag         bool                   `json:"levelFiveExpensesFlag,omitempty"`
	LevelFiveProductsFlag         bool                   `json:"levelFiveProductsFlag,omitempty"`
	LevelFiveRate                 float64                `json:"levelFiveRate,omitempty"`
	LevelFiveRateType             string                 `json:"levelFiveRateType,omitempty"`
	LevelFiveServicesFlag         bool                   `json:"levelFiveServicesFlag,omitempty"`
	LevelFiveTaxCodeXref          string                 `json:"levelFiveTaxCodeXref,omitempty"`
	LevelFiveTaxableMax           float64                `json:"levelFiveTaxableMax,omitempty"`
	LevelFourAgencyXref           string                 `json:"levelFourAgencyXref,omitempty"`
	LevelFourApplySingleUnitFlag  bool                   `json:"levelFourApplySingleUnitFlag,omitempty"`
	LevelFourApplySingleUnitMax   float64                `json:"levelFourApplySingleUnitMax,omitempty"`
	LevelFourApplySingleUnitMin   float64                `json:"levelFourApplySingleUnitMin,omitempty"`
	LevelFourCaption              string                 `json:"levelFourCaption,omitempty"`
	LevelFourExpensesFlag         bool                   `json:"levelFourExpensesFlag,omitempty"`
	LevelFourProductsFlag         bool                   `json:"levelFourProductsFlag,omitempty"`
	LevelFourRate                 float64                `json:"levelFourRate,omitempty"`
	LevelFourRateType             string                 `json:"levelFourRateType,omitempty"`
	LevelFourServicesFlag         bool                   `json:"levelFourServicesFlag,omitempty"`
	LevelFourTaxCodeXref          string                 `json:"levelFourTaxCodeXref,omitempty"`
	LevelFourTaxableMax           float64                `json:"levelFourTaxableMax,omitempty"`
	LevelOneAgencyXref            string                 `json:"levelOneAgencyXref,omitempty"`
	LevelOneApplySingleUnitFlag   bool                   `json:"levelOneApplySingleUnitFlag,omitempty"`
	LevelOneApplySingleUnitMax    float64                `json:"levelOneApplySingleUnitMax,omitempty"`
	LevelOneApplySingleUnitMin    float64                `json:"levelOneApplySingleUnitMin,omitempty"`
	LevelOneCaption               string                 `json:"levelOneCaption,omitempty"`
	LevelOneExpensesFlag          bool                   `json:"levelOneExpensesFlag,omitempty"`
	LevelOneProductsFlag          bool                   `json:"levelOneProductsFlag,omitempty"`
	LevelOneRate                  float64                `json:"levelOneRate,omitempty"`
	LevelOneRateType              string                 `json:"levelOneRateType,omitempty"`
	LevelOneServicesFlag          bool                   `json:"levelOneServicesFlag,omitempty"`
	LevelOneTaxCodeXref           string                 `json:"levelOneTaxCodeXref,omitempty"`
	LevelOneTaxableMax            float64                `json:"levelOneTaxableMax,omitempty"`
	LevelSixAgencyXref            string                 `json:"levelSixAgencyXref,omitempty"`
	LevelSixApplySingleUnitFlag   bool                   `json:"levelSixApplySingleUnitFlag,omitempty"`
	LevelSixApplySingleUnitMax    float64                `json:"levelSixApplySingleUnitMax,omitempty"`
	LevelSixApplySingleUnitMin    float64                `json:"levelSixApplySingleUnitMin,omitempty"`
	LevelSixCaption               string                 `json:"levelSixCaption,omitempty"`
	LevelSixExpensesFlag          bool                   `json:"levelSixExpensesFlag,omitempty"`
	LevelSixProductsFlag          bool                   `json:"levelSixProductsFlag,omitempty"`
	LevelSixRate                  float64                `json:"levelSixRate,omitempty"`
	LevelSixRateType              string                 `json:"levelSixRateType,omitempty"`
	LevelSixServicesFlag          bool                   `json:"levelSixServicesFlag,omitempty"`
	LevelSixTaxCodeXref           string                 `json:"levelSixTaxCodeXref,omitempty"`
	LevelSixTaxableMax            float64                `json:"levelSixTaxableMax,omitempty"`
	LevelThreeAgencyXref          string                 `json:"levelThreeAgencyXref,omitempty"`
	LevelThreeApplySingleUnitFlag bool                   `json:"levelThreeApplySingleUnitFlag,omitempty"`
	LevelThreeApplySingleUnitMax  float64                `json:"levelThreeApplySingleUnitMax,omitempty"`
	LevelThreeApplySingleUnitMin  float64                `json:"levelThreeApplySingleUnitMin,omitempty"`
	LevelThreeCaption             string                 `json:"levelThreeCaption,omitempty"`
	LevelThreeExpensesFlag        bool                   `json:"levelThreeExpensesFlag,omitempty"`
	LevelThreeProductsFlag        bool                   `json:"levelThreeProductsFlag,omitempty"`
	LevelThreeRate                float64                `json:"levelThreeRate,omitempty"`
	LevelThreeRateType            string                 `json:"levelThreeRateType,omitempty"`
	LevelThreeServicesFlag        bool                   `json:"levelThreeServicesFlag,omitempty"`
	LevelThreeTaxCodeXref         string                 `json:"levelThreeTaxCodeXref,omitempty"`
	LevelThreeTaxableMax          float64                `json:"levelThreeTaxableMax,omitempty"`
	LevelTwoAgencyXref            string                 `json:"levelTwoAgencyXref,omitempty"`
	LevelTwoApplySingleUnitFlag   bool                   `json:"levelTwoApplySingleUnitFlag,omitempty"`
	LevelTwoApplySingleUnitMax    float64                `json:"levelTwoApplySingleUnitMax,omitempty"`
	LevelTwoApplySingleUnitMin    float64                `json:"levelTwoApplySingleUnitMin,omitempty"`
	LevelTwoCaption               string                 `json:"levelTwoCaption,omitempty"`
	LevelTwoExpensesFlag          bool                   `json:"levelTwoExpensesFlag,omitempty"`
	LevelTwoProductsFlag          bool                   `json:"levelTwoProductsFlag,omitempty"`
	LevelTwoRate                  float64                `json:"levelTwoRate,omitempty"`
	LevelTwoRateType              string                 `json:"levelTwoRateType,omitempty"`
	LevelTwoServicesFlag          bool                   `json:"levelTwoServicesFlag,omitempty"`
	LevelTwoTaxCodeXref           string                 `json:"levelTwoTaxCodeXref,omitempty"`
	LevelTwoTaxableMax            float64                `json:"levelTwoTaxableMax,omitempty"`
	ProductTypeIds                []int                  `json:"productTypeIds,omitempty"`
	RemoveAllExpenseTypes         bool                   `json:"removeAllExpenseTypes,omitempty"`
	RemoveAllProductTypes         bool                   `json:"removeAllProductTypes,omitempty"`
	RemoveAllWorkRoles            bool                   `json:"removeAllWorkRoles,omitempty"`
	WorkRoleIds                   []int                  `json:"workRoleIds,omitempty"`
}

type TaxCodeInfo struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	CancelDate     string                 `json:"cancelDate,omitempty"`
	DefaultFlag    bool                   `json:"defaultFlag,omitempty"`
	Description    string                 `json:"description,omitempty"`
	EffectiveDate  string                 `json:"effectiveDate,omitempty"`
	ID             int                    `json:"id,omitempty"`
	Identifier     string                 `json:"identifier,omitempty"`
	LevelFiveRate  float64                `json:"levelFiveRate,omitempty"`
	LevelFourRate  float64                `json:"levelFourRate,omitempty"`
	LevelOneRate   float64                `json:"levelOneRate,omitempty"`
	LevelSixRate   float64                `json:"levelSixRate,omitempty"`
	LevelThreeRate float64                `json:"levelThreeRate,omitempty"`
	LevelTwoRate   float64                `json:"levelTwoRate,omitempty"`
}

type TaxCodeLevel struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	AgencyXref        string                 `json:"agencyXref,omitempty"`
	Caption           string                 `json:"caption,omitempty"`
	ID                int                    `json:"id,omitempty"`
	RateType          string                 `json:"rateType,omitempty"`
	SingleUnitFlag    bool                   `json:"singleUnitFlag,omitempty"`
	SingleUnitMaximum float64                `json:"singleUnitMaximum,omitempty"`
	SingleUnitMinimum float64                `json:"singleUnitMinimum,omitempty"`
	TaxCodeXref       string                 `json:"taxCodeXref,omitempty"`
	TaxExpensesFlag   bool                   `json:"taxExpensesFlag,omitempty"`
	TaxLevel          int                    `json:"taxLevel,omitempty"`
	TaxProductsFlag   bool                   `json:"taxProductsFlag,omitempty"`
	TaxRate           float64                `json:"taxRate,omitempty"`
	TaxServicesFlag   bool                   `json:"taxServicesFlag,omitempty"`
	TaxableMax        float64                `json:"taxableMax,omitempty"`
}

type TaxCodeXRef struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag   bool                   `json:"defaultFlag,omitempty"`
	Description   string                 `json:"description"`
	ID            int                    `json:"id,omitempty"`
	LevelFive     string                 `json:"levelFive,omitempty"`
	LevelFour     string                 `json:"levelFour,omitempty"`
	LevelOne      string                 `json:"levelOne,omitempty"`
	LevelSix      string                 `json:"levelSix,omitempty"`
	LevelThree    string                 `json:"levelThree,omitempty"`
	LevelTwo      string                 `json:"levelTwo,omitempty"`
	TaxCode       TaxCode                `json:"taxCode,omitempty"`
	TaxableLevels []int                  `json:"taxableLevels,omitempty"`
}

type TaxIntegration struct {
	Info                      map[string]interface{} `json:"_info,omitempty"`
	AccountNumber             string                 `json:"accountNumber,omitempty"`
	AccountingIntegrationFlag bool                   `json:"accountingIntegrationFlag,omitempty"`
	CommitTransactionsFlag    bool                   `json:"commitTransactionsFlag,omitempty"`
	CompanyCode               string                 `json:"companyCode,omitempty"`
	EnabledFlag               bool                   `json:"enabledFlag,omitempty"`
	ExpenseTaxCode            string                 `json:"expenseTaxCode,omitempty"`
	FreightTaxCode            string                 `json:"freightTaxCode,omitempty"`
	ID                        int                    `json:"id,omitempty"`
	InvoiceAmountTaxCode      string                 `json:"invoiceAmountTaxCode,omitempty"`
	LicenseKey                string                 `json:"licenseKey,omitempty"`
	ProductTaxCode            string                 `json:"productTaxCode,omitempty"`
	SalesInvoiceFlag          bool                   `json:"salesInvoiceFlag,omitempty"`
	ServiceUrl                string                 `json:"serviceUrl,omitempty"`
	TaxIntegrationType        string                 `json:"taxIntegrationType,omitempty"`
	TaxLineFlag               bool                   `json:"taxLineFlag,omitempty"`
	TimeTaxCode               string                 `json:"timeTaxCode,omitempty"`
}

type TaxIntegrationInfo struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	EnabledFlag        bool                   `json:"enabledFlag,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	TaxIntegrationType string                 `json:"taxIntegrationType,omitempty"`
}

type TaxableExpenseTypeLevel struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	TaxCodeLevel TaxCodeLevel           `json:"taxCodeLevel"`
}

type TaxableProductTypeLevel struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	TaxCodeLevel TaxCodeLevel           `json:"taxCodeLevel"`
}

type TaxableWorkRoleLevel struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	TaxCodeLevel TaxCodeLevel           `json:"taxCodeLevel"`
}

type TaxableXRefLevel struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	TaxCodeLevel TaxCodeLevel           `json:"taxCodeLevel"`
}

type UnpostedExpense struct {
	Info                   map[string]interface{} `json:"_info,omitempty"`
	AccountNumber          string                 `json:"accountNumber,omitempty"`
	Agreement              Agreement              `json:"agreement,omitempty"`
	AgreementAmountCovered float64                `json:"agreementAmountCovered,omitempty"`
	AvalaraTaxFlag         bool                   `json:"avalaraTaxFlag,omitempty"`
	BillableAmount         float64                `json:"billableAmount,omitempty"`
	ChargeCode             ChargeCode             `json:"chargeCode,omitempty"`
	ChargeDescription      string                 `json:"chargeDescription,omitempty"`
	CityTaxAmount          float64                `json:"cityTaxAmount,omitempty"`
	CityTaxFlag            bool                   `json:"cityTaxFlag,omitempty"`
	CityTaxXref            string                 `json:"cityTaxXref,omitempty"`
	Classification         string                 `json:"classification,omitempty"`
	Company                Company                `json:"company,omitempty"`
	CompositeTaxAmount     float64                `json:"compositeTaxAmount,omitempty"`
	CompositeTaxFlag       bool                   `json:"compositeTaxFlag,omitempty"`
	CompositeTaxXref       string                 `json:"compositeTaxXref,omitempty"`
	CountryTaxAmount       float64                `json:"countryTaxAmount,omitempty"`
	CountryTaxFlag         bool                   `json:"countryTaxFlag,omitempty"`
	CountryTaxXref         string                 `json:"countryTaxXref,omitempty"`
	CountyTaxAmount        float64                `json:"countyTaxAmount,omitempty"`
	CountyTaxFlag          bool                   `json:"countyTaxFlag,omitempty"`
	CountyTaxXref          string                 `json:"countyTaxXref,omitempty"`
	CreditAccount          string                 `json:"creditAccount,omitempty"`
	Currency               Currency               `json:"currency,omitempty"`
	DateClosed             string                 `json:"dateClosed,omitempty"`
	DateExpense            string                 `json:"dateExpense,omitempty"`
	DepartmentId           int                    `json:"departmentId,omitempty"`
	ExpenseDetailId        int                    `json:"expenseDetailId,omitempty"`
	ExpenseType            ExpenseType            `json:"expenseType,omitempty"`
	GlType                 string                 `json:"glType,omitempty"`
	ID                     int                    `json:"id,omitempty"`
	InPolicy               bool                   `json:"inPolicy,omitempty"`
	ItemTaxableFlag        bool                   `json:"itemTaxableFlag,omitempty"`
	LevelSixTaxAmount      float64                `json:"levelSixTaxAmount,omitempty"`
	LevelSixTaxFlag        bool                   `json:"levelSixTaxFlag,omitempty"`
	LevelSixTaxXref        string                 `json:"levelSixTaxXref,omitempty"`
	LocationId             int                    `json:"locationId,omitempty"`
	Member                 Member                 `json:"member,omitempty"`
	NonBillableAmount      float64                `json:"nonBillableAmount,omitempty"`
	PaymentMethod          PaymentMethod          `json:"paymentMethod,omitempty"`
	Project                Project                `json:"project,omitempty"`
	ProjectPhase           ProjectPhase           `json:"projectPhase,omitempty"`
	SalesTaxAmount         float64                `json:"salesTaxAmount,omitempty"`
	StateTaxAmount         float64                `json:"stateTaxAmount,omitempty"`
	StateTaxFlag           bool                   `json:"stateTaxFlag,omitempty"`
	StateTaxXref           string                 `json:"stateTaxXref,omitempty"`
	TaxCode                TaxCode                `json:"taxCode,omitempty"`
	Ticket                 Ticket                 `json:"ticket,omitempty"`
	Total                  float64                `json:"total,omitempty"`
}

type UnpostedExpenseTaxableLevel struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	ID          int                    `json:"id,omitempty"`
	TaxAmount   float64                `json:"taxAmount,omitempty"`
	TaxCodeXref string                 `json:"taxCodeXref,omitempty"`
	TaxLevel    int                    `json:"taxLevel,omitempty"`
}

type UnpostedInvoice struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	AccountNumber      string                 `json:"accountNumber,omitempty"`
	AvalaraTaxFlag     bool                   `json:"avalaraTaxFlag,omitempty"`
	BillToCompany      Company                `json:"billToCompany,omitempty"`
	BillToSite         Site                   `json:"billToSite,omitempty"`
	BillingLogId       int                    `json:"billingLogId,omitempty"`
	BillingTerms       BillingTerms           `json:"billingTerms,omitempty"`
	CityTaxAmount      float64                `json:"cityTaxAmount,omitempty"`
	CityTaxFlag        bool                   `json:"cityTaxFlag,omitempty"`
	CityTaxXref        string                 `json:"cityTaxXref,omitempty"`
	Company            Company                `json:"company,omitempty"`
	CompositeTaxAmount float64                `json:"compositeTaxAmount,omitempty"`
	CompositeTaxFlag   bool                   `json:"compositeTaxFlag,omitempty"`
	CompositeTaxXref   string                 `json:"compositeTaxXref,omitempty"`
	CountryTaxAmount   float64                `json:"countryTaxAmount,omitempty"`
	CountryTaxFlag     bool                   `json:"countryTaxFlag,omitempty"`
	CountryTaxXref     string                 `json:"countryTaxXref,omitempty"`
	CountyTaxAmount    float64                `json:"countyTaxAmount,omitempty"`
	CountyTaxFlag      bool                   `json:"countyTaxFlag,omitempty"`
	CountyTaxXref      string                 `json:"countyTaxXref,omitempty"`
	CreatedBy          string                 `json:"createdBy,omitempty"`
	Currency           Currency               `json:"currency,omitempty"`
	DateClosed         string                 `json:"dateClosed,omitempty"`
	Department         SystemDepartment       `json:"department,omitempty"`
	DepartmentId       int                    `json:"departmentId,omitempty"`
	Description        string                 `json:"description,omitempty"`
	DueDate            string                 `json:"dueDate,omitempty"`
	DueDays            string                 `json:"dueDays,omitempty"`
	HasExpenses        bool                   `json:"hasExpenses,omitempty"`
	HasProducts        bool                   `json:"hasProducts,omitempty"`
	HasTime            bool                   `json:"hasTime,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	InvoiceDate        string                 `json:"invoiceDate,omitempty"`
	InvoiceNumber      string                 `json:"invoiceNumber,omitempty"`
	InvoiceTaxableFlag bool                   `json:"invoiceTaxableFlag,omitempty"`
	InvoiceType        string                 `json:"invoiceType,omitempty"`
	ItemTaxableFlag    bool                   `json:"itemTaxableFlag,omitempty"`
	LevelSixTaxAmount  float64                `json:"levelSixTaxAmount,omitempty"`
	LevelSixTaxFlag    bool                   `json:"levelSixTaxFlag,omitempty"`
	LevelSixTaxXref    string                 `json:"levelSixTaxXref,omitempty"`
	Location           OwnerLevel             `json:"location,omitempty"`
	LocationId         int                    `json:"locationId,omitempty"`
	SalesTaxAmount     float64                `json:"salesTaxAmount,omitempty"`
	ShipToCompany      Company                `json:"shipToCompany,omitempty"`
	ShipToSite         Site                   `json:"shipToSite,omitempty"`
	StateTaxAmount     float64                `json:"stateTaxAmount,omitempty"`
	StateTaxFlag       bool                   `json:"stateTaxFlag,omitempty"`
	StateTaxXref       string                 `json:"stateTaxXref,omitempty"`
	SubTotal           float64                `json:"subTotal,omitempty"`
	TaxCode            TaxCode                `json:"taxCode,omitempty"`
	Total              float64                `json:"total,omitempty"`
}

type UnpostedInvoiceTaxableLevel struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	ID          int                    `json:"id,omitempty"`
	TaxAmount   float64                `json:"taxAmount,omitempty"`
	TaxCodeXref string                 `json:"taxCodeXref,omitempty"`
	TaxLevel    int                    `json:"taxLevel,omitempty"`
}

type UnpostedPayments struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	ARPaymentAccount  string                 `json:"aRPaymentAccount,omitempty"`
	Amount            float64                `json:"amount,omitempty"`
	AppliedBy         string                 `json:"appliedBy,omitempty"`
	ID                int                    `json:"id,omitempty"`
	Invoice           Invoice                `json:"invoice,omitempty"`
	PaymentAccount    string                 `json:"paymentAccount,omitempty"`
	PaymentDate       string                 `json:"paymentDate,omitempty"`
	PaymentSyncDate   string                 `json:"paymentSyncDate,omitempty"`
	PaymentSyncStatus string                 `json:"paymentSyncStatus,omitempty"`
	Source            string                 `json:"source,omitempty"`
	Type              string                 `json:"type,omitempty"`
	WisePayPayment    WisePayPayment         `json:"wisePayPayment,omitempty"`
}

type UnpostedProcurement struct {
	Info                     map[string]interface{} `json:"_info,omitempty"`
	AvalaraTaxFlag           bool                   `json:"avalaraTaxFlag,omitempty"`
	BillingTerms             BillingTerms           `json:"billingTerms,omitempty"`
	CityTaxAmount            float64                `json:"cityTaxAmount,omitempty"`
	CityTaxFlag              bool                   `json:"cityTaxFlag,omitempty"`
	CityTaxXref              string                 `json:"cityTaxXref,omitempty"`
	CompositeTaxAmount       float64                `json:"compositeTaxAmount,omitempty"`
	CompositeTaxFlag         bool                   `json:"compositeTaxFlag,omitempty"`
	CompositeTaxXref         string                 `json:"compositeTaxXref,omitempty"`
	CountryTaxAmount         float64                `json:"countryTaxAmount,omitempty"`
	CountryTaxFlag           bool                   `json:"countryTaxFlag,omitempty"`
	CountryTaxXref           string                 `json:"countryTaxXref,omitempty"`
	CountyTaxAmount          float64                `json:"countyTaxAmount,omitempty"`
	CountyTaxFlag            bool                   `json:"countyTaxFlag,omitempty"`
	CountyTaxXref            string                 `json:"countyTaxXref,omitempty"`
	Currency                 Currency               `json:"currency,omitempty"`
	Customer                 Company                `json:"customer,omitempty"`
	DateClosed               string                 `json:"dateClosed,omitempty"`
	DepartmentId             int                    `json:"departmentId,omitempty"`
	Description              string                 `json:"description,omitempty"`
	FreightCost              float64                `json:"freightCost,omitempty"`
	FreightTaxTotal          float64                `json:"freightTaxTotal,omitempty"`
	ID                       int                    `json:"id,omitempty"`
	ItemTaxableFlag          bool                   `json:"itemTaxableFlag,omitempty"`
	LevelSixTaxAmount        float64                `json:"levelSixTaxAmount,omitempty"`
	LevelSixTaxFlag          bool                   `json:"levelSixTaxFlag,omitempty"`
	LevelSixTaxXref          string                 `json:"levelSixTaxXref,omitempty"`
	LocationId               int                    `json:"locationId,omitempty"`
	ProcurementType          string                 `json:"procurementType,omitempty"`
	PurchaseDate             string                 `json:"purchaseDate,omitempty"`
	PurchaseOrder            PurchaseOrder          `json:"purchaseOrder,omitempty"`
	PurchaseOrderTaxableFlag bool                   `json:"purchaseOrderTaxableFlag,omitempty"`
	StateTaxAmount           float64                `json:"stateTaxAmount,omitempty"`
	StateTaxFlag             bool                   `json:"stateTaxFlag,omitempty"`
	StateTaxXref             string                 `json:"stateTaxXref,omitempty"`
	TaxCode                  TaxCode                `json:"taxCode,omitempty"`
	TaxFreightFlag           bool                   `json:"taxFreightFlag,omitempty"`
	TaxTotal                 float64                `json:"taxTotal,omitempty"`
	Total                    float64                `json:"total,omitempty"`
	TrackingNumber           string                 `json:"trackingNumber,omitempty"`
	UnpostedProductId        string                 `json:"unpostedProductId,omitempty"`
	Vendor                   Company                `json:"vendor,omitempty"`
	VendorAccountNumber      string                 `json:"vendorAccountNumber,omitempty"`
	VendorInvoiceDate        string                 `json:"vendorInvoiceDate,omitempty"`
	VendorInvoiceNumber      string                 `json:"vendorInvoiceNumber,omitempty"`
}

type UnpostedProcurementTaxableLevel struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	ID          int                    `json:"id,omitempty"`
	TaxAmount   float64                `json:"taxAmount,omitempty"`
	TaxCodeXref string                 `json:"taxCodeXref,omitempty"`
	TaxLevel    int                    `json:"taxLevel,omitempty"`
}

type WorkRoleExemption struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	ID            int                    `json:"id,omitempty"`
	TaxableLevels []int                  `json:"taxableLevels,omitempty"`
	WorkRole      WorkRole               `json:"workRole"`
}
