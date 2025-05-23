package types

import (
	"time"
)

type Activity struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	Agreement      Agreement              `json:"agreement,omitempty"`
	AssignTo       Member                 `json:"assignTo"`
	AssignedBy     Member                 `json:"assignedBy,omitempty"`
	Campaign       Campaign               `json:"campaign,omitempty"`
	Company        Company                `json:"company,omitempty"`
	Contact        Contact                `json:"contact,omitempty"`
	Currency       Currency               `json:"currency,omitempty"`
	CustomFields   []CustomFieldValue     `json:"customFields,omitempty"`
	DateEnd        time.Time              `json:"dateEnd,omitempty"`
	DateStart      time.Time              `json:"dateStart,omitempty"`
	Email          string                 `json:"email,omitempty"`
	ID             int                    `json:"id,omitempty"`
	MobileGuid     string                 `json:"mobileGuid,omitempty"`
	Name           string                 `json:"name"`
	Notes          string                 `json:"notes,omitempty"`
	NotifyFlag     bool                   `json:"notifyFlag,omitempty"`
	Opportunity    Opportunity            `json:"opportunity,omitempty"`
	PhoneNumber    string                 `json:"phoneNumber,omitempty"`
	Reminder       Reminder               `json:"reminder,omitempty"`
	ScheduleStatus ScheduleStatus         `json:"scheduleStatus,omitempty"`
	Status         ActivityStatus         `json:"status,omitempty"`
	Ticket         Ticket                 `json:"ticket,omitempty"`
	Type           ActivityType           `json:"type,omitempty"`
	Where          ServiceLocation        `json:"where,omitempty"`
}

type ActivityStatus struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag        bool                   `json:"closedFlag,omitempty"`
	DefaultFlag       bool                   `json:"defaultFlag,omitempty"`
	ID                int                    `json:"id,omitempty"`
	InactiveFlag      bool                   `json:"inactiveFlag,omitempty"`
	Name              string                 `json:"name"`
	SpawnFollowupFlag bool                   `json:"spawnFollowupFlag,omitempty"`
}

type ActivityStatusInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag   bool                   `json:"closedFlag,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type ActivityType struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	EmailFlag    bool                   `json:"emailFlag,omitempty"`
	HistoryFlag  bool                   `json:"historyFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	MemoFlag     bool                   `json:"memoFlag,omitempty"`
	Name         string                 `json:"name"`
	Points       int                    `json:"points,omitempty"`
}

type Commission struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	Agreement           Agreement              `json:"agreement,omitempty"`
	AgreementType       AgreementType          `json:"agreementType,omitempty"`
	AgreementsFlag      bool                   `json:"agreementsFlag,omitempty"`
	BillingMethod       string                 `json:"billingMethod,omitempty"`
	CommissionBasis     string                 `json:"commissionBasis,omitempty"`
	CommissionPercent   float64                `json:"commissionPercent,omitempty"`
	Company             Company                `json:"company,omitempty"`
	DateEnd             time.Time              `json:"dateEnd,omitempty"`
	DateStart           time.Time              `json:"dateStart,omitempty"`
	Department          SystemDepartment       `json:"department,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	InvoiceOption       string                 `json:"invoiceOption,omitempty"`
	Item                IvItem                 `json:"item,omitempty"`
	Location            SystemLocation         `json:"location,omitempty"`
	Member              Member                 `json:"member"`
	MyOpportunitiesFlag bool                   `json:"myOpportunitiesFlag,omitempty"`
	NumberOfMonths      int                    `json:"numberOfMonths,omitempty"`
	ProductCategory     ProductCategory        `json:"productCategory,omitempty"`
	ProductSubCategory  ProductSubCategory     `json:"productSubCategory,omitempty"`
	ProductsFlag        bool                   `json:"productsFlag,omitempty"`
	Project             Project                `json:"project,omitempty"`
	ProjectBoard        ProjectBoard           `json:"projectBoard,omitempty"`
	ProjectType         ProjectType            `json:"projectType,omitempty"`
	ServiceBoard        Board                  `json:"serviceBoard,omitempty"`
	ServiceType         ServiceType            `json:"serviceType,omitempty"`
	ServicesFlag        bool                   `json:"servicesFlag,omitempty"`
	Site                Site                   `json:"site,omitempty"`
	Territory           SystemLocation         `json:"territory,omitempty"`
	Ticket              Ticket                 `json:"ticket,omitempty"`
}

type Forecast struct {
	Info                   map[string]interface{} `json:"_info,omitempty"`
	AgreementRevenue       AgreementRevenue       `json:"agreementRevenue,omitempty"`
	BillingTerms           BillingTerms           `json:"billingTerms,omitempty"`
	Currency               Currency               `json:"currency,omitempty"`
	ExpectedProbability    int                    `json:"expectedProbability,omitempty"`
	ExpenseRevenue         ExpenseRevenue         `json:"expenseRevenue,omitempty"`
	ForecastItems          []ForecastItem         `json:"forecastItems,omitempty"`
	ForecastRevenueTotals  ForecastRevenue        `json:"forecastRevenueTotals,omitempty"`
	ForecastTotalWithTaxes float64                `json:"forecastTotalWithTaxes,omitempty"`
	ID                     int                    `json:"id,omitempty"`
	InclusiveRevenueTotals InclusiveRevenue       `json:"inclusiveRevenueTotals,omitempty"`
	LostRevenue            LostRevenue            `json:"lostRevenue,omitempty"`
	OpenRevenue            OpenRevenue            `json:"openRevenue,omitempty"`
	OtherRevenue1          OtherRevenue1          `json:"otherRevenue1,omitempty"`
	OtherRevenue2          OtherRevenue2          `json:"otherRevenue2,omitempty"`
	ProductRevenue         ProductRevenue         `json:"productRevenue,omitempty"`
	RecurringTotal         float64                `json:"recurringTotal,omitempty"`
	SalesTaxRevenue        float64                `json:"salesTaxRevenue,omitempty"`
	ServiceRevenue         ServiceRevenue         `json:"serviceRevenue,omitempty"`
	TaxCode                TaxCode                `json:"taxCode,omitempty"`
	TimeRevenue            TimeRevenue            `json:"timeRevenue,omitempty"`
	WonRevenue             WonRevenue             `json:"wonRevenue,omitempty"`
}

type ForecastItem struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	BillCycle           BillingCycle           `json:"billCycle,omitempty"`
	CatalogItem         IvItem                 `json:"catalogItem,omitempty"`
	Cost                float64                `json:"cost,omitempty"`
	CycleBasis          string                 `json:"cycleBasis,omitempty"`
	Cycles              int                    `json:"cycles,omitempty"`
	ForecastDescription string                 `json:"forecastDescription,omitempty"`
	ForecastType        string                 `json:"forecastType,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	IncludeFlag         bool                   `json:"includeFlag,omitempty"`
	LinkFlag            bool                   `json:"linkFlag,omitempty"`
	Margin              float64                `json:"margin,omitempty"`
	Opportunity         Opportunity            `json:"opportunity"`
	Percentage          int                    `json:"percentage,omitempty"`
	ProductClass        string                 `json:"productClass,omitempty"`
	ProductDescription  string                 `json:"productDescription,omitempty"`
	Quantity            float64                `json:"quantity,omitempty"`
	QuoteWerksDocName   string                 `json:"quoteWerksDocName,omitempty"`
	QuoteWerksDocNo     string                 `json:"quoteWerksDocNo,omitempty"`
	QuoteWerksQuantity  int                    `json:"quoteWerksQuantity,omitempty"`
	RecurringCost       float64                `json:"recurringCost,omitempty"`
	RecurringDateEnd    time.Time              `json:"recurringDateEnd,omitempty"`
	RecurringDateStart  time.Time              `json:"recurringDateStart,omitempty"`
	RecurringFlag       bool                   `json:"recurringFlag,omitempty"`
	RecurringRevenue    float64                `json:"recurringRevenue,omitempty"`
	Revenue             float64                `json:"revenue,omitempty"`
	SequenceNumber      float64                `json:"sequenceNumber,omitempty"`
	Status              OpportunityStatus      `json:"status"`
	SubNumber           int                    `json:"subNumber,omitempty"`
	TaxableFlag         bool                   `json:"taxableFlag,omitempty"`
}

type Opportunity struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	BillToCompany      Company                `json:"billToCompany,omitempty"`
	BillToContact      Contact                `json:"billToContact,omitempty"`
	BillToSite         Site                   `json:"billToSite,omitempty"`
	BillingTerms       BillingTerms           `json:"billingTerms,omitempty"`
	BusinessUnitId     int                    `json:"businessUnitId,omitempty"`
	Campaign           Campaign               `json:"campaign,omitempty"`
	ClosedBy           Member                 `json:"closedBy,omitempty"`
	ClosedDate         time.Time              `json:"closedDate,omitempty"`
	Company            Company                `json:"company"`
	CompanyLocationId  int                    `json:"companyLocationId,omitempty"`
	Contact            Contact                `json:"contact"`
	Currency           Currency               `json:"currency,omitempty"`
	CustomFields       []CustomFieldValue     `json:"customFields,omitempty"`
	CustomerPO         string                 `json:"customerPO,omitempty"`
	DateBecameLead     time.Time              `json:"dateBecameLead,omitempty"`
	ExpectedCloseDate  time.Time              `json:"expectedCloseDate,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	LocationId         int                    `json:"locationId,omitempty"`
	Name               string                 `json:"name"`
	Notes              string                 `json:"notes,omitempty"`
	PipelineChangeDate time.Time              `json:"pipelineChangeDate,omitempty"`
	PrimarySalesRep    Member                 `json:"primarySalesRep"`
	Priority           OpportunityPriority    `json:"priority,omitempty"`
	Probability        OpportunityProbability `json:"probability,omitempty"`
	Rating             OpportunityRating      `json:"rating,omitempty"`
	SecondarySalesRep  Member                 `json:"secondarySalesRep,omitempty"`
	ShipToCompany      Company                `json:"shipToCompany,omitempty"`
	ShipToContact      Contact                `json:"shipToContact,omitempty"`
	ShipToSite         Site                   `json:"shipToSite,omitempty"`
	Site               Site                   `json:"site,omitempty"`
	Source             string                 `json:"source,omitempty"`
	Stage              OpportunityStage       `json:"stage,omitempty"`
	Status             OpportunityStatus      `json:"status,omitempty"`
	TaxCode            TaxCode                `json:"taxCode,omitempty"`
	TechnicalContact   Contact                `json:"technicalContact,omitempty"`
	TotalSalesTax      float64                `json:"totalSalesTax,omitempty"`
	Type               OpportunityType        `json:"type,omitempty"`
}

type OpportunityContact struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	Company       Company                `json:"company,omitempty"`
	Contact       Contact                `json:"contact"`
	EmailAddress  string                 `json:"emailAddress,omitempty"`
	ID            int                    `json:"id,omitempty"`
	Notes         string                 `json:"notes,omitempty"`
	OpportunityId int                    `json:"opportunityId,omitempty"`
	PhoneNumber   string                 `json:"phoneNumber,omitempty"`
	ReferralFlag  bool                   `json:"referralFlag,omitempty"`
	Role          OpportunitySalesRole   `json:"role,omitempty"`
}

type OpportunityNote struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	EnteredBy     string                 `json:"enteredBy,omitempty"`
	Flagged       bool                   `json:"flagged,omitempty"`
	ID            int                    `json:"id,omitempty"`
	MobileGuid    string                 `json:"mobileGuid,omitempty"`
	OpportunityId int                    `json:"opportunityId,omitempty"`
	Text          string                 `json:"text"`
	Type          NoteType               `json:"type,omitempty"`
}

type OpportunityRating struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	ID        int                    `json:"id,omitempty"`
	Name      string                 `json:"name"`
	SortOrder int                    `json:"sortOrder,omitempty"`
}

type OpportunityRatingInfo struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	ID        int                    `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	SortOrder int                    `json:"sortOrder,omitempty"`
}

type OpportunityStage struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	Color          string                 `json:"color,omitempty"`
	ID             int                    `json:"id,omitempty"`
	Name           string                 `json:"name"`
	Probability    OpportunityProbability `json:"probability,omitempty"`
	SequenceNumber int                    `json:"sequenceNumber,omitempty"`
}

type OpportunityStageInfo struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	Color          string                 `json:"color,omitempty"`
	ID             int                    `json:"id,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Probability    OpportunityProbability `json:"probability,omitempty"`
	SequenceNumber int                    `json:"sequenceNumber,omitempty"`
}

type OpportunityStatus struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag   bool                   `json:"closedFlag,omitempty"`
	DateEntered  time.Time              `json:"dateEntered,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	EnteredBy    string                 `json:"enteredBy,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	LostFlag     bool                   `json:"lostFlag,omitempty"`
	Name         string                 `json:"name"`
	WonFlag      bool                   `json:"wonFlag,omitempty"`
}

type OpportunityStatusInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag   bool                   `json:"closedFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type OpportunityToAgreementConversion struct {
	AgreementId             int           `json:"agreementId,omitempty"`
	BillCycleId             int           `json:"billCycleId,omitempty"`
	BillOneTimeFlag         bool          `json:"billOneTimeFlag,omitempty"`
	BusinessUnitId          int           `json:"businessUnitId,omitempty"`
	EndDate                 string        `json:"endDate,omitempty"`
	IncludeAllDocumentsFlag bool          `json:"includeAllDocumentsFlag,omitempty"`
	IncludeAllNotesFlag     bool          `json:"includeAllNotesFlag,omitempty"`
	IncludeAllProductsFlag  bool          `json:"includeAllProductsFlag,omitempty"`
	IncludeDocumentIds      []int         `json:"includeDocumentIds,omitempty"`
	IncludeNoteIds          []int         `json:"includeNoteIds,omitempty"`
	IncludeProductIds       []int         `json:"includeProductIds,omitempty"`
	LocationId              int           `json:"locationId,omitempty"`
	Name                    string        `json:"name,omitempty"`
	NoEndingDateFlag        bool          `json:"noEndingDateFlag,omitempty"`
	StartDate               string        `json:"startDate,omitempty"`
	Type                    AgreementType `json:"type,omitempty"`
}

type OpportunityToProjectConversion struct {
	Board                   ProjectBoard  `json:"board,omitempty"`
	BusinessUnitId          int           `json:"businessUnitId,omitempty"`
	EstimatedEnd            string        `json:"estimatedEnd,omitempty"`
	EstimatedStart          string        `json:"estimatedStart,omitempty"`
	IncludeAllDocumentsFlag bool          `json:"includeAllDocumentsFlag,omitempty"`
	IncludeAllNotesFlag     bool          `json:"includeAllNotesFlag,omitempty"`
	IncludeAllProductsFlag  bool          `json:"includeAllProductsFlag,omitempty"`
	IncludeDocumentIds      []int         `json:"includeDocumentIds,omitempty"`
	IncludeNoteIds          []int         `json:"includeNoteIds,omitempty"`
	IncludeProductIds       []int         `json:"includeProductIds,omitempty"`
	LocationId              int           `json:"locationId,omitempty"`
	Manager                 Member        `json:"manager,omitempty"`
	Name                    string        `json:"name,omitempty"`
	ProjectId               int           `json:"projectId,omitempty"`
	Status                  ProjectStatus `json:"status,omitempty"`
}

type OpportunityToSalesOrderConversion struct {
	IncludeAllDocumentsFlag bool   `json:"includeAllDocumentsFlag,omitempty"`
	IncludeAllNotesFlag     bool   `json:"includeAllNotesFlag,omitempty"`
	IncludeAllProductsFlag  bool   `json:"includeAllProductsFlag,omitempty"`
	IncludeDocumentIds      []int  `json:"includeDocumentIds,omitempty"`
	IncludeNoteIds          []int  `json:"includeNoteIds,omitempty"`
	IncludeProductIds       []int  `json:"includeProductIds,omitempty"`
	Name                    string `json:"name,omitempty"`
	SalesOrderId            int    `json:"salesOrderId,omitempty"`
}

type OpportunityToServiceTicketConversion struct {
	IncludeAllDocumentsFlag bool   `json:"includeAllDocumentsFlag,omitempty"`
	IncludeAllNotesFlag     bool   `json:"includeAllNotesFlag,omitempty"`
	IncludeAllProductsFlag  bool   `json:"includeAllProductsFlag,omitempty"`
	IncludeDocumentIds      []int  `json:"includeDocumentIds,omitempty"`
	IncludeNoteIds          []int  `json:"includeNoteIds,omitempty"`
	IncludeProductIds       []int  `json:"includeProductIds,omitempty"`
	Summary                 string `json:"summary,omitempty"`
	TicketId                int    `json:"ticketId,omitempty"`
}

type OpportunityType struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	Description  string                 `json:"description"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
}

type OpportunityTypeInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	Description  string                 `json:"description,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
}

type Order struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	BillClosedFlag          bool                   `json:"billClosedFlag,omitempty"`
	BillShippedFlag         bool                   `json:"billShippedFlag,omitempty"`
	BillToCompany           Company                `json:"billToCompany,omitempty"`
	BillToContact           Contact                `json:"billToContact,omitempty"`
	BillToSite              Site                   `json:"billToSite,omitempty"`
	BillingTerms            BillingTerms           `json:"billingTerms,omitempty"`
	BottomCommentFlag       bool                   `json:"bottomCommentFlag,omitempty"`
	Company                 Company                `json:"company"`
	CompanyLocation         SystemLocation         `json:"companyLocation,omitempty"`
	ConfigIds               []int                  `json:"configIds,omitempty"`
	Contact                 Contact                `json:"contact,omitempty"`
	Currency                Currency               `json:"currency,omitempty"`
	CustomFields            []CustomFieldValue     `json:"customFields,omitempty"`
	Department              SystemDepartment       `json:"department,omitempty"`
	Description             string                 `json:"description,omitempty"`
	DocumentIds             []int                  `json:"documentIds,omitempty"`
	DueDate                 time.Time              `json:"dueDate,omitempty"`
	Email                   string                 `json:"email,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	InvoiceIds              []int                  `json:"invoiceIds,omitempty"`
	Location                SystemLocation         `json:"location,omitempty"`
	Notes                   string                 `json:"notes,omitempty"`
	Opportunity             Opportunity            `json:"opportunity,omitempty"`
	OrderDate               time.Time              `json:"orderDate,omitempty"`
	Phone                   string                 `json:"phone,omitempty"`
	PhoneExt                string                 `json:"phoneExt,omitempty"`
	PoNumber                string                 `json:"poNumber,omitempty"`
	ProductIds              []int                  `json:"productIds,omitempty"`
	RestrictDownpaymentFlag bool                   `json:"restrictDownpaymentFlag,omitempty"`
	SalesRep                Member                 `json:"salesRep"`
	ShipToCompany           Company                `json:"shipToCompany,omitempty"`
	ShipToContact           Contact                `json:"shipToContact,omitempty"`
	ShipToSite              Site                   `json:"shipToSite,omitempty"`
	Site                    Site                   `json:"site,omitempty"`
	Status                  OrderStatus            `json:"status"`
	SubTotal                float64                `json:"subTotal,omitempty"`
	TaxCode                 TaxCode                `json:"taxCode,omitempty"`
	TaxTotal                float64                `json:"taxTotal,omitempty"`
	TopCommentFlag          bool                   `json:"topCommentFlag,omitempty"`
	Total                   float64                `json:"total,omitempty"`
}

type OrderStatus struct {
	Info          map[string]interface{}    `json:"_info,omitempty"`
	ClosedFlag    bool                      `json:"closedFlag,omitempty"`
	DefaultFlag   bool                      `json:"defaultFlag,omitempty"`
	EmailTemplate *OrderStatusEmailTemplate `json:"emailTemplate,omitempty"`
	ID            int                       `json:"id,omitempty"`
	InactiveFlag  bool                      `json:"inactiveFlag,omitempty"`
	Name          string                    `json:"name"`
	SortOrder     int                       `json:"sortOrder,omitempty"`
}

type OrderStatusEmailTemplate struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	Body           string                 `json:"body"`
	CopySenderFlag bool                   `json:"copySenderFlag,omitempty"`
	EmailAddress   string                 `json:"emailAddress,omitempty"`
	FirstName      string                 `json:"firstName,omitempty"`
	ID             int                    `json:"id,omitempty"`
	LastName       string                 `json:"lastName,omitempty"`
	Status         OrderStatus            `json:"status,omitempty"`
	Subject        string                 `json:"subject"`
	UseSenderFlag  bool                   `json:"useSenderFlag,omitempty"`
}

type OrderStatusInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type OrderStatusNotification struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	Email        string                 `json:"email,omitempty"`
	ID           int                    `json:"id,omitempty"`
	Member       Member                 `json:"member,omitempty"`
	NotifyWho    NotificationRecipient  `json:"notifyWho"`
	Status       OrderStatus            `json:"status,omitempty"`
	WorkflowStep int                    `json:"workflowStep,omitempty"`
}

type Role struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name"`
}

type SalesConversion struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	ConvertedTo ConversionType         `json:"convertedTo,omitempty"`
	ParentType  string                 `json:"parentType,omitempty"`
}

type SalesOrderRecap struct {
	BillableAmount float64 `json:"billableAmount,omitempty"`
	Cost           float64 `json:"cost,omitempty"`
	ID             int     `json:"id,omitempty"`
	Margin         float64 `json:"margin,omitempty"`
	Percent        float64 `json:"percent,omitempty"`
}

type SalesOrdersLineItem struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	BillStatus          string                 `json:"billStatus,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	PurchaseOrderNumber string                 `json:"purchaseOrderNumber,omitempty"`
	Quantity            int                    `json:"quantity,omitempty"`
	QuantityCancelled   int                    `json:"quantityCancelled,omitempty"`
	SalesOrder          SalesOrder             `json:"salesOrder"`
}

type SalesProbability struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Probability int                    `json:"probability"`
}

type SalesProbabilityInfo struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Probability int                    `json:"probability,omitempty"`
}

type SalesQuota struct {
	Info             map[string]interface{} `json:"_info,omitempty"`
	AprilMargin      float64                `json:"aprilMargin,omitempty"`
	AprilRevenue     float64                `json:"aprilRevenue,omitempty"`
	AugustMargin     float64                `json:"augustMargin,omitempty"`
	AugustRevenue    float64                `json:"augustRevenue,omitempty"`
	Category         ProductCategory        `json:"category,omitempty"`
	Currency         Currency               `json:"currency,omitempty"`
	DecemberMargin   float64                `json:"decemberMargin,omitempty"`
	DecemberRevenue  float64                `json:"decemberRevenue,omitempty"`
	Department       SystemDepartment       `json:"department,omitempty"`
	FebruaryMargin   float64                `json:"februaryMargin,omitempty"`
	FebruaryRevenue  float64                `json:"februaryRevenue,omitempty"`
	ForecastYear     int                    `json:"forecastYear,omitempty"`
	ID               int                    `json:"id,omitempty"`
	JanuaryMargin    float64                `json:"januaryMargin,omitempty"`
	JanuaryRevenue   float64                `json:"januaryRevenue,omitempty"`
	JulyMargin       float64                `json:"julyMargin,omitempty"`
	JulyRevenue      float64                `json:"julyRevenue,omitempty"`
	JuneMargin       float64                `json:"juneMargin,omitempty"`
	JuneRevenue      float64                `json:"juneRevenue,omitempty"`
	Location         SystemLocation         `json:"location"`
	MarchMargin      float64                `json:"marchMargin,omitempty"`
	MarchRevenue     float64                `json:"marchRevenue,omitempty"`
	MayMargin        float64                `json:"mayMargin,omitempty"`
	MayRevenue       float64                `json:"mayRevenue,omitempty"`
	Member           Member                 `json:"member"`
	NovemberMargin   float64                `json:"novemberMargin,omitempty"`
	NovemberRevenue  float64                `json:"novemberRevenue,omitempty"`
	OctoberMargin    float64                `json:"octoberMargin,omitempty"`
	OctoberRevenue   float64                `json:"octoberRevenue,omitempty"`
	SeptemberMargin  float64                `json:"septemberMargin,omitempty"`
	SeptemberRevenue float64                `json:"septemberRevenue,omitempty"`
	SubCategory      ProductSubCategory     `json:"subCategory,omitempty"`
}

type SalesTeam struct {
	Info                 map[string]interface{} `json:"_info,omitempty"`
	ID                   int                    `json:"id,omitempty"`
	InactiveFlag         bool                   `json:"inactiveFlag,omitempty"`
	SalesTeamDescription string                 `json:"salesTeamDescription"`
	SalesTeamIdentifier  string                 `json:"salesTeamIdentifier"`
	SalesTeamLocation    SystemLocation         `json:"salesTeamLocation"`
}

type SalesTeamMember struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	AllowAccessFlag bool                   `json:"allowAccessFlag,omitempty"`
	Department      SystemDepartment       `json:"department,omitempty"`
	ID              int                    `json:"id,omitempty"`
	Location        SystemLocation         `json:"location,omitempty"`
	Member          Member                 `json:"member"`
}

type Team struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	CommissionPercent int                    `json:"commissionPercent,omitempty"`
	ID                int                    `json:"id,omitempty"`
	Member            Member                 `json:"member,omitempty"`
	OpportunityId     int                    `json:"opportunityId,omitempty"`
	ReferralFlag      bool                   `json:"referralFlag,omitempty"`
	ResponsibleFlag   bool                   `json:"responsibleFlag,omitempty"`
	SalesTeam         SalesTeam              `json:"salesTeam,omitempty"`
	Type              string                 `json:"type,omitempty"`
}
