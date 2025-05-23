package types

type AccountingPackageReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type ActivityStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ActivityTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type AddressFormatReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type AdjustmentDetailReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type AdjustmentReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type AdjustmentTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type AgreementApplicationAviablePer struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag string `json:"tag,omitempty"`
}

type AgreementApplicationBillingCycle struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag string `json:"tag,omitempty"`
}

type AgreementApplicationLimit struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag string `json:"tag,omitempty"`
}

type AgreementApplicationUnit struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag string `json:"tag,omitempty"`
}

type AgreementBillingInfo struct {
	AgreementAmount float64 `json:"agreementAmount,omitempty"`
	AgreementName string `json:"agreementName,omitempty"`
	AgreementRecId int `json:"agreementRecId,omitempty"`
	AgreementType string `json:"agreementType,omitempty"`
	ParentRecId int `json:"parentRecId,omitempty"`
}

type AgreementReference struct {
	Info interface{} `json:"_info,omitempty"`
	ChargeFirmFlag bool `json:"chargeFirmFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type AgreementRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type AgreementTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ApiRequest struct {
	Entity IRestIdentifiedItem `json:"entity,omitempty"`
	ExternalId string `json:"externalId,omitempty"`
	Fields string `json:"fields,omitempty"`
	Filters FilterValues `json:"filters,omitempty"`
	GrandParentId int `json:"grandParentId,omitempty"`
	ID int `json:"id,omitempty"`
	MemberContext string `json:"memberContext,omitempty"`
	MiscProperties interface{} `json:"miscProperties,omitempty"`
	Page PageValues `json:"page,omitempty"`
	ParentId int `json:"parentId,omitempty"`
	UpdateOnlyCesProperties bool `json:"updateOnlyCesProperties,omitempty"`
}

type AutomateScriptReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BatchReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BillingCycleReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BillingDeliveryReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BillingSetupReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BillingStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	IsClosed bool `json:"isClosed,omitempty"`
	Name string `json:"name,omitempty"`
}

type BillingTermsReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BillingUnitReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BoardReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BundleRequest struct {
	ApiRequest ApiRequest `json:"apiRequest,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	SequenceNumber int `json:"sequenceNumber,omitempty"`
	Version string `json:"version,omitempty"`
}

type BundleResult struct {
	Count int `json:"count,omitempty"`
	Entities []IRestIdentifiedItem `json:"entities,omitempty"`
	Error ErrorResponseMessage `json:"error,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	SequenceNumber int `json:"sequenceNumber,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
	Success bool `json:"success,omitempty"`
	Version string `json:"version,omitempty"`
}

type CalendarReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CalendarSetupReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
}

type CampaignReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CampaignStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CampaignSubTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CampaignTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CatalogItemReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type CertificationReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ChargeCodeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ClassificationReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CommunicationTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CompanyReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type CompanyStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CompanyTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ConfigurationQuestion struct {
	Answer interface{} `json:"answer,omitempty"`
	AnswerId int `json:"answerId,omitempty"`
	FieldType string `json:"fieldType,omitempty"`
	NumberOfDecimals int `json:"numberOfDecimals,omitempty"`
	Question string `json:"question,omitempty"`
	QuestionId int `json:"questionId,omitempty"`
	RequiredFlag bool `json:"requiredFlag,omitempty"`
	SequenceNumber float64 `json:"sequenceNumber,omitempty"`
}

type ConfigurationStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ConfigurationTypeQuestionReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
}

type ConfigurationTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ContactCommunicationItem struct {
	CommunicationType string `json:"communicationType,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	Domain string `json:"domain,omitempty"`
	Extension string `json:"extension,omitempty"`
	ID int `json:"id,omitempty"`
	Type CommunicationTypeReference `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type ContactDepartmentReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ContactReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ContactTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ConversionTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CorporateStructureLevelReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CountryReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type CurrencyCodeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CurrencyReference struct {
	Info interface{} `json:"_info,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
	CurrencyIdentifier string `json:"currencyIdentifier,omitempty"`
	DecimalSeparator string `json:"decimalSeparator,omitempty"`
	DisplayIdFlag bool `json:"displayIdFlag,omitempty"`
	DisplaySymbolFlag bool `json:"displaySymbolFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	NegativeParenthesesFlag bool `json:"negativeParenthesesFlag,omitempty"`
	NumberOfDecimals int `json:"numberOfDecimals,omitempty"`
	RightAlign bool `json:"rightAlign,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	ThousandsSeparator string `json:"thousandsSeparator,omitempty"`
}

type CustomFieldValue struct {
	Caption string `json:"caption,omitempty"`
	ConnectWiseId string `json:"connectWiseId,omitempty"`
	EntryMethod string `json:"entryMethod,omitempty"`
	ID int `json:"id,omitempty"`
	NumberOfDecimals int `json:"numberOfDecimals,omitempty"`
	Type string `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type CustomReportReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DirectionalSyncReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DocumentTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EmailConnectorParsingStyleReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EmailConnectorParsingTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EmailConnectorParsingVariableReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type EmailConnectorReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EmailTemplateReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EntityTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ErrorResponseMessage struct {
	Code string `json:"code,omitempty"`
	Errors []ValidationError `json:"errors,omitempty"`
	Message string `json:"message,omitempty"`
}

type ExistingTenantReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ExpenseDetailReference struct {
	Info interface{} `json:"_info,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	ID int `json:"id,omitempty"`
}

type ExpenseReportReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ExpenseRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type ExpenseTax struct {
	Amount float64 `json:"amount,omitempty"`
	ID int `json:"id,omitempty"`
	Type ExpenseTaxTypeReference `json:"type,omitempty"`
}

type ExpenseTaxTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ExpenseTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ForecastRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type GLExportAdjustmentTransaction struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	AdjustmentDescription string `json:"adjustmentDescription,omitempty"`
	AdjustmentDetail []GLExportAdjustmentTransactionDetail `json:"adjustmentDetail,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DocumentType string `json:"documentType,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	GlTypeID string `json:"glTypeID,omitempty"`
	ID string `json:"id,omitempty"`
	Memo string `json:"memo,omitempty"`
}

type GLExportAdjustmentTransactionDetail struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	CostAccountNumber string `json:"costAccountNumber,omitempty"`
	Description string `json:"description,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	InventoryAccountNumber string `json:"inventoryAccountNumber,omitempty"`
	Item IvItemReference `json:"item,omitempty"`
	Memo string `json:"memo,omitempty"`
	ProductAccountNumber string `json:"productAccountNumber,omitempty"`
	Quantity int `json:"quantity,omitempty"`
	Total float64 `json:"total,omitempty"`
}

type GLExportCustomer struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	BillingTerms BillingTermsReference `json:"billingTerms,omitempty"`
	BillingTermsXref string `json:"billingTermsXref,omitempty"`
	CityTaxAgencyXref string `json:"cityTaxAgencyXref,omitempty"`
	CityTaxRate float64 `json:"cityTaxRate,omitempty"`
	CityTaxXref string `json:"cityTaxXref,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	CompanyType CompanyTypeReference `json:"companyType,omitempty"`
	CompositeTaxAgencyXref string `json:"compositeTaxAgencyXref,omitempty"`
	CompositeTaxRate float64 `json:"compositeTaxRate,omitempty"`
	CompositeTaxXref string `json:"compositeTaxXref,omitempty"`
	Contact ContactReference `json:"contact,omitempty"`
	CountryTaxAgencyXref string `json:"countryTaxAgencyXref,omitempty"`
	CountryTaxRate float64 `json:"countryTaxRate,omitempty"`
	CountryTaxXref string `json:"countryTaxXref,omitempty"`
	CountyTaxAgencyXref string `json:"countyTaxAgencyXref,omitempty"`
	CountyTaxRate float64 `json:"countyTaxRate,omitempty"`
	CountyTaxXref string `json:"countyTaxXref,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	DueDays int `json:"dueDays,omitempty"`
	Site SiteReference `json:"site,omitempty"`
	StateTaxAgencyXref string `json:"stateTaxAgencyXref,omitempty"`
	StateTaxRate float64 `json:"stateTaxRate,omitempty"`
	StateTaxXref string `json:"stateTaxXref,omitempty"`
	TaxAgencyXref string `json:"taxAgencyXref,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	TaxGroupRate float64 `json:"taxGroupRate,omitempty"`
	TaxLevels []GLExportCustomerTaxLevel `json:"taxLevels,omitempty"`
	Taxable bool `json:"taxable,omitempty"`
}

type GLExportCustomerTaxLevel struct {
	AgencyXref string `json:"agencyXref,omitempty"`
	TaxCodeXref string `json:"taxCodeXref,omitempty"`
	TaxLevel int `json:"taxLevel,omitempty"`
	TaxRate float64 `json:"taxRate,omitempty"`
}

type GLExportExpense struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	ApAccountNumber string `json:"apAccountNumber,omitempty"`
	ApClass string `json:"apClass,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	CompanyAccountNumber string `json:"companyAccountNumber,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DocumentType string `json:"documentType,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	GlTypeId string `json:"glTypeId,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Memo string `json:"memo,omitempty"`
	Offset GLExportExpenseOffset `json:"offset,omitempty"`
	PeriodEndDate string `json:"periodEndDate,omitempty"`
	PeriodStartDate string `json:"periodStartDate,omitempty"`
	Project ProjectReference `json:"project,omitempty"`
	Total float64 `json:"total,omitempty"`
	VendorNumber string `json:"vendorNumber,omitempty"`
}

type GLExportExpenseBill struct {
	ApAccountNumber string `json:"apAccountNumber,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	Detail []GLExportExpenseBillDetail `json:"detail,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DocumentNumber string `json:"documentNumber,omitempty"`
	DocumentType string `json:"documentType,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Memo string `json:"memo,omitempty"`
	Total float64 `json:"total,omitempty"`
	VendorNumber string `json:"vendorNumber,omitempty"`
}

type GLExportExpenseBillDetail struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	Billable bool `json:"billable,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	CompanyAdvance bool `json:"companyAdvance,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	ExpenseClass ClassificationReference `json:"expenseClass,omitempty"`
	GlTypeId string `json:"glTypeId,omitempty"`
	ID []int `json:"id,omitempty"`
	Memo string `json:"memo,omitempty"`
	Reimbursable bool `json:"reimbursable,omitempty"`
	Total float64 `json:"total,omitempty"`
}

type GLExportExpenseOffset struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DocumentType string `json:"documentType,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	GlTypeId string `json:"glTypeId,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Memo string `json:"memo,omitempty"`
	Total float64 `json:"total,omitempty"`
}

type GLExportInventoryTransfer struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	Bin WarehouseBinReference `json:"bin,omitempty"`
	CogsXref string `json:"cogsXref,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	CostAccountNumber string `json:"costAccountNumber,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DocumentType string `json:"documentType,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	GlItemId string `json:"glItemId,omitempty"`
	GlTypeId string `json:"glTypeId,omitempty"`
	ID string `json:"id,omitempty"`
	InventoryAccountNumber string `json:"inventoryAccountNumber,omitempty"`
	InventoryXref string `json:"inventoryXref,omitempty"`
	Item IvItemReference `json:"item,omitempty"`
	ItemDescription string `json:"itemDescription,omitempty"`
	ItemPrice float64 `json:"itemPrice,omitempty"`
	ItemTypeXref string `json:"itemTypeXref,omitempty"`
	LocationXref string `json:"locationXref,omitempty"`
	Memo string `json:"memo,omitempty"`
	Offset GLExportInventoryTransferOffset `json:"offset,omitempty"`
	PriceLevelXref string `json:"priceLevelXref,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	SalesCode string `json:"salesCode,omitempty"`
	SalesDescription string `json:"salesDescription,omitempty"`
	SerialNumbers string `json:"serialNumbers,omitempty"`
	SerializedFlag bool `json:"serializedFlag,omitempty"`
	SubCategory ProductSubCategoryReference `json:"subCategory,omitempty"`
	TaxNote string `json:"taxNote,omitempty"`
	Taxable bool `json:"taxable,omitempty"`
	Total float64 `json:"total,omitempty"`
	TransferFromBin WarehouseBinReference `json:"transferFromBin,omitempty"`
	TransferFromLocationXref string `json:"transferFromLocationXref,omitempty"`
	TransferId int `json:"transferId,omitempty"`
	TransferToBin WarehouseBinReference `json:"transferToBin,omitempty"`
	TransferToLocationXref string `json:"transferToLocationXref,omitempty"`
	UnitOfMeasure UnitOfMeasureReference `json:"unitOfMeasure,omitempty"`
	UomScheduleXref string `json:"uomScheduleXref,omitempty"`
	Warehouse WarehouseReference `json:"warehouse,omitempty"`
}

type GLExportInventoryTransferOffset struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DocumentType string `json:"documentType,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	GlTypeId string `json:"glTypeId,omitempty"`
	ID int `json:"id,omitempty"`
	Memo string `json:"memo,omitempty"`
	Total float64 `json:"total,omitempty"`
}

type GLExportPurchaseTransaction struct {
	ApAccountNumber string `json:"apAccountNumber,omitempty"`
	BillingTerms BillingTermsReference `json:"billingTerms,omitempty"`
	BillingTermsXref string `json:"billingTermsXref,omitempty"`
	CityTaxXref string `json:"cityTaxXref,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	CompanyType CompanyTypeReference `json:"companyType,omitempty"`
	Contact ContactReference `json:"contact,omitempty"`
	CountyTaxXref string `json:"countyTaxXref,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DocumentNumber string `json:"documentNumber,omitempty"`
	DropshipFlag bool `json:"dropshipFlag,omitempty"`
	DueDays int `json:"dueDays,omitempty"`
	FreightAmount float64 `json:"freightAmount,omitempty"`
	FreightPackingSlip string `json:"freightPackingSlip,omitempty"`
	ID string `json:"id,omitempty"`
	Memo string `json:"memo,omitempty"`
	PackingSlip string `json:"packingSlip,omitempty"`
	PurchaseClass string `json:"purchaseClass,omitempty"`
	PurchaseDate string `json:"purchaseDate,omitempty"`
	PurchaseDetail []GLExportPurchaseTransactionDetail `json:"purchaseDetail,omitempty"`
	PurchaseDetailTax []GLExportPurchaseTransactionDetailTax `json:"purchaseDetailTax,omitempty"`
	PurchaseHeaderFreightTaxableFlag bool `json:"purchaseHeaderFreightTaxableFlag,omitempty"`
	PurchaseHeaderTaxGroup string `json:"purchaseHeaderTaxGroup,omitempty"`
	PurchaseHeaderTaxableFlag bool `json:"purchaseHeaderTaxableFlag,omitempty"`
	ShipToCompany CompanyReference `json:"shipToCompany,omitempty"`
	ShipToCompanyAccountNumber string `json:"shipToCompanyAccountNumber,omitempty"`
	ShipToCompanyType CompanyTypeReference `json:"shipToCompanyType,omitempty"`
	ShipToContact ContactReference `json:"shipToContact,omitempty"`
	ShipToSite SiteReference `json:"shipToSite,omitempty"`
	ShipToTaxGroup string `json:"shipToTaxGroup,omitempty"`
	Site SiteReference `json:"site,omitempty"`
	StateTaxXref string `json:"stateTaxXref,omitempty"`
	TaxAgencyXref string `json:"taxAgencyXref,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	TaxGroupRate float64 `json:"taxGroupRate,omitempty"`
	TaxLevels []GLExportPurchaseTransactionTaxLevel `json:"taxLevels,omitempty"`
	Total float64 `json:"total,omitempty"`
	UseAvalaraTaxFlag bool `json:"useAvalaraTaxFlag,omitempty"`
	VendorAccountNumber string `json:"vendorAccountNumber,omitempty"`
	VendorInvoiceDate string `json:"vendorInvoiceDate,omitempty"`
	VendorInvoiceNumber string `json:"vendorInvoiceNumber,omitempty"`
	VendorNumber string `json:"vendorNumber,omitempty"`
}

type GLExportPurchaseTransactionDetail struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	CogsXref string `json:"cogsXref,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	CostAccountNumber string `json:"costAccountNumber,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DropShippedFlag bool `json:"dropShippedFlag,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	GlItemId string `json:"glItemId,omitempty"`
	GlTypeId string `json:"glTypeId,omitempty"`
	ID int `json:"id,omitempty"`
	InventoryAccountNumber string `json:"inventoryAccountNumber,omitempty"`
	InventoryXref string `json:"inventoryXref,omitempty"`
	Item IvItemReference `json:"item,omitempty"`
	ItemCost float64 `json:"itemCost,omitempty"`
	ItemDescription string `json:"itemDescription,omitempty"`
	ItemPrice float64 `json:"itemPrice,omitempty"`
	ItemTypeXref string `json:"itemTypeXref,omitempty"`
	LineNumber int `json:"lineNumber,omitempty"`
	LocationXref string `json:"locationXref,omitempty"`
	Memo string `json:"memo,omitempty"`
	PriceLevelXref string `json:"priceLevelXref,omitempty"`
	PurchaseHeaderTaxGroup string `json:"purchaseHeaderTaxGroup,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	SalesCode string `json:"salesCode,omitempty"`
	SalesDescription string `json:"salesDescription,omitempty"`
	SerialNumbers string `json:"serialNumbers,omitempty"`
	SerializedFlag bool `json:"serializedFlag,omitempty"`
	ShipmentMethod ShipmentMethodReference `json:"shipmentMethod,omitempty"`
	SubCategory ProductSubCategoryReference `json:"subCategory,omitempty"`
	TaxAgencyXref string `json:"taxAgencyXref,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	TaxCodeXref string `json:"taxCodeXref,omitempty"`
	TaxNote string `json:"taxNote,omitempty"`
	TaxRate float64 `json:"taxRate,omitempty"`
	Taxable bool `json:"taxable,omitempty"`
	Total float64 `json:"total,omitempty"`
	UnitOfMeasure UnitOfMeasureReference `json:"unitOfMeasure,omitempty"`
	UomScheduleXref string `json:"uomScheduleXref,omitempty"`
	VendorAccountNumber string `json:"vendorAccountNumber,omitempty"`
	VendorNumber string `json:"vendorNumber,omitempty"`
	WarehouseBin WarehouseBinReference `json:"warehouseBin,omitempty"`
	WarehouseSite SiteReference `json:"warehouseSite,omitempty"`
}

type GLExportPurchaseTransactionDetailTax struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	CogsXref string `json:"cogsXref,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	CostAccountNumber string `json:"costAccountNumber,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DropShippedFlag bool `json:"dropShippedFlag,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	GlItemId string `json:"glItemId,omitempty"`
	GlTypeId string `json:"glTypeId,omitempty"`
	ID int `json:"id,omitempty"`
	InventoryAccountNumber string `json:"inventoryAccountNumber,omitempty"`
	InventoryXref string `json:"inventoryXref,omitempty"`
	Item IvItemReference `json:"item,omitempty"`
	ItemCost float64 `json:"itemCost,omitempty"`
	ItemDescription string `json:"itemDescription,omitempty"`
	ItemPrice float64 `json:"itemPrice,omitempty"`
	ItemTypeXref string `json:"itemTypeXref,omitempty"`
	LineNumber int `json:"lineNumber,omitempty"`
	LocationXref string `json:"locationXref,omitempty"`
	Memo string `json:"memo,omitempty"`
	PriceLevelXref string `json:"priceLevelXref,omitempty"`
	PurchaseHeaderTaxGroup string `json:"purchaseHeaderTaxGroup,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	SalesCode string `json:"salesCode,omitempty"`
	SalesDescription string `json:"salesDescription,omitempty"`
	SerialNumbers string `json:"serialNumbers,omitempty"`
	SerializedFlag bool `json:"serializedFlag,omitempty"`
	ShipmentMethod ShipmentMethodReference `json:"shipmentMethod,omitempty"`
	SubCategory ProductSubCategoryReference `json:"subCategory,omitempty"`
	TaxAgencyXref string `json:"taxAgencyXref,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	TaxNote string `json:"taxNote,omitempty"`
	TaxRate float64 `json:"taxRate,omitempty"`
	TaxRatePercent float64 `json:"taxRatePercent,omitempty"`
	TaxableFlag bool `json:"taxableFlag,omitempty"`
	Total float64 `json:"total,omitempty"`
	UnitOfMeasure UnitOfMeasureReference `json:"unitOfMeasure,omitempty"`
	UomScheduleXref string `json:"uomScheduleXref,omitempty"`
	VendorAccountNumber string `json:"vendorAccountNumber,omitempty"`
	VendorNumber string `json:"vendorNumber,omitempty"`
	WarehouseBin WarehouseBinReference `json:"warehouseBin,omitempty"`
	WarehouseSite SiteReference `json:"warehouseSite,omitempty"`
}

type GLExportPurchaseTransactionTaxLevel struct {
	TaxCodeXref string `json:"taxCodeXref,omitempty"`
	TaxLevel int `json:"taxLevel,omitempty"`
}

type GLExportSettings struct {
}

type GLExportTransaction struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	AgreementPrePaymentFlag bool `json:"agreementPrePaymentFlag,omitempty"`
	Attention string `json:"attention,omitempty"`
	BillingTerms BillingTermsReference `json:"billingTerms,omitempty"`
	BillingTermsXref string `json:"billingTermsXref,omitempty"`
	BillingType string `json:"billingType,omitempty"`
	CityTax float64 `json:"cityTax,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	CompanyAccountNumber string `json:"companyAccountNumber,omitempty"`
	CompanyType CompanyTypeReference `json:"companyType,omitempty"`
	CountyTax float64 `json:"countyTax,omitempty"`
	CountyTaxXref string `json:"countyTaxXref,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	Description string `json:"description,omitempty"`
	Detail []GLExportTransactionDetail `json:"detail,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DocumentNumber string `json:"documentNumber,omitempty"`
	DocumentType string `json:"documentType,omitempty"`
	DueDate string `json:"dueDate,omitempty"`
	DueDays int `json:"dueDays,omitempty"`
	EmailDeliveryFlag bool `json:"emailDeliveryFlag,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	GlEntryIds string `json:"glEntryIds,omitempty"`
	GlTypeId string `json:"glTypeId,omitempty"`
	ID int `json:"id,omitempty"`
	Memo string `json:"memo,omitempty"`
	PiggyBackFlag bool `json:"piggyBackFlag,omitempty"`
	PrintDeliveryFlag bool `json:"printDeliveryFlag,omitempty"`
	Project ProjectReference `json:"project,omitempty"`
	PurchaseOrder PurchaseOrderReference `json:"purchaseOrder,omitempty"`
	SalesRepId string `json:"salesRepId,omitempty"`
	SalesRepName string `json:"salesRepName,omitempty"`
	SalesTax float64 `json:"salesTax,omitempty"`
	SalesTerritory string `json:"salesTerritory,omitempty"`
	SendAvalaraTaxFlag bool `json:"sendAvalaraTaxFlag,omitempty"`
	ShipContact string `json:"shipContact,omitempty"`
	ShipSite SiteReference `json:"shipSite,omitempty"`
	ShipToCompany CompanyReference `json:"shipToCompany,omitempty"`
	ShipToCompanyAccountNumber string `json:"shipToCompanyAccountNumber,omitempty"`
	ShipToCompanyType CompanyTypeReference `json:"shipToCompanyType,omitempty"`
	ShipToTaxId string `json:"shipToTaxId,omitempty"`
	Site SiteReference `json:"site,omitempty"`
	StateTax float64 `json:"stateTax,omitempty"`
	StateTaxXref string `json:"stateTaxXref,omitempty"`
	TaxAccountNumber string `json:"taxAccountNumber,omitempty"`
	TaxAgencyXref string `json:"taxAgencyXref,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	TaxDpAppliedFlag bool `json:"taxDpAppliedFlag,omitempty"`
	TaxGroupRate float64 `json:"taxGroupRate,omitempty"`
	TaxId string `json:"taxId,omitempty"`
	TaxLevels []GLExportTransactionTaxLevel `json:"taxLevels,omitempty"`
	Taxable bool `json:"taxable,omitempty"`
	TaxableAmount1 float64 `json:"taxableAmount1,omitempty"`
	TaxableAmount2 float64 `json:"taxableAmount2,omitempty"`
	TaxableAmount3 float64 `json:"taxableAmount3,omitempty"`
	TaxableAmount4 float64 `json:"taxableAmount4,omitempty"`
	TaxableAmount5 float64 `json:"taxableAmount5,omitempty"`
	TaxableTotal float64 `json:"taxableTotal,omitempty"`
	Total float64 `json:"total,omitempty"`
	UseAvalaraFlag bool `json:"useAvalaraFlag,omitempty"`
}

type GLExportTransactionDetail struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	CogsXref string `json:"cogsXref,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	CostAccountNumber string `json:"costAccountNumber,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentDate string `json:"documentDate,omitempty"`
	DocumentType string `json:"documentType,omitempty"`
	DropShippedFlag bool `json:"dropShippedFlag,omitempty"`
	GlClass string `json:"glClass,omitempty"`
	GlItemId string `json:"glItemId,omitempty"`
	GlTypeId string `json:"glTypeId,omitempty"`
	ID int `json:"id,omitempty"`
	InventoryAccountNumber string `json:"inventoryAccountNumber,omitempty"`
	InventoryXref string `json:"inventoryXref,omitempty"`
	InvoiceSummaryOption string `json:"invoiceSummaryOption,omitempty"`
	Item IvItemReference `json:"item,omitempty"`
	ItemCost float64 `json:"itemCost,omitempty"`
	ItemDescription string `json:"itemDescription,omitempty"`
	ItemPrice float64 `json:"itemPrice,omitempty"`
	ItemTaxableFlag bool `json:"itemTaxableFlag,omitempty"`
	ItemTypeXref string `json:"itemTypeXref,omitempty"`
	LocationXref string `json:"locationXref,omitempty"`
	Memo string `json:"memo,omitempty"`
	PriceLevelXref string `json:"priceLevelXref,omitempty"`
	Product ProductReference `json:"product,omitempty"`
	ProductAccountNumber string `json:"productAccountNumber,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	SalesCode string `json:"salesCode,omitempty"`
	SalesDescription string `json:"salesDescription,omitempty"`
	SerialNumbers string `json:"serialNumbers,omitempty"`
	SerializedFlag bool `json:"serializedFlag,omitempty"`
	ShipmentMethod ShipmentMethodReference `json:"shipmentMethod,omitempty"`
	SubCategory ProductSubCategoryReference `json:"subCategory,omitempty"`
	TaxAgencyXref string `json:"taxAgencyXref,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	TaxCodeXref string `json:"taxCodeXref,omitempty"`
	TaxLevels []GLExportTransactionDetailTaxLevel `json:"taxLevels,omitempty"`
	TaxNote string `json:"taxNote,omitempty"`
	TaxRate float64 `json:"taxRate,omitempty"`
	TaxRatePercent float64 `json:"taxRatePercent,omitempty"`
	Taxable2Flag bool `json:"taxable2Flag,omitempty"`
	Taxable3Flag bool `json:"taxable3Flag,omitempty"`
	Taxable4Flag bool `json:"taxable4Flag,omitempty"`
	Taxable5Flag bool `json:"taxable5Flag,omitempty"`
	TaxableFlag bool `json:"taxableFlag,omitempty"`
	TimeEntry TimeEntryReference `json:"timeEntry,omitempty"`
	Total float64 `json:"total,omitempty"`
	UnitOfMeasure UnitOfMeasureReference `json:"unitOfMeasure,omitempty"`
	UomScheduleXref string `json:"uomScheduleXref,omitempty"`
	WarehouseBin WarehouseBinReference `json:"warehouseBin,omitempty"`
	WarehouseSite SiteReference `json:"warehouseSite,omitempty"`
}

type GLExportTransactionDetailTaxLevel struct {
	TaxLevel int `json:"taxLevel,omitempty"`
	TaxableFlag bool `json:"taxableFlag,omitempty"`
}

type GLExportTransactionTaxLevel struct {
	TaxAmount float64 `json:"taxAmount,omitempty"`
	TaxCodeXref string `json:"taxCodeXref,omitempty"`
	TaxLevel int `json:"taxLevel,omitempty"`
	TaxableAmount float64 `json:"taxableAmount,omitempty"`
}

type GLExportVendor struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	BillingTerms BillingTermsReference `json:"billingTerms,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	Contact ContactReference `json:"contact,omitempty"`
	DueDays int `json:"dueDays,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Site SiteReference `json:"site,omitempty"`
	TaxCode TaxCodeReference `json:"taxCode,omitempty"`
	Vendor CompanyReference `json:"vendor,omitempty"`
	VendorNumber string `json:"vendorNumber,omitempty"`
}

type GenericBoardTeamReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	IsProjectTeamFlag bool `json:"isProjectTeamFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type GenericIdIdentifierReference struct {
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type GenericNameIdDTO struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag string `json:"tag,omitempty"`
}

type GoogleEmailSetupReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GroupReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type HolidayListReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type HttpContent struct {
	Headers []interface{} `json:"headers,omitempty"`
}

type HttpMethod struct {
	Delete *HttpMethod `json:"delete,omitempty"`
	Get *HttpMethod `json:"get,omitempty"`
	Head *HttpMethod `json:"head,omitempty"`
	Method string `json:"method,omitempty"`
	Options *HttpMethod `json:"options,omitempty"`
	Post *HttpMethod `json:"post,omitempty"`
	Put *HttpMethod `json:"put,omitempty"`
	Trace *HttpMethod `json:"trace,omitempty"`
}

type HttpRequestMessage struct {
	Content HttpContent `json:"content,omitempty"`
	Headers []interface{} `json:"headers,omitempty"`
	Method *HttpMethod `json:"method,omitempty"`
	Properties interface{} `json:"properties,omitempty"`
	RequestUri string `json:"requestUri,omitempty"`
	Version Version `json:"version,omitempty"`
}

type IRestIdentifiedItem struct {
	ID int `json:"id,omitempty"`
}

type ImapSetupReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type InOutTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type InclusiveRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type IntegratorLoginReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type IntegratorTagCollection struct {
	Tags []string `json:"tags,omitempty"`
}

type InvoiceGroupingReference struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description,omitempty"`
	GroupParentChildAdditionsFlag bool `json:"groupParentChildAdditionsFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ShowPriceFlag bool `json:"showPriceFlag,omitempty"`
	ShowSubItemsFlag bool `json:"showSubItemsFlag,omitempty"`
}

type InvoiceReference struct {
	Info interface{} `json:"_info,omitempty"`
	ApplyToType string `json:"applyToType,omitempty"`
	BillingType string `json:"billingType,omitempty"`
	ChargeFirmFlag bool `json:"chargeFirmFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	InvoiceDate string `json:"invoiceDate,omitempty"`
}

type InvoiceTemplateDetailReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type InvoiceTemplateReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type IvItemReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	SerializedFlag bool `json:"serializedFlag,omitempty"`
}

type KBCategoryReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type KPICategoryReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type KPIReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type LdapConfigurationReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Server string `json:"server,omitempty"`
}

type LicenseBit struct {
	ActiveFlag bool `json:"activeFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type LocaleReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type LostRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type ManagedDevicesIntegrationReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ManagementSolutionReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	SetupName string `json:"setupName,omitempty"`
}

type ManufacturerReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MappedRecordReference struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MappedTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MarketDescriptionReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MemberDeactivationCompanyTeam struct {
	Count int `json:"count,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ReAssignToContact ContactReference `json:"reAssignToContact,omitempty"`
	ReAssignToMember MemberReference `json:"reAssignToMember,omitempty"`
}

type MemberDeactivationItem struct {
	Count int `json:"count,omitempty"`
	ReAssignToMember MemberReference `json:"reAssignToMember,omitempty"`
}

type MemberDeactivationStatusWorkflow struct {
	Count int `json:"count,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ReAssignToMember MemberReference `json:"reAssignToMember,omitempty"`
}

type MemberOffice365 struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MemberReference struct {
	Info interface{} `json:"_info,omitempty"`
	DailyCapacity float64 `json:"dailyCapacity,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type MemberSsoSettingsReference struct {
	Info interface{} `json:"_info,omitempty"`
	Email string `json:"email,omitempty"`
	ID int `json:"id,omitempty"`
	SsoUserId string `json:"ssoUserId,omitempty"`
	UserName string `json:"userName,omitempty"`
}

type MemberTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MenuLocationReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type NoteTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type NotificationRecipientReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type NotifyTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type Office365EmailSetupReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OnHandSerialNumberReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
}

type OpenRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type OpportunityPriorityReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OpportunityProbabilityReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OpportunityRatingReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OpportunitySalesRoleReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OpportunityStageReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OpportunityStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OpportunityTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OrderStatusEmailTemplateReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OrderStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Other1RevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type Other2RevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type OwnerLevelReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type OwnershipTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PageValues struct {
	Page int `json:"page,omitempty"`
	PageId int `json:"pageId,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type PaymentMethodReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PhaseStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PortalConfigurationPaymentProcessorReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PortalConfigurationReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PricingScheduleReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PriorityReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Level string `json:"level,omitempty"`
	Name string `json:"name,omitempty"`
	Sort int `json:"sort,omitempty"`
}

type ProductCategoryReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductDemand struct {
	Cost float64 `json:"cost,omitempty"`
	ProductRecId int `json:"productRecId,omitempty"`
	Quantity int `json:"quantity,omitempty"`
}

type ProductItemReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductRecurring struct {
	AgreementType AgreementTypeReference `json:"agreementType,omitempty"`
	BillCycle BillingCycleReference `json:"billCycle,omitempty"`
	BillCycleId int `json:"billCycleId,omitempty"`
	CycleType string `json:"cycleType,omitempty"`
	Cycles int `json:"cycles,omitempty"`
	EndDate string `json:"endDate,omitempty"`
	RecurringCost float64 `json:"recurringCost,omitempty"`
	RecurringRevenue float64 `json:"recurringRevenue,omitempty"`
	StartDate string `json:"startDate,omitempty"`
}

type ProductRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type ProductSubCategoryReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProjectBoardKanbanStatus struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Order int `json:"order,omitempty"`
	SrStatusId int `json:"srStatusId,omitempty"`
}

type ProjectBoardReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProjectPhaseReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProjectReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProjectRoleReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type ProjectStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProjectTemplatePhaseReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProjectTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProjectWorkplanProjectPhase struct {
	Info interface{} `json:"_info,omitempty"`
	ActualHours float64 `json:"actualHours,omitempty"`
	BillPhaseSeparately bool `json:"billPhaseSeparately,omitempty"`
	BillableHours float64 `json:"billableHours,omitempty"`
	BudgetHours float64 `json:"budgetHours,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
	Description string `json:"description,omitempty"`
	EndDate string `json:"endDate,omitempty"`
	ID int `json:"id,omitempty"`
	MarkAsMilestoneFlag bool `json:"markAsMilestoneFlag,omitempty"`
	Notes string `json:"notes,omitempty"`
	ParentPhase ProjectPhaseReference `json:"parentPhase,omitempty"`
	ProjectId int `json:"projectId,omitempty"`
	ScheduledDuration int `json:"scheduled_Duration,omitempty"`
	ScheduledEnd string `json:"scheduled_End,omitempty"`
	ScheduledHours float64 `json:"scheduled_Hours,omitempty"`
	ScheduledStart string `json:"scheduled_Start,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	Status PhaseStatusReference `json:"status,omitempty"`
	WbsCode string `json:"wbsCode,omitempty"`
}

type PurchaseOrderLineItemReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type PurchaseOrderReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PurchaseOrderStatusEmailTemplateReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PurchaseOrderStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RelationshipReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ReminderReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ReportCardReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ResultInfo struct {
	Data IRestIdentifiedItem `json:"data,omitempty"`
	Error ErrorResponseMessage `json:"error,omitempty"`
	OriginalIndex int `json:"originalIndex,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
	Success bool `json:"success,omitempty"`
}

type RmaActionReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RmaDispositionReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RmaStatusEmailTemplateReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RmaStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SLAReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SalesOrderReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type SalesTeamReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type ScheduleSpanReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type ScheduleStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ScheduleTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type SecurityRoleReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceCodeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceEmailTemplateReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type ServiceItemReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceLocationReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type ServiceSignoffReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceSourceReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceStatusReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Sort int `json:"sort,omitempty"`
}

type ServiceSubTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceSurveyQuestionOption struct {
	Caption string `json:"caption,omitempty"`
	IncludeFlag bool `json:"includeFlag,omitempty"`
	Points int `json:"points,omitempty"`
}

type ServiceSurveyReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceTeamReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ServiceTemplateReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type ServiceTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ShipmentMethodReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SicCodeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SiteReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SkillCategoryReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SkillReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type StateReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type StatusIndicatorReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type StructureReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SurveyQuestionReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
}

type SurveyReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SurveyResultDetail struct {
	Answer interface{} `json:"answer,omitempty"`
	QuestionId int `json:"questionId,omitempty"`
}

type SystemDepartmentReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type SystemLocationReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SystemMenuEntryReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TaxCodeLevelReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TaxCodeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TeamRoleReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TemplatePhase struct {
	Info interface{} `json:"_info,omitempty"`
	BillPhaseSeparately bool `json:"billPhaseSeparately,omitempty"`
	Description string `json:"description,omitempty"`
	ID int `json:"id,omitempty"`
	MarkAsMilestoneFlag bool `json:"markAsMilestoneFlag,omitempty"`
	Notes string `json:"notes,omitempty"`
	ParentPhase ProjectTemplatePhaseReference `json:"parentPhase,omitempty"`
	TemplateId int `json:"templateId,omitempty"`
	WbsCode string `json:"wbsCode,omitempty"`
}

type TicketReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type TimeAccrualReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TimePeriodSetupReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
}

type TimeRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type TimeSheetReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TimeZoneReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TimeZoneSetupReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TrackReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UnitOfMeasureReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserDefinedFieldOption struct {
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	OptionValue string `json:"optionValue,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type UserDefinedFieldReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserDefinedFieldValueModel struct {
	Filtered bool `json:"filtered,omitempty"`
	RowNum int `json:"rowNum,omitempty"`
	SkipLocationAndBillingUnit bool `json:"skipLocationAndBillingUnit,omitempty"`
	UserDefinedFieldRecId int `json:"userDefinedFieldRecId,omitempty"`
	Value string `json:"value,omitempty"`
}

type ValidationError struct {
	Code string `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
	Field string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Resource string `json:"resource,omitempty"`
}

type Version struct {
	Build int `json:"build,omitempty"`
	Major int `json:"major,omitempty"`
	MajorRevision int `json:"majorRevision,omitempty"`
	Minor int `json:"minor,omitempty"`
	MinorRevision int `json:"minorRevision,omitempty"`
	Revision int `json:"revision,omitempty"`
}

type WarehouseBinReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type WarehouseReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	LockedFlag bool `json:"lockedFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type WisePayBatchPayment struct {
	Amount float64 `json:"amount,omitempty"`
	WisePayHref string `json:"wisePayHref,omitempty"`
}

type WisePayFeeInvoice struct {
	Amount float64 `json:"amount,omitempty"`
	ID int `json:"id,omitempty"`
	InvoiceHref string `json:"invoiceHref,omitempty"`
	InvoiceNumber string `json:"invoiceNumber,omitempty"`
}

type WisePayPayment struct {
	BatchPayment WisePayBatchPayment `json:"batchPayment,omitempty"`
	FeeInvoice WisePayFeeInvoice `json:"feeInvoice,omitempty"`
	PaymentDateUtc string `json:"paymentDateUtc,omitempty"`
	WisePay string `json:"wisePayReference,omitempty"`
}

type WonRevenueReference struct {
	Info interface{} `json:"_info,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	ID int `json:"id,omitempty"`
	Margin float64 `json:"margin,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
}

type WorkRoleReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type WorkTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	UtilizationFlag bool `json:"utilizationFlag,omitempty"`
}

type WorkflowTableTypeReference struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

