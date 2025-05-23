package types

import (
	"time"
)

type AllowedFileType struct {
	Info interface{} `json:"_info,omitempty"`
	FileType string `json:"fileType"`
	ID int `json:"id,omitempty"`
	SizeLimit int `json:"sizeLimit,omitempty"`
}

type AllowedOrigin struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description"`
	ID int `json:"id,omitempty"`
	LastUpdateUtc time.Time `json:"lastUpdateUtc,omitempty"`
	Origin string `json:"origin"`
	UpdatedBy string `json:"updatedBy,omitempty"`
}

type ApiMember struct {
	Info interface{} `json:"_info,omitempty"`
	BlockCostFlag bool `json:"blockCostFlag,omitempty"`
	BlockPriceFlag bool `json:"blockPriceFlag,omitempty"`
	DefaultDepartment SystemDepartmentReference `json:"defaultDepartment,omitempty"`
	DefaultLocation SystemLocationReference `json:"defaultLocation,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	ExcludedServiceBoardIds []int `json:"excludedServiceBoardIds,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier"`
	InactiveDate time.Time `json:"inactiveDate,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
	Notes string `json:"notes,omitempty"`
	SalesDefaultLocation SystemLocationReference `json:"salesDefaultLocation,omitempty"`
	SecurityLocation SystemLocationReference `json:"securityLocation,omitempty"`
	SecurityRole SecurityRoleReference `json:"securityRole,omitempty"`
	ServiceDefaultBoard BoardReference `json:"serviceDefaultBoard,omitempty"`
	StructureLevel StructureReference `json:"structureLevel,omitempty"`
	TimeZone TimeZoneSetupReference `json:"timeZone,omitempty"`
}

type AuditTrailEntry struct {
	AuditSource string `json:"auditSource,omitempty"`
	AuditSubType string `json:"auditSubType,omitempty"`
	AuditType string `json:"auditType,omitempty"`
	EnteredBy string `json:"enteredBy,omitempty"`
	EnteredDate string `json:"enteredDate,omitempty"`
	Text string `json:"text,omitempty"`
}

type AuthAnvil struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	ServerLocationUrl string `json:"serverLocationUrl"`
	SiteId int `json:"siteId,omitempty"`
}

type AutoSyncTime struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	SyncTime string `json:"syncTime"`
	TimeZone TimeZoneSetupReference `json:"timeZone"`
}

type BillableOptionsInfo struct {
	BillableFlag bool `json:"billableFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	EnumId string `json:"enumId,omitempty"`
	ExpenseFlag bool `json:"expenseFlag,omitempty"`
	ID int `json:"id,omitempty"`
	IncludeNoDefaultFlag bool `json:"includeNoDefaultFlag,omitempty"`
	InvoiceFlag bool `json:"invoiceFlag,omitempty"`
	Name string `json:"name,omitempty"`
	OptionId string `json:"optionId,omitempty"`
	ProductFlag bool `json:"productFlag,omitempty"`
	TimeFlag bool `json:"timeFlag,omitempty"`
}

type BundleRequestsCollection struct {
	Requests []BundleRequest `json:"requests"`
}

type BundleResultsCollection struct {
	Info interface{} `json:"_info,omitempty"`
	Results []BundleResult `json:"results,omitempty"`
}

type CallbackEntry struct {
	Info interface{} `json:"_info,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	Description string `json:"description,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	IsSelfSuppressedFlag bool `json:"isSelfSuppressedFlag,omitempty"`
	IsSoapCallbackFlag bool `json:"isSoapCallbackFlag,omitempty"`
	Level string `json:"level"`
	MemberId int `json:"memberId,omitempty"`
	ObjectId int `json:"objectId,omitempty"`
	PayloadVersion string `json:"payloadVersion,omitempty"`
	Type string `json:"type"`
	URL string `json:"url"`
}

type Certification struct {
	Info interface{} `json:"_info,omitempty"`
	Company CompanyReference `json:"company"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
}

type ConnectWiseHostedScreen struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ScreenId string `json:"screenId,omitempty"`
}

type ConnectWiseHostedSetup struct {
	Info interface{} `json:"_info,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	DateCreated string `json:"dateCreated,omitempty"`
	Description string `json:"description"`
	DisabledFlag bool `json:"disabledFlag,omitempty"`
	ID int `json:"id,omitempty"`
	LocationIds []int `json:"locationIds,omitempty"`
	LocationsEnabledFlag bool `json:"locationsEnabledFlag,omitempty"`
	Origin string `json:"origin,omitempty"`
	PodHeight int `json:"podHeight,omitempty"`
	ScreenId int `json:"screenId,omitempty"`
	ToolbarButtonDialogHeight int `json:"toolbarButtonDialogHeight,omitempty"`
	ToolbarButtonDialogWidth int `json:"toolbarButtonDialogWidth,omitempty"`
	ToolbarButtonIconDocumentId int `json:"toolbarButtonIconDocumentId,omitempty"`
	ToolbarButtonText string `json:"toolbarButtonText,omitempty"`
	ToolbarButtonToolTip string `json:"toolbarButtonToolTip,omitempty"`
	Type string `json:"type,omitempty"`
	URL string `json:"url"`
}

type CorporateStructure struct {
	Info interface{} `json:"_info,omitempty"`
	BaseCurrency CurrencyReference `json:"baseCurrency"`
	ChiefOperatingOfficer MemberReference `json:"chiefOperatingOfficer,omitempty"`
	Controller MemberReference `json:"controller,omitempty"`
	Dispatcher MemberReference `json:"dispatcher,omitempty"`
	DutyManager MemberReference `json:"dutyManager,omitempty"`
	FiscalYearStart string `json:"fiscalYearStart,omitempty"`
	GroupCaption string `json:"groupCaption"`
	ID int `json:"id,omitempty"`
	Level1Name string `json:"level1Name,omitempty"`
	Level2Name string `json:"level2Name,omitempty"`
	Level3Name string `json:"level3Name,omitempty"`
	Level4Name string `json:"level4Name,omitempty"`
	Level5Name string `json:"level5Name,omitempty"`
	LevelCount string `json:"levelCount,omitempty"`
	LocationCaption string `json:"locationCaption"`
	President MemberReference `json:"president,omitempty"`
	ServiceManager MemberReference `json:"serviceManager,omitempty"`
}

type CorporateStructureInfo struct {
	Info interface{} `json:"_info,omitempty"`
	BaseCurrency CurrencyReference `json:"baseCurrency,omitempty"`
	GroupCaption string `json:"groupCaption,omitempty"`
	ID int `json:"id,omitempty"`
	LocationCaption string `json:"locationCaption,omitempty"`
}

type CorporateStructureLevel struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Count struct {
	Count int `json:"count,omitempty"`
}

type Crm struct {
	Info interface{} `json:"_info,omitempty"`
	AccountManagerRole TeamRoleReference `json:"accountManagerRole"`
	CompanyIdGenerationFlag bool `json:"companyIdGenerationFlag,omitempty"`
	CompanyListCount int `json:"companyListCount,omitempty"`
	DefaultYear bool `json:"defaultYear,omitempty"`
	ExcludeSpacesFlag bool `json:"excludeSpacesFlag,omitempty"`
	Field10Caption string `json:"field10Caption,omitempty"`
	Field1Caption string `json:"field1Caption,omitempty"`
	Field2Caption string `json:"field2Caption,omitempty"`
	Field3Caption string `json:"field3Caption,omitempty"`
	Field4Caption string `json:"field4Caption,omitempty"`
	Field5Caption string `json:"field5Caption,omitempty"`
	Field6Caption string `json:"field6Caption,omitempty"`
	Field7Caption string `json:"field7Caption,omitempty"`
	Field8Caption string `json:"field8Caption,omitempty"`
	Field9Caption string `json:"field9Caption,omitempty"`
	ID int `json:"id,omitempty"`
	LockProbabilityFlag bool `json:"lockProbabilityFlag,omitempty"`
	Other1Caption string `json:"other1Caption,omitempty"`
	Other2Caption string `json:"other2Caption,omitempty"`
	PrimaryRepCaption string `json:"primaryRepCaption,omitempty"`
	SalesRepRole TeamRoleReference `json:"salesRepRole"`
	SecondaryRepCaption string `json:"secondaryRepCaption,omitempty"`
	TechnicalContactRole TeamRoleReference `json:"technicalContactRole"`
}

type CrmInfo struct {
	Info interface{} `json:"_info,omitempty"`
	AccountManagerRole TeamRoleReference `json:"accountManagerRole,omitempty"`
	Field10Caption string `json:"field10Caption,omitempty"`
	Field1Caption string `json:"field1Caption,omitempty"`
	Field2Caption string `json:"field2Caption,omitempty"`
	Field3Caption string `json:"field3Caption,omitempty"`
	Field4Caption string `json:"field4Caption,omitempty"`
	Field5Caption string `json:"field5Caption,omitempty"`
	Field6Caption string `json:"field6Caption,omitempty"`
	Field7Caption string `json:"field7Caption,omitempty"`
	Field8Caption string `json:"field8Caption,omitempty"`
	Field9Caption string `json:"field9Caption,omitempty"`
	ID int `json:"id,omitempty"`
	PrimaryRepCaption string `json:"primaryRepCaption,omitempty"`
	SalesRepRole TeamRoleReference `json:"salesRepRole,omitempty"`
	SecondaryRepCaption string `json:"secondaryRepCaption,omitempty"`
	TechnicalContactRole TeamRoleReference `json:"technicalContactRole,omitempty"`
}

type CustomReport struct {
	Info interface{} `json:"_info,omitempty"`
	AgreementFlag bool `json:"agreementFlag,omitempty"`
	AgreementOverride string `json:"agreementOverride,omitempty"`
	AgreementParamId int `json:"agreementParamId,omitempty"`
	AgreementTypeFlag bool `json:"agreementTypeFlag,omitempty"`
	AgreementTypeOverride string `json:"agreementTypeOverride,omitempty"`
	AgreementTypeParamId int `json:"agreementTypeParamId,omitempty"`
	CompanyFlag bool `json:"companyFlag,omitempty"`
	CompanyOverride string `json:"companyOverride,omitempty"`
	CompanyParamId int `json:"companyParamId,omitempty"`
	DepartmentDefaultFlag bool `json:"departmentDefaultFlag,omitempty"`
	DepartmentFlag bool `json:"departmentFlag,omitempty"`
	DepartmentOverride string `json:"departmentOverride,omitempty"`
	DepartmentParamId int `json:"departmentParamId,omitempty"`
	Description string `json:"description"`
	EndDateFlag bool `json:"endDateFlag,omitempty"`
	EndDateOverride string `json:"endDateOverride,omitempty"`
	EndDateParamId int `json:"endDateParamId,omitempty"`
	GeneratedFlag bool `json:"generatedFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InvoiceFlag bool `json:"invoiceFlag,omitempty"`
	InvoiceOverride string `json:"invoiceOverride,omitempty"`
	InvoiceParamId int `json:"invoiceParamId,omitempty"`
	LocationDefaultFlag bool `json:"locationDefaultFlag,omitempty"`
	LocationFlag bool `json:"locationFlag,omitempty"`
	LocationOverride string `json:"locationOverride,omitempty"`
	LocationParamId int `json:"locationParamId,omitempty"`
	MarketingCampaignFlag bool `json:"marketingCampaignFlag,omitempty"`
	MarketingCampaignOverride string `json:"marketingCampaignOverride,omitempty"`
	MarketingCampaignParamId int `json:"marketingCampaignParamId,omitempty"`
	MemberFlag bool `json:"memberFlag,omitempty"`
	MemberOverride string `json:"memberOverride,omitempty"`
	MemberParamId int `json:"memberParamId,omitempty"`
	Module string `json:"module,omitempty"`
	Name string `json:"name"`
	OppTypeFlag bool `json:"oppTypeFlag,omitempty"`
	OppTypeOverride string `json:"oppTypeOverride,omitempty"`
	OppTypeParamId int `json:"oppTypeParamId,omitempty"`
	OpportunityFlag bool `json:"opportunityFlag,omitempty"`
	OpportunityOverride string `json:"opportunityOverride,omitempty"`
	OpportunityParamId int `json:"opportunityParamId,omitempty"`
	ParameterNameSeparator string `json:"parameterNameSeparator,omitempty"`
	ParameterPrefix string `json:"parameterPrefix,omitempty"`
	ParameterSeparator string `json:"parameterSeparator,omitempty"`
	ParameterSuffix string `json:"parameterSuffix,omitempty"`
	ProjectFlag bool `json:"projectFlag,omitempty"`
	ProjectOverride string `json:"projectOverride,omitempty"`
	ProjectParamId int `json:"projectParamId,omitempty"`
	ProjectTypeFlag bool `json:"projectTypeFlag,omitempty"`
	ProjectTypeOverride string `json:"projectTypeOverride,omitempty"`
	ProjectTypeParamId int `json:"projectTypeParamId,omitempty"`
	ReportLink string `json:"reportLink"`
	ServiceBoardDefaultFlag bool `json:"serviceBoardDefaultFlag,omitempty"`
	ServiceBoardFlag bool `json:"serviceBoardFlag,omitempty"`
	ServiceBoardOverride string `json:"serviceBoardOverride,omitempty"`
	ServiceBoardParamId int `json:"serviceBoardParamId,omitempty"`
	ServiceStatusFlag bool `json:"serviceStatusFlag,omitempty"`
	ServiceStatusOverride string `json:"serviceStatusOverride,omitempty"`
	ServiceStatusParamId int `json:"serviceStatusParamId,omitempty"`
	ServiceTypeFlag bool `json:"serviceTypeFlag,omitempty"`
	ServiceTypeOverride string `json:"serviceTypeOverride,omitempty"`
	ServiceTypeParamId int `json:"serviceTypeParamId,omitempty"`
	StartDateFlag bool `json:"startDateFlag,omitempty"`
	StartDateOverride string `json:"startDateOverride,omitempty"`
	StartDateParamId int `json:"startDateParamId,omitempty"`
	TerritoryDefaultFlag bool `json:"territoryDefaultFlag,omitempty"`
	TerritoryFlag bool `json:"territoryFlag,omitempty"`
	TerritoryOverride string `json:"territoryOverride,omitempty"`
	TerritoryParamId int `json:"territoryParamId,omitempty"`
	WorkRoleFlag bool `json:"workRoleFlag,omitempty"`
	WorkRoleOverride string `json:"workRoleOverride,omitempty"`
	WorkRoleParamId int `json:"workRoleParamId,omitempty"`
	WorkTypeFlag bool `json:"workTypeFlag,omitempty"`
	WorkTypeOverride string `json:"workTypeOverride,omitempty"`
	WorkTypeParamId int `json:"workTypeParamId,omitempty"`
}

type CustomReportParameter struct {
	Info interface{} `json:"_info,omitempty"`
	CaptionName string `json:"captionName,omitempty"`
	CustomReport CustomReportReference `json:"customReport,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CwTimeZone struct {
	Info interface{} `json:"_info,omitempty"`
	DaylightSavingsFlag bool `json:"daylightSavingsFlag,omitempty"`
	EndDate string `json:"endDate,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Offset float64 `json:"offset,omitempty"`
	StartDate string `json:"startDate,omitempty"`
}

type Department struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier"`
	Name string `json:"name"`
}

type DepartmentInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DepartmentLocation struct {
	Info interface{} `json:"_info,omitempty"`
	AddAllLocations bool `json:"addAllLocations,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	DepartmentManager MemberReference `json:"departmentManager,omitempty"`
	Dispatch MemberReference `json:"dispatch,omitempty"`
	DutyManager MemberReference `json:"dutyManager,omitempty"`
	ID int `json:"id,omitempty"`
	LdapConfig LdapConfigurationReference `json:"ldapConfig,omitempty"`
	Location SystemLocationReference `json:"location"`
	RemoveAllLocations bool `json:"removeAllLocations,omitempty"`
	ServiceManager MemberReference `json:"serviceManager,omitempty"`
}

type DepartmentLocationInfo struct {
	Info interface{} `json:"_info,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
}

type DirectionalSyncInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DocumentFormData struct {
	File string `json:"file,omitempty"`
	IsAvatar bool `json:"isAvatar,omitempty"`
	PrivateFlag bool `json:"privateFlag,omitempty"`
	ReadOnlyFlay bool `json:"readOnlyFlay,omitempty"`
	RecordId int `json:"recordId,omitempty"`
	RecordType string `json:"recordType,omitempty"`
	Title string `json:"title,omitempty"`
	URL string `json:"url,omitempty"`
}

type DocumentInfo struct {
	Info interface{} `json:"_info,omitempty"`
	CreatedOnDate string `json:"createdOnDate,omitempty"`
	DocumentType DocumentTypeReference `json:"documentType,omitempty"`
	FileName string `json:"fileName,omitempty"`
	Guid string `json:"guid,omitempty"`
	HtmlTemplateFlag bool `json:"htmlTemplateFlag,omitempty"`
	ID int `json:"id,omitempty"`
	ImageFlag bool `json:"imageFlag,omitempty"`
	LinkFlag bool `json:"linkFlag,omitempty"`
	Owner string `json:"owner,omitempty"`
	PublicFlag bool `json:"publicFlag,omitempty"`
	ReadOnlyFlag bool `json:"readOnlyFlag,omitempty"`
	ServerFileName string `json:"serverFileName,omitempty"`
	Size int `json:"size,omitempty"`
	Title string `json:"title,omitempty"`
	UrlFlag bool `json:"urlFlag,omitempty"`
}

type DocumentSetup struct {
	Info interface{} `json:"_info,omitempty"`
	DocPath string `json:"docPath,omitempty"`
	ID int `json:"id,omitempty"`
	IsPublicFlag bool `json:"isPublicFlag,omitempty"`
	TemplateOutputPath string `json:"templateOutputPath,omitempty"`
	TemplatePath string `json:"templatePath,omitempty"`
	UploadAsLinkFlag bool `json:"uploadAsLinkFlag,omitempty"`
}

type DocumentType struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description,omitempty"`
	FileExtension string `json:"fileExtension,omitempty"`
	Icon string `json:"icon,omitempty"`
	ID int `json:"id,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
}

type EPayConfiguration struct {
	Info interface{} `json:"_info,omitempty"`
	Currency CurrencyReference `json:"currency"`
	EncryptionKey string `json:"encryptionKey,omitempty"`
	ID int `json:"id,omitempty"`
	InitializationVector string `json:"initializationVector,omitempty"`
	Location SystemLocationReference `json:"location"`
	StoreIdentifier string `json:"storeIdentifier"`
	URL string `json:"url"`
}

type EmailConnector struct {
	Info interface{} `json:"_info,omitempty"`
	AddCcFlag bool `json:"addCcFlag,omitempty"`
	Asio365EmailSetup Office365EmailSetupReference `json:"asio365EmailSetup,omitempty"`
	BccEmailTo string `json:"bccEmailTo,omitempty"`
	CreateContactFlag bool `json:"createContactFlag,omitempty"`
	DefaultCompany CompanyReference `json:"defaultCompany"`
	DefaultMember MemberReference `json:"defaultMember,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	DiscardDuplicatesFlag bool `json:"discardDuplicatesFlag,omitempty"`
	EmailErrorsTo string `json:"emailErrorsTo"`
	EmailNotifyFrom string `json:"emailNotifyFrom,omitempty"`
	EmailServerType string `json:"emailServerType,omitempty"`
	GoogleEmailSetup GoogleEmailSetupReference `json:"googleEmailSetup,omitempty"`
	ID int `json:"id,omitempty"`
	ImapSetup ImapSetupReference `json:"imapSetup,omitempty"`
	InboundTicketMailboxId string `json:"inboundTicketMailboxId,omitempty"`
	ItemOverride ServiceItemReference `json:"itemOverride,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	NeverRespondFlag bool `json:"neverRespondFlag,omitempty"`
	NoResponseFlag bool `json:"noResponseFlag,omitempty"`
	Office365EmailSetup Office365EmailSetupReference `json:"office365EmailSetup,omitempty"`
	PostRepliesToTicketFlag bool `json:"postRepliesToTicketFlag,omitempty"`
	PriorityOverride PriorityReference `json:"priorityOverride,omitempty"`
	ResponseEmailForExisting string `json:"responseEmailForExisting,omitempty"`
	ResponseEmailForNew string `json:"responseEmailForNew,omitempty"`
	ServiceBoard BoardReference `json:"serviceBoard"`
	SetEmailToDefaultContactFlag bool `json:"setEmailToDefaultContactFlag,omitempty"`
	SourceOverride ServiceSourceReference `json:"sourceOverride,omitempty"`
	StatusOverride ServiceStatusReference `json:"statusOverride,omitempty"`
	SubTypeOverride ServiceSubTypeReference `json:"subTypeOverride,omitempty"`
	TypeOverride ServiceTypeReference `json:"typeOverride,omitempty"`
	UseEmailMessageIdFlag bool `json:"useEmailMessageIdFlag,omitempty"`
}

type EmailConnectorInfo struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultCompany CompanyReference `json:"defaultCompany,omitempty"`
	ID int `json:"id,omitempty"`
	ImapSetup ImapSetupReference `json:"imapSetup,omitempty"`
}

type EmailConnectorParsingRule struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	ParsingStyle EmailConnectorParsingStyleReference `json:"parsingStyle,omitempty"`
	ParsingVariable EmailConnectorParsingVariableReference `json:"parsingVariable"`
	Priority int `json:"priority,omitempty"`
	SearchTerm string `json:"searchTerm"`
	ServiceBoard BoardReference `json:"serviceBoard,omitempty"`
	ServiceItem ServiceItemReference `json:"serviceItem,omitempty"`
	ServicePriority PriorityReference `json:"servicePriority,omitempty"`
	ServiceStatus ServiceStatusReference `json:"serviceStatus,omitempty"`
	ServiceSubType ServiceSubTypeReference `json:"serviceSubType,omitempty"`
	ServiceType ServiceTypeReference `json:"serviceType,omitempty"`
}

type EmailConnectorParsingStyle struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	ParseRule string `json:"parseRule"`
	ParsingType EmailConnectorParsingTypeReference `json:"parsingType"`
	Priority int `json:"priority,omitempty"`
}

type EmailExclusion struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description"`
	ID int `json:"id,omitempty"`
}

type EmailToken struct {
	AddressFlag bool `json:"addressFlag,omitempty"`
	AgreementFlag bool `json:"agreementFlag,omitempty"`
	CompanyFlag bool `json:"companyFlag,omitempty"`
	ConfigFlag bool `json:"configFlag,omitempty"`
	ContactFlag bool `json:"contactFlag,omitempty"`
	Description string `json:"description,omitempty"`
	ID int `json:"id,omitempty"`
	InvoiceFlag bool `json:"invoiceFlag,omitempty"`
	PortalPasswordFlag bool `json:"portalPasswordFlag,omitempty"`
	PurchaseOrderFlag bool `json:"purchaseOrderFlag,omitempty"`
	PurchaseOrderStatusFlag bool `json:"purchaseOrderStatusFlag,omitempty"`
	RmaFlag bool `json:"rmaFlag,omitempty"`
	SalesFlag bool `json:"salesFlag,omitempty"`
	ServiceFlag bool `json:"serviceFlag,omitempty"`
	Token string `json:"token,omitempty"`
	TracksFlag bool `json:"tracksFlag,omitempty"`
	WorkflowFlag bool `json:"workflowFlag,omitempty"`
}

type Experiment struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description,omitempty"`
	ExperimentId string `json:"experimentId,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	MemberInactiveFlag bool `json:"memberInactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
	Properties string `json:"properties,omitempty"`
}

type FileUploadSetting struct {
	Info interface{} `json:"_info,omitempty"`
	GlobalFileSizeLimit int `json:"globalFileSizeLimit,omitempty"`
	ID int `json:"id,omitempty"`
	RestrictFileTypesFlag bool `json:"restrictFileTypesFlag,omitempty"`
}

type GoogleEmailSetup struct {
	Info interface{} `json:"_info,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	EmailConnector EmailConnectorReference `json:"emailConnector,omitempty"`
	FailedFolder string `json:"failedFolder"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	InboxFolder string `json:"inboxFolder"`
	Name string `json:"name"`
	PrivateKey string `json:"privateKey,omitempty"`
	ProcessedFolder string `json:"processedFolder"`
	Username string `json:"username"`
}

type IdCollection struct {
	Ids []int `json:"ids,omitempty"`
}

type Imap struct {
	Info interface{} `json:"_info,omitempty"`
	EmailConnector EmailConnectorReference `json:"emailConnector,omitempty"`
	FailedFolder string `json:"failedFolder"`
	ID int `json:"id,omitempty"`
	ImapName string `json:"imapName"`
	Name string `json:"name"`
	Password string `json:"password,omitempty"`
	Port int `json:"port,omitempty"`
	ProcessedName string `json:"processedName"`
	Server string `json:"server"`
	SslFlag bool `json:"sslFlag,omitempty"`
	UserName string `json:"userName"`
}

type ImapInfo struct {
	Info interface{} `json:"_info,omitempty"`
	EmailConnector EmailConnectorReference `json:"emailConnector,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ImportMassMaintenance struct {
	DeletedCompanyCount int `json:"deletedCompanyCount,omitempty"`
	DeletedContactCount int `json:"deletedContactCount,omitempty"`
	Message string `json:"message,omitempty"`
	Success bool `json:"success,omitempty"`
}

type InOutBoard struct {
	Info interface{} `json:"_info,omitempty"`
	AdditionalInfo string `json:"additionalInfo,omitempty"`
	DateBack time.Time `json:"dateBack"`
	ID int `json:"id,omitempty"`
	InOutType InOutTypeReference `json:"inOutType"`
	Member MemberReference `json:"member"`
}

type InOutType struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description"`
	ID int `json:"id,omitempty"`
}

type InOutTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description,omitempty"`
	ID int `json:"id,omitempty"`
}

type Info struct {
	CloudRegion string `json:"cloudRegion,omitempty"`
	IsCloud bool `json:"isCloud,omitempty"`
	LicenseBits []LicenseBit `json:"licenseBits,omitempty"`
	MaxWorkFlowRecordsAllowed int `json:"maxWorkFlowRecordsAllowed,omitempty"`
	ServerTimeZone string `json:"serverTimeZone,omitempty"`
	Version string `json:"version,omitempty"`
}

type IntegratorLogin struct {
	Info interface{} `json:"_info,omitempty"`
	ActivityApiFlag bool `json:"activityApiFlag,omitempty"`
	ActivityCallbackUrl string `json:"activityCallbackUrl,omitempty"`
	ActivityLegacyCallbackFlag bool `json:"activityLegacyCallbackFlag,omitempty"`
	AgreementApiFlag bool `json:"agreementApiFlag,omitempty"`
	AgreementCallbackLegacyFlag bool `json:"agreementCallbackLegacyFlag,omitempty"`
	AgreementCallbackUrl string `json:"agreementCallbackUrl,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	CanAccessAllApisFlag bool `json:"canAccessAllApisFlag,omitempty"`
	CanAccessAllRecordsFlag bool `json:"canAccessAllRecordsFlag,omitempty"`
	CompanyApiFlag bool `json:"companyApiFlag,omitempty"`
	CompanyCallbackUrl string `json:"companyCallbackUrl,omitempty"`
	CompanyLegacyCallbackFlag bool `json:"companyLegacyCallbackFlag,omitempty"`
	ConfigurationApiFlag bool `json:"configurationApiFlag,omitempty"`
	ConfigurationAutoChildFlag bool `json:"configurationAutoChildFlag,omitempty"`
	ConfigurationCallbackUrl string `json:"configurationCallbackUrl,omitempty"`
	ConfigurationChildlingFlag bool `json:"configurationChildlingFlag,omitempty"`
	ConfigurationLegacyCallbackFlag bool `json:"configurationLegacyCallbackFlag,omitempty"`
	ContactApiFlag bool `json:"contactApiFlag,omitempty"`
	ContactCallbackUrl string `json:"contactCallbackUrl,omitempty"`
	ContactLegacyCallbackFlag bool `json:"contactLegacyCallbackFlag,omitempty"`
	DateInactivated time.Time `json:"dateInactivated,omitempty"`
	DocumentApiFlag bool `json:"documentApiFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactivatedBy MemberReference `json:"inactivatedBy,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	InvoiceApiFlag bool `json:"invoiceApiFlag,omitempty"`
	ManagedServicesApiFlag bool `json:"managedServicesApiFlag,omitempty"`
	ManagedServicesAutoChildFlag bool `json:"managedServicesAutoChildFlag,omitempty"`
	ManagedServicesChildingFlag bool `json:"managedServicesChildingFlag,omitempty"`
	MarketingApiFlag bool `json:"marketingApiFlag,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	MemberApiFlag bool `json:"memberApiFlag,omitempty"`
	OpportunityApiFlag bool `json:"opportunityApiFlag,omitempty"`
	OpportunityCallbackUrl string `json:"opportunityCallbackUrl,omitempty"`
	OpportunityConversionApiFlag bool `json:"opportunityConversionApiFlag,omitempty"`
	OpportunityLegacyCallbackFlag bool `json:"opportunityLegacyCallbackFlag,omitempty"`
	Password string `json:"password,omitempty"`
	ProductApiFlag bool `json:"productApiFlag,omitempty"`
	ProductCallbackUrl string `json:"productCallbackUrl,omitempty"`
	ProductLegacyCallbackFlag bool `json:"productLegacyCallbackFlag,omitempty"`
	ProjectApiFlag bool `json:"projectApiFlag,omitempty"`
	ProjectCallbackUrl string `json:"projectCallbackUrl,omitempty"`
	ProjectLegacyCallbackFlag bool `json:"projectLegacyCallbackFlag,omitempty"`
	PurchasingApiFlag bool `json:"purchasingApiFlag,omitempty"`
	PurchasingCallbackUrl string `json:"purchasingCallbackUrl,omitempty"`
	PurchasingLegacyCallbackFlag bool `json:"purchasingLegacyCallbackFlag,omitempty"`
	ReportingApiFlag bool `json:"reportingApiFlag,omitempty"`
	ScheduleApiFlag bool `json:"scheduleApiFlag,omitempty"`
	ScheduleCallbackUrl string `json:"scheduleCallbackUrl,omitempty"`
	ScheduleLegacyCallbackFlag bool `json:"scheduleLegacyCallbackFlag,omitempty"`
	ServiceBoardCallbackUrl string `json:"serviceBoardCallbackUrl,omitempty"`
	ServiceBoardLegacyCallbackFlag bool `json:"serviceBoardLegacyCallbackFlag,omitempty"`
	ServiceTicketApiFlag bool `json:"serviceTicketApiFlag,omitempty"`
	SystemApiFlag bool `json:"systemApiFlag,omitempty"`
	TimeEntryApiFlag bool `json:"timeEntryApiFlag,omitempty"`
	TimeEntryCallbackUrl string `json:"timeEntryCallbackUrl,omitempty"`
	TimeEntryLegacyCallbackFlag bool `json:"timeEntryLegacyCallbackFlag,omitempty"`
	Username string `json:"username"`
}

type IntegratorTag struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Text string `json:"text"`
}

type KPI struct {
	Category KPICategoryReference `json:"category,omitempty"`
	DateFilter string `json:"dateFilter,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type KPICategory struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type LdapConfiguration struct {
	Info interface{} `json:"_info,omitempty"`
	Domain string `json:"domain"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Server string `json:"server"`
}

type LdapConfigurationInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type LdapConfigurationTestLink struct {
	Server string `json:"server,omitempty"`
}

type Link struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	ScreenLink string `json:"screenLink,omitempty"`
	TableReferenceId int `json:"tableReferenceId,omitempty"`
	URL string `json:"url,omitempty"`
}

type LinkInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ScreenLink string `json:"screenLink,omitempty"`
}

type LinkResolveUrlInfo struct {
	ReferenceId int `json:"referenceId,omitempty"`
	URL string `json:"url,omitempty"`
}

type LocaleInfo struct {
	ID int `json:"id,omitempty"`
	LocaleCode string `json:"localeCode,omitempty"`
	Name string `json:"name,omitempty"`
}

type Location struct {
	Info interface{} `json:"_info,omitempty"`
	Calendar CalendarReference `json:"calendar,omitempty"`
	ClientFlag bool `json:"clientFlag,omitempty"`
	DepartmentIds []int `json:"departmentIds,omitempty"`
	ID int `json:"id,omitempty"`
	LocationFlag bool `json:"locationFlag,omitempty"`
	Manager MemberReference `json:"manager,omitempty"`
	Name string `json:"name"`
	OverrideAddressLine1 string `json:"overrideAddressLine1,omitempty"`
	OverrideAddressLine2 string `json:"overrideAddressLine2,omitempty"`
	OverrideCity string `json:"overrideCity,omitempty"`
	OverrideCountry CountryReference `json:"overrideCountry,omitempty"`
	OverrideFaxNumber string `json:"overrideFaxNumber,omitempty"`
	OverridePhoneNumber string `json:"overridePhoneNumber,omitempty"`
	OverrideState string `json:"overrideState,omitempty"`
	OverrideZip string `json:"overrideZip,omitempty"`
	OwaUrl string `json:"owaUrl,omitempty"`
	OwnerLevelId int `json:"ownerLevelId,omitempty"`
	PayrollXref string `json:"payrollXref,omitempty"`
	ReportsTo SystemLocationReference `json:"reportsTo,omitempty"`
	SalesRep string `json:"salesRep,omitempty"`
	StructureLevel CorporateStructureLevelReference `json:"structureLevel"`
	TimeZoneSetup TimeZoneSetupReference `json:"timeZoneSetup,omitempty"`
	WorkRoleIds []int `json:"workRoleIds,omitempty"`
}

type LocationDepartment struct {
	Info interface{} `json:"_info,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
}

type LocationInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	LocationFlag bool `json:"location_flag,omitempty"`
	Name string `json:"name,omitempty"`
	StructureLevel CorporateStructureLevelReference `json:"structureLevel,omitempty"`
}

type LocationWorkRole struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	WorkRole WorkRoleReference `json:"workRole,omitempty"`
	WorkRoleInactiveFlag bool `json:"workRoleInactiveFlag,omitempty"`
}

type M365ContactSyncInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type M365ContactSyncMonitoring struct {
	Info interface{} `json:"_info,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	ID int `json:"id,omitempty"`
	MonitoringTypeId int `json:"monitoringTypeId,omitempty"`
	ServiceBoardId int `json:"serviceBoardId,omitempty"`
	ServiceBoardStatusId int `json:"serviceBoardStatusId,omitempty"`
}

type ManagedDeviceAccount struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	ManagedDevicesIntegration ManagedDevicesIntegrationReference `json:"managedDevicesIntegration,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}

type ManagementNetworkSecurity struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Password string `json:"password,omitempty"`
	Site string `json:"site"`
	Username string `json:"username,omitempty"`
}

type MarketplaceImport struct {
	ID int `json:"id,omitempty"`
	MarketplaceImportType string `json:"marketplaceImportType,omitempty"`
	MarketplaceObject []interface{} `json:"marketplaceObject,omitempty"`
	RequiredFields []string `json:"requiredFields,omitempty"`
}

type Member struct {
	Info interface{} `json:"_info,omitempty"`
	AdminFlag bool `json:"adminFlag,omitempty"`
	AgreementInvoicingDisplayOptions string `json:"agreementInvoicingDisplayOptions,omitempty"`
	AllowExpensesEnteredAgainstCompaniesFlag bool `json:"allowExpensesEnteredAgainstCompaniesFlag,omitempty"`
	AllowInCellEntryOnTimeSheet bool `json:"allowInCellEntryOnTimeSheet,omitempty"`
	AuthenticationServiceType string `json:"authenticationServiceType,omitempty"`
	AutoPopupQuickNotesWithStopwatch bool `json:"autoPopupQuickNotesWithStopwatch,omitempty"`
	AutoStartStopwatch bool `json:"autoStartStopwatch,omitempty"`
	BillableForecast float64 `json:"billableForecast,omitempty"`
	Calendar CalendarReference `json:"calendar,omitempty"`
	CalendarSyncIntegrationFlag bool `json:"calendarSyncIntegrationFlag,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	CompanyActivityTabFormat string `json:"companyActivityTabFormat,omitempty"`
	CopyColumnLayoutsAndFilters bool `json:"copyColumnLayoutsAndFilters,omitempty"`
	CopyPodLayouts bool `json:"copyPodLayouts,omitempty"`
	CopySharedDefaultViews bool `json:"copySharedDefaultViews,omitempty"`
	Country CountryReference `json:"country,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
	DailyCapacity float64 `json:"dailyCapacity,omitempty"`
	DaysTolerance int `json:"daysTolerance,omitempty"`
	DefaultDepartment SystemDepartmentReference `json:"defaultDepartment,omitempty"`
	DefaultEmail string `json:"defaultEmail,omitempty"`
	DefaultLocation SystemLocationReference `json:"defaultLocation,omitempty"`
	DefaultPhone string `json:"defaultPhone,omitempty"`
	DirectionalSync DirectionalSyncReference `json:"directionalSync,omitempty"`
	DisableOnlineFlag bool `json:"disableOnlineFlag,omitempty"`
	EmployeeIdentifer string `json:"employeeIdentifer,omitempty"`
	EnableLdapAuthenticationFlag bool `json:"enableLdapAuthenticationFlag,omitempty"`
	EnableMobileFlag bool `json:"enableMobileFlag,omitempty"`
	EnableMobileGpsFlag bool `json:"enableMobileGpsFlag,omitempty"`
	EnterTimeAgainstCompanyFlag bool `json:"enterTimeAgainstCompanyFlag,omitempty"`
	ExcludedProjectBoardIds []int `json:"excludedProjectBoardIds,omitempty"`
	ExcludedServiceBoardIds []int `json:"excludedServiceBoardIds,omitempty"`
	ExpenseApprover MemberReference `json:"expenseApprover,omitempty"`
	FirstName string `json:"firstName"`
	FromMemberRecId int `json:"fromMemberRecId,omitempty"`
	FromMemberTemplateRecId int `json:"fromMemberTemplateRecId,omitempty"`
	GlobalSearchDefaultSort string `json:"globalSearchDefaultSort,omitempty"`
	GlobalSearchDefaultTicketFilter string `json:"globalSearchDefaultTicketFilter,omitempty"`
	HideMemberInDispatchPortalFlag bool `json:"hideMemberInDispatchPortalFlag,omitempty"`
	HireDate time.Time `json:"hireDate"`
	HomeEmail string `json:"homeEmail,omitempty"`
	HomeExtension string `json:"homeExtension,omitempty"`
	HomePhone string `json:"homePhone,omitempty"`
	HourlyCost float64 `json:"hourlyCost,omitempty"`
	HourlyRate float64 `json:"hourlyRate,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier"`
	InactiveDate time.Time `json:"inactiveDate,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	IncludeInUtilizationReportingFlag bool `json:"includeInUtilizationReportingFlag,omitempty"`
	InvoiceScreenDefaultTabFormat string `json:"invoiceScreenDefaultTabFormat,omitempty"`
	InvoiceTimeTabFormat string `json:"invoiceTimeTabFormat,omitempty"`
	InvoicingDisplayOptions string `json:"invoicingDisplayOptions,omitempty"`
	LastLogin string `json:"lastLogin,omitempty"`
	LastName string `json:"lastName"`
	LdapConfiguration LdapConfigurationReference `json:"ldapConfiguration,omitempty"`
	LdapUserName string `json:"ldapUserName,omitempty"`
	LicenseClass string `json:"licenseClass,omitempty"`
	MapiName string `json:"mapiName,omitempty"`
	MemberPersonas []int `json:"memberPersonas,omitempty"`
	MiddleInitial string `json:"middleInitial,omitempty"`
	MinimumHours float64 `json:"minimumHours,omitempty"`
	MobileEmail string `json:"mobileEmail,omitempty"`
	MobileExtension string `json:"mobileExtension,omitempty"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	Notes string `json:"notes,omitempty"`
	Office365 MemberOffice365 `json:"office365,omitempty"`
	OfficeEmail string `json:"officeEmail,omitempty"`
	OfficeExtension string `json:"officeExtension,omitempty"`
	OfficePhone string `json:"officePhone,omitempty"`
	PartnerPortalFlag bool `json:"partnerPortalFlag,omitempty"`
	Password string `json:"password,omitempty"`
	PhoneIntegrationType string `json:"phoneIntegrationType,omitempty"`
	PhoneSource string `json:"phoneSource,omitempty"`
	Photo DocumentReference `json:"photo,omitempty"`
	PrimaryEmail string `json:"primaryEmail,omitempty"`
	ProjectDefaultBoard ProjectBoardReference `json:"projectDefaultBoard,omitempty"`
	ProjectDefaultDepartment SystemDepartmentReference `json:"projectDefaultDepartment,omitempty"`
	ProjectDefaultLocation SystemLocationReference `json:"projectDefaultLocation,omitempty"`
	ReportCard ReportCardReference `json:"reportCard,omitempty"`
	ReportsTo MemberReference `json:"reportsTo,omitempty"`
	RequireExpenseEntryFlag bool `json:"requireExpenseEntryFlag,omitempty"`
	RequireStartAndEndTimeOnTimeEntryFlag bool `json:"requireStartAndEndTimeOnTimeEntryFlag,omitempty"`
	RequireTimeSheetEntryFlag bool `json:"requireTimeSheetEntryFlag,omitempty"`
	RestrictDefaultSalesTerritoryFlag bool `json:"restrictDefaultSalesTerritoryFlag,omitempty"`
	RestrictDefaultWarehouseBinFlag bool `json:"restrictDefaultWarehouseBinFlag,omitempty"`
	RestrictDefaultWarehouseFlag bool `json:"restrictDefaultWarehouseFlag,omitempty"`
	RestrictDepartmentFlag bool `json:"restrictDepartmentFlag,omitempty"`
	RestrictLocationFlag bool `json:"restrictLocationFlag,omitempty"`
	RestrictProjectDefaultDepartmentFlag bool `json:"restrictProjectDefaultDepartmentFlag,omitempty"`
	RestrictProjectDefaultLocationFlag bool `json:"restrictProjectDefaultLocationFlag,omitempty"`
	RestrictScheduleFlag bool `json:"restrictScheduleFlag,omitempty"`
	RestrictServiceDefaultDepartmentFlag bool `json:"restrictServiceDefaultDepartmentFlag,omitempty"`
	RestrictServiceDefaultLocationFlag bool `json:"restrictServiceDefaultLocationFlag,omitempty"`
	SalesDefaultLocation SystemLocationReference `json:"salesDefaultLocation,omitempty"`
	ScheduleCapacity float64 `json:"scheduleCapacity,omitempty"`
	ScheduleDefaultDepartment SystemDepartmentReference `json:"scheduleDefaultDepartment,omitempty"`
	ScheduleDefaultLocation SystemLocationReference `json:"scheduleDefaultLocation,omitempty"`
	SecurityLocation SystemLocationReference `json:"securityLocation,omitempty"`
	SecurityRole SecurityRoleReference `json:"securityRole"`
	ServiceBoardTeamIds []int `json:"serviceBoardTeamIds,omitempty"`
	ServiceDefaultBoard BoardReference `json:"serviceDefaultBoard,omitempty"`
	ServiceDefaultDepartment SystemDepartmentReference `json:"serviceDefaultDepartment,omitempty"`
	ServiceDefaultLocation SystemLocationReference `json:"serviceDefaultLocation,omitempty"`
	ServiceLocation ServiceLocationReference `json:"serviceLocation,omitempty"`
	Signature string `json:"signature,omitempty"`
	SsoSettings MemberSsoSettingsReference `json:"ssoSettings,omitempty"`
	StructureLevel StructureReference `json:"structureLevel,omitempty"`
	StsUserAdminUrl string `json:"stsUserAdminUrl,omitempty"`
	Teams []int `json:"teams,omitempty"`
	TimeApprover MemberReference `json:"timeApprover,omitempty"`
	TimeReminderEmailFlag bool `json:"timeReminderEmailFlag,omitempty"`
	TimeSheetStartDate string `json:"timeSheetStartDate,omitempty"`
	TimeZone TimeZoneSetupReference `json:"timeZone,omitempty"`
	TimebasedOneTimePasswordActivated bool `json:"timebasedOneTimePasswordActivated,omitempty"`
	Title string `json:"title,omitempty"`
	ToastNotificationFlag bool `json:"toastNotificationFlag,omitempty"`
	Token string `json:"token,omitempty"`
	Type MemberTypeReference `json:"type,omitempty"`
	UseBrowserLanguageFlag bool `json:"useBrowserLanguageFlag,omitempty"`
	VendorNumber string `json:"vendorNumber,omitempty"`
	Warehouse WarehouseReference `json:"warehouse,omitempty"`
	WarehouseBin WarehouseBinReference `json:"warehouseBin,omitempty"`
	WorkRole WorkRoleReference `json:"workRole,omitempty"`
	WorkType WorkTypeReference `json:"workType,omitempty"`
}

type MemberAccrual struct {
	Info interface{} `json:"_info,omitempty"`
	AccrualType string `json:"accrualType,omitempty"`
	Hours float64 `json:"hours,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Reason string `json:"reason"`
	Year int `json:"year,omitempty"`
}

type MemberCertification struct {
	Info interface{} `json:"_info,omitempty"`
	Certification CertificationReference `json:"certification"`
	CertificationNumber string `json:"certificationNumber,omitempty"`
	Company CompanyReference `json:"company,omitempty"`
	DateExpires time.Time `json:"dateExpires,omitempty"`
	DateReceived time.Time `json:"dateReceived,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Notes string `json:"notes,omitempty"`
	PercentComplete int `json:"percentComplete,omitempty"`
}

type MemberDeactivation struct {
	Activity MemberDeactivationItem `json:"activity,omitempty"`
	CompanyTeam []MemberDeactivationCompanyTeam `json:"companyTeam,omitempty"`
	DeleteOpenTimeSheetsFlag bool `json:"deleteOpenTimeSheetsFlag,omitempty"`
	DepartmentManager MemberDeactivationItem `json:"departmentManager,omitempty"`
	DispatchMember MemberDeactivationItem `json:"dispatchMember,omitempty"`
	DutyManager MemberDeactivationItem `json:"dutyManager,omitempty"`
	KnowledgeBaseArticle MemberDeactivationItem `json:"knowledgeBaseArticle,omitempty"`
	MyCompanyCOO MemberDeactivationItem `json:"myCompanyCOO,omitempty"`
	MyCompanyController MemberDeactivationItem `json:"myCompanyController,omitempty"`
	MyCompanyDispatch MemberDeactivationItem `json:"myCompanyDispatch,omitempty"`
	MyCompanyDutyManagerRole MemberDeactivationItem `json:"myCompanyDutyManagerRole,omitempty"`
	MyCompanyPresident MemberDeactivationItem `json:"myCompanyPresident,omitempty"`
	MyCompanyServiceManager MemberDeactivationItem `json:"myCompanyServiceManager,omitempty"`
	Opportunity MemberDeactivationItem `json:"opportunity,omitempty"`
	ProjectExpenseApprover MemberDeactivationItem `json:"projectExpenseApprover,omitempty"`
	ProjectManager MemberDeactivationItem `json:"projectManager,omitempty"`
	ProjectTimeApprover MemberDeactivationItem `json:"projectTimeApprover,omitempty"`
	SalesTeam MemberDeactivationItem `json:"salesTeam,omitempty"`
	SendFromEmailNotify MemberDeactivationItem `json:"sendFromEmailNotify,omitempty"`
	ServiceManager MemberDeactivationItem `json:"serviceManager,omitempty"`
	ServiceStatusWorkflow []MemberDeactivationStatusWorkflow `json:"serviceStatusWorkflow,omitempty"`
	ServiceTeam MemberDeactivationItem `json:"serviceTeam,omitempty"`
	TicketTemplate MemberDeactivationItem `json:"ticketTemplate,omitempty"`
	WorkflowEmail MemberDeactivationItem `json:"workflowEmail,omitempty"`
}

type MemberDelegation struct {
	Info interface{} `json:"_info,omitempty"`
	DateEnd time.Time `json:"dateEnd"`
	DateStart time.Time `json:"dateStart"`
	DelegatedTo MemberReference `json:"delegatedTo"`
	DelegationType string `json:"delegationType,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
}

type MemberForCalSync struct {
	CalendarSyncIntegrationFlag bool `json:"calendarSyncIntegrationFlag,omitempty"`
	ID int `json:"id,omitempty"`
	MapiName string `json:"mapiName,omitempty"`
	MemberId string `json:"memberId,omitempty"`
	Office365Id string `json:"office365Id,omitempty"`
}

type MemberInfo struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultEmail string `json:"defaultEmail,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	FullName string `json:"fullName,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	LastName string `json:"lastName,omitempty"`
	LicenseClass string `json:"licenseClass,omitempty"`
	MiddleInitial string `json:"middleInitial,omitempty"`
	Photo DocumentReference `json:"photo,omitempty"`
}

type MemberLinkSsoUser struct {
	SsoUserId string `json:"ssoUserId,omitempty"`
}

type MemberNotificationSetting struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	NotificationTrigger string `json:"notificationTrigger,omitempty"`
	NotificationType string `json:"notificationType,omitempty"`
}

type MemberPersona struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	JobRolePercentage int `json:"jobRolePercentage,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Name string `json:"name"`
	PersonaId int `json:"personaId"`
}

type MemberSkill struct {
	Info interface{} `json:"_info,omitempty"`
	CertifiedFlag bool `json:"certifiedFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	Notes string `json:"notes,omitempty"`
	Skill SkillReference `json:"skill"`
	SkillLevel string `json:"skillLevel,omitempty"`
	YearsExperience int `json:"yearsExperience,omitempty"`
}

type MemberSsoToken struct {
	Token string `json:"token,omitempty"`
}

type MemberTemplate struct {
	Info interface{} `json:"_info,omitempty"`
	AdminFlag bool `json:"adminFlag,omitempty"`
	AgreementInvoicingDisplayOptions string `json:"agreementInvoicingDisplayOptions,omitempty"`
	AllowExpensesEnteredAgainstCompaniesFlag bool `json:"allowExpensesEnteredAgainstCompaniesFlag,omitempty"`
	AllowInCellEntryOnTimeSheet bool `json:"allowInCellEntryOnTimeSheet,omitempty"`
	AutoPopupQuickNotesWithStopwatch bool `json:"autoPopupQuickNotesWithStopwatch,omitempty"`
	AutoStartStopwatch bool `json:"autoStartStopwatch,omitempty"`
	BillableForecast float64 `json:"billableForecast,omitempty"`
	Calendar CalendarReference `json:"calendar,omitempty"`
	CompanyActivityTabFormat string `json:"companyActivityTabFormat,omitempty"`
	CopyColumnLayoutsAndFilters bool `json:"copyColumnLayoutsAndFilters,omitempty"`
	CopyPodLayouts bool `json:"copyPodLayouts,omitempty"`
	CopySharedDefaultViews bool `json:"copySharedDefaultViews,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
	DailyCapacity float64 `json:"dailyCapacity,omitempty"`
	DaysTolerance int `json:"daysTolerance,omitempty"`
	DefaultDepartment SystemDepartmentReference `json:"defaultDepartment,omitempty"`
	DefaultLocation SystemLocationReference `json:"defaultLocation,omitempty"`
	EnableMobileFlag bool `json:"enableMobileFlag,omitempty"`
	EnterTimeAgainstCompanyFlag bool `json:"enterTimeAgainstCompanyFlag,omitempty"`
	ExcludedProjectBoardIds []int `json:"excludedProjectBoardIds,omitempty"`
	ExcludedServiceBoardIds []int `json:"excludedServiceBoardIds,omitempty"`
	ExpenseApprover MemberReference `json:"expenseApprover,omitempty"`
	FromMemberRecId int `json:"fromMemberRecId,omitempty"`
	FromMemberTemplateRecId int `json:"fromMemberTemplateRecId,omitempty"`
	GlobalSearchDefaultSort string `json:"globalSearchDefaultSort,omitempty"`
	GlobalSearchDefaultTicketFilter string `json:"globalSearchDefaultTicketFilter,omitempty"`
	HideMemberInDispatchPortalFlag bool `json:"hideMemberInDispatchPortalFlag,omitempty"`
	HourlyCost float64 `json:"hourlyCost,omitempty"`
	HourlyRate float64 `json:"hourlyRate,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier"`
	IncludeInUtilizationReportingFlag bool `json:"includeInUtilizationReportingFlag,omitempty"`
	InvoiceScreenDefaultTabFormat string `json:"invoiceScreenDefaultTabFormat,omitempty"`
	InvoiceTimeTabFormat string `json:"invoiceTimeTabFormat,omitempty"`
	InvoicingDisplayOptions string `json:"invoicingDisplayOptions,omitempty"`
	MemberPersonas []int `json:"memberPersonas,omitempty"`
	MinimumHours float64 `json:"minimumHours,omitempty"`
	PartnerPortalFlag bool `json:"partnerPortalFlag,omitempty"`
	PhoneSource string `json:"phoneSource,omitempty"`
	ProjectDefaultBoard ProjectBoardReference `json:"projectDefaultBoard,omitempty"`
	ProjectDefaultDepartment SystemDepartmentReference `json:"projectDefaultDepartment,omitempty"`
	ProjectDefaultLocation SystemLocationReference `json:"projectDefaultLocation,omitempty"`
	ReportCard ReportCardReference `json:"reportCard,omitempty"`
	ReportsTo MemberReference `json:"reportsTo,omitempty"`
	RequireExpenseEntryFlag bool `json:"requireExpenseEntryFlag,omitempty"`
	RequireStartAndEndTimeOnTimeEntryFlag bool `json:"requireStartAndEndTimeOnTimeEntryFlag,omitempty"`
	RequireTimeSheetEntryFlag bool `json:"requireTimeSheetEntryFlag,omitempty"`
	RestrictDefaultSalesTerritoryFlag bool `json:"restrictDefaultSalesTerritoryFlag,omitempty"`
	RestrictDefaultWarehouseBinFlag bool `json:"restrictDefaultWarehouseBinFlag,omitempty"`
	RestrictDefaultWarehouseFlag bool `json:"restrictDefaultWarehouseFlag,omitempty"`
	RestrictDepartmentFlag bool `json:"restrictDepartmentFlag,omitempty"`
	RestrictLocationFlag bool `json:"restrictLocationFlag,omitempty"`
	RestrictProjectDefaultDepartmentFlag bool `json:"restrictProjectDefaultDepartmentFlag,omitempty"`
	RestrictProjectDefaultLocationFlag bool `json:"restrictProjectDefaultLocationFlag,omitempty"`
	RestrictScheduleFlag bool `json:"restrictScheduleFlag,omitempty"`
	RestrictServiceDefaultDepartmentFlag bool `json:"restrictServiceDefaultDepartmentFlag,omitempty"`
	RestrictServiceDefaultLocationFlag bool `json:"restrictServiceDefaultLocationFlag,omitempty"`
	SalesDefaultLocation SystemLocationReference `json:"salesDefaultLocation,omitempty"`
	ScheduleCapacity float64 `json:"scheduleCapacity,omitempty"`
	ScheduleDefaultDepartment SystemDepartmentReference `json:"scheduleDefaultDepartment,omitempty"`
	ScheduleDefaultLocation SystemLocationReference `json:"scheduleDefaultLocation,omitempty"`
	SecurityLocation SystemLocationReference `json:"securityLocation,omitempty"`
	ServiceBoardTeamIds []int `json:"serviceBoardTeamIds,omitempty"`
	ServiceDefaultBoard BoardReference `json:"serviceDefaultBoard,omitempty"`
	ServiceDefaultDepartment SystemDepartmentReference `json:"serviceDefaultDepartment,omitempty"`
	ServiceDefaultLocation SystemLocationReference `json:"serviceDefaultLocation,omitempty"`
	ServiceLocation ServiceLocationReference `json:"serviceLocation,omitempty"`
	StructureLevel StructureReference `json:"structureLevel,omitempty"`
	StsUserAdminUrl string `json:"stsUserAdminUrl,omitempty"`
	Teams []int `json:"teams,omitempty"`
	TemplateDescription string `json:"templateDescription,omitempty"`
	TimeApprover MemberReference `json:"timeApprover,omitempty"`
	TimeReminderEmailFlag bool `json:"timeReminderEmailFlag,omitempty"`
	TimeSheetStartDate string `json:"timeSheetStartDate,omitempty"`
	TimeZone TimeZoneSetupReference `json:"timeZone,omitempty"`
	Title string `json:"title,omitempty"`
	ToastNotificationFlag bool `json:"toastNotificationFlag,omitempty"`
	Type MemberTypeReference `json:"type,omitempty"`
	Warehouse WarehouseReference `json:"warehouse,omitempty"`
	WarehouseBin WarehouseBinReference `json:"warehouseBin,omitempty"`
	WorkRole WorkRoleReference `json:"workRole,omitempty"`
	WorkType WorkTypeReference `json:"workType,omitempty"`
}

type MemberType struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name"`
}

type MemberTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type MenuEntry struct {
	Info interface{} `json:"_info,omitempty"`
	AddAllLocations bool `json:"addAllLocations,omitempty"`
	Caption string `json:"caption"`
	ClientId string `json:"clientId,omitempty"`
	ID int `json:"id,omitempty"`
	LargeMenuIconId int `json:"largeMenuIconId,omitempty"`
	Link string `json:"link"`
	LocationIds []int `json:"locationIds,omitempty"`
	MenuLocation MenuLocationReference `json:"menuLocation"`
	NewWindowFlag bool `json:"newWindowFlag,omitempty"`
	Origin string `json:"origin,omitempty"`
	RemoveAllLocations bool `json:"removeAllLocations,omitempty"`
	SmallMenuIconId int `json:"smallMenuIconId,omitempty"`
}

type MenuEntryLocation struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location"`
	MenuEntry SystemMenuEntryReference `json:"menuEntry,omitempty"`
}

type MyAccount struct {
	Info interface{} `json:"_info,omitempty"`
	AgreementInvoicingDisplayOptions string `json:"agreementInvoicingDisplayOptions,omitempty"`
	AllowExpensesEnteredAgainstCompaniesFlag bool `json:"allowExpensesEnteredAgainstCompaniesFlag,omitempty"`
	AllowInCellEntryOnTimeSheet bool `json:"allowInCellEntryOnTimeSheet,omitempty"`
	AuthenticationServiceType string `json:"authenticationServiceType,omitempty"`
	AutoPopupQuickNotesWithStopwatch bool `json:"autoPopupQuickNotesWithStopwatch,omitempty"`
	AutoStartStopwatch bool `json:"autoStartStopwatch,omitempty"`
	BillableForecast float64 `json:"billableForecast,omitempty"`
	Calendar CalendarReference `json:"calendar,omitempty"`
	CalendarSyncIntegrationFlag bool `json:"calendarSyncIntegrationFlag,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	CompanyActivityTabFormat string `json:"companyActivityTabFormat,omitempty"`
	CopyColumnLayoutsAndFilters bool `json:"copyColumnLayoutsAndFilters,omitempty"`
	CopyPodLayouts bool `json:"copyPodLayouts,omitempty"`
	CopySharedDefaultViews bool `json:"copySharedDefaultViews,omitempty"`
	Country CountryReference `json:"country,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
	DailyCapacity float64 `json:"dailyCapacity,omitempty"`
	DaysTolerance int `json:"daysTolerance,omitempty"`
	DefaultDepartment SystemDepartmentReference `json:"defaultDepartment"`
	DefaultEmail string `json:"defaultEmail,omitempty"`
	DefaultLocation SystemLocationReference `json:"defaultLocation"`
	DefaultPhone string `json:"defaultPhone,omitempty"`
	DirectionalSync DirectionalSyncReference `json:"directionalSync,omitempty"`
	DisableOnlineFlag bool `json:"disableOnlineFlag,omitempty"`
	EmployeeIdentifer string `json:"employeeIdentifer,omitempty"`
	EnableMobileFlag bool `json:"enableMobileFlag,omitempty"`
	EnableMobileGpsFlag bool `json:"enableMobileGpsFlag,omitempty"`
	EnterTimeAgainstCompanyFlag bool `json:"enterTimeAgainstCompanyFlag,omitempty"`
	ExpenseApprover MemberReference `json:"expenseApprover"`
	FirstName string `json:"firstName"`
	FromMemberRecId int `json:"fromMemberRecId,omitempty"`
	GlobalSearchDefaultSort string `json:"globalSearchDefaultSort,omitempty"`
	GlobalSearchDefaultTicketFilter string `json:"globalSearchDefaultTicketFilter,omitempty"`
	HideMemberInDispatchPortalFlag bool `json:"hideMemberInDispatchPortalFlag,omitempty"`
	HireDate time.Time `json:"hireDate"`
	HomeEmail string `json:"homeEmail,omitempty"`
	HomeExtension string `json:"homeExtension,omitempty"`
	HomePhone string `json:"homePhone,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier"`
	InactiveDate time.Time `json:"inactiveDate,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	IncludeInUtilizationReportingFlag bool `json:"includeInUtilizationReportingFlag,omitempty"`
	InvoiceScreenDefaultTabFormat string `json:"invoiceScreenDefaultTabFormat,omitempty"`
	InvoiceTimeTabFormat string `json:"invoiceTimeTabFormat,omitempty"`
	InvoicingDisplayOptions string `json:"invoicingDisplayOptions,omitempty"`
	LastLogin string `json:"lastLogin,omitempty"`
	LastName string `json:"lastName"`
	LicenseClass string `json:"licenseClass,omitempty"`
	MapiName string `json:"mapiName,omitempty"`
	MemberPersonas []int `json:"memberPersonas,omitempty"`
	MiddleInitial string `json:"middleInitial,omitempty"`
	MinimumHours float64 `json:"minimumHours,omitempty"`
	MobileEmail string `json:"mobileEmail,omitempty"`
	MobileExtension string `json:"mobileExtension,omitempty"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	Notes string `json:"notes,omitempty"`
	Office365 MemberOffice365 `json:"office365,omitempty"`
	OfficeEmail string `json:"officeEmail,omitempty"`
	OfficeExtension string `json:"officeExtension,omitempty"`
	OfficePhone string `json:"officePhone,omitempty"`
	PartnerPortalFlag bool `json:"partnerPortalFlag,omitempty"`
	Password string `json:"password,omitempty"`
	PhoneIntegrationType string `json:"phoneIntegrationType,omitempty"`
	PhoneSource string `json:"phoneSource,omitempty"`
	Photo DocumentReference `json:"photo,omitempty"`
	PrimaryEmail string `json:"primaryEmail,omitempty"`
	ProjectDefaultBoard ProjectBoardReference `json:"projectDefaultBoard,omitempty"`
	ProjectDefaultDepartment SystemDepartmentReference `json:"projectDefaultDepartment,omitempty"`
	ProjectDefaultLocation SystemLocationReference `json:"projectDefaultLocation,omitempty"`
	ReportCard ReportCardReference `json:"reportCard,omitempty"`
	ReportsTo MemberReference `json:"reportsTo,omitempty"`
	RequireExpenseEntryFlag bool `json:"requireExpenseEntryFlag,omitempty"`
	RequireStartAndEndTimeOnTimeEntryFlag bool `json:"requireStartAndEndTimeOnTimeEntryFlag,omitempty"`
	RequireTimeSheetEntryFlag bool `json:"requireTimeSheetEntryFlag,omitempty"`
	SalesDefaultLocation SystemLocationReference `json:"salesDefaultLocation"`
	ScheduleCapacity float64 `json:"scheduleCapacity,omitempty"`
	ScheduleDefaultDepartment SystemDepartmentReference `json:"scheduleDefaultDepartment,omitempty"`
	ScheduleDefaultLocation SystemLocationReference `json:"scheduleDefaultLocation,omitempty"`
	ServiceBoardTeamIds []int `json:"serviceBoardTeamIds,omitempty"`
	ServiceDefaultBoard BoardReference `json:"serviceDefaultBoard,omitempty"`
	ServiceDefaultDepartment SystemDepartmentReference `json:"serviceDefaultDepartment,omitempty"`
	ServiceDefaultLocation SystemLocationReference `json:"serviceDefaultLocation,omitempty"`
	ServiceLocation ServiceLocationReference `json:"serviceLocation,omitempty"`
	Signature string `json:"signature,omitempty"`
	StsUserAdminUrl string `json:"stsUserAdminUrl,omitempty"`
	TimeApprover MemberReference `json:"timeApprover"`
	TimeReminderEmailFlag bool `json:"timeReminderEmailFlag,omitempty"`
	TimeSheetStartDate time.Time `json:"timeSheetStartDate,omitempty"`
	TimeZone TimeZoneSetupReference `json:"timeZone"`
	TimebasedOneTimePasswordActivated bool `json:"timebasedOneTimePasswordActivated,omitempty"`
	Title string `json:"title,omitempty"`
	ToastNotificationFlag bool `json:"toastNotificationFlag,omitempty"`
	Token string `json:"token,omitempty"`
	Type MemberTypeReference `json:"type,omitempty"`
	UseBrowserLanguageFlag bool `json:"useBrowserLanguageFlag,omitempty"`
	VendorNumber string `json:"vendorNumber,omitempty"`
	Warehouse WarehouseReference `json:"warehouse,omitempty"`
	WarehouseBin WarehouseBinReference `json:"warehouseBin,omitempty"`
	WorkRole WorkRoleReference `json:"workRole"`
	WorkType WorkTypeReference `json:"workType,omitempty"`
}

type MyMember struct {
	Info interface{} `json:"_info,omitempty"`
	AdminFlag bool `json:"adminFlag,omitempty"`
	AgreementInvoicingDisplayOptions string `json:"agreementInvoicingDisplayOptions,omitempty"`
	AllowExpensesEnteredAgainstCompaniesFlag bool `json:"allowExpensesEnteredAgainstCompaniesFlag,omitempty"`
	AllowInCellEntryOnTimeSheet bool `json:"allowInCellEntryOnTimeSheet,omitempty"`
	AuthenticationServiceType string `json:"authenticationServiceType,omitempty"`
	BillableForecast float64 `json:"billableForecast,omitempty"`
	Calendar CalendarReference `json:"calendar,omitempty"`
	CalendarSyncIntegrationFlag bool `json:"calendarSyncIntegrationFlag,omitempty"`
	CompanyActivityTabFormat string `json:"companyActivityTabFormat,omitempty"`
	CorelyticsPassword string `json:"corelyticsPassword,omitempty"`
	CorelyticsUsername string `json:"corelyticsUsername,omitempty"`
	Country CountryReference `json:"country,omitempty"`
	DailyCapacity float64 `json:"dailyCapacity,omitempty"`
	DaysTolerance int `json:"daysTolerance,omitempty"`
	DefaultDepartment SystemDepartmentReference `json:"defaultDepartment,omitempty"`
	DefaultEmail string `json:"defaultEmail,omitempty"`
	DefaultLocation SystemLocationReference `json:"defaultLocation,omitempty"`
	DefaultPhone string `json:"defaultPhone,omitempty"`
	DirectionalSync DirectionalSyncReference `json:"directionalSync,omitempty"`
	DisableOnlineFlag bool `json:"disableOnlineFlag,omitempty"`
	EmployeeIdentifer string `json:"employeeIdentifer,omitempty"`
	EnableLdapAuthenticationFlag bool `json:"enableLdapAuthenticationFlag,omitempty"`
	EnableMobileFlag bool `json:"enableMobileFlag,omitempty"`
	EnableMobileGpsFlag bool `json:"enableMobileGpsFlag,omitempty"`
	EnterTimeAgainstCompanyFlag bool `json:"enterTimeAgainstCompanyFlag,omitempty"`
	ExcludedProjectBoardIds []int `json:"excludedProjectBoardIds,omitempty"`
	ExcludedServiceBoardIds []int `json:"excludedServiceBoardIds,omitempty"`
	ExpenseApprover MemberReference `json:"expenseApprover,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	HideMemberInDispatchPortalFlag bool `json:"hideMemberInDispatchPortalFlag,omitempty"`
	HireDate string `json:"hireDate,omitempty"`
	HomeEmail string `json:"homeEmail,omitempty"`
	HomeExtension string `json:"homeExtension,omitempty"`
	HomePhone string `json:"homePhone,omitempty"`
	HourlyCost float64 `json:"hourlyCost,omitempty"`
	HourlyRate float64 `json:"hourlyRate,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	InactiveDate string `json:"inactiveDate,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	IncludeInUtilizationReportingFlag bool `json:"includeInUtilizationReportingFlag,omitempty"`
	InvoiceScreenDefaultTabFormat string `json:"invoiceScreenDefaultTabFormat,omitempty"`
	InvoiceTimeTabFormat string `json:"invoiceTimeTabFormat,omitempty"`
	InvoicingDisplayOptions string `json:"invoicingDisplayOptions,omitempty"`
	LastLogin string `json:"lastLogin,omitempty"`
	LastName string `json:"lastName,omitempty"`
	LdapConfiguration LdapConfigurationReference `json:"ldapConfiguration,omitempty"`
	LdapUserName string `json:"ldapUserName,omitempty"`
	LicenseClass string `json:"licenseClass,omitempty"`
	MapiName string `json:"mapiName,omitempty"`
	MiddleInitial string `json:"middleInitial,omitempty"`
	MinimumHours float64 `json:"minimumHours,omitempty"`
	MobileEmail string `json:"mobileEmail,omitempty"`
	MobileExtension string `json:"mobileExtension,omitempty"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	Notes string `json:"notes,omitempty"`
	OfficeEmail string `json:"officeEmail,omitempty"`
	OfficeExtension string `json:"officeExtension,omitempty"`
	OfficePhone string `json:"officePhone,omitempty"`
	Password string `json:"password,omitempty"`
	Photo DocumentReference `json:"photo,omitempty"`
	ProjectDefaultBoard ProjectBoardReference `json:"projectDefaultBoard,omitempty"`
	ProjectDefaultDepartment SystemDepartmentReference `json:"projectDefaultDepartment,omitempty"`
	ProjectDefaultLocation SystemLocationReference `json:"projectDefaultLocation,omitempty"`
	ReportCard ReportCardReference `json:"reportCard,omitempty"`
	ReportsTo MemberReference `json:"reportsTo,omitempty"`
	RequireExpenseEntryFlag bool `json:"requireExpenseEntryFlag,omitempty"`
	RequireStartAndEndTimeOnTimeEntryFlag bool `json:"requireStartAndEndTimeOnTimeEntryFlag,omitempty"`
	RequireTimeSheetEntryFlag bool `json:"requireTimeSheetEntryFlag,omitempty"`
	RestrictDefaultSalesTerritoryFlag bool `json:"restrictDefaultSalesTerritoryFlag,omitempty"`
	RestrictDefaultWarehouseBinFlag bool `json:"restrictDefaultWarehouseBinFlag,omitempty"`
	RestrictDefaultWarehouseFlag bool `json:"restrictDefaultWarehouseFlag,omitempty"`
	RestrictDepartmentFlag bool `json:"restrictDepartmentFlag,omitempty"`
	RestrictLocationFlag bool `json:"restrictLocationFlag,omitempty"`
	RestrictProjectDefaultDepartmentFlag bool `json:"restrictProjectDefaultDepartmentFlag,omitempty"`
	RestrictProjectDefaultLocationFlag bool `json:"restrictProjectDefaultLocationFlag,omitempty"`
	RestrictScheduleFlag bool `json:"restrictScheduleFlag,omitempty"`
	RestrictServiceDefaultDepartmentFlag bool `json:"restrictServiceDefaultDepartmentFlag,omitempty"`
	RestrictServiceDefaultLocationFlag bool `json:"restrictServiceDefaultLocationFlag,omitempty"`
	SalesDefaultLocation SystemLocationReference `json:"salesDefaultLocation,omitempty"`
	ScheduleCapacity float64 `json:"scheduleCapacity,omitempty"`
	ScheduleDefaultDepartment SystemDepartmentReference `json:"scheduleDefaultDepartment,omitempty"`
	ScheduleDefaultLocation SystemLocationReference `json:"scheduleDefaultLocation,omitempty"`
	SecurityLocation SystemLocationReference `json:"securityLocation,omitempty"`
	SecurityRole SecurityRoleReference `json:"securityRole,omitempty"`
	ServiceBoardTeamIds []int `json:"serviceBoardTeamIds,omitempty"`
	ServiceDefaultBoard BoardReference `json:"serviceDefaultBoard,omitempty"`
	ServiceDefaultDepartment SystemDepartmentReference `json:"serviceDefaultDepartment,omitempty"`
	ServiceDefaultLocation SystemLocationReference `json:"serviceDefaultLocation,omitempty"`
	ServiceLocation ServiceLocationReference `json:"serviceLocation,omitempty"`
	SsoClientId string `json:"ssoClientId,omitempty"`
	SsoSessionFlag bool `json:"ssoSessionFlag,omitempty"`
	StructureLevel StructureReference `json:"structureLevel,omitempty"`
	TimeApprover MemberReference `json:"timeApprover,omitempty"`
	TimeReminderEmailFlag bool `json:"timeReminderEmailFlag,omitempty"`
	TimeSheetStartDate string `json:"timeSheetStartDate,omitempty"`
	TimeZone TimeZoneSetupReference `json:"timeZone,omitempty"`
	TimebasedOneTimePasswordActivated bool `json:"timebasedOneTimePasswordActivated,omitempty"`
	Title string `json:"title,omitempty"`
	ToastNotificationFlag bool `json:"toastNotificationFlag,omitempty"`
	Type MemberTypeReference `json:"type,omitempty"`
	VendorNumber string `json:"vendorNumber,omitempty"`
	Warehouse WarehouseReference `json:"warehouse,omitempty"`
	WarehouseBin WarehouseBinReference `json:"warehouseBin,omitempty"`
	WorkRole WorkRoleReference `json:"workRole,omitempty"`
	WorkType WorkTypeReference `json:"workType,omitempty"`
}

type MyMemberInfo struct {
	Info interface{} `json:"_info,omitempty"`
	AllowExpensesEnteredAgainstCompaniesFlag bool `json:"allowExpensesEnteredAgainstCompaniesFlag,omitempty"`
	DailyCapacity float64 `json:"dailyCapacity,omitempty"`
	DefaultDepartment SystemDepartmentReference `json:"defaultDepartment,omitempty"`
	DefaultEmail string `json:"defaultEmail,omitempty"`
	DefaultLocation SystemLocationReference `json:"defaultLocation,omitempty"`
	EnterTimeAgainstCompanyFlag bool `json:"enterTimeAgainstCompanyFlag,omitempty"`
	ExcludedProjectBoardIds []int `json:"excludedProjectBoardIds,omitempty"`
	ExcludedServiceBoardIds []int `json:"excludedServiceBoardIds,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	FullName string `json:"fullName,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	LastName string `json:"lastName,omitempty"`
	LicenseClass string `json:"licenseClass,omitempty"`
	MiddleInitial string `json:"middleInitial,omitempty"`
	Photo DocumentReference `json:"photo,omitempty"`
	ProjectDefaultBoard ProjectBoardReference `json:"projectDefaultBoard,omitempty"`
	ProjectDefaultDepartment SystemDepartmentReference `json:"projectDefaultDepartment,omitempty"`
	ProjectDefaultLocation SystemLocationReference `json:"projectDefaultLocation,omitempty"`
	RequireExpenseEntryFlag bool `json:"requireExpenseEntryFlag,omitempty"`
	RequireStartAndEndTimeOnTimeEntryFlag bool `json:"requireStartAndEndTimeOnTimeEntryFlag,omitempty"`
	RequireTimeSheetEntryFlag bool `json:"requireTimeSheetEntryFlag,omitempty"`
	RestrictDefaultWarehouseBinFlag bool `json:"restrictDefaultWarehouseBinFlag,omitempty"`
	RestrictDefaultWarehouseFlag bool `json:"restrictDefaultWarehouseFlag,omitempty"`
	RestrictProjectDefaultDepartmentFlag bool `json:"restrictProjectDefaultDepartmentFlag,omitempty"`
	RestrictProjectDefaultLocationFlag bool `json:"restrictProjectDefaultLocationFlag,omitempty"`
	RestrictServiceDefaultDepartmentFlag bool `json:"restrictServiceDefaultDepartmentFlag,omitempty"`
	RestrictServiceDefaultLocationFlag bool `json:"restrictServiceDefaultLocationFlag,omitempty"`
	SalesDefaultLocation SystemLocationReference `json:"salesDefaultLocation,omitempty"`
	ScheduleCapacity float64 `json:"scheduleCapacity,omitempty"`
	ScheduleDefaultDepartment SystemDepartmentReference `json:"scheduleDefaultDepartment,omitempty"`
	ScheduleDefaultLocation SystemLocationReference `json:"scheduleDefaultLocation,omitempty"`
	ServiceDefaultBoard BoardReference `json:"serviceDefaultBoard,omitempty"`
	ServiceDefaultDepartment SystemDepartmentReference `json:"serviceDefaultDepartment,omitempty"`
	ServiceDefaultLocation SystemLocationReference `json:"serviceDefaultLocation,omitempty"`
	ServiceLocation ServiceLocationReference `json:"serviceLocation,omitempty"`
	SsoClientId string `json:"ssoClientId,omitempty"`
	SsoSessionFlag bool `json:"ssoSessionFlag,omitempty"`
	TimeZone TimeZoneSetupReference `json:"timeZone,omitempty"`
	UseBrowserLanguageFlag bool `json:"useBrowserLanguageFlag,omitempty"`
	Warehouse WarehouseReference `json:"warehouse,omitempty"`
	WarehouseBin WarehouseBinReference `json:"warehouseBin,omitempty"`
	WorkRole WorkRoleReference `json:"workRole,omitempty"`
	WorkType WorkTypeReference `json:"workType,omitempty"`
}

type MySecurity struct {
	Info interface{} `json:"_info,omitempty"`
	AddLevel string `json:"addLevel,omitempty"`
	CustomFlag bool `json:"customFlag,omitempty"`
	DeleteLevel string `json:"deleteLevel,omitempty"`
	EditLevel string `json:"editLevel,omitempty"`
	ID int `json:"id,omitempty"`
	InquireLevel string `json:"inquireLevel,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	ModuleDescription string `json:"moduleDescription,omitempty"`
	ModuleFunctionDescription string `json:"moduleFunctionDescription,omitempty"`
	ModuleFunctionIdentifier string `json:"moduleFunctionIdentifier,omitempty"`
	ModuleFunctionName string `json:"moduleFunctionName,omitempty"`
	ModuleIdentifier string `json:"moduleIdentifier,omitempty"`
	ModuleName string `json:"moduleName,omitempty"`
	MyAllFlag bool `json:"myAllFlag,omitempty"`
	ReportFlag bool `json:"reportFlag,omitempty"`
	RestrictFlag bool `json:"restrictFlag,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type MySecurityCustomizeItem struct {
	CustomizeIdentifier string `json:"customizeIdentifier,omitempty"`
	ID int `json:"id,omitempty"`
	ItemIdentifier string `json:"itemIdentifier,omitempty"`
}

type NotificationRecipient struct {
	Info interface{} `json:"_info,omitempty"`
	AgreementFlag bool `json:"agreementFlag,omitempty"`
	ConfigFlag bool `json:"configFlag,omitempty"`
	ExternalFlag bool `json:"externalFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	InvoiceFlag bool `json:"invoiceFlag,omitempty"`
	KnowledgeBaseFlag bool `json:"knowledgeBaseFlag,omitempty"`
	MemberFlag bool `json:"memberFlag,omitempty"`
	MspFlag bool `json:"mspFlag,omitempty"`
	Name string `json:"name,omitempty"`
	ProcurementFlag bool `json:"procurementFlag,omitempty"`
	ProjectFlag bool `json:"projectFlag,omitempty"`
	SalesFlag bool `json:"salesFlag,omitempty"`
	ServiceFlag bool `json:"serviceFlag,omitempty"`
	TrackFlag bool `json:"trackFlag,omitempty"`
}

type Office365EmailApplicationInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Office365EmailSetup struct {
	Info interface{} `json:"_info,omitempty"`
	AuthorizedFlag bool `json:"authorizedFlag,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	EmailConnector EmailConnectorReference `json:"emailConnector,omitempty"`
	ExistingTenant ExistingTenantReference `json:"existingTenant,omitempty"`
	FailedFolder string `json:"failedFolder"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	InboxFolder string `json:"inboxFolder"`
	Name string `json:"name"`
	ProcessedFolder string `json:"processedFolder"`
	Source int `json:"source,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	UseExistingTenantFlag bool `json:"useExistingTenantFlag,omitempty"`
	Username string `json:"username,omitempty"`
}

type OnPremiseSearchSetting struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Password string `json:"password"`
}

type OsGradeWeight struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	OsGradeWeight float64 `json:"osGradeWeight,omitempty"`
	OsName string `json:"osName,omitempty"`
}

type Other struct {
	Info interface{} `json:"_info,omitempty"`
	ContactSync string `json:"contactSync,omitempty"`
	DefaultAddressFormat AddressFormatReference `json:"defaultAddressFormat"`
	DefaultCalendar CalendarReference `json:"defaultCalendar"`
	DefaultFromAddress string `json:"defaultFromAddress"`
	DefaultLdap LdapConfigurationReference `json:"defaultLdap,omitempty"`
	DisableZAdminLoginFlag bool `json:"disableZAdminLoginFlag,omitempty"`
	ID int `json:"id,omitempty"`
	IncludePortalLinkFlag bool `json:"includePortalLinkFlag,omitempty"`
	Locale LocaleReference `json:"locale"`
	LogoPath string `json:"logoPath,omitempty"`
	PortalUrlOverride string `json:"portalUrlOverride"`
	ServerTimeZone TimeZoneSetupReference `json:"serverTimeZone"`
	SiteUrl string `json:"siteUrl"`
	SyncLeadsFlag bool `json:"syncLeadsFlag,omitempty"`
	UpdateMemberTimeZonesFlag bool `json:"updateMemberTimeZonesFlag,omitempty"`
	UseExpandedFormatActivityFlag bool `json:"useExpandedFormatActivityFlag,omitempty"`
	UseExpandedFormatTimeFlag bool `json:"useExpandedFormatTimeFlag,omitempty"`
	UseSslFlag bool `json:"useSslFlag,omitempty"`
}

type ParsingType struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ParseRule string `json:"parseRule,omitempty"`
}

type ParsingVariable struct {
	Info interface{} `json:"_info,omitempty"`
	Code string `json:"code,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PatchOperation struct {
	Op string `json:"op,omitempty"`
	Path string `json:"path,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type PersonasInfo struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PortalReport struct {
	Info interface{} `json:"_info,omitempty"`
	CustomFlag bool `json:"customFlag,omitempty"`
	DisplayFlag bool `json:"displayFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	OpenSameWindowFlag bool `json:"openSameWindowFlag,omitempty"`
	PortalConfiguration PortalConfigurationReference `json:"portalConfiguration,omitempty"`
	URL string `json:"url"`
}

type QuoteLink struct {
	Info interface{} `json:"_info,omitempty"`
	AllLocationsFlag bool `json:"allLocationsFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Link string `json:"link"`
	Location SystemLocationReference `json:"location,omitempty"`
	NewWindowFlag bool `json:"newWindowFlag,omitempty"`
}

type Report struct {
	Name string `json:"name,omitempty"`
}

type ReportCard struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
}

type ReportCardDetail struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Kpi KPIReference `json:"kpi"`
	ReportCard ReportCardReference `json:"reportCard,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type ReportCardInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ReportColumnDefinition struct {
	IdentityColumn bool `json:"identityColumn,omitempty"`
	IsNullable bool `json:"isNullable,omitempty"`
	Type string `json:"type,omitempty"`
}

type ReportDataResponse struct {
	ColumnDefinitions []interface{} `json:"column_definitions,omitempty"`
	RowValues [][]interface{} `json:"row_values,omitempty"`
}

type ReportingService struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	ReportingDomain string `json:"reportingDomain,omitempty"`
	ReportingPassword string `json:"reportingPassword,omitempty"`
	ReportingUrl string `json:"reportingUrl,omitempty"`
	ReportingUserName string `json:"reportingUserName,omitempty"`
}

type SecurityRole struct {
	Info interface{} `json:"_info,omitempty"`
	AdminFlag bool `json:"adminFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name"`
	RoleType string `json:"roleType,omitempty"`
}

type SecurityRoleInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
	RoleType string `json:"roleType,omitempty"`
}

type SecurityRoleSetting struct {
	Info interface{} `json:"_info,omitempty"`
	AddLevel string `json:"addLevel,omitempty"`
	CustomFlag bool `json:"customFlag,omitempty"`
	DeleteLevel string `json:"deleteLevel,omitempty"`
	EditLevel string `json:"editLevel,omitempty"`
	ID int `json:"id,omitempty"`
	InquireLevel string `json:"inquireLevel,omitempty"`
	ModuleDescription string `json:"moduleDescription,omitempty"`
	ModuleFunctionDescription string `json:"moduleFunctionDescription,omitempty"`
	ModuleFunctionIdentifier string `json:"moduleFunctionIdentifier,omitempty"`
	ModuleFunctionName string `json:"moduleFunctionName,omitempty"`
	ModuleIdentifier string `json:"moduleIdentifier,omitempty"`
	ModuleName string `json:"moduleName,omitempty"`
	MyAllFlag bool `json:"myAllFlag,omitempty"`
	ReportFlag bool `json:"reportFlag,omitempty"`
	RestrictFlag bool `json:"restrictFlag,omitempty"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type Service struct {
	Info interface{} `json:"_info,omitempty"`
	AllowCCFlag bool `json:"allowCCFlag,omitempty"`
	AllowTOFlag bool `json:"allowTOFlag,omitempty"`
	CalendarSetup CalendarSetupReference `json:"calendarSetup,omitempty"`
	ContactColor string `json:"contactColor,omitempty"`
	ContactColorDisableFlag bool `json:"contactColorDisableFlag,omitempty"`
	HeaderColor string `json:"headerColor,omitempty"`
	HeaderColorDisableFlag bool `json:"headerColorDisableFlag,omitempty"`
	HideDelimiterFlag bool `json:"hideDelimiterFlag,omitempty"`
	ID int `json:"id,omitempty"`
	MemberColor string `json:"memberColor,omitempty"`
	MemberColorDisableFlag bool `json:"memberColorDisableFlag,omitempty"`
	ScheduleSpan string `json:"scheduleSpan"`
	SrNotify string `json:"srNotify,omitempty"`
	UnknownColor string `json:"unknownColor,omitempty"`
	UnknownColorDisableFlag bool `json:"unknownColorDisableFlag,omitempty"`
}

type ServiceInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ContactColor string `json:"contactColor,omitempty"`
	HeaderColor string `json:"headerColor,omitempty"`
	ID int `json:"id,omitempty"`
	MemberColor string `json:"memberColor,omitempty"`
	UnknownColor string `json:"unknownColor,omitempty"`
}

type SetupScreen struct {
	Info interface{} `json:"_info,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	ID int `json:"id,omitempty"`
	ModuleDescription string `json:"moduleDescription,omitempty"`
	ModuleIdentifier string `json:"moduleIdentifier,omitempty"`
	ModuleName string `json:"moduleName,omitempty"`
	Name string `json:"name,omitempty"`
}

type Skill struct {
	Info interface{} `json:"_info,omitempty"`
	Category SkillCategoryReference `json:"category"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
}

type SkillCategory struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
}

type SkillInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SsoConfiguration struct {
	Info interface{} `json:"_info,omitempty"`
	AllMembersSubmitted bool `json:"allMembersSubmitted,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	IsSsoOnByDefault bool `json:"isSsoOnByDefault,omitempty"`
	LocationIds []int `json:"locationIds"`
	Name string `json:"name"`
	SamlCertificateIssuedTo string `json:"samlCertificateIssuedTo,omitempty"`
	SamlCertificateName string `json:"samlCertificateName,omitempty"`
	SamlCertificateThumbprint string `json:"samlCertificateThumbprint,omitempty"`
	SamlCertificateValidFrom time.Time `json:"samlCertificateValidFrom,omitempty"`
	SamlCertificateValidTo time.Time `json:"samlCertificateValidTo,omitempty"`
	SamlEntityId string `json:"samlEntityId,omitempty"`
	SamlIdpCertificate string `json:"samlIdpCertificate,omitempty"`
	SamlSignInUrl string `json:"samlSignInUrl,omitempty"`
	SsoType string `json:"ssoType,omitempty"`
	StsBaseUrl string `json:"stsBaseUrl,omitempty"`
	StsUserAdminUrl string `json:"stsUserAdminUrl,omitempty"`
	SubmittedMemberCount int `json:"submittedMemberCount,omitempty"`
	Token string `json:"token,omitempty"`
}

type SsoUser struct {
	DateEntered string `json:"dateEntered,omitempty"`
	DisabledFlag bool `json:"disabledFlag,omitempty"`
	Email string `json:"email,omitempty"`
	EmailConfirmed bool `json:"emailConfirmed,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	ID int `json:"id,omitempty"`
	LastName string `json:"lastName,omitempty"`
	LastUpdated string `json:"lastUpdated,omitempty"`
	LinkedFlag bool `json:"linkedFlag,omitempty"`
	LinkedMember MemberReference `json:"linkedMember,omitempty"`
	SsoUserId string `json:"ssoUserId,omitempty"`
	UserName string `json:"userName,omitempty"`
}

type StandardNote struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	Contents string `json:"contents"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name"`
}

type StandardNoteInfo struct {
	Info interface{} `json:"_info,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	Contents string `json:"contents,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
	Success bool `json:"success,omitempty"`
}

type Survey struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Instructions string `json:"instructions,omitempty"`
	Name string `json:"name"`
}

type SurveyInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type SurveyQuestion struct {
	Info interface{} `json:"_info,omitempty"`
	EntryType string `json:"entryType,omitempty"`
	FieldType string `json:"fieldType,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	NumberOfDecimals int `json:"numberOfDecimals,omitempty"`
	Question string `json:"question"`
	RequiredFlag bool `json:"requiredFlag,omitempty"`
	SequenceNumber float64 `json:"sequenceNumber,omitempty"`
	Survey SurveyReference `json:"survey,omitempty"`
}

type SurveyQuestionValue struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	PointValue int `json:"pointValue,omitempty"`
	Question SurveyQuestionReference `json:"question,omitempty"`
	Survey SurveyReference `json:"survey,omitempty"`
	Value string `json:"value"`
}

type SystemSetting struct {
	Info interface{} `json:"_info,omitempty"`
	Description string `json:"description,omitempty"`
	ID int `json:"id,omitempty"`
	Value string `json:"value"`
	ValueType string `json:"valueType,omitempty"`
}

type TimeExpense struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultSpecialInvoiceType string `json:"defaultSpecialInvoiceType,omitempty"`
	DisableTimeEntryFlag bool `json:"disableTimeEntryFlag,omitempty"`
	ID int `json:"id,omitempty"`
	InternalCompany CompanyReference `json:"internalCompany"`
	InvoiceStart int `json:"invoiceStart,omitempty"`
	RequireExpenseNoteFlag bool `json:"requireExpenseNoteFlag,omitempty"`
	RequireTimeNoteFlag bool `json:"requireTimeNoteFlag,omitempty"`
	RoundingFactor float64 `json:"roundingFactor,omitempty"`
	Tier1ApprovalFlag bool `json:"tier1ApprovalFlag,omitempty"`
	Tier2ApprovalFlag bool `json:"tier2ApprovalFlag,omitempty"`
}

type TimeZoneSetup struct {
	Info interface{} `json:"_info,omitempty"`
	DaylightSavingsFlag bool `json:"daylightSavingsFlag,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Offset float64 `json:"offset,omitempty"`
	TimeZone TimeZoneReference `json:"timeZone"`
}

type TimeZoneSetupInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Offset float64 `json:"offset,omitempty"`
}

type TodayPageCategory struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name"`
	SortOrder int `json:"sortOrder,omitempty"`
}

type Token struct {
	Expiration string `json:"expiration,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

type UserDefinedField struct {
	Info interface{} `json:"_info,omitempty"`
	AddAllBusinessUnits bool `json:"addAllBusinessUnits,omitempty"`
	AddAllLocations bool `json:"addAllLocations,omitempty"`
	BusinessUnitIds []int `json:"businessUnitIds,omitempty"`
	ButtonUrl string `json:"buttonUrl,omitempty"`
	Caption string `json:"caption"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	DisplayOnScreenFlag bool `json:"displayOnScreenFlag,omitempty"`
	EntryTypeIdentifier string `json:"entryTypeIdentifier,omitempty"`
	FieldTypeIdentifier string `json:"fieldTypeIdentifier,omitempty"`
	HelpText string `json:"helpText,omitempty"`
	ID int `json:"id,omitempty"`
	ListViewFlag bool `json:"listViewFlag,omitempty"`
	LocationIds []int `json:"locationIds,omitempty"`
	NumberDecimals int `json:"numberDecimals,omitempty"`
	Options []UserDefinedFieldOption `json:"options,omitempty"`
	PodId int `json:"podId,omitempty"`
	ReadOnlyFlag bool `json:"readOnlyFlag,omitempty"`
	RemoveAllBusinessUnits bool `json:"removeAllBusinessUnits,omitempty"`
	RemoveAllLocations bool `json:"removeAllLocations,omitempty"`
	RequiredFlag bool `json:"requiredFlag,omitempty"`
	ScreenId string `json:"screenId,omitempty"`
	SequenceNumber int `json:"sequenceNumber,omitempty"`
}

type UserDefinedFieldInfo struct {
	Info interface{} `json:"_info,omitempty"`
	BusinessUnitIds []int `json:"businessUnitIds,omitempty"`
	ButtonUrl string `json:"buttonUrl,omitempty"`
	Caption string `json:"caption,omitempty"`
	DateCreated string `json:"dateCreated,omitempty"`
	DisplayOnScreenFlag bool `json:"displayOnScreenFlag,omitempty"`
	EntryTypeIdentifier string `json:"entryTypeIdentifier,omitempty"`
	FieldTypeIdentifier string `json:"fieldTypeIdentifier,omitempty"`
	HelpText string `json:"helpText,omitempty"`
	ID int `json:"id,omitempty"`
	ListViewFlag bool `json:"listViewFlag,omitempty"`
	LocationIds []int `json:"locationIds,omitempty"`
	NumberDecimals int `json:"numberDecimals,omitempty"`
	Options []UserDefinedFieldOption `json:"options,omitempty"`
	PodId int `json:"podId,omitempty"`
	ReadOnlyFlag bool `json:"readOnlyFlag,omitempty"`
	RequiredFlag bool `json:"requiredFlag,omitempty"`
	SequenceNumber int `json:"sequenceNumber,omitempty"`
}

type UserEmail struct {
	DisplayName string `json:"displayName,omitempty"`
	ID string `json:"id,omitempty"`
	UserPrincipalName string `json:"userPrincipalName,omitempty"`
}

type Workflow struct {
	Info interface{} `json:"_info,omitempty"`
	ActivateFlag bool `json:"activateFlag,omitempty"`
	BatchFrequencyUnit string `json:"batchFrequencyUnit,omitempty"`
	BatchInterval int `json:"batchInterval,omitempty"`
	BatchLastRan time.Time `json:"batchLastRan,omitempty"`
	BatchSchedule string `json:"batchSchedule,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	Department SystemDepartmentReference `json:"department,omitempty"`
	ID int `json:"id,omitempty"`
	Location SystemLocationReference `json:"location,omitempty"`
	Name string `json:"name"`
	TableType WorkflowTableTypeReference `json:"tableType"`
}

type WorkflowAction struct {
	Info interface{} `json:"_info,omitempty"`
	ActivityStatus ActivityStatusReference `json:"activityStatus,omitempty"`
	ActivityType ActivityTypeReference `json:"activityType,omitempty"`
	AttachConfigurationsFor string `json:"attachConfigurationsFor,omitempty"`
	AttachedTrack TrackReference `json:"attachedTrack,omitempty"`
	Attachments []int `json:"attachments,omitempty"`
	AuditNotesFlag bool `json:"auditNotesFlag,omitempty"`
	AutomateScript AutomateScriptReference `json:"automateScript,omitempty"`
	BccContact ContactReference `json:"bccContact,omitempty"`
	Board BoardReference `json:"board,omitempty"`
	BoardStatus ServiceStatusReference `json:"boardStatus,omitempty"`
	CcContact ContactReference `json:"ccContact,omitempty"`
	CompanyStatus CompanyStatusReference `json:"companyStatus,omitempty"`
	ConfigurationStatus ConfigurationStatusReference `json:"configurationStatus,omitempty"`
	ConfigurationType ConfigurationTypeReference `json:"configurationType,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	DaysToExecute int `json:"daysToExecute,omitempty"`
	DetailNotesFlag bool `json:"detailNotesFlag,omitempty"`
	EmailFrom string `json:"emailFrom,omitempty"`
	EmailRecipient string `json:"emailRecipient,omitempty"`
	GrandParentConnectWiseId string `json:"grandParentConnectWiseId,omitempty"`
	GrandParentId int `json:"grandParentId,omitempty"`
	Group GroupReference `json:"group,omitempty"`
	ID int `json:"id,omitempty"`
	InternalNotesFlag bool `json:"internalNotesFlag,omitempty"`
	InvoiceMinDays int `json:"invoiceMinDays,omitempty"`
	Notes string `json:"notes,omitempty"`
	NotifyFrom NotificationRecipientReference `json:"notifyFrom,omitempty"`
	NotifyType NotifyTypeReference `json:"notifyType"`
	NotifyWho NotificationRecipientReference `json:"notifyWho,omitempty"`
	ParentConnectWiseId string `json:"parentConnectWiseId,omitempty"`
	ParentId int `json:"parentId,omitempty"`
	ProjectStatus ProjectStatusReference `json:"projectStatus,omitempty"`
	SalesOrderStatus OrderStatusReference `json:"salesOrderStatus,omitempty"`
	ScriptFailStatus ServiceStatusReference `json:"scriptFailStatus,omitempty"`
	ScriptSuccessStatus ServiceStatusReference `json:"scriptSuccessStatus,omitempty"`
	ServiceItem ServiceItemReference `json:"serviceItem,omitempty"`
	ServicePriority PriorityReference `json:"servicePriority,omitempty"`
	ServiceSubType ServiceSubTypeReference `json:"serviceSubType,omitempty"`
	ServiceSurvey ServiceSurveyReference `json:"serviceSurvey,omitempty"`
	ServiceTemplate ServiceTemplateReference `json:"serviceTemplate,omitempty"`
	ServiceType ServiceTypeReference `json:"serviceType,omitempty"`
	SpecificMemberFrom MemberReference `json:"specificMemberFrom,omitempty"`
	SpecificMemberTo MemberReference `json:"specificMemberTo,omitempty"`
	SpecificTeamTo GenericBoardTeamReference `json:"specificTeamTo,omitempty"`
	Subject string `json:"subject,omitempty"`
	UpdateOwnerFlag bool `json:"updateOwnerFlag,omitempty"`
}

type WorkflowActionAutomateParameter struct {
	Info interface{} `json:"_info,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	ParentConnectWiseId string `json:"parentConnectWiseId,omitempty"`
	ParentId int `json:"parentId,omitempty"`
	Value string `json:"value,omitempty"`
}

type WorkflowActionUserDefinedField struct {
	Info interface{} `json:"_info,omitempty"`
	ActionId int `json:"actionId,omitempty"`
	Caption string `json:"caption,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	EntryTypeId string `json:"entryTypeId,omitempty"`
	EventId int `json:"eventId,omitempty"`
	FieldTypeId string `json:"fieldTypeId,omitempty"`
	GrandParentConnectWiseId string `json:"grandParentConnectWiseId,omitempty"`
	GrandParentId int `json:"grandParentId,omitempty"`
	ID int `json:"id,omitempty"`
	InactiveFlag bool `json:"inactiveFlag,omitempty"`
	OverwriteFlag bool `json:"overwriteFlag,omitempty"`
	ParentConnectWiseId string `json:"parentConnectWiseId,omitempty"`
	ParentId int `json:"parentId,omitempty"`
	PodDescription string `json:"podDescription,omitempty"`
	RequiredFlag bool `json:"requiredFlag,omitempty"`
	UserDefinedFieldId int `json:"userDefinedFieldId,omitempty"`
	Value string `json:"value,omitempty"`
}

type WorkflowAttachment struct {
	Info interface{} `json:"_info,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ParentConnectWiseId string `json:"parentConnectWiseId,omitempty"`
	ParentId int `json:"parentId,omitempty"`
}

type WorkflowEvent struct {
	Info interface{} `json:"_info,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	DateTestedUTC time.Time `json:"dateTestedUTC,omitempty"`
	EventCondition string `json:"eventCondition"`
	ExecutionTime string `json:"executionTime,omitempty"`
	FrequencyOfExecution int `json:"frequencyOfExecution,omitempty"`
	FrequencyUnit string `json:"frequencyUnit,omitempty"`
	ID int `json:"id,omitempty"`
	MaxNumberOfExecution int `json:"maxNumberOfExecution,omitempty"`
	Name string `json:"name,omitempty"`
	ParentConnectWiseId string `json:"parentConnectWiseId,omitempty"`
	ParentId int `json:"parentId,omitempty"`
	TestRecordsMatched int `json:"testRecordsMatched,omitempty"`
}

type WorkflowNotifyType struct {
	Info interface{} `json:"_info,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	ExternalFlag bool `json:"externalFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	IsSetupFlag bool `json:"isSetupFlag,omitempty"`
	Name string `json:"name,omitempty"`
	ParentConnectWiseId string `json:"parentConnectWiseId,omitempty"`
	ParentId int `json:"parentId,omitempty"`
}

type WorkflowNotifyTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	IsSetupFlag bool `json:"isSetupFlag,omitempty"`
	Name string `json:"name,omitempty"`
}

type WorkflowTableType struct {
	Info interface{} `json:"_info,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type WorkflowTableTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
}

type WorkflowTrigger struct {
	Info interface{} `json:"_info,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	CustomField UserDefinedFieldReference `json:"customField,omitempty"`
	Description string `json:"description,omitempty"`
	ExpectedType string `json:"expectedType,omitempty"`
	HasOperatorFlag bool `json:"hasOperatorFlag,omitempty"`
	HasOptionsFlag bool `json:"hasOptionsFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ParentConnectWiseId string `json:"parentConnectWiseId,omitempty"`
	ParentId int `json:"parentId,omitempty"`
}

type WorkflowTriggerOption struct {
	Info interface{} `json:"_info,omitempty"`
	ConnectWiseID string `json:"connectWiseID,omitempty"`
	CustomField UserDefinedFieldReference `json:"customField,omitempty"`
	GrandParentConnectWiseId string `json:"grandParentConnectWiseId,omitempty"`
	GrandParentId int `json:"grandParentId,omitempty"`
	Name string `json:"name,omitempty"`
	ParentConnectWiseId string `json:"parentConnectWiseId,omitempty"`
	ParentId int `json:"parentId,omitempty"`
	Value string `json:"value,omitempty"`
}

