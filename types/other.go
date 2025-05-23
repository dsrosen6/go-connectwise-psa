package types

type Adjustment struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type AgreementApplicationAviablePer struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
}

type AgreementApplicationBillingCycle struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
}

type AgreementApplicationLimit struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
}

type AgreementApplicationUnit struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
}

type AgreementBillingInfo struct {
	AgreementAmount float64 `json:"agreementAmount,omitempty"`
	AgreementName   string  `json:"agreementName,omitempty"`
	AgreementRecId  int     `json:"agreementRecId,omitempty"`
	AgreementType   string  `json:"agreementType,omitempty"`
	ParentRecId     int     `json:"parentRecId,omitempty"`
}

type AgreementRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type ApiRequest struct {
	Entity                  IRestIdentifiedItem    `json:"entity,omitempty"`
	ExternalId              string                 `json:"externalId,omitempty"`
	Fields                  string                 `json:"fields,omitempty"`
	Filters                 FilterValues           `json:"filters,omitempty"`
	GrandParentId           int                    `json:"grandParentId,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	MemberContext           string                 `json:"memberContext,omitempty"`
	MiscProperties          map[string]interface{} `json:"miscProperties,omitempty"`
	Page                    PageValues             `json:"page,omitempty"`
	ParentId                int                    `json:"parentId,omitempty"`
	UpdateOnlyCesProperties bool                   `json:"updateOnlyCesProperties,omitempty"`
}

type AutomateScript struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type Batch struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type BillingDelivery struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type BillingTerms struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type BillingUnit struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type BundleRequest struct {
	ApiRequest     ApiRequest `json:"apiRequest,omitempty"`
	ResourceType   string     `json:"resourceType,omitempty"`
	SequenceNumber int        `json:"sequenceNumber,omitempty"`
	Version        string     `json:"version,omitempty"`
}

type BundleResult struct {
	Count          int                   `json:"count,omitempty"`
	Entities       []IRestIdentifiedItem `json:"entities,omitempty"`
	Error          ErrorResponseMessage  `json:"error,omitempty"`
	ResourceType   string                `json:"resourceType,omitempty"`
	SequenceNumber int                   `json:"sequenceNumber,omitempty"`
	StatusCode     int                   `json:"statusCode,omitempty"`
	Success        bool                  `json:"success,omitempty"`
	Version        string                `json:"version,omitempty"`
}

type CalendarSetup struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
}

type CampaignSubType struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ConfigurationQuestion struct {
	Answer           map[string]interface{} `json:"answer,omitempty"`
	AnswerId         int                    `json:"answerId,omitempty"`
	FieldType        string                 `json:"fieldType,omitempty"`
	NumberOfDecimals int                    `json:"numberOfDecimals,omitempty"`
	Question         string                 `json:"question,omitempty"`
	QuestionId       int                    `json:"questionId,omitempty"`
	RequiredFlag     bool                   `json:"requiredFlag,omitempty"`
	SequenceNumber   float64                `json:"sequenceNumber,omitempty"`
}

type ContactCommunicationItem struct {
	CommunicationType string            `json:"communicationType,omitempty"`
	DefaultFlag       bool              `json:"defaultFlag,omitempty"`
	Domain            string            `json:"domain,omitempty"`
	Extension         string            `json:"extension,omitempty"`
	ID                int               `json:"id,omitempty"`
	Type              CommunicationType `json:"type,omitempty"`
	Value             string            `json:"value,omitempty"`
}

type ConversionType struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type Currency struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	CurrencyCode            string                 `json:"currencyCode,omitempty"`
	CurrencyIdentifier      string                 `json:"currencyIdentifier,omitempty"`
	DecimalSeparator        string                 `json:"decimalSeparator,omitempty"`
	DisplayIdFlag           bool                   `json:"displayIdFlag,omitempty"`
	DisplaySymbolFlag       bool                   `json:"displaySymbolFlag,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	Name                    string                 `json:"name,omitempty"`
	NegativeParenthesesFlag bool                   `json:"negativeParenthesesFlag,omitempty"`
	NumberOfDecimals        int                    `json:"numberOfDecimals,omitempty"`
	RightAlign              bool                   `json:"rightAlign,omitempty"`
	Symbol                  string                 `json:"symbol,omitempty"`
	ThousandsSeparator      string                 `json:"thousandsSeparator,omitempty"`
}

type CustomFieldValue struct {
	Caption          string                 `json:"caption,omitempty"`
	ConnectWiseId    string                 `json:"connectWiseId,omitempty"`
	EntryMethod      string                 `json:"entryMethod,omitempty"`
	ID               int                    `json:"id,omitempty"`
	NumberOfDecimals int                    `json:"numberOfDecimals,omitempty"`
	Type             string                 `json:"type,omitempty"`
	Value            map[string]interface{} `json:"value,omitempty"`
}

type EmailConnectorParsingType struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type EmailConnectorParsingVariable struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier,omitempty"`
	Name       string                 `json:"name,omitempty"`
}

type EmailTemplate struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ErrorResponseMessage struct {
	Code    string            `json:"code,omitempty"`
	Errors  []ValidationError `json:"errors,omitempty"`
	Message string            `json:"message,omitempty"`
}

type ExistingTenant struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ExpenseDetail struct {
	Info   map[string]interface{} `json:"_info,omitempty"`
	Amount float64                `json:"amount,omitempty"`
	ID     int                    `json:"id,omitempty"`
}

type ExpenseRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type ExpenseTax struct {
	Amount float64        `json:"amount,omitempty"`
	ID     int            `json:"id,omitempty"`
	Type   ExpenseTaxType `json:"type,omitempty"`
}

type ExpenseTaxType struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ForecastRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type GLExportAdjustmentTransaction struct {
	AccountNumber         string                                `json:"accountNumber,omitempty"`
	AdjustmentDescription string                                `json:"adjustmentDescription,omitempty"`
	AdjustmentDetail      []GLExportAdjustmentTransactionDetail `json:"adjustmentDetail,omitempty"`
	DocumentDate          string                                `json:"documentDate,omitempty"`
	DocumentType          string                                `json:"documentType,omitempty"`
	GlClass               string                                `json:"glClass,omitempty"`
	GlTypeID              string                                `json:"glTypeID,omitempty"`
	ID                    string                                `json:"id,omitempty"`
	Memo                  string                                `json:"memo,omitempty"`
}

type GLExportAdjustmentTransactionDetail struct {
	AccountNumber          string  `json:"accountNumber,omitempty"`
	Cost                   float64 `json:"cost,omitempty"`
	CostAccountNumber      string  `json:"costAccountNumber,omitempty"`
	Description            string  `json:"description,omitempty"`
	GlClass                string  `json:"glClass,omitempty"`
	InventoryAccountNumber string  `json:"inventoryAccountNumber,omitempty"`
	Item                   IvItem  `json:"item,omitempty"`
	Memo                   string  `json:"memo,omitempty"`
	ProductAccountNumber   string  `json:"productAccountNumber,omitempty"`
	Quantity               int     `json:"quantity,omitempty"`
	Total                  float64 `json:"total,omitempty"`
}

type GLExportCustomer struct {
	AccountNumber          string                     `json:"accountNumber,omitempty"`
	BillingTerms           BillingTerms               `json:"billingTerms,omitempty"`
	BillingTermsXref       string                     `json:"billingTermsXref,omitempty"`
	CityTaxAgencyXref      string                     `json:"cityTaxAgencyXref,omitempty"`
	CityTaxRate            float64                    `json:"cityTaxRate,omitempty"`
	CityTaxXref            string                     `json:"cityTaxXref,omitempty"`
	Company                Company                    `json:"company,omitempty"`
	CompanyType            CompanyType                `json:"companyType,omitempty"`
	CompositeTaxAgencyXref string                     `json:"compositeTaxAgencyXref,omitempty"`
	CompositeTaxRate       float64                    `json:"compositeTaxRate,omitempty"`
	CompositeTaxXref       string                     `json:"compositeTaxXref,omitempty"`
	Contact                Contact                    `json:"contact,omitempty"`
	CountryTaxAgencyXref   string                     `json:"countryTaxAgencyXref,omitempty"`
	CountryTaxRate         float64                    `json:"countryTaxRate,omitempty"`
	CountryTaxXref         string                     `json:"countryTaxXref,omitempty"`
	CountyTaxAgencyXref    string                     `json:"countyTaxAgencyXref,omitempty"`
	CountyTaxRate          float64                    `json:"countyTaxRate,omitempty"`
	CountyTaxXref          string                     `json:"countyTaxXref,omitempty"`
	Currency               Currency                   `json:"currency,omitempty"`
	DueDays                int                        `json:"dueDays,omitempty"`
	Site                   Site                       `json:"site,omitempty"`
	StateTaxAgencyXref     string                     `json:"stateTaxAgencyXref,omitempty"`
	StateTaxRate           float64                    `json:"stateTaxRate,omitempty"`
	StateTaxXref           string                     `json:"stateTaxXref,omitempty"`
	TaxAgencyXref          string                     `json:"taxAgencyXref,omitempty"`
	TaxCode                TaxCode                    `json:"taxCode,omitempty"`
	TaxGroupRate           float64                    `json:"taxGroupRate,omitempty"`
	TaxLevels              []GLExportCustomerTaxLevel `json:"taxLevels,omitempty"`
	Taxable                bool                       `json:"taxable,omitempty"`
}

type GLExportCustomerTaxLevel struct {
	AgencyXref  string  `json:"agencyXref,omitempty"`
	TaxCodeXref string  `json:"taxCodeXref,omitempty"`
	TaxLevel    int     `json:"taxLevel,omitempty"`
	TaxRate     float64 `json:"taxRate,omitempty"`
}

type GLExportExpense struct {
	AccountNumber        string                `json:"accountNumber,omitempty"`
	ApAccountNumber      string                `json:"apAccountNumber,omitempty"`
	ApClass              string                `json:"apClass,omitempty"`
	Company              Company               `json:"company,omitempty"`
	CompanyAccountNumber string                `json:"companyAccountNumber,omitempty"`
	Currency             Currency              `json:"currency,omitempty"`
	Description          string                `json:"description,omitempty"`
	DocumentDate         string                `json:"documentDate,omitempty"`
	DocumentType         string                `json:"documentType,omitempty"`
	GlClass              string                `json:"glClass,omitempty"`
	GlTypeId             string                `json:"glTypeId,omitempty"`
	ID                   int                   `json:"id,omitempty"`
	Member               Member                `json:"member,omitempty"`
	Memo                 string                `json:"memo,omitempty"`
	Offset               GLExportExpenseOffset `json:"offset,omitempty"`
	PeriodEndDate        string                `json:"periodEndDate,omitempty"`
	PeriodStartDate      string                `json:"periodStartDate,omitempty"`
	Project              Project               `json:"project,omitempty"`
	Total                float64               `json:"total,omitempty"`
	VendorNumber         string                `json:"vendorNumber,omitempty"`
}

type GLExportExpenseBill struct {
	ApAccountNumber string                      `json:"apAccountNumber,omitempty"`
	Currency        Currency                    `json:"currency,omitempty"`
	Detail          []GLExportExpenseBillDetail `json:"detail,omitempty"`
	DocumentDate    string                      `json:"documentDate,omitempty"`
	DocumentNumber  string                      `json:"documentNumber,omitempty"`
	DocumentType    string                      `json:"documentType,omitempty"`
	GlClass         string                      `json:"glClass,omitempty"`
	ID              int                         `json:"id,omitempty"`
	Member          Member                      `json:"member,omitempty"`
	Memo            string                      `json:"memo,omitempty"`
	Total           float64                     `json:"total,omitempty"`
	VendorNumber    string                      `json:"vendorNumber,omitempty"`
}

type GLExportExpenseBillDetail struct {
	AccountNumber  string         `json:"accountNumber,omitempty"`
	Billable       bool           `json:"billable,omitempty"`
	Company        Company        `json:"company,omitempty"`
	CompanyAdvance bool           `json:"companyAdvance,omitempty"`
	Currency       Currency       `json:"currency,omitempty"`
	DocumentDate   string         `json:"documentDate,omitempty"`
	ExpenseClass   Classification `json:"expenseClass,omitempty"`
	GlTypeId       string         `json:"glTypeId,omitempty"`
	ID             []int          `json:"id,omitempty"`
	Memo           string         `json:"memo,omitempty"`
	Reimbursable   bool           `json:"reimbursable,omitempty"`
	Total          float64        `json:"total,omitempty"`
}

type GLExportExpenseOffset struct {
	AccountNumber string  `json:"accountNumber,omitempty"`
	Description   string  `json:"description,omitempty"`
	DocumentDate  string  `json:"documentDate,omitempty"`
	DocumentType  string  `json:"documentType,omitempty"`
	GlClass       string  `json:"glClass,omitempty"`
	GlTypeId      string  `json:"glTypeId,omitempty"`
	ID            int     `json:"id,omitempty"`
	Member        Member  `json:"member,omitempty"`
	Memo          string  `json:"memo,omitempty"`
	Total         float64 `json:"total,omitempty"`
}

type GLExportInventoryTransfer struct {
	AccountNumber            string                          `json:"accountNumber,omitempty"`
	Bin                      WarehouseBin                    `json:"bin,omitempty"`
	CogsXref                 string                          `json:"cogsXref,omitempty"`
	Cost                     float64                         `json:"cost,omitempty"`
	CostAccountNumber        string                          `json:"costAccountNumber,omitempty"`
	Currency                 Currency                        `json:"currency,omitempty"`
	Description              string                          `json:"description,omitempty"`
	DocumentDate             string                          `json:"documentDate,omitempty"`
	DocumentType             string                          `json:"documentType,omitempty"`
	GlClass                  string                          `json:"glClass,omitempty"`
	GlItemId                 string                          `json:"glItemId,omitempty"`
	GlTypeId                 string                          `json:"glTypeId,omitempty"`
	ID                       string                          `json:"id,omitempty"`
	InventoryAccountNumber   string                          `json:"inventoryAccountNumber,omitempty"`
	InventoryXref            string                          `json:"inventoryXref,omitempty"`
	Item                     IvItem                          `json:"item,omitempty"`
	ItemDescription          string                          `json:"itemDescription,omitempty"`
	ItemPrice                float64                         `json:"itemPrice,omitempty"`
	ItemTypeXref             string                          `json:"itemTypeXref,omitempty"`
	LocationXref             string                          `json:"locationXref,omitempty"`
	Memo                     string                          `json:"memo,omitempty"`
	Offset                   GLExportInventoryTransferOffset `json:"offset,omitempty"`
	PriceLevelXref           string                          `json:"priceLevelXref,omitempty"`
	Quantity                 float64                         `json:"quantity,omitempty"`
	SalesCode                string                          `json:"salesCode,omitempty"`
	SalesDescription         string                          `json:"salesDescription,omitempty"`
	SerialNumbers            string                          `json:"serialNumbers,omitempty"`
	SerializedFlag           bool                            `json:"serializedFlag,omitempty"`
	SubCategory              ProductSubCategory              `json:"subCategory,omitempty"`
	TaxNote                  string                          `json:"taxNote,omitempty"`
	Taxable                  bool                            `json:"taxable,omitempty"`
	Total                    float64                         `json:"total,omitempty"`
	TransferFromBin          WarehouseBin                    `json:"transferFromBin,omitempty"`
	TransferFromLocationXref string                          `json:"transferFromLocationXref,omitempty"`
	TransferId               int                             `json:"transferId,omitempty"`
	TransferToBin            WarehouseBin                    `json:"transferToBin,omitempty"`
	TransferToLocationXref   string                          `json:"transferToLocationXref,omitempty"`
	UnitOfMeasure            UnitOfMeasure                   `json:"unitOfMeasure,omitempty"`
	UomScheduleXref          string                          `json:"uomScheduleXref,omitempty"`
	Warehouse                Warehouse                       `json:"warehouse,omitempty"`
}

type GLExportInventoryTransferOffset struct {
	AccountNumber string  `json:"accountNumber,omitempty"`
	Description   string  `json:"description,omitempty"`
	DocumentDate  string  `json:"documentDate,omitempty"`
	DocumentType  string  `json:"documentType,omitempty"`
	GlClass       string  `json:"glClass,omitempty"`
	GlTypeId      string  `json:"glTypeId,omitempty"`
	ID            int     `json:"id,omitempty"`
	Memo          string  `json:"memo,omitempty"`
	Total         float64 `json:"total,omitempty"`
}

type GLExportPurchaseTransaction struct {
	ApAccountNumber                  string                                 `json:"apAccountNumber,omitempty"`
	BillingTerms                     BillingTerms                           `json:"billingTerms,omitempty"`
	BillingTermsXref                 string                                 `json:"billingTermsXref,omitempty"`
	CityTaxXref                      string                                 `json:"cityTaxXref,omitempty"`
	Company                          Company                                `json:"company,omitempty"`
	CompanyType                      CompanyType                            `json:"companyType,omitempty"`
	Contact                          Contact                                `json:"contact,omitempty"`
	CountyTaxXref                    string                                 `json:"countyTaxXref,omitempty"`
	Currency                         Currency                               `json:"currency,omitempty"`
	Description                      string                                 `json:"description,omitempty"`
	DocumentDate                     string                                 `json:"documentDate,omitempty"`
	DocumentNumber                   string                                 `json:"documentNumber,omitempty"`
	DropshipFlag                     bool                                   `json:"dropshipFlag,omitempty"`
	DueDays                          int                                    `json:"dueDays,omitempty"`
	FreightAmount                    float64                                `json:"freightAmount,omitempty"`
	FreightPackingSlip               string                                 `json:"freightPackingSlip,omitempty"`
	ID                               string                                 `json:"id,omitempty"`
	Memo                             string                                 `json:"memo,omitempty"`
	PackingSlip                      string                                 `json:"packingSlip,omitempty"`
	PurchaseClass                    string                                 `json:"purchaseClass,omitempty"`
	PurchaseDate                     string                                 `json:"purchaseDate,omitempty"`
	PurchaseDetail                   []GLExportPurchaseTransactionDetail    `json:"purchaseDetail,omitempty"`
	PurchaseDetailTax                []GLExportPurchaseTransactionDetailTax `json:"purchaseDetailTax,omitempty"`
	PurchaseHeaderFreightTaxableFlag bool                                   `json:"purchaseHeaderFreightTaxableFlag,omitempty"`
	PurchaseHeaderTaxGroup           string                                 `json:"purchaseHeaderTaxGroup,omitempty"`
	PurchaseHeaderTaxableFlag        bool                                   `json:"purchaseHeaderTaxableFlag,omitempty"`
	ShipToCompany                    Company                                `json:"shipToCompany,omitempty"`
	ShipToCompanyAccountNumber       string                                 `json:"shipToCompanyAccountNumber,omitempty"`
	ShipToCompanyType                CompanyType                            `json:"shipToCompanyType,omitempty"`
	ShipToContact                    Contact                                `json:"shipToContact,omitempty"`
	ShipToSite                       Site                                   `json:"shipToSite,omitempty"`
	ShipToTaxGroup                   string                                 `json:"shipToTaxGroup,omitempty"`
	Site                             Site                                   `json:"site,omitempty"`
	StateTaxXref                     string                                 `json:"stateTaxXref,omitempty"`
	TaxAgencyXref                    string                                 `json:"taxAgencyXref,omitempty"`
	TaxCode                          TaxCode                                `json:"taxCode,omitempty"`
	TaxGroupRate                     float64                                `json:"taxGroupRate,omitempty"`
	TaxLevels                        []GLExportPurchaseTransactionTaxLevel  `json:"taxLevels,omitempty"`
	Total                            float64                                `json:"total,omitempty"`
	UseAvalaraTaxFlag                bool                                   `json:"useAvalaraTaxFlag,omitempty"`
	VendorAccountNumber              string                                 `json:"vendorAccountNumber,omitempty"`
	VendorInvoiceDate                string                                 `json:"vendorInvoiceDate,omitempty"`
	VendorInvoiceNumber              string                                 `json:"vendorInvoiceNumber,omitempty"`
	VendorNumber                     string                                 `json:"vendorNumber,omitempty"`
}

type GLExportPurchaseTransactionDetail struct {
	AccountNumber          string             `json:"accountNumber,omitempty"`
	CogsXref               string             `json:"cogsXref,omitempty"`
	Cost                   float64            `json:"cost,omitempty"`
	CostAccountNumber      string             `json:"costAccountNumber,omitempty"`
	Currency               Currency           `json:"currency,omitempty"`
	Description            string             `json:"description,omitempty"`
	DocumentDate           string             `json:"documentDate,omitempty"`
	DropShippedFlag        bool               `json:"dropShippedFlag,omitempty"`
	GlClass                string             `json:"glClass,omitempty"`
	GlItemId               string             `json:"glItemId,omitempty"`
	GlTypeId               string             `json:"glTypeId,omitempty"`
	ID                     int                `json:"id,omitempty"`
	InventoryAccountNumber string             `json:"inventoryAccountNumber,omitempty"`
	InventoryXref          string             `json:"inventoryXref,omitempty"`
	Item                   IvItem             `json:"item,omitempty"`
	ItemCost               float64            `json:"itemCost,omitempty"`
	ItemDescription        string             `json:"itemDescription,omitempty"`
	ItemPrice              float64            `json:"itemPrice,omitempty"`
	ItemTypeXref           string             `json:"itemTypeXref,omitempty"`
	LineNumber             int                `json:"lineNumber,omitempty"`
	LocationXref           string             `json:"locationXref,omitempty"`
	Memo                   string             `json:"memo,omitempty"`
	PriceLevelXref         string             `json:"priceLevelXref,omitempty"`
	PurchaseHeaderTaxGroup string             `json:"purchaseHeaderTaxGroup,omitempty"`
	Quantity               float64            `json:"quantity,omitempty"`
	SalesCode              string             `json:"salesCode,omitempty"`
	SalesDescription       string             `json:"salesDescription,omitempty"`
	SerialNumbers          string             `json:"serialNumbers,omitempty"`
	SerializedFlag         bool               `json:"serializedFlag,omitempty"`
	ShipmentMethod         ShipmentMethod     `json:"shipmentMethod,omitempty"`
	SubCategory            ProductSubCategory `json:"subCategory,omitempty"`
	TaxAgencyXref          string             `json:"taxAgencyXref,omitempty"`
	TaxCode                TaxCode            `json:"taxCode,omitempty"`
	TaxCodeXref            string             `json:"taxCodeXref,omitempty"`
	TaxNote                string             `json:"taxNote,omitempty"`
	TaxRate                float64            `json:"taxRate,omitempty"`
	Taxable                bool               `json:"taxable,omitempty"`
	Total                  float64            `json:"total,omitempty"`
	UnitOfMeasure          UnitOfMeasure      `json:"unitOfMeasure,omitempty"`
	UomScheduleXref        string             `json:"uomScheduleXref,omitempty"`
	VendorAccountNumber    string             `json:"vendorAccountNumber,omitempty"`
	VendorNumber           string             `json:"vendorNumber,omitempty"`
	WarehouseBin           WarehouseBin       `json:"warehouseBin,omitempty"`
	WarehouseSite          Site               `json:"warehouseSite,omitempty"`
}

type GLExportPurchaseTransactionDetailTax struct {
	AccountNumber          string             `json:"accountNumber,omitempty"`
	CogsXref               string             `json:"cogsXref,omitempty"`
	Cost                   float64            `json:"cost,omitempty"`
	CostAccountNumber      string             `json:"costAccountNumber,omitempty"`
	Currency               Currency           `json:"currency,omitempty"`
	DocumentDate           string             `json:"documentDate,omitempty"`
	DropShippedFlag        bool               `json:"dropShippedFlag,omitempty"`
	GlClass                string             `json:"glClass,omitempty"`
	GlItemId               string             `json:"glItemId,omitempty"`
	GlTypeId               string             `json:"glTypeId,omitempty"`
	ID                     int                `json:"id,omitempty"`
	InventoryAccountNumber string             `json:"inventoryAccountNumber,omitempty"`
	InventoryXref          string             `json:"inventoryXref,omitempty"`
	Item                   IvItem             `json:"item,omitempty"`
	ItemCost               float64            `json:"itemCost,omitempty"`
	ItemDescription        string             `json:"itemDescription,omitempty"`
	ItemPrice              float64            `json:"itemPrice,omitempty"`
	ItemTypeXref           string             `json:"itemTypeXref,omitempty"`
	LineNumber             int                `json:"lineNumber,omitempty"`
	LocationXref           string             `json:"locationXref,omitempty"`
	Memo                   string             `json:"memo,omitempty"`
	PriceLevelXref         string             `json:"priceLevelXref,omitempty"`
	PurchaseHeaderTaxGroup string             `json:"purchaseHeaderTaxGroup,omitempty"`
	Quantity               float64            `json:"quantity,omitempty"`
	SalesCode              string             `json:"salesCode,omitempty"`
	SalesDescription       string             `json:"salesDescription,omitempty"`
	SerialNumbers          string             `json:"serialNumbers,omitempty"`
	SerializedFlag         bool               `json:"serializedFlag,omitempty"`
	ShipmentMethod         ShipmentMethod     `json:"shipmentMethod,omitempty"`
	SubCategory            ProductSubCategory `json:"subCategory,omitempty"`
	TaxAgencyXref          string             `json:"taxAgencyXref,omitempty"`
	TaxCode                TaxCode            `json:"taxCode,omitempty"`
	TaxNote                string             `json:"taxNote,omitempty"`
	TaxRate                float64            `json:"taxRate,omitempty"`
	TaxRatePercent         float64            `json:"taxRatePercent,omitempty"`
	TaxableFlag            bool               `json:"taxableFlag,omitempty"`
	Total                  float64            `json:"total,omitempty"`
	UnitOfMeasure          UnitOfMeasure      `json:"unitOfMeasure,omitempty"`
	UomScheduleXref        string             `json:"uomScheduleXref,omitempty"`
	VendorAccountNumber    string             `json:"vendorAccountNumber,omitempty"`
	VendorNumber           string             `json:"vendorNumber,omitempty"`
	WarehouseBin           WarehouseBin       `json:"warehouseBin,omitempty"`
	WarehouseSite          Site               `json:"warehouseSite,omitempty"`
}

type GLExportPurchaseTransactionTaxLevel struct {
	TaxCodeXref string `json:"taxCodeXref,omitempty"`
	TaxLevel    int    `json:"taxLevel,omitempty"`
}

type GLExportSettings struct {
}

type GLExportTransaction struct {
	AccountNumber              string                        `json:"accountNumber,omitempty"`
	AgreementPrePaymentFlag    bool                          `json:"agreementPrePaymentFlag,omitempty"`
	Attention                  string                        `json:"attention,omitempty"`
	BillingTerms               BillingTerms                  `json:"billingTerms,omitempty"`
	BillingTermsXref           string                        `json:"billingTermsXref,omitempty"`
	BillingType                string                        `json:"billingType,omitempty"`
	CityTax                    float64                       `json:"cityTax,omitempty"`
	Company                    Company                       `json:"company,omitempty"`
	CompanyAccountNumber       string                        `json:"companyAccountNumber,omitempty"`
	CompanyType                CompanyType                   `json:"companyType,omitempty"`
	CountyTax                  float64                       `json:"countyTax,omitempty"`
	CountyTaxXref              string                        `json:"countyTaxXref,omitempty"`
	Currency                   Currency                      `json:"currency,omitempty"`
	Description                string                        `json:"description,omitempty"`
	Detail                     []GLExportTransactionDetail   `json:"detail,omitempty"`
	DocumentDate               string                        `json:"documentDate,omitempty"`
	DocumentNumber             string                        `json:"documentNumber,omitempty"`
	DocumentType               string                        `json:"documentType,omitempty"`
	DueDate                    string                        `json:"dueDate,omitempty"`
	DueDays                    int                           `json:"dueDays,omitempty"`
	EmailDeliveryFlag          bool                          `json:"emailDeliveryFlag,omitempty"`
	GlClass                    string                        `json:"glClass,omitempty"`
	GlEntryIds                 string                        `json:"glEntryIds,omitempty"`
	GlTypeId                   string                        `json:"glTypeId,omitempty"`
	ID                         int                           `json:"id,omitempty"`
	Memo                       string                        `json:"memo,omitempty"`
	PiggyBackFlag              bool                          `json:"piggyBackFlag,omitempty"`
	PrintDeliveryFlag          bool                          `json:"printDeliveryFlag,omitempty"`
	Project                    Project                       `json:"project,omitempty"`
	PurchaseOrder              PurchaseOrder                 `json:"purchaseOrder,omitempty"`
	SalesRepId                 string                        `json:"salesRepId,omitempty"`
	SalesRepName               string                        `json:"salesRepName,omitempty"`
	SalesTax                   float64                       `json:"salesTax,omitempty"`
	SalesTerritory             string                        `json:"salesTerritory,omitempty"`
	SendAvalaraTaxFlag         bool                          `json:"sendAvalaraTaxFlag,omitempty"`
	ShipContact                string                        `json:"shipContact,omitempty"`
	ShipSite                   Site                          `json:"shipSite,omitempty"`
	ShipToCompany              Company                       `json:"shipToCompany,omitempty"`
	ShipToCompanyAccountNumber string                        `json:"shipToCompanyAccountNumber,omitempty"`
	ShipToCompanyType          CompanyType                   `json:"shipToCompanyType,omitempty"`
	ShipToTaxId                string                        `json:"shipToTaxId,omitempty"`
	Site                       Site                          `json:"site,omitempty"`
	StateTax                   float64                       `json:"stateTax,omitempty"`
	StateTaxXref               string                        `json:"stateTaxXref,omitempty"`
	TaxAccountNumber           string                        `json:"taxAccountNumber,omitempty"`
	TaxAgencyXref              string                        `json:"taxAgencyXref,omitempty"`
	TaxCode                    TaxCode                       `json:"taxCode,omitempty"`
	TaxDpAppliedFlag           bool                          `json:"taxDpAppliedFlag,omitempty"`
	TaxGroupRate               float64                       `json:"taxGroupRate,omitempty"`
	TaxId                      string                        `json:"taxId,omitempty"`
	TaxLevels                  []GLExportTransactionTaxLevel `json:"taxLevels,omitempty"`
	Taxable                    bool                          `json:"taxable,omitempty"`
	TaxableAmount1             float64                       `json:"taxableAmount1,omitempty"`
	TaxableAmount2             float64                       `json:"taxableAmount2,omitempty"`
	TaxableAmount3             float64                       `json:"taxableAmount3,omitempty"`
	TaxableAmount4             float64                       `json:"taxableAmount4,omitempty"`
	TaxableAmount5             float64                       `json:"taxableAmount5,omitempty"`
	TaxableTotal               float64                       `json:"taxableTotal,omitempty"`
	Total                      float64                       `json:"total,omitempty"`
	UseAvalaraFlag             bool                          `json:"useAvalaraFlag,omitempty"`
}

type GLExportTransactionDetail struct {
	AccountNumber          string                              `json:"accountNumber,omitempty"`
	CogsXref               string                              `json:"cogsXref,omitempty"`
	Cost                   float64                             `json:"cost,omitempty"`
	CostAccountNumber      string                              `json:"costAccountNumber,omitempty"`
	Currency               Currency                            `json:"currency,omitempty"`
	Description            string                              `json:"description,omitempty"`
	DocumentDate           string                              `json:"documentDate,omitempty"`
	DocumentType           string                              `json:"documentType,omitempty"`
	DropShippedFlag        bool                                `json:"dropShippedFlag,omitempty"`
	GlClass                string                              `json:"glClass,omitempty"`
	GlItemId               string                              `json:"glItemId,omitempty"`
	GlTypeId               string                              `json:"glTypeId,omitempty"`
	ID                     int                                 `json:"id,omitempty"`
	InventoryAccountNumber string                              `json:"inventoryAccountNumber,omitempty"`
	InventoryXref          string                              `json:"inventoryXref,omitempty"`
	InvoiceSummaryOption   string                              `json:"invoiceSummaryOption,omitempty"`
	Item                   IvItem                              `json:"item,omitempty"`
	ItemCost               float64                             `json:"itemCost,omitempty"`
	ItemDescription        string                              `json:"itemDescription,omitempty"`
	ItemPrice              float64                             `json:"itemPrice,omitempty"`
	ItemTaxableFlag        bool                                `json:"itemTaxableFlag,omitempty"`
	ItemTypeXref           string                              `json:"itemTypeXref,omitempty"`
	LocationXref           string                              `json:"locationXref,omitempty"`
	Memo                   string                              `json:"memo,omitempty"`
	PriceLevelXref         string                              `json:"priceLevelXref,omitempty"`
	Product                Product                             `json:"product,omitempty"`
	ProductAccountNumber   string                              `json:"productAccountNumber,omitempty"`
	Quantity               float64                             `json:"quantity,omitempty"`
	SalesCode              string                              `json:"salesCode,omitempty"`
	SalesDescription       string                              `json:"salesDescription,omitempty"`
	SerialNumbers          string                              `json:"serialNumbers,omitempty"`
	SerializedFlag         bool                                `json:"serializedFlag,omitempty"`
	ShipmentMethod         ShipmentMethod                      `json:"shipmentMethod,omitempty"`
	SubCategory            ProductSubCategory                  `json:"subCategory,omitempty"`
	TaxAgencyXref          string                              `json:"taxAgencyXref,omitempty"`
	TaxCode                TaxCode                             `json:"taxCode,omitempty"`
	TaxCodeXref            string                              `json:"taxCodeXref,omitempty"`
	TaxLevels              []GLExportTransactionDetailTaxLevel `json:"taxLevels,omitempty"`
	TaxNote                string                              `json:"taxNote,omitempty"`
	TaxRate                float64                             `json:"taxRate,omitempty"`
	TaxRatePercent         float64                             `json:"taxRatePercent,omitempty"`
	Taxable2Flag           bool                                `json:"taxable2Flag,omitempty"`
	Taxable3Flag           bool                                `json:"taxable3Flag,omitempty"`
	Taxable4Flag           bool                                `json:"taxable4Flag,omitempty"`
	Taxable5Flag           bool                                `json:"taxable5Flag,omitempty"`
	TaxableFlag            bool                                `json:"taxableFlag,omitempty"`
	TimeEntry              TimeEntry                           `json:"timeEntry,omitempty"`
	Total                  float64                             `json:"total,omitempty"`
	UnitOfMeasure          UnitOfMeasure                       `json:"unitOfMeasure,omitempty"`
	UomScheduleXref        string                              `json:"uomScheduleXref,omitempty"`
	WarehouseBin           WarehouseBin                        `json:"warehouseBin,omitempty"`
	WarehouseSite          Site                                `json:"warehouseSite,omitempty"`
}

type GLExportTransactionDetailTaxLevel struct {
	TaxLevel    int  `json:"taxLevel,omitempty"`
	TaxableFlag bool `json:"taxableFlag,omitempty"`
}

type GLExportTransactionTaxLevel struct {
	TaxAmount     float64 `json:"taxAmount,omitempty"`
	TaxCodeXref   string  `json:"taxCodeXref,omitempty"`
	TaxLevel      int     `json:"taxLevel,omitempty"`
	TaxableAmount float64 `json:"taxableAmount,omitempty"`
}

type GLExportVendor struct {
	AccountNumber string       `json:"accountNumber,omitempty"`
	BillingTerms  BillingTerms `json:"billingTerms,omitempty"`
	Company       Company      `json:"company,omitempty"`
	Contact       Contact      `json:"contact,omitempty"`
	DueDays       int          `json:"dueDays,omitempty"`
	Member        Member       `json:"member,omitempty"`
	Site          Site         `json:"site,omitempty"`
	TaxCode       TaxCode      `json:"taxCode,omitempty"`
	Vendor        Company      `json:"vendor,omitempty"`
	VendorNumber  string       `json:"vendorNumber,omitempty"`
}

type GenericBoardTeam struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	ID                int                    `json:"id,omitempty"`
	IsProjectTeamFlag bool                   `json:"isProjectTeamFlag,omitempty"`
	Name              string                 `json:"name,omitempty"`
}

type GenericIdentifier struct {
	ID         int    `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name       string `json:"name,omitempty"`
}

type GenericNameIdDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
}

type HttpContent struct {
	Headers []interface{} `json:"headers,omitempty"`
}

type HttpMethod struct {
	Delete  *HttpMethod `json:"delete,omitempty"`
	Get     *HttpMethod `json:"get,omitempty"`
	Head    *HttpMethod `json:"head,omitempty"`
	Method  string      `json:"method,omitempty"`
	Options *HttpMethod `json:"options,omitempty"`
	Post    *HttpMethod `json:"post,omitempty"`
	Put     *HttpMethod `json:"put,omitempty"`
	Trace   *HttpMethod `json:"trace,omitempty"`
}

type HttpRequestMessage struct {
	Content    HttpContent            `json:"content,omitempty"`
	Headers    []interface{}          `json:"headers,omitempty"`
	Method     *HttpMethod            `json:"method,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	RequestUri string                 `json:"requestUri,omitempty"`
	Version    Version                `json:"version,omitempty"`
}

type IRestIdentifiedItem struct {
	ID int `json:"id,omitempty"`
}

type ImapSetup struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type InclusiveRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type IntegratorTagCollection struct {
	Tags []string `json:"tags,omitempty"`
}

type InvoiceTemplateDetail struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type IvItem struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	ID             int                    `json:"id,omitempty"`
	Identifier     string                 `json:"identifier,omitempty"`
	SerializedFlag bool                   `json:"serializedFlag,omitempty"`
}

type KBCategory struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type LicenseBit struct {
	ActiveFlag bool   `json:"activeFlag,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Locale struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type LostRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type ManagementSolution struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	ID        int                    `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	SetupName string                 `json:"setupName,omitempty"`
}

type MappedRecord struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MemberDeactivationCompanyTeam struct {
	Count             int     `json:"count,omitempty"`
	ID                int     `json:"id,omitempty"`
	Name              string  `json:"name,omitempty"`
	ReAssignToContact Contact `json:"reAssignToContact,omitempty"`
	ReAssignToMember  Member  `json:"reAssignToMember,omitempty"`
}

type MemberDeactivationItem struct {
	Count            int    `json:"count,omitempty"`
	ReAssignToMember Member `json:"reAssignToMember,omitempty"`
}

type MemberDeactivationStatusWorkflow struct {
	Count            int    `json:"count,omitempty"`
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	ReAssignToMember Member `json:"reAssignToMember,omitempty"`
}

type MemberOffice365 struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MemberSsoSettings struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	Email     string                 `json:"email,omitempty"`
	ID        int                    `json:"id,omitempty"`
	SsoUserId string                 `json:"ssoUserId,omitempty"`
	UserName  string                 `json:"userName,omitempty"`
}

type MenuLocation struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type NoteType struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type NotifyType struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier,omitempty"`
	Name       string                 `json:"name,omitempty"`
}

type OpenRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type OpportunityPriority struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type OpportunityProbability struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type OpportunitySalesRole struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type OtherRevenue1 struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type OtherRevenue2 struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type OwnerLevel struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type PageValues struct {
	Page     int `json:"page,omitempty"`
	PageId   int `json:"pageId,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type PaymentMethod struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ProductCategory struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ProductDemand struct {
	Cost         float64 `json:"cost,omitempty"`
	ProductRecId int     `json:"productRecId,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
}

type ProductRecurring struct {
	AgreementType    AgreementType `json:"agreementType,omitempty"`
	BillCycle        BillingCycle  `json:"billCycle,omitempty"`
	BillCycleId      int           `json:"billCycleId,omitempty"`
	CycleType        string        `json:"cycleType,omitempty"`
	Cycles           int           `json:"cycles,omitempty"`
	EndDate          string        `json:"endDate,omitempty"`
	RecurringCost    float64       `json:"recurringCost,omitempty"`
	RecurringRevenue float64       `json:"recurringRevenue,omitempty"`
	StartDate        string        `json:"startDate,omitempty"`
}

type ProductRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type ProductSubCategory struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ProjectBoardKanbanStatus struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Order      int    `json:"order,omitempty"`
	SrStatusId int    `json:"srStatusId,omitempty"`
}

type ProjectBoard struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ProjectRole struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier,omitempty"`
}

type ProjectWorkplanProjectPhase struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	ActualHours         float64                `json:"actualHours,omitempty"`
	BillPhaseSeparately bool                   `json:"billPhaseSeparately,omitempty"`
	BillableHours       float64                `json:"billableHours,omitempty"`
	BudgetHours         float64                `json:"budgetHours,omitempty"`
	CustomFields        []CustomFieldValue     `json:"customFields,omitempty"`
	Description         string                 `json:"description,omitempty"`
	EndDate             string                 `json:"endDate,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	MarkAsMilestoneFlag bool                   `json:"markAsMilestoneFlag,omitempty"`
	Notes               string                 `json:"notes,omitempty"`
	ParentPhase         ProjectPhase           `json:"parentPhase,omitempty"`
	ProjectId           int                    `json:"projectId,omitempty"`
	ScheduledDuration   int                    `json:"scheduled_Duration,omitempty"`
	ScheduledEnd        string                 `json:"scheduled_End,omitempty"`
	ScheduledHours      float64                `json:"scheduled_Hours,omitempty"`
	ScheduledStart      string                 `json:"scheduled_Start,omitempty"`
	StartDate           string                 `json:"startDate,omitempty"`
	Status              PhaseStatus            `json:"status,omitempty"`
	WbsCode             string                 `json:"wbsCode,omitempty"`
}

type Relationship struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type Reminder struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ResultInfo struct {
	Data          IRestIdentifiedItem  `json:"data,omitempty"`
	Error         ErrorResponseMessage `json:"error,omitempty"`
	OriginalIndex int                  `json:"originalIndex,omitempty"`
	StatusCode    int                  `json:"statusCode,omitempty"`
	Success       bool                 `json:"success,omitempty"`
}

type SalesOrder struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier,omitempty"`
}

type ScheduleSpan struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier,omitempty"`
}

type ServiceCode struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ServiceItem struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ServiceRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type ServiceSource struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ServiceStatus struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
	Sort int                    `json:"sort,omitempty"`
}

type ServiceSubType struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ServiceSurveyQuestionOption struct {
	Caption     string `json:"caption,omitempty"`
	IncludeFlag bool   `json:"includeFlag,omitempty"`
	Points      int    `json:"points,omitempty"`
}

type ServiceType struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type SicCode struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type Site struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type Structure struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type SurveyResultDetail struct {
	Answer     map[string]interface{} `json:"answer,omitempty"`
	QuestionId int                    `json:"questionId,omitempty"`
}

type SystemDepartment struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier,omitempty"`
	Name       string                 `json:"name,omitempty"`
}

type SystemLocation struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type SystemMenuEntry struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type TemplatePhase struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	BillPhaseSeparately bool                   `json:"billPhaseSeparately,omitempty"`
	Description         string                 `json:"description,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	MarkAsMilestoneFlag bool                   `json:"markAsMilestoneFlag,omitempty"`
	Notes               string                 `json:"notes,omitempty"`
	ParentPhase         ProjectTemplatePhase   `json:"parentPhase,omitempty"`
	TemplateId          int                    `json:"templateId,omitempty"`
	WbsCode             string                 `json:"wbsCode,omitempty"`
}

type TimeRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}

type TimeZone struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type UserDefinedFieldOption struct {
	DefaultFlag  bool   `json:"defaultFlag,omitempty"`
	ID           int    `json:"id,omitempty"`
	InactiveFlag bool   `json:"inactiveFlag,omitempty"`
	OptionValue  string `json:"optionValue,omitempty"`
	SortOrder    int    `json:"sortOrder,omitempty"`
}

type UserDefinedFieldValueModel struct {
	Filtered                   bool   `json:"filtered,omitempty"`
	RowNum                     int    `json:"rowNum,omitempty"`
	SkipLocationAndBillingUnit bool   `json:"skipLocationAndBillingUnit,omitempty"`
	UserDefinedFieldRecId      int    `json:"userDefinedFieldRecId,omitempty"`
	Value                      string `json:"value,omitempty"`
}

type ValidationError struct {
	Code     string `json:"code,omitempty"`
	Details  string `json:"details,omitempty"`
	Field    string `json:"field,omitempty"`
	Message  string `json:"message,omitempty"`
	Resource string `json:"resource,omitempty"`
}

type Version struct {
	Build         int `json:"build,omitempty"`
	Major         int `json:"major,omitempty"`
	MajorRevision int `json:"majorRevision,omitempty"`
	Minor         int `json:"minor,omitempty"`
	MinorRevision int `json:"minorRevision,omitempty"`
	Revision      int `json:"revision,omitempty"`
}

type WisePayBatchPayment struct {
	Amount      float64 `json:"amount,omitempty"`
	WisePayHref string  `json:"wisePayHref,omitempty"`
}

type WisePayFeeInvoice struct {
	Amount        float64 `json:"amount,omitempty"`
	ID            int     `json:"id,omitempty"`
	InvoiceHref   string  `json:"invoiceHref,omitempty"`
	InvoiceNumber string  `json:"invoiceNumber,omitempty"`
}

type WisePayPayment struct {
	BatchPayment   WisePayBatchPayment `json:"batchPayment,omitempty"`
	FeeInvoice     WisePayFeeInvoice   `json:"feeInvoice,omitempty"`
	PaymentDateUtc string              `json:"paymentDateUtc,omitempty"`
	WisePay        string              `json:"wisePay,omitempty"`
}

type WonRevenue struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Cost       float64                `json:"cost,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Margin     float64                `json:"margin,omitempty"`
	Percentage float64                `json:"percentage,omitempty"`
	Revenue    float64                `json:"revenue,omitempty"`
}
