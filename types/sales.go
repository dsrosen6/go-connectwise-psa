package types

import (
	"time"
)

type Activity struct {
	Info interface{} `json:"_info,omitempty"`
	Agreement AgreementReference `json:"agreement,omitempty"`
	AssignTo MemberReference `json:"assignTo"`
	AssignedBy MemberReference `json:"assignedBy,omitempty"`
	Campaign CampaignReference `json:"campaign,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	Contact ContactReference `json:"contact,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
	DateEnd time.Time `json:"dateEnd,omitempty"`
	DateStart time.Time `json:"dateStart,omitempty"`
	Email string `json:"email,omitempty"`
	ID int `json:"id,omitempty"`
	MobileGuid string `json:"mobileGuid,omitempty"`
	Name string `json:"name"`
	Notes string `json:"notes,omitempty"`
	NotifyFlag bool `json:"notifyFlag,omitempty"`
	Opportunity OpportunityReference `json:"opportunity,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Reminder ReminderReference `json:"reminder,omitempty"`
	ScheduleStatus ScheduleStatusReference `json:"scheduleStatus,omitempty"`
	Status ActivityStatusReference `json:"status,omitempty"`
	Ticket TicketReference `json:"ticket,omitempty"`
	Type ActivityTypeReference `json:"type,omitempty"`
	Where ServiceLocationReference `json:"where,omitempty"`
}

type ActivityStatus struct {
	Info interface{} `json:"_info,omitempty"`
	ClosedFlag bool `json:"closedFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name"`
	SpawnFollowupFlag bool `json:"spawnFollowupFlag,omitempty"`
}

type ActivityStatusInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ClosedFlag bool `json:"closedFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type ActivityType struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	EmailFlag bool `json:"emailFlag,omitempty"`
	HistoryFlag bool `json:"historyFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	MemoFlag bool `json:"memoFlag,omitempty"`
	Name string `json:"name"`
	Points int `json:"points,omitempty"`
}

type Commission struct {
	Info interface{} `json:"_info,omitempty"`
	Agreement AgreementReference `json:"agreement,omitempty"`
	AgreementType AgreementTypeReference `json:"agreementType,omitempty"`
	AgreementsFlag bool `json:"agreementsFlag,omitempty"`
	BillingMethod string `json:"billingMethod,omitempty"`
	CommissionBasis string `json:"commissionBasis,omitempty"`
	CommissionPercent float64 `json:"commissionPercent,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	DateEnd time.Time `json:"dateEnd,omitempty"`
	DateStart time.Time `json:"dateStart,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	InvoiceOption string `json:"invoiceOption,omitempty"`
	Item IvItemReference `json:"item,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Member MemberReference `json:"member"`
	MyOpportunitiesFlag bool `json:"myOpportunitiesFlag,omitempty"`
	NumberOfMonths int `json:"numberOfMonths,omitempty"`
	ProductCategory ProductCategoryReference `json:"productCategory,omitempty"`
	ProductSubCategory ProductSubCategoryReference `json:"productSubCategory,omitempty"`
	ProductsFlag bool `json:"productsFlag,omitempty"`
	Project ProjectReference `json:"project,omitempty"`
	ProjectBoard ProjectBoardReference `json:"projectBoard,omitempty"`
	ProjectType ProjectTypeReference `json:"projectType,omitempty"`
	ServiceBoard BoardReference `json:"serviceBoard,omitempty"`
	ServiceType ServiceTypeReference `json:"serviceType,omitempty"`
	ServicesFlag bool `json:"servicesFlag,omitempty"`
	Site SiteReference `json:"site,omitempty"`
	Territory SystemLocationReference `json:"territory,omitempty"`
	Ticket TicketReference `json:"ticket,omitempty"`
}

type Forecast struct {
	Info interface{} `json:"_info,omitempty"`
	AgreementRevenue AgreementRevenueReference `json:"agreementRevenue,omitempty"`
	BillingTerms BillingTermsReference `json:"billingTerms,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	ExpectedProbability int `json:"expectedProbability,omitempty"`
	ExpenseRevenue ExpenseRevenueReference `json:"expenseRevenue,omitempty"`
	ForecastItems []ForecastItem `json:"forecastItems,omitempty"`
	ForecastRevenueTotals ForecastRevenueReference `json:"forecastRevenueTotals,omitempty"`
	ForecastTotalWithTaxes float64 `json:"forecastTotalWithTaxes,omitempty"`
	ID int `json:"id,omitempty"`
	InclusiveRevenueTotals InclusiveRevenueReference `json:"inclusiveRevenueTotals,omitempty"`
	LostRevenue LostRevenueReference `json:"lostRevenue,omitempty"`
	OpenRevenue OpenRevenueReference `json:"openRevenue,omitempty"`
	OtherRevenue1 Other1RevenueReference `json:"otherRevenue1,omitempty"`
	OtherRevenue2 Other2RevenueReference `json:"otherRevenue2,omitempty"`
	ProductRevenue ProductRevenueReference `json:"productRevenue,omitempty"`
	RecurringTotal float64 `json:"recurringTotal,omitempty"`
	SalesTaxRevenue float64 `json:"salesTaxRevenue,omitempty"`
	ServiceRevenue ServiceRevenueReference `json:"serviceRevenue,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	TimeRevenue TimeRevenueReference `json:"timeRevenue,omitempty"`
	WonRevenue WonRevenueReference `json:"wonRevenue,omitempty"`
}

type ForecastItem struct {
	Info interface{} `json:"_info,omitempty"`
	BillCycle BillingCycleReference `json:"billCycle,omitempty"`
	CatalogItem IvItemReference `json:"catalogItem,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	CycleBasis string `json:"cycleBasis,omitempty"`
	Cycles int `json:"cycles,omitempty"`
	ForecastDescription string `json:"forecastDescription,omitempty"`
	ForecastType string `json:"forecastType,omitempty"`
	ID int `json:"id,omitempty"`
	IncludeFlag bool `json:"includeFlag,omitempty"`
	LinkFlag bool `json:"linkFlag,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Opportunity OpportunityReference `json:"opportunity"`
	Percentage int `json:"percentage,omitempty"`
	ProductClass string `json:"productClass,omitempty"`
	ProductDescription string `json:"productDescription,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	QuoteWerksDocName string `json:"quoteWerksDocName,omitempty"`
	QuoteWerksDocNo string `json:"quoteWerksDocNo,omitempty"`
	QuoteWerksQuantity int `json:"quoteWerksQuantity,omitempty"`
	RecurringCost float64 `json:"recurringCost,omitempty"`
	RecurringDateEnd time.Time `json:"recurringDateEnd,omitempty"`
	RecurringDateStart time.Time `json:"recurringDateStart,omitempty"`
	RecurringFlag bool `json:"recurringFlag,omitempty"`
	RecurringRevenue float64 `json:"recurringRevenue,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
	SequenceNumber float64 `json:"sequenceNumber,omitempty"`
	Status OpportunityStatusReference `json:"status"`
	SubNumber int `json:"subNumber,omitempty"`
	TaxableFlag bool `json:"taxableFlag,omitempty"`
}

type Opportunity struct {
	Info interface{} `json:"_info,omitempty"`
	BillToCompany CompanyReference `json:"billToCompany,omitempty"`
	BillToContact ContactReference `json:"billToContact,omitempty"`
	BillToSite SiteReference `json:"billToSite,omitempty"`
	BillingTerms BillingTermsReference `json:"billingTerms,omitempty"`
	BusinessUnitId int `json:"businessUnitId,omitempty"`
	Campaign CampaignReference `json:"campaign,omitempty"`
	ClosedBy MemberReference `json:"closedBy,omitempty"`
	ClosedDate time.Time `json:"closedDate,omitempty"`
	Company CompanyReference `json:"company"`
	CompanyLocationId int `json:"companyLocationId,omitempty"`
	Contact ContactReference `json:"contact"`
	Currency CurrencyReference `json:"currency,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
	CustomerPO string `json:"customerPO,omitempty"`
	DateBecameLead time.Time `json:"dateBecameLead,omitempty"`
	ExpectedCloseDate time.Time `json:"expectedCloseDate,omitempty"`
	ID int `json:"id,omitempty"`
	LocationId int `json:"locationId,omitempty"`
	Name string `json:"name"`
	Notes string `json:"notes,omitempty"`
	PipelineChangeDate time.Time `json:"pipelineChangeDate,omitempty"`
	PrimarySalesRep MemberReference `json:"primarySalesRep"`
	Priority OpportunityPriorityReference `json:"priority,omitempty"`
	Probability OpportunityProbabilityReference `json:"probability,omitempty"`
	Rating OpportunityRatingReference `json:"rating,omitempty"`
	SecondarySalesRep MemberReference `json:"secondarySalesRep,omitempty"`
	ShipToCompany CompanyReference `json:"shipToCompany,omitempty"`
	ShipToContact ContactReference `json:"shipToContact,omitempty"`
	ShipToSite SiteReference `json:"shipToSite,omitempty"`
	Site SiteReference `json:"site,omitempty"`
	Source string `json:"source,omitempty"`
	Stage OpportunityStageReference `json:"stage,omitempty"`
	Status OpportunityStatusReference `json:"status,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	TechnicalContact ContactReference `json:"technicalContact,omitempty"`
	TotalSalesTax float64 `json:"totalSalesTax,omitempty"`
	Type OpportunityTypeReference `json:"type,omitempty"`
}

type OpportunityContact struct {
	Info interface{} `json:"_info,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	Contact ContactReference `json:"contact"`
	EmailAddress string `json:"emailAddress,omitempty"`
	ID int `json:"id,omitempty"`
	Notes string `json:"notes,omitempty"`
	OpportunityId int `json:"opportunityId,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	ReferralFlag bool `json:"referralFlag,omitempty"`
	Role OpportunitySalesRoleReference `json:"role,omitempty"`
}

type OpportunityNote struct {
	Info interface{} `json:"_info,omitempty"`
	EnteredBy string `json:"enteredBy,omitempty"`
	Flagged bool `json:"flagged,omitempty"`
	ID int `json:"id,omitempty"`
	MobileGuid string `json:"mobileGuid,omitempty"`
	OpportunityId int `json:"opportunityId,omitempty"`
	Text string `json:"text"`
	Type NoteTypeReference `json:"type,omitempty"`
}

type OpportunityRating struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type OpportunityRatingInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type OpportunityStage struct {
	Info interface{} `json:"_info,omitempty"`
	Color string `json:"color,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Probability OpportunityProbabilityReference `json:"probability,omitempty"`
	SequenceNumber int `json:"sequenceNumber,omitempty"`
}

type OpportunityStageInfo struct {
	Info interface{} `json:"_info,omitempty"`
	Color string `json:"color,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Probability OpportunityProbabilityReference `json:"probability,omitempty"`
	SequenceNumber int `json:"sequenceNumber,omitempty"`
}

type OpportunityStatus struct {
	Info interface{} `json:"_info,omitempty"`
	ClosedFlag bool `json:"closedFlag,omitempty"`
	DateEntered time.Time `json:"dateEntered,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	EnteredBy string `json:"enteredBy,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	LostFlag bool `json:"lostFlag,omitempty"`
	Name string `json:"name"`
	WonFlag bool `json:"wonFlag,omitempty"`
}

type OpportunityStatusInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ClosedFlag bool `json:"closedFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type OpportunityToAgreementConversion struct {
	AgreementId int `json:"agreementId,omitempty"`
	BillCycleId int `json:"billCycleId,omitempty"`
	BillOneTimeFlag bool `json:"billOneTimeFlag,omitempty"`
	BusinessUnitId int `json:"businessUnitId,omitempty"`
	EndDate string `json:"endDate,omitempty"`
	IncludeAllDocumentsFlag bool `json:"includeAllDocumentsFlag,omitempty"`
	IncludeAllNotesFlag bool `json:"includeAllNotesFlag,omitempty"`
	IncludeAllProductsFlag bool `json:"includeAllProductsFlag,omitempty"`
	IncludeDocumentIds []int `json:"includeDocumentIds,omitempty"`
	IncludeNoteIds []int `json:"includeNoteIds,omitempty"`
	IncludeProductIds []int `json:"includeProductIds,omitempty"`
	LocationId int `json:"locationId,omitempty"`
	Name string `json:"name,omitempty"`
	NoEndingDateFlag bool `json:"noEndingDateFlag,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	Type AgreementTypeReference `json:"type,omitempty"`
}

type OpportunityToProjectConversion struct {
	Board ProjectBoardReference `json:"board,omitempty"`
	BusinessUnitId int `json:"businessUnitId,omitempty"`
	EstimatedEnd string `json:"estimatedEnd,omitempty"`
	EstimatedStart string `json:"estimatedStart,omitempty"`
	IncludeAllDocumentsFlag bool `json:"includeAllDocumentsFlag,omitempty"`
	IncludeAllNotesFlag bool `json:"includeAllNotesFlag,omitempty"`
	IncludeAllProductsFlag bool `json:"includeAllProductsFlag,omitempty"`
	IncludeDocumentIds []int `json:"includeDocumentIds,omitempty"`
	IncludeNoteIds []int `json:"includeNoteIds,omitempty"`
	IncludeProductIds []int `json:"includeProductIds,omitempty"`
	LocationId int `json:"locationId,omitempty"`
	Manager MemberReference `json:"manager,omitempty"`
	Name string `json:"name,omitempty"`
	ProjectId int `json:"projectId,omitempty"`
	Status ProjectStatusReference `json:"status,omitempty"`
}

type OpportunityToSalesOrderConversion struct {
	IncludeAllDocumentsFlag bool `json:"includeAllDocumentsFlag,omitempty"`
	IncludeAllNotesFlag bool `json:"includeAllNotesFlag,omitempty"`
	IncludeAllProductsFlag bool `json:"includeAllProductsFlag,omitempty"`
	IncludeDocumentIds []int `json:"includeDocumentIds,omitempty"`
	IncludeNoteIds []int `json:"includeNoteIds,omitempty"`
	IncludeProductIds []int `json:"includeProductIds,omitempty"`
	Name string `json:"name,omitempty"`
	SalesOrderId int `json:"salesOrderId,omitempty"`
}

type OpportunityToServiceTicketConversion struct {
	IncludeAllDocumentsFlag bool `json:"includeAllDocumentsFlag,omitempty"`
	IncludeAllNotesFlag bool `json:"includeAllNotesFlag,omitempty"`
	IncludeAllProductsFlag bool `json:"includeAllProductsFlag,omitempty"`
	IncludeDocumentIds []int `json:"includeDocumentIds,omitempty"`
	IncludeNoteIds []int `json:"includeNoteIds,omitempty"`
	IncludeProductIds []int `json:"includeProductIds,omitempty"`
	Summary string `json:"summary,omitempty"`
	TicketId int `json:"ticketId,omitempty"`
}

type OpportunityType struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
}

type OpportunityTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
}

type Order struct {
	Info interface{} `json:"_info,omitempty"`
	BillClosedFlag bool `json:"billClosedFlag,omitempty"`
	BillShippedFlag bool `json:"billShippedFlag,omitempty"`
	BillToCompany CompanyReference `json:"billToCompany,omitempty"`
	BillToContact ContactReference `json:"billToContact,omitempty"`
	BillToSite SiteReference `json:"billToSite,omitempty"`
	BillingTerms BillingTermsReference `json:"billingTerms,omitempty"`
	BottomCommentFlag bool `json:"bottomCommentFlag,omitempty"`
	Company CompanyReference `json:"company"`
	CompanyLocation SystemLocationReference `json:"companyLocation,omitempty"`
	ConfigIds []int `json:"configIds,omitempty"`
	Contact ContactReference `json:"contact,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentIds []int `json:"documentIds,omitempty"`
	DueDate time.Time `json:"dueDate,omitempty"`
	Email string `json:"email,omitempty"`
	ID int `json:"id,omitempty"`
	InvoiceIds []int `json:"invoiceIds,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Notes string `json:"notes,omitempty"`
	Opportunity OpportunityReference `json:"opportunity,omitempty"`
	OrderDate time.Time `json:"orderDate,omitempty"`
	Phone string `json:"phone,omitempty"`
	PhoneExt string `json:"phoneExt,omitempty"`
	PoNumber string `json:"poNumber,omitempty"`
	ProductIds []int `json:"productIds,omitempty"`
	RestrictDownpaymentFlag bool `json:"restrictDownpaymentFlag,omitempty"`
	SalesRep MemberReference `json:"salesRep"`
	ShipToCompany CompanyReference `json:"shipToCompany,omitempty"`
	ShipToContact ContactReference `json:"shipToContact,omitempty"`
	ShipToSite SiteReference `json:"shipToSite,omitempty"`
	Site SiteReference `json:"site,omitempty"`
	Status OrderStatusReference `json:"status"`
	SubTotal float64 `json:"subTotal,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	TaxTotal float64 `json:"taxTotal,omitempty"`
	TopCommentFlag bool `json:"topCommentFlag,omitempty"`
	Total float64 `json:"total,omitempty"`
}

type OrderStatus struct {
	Info interface{} `json:"_info,omitempty"`
	ClosedFlag bool `json:"closedFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	EmailTemplate OrderStatusEmailTemplateReference `json:"emailTemplate,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type OrderStatusEmailTemplate struct {
	Info interface{} `json:"_info,omitempty"`
	Body string `json:"body"`
	CopySenderFlag bool `json:"copySenderFlag,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	ID int `json:"id,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Status OrderStatusReference `json:"status,omitempty"`
	Subject string `json:"subject"`
	UseSenderFlag bool `json:"useSenderFlag,omitempty"`
}

type OrderStatusInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type OrderStatusNotification struct {
	Info interface{} `json:"_info,omitempty"`
	Email string `json:"email,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	NotifyWho NotificationRecipientReference `json:"notifyWho"`
	Status OrderStatusReference `json:"status,omitempty"`
	WorkflowStep int `json:"workflowStep,omitempty"`
}

type Role struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
}

type SalesConversion struct {
	Info interface{} `json:"_info,omitempty"`
	ConvertedTo ConversionTypeReference `json:"convertedTo,omitempty"`
	ParentType string `json:"parentType,omitempty"`
}

type SalesOrderRecap struct {
	BillableAmount float64 `json:"billableAmount,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percent float64 `json:"percent,omitempty"`
}

type SalesOrdersLineItem struct {
	Info interface{} `json:"_info,omitempty"`
	BillStatus string `json:"billStatus,omitempty"`
	ID int `json:"id,omitempty"`
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`
	Quantity int `json:"quantity,omitempty"`
	QuantityCancelled int `json:"quantityCancelled,omitempty"`
	SalesOrder SalesOrderReference `json:"salesOrder"`
}

type SalesProbability struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Probability int `json:"probability"`
}

type SalesProbabilityInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Probability int `json:"probability,omitempty"`
}

type SalesQuota struct {
	Info interface{} `json:"_info,omitempty"`
	AprilMargin float64 `json:"aprilMargin,omitempty"`
	AprilRevenue float64 `json:"aprilRevenue,omitempty"`
	AugustMargin float64 `json:"augustMargin,omitempty"`
	AugustRevenue float64 `json:"augustRevenue,omitempty"`
	Category ProductCategoryReference `json:"category,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	DecemberMargin float64 `json:"decemberMargin,omitempty"`
	DecemberRevenue float64 `json:"decemberRevenue,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	FebruaryMargin float64 `json:"februaryMargin,omitempty"`
	FebruaryRevenue float64 `json:"februaryRevenue,omitempty"`
	ForecastYear int `json:"forecastYear,omitempty"`
	ID int `json:"id,omitempty"`
	JanuaryMargin float64 `json:"januaryMargin,omitempty"`
	JanuaryRevenue float64 `json:"januaryRevenue,omitempty"`
	JulyMargin float64 `json:"julyMargin,omitempty"`
	JulyRevenue float64 `json:"julyRevenue,omitempty"`
	JuneMargin float64 `json:"juneMargin,omitempty"`
	JuneRevenue float64 `json:"juneRevenue,omitempty"`
	Location SystemLocationReference `json:"location"`
	MarchMargin float64 `json:"marchMargin,omitempty"`
	MarchRevenue float64 `json:"marchRevenue,omitempty"`
	MayMargin float64 `json:"mayMargin,omitempty"`
	MayRevenue float64 `json:"mayRevenue,omitempty"`
	Member MemberReference `json:"member"`
	NovemberMargin float64 `json:"novemberMargin,omitempty"`
	NovemberRevenue float64 `json:"novemberRevenue,omitempty"`
	OctoberMargin float64 `json:"octoberMargin,omitempty"`
	OctoberRevenue float64 `json:"octoberRevenue,omitempty"`
	SeptemberMargin float64 `json:"septemberMargin,omitempty"`
	SeptemberRevenue float64 `json:"septemberRevenue,omitempty"`
	SubCategory ProductSubCategoryReference `json:"subCategory,omitempty"`
}

type SalesTeam struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	SalesTeamDescription string `json:"salesTeamDescription"`
	SalesTeamIdentifier string `json:"salesTeamIdentifier"`
	SalesTeamLocation SystemLocationReference `json:"salesTeamLocation"`
}

type SalesTeamMember struct {
	Info interface{} `json:"_info,omitempty"`
	AllowAccessFlag bool `json:"allowAccessFlag,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Member MemberReference `json:"member"`
}

type Team struct {
	Info interface{} `json:"_info,omitempty"`
	CommissionPercent int `json:"commissionPercent,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	OpportunityId int `json:"opportunityId,omitempty"`
	ReferralFlag bool `json:"referralFlag,omitempty"`
	ResponsibleFlag bool `json:"responsibleFlag,omitempty"`
	SalesTeam SalesTeamReference `json:"salesTeam,omitempty"`
	Type string `json:"type,omitempty"`
}

