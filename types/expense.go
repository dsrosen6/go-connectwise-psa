package types

import (
	"time"
)

type Classification struct {
	Info interface{} `json:"_info,omitempty"`
	CompanyFlag bool `json:"companyFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	EmployeeFlag bool `json:"employeeFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Multiplier float64 `json:"multiplier,omitempty"`
	Name string `json:"name,omitempty"`
}

type ExpenseEntry struct {
	Info interface{} `json:"_info,omitempty"`
	Agreement AgreementReference `json:"agreement,omitempty"`
	AgreementAmount float64 `json:"agreementAmount,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	BillAmount float64 `json:"billAmount,omitempty"`
	BillableOption string `json:"billableOption,omitempty"`
	BusinessUnitId int `json:"businessUnitId,omitempty"`
	ChargeToId int `json:"chargeToId,omitempty"`
	ChargeToType string `json:"chargeToType,omitempty"`
	Classification ClassificationReference `json:"classification,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	Currency CurrencyReference `json:"currency,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
	Date time.Time `json:"date"`
	ExpenseReport ExpenseReportReference `json:"expenseReport,omitempty"`
	ID int `json:"id,omitempty"`
	Invoice InvoiceReference `json:"invoice,omitempty"`
	InvoiceAmount float64 `json:"invoiceAmount,omitempty"`
	LocationId int `json:"locationId,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	MobileGuid string `json:"mobileGuid,omitempty"`
	Notes string `json:"notes,omitempty"`
	OdometerEnd float64 `json:"odometerEnd,omitempty"`
	OdometerStart float64 `json:"odometerStart,omitempty"`
	PaymentMethod PaymentMethodReference `json:"paymentMethod,omitempty"`
	Phase ProjectPhaseReference `json:"phase,omitempty"`
	Project ProjectReference `json:"project,omitempty"`
	Status string `json:"status,omitempty"`
	Taxes []ExpenseTax `json:"taxes,omitempty"`
	Ticket TicketReference `json:"ticket,omitempty"`
	Type ExpenseTypeReference `json:"type"`
}

type ExpenseEntryAudit struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Message string `json:"message,omitempty"`
	NewValue string `json:"newValue,omitempty"`
	OldValue string `json:"oldValue,omitempty"`
	Source string `json:"source,omitempty"`
	Type string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type ExpenseReport struct {
	Info interface{} `json:"_info,omitempty"`
	DateEnd string `json:"dateEnd,omitempty"`
	DateStart string `json:"dateStart,omitempty"`
	DueDate string `json:"dueDate,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Period int `json:"period,omitempty"`
	Status string `json:"status,omitempty"`
	Total float64 `json:"total,omitempty"`
	Year int `json:"year,omitempty"`
}

type ExpenseReportAudit struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Message string `json:"message,omitempty"`
	NewValue string `json:"newValue,omitempty"`
	OldValue string `json:"oldValue,omitempty"`
	Source string `json:"source,omitempty"`
	Type string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type ExpenseReportTierUpdate struct {
	ApprovalType string `json:"approvalType,omitempty"`
	ID int `json:"id,omitempty"`
}

type ExpenseTaxTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Inactive bool `json:"inactive,omitempty"`
	Name string `json:"name,omitempty"`
}

type ExpenseType struct {
	Info interface{} `json:"_info,omitempty"`
	AdvancedAmountFlag bool `json:"advancedAmountFlag,omitempty"`
	AmountCaption string `json:"amountCaption"`
	BillExpenses string `json:"billExpenses,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	IntegrationXRef string `json:"integrationXRef,omitempty"`
	InvoiceMarkupAmount float64 `json:"invoiceMarkupAmount,omitempty"`
	InvoiceMarkupOption string `json:"invoiceMarkupOption,omitempty"`
	MaxAmount float64 `json:"maxAmount,omitempty"`
	MileageFlag bool `json:"mileageFlag,omitempty"`
	Name string `json:"name"`
	QuantityFlag bool `json:"quantityFlag,omitempty"`
	ReimbursementRate float64 `json:"reimbursementRate,omitempty"`
}

type ExpenseTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	AmountCaption string `json:"amountCaption,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	MileageFlag bool `json:"mileageFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type PaymentType struct {
	Info interface{} `json:"_info,omitempty"`
	Classification ClassificationReference `json:"classification"`
	CompanyFlag bool `json:"companyFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
}

type PaymentTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

