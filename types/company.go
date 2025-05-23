package types

import (
	"time"
)

type AddressFormat struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	AddAllCountries    bool                   `json:"addAllCountries,omitempty"`
	CountryIds         []int                  `json:"countryIds,omitempty"`
	DefaultFlag        bool                   `json:"defaultFlag,omitempty"`
	Format             string                 `json:"format"`
	ID                 int                    `json:"id,omitempty"`
	Name               string                 `json:"name"`
	RemoveAllCountries bool                   `json:"removeAllCountries,omitempty"`
}

type AddressFormatInfo struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
}

type ClearPickerRequest struct {
	Member Member `json:"member,omitempty"`
	Type   string `json:"type,omitempty"`
}

type CommunicationType struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	AndroidXref  string                 `json:"androidXref,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	Description  string                 `json:"description"`
	EmailFlag    bool                   `json:"emailFlag,omitempty"`
	ExchangeXref string                 `json:"exchangeXref,omitempty"`
	FaxFlag      bool                   `json:"faxFlag,omitempty"`
	GoogleXref   string                 `json:"googleXref,omitempty"`
	ID           int                    `json:"id,omitempty"`
	IphoneXref   string                 `json:"iphoneXref,omitempty"`
	PhoneFlag    bool                   `json:"phoneFlag,omitempty"`
}

type CommunicationTypeInfo struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	Description string                 `json:"description,omitempty"`
	EmailFlag   bool                   `json:"emailFlag,omitempty"`
	FaxFlag     bool                   `json:"faxFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	PhoneFlag   bool                   `json:"phoneFlag,omitempty"`
}

type Company struct {
	Info                  map[string]interface{} `json:"_info,omitempty"`
	AccountNumber         string                 `json:"accountNumber,omitempty"`
	AddressLine1          string                 `json:"addressLine1,omitempty"`
	AddressLine2          string                 `json:"addressLine2,omitempty"`
	AnnualRevenue         float64                `json:"annualRevenue,omitempty"`
	BillToCompany         *Company               `json:"billToCompany,omitempty"`
	BillingContact        *Contact               `json:"billingContact,omitempty"`
	BillingSite           Site                   `json:"billingSite,omitempty"`
	BillingTerms          BillingTerms           `json:"billingTerms,omitempty"`
	Calendar              Calendar               `json:"calendar,omitempty"`
	City                  string                 `json:"city,omitempty"`
	CompanyEntityType     EntityType             `json:"companyEntityType,omitempty"`
	Country               Country                `json:"country,omitempty"`
	Currency              Currency               `json:"currency,omitempty"`
	CustomFields          []CustomFieldValue     `json:"customFields,omitempty"`
	DateAcquired          time.Time              `json:"dateAcquired,omitempty"`
	DateDeleted           time.Time              `json:"dateDeleted,omitempty"`
	DefaultContact        *Contact               `json:"defaultContact,omitempty"`
	DeletedBy             string                 `json:"deletedBy,omitempty"`
	DeletedFlag           bool                   `json:"deletedFlag,omitempty"`
	FacebookUrl           string                 `json:"facebookUrl,omitempty"`
	FaxNumber             string                 `json:"faxNumber,omitempty"`
	ID                    int                    `json:"id,omitempty"`
	Identifier            string                 `json:"identifier"`
	IntegratorTags        []string               `json:"integratorTags,omitempty"`
	InvoiceCCEmailAddress string                 `json:"invoiceCCEmailAddress,omitempty"`
	InvoiceDeliveryMethod BillingDelivery        `json:"invoiceDeliveryMethod,omitempty"`
	InvoiceTemplate       InvoiceTemplate        `json:"invoiceTemplate,omitempty"`
	InvoiceToEmailAddress string                 `json:"invoiceToEmailAddress,omitempty"`
	IsVendorFlag          bool                   `json:"isVendorFlag,omitempty"`
	LeadFlag              bool                   `json:"leadFlag,omitempty"`
	LeadSource            string                 `json:"leadSource,omitempty"`
	LinkedInUrl           string                 `json:"linkedInUrl,omitempty"`
	Market                MarketDescription      `json:"market,omitempty"`
	MobileGuid            string                 `json:"mobileGuid,omitempty"`
	Name                  string                 `json:"name"`
	NumberOfEmployees     int                    `json:"numberOfEmployees,omitempty"`
	OwnershipType         OwnershipType          `json:"ownershipType,omitempty"`
	ParentCompany         *Company               `json:"parentCompany,omitempty"`
	PhoneNumber           string                 `json:"phoneNumber,omitempty"`
	PricingSchedule       PricingSchedule        `json:"pricingSchedule,omitempty"`
	ResellerIdentifier    string                 `json:"resellerIdentifier,omitempty"`
	RevenueYear           int                    `json:"revenueYear,omitempty"`
	SicCode               SicCode                `json:"sicCode,omitempty"`
	Site                  Site                   `json:"site,omitempty"`
	State                 string                 `json:"state,omitempty"`
	Status                CompanyStatus          `json:"status,omitempty"`
	TaxCode               TaxCode                `json:"taxCode,omitempty"`
	TaxIdentifier         string                 `json:"taxIdentifier,omitempty"`
	Territory             SystemLocation         `json:"territory,omitempty"`
	TerritoryManager      Member                 `json:"territoryManager,omitempty"`
	TimeZoneSetup         TimeZoneSetup          `json:"timeZoneSetup,omitempty"`
	TwitterUrl            string                 `json:"twitterUrl,omitempty"`
	Types                 []CompanyType          `json:"types,omitempty"`
	UnsubscribeFlag       bool                   `json:"unsubscribeFlag,omitempty"`
	UserDefinedField1     string                 `json:"userDefinedField1,omitempty"`
	UserDefinedField10    string                 `json:"userDefinedField10,omitempty"`
	UserDefinedField2     string                 `json:"userDefinedField2,omitempty"`
	UserDefinedField3     string                 `json:"userDefinedField3,omitempty"`
	UserDefinedField4     string                 `json:"userDefinedField4,omitempty"`
	UserDefinedField5     string                 `json:"userDefinedField5,omitempty"`
	UserDefinedField6     string                 `json:"userDefinedField6,omitempty"`
	UserDefinedField7     string                 `json:"userDefinedField7,omitempty"`
	UserDefinedField8     string                 `json:"userDefinedField8,omitempty"`
	UserDefinedField9     string                 `json:"userDefinedField9,omitempty"`
	VendorIdentifier      string                 `json:"vendorIdentifier,omitempty"`
	Website               string                 `json:"website,omitempty"`
	YearEstablished       int                    `json:"yearEstablished,omitempty"`
	Zip                   string                 `json:"zip,omitempty"`
}

type CompanyCompanyTypeAssociation struct {
	Info    map[string]interface{} `json:"_info,omitempty"`
	Company Company                `json:"company"`
	ID      int                    `json:"id,omitempty"`
	Type    CompanyType            `json:"type"`
}

type CompanyConfiguration struct {
	Info                      map[string]interface{}  `json:"_info,omitempty"`
	ActiveFlag                bool                    `json:"activeFlag,omitempty"`
	BackupBillableSpaceGb     float64                 `json:"backupBillableSpaceGb,omitempty"`
	BackupFailed              int                     `json:"backupFailed,omitempty"`
	BackupIncomplete          int                     `json:"backupIncomplete,omitempty"`
	BackupMonth               int                     `json:"backupMonth,omitempty"`
	BackupProtectedDeviceList string                  `json:"backupProtectedDeviceList,omitempty"`
	BackupRestores            int                     `json:"backupRestores,omitempty"`
	BackupServerName          string                  `json:"backupServerName,omitempty"`
	BackupSuccesses           int                     `json:"backupSuccesses,omitempty"`
	BackupYear                int                     `json:"backupYear,omitempty"`
	BillFlag                  bool                    `json:"billFlag,omitempty"`
	BusinessUnitId            int                     `json:"businessUnitId,omitempty"`
	Company                   Company                 `json:"company"`
	CompanyLocationId         int                     `json:"companyLocationId,omitempty"`
	Contact                   Contact                 `json:"contact,omitempty"`
	CpuSpeed                  string                  `json:"cpuSpeed,omitempty"`
	CustomFields              []CustomFieldValue      `json:"customFields,omitempty"`
	DefaultGateway            string                  `json:"defaultGateway,omitempty"`
	Department                SystemDepartment        `json:"department,omitempty"`
	DeviceIdentifier          string                  `json:"deviceIdentifier,omitempty"`
	DisplayVendorFlag         bool                    `json:"displayVendorFlag,omitempty"`
	ID                        int                     `json:"id,omitempty"`
	InstallationDate          time.Time               `json:"installationDate,omitempty"`
	InstalledBy               Member                  `json:"installedBy,omitempty"`
	IpAddress                 string                  `json:"ipAddress,omitempty"`
	LastBackupDate            time.Time               `json:"lastBackupDate,omitempty"`
	LastLoginName             string                  `json:"lastLoginName,omitempty"`
	LocalHardDrives           string                  `json:"localHardDrives,omitempty"`
	Location                  SystemLocation          `json:"location,omitempty"`
	LocationId                int                     `json:"locationId,omitempty"`
	MacAddress                string                  `json:"macAddress,omitempty"`
	ManagementLink            string                  `json:"managementLink,omitempty"`
	Manufacturer              Manufacturer            `json:"manufacturer,omitempty"`
	ManufacturerPartNumber    string                  `json:"manufacturerPartNumber,omitempty"`
	MobileGuid                string                  `json:"mobileGuid,omitempty"`
	ModelNumber               string                  `json:"modelNumber,omitempty"`
	Name                      string                  `json:"name"`
	NeedsRenewalFlag          bool                    `json:"needsRenewalFlag,omitempty"`
	Notes                     string                  `json:"notes,omitempty"`
	OsInfo                    string                  `json:"osInfo,omitempty"`
	OsType                    string                  `json:"osType,omitempty"`
	ParentConfigurationId     int                     `json:"parentConfigurationId,omitempty"`
	PurchaseDate              time.Time               `json:"purchaseDate,omitempty"`
	Questions                 []ConfigurationQuestion `json:"questions,omitempty"`
	Ram                       string                  `json:"ram,omitempty"`
	RemoteLink                string                  `json:"remoteLink,omitempty"`
	SerialNumber              string                  `json:"serialNumber,omitempty"`
	ShowAutomateFlag          bool                    `json:"showAutomateFlag,omitempty"`
	ShowRemoteFlag            bool                    `json:"showRemoteFlag,omitempty"`
	Site                      Site                    `json:"site,omitempty"`
	Sla                       SLA                     `json:"sla,omitempty"`
	Status                    ConfigurationStatus     `json:"status,omitempty"`
	TagNumber                 string                  `json:"tagNumber,omitempty"`
	Type                      ConfigurationType       `json:"type"`
	Vendor                    Company                 `json:"vendor,omitempty"`
	VendorNotes               string                  `json:"vendorNotes,omitempty"`
	WarrantyExpirationDate    time.Time               `json:"warrantyExpirationDate,omitempty"`
}

type CompanyContactTypeAssociation struct {
	Info    map[string]interface{} `json:"_info,omitempty"`
	Contact Contact                `json:"contact"`
	ID      int                    `json:"id,omitempty"`
	Type    ContactType            `json:"type"`
}

type CompanyCompanyTypeAssociationCompanyTypeAssociation struct {
	Info    map[string]interface{} `json:"_info,omitempty"`
	Company Company                `json:"company"`
	ID      int                    `json:"id,omitempty"`
	Type    CompanyType            `json:"type"`
}

type CompanyCustomNote struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Company    Company                `json:"company,omitempty"`
	CustomNote string                 `json:"customNote"`
	ID         int                    `json:"id,omitempty"`
	Status     CompanyStatus          `json:"status"`
}

type CompanyGroup struct {
	Info                  map[string]interface{} `json:"_info,omitempty"`
	AllContactsFlag       bool                   `json:"allContactsFlag,omitempty"`
	Company               Company                `json:"company,omitempty"`
	ContactIds            []int                  `json:"contactIds,omitempty"`
	DefaultContactFlag    bool                   `json:"defaultContactFlag,omitempty"`
	Group                 Group                  `json:"group"`
	ID                    int                    `json:"id,omitempty"`
	RemoveAllContactsFlag bool                   `json:"removeAllContactsFlag,omitempty"`
	UnsubscribeFlag       bool                   `json:"unsubscribeFlag,omitempty"`
}

type CompanyInfo struct {
	Info             map[string]interface{} `json:"_info,omitempty"`
	AddressLine1     string                 `json:"addressLine1,omitempty"`
	AddressLine2     string                 `json:"addressLine2,omitempty"`
	BillToCompany    Company                `json:"billToCompany,omitempty"`
	BillingContact   Contact                `json:"billingContact,omitempty"`
	BillingSite      Site                   `json:"billingSite,omitempty"`
	BillingTerms     BillingTerms           `json:"billingTerms,omitempty"`
	City             string                 `json:"city,omitempty"`
	Country          Country                `json:"country,omitempty"`
	Currency         Currency               `json:"currency,omitempty"`
	DefaultContact   Contact                `json:"defaultContact,omitempty"`
	DeletedFlag      bool                   `json:"deletedFlag,omitempty"`
	FacebookUrl      string                 `json:"facebookUrl,omitempty"`
	FaxNumber        string                 `json:"faxNumber,omitempty"`
	ID               int                    `json:"id,omitempty"`
	Identifier       string                 `json:"identifier,omitempty"`
	IsVendorFlag     bool                   `json:"isVendorFlag,omitempty"`
	LeadFlag         bool                   `json:"leadFlag,omitempty"`
	LinkedInUrl      string                 `json:"linkedInUrl,omitempty"`
	Name             string                 `json:"name,omitempty"`
	NoServiceFlag    bool                   `json:"noServiceFlag,omitempty"`
	PhoneNumber      string                 `json:"phoneNumber,omitempty"`
	Site             Site                   `json:"site,omitempty"`
	State            string                 `json:"state,omitempty"`
	Status           CompanyStatus          `json:"status,omitempty"`
	TaxCode          TaxCode                `json:"taxCode,omitempty"`
	TaxIdentifier    string                 `json:"taxIdentifier,omitempty"`
	Territory        SystemLocation         `json:"territory,omitempty"`
	TwitterUrl       string                 `json:"twitterUrl,omitempty"`
	Types            []CompanyType          `json:"types,omitempty"`
	VendorIdentifier string                 `json:"vendorIdentifier,omitempty"`
	Zip              string                 `json:"zip,omitempty"`
}

type CompanyManagementSummary struct {
	Info                            map[string]interface{} `json:"_info,omitempty"`
	Agreement                       Agreement              `json:"agreement,omitempty"`
	AlertsGenerated                 string                 `json:"alertsGenerated,omitempty"`
	Company                         Company                `json:"company,omitempty"`
	CpuUtilization                  float64                `json:"cpuUtilization,omitempty"`
	DeviceType                      string                 `json:"deviceType,omitempty"`
	DiskCleanups                    int                    `json:"diskCleanups,omitempty"`
	DiskDefragmentations            int                    `json:"diskDefragmentations,omitempty"`
	DiskSpaceCleanedMb              int                    `json:"diskSpaceCleanedMb,omitempty"`
	FailedBackupJobs                int                    `json:"failedBackupJobs,omitempty"`
	FullyPatchedMachines            int                    `json:"fullyPatchedMachines,omitempty"`
	GroupIdentifier                 string                 `json:"groupIdentifier"`
	ID                              int                    `json:"id,omitempty"`
	InternetConnectivity            float64                `json:"internetConnectivity,omitempty"`
	ManagementSolution              ManagementSolution     `json:"managementSolution,omitempty"`
	MemoryUtilization               float64                `json:"memoryUtilization,omitempty"`
	MissingMoreFivePatchesMachines  int                    `json:"missingMoreFivePatchesMachines,omitempty"`
	MissingOneTwoPatchesMachines    int                    `json:"missingOneTwoPatchesMachines,omitempty"`
	MissingSecurityPatches          string                 `json:"missingSecurityPatches,omitempty"`
	MissingThreeFivePatchesMachines int                    `json:"missingThreeFivePatchesMachines,omitempty"`
	MissingUnscannedPatchesMachines int                    `json:"missingUnscannedPatchesMachines,omitempty"`
	ServerAvailability              int                    `json:"serverAvailability,omitempty"`
	ServersDiskSpaceLow             int                    `json:"serversDiskSpaceLow,omitempty"`
	ServersOffline                  int                    `json:"serversOffline,omitempty"`
	SnmpMachines                    int                    `json:"snmpMachines,omitempty"`
	SpywareItemsRemoved             int                    `json:"spywareItemsRemoved,omitempty"`
	SuccessfulBackupJobs            int                    `json:"successfulBackupJobs,omitempty"`
	TotalManagedMachines            int                    `json:"totalManagedMachines,omitempty"`
	TotalNotifications              int                    `json:"totalNotifications,omitempty"`
	TotalServers                    int                    `json:"totalServers,omitempty"`
	TotalWindowsServers             int                    `json:"totalWindowsServers,omitempty"`
	TotalWindowsWorkstations        int                    `json:"totalWindowsWorkstations,omitempty"`
	TotalWorkstations               int                    `json:"totalWorkstations,omitempty"`
	VirusesRemoved                  int                    `json:"virusesRemoved,omitempty"`
	WindowsPatchesInstalled         int                    `json:"windowsPatchesInstalled,omitempty"`
}

type CompanyMerge struct {
	AccountNumber      string `json:"accountNumber,omitempty"`
	Activities         string `json:"activities,omitempty"`
	BillingAddress     string `json:"billingAddress,omitempty"`
	BillingContact     string `json:"billingContact,omitempty"`
	BillingTerms       string `json:"billingTerms,omitempty"`
	Contacts           string `json:"contacts,omitempty"`
	DateAcquired       string `json:"dateAcquired,omitempty"`
	Documents          string `json:"documents,omitempty"`
	Fax                string `json:"fax,omitempty"`
	Identifier         string `json:"identifier,omitempty"`
	Market             string `json:"market,omitempty"`
	Name               string `json:"name,omitempty"`
	Notes              string `json:"notes,omitempty"`
	NumberOfEmployees  string `json:"numberOfEmployees,omitempty"`
	Opportunities      string `json:"opportunities,omitempty"`
	Phone              string `json:"phone,omitempty"`
	PrimaryAddress     string `json:"primaryAddress,omitempty"`
	PrimaryContact     string `json:"primaryContact,omitempty"`
	Projects           string `json:"projects,omitempty"`
	Revenue            string `json:"revenue,omitempty"`
	RevenueYear        string `json:"revenueYear,omitempty"`
	Services           string `json:"services,omitempty"`
	SicCode            string `json:"sicCode,omitempty"`
	Sites              string `json:"sites,omitempty"`
	SourceList         string `json:"sourceList,omitempty"`
	Status             string `json:"status,omitempty"`
	TaxCode            string `json:"taxCode,omitempty"`
	Territory          string `json:"territory,omitempty"`
	TimeZone           string `json:"timeZone,omitempty"`
	ToCompanyId        int    `json:"toCompanyId"`
	Type               string `json:"type,omitempty"`
	UserDefinedField1  string `json:"userDefinedField1,omitempty"`
	UserDefinedField10 string `json:"userDefinedField10,omitempty"`
	UserDefinedField2  string `json:"userDefinedField2,omitempty"`
	UserDefinedField3  string `json:"userDefinedField3,omitempty"`
	UserDefinedField4  string `json:"userDefinedField4,omitempty"`
	UserDefinedField5  string `json:"userDefinedField5,omitempty"`
	UserDefinedField6  string `json:"userDefinedField6,omitempty"`
	UserDefinedField7  string `json:"userDefinedField7,omitempty"`
	UserDefinedField8  string `json:"userDefinedField8,omitempty"`
	UserDefinedField9  string `json:"userDefinedField9,omitempty"`
	Website            string `json:"website,omitempty"`
}

type CompanyNote struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	Company   Company                `json:"company,omitempty"`
	EnteredBy string                 `json:"enteredBy,omitempty"`
	Flagged   bool                   `json:"flagged,omitempty"`
	ID        int                    `json:"id,omitempty"`
	Text      string                 `json:"text"`
	Type      NoteType               `json:"type,omitempty"`
}

type CompanyNoteType struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Identifier  string                 `json:"identifier,omitempty"`
	ImportFlag  bool                   `json:"importFlag,omitempty"`
	Name        string                 `json:"name"`
}

type CompanyNoteTypeInfo struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Identifier  string                 `json:"identifier,omitempty"`
	Name        string                 `json:"name,omitempty"`
}

type CompanyPickerItem struct {
	Info             map[string]interface{} `json:"_info,omitempty"`
	Company          Company                `json:"company"`
	CompanyCountry   Country                `json:"companyCountry,omitempty"`
	CompanyLocation  SystemLocation         `json:"companyLocation,omitempty"`
	CompanySite      Site                   `json:"companySite,omitempty"`
	CompanyStatus    CompanyStatus          `json:"companyStatus,omitempty"`
	CompanyType      CompanyType            `json:"companyType,omitempty"`
	ID               int                    `json:"id,omitempty"`
	Member           Member                 `json:"member,omitempty"`
	VendorPickerFlag bool                   `json:"vendorPickerFlag,omitempty"`
}

type CompanySite struct {
	Info                 map[string]interface{} `json:"_info,omitempty"`
	AddressFormat        string                 `json:"addressFormat,omitempty"`
	AddressLine1         string                 `json:"addressLine1,omitempty"`
	AddressLine2         string                 `json:"addressLine2,omitempty"`
	BillSeparateFlag     bool                   `json:"billSeparateFlag,omitempty"`
	Calendar             Calendar               `json:"calendar,omitempty"`
	City                 string                 `json:"city,omitempty"`
	Company              Company                `json:"company,omitempty"`
	Country              Country                `json:"country,omitempty"`
	CustomFields         []CustomFieldValue     `json:"customFields,omitempty"`
	DefaultBillingFlag   bool                   `json:"defaultBillingFlag,omitempty"`
	DefaultMailingFlag   bool                   `json:"defaultMailingFlag,omitempty"`
	DefaultShippingFlag  bool                   `json:"defaultShippingFlag,omitempty"`
	EntityType           EntityType             `json:"entityType,omitempty"`
	ExpenseReimbursement float64                `json:"expenseReimbursement,omitempty"`
	FaxNumber            string                 `json:"faxNumber,omitempty"`
	ID                   int                    `json:"id,omitempty"`
	InactiveFlag         bool                   `json:"inactiveFlag,omitempty"`
	MobileGuid           string                 `json:"mobileGuid,omitempty"`
	Name                 string                 `json:"name"`
	PhoneNumber          string                 `json:"phoneNumber,omitempty"`
	PhoneNumberExt       string                 `json:"phoneNumberExt,omitempty"`
	PrimaryAddressFlag   bool                   `json:"primaryAddressFlag,omitempty"`
	State                State                  `json:"state,omitempty"`
	TaxCode              TaxCode                `json:"taxCode,omitempty"`
	TimeZone             TimeZoneSetup          `json:"timeZone,omitempty"`
	Zip                  string                 `json:"zip,omitempty"`
}

type CompanySiteInfo struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	AddressFormat       string                 `json:"addressFormat,omitempty"`
	AddressLine1        string                 `json:"addressLine1,omitempty"`
	AddressLine2        string                 `json:"addressLine2,omitempty"`
	City                string                 `json:"city,omitempty"`
	Country             Country                `json:"country,omitempty"`
	DefaultBillingFlag  bool                   `json:"defaultBillingFlag,omitempty"`
	DefaultShippingFlag bool                   `json:"defaultShippingFlag,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	InactiveFlag        bool                   `json:"inactiveFlag,omitempty"`
	Name                string                 `json:"name,omitempty"`
	PhoneNumber         string                 `json:"phoneNumber,omitempty"`
	PhoneNumberExt      string                 `json:"phoneNumberExt,omitempty"`
	PrimaryAddressFlag  bool                   `json:"primaryAddressFlag,omitempty"`
	State               State                  `json:"state,omitempty"`
	TaxCode             TaxCode                `json:"taxCode,omitempty"`
	Zip                 string                 `json:"zip,omitempty"`
}

type CompanyStatus struct {
	Info                 map[string]interface{} `json:"_info,omitempty"`
	CancelOpenTracksFlag bool                   `json:"cancelOpenTracksFlag,omitempty"`
	CustomNoteFlag       bool                   `json:"customNoteFlag,omitempty"`
	DefaultFlag          bool                   `json:"defaultFlag,omitempty"`
	DisallowSavingFlag   bool                   `json:"disallowSavingFlag,omitempty"`
	ID                   int                    `json:"id,omitempty"`
	InactiveFlag         bool                   `json:"inactiveFlag,omitempty"`
	Name                 string                 `json:"name"`
	NotificationMessage  string                 `json:"notificationMessage,omitempty"`
	NotifyFlag           bool                   `json:"notifyFlag,omitempty"`
	Track                Track                  `json:"track,omitempty"`
}

type CompanyTeam struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	AccountManagerFlag bool                   `json:"accountManagerFlag,omitempty"`
	Company            Company                `json:"company,omitempty"`
	Contact            Contact                `json:"contact,omitempty"`
	Department         SystemDepartment       `json:"department,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	Location           SystemLocation         `json:"location,omitempty"`
	Member             Member                 `json:"member,omitempty"`
	SalesFlag          bool                   `json:"salesFlag,omitempty"`
	TeamRole           TeamRole               `json:"teamRole"`
	TechFlag           bool                   `json:"techFlag,omitempty"`
}

type CompanyType struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag         bool                   `json:"defaultFlag,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	Name                string                 `json:"name"`
	ServiceAlertFlag    bool                   `json:"serviceAlertFlag,omitempty"`
	ServiceAlertMessage string                 `json:"serviceAlertMessage,omitempty"`
	VendorFlag          bool                   `json:"vendorFlag,omitempty"`
}

type CompanyTypeInfo struct {
	Info     map[string]interface{} `json:"_info,omitempty"`
	ID       int                    `json:"id,omitempty"`
	IsVendor bool                   `json:"isVendor,omitempty"`
	Name     string                 `json:"name,omitempty"`
}

type ConfigurationStatus struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag  bool                   `json:"closedFlag,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	Description string                 `json:"description"`
	ID          int                    `json:"id,omitempty"`
}

type ConfigurationStatusInfo struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag  bool                   `json:"closedFlag,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	Description string                 `json:"description,omitempty"`
	ID          int                    `json:"id,omitempty"`
}

type ConfigurationTabsCount struct {
}

type ConfigurationType struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name"`
	SystemFlag   bool                   `json:"systemFlag,omitempty"`
}

type ConfigurationTypeCopy struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ConfigurationTypeInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
	SystemFlag   bool                   `json:"systemFlag,omitempty"`
}

type ConfigurationTypeQuestion struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	ConfigurationType ConfigurationType      `json:"configurationType,omitempty"`
	EntryType         string                 `json:"entryType,omitempty"`
	FieldType         string                 `json:"fieldType,omitempty"`
	ID                int                    `json:"id,omitempty"`
	InactiveFlag      bool                   `json:"inactiveFlag,omitempty"`
	NumberOfDecimals  int                    `json:"numberOfDecimals,omitempty"`
	Question          string                 `json:"question"`
	RequiredFlag      bool                   `json:"requiredFlag,omitempty"`
	SequenceNumber    float64                `json:"sequenceNumber,omitempty"`
}

type ConfigurationTypeQuestionValue struct {
	Info              map[string]interface{}    `json:"_info,omitempty"`
	ConfigurationType ConfigurationType         `json:"configurationType,omitempty"`
	DefaultFlag       bool                      `json:"defaultFlag,omitempty"`
	ID                int                       `json:"id,omitempty"`
	InactiveFlag      bool                      `json:"inactiveFlag,omitempty"`
	Question          ConfigurationTypeQuestion `json:"question,omitempty"`
	Value             string                    `json:"value"`
}

type Contact struct {
	Info                   map[string]interface{}     `json:"_info,omitempty"`
	AddressLine1           string                     `json:"addressLine1,omitempty"`
	AddressLine2           string                     `json:"addressLine2,omitempty"`
	Anniversary            string                     `json:"anniversary,omitempty"`
	AssistantContact       *Contact                   `json:"assistantContact,omitempty"`
	BirthDay               string                     `json:"birthDay,omitempty"`
	Children               string                     `json:"children,omitempty"`
	ChildrenFlag           bool                       `json:"childrenFlag,omitempty"`
	City                   string                     `json:"city,omitempty"`
	CommunicationItems     []ContactCommunicationItem `json:"communicationItems,omitempty"`
	Company                Company                    `json:"company,omitempty"`
	CompanyLocation        SystemLocation             `json:"companyLocation,omitempty"`
	Country                Country                    `json:"country,omitempty"`
	CustomFields           []CustomFieldValue         `json:"customFields,omitempty"`
	DefaultBillingFlag     bool                       `json:"defaultBillingFlag,omitempty"`
	DefaultFlag            bool                       `json:"defaultFlag,omitempty"`
	DefaultMergeContactId  int                        `json:"defaultMergeContactId,omitempty"`
	DefaultPhoneExtension  string                     `json:"defaultPhoneExtension,omitempty"`
	DefaultPhoneNbr        string                     `json:"defaultPhoneNbr,omitempty"`
	DefaultPhoneType       string                     `json:"defaultPhoneType,omitempty"`
	Department             ContactDepartment          `json:"department,omitempty"`
	DisablePortalLoginFlag bool                       `json:"disablePortalLoginFlag,omitempty"`
	FacebookUrl            string                     `json:"facebookUrl,omitempty"`
	FirstName              string                     `json:"firstName,omitempty"`
	Gender                 string                     `json:"gender,omitempty"`
	ID                     int                        `json:"id,omitempty"`
	IgnoreDuplicates       bool                       `json:"ignoreDuplicates,omitempty"`
	InactiveFlag           bool                       `json:"inactiveFlag,omitempty"`
	IntegratorTags         []string                   `json:"integratorTags,omitempty"`
	LastName               string                     `json:"lastName,omitempty"`
	LinkedInUrl            string                     `json:"linkedInUrl,omitempty"`
	ManagerContact         *Contact                   `json:"managerContact,omitempty"`
	MarriedFlag            bool                       `json:"marriedFlag,omitempty"`
	MobileGuid             string                     `json:"mobileGuid,omitempty"`
	NickName               string                     `json:"nickName,omitempty"`
	Photo                  Document                   `json:"photo,omitempty"`
	PortalPassword         string                     `json:"portalPassword,omitempty"`
	PortalSecurityLevel    int                        `json:"portalSecurityLevel,omitempty"`
	Presence               string                     `json:"presence,omitempty"`
	Relationship           Relationship               `json:"relationship,omitempty"`
	RelationshipOverride   string                     `json:"relationshipOverride,omitempty"`
	School                 string                     `json:"school,omitempty"`
	SecurityIdentifier     string                     `json:"securityIdentifier,omitempty"`
	SignificantOther       string                     `json:"significantOther,omitempty"`
	Site                   Site                       `json:"site,omitempty"`
	State                  string                     `json:"state,omitempty"`
	Title                  string                     `json:"title,omitempty"`
	TwitterUrl             string                     `json:"twitterUrl,omitempty"`
	TypeIds                []int                      `json:"typeIds,omitempty"`
	Types                  []ContactType              `json:"types,omitempty"`
	UnsubscribeFlag        bool                       `json:"unsubscribeFlag,omitempty"`
	UserDefinedField1      string                     `json:"userDefinedField1,omitempty"`
	UserDefinedField10     string                     `json:"userDefinedField10,omitempty"`
	UserDefinedField2      string                     `json:"userDefinedField2,omitempty"`
	UserDefinedField3      string                     `json:"userDefinedField3,omitempty"`
	UserDefinedField4      string                     `json:"userDefinedField4,omitempty"`
	UserDefinedField5      string                     `json:"userDefinedField5,omitempty"`
	UserDefinedField6      string                     `json:"userDefinedField6,omitempty"`
	UserDefinedField7      string                     `json:"userDefinedField7,omitempty"`
	UserDefinedField8      string                     `json:"userDefinedField8,omitempty"`
	UserDefinedField9      string                     `json:"userDefinedField9,omitempty"`
	Zip                    string                     `json:"zip,omitempty"`
}

type ContactCommunication struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	CommunicationType string                 `json:"communicationType,omitempty"`
	ContactId         int                    `json:"contactId,omitempty"`
	DefaultFlag       bool                   `json:"defaultFlag,omitempty"`
	Domain            string                 `json:"domain,omitempty"`
	Extension         string                 `json:"extension,omitempty"`
	ID                int                    `json:"id,omitempty"`
	MobileGuid        string                 `json:"mobileGuid,omitempty"`
	Type              CommunicationType      `json:"type"`
	Value             string                 `json:"value"`
}

type ContactContactTypeAssociationContactTypeAssociation struct {
	Info    map[string]interface{} `json:"_info,omitempty"`
	Contact Contact                `json:"contact"`
	ID      int                    `json:"id,omitempty"`
	Type    ContactType            `json:"type"`
}

type ContactDepartment struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name"`
}

type ContactDepartmentInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type ContactGroup struct {
	Info                                 map[string]interface{} `json:"_info,omitempty"`
	CompanyGroupUnsubscribedEmailMessage string                 `json:"companyGroupUnsubscribedEmailMessage,omitempty"`
	CompanyUnsubcribedEmailMessage       string                 `json:"companyUnsubcribedEmailMessage,omitempty"`
	Contact                              Contact                `json:"contact,omitempty"`
	ContactGroupUnsubscribedEmailMessage string                 `json:"contactGroupUnsubscribedEmailMessage,omitempty"`
	ContactUnsubscribedEmailMessage      string                 `json:"contactUnsubscribedEmailMessage,omitempty"`
	Description                          string                 `json:"description,omitempty"`
	Group                                Group                  `json:"group"`
	ID                                   int                    `json:"id,omitempty"`
	UnsubscribeFlag                      bool                   `json:"unsubscribeFlag,omitempty"`
}

type ContactInfo struct {
	Info               map[string]interface{}     `json:"_info,omitempty"`
	AddressLine1       string                     `json:"addressLine1,omitempty"`
	AddressLine2       string                     `json:"addressLine2,omitempty"`
	City               string                     `json:"city,omitempty"`
	CommunicationItems []ContactCommunicationItem `json:"communicationItems,omitempty"`
	Company            Company                    `json:"company,omitempty"`
	CompanyLocation    SystemLocation             `json:"companyLocation,omitempty"`
	Country            Country                    `json:"country,omitempty"`
	DefaultBillingFlag bool                       `json:"defaultBillingFlag,omitempty"`
	DefaultFlag        bool                       `json:"defaultFlag,omitempty"`
	DefaultPhoneNbr    string                     `json:"defaultPhoneNbr,omitempty"`
	DefaultPhoneType   string                     `json:"defaultPhoneType,omitempty"`
	Department         ContactDepartment          `json:"department,omitempty"`
	FacebookUrl        string                     `json:"facebookUrl,omitempty"`
	FirstName          string                     `json:"firstName,omitempty"`
	ID                 int                        `json:"id,omitempty"`
	InactiveFlag       bool                       `json:"inactiveFlag,omitempty"`
	LastName           string                     `json:"lastName,omitempty"`
	LinkedInUrl        string                     `json:"linkedInUrl,omitempty"`
	Site               Site                       `json:"site,omitempty"`
	State              string                     `json:"state,omitempty"`
	Title              string                     `json:"title,omitempty"`
	TwitterUrl         string                     `json:"twitterUrl,omitempty"`
	Types              []ContactType              `json:"types,omitempty"`
	Zip                string                     `json:"zip,omitempty"`
}

type ContactNote struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	ContactId int                    `json:"contactId,omitempty"`
	EnteredBy string                 `json:"enteredBy,omitempty"`
	Flagged   bool                   `json:"flagged,omitempty"`
	ID        int                    `json:"id,omitempty"`
	Text      string                 `json:"text"`
	Type      NoteType               `json:"type,omitempty"`
}

type ContactRelationship struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name"`
}

type ContactTrack struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	ActionRemaining int                    `json:"actionRemaining,omitempty"`
	ActionTaken     int                    `json:"actionTaken,omitempty"`
	Company         Company                `json:"company,omitempty"`
	Contact         Contact                `json:"contact,omitempty"`
	EndDate         string                 `json:"endDate,omitempty"`
	ID              int                    `json:"id,omitempty"`
	Name            string                 `json:"name,omitempty"`
	StartDate       string                 `json:"startDate,omitempty"`
	StartedBy       string                 `json:"startedBy,omitempty"`
	TrackId         int                    `json:"trackId,omitempty"`
}

type ContactType struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag         bool                   `json:"defaultFlag,omitempty"`
	Description         string                 `json:"description"`
	ID                  int                    `json:"id,omitempty"`
	ServiceAlertFlag    bool                   `json:"serviceAlertFlag,omitempty"`
	ServiceAlertMessage string                 `json:"serviceAlertMessage,omitempty"`
}

type ContactTypeInfo struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag         bool                   `json:"defaultFlag,omitempty"`
	Description         string                 `json:"description,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	ServiceAlertFlag    bool                   `json:"serviceAlertFlag,omitempty"`
	ServiceAlertMessage string                 `json:"serviceAlertMessage,omitempty"`
}

type Country struct {
	Info                   map[string]interface{} `json:"_info,omitempty"`
	AddressFormat          AddressFormat          `json:"addressFormat,omitempty"`
	CityCaption            string                 `json:"cityCaption,omitempty"`
	CoreEntityCountryCode  string                 `json:"coreEntityCountryCode,omitempty"`
	CountryCode            string                 `json:"countryCode,omitempty"`
	Currency               Currency               `json:"currency"`
	DefaultFlag            bool                   `json:"defaultFlag,omitempty"`
	DialingPrefix          string                 `json:"dialingPrefix,omitempty"`
	ID                     int                    `json:"id,omitempty"`
	LocalizationCaptionOne string                 `json:"localizationCaptionOne,omitempty"`
	LocalizationValueOne   string                 `json:"localizationValueOne,omitempty"`
	Name                   string                 `json:"name"`
	StateCaption           string                 `json:"stateCaption,omitempty"`
	ZipCaption             string                 `json:"zipCaption,omitempty"`
	ZipMinimumLength       int                    `json:"zipMinimumLength,omitempty"`
}

type CountryInfo struct {
	Info                   map[string]interface{} `json:"_info,omitempty"`
	CityCaption            string                 `json:"cityCaption,omitempty"`
	Currency               Currency               `json:"currency,omitempty"`
	DefaultFlag            bool                   `json:"defaultFlag,omitempty"`
	DialingPrefix          string                 `json:"dialingPrefix,omitempty"`
	ID                     int                    `json:"id,omitempty"`
	LocalizationCaptionOne string                 `json:"localizationCaptionOne,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	StateCaption           string                 `json:"stateCaption,omitempty"`
	ZipCaption             string                 `json:"zipCaption,omitempty"`
}

type EntityType struct {
	Code string `json:"code,omitempty"`
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EntityTypeInfo struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type M365Contact struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	AssignedLicenses    string                 `json:"assignedLicenses,omitempty"`
	ContactRecId        int                    `json:"contactRecId,omitempty"`
	Department          string                 `json:"department,omitempty"`
	DirectoryRoles      string                 `json:"directoryRoles,omitempty"`
	DisplayName         string                 `json:"displayName,omitempty"`
	EmployeeType        string                 `json:"employeeType,omitempty"`
	Groups              string                 `json:"groups,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	M365ContactId       string                 `json:"m365ContactId,omitempty"`
	ManagerId           string                 `json:"managerId,omitempty"`
	ProxyAddresses      string                 `json:"proxyAddresses,omitempty"`
	ProxyAddressesPlain string                 `json:"proxyAddressesPlain,omitempty"`
	TenantId            string                 `json:"tenantId,omitempty"`
	UserPrincipalName   string                 `json:"userPrincipalName,omitempty"`
}

type M365ContactSyncProperty struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	CompanyRecID       int                    `json:"companyRecID,omitempty"`
	ExcludeIncludeFlag bool                   `json:"excludeIncludeFlag,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	IncludeExcludeType string                 `json:"includeExcludeType,omitempty"`
	PropertyType       string                 `json:"propertyType,omitempty"`
	WildCard           string                 `json:"wildCard,omitempty"`
}

type ManagedDevicesIntegration struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	ConfigBillCustomerFlag  bool                   `json:"configBillCustomerFlag,omitempty"`
	DefaultBillingLevel     string                 `json:"defaultBillingLevel,omitempty"`
	DefaultDepartment       SystemDepartment       `json:"defaultDepartment,omitempty"`
	DefaultLocation         SystemLocation         `json:"defaultLocation,omitempty"`
	DisableNewCrosssFlag    bool                   `json:"disableNewCrosssFlag,omitempty"`
	GlobalLoginPassword     string                 `json:"globalLoginPassword,omitempty"`
	GlobalLoginUsername     string                 `json:"globalLoginUsername,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	IntegratorLogin         IntegratorLogin        `json:"integratorLogin,omitempty"`
	LoginBy                 string                 `json:"loginBy,omitempty"`
	ManagementItSetupType   string                 `json:"managementItSetupType,omitempty"`
	MatchOnSerialNumberFlag bool                   `json:"matchOnSerialNumberFlag,omitempty"`
	Name                    string                 `json:"name"`
	PortalUrl               string                 `json:"portalUrl,omitempty"`
	Solution                string                 `json:"solution"`
}

type ManagedDevicesIntegrationCross struct {
	Info                      map[string]interface{}    `json:"_info,omitempty"`
	AgreementType             AgreementType             `json:"agreementType,omitempty"`
	ConfigurationType         ConfigurationType         `json:"configurationType,omitempty"`
	ID                        int                       `json:"id,omitempty"`
	InactiveFlag              bool                      `json:"inactiveFlag,omitempty"`
	ManagedDevicesIntegration ManagedDevicesIntegration `json:"managedDevicesIntegration,omitempty"`
	Product                   IvItem                    `json:"product,omitempty"`
	VendorLevel               string                    `json:"vendorLevel,omitempty"`
	VendorType                string                    `json:"vendorType,omitempty"`
}

type ManagedDevicesIntegrationInfo struct {
	Info                  map[string]interface{} `json:"_info,omitempty"`
	ID                    int                    `json:"id,omitempty"`
	ManagementItSetupType string                 `json:"managementItSetupType,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	Solution              string                 `json:"solution,omitempty"`
}

type ManagedDevicesIntegrationLogin struct {
	Info                      map[string]interface{}    `json:"_info,omitempty"`
	ID                        int                       `json:"id,omitempty"`
	ManagedDevicesIntegration ManagedDevicesIntegration `json:"managedDevicesIntegration,omitempty"`
	Member                    Member                    `json:"member"`
	Password                  string                    `json:"password,omitempty"`
	Username                  string                    `json:"username"`
}

type ManagedDevicesIntegrationNotification struct {
	Info                      map[string]interface{}    `json:"_info,omitempty"`
	ID                        int                       `json:"id,omitempty"`
	LogType                   string                    `json:"logType,omitempty"`
	ManagedDevicesIntegration ManagedDevicesIntegration `json:"managedDevicesIntegration,omitempty"`
	Member                    Member                    `json:"member,omitempty"`
	NotifyWho                 NotificationRecipient     `json:"notifyWho"`
}

type ManagedInformation struct {
	ChildConfigurationsMatchingOn      string `json:"childConfigurationsMatchingOn,omitempty"`
	InactivateConfigurationsMatchingOn string `json:"inactivateConfigurationsMatchingOn,omitempty"`
	InactiveConfigurationStatusId      int    `json:"inactiveConfigurationStatusId,omitempty"`
	Level                              string `json:"level,omitempty"`
	ManagedIdentifier                  string `json:"managedIdentifier,omitempty"`
	ManagementSolutionName             string `json:"managementSolutionName,omitempty"`
	Type                               string `json:"type,omitempty"`
}

type Management struct {
	Info                                 map[string]interface{} `json:"_info,omitempty"`
	AddedConfigurationStatus             ConfigurationStatus    `json:"addedConfigurationStatus"`
	DeletedConfigurationStatus           ConfigurationStatus    `json:"deletedConfigurationStatus"`
	ExecutiveSummaryReportScheduleDay    int                    `json:"executiveSummaryReportScheduleDay,omitempty"`
	ExecutiveSummaryReportScheduleHour   int                    `json:"executiveSummaryReportScheduleHour,omitempty"`
	ExecutiveSummaryReportScheduleMinute int                    `json:"executiveSummaryReportScheduleMinute,omitempty"`
	ID                                   int                    `json:"id,omitempty"`
	IntegratorLogin                      IntegratorLogin        `json:"integratorLogin"`
	RunTime                              time.Time              `json:"runTime,omitempty"`
	ScheduleExecutiveSummaryReportFlag   bool                   `json:"scheduleExecutiveSummaryReportFlag,omitempty"`
}

type ManagementBackup struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	BillingLevel string                 `json:"billingLevel,omitempty"`
	ID           int                    `json:"id,omitempty"`
	Item         CatalogItem            `json:"item"`
	Type         AgreementType          `json:"type"`
}

type ManagementItSolution struct {
	Info                           map[string]interface{} `json:"_info,omitempty"`
	ContinuumApiPassword           string                 `json:"continuumApiPassword,omitempty"`
	ContinuumApiUsername           string                 `json:"continuumApiUsername,omitempty"`
	GlobalLoginFlag                bool                   `json:"globalLoginFlag,omitempty"`
	GlobalLoginPassword            string                 `json:"globalLoginPassword,omitempty"`
	GlobalLoginUsername            string                 `json:"globalLoginUsername,omitempty"`
	ID                             int                    `json:"id,omitempty"`
	LevelApiPassword               string                 `json:"levelApiPassword,omitempty"`
	LevelApiUsername               string                 `json:"levelApiUsername,omitempty"`
	LevelVarDomain                 string                 `json:"levelVarDomain,omitempty"`
	ManagementItSolutionType       string                 `json:"managementItSolutionType,omitempty"`
	ManagementServerUrl            string                 `json:"managementServerUrl,omitempty"`
	ManagementSolutionName         string                 `json:"managementSolutionName,omitempty"`
	NAblePassword                  string                 `json:"nAblePassword,omitempty"`
	NAbleUsername                  string                 `json:"nAbleUsername,omitempty"`
	Name                           string                 `json:"name"`
	NoDisplayFlag                  bool                   `json:"noDisplayFlag,omitempty"`
	OverrideLoginLocationFlag      bool                   `json:"overrideLoginLocationFlag,omitempty"`
	OverrideWebServiceLocationFlag bool                   `json:"overrideWebServiceLocationFlag,omitempty"`
	PortalOverrideLoginUrl         string                 `json:"portalOverrideLoginUrl,omitempty"`
	UsingSslFlag                   bool                   `json:"usingSslFlag,omitempty"`
	WebserviceOverrideUrl          string                 `json:"webserviceOverrideUrl,omitempty"`
}

type ManagementItSolutionAgreementInterfaceParameter struct {
	Info                      map[string]interface{}    `json:"_info,omitempty"`
	AgreementType             AgreementType             `json:"agreementType"`
	ID                        int                       `json:"id,omitempty"`
	ManagedDevicesIntegration ManagedDevicesIntegration `json:"managedDevicesIntegration,omitempty"`
	ServerProduct             IvItem                    `json:"serverProduct,omitempty"`
	SpamStatsProduct          IvItem                    `json:"spamStatsProduct,omitempty"`
	WorkstationProduct        IvItem                    `json:"workstationProduct,omitempty"`
}

type ManagementLogDocumentInfo struct {
	FileSize         string `json:"fileSize,omitempty"`
	FullPathFileName string `json:"fullPathFileName,omitempty"`
}

type ManagementReportNotification struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Company    Company                `json:"company,omitempty"`
	Email      string                 `json:"email,omitempty"`
	GlobalFlag bool                   `json:"globalFlag,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Member     Member                 `json:"member,omitempty"`
	NotifyWho  NotificationRecipient  `json:"notifyWho"`
}

type ManagementReportSetup struct {
	Info                        map[string]interface{} `json:"_info,omitempty"`
	ID                          int                    `json:"id,omitempty"`
	ScheduledReportDisabledFlag bool                   `json:"scheduledReportDisabledFlag"`
}

type MarketDescription struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name"`
}

type MarketDescriptionInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type OwnershipType struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name"`
}

type OwnershipTypeInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type PortalConfiguration struct {
	Info                       map[string]interface{} `json:"_info,omitempty"`
	AgreementTypeIds           []int                  `json:"agreementTypeIds,omitempty"`
	BoardIds                   []int                  `json:"boardIds,omitempty"`
	ButtonColor                string                 `json:"buttonColor,omitempty"`
	Company                    Company                `json:"company,omitempty"`
	ConfigTypeIds              []int                  `json:"configTypeIds,omitempty"`
	DefaultFlag                bool                   `json:"defaultFlag,omitempty"`
	DisplayVendorFlag          bool                   `json:"displayVendorFlag,omitempty"`
	HeaderColor                string                 `json:"headerColor,omitempty"`
	ID                         int                    `json:"id,omitempty"`
	Language                   string                 `json:"language,omitempty"`
	LocationIds                []int                  `json:"locationIds,omitempty"`
	LoginBackgroundColor       string                 `json:"loginBackgroundColor,omitempty"`
	MenuColor                  string                 `json:"menuColor,omitempty"`
	Name                       string                 `json:"name"`
	PortalBackgroundColor      string                 `json:"portalBackgroundColor,omitempty"`
	PortalImageCopySuccessFlag bool                   `json:"portalImageCopySuccessFlag,omitempty"`
	URL                        string                 `json:"url,omitempty"`
	WelcomeText                string                 `json:"welcomeText,omitempty"`
}

type PortalConfigurationInvoiceSetup struct {
	Info                map[string]interface{}              `json:"_info,omitempty"`
	AddAllStatuses      bool                                `json:"addAllStatuses,omitempty"`
	AllowInvPmtFlag     bool                                `json:"allowInvPmtFlag,omitempty"`
	BillingStatusIds    []int                               `json:"billingStatusIds,omitempty"`
	DisplayInvPmtFlag   bool                                `json:"displayInvPmtFlag,omitempty"`
	ID                  int                                 `json:"id,omitempty"`
	Location            SystemLocation                      `json:"location,omitempty"`
	Login               string                              `json:"login,omitempty"`
	Password            string                              `json:"password,omitempty"`
	PaymentProcessor    PortalConfigurationPaymentProcessor `json:"paymentProcessor,omitempty"`
	PortalConfiguration PortalConfiguration                 `json:"portalConfiguration,omitempty"`
	RemoveAllStatuses   bool                                `json:"removeAllStatuses,omitempty"`
	UrlOverride         string                              `json:"urlOverride,omitempty"`
}

type PortalConfigurationOpportunitySetup struct {
	Info                                               map[string]interface{} `json:"_info,omitempty"`
	AcceptanceChangeStatusFlag                         bool                   `json:"acceptanceChangeStatusFlag,omitempty"`
	AcceptanceCreateActivityFlag                       bool                   `json:"acceptanceCreateActivityFlag,omitempty"`
	AcceptanceEmailActivityType                        ActivityType           `json:"acceptanceEmailActivityType,omitempty"`
	AcceptanceEmailAssignedByMember                    Member                 `json:"acceptanceEmailAssignedByMember,omitempty"`
	AcceptanceEmailBody                                string                 `json:"acceptanceEmailBody,omitempty"`
	AcceptanceEmailFromFirstName                       string                 `json:"acceptanceEmailFromFirstName,omitempty"`
	AcceptanceEmailFromLastName                        string                 `json:"acceptanceEmailFromLastName,omitempty"`
	AcceptanceEmailSubject                             string                 `json:"acceptanceEmailSubject,omitempty"`
	AcceptanceFromEmail                                string                 `json:"acceptanceFromEmail,omitempty"`
	AcceptanceOpportunityStatus                        OpportunityStatus      `json:"acceptanceOpportunityStatus,omitempty"`
	AcceptanceSendEmailFlag                            bool                   `json:"acceptanceSendEmailFlag,omitempty"`
	AddAllOpportunityStatuses                          bool                   `json:"addAllOpportunityStatuses,omitempty"`
	AddAllOpportunityTypes                             bool                   `json:"addAllOpportunityTypes,omitempty"`
	ConfirmationEmailBody                              string                 `json:"confirmationEmailBody,omitempty"`
	ConfirmationEmailFromFirstName                     string                 `json:"confirmationEmailFromFirstName,omitempty"`
	ConfirmationEmailFromLastName                      string                 `json:"confirmationEmailFromLastName,omitempty"`
	ConfirmationEmailSubject                           string                 `json:"confirmationEmailSubject,omitempty"`
	ConfirmationEmailUseDefaultCompanyEmailAddressFlag bool                   `json:"confirmationEmailUseDefaultCompanyEmailAddressFlag,omitempty"`
	ConfirmationFromEmail                              string                 `json:"confirmationFromEmail,omitempty"`
	ConfirmationSendEmailFlag                          bool                   `json:"confirmationSendEmailFlag,omitempty"`
	ID                                                 int                    `json:"id,omitempty"`
	OpportunityStatusRecIDs                            []int                  `json:"opportunityStatusRecIDs,omitempty"`
	OpportunityTypeRecIDs                              []int                  `json:"opportunityTypeRecIDs,omitempty"`
	RejectionChangeStatusFlag                          bool                   `json:"rejectionChangeStatusFlag,omitempty"`
	RejectionCreateActivityFlag                        bool                   `json:"rejectionCreateActivityFlag,omitempty"`
	RejectionEmailActivityType                         ActivityType           `json:"rejectionEmailActivityType,omitempty"`
	RejectionEmailAssignedByMember                     Member                 `json:"rejectionEmailAssignedByMember,omitempty"`
	RejectionEmailBody                                 string                 `json:"rejectionEmailBody,omitempty"`
	RejectionEmailFromFirstName                        string                 `json:"rejectionEmailFromFirstName,omitempty"`
	RejectionEmailFromLastName                         string                 `json:"rejectionEmailFromLastName,omitempty"`
	RejectionEmailSubject                              string                 `json:"rejectionEmailSubject,omitempty"`
	RejectionFromEmail                                 string                 `json:"rejectionFromEmail,omitempty"`
	RejectionOpportunityStatus                         OpportunityStatus      `json:"rejectionOpportunityStatus,omitempty"`
	RejectionSendEmailFlag                             bool                   `json:"rejectionSendEmailFlag,omitempty"`
	RemoveAllOpportunityStatuses                       bool                   `json:"removeAllOpportunityStatuses,omitempty"`
	RemoveAllOpportunityTypes                          bool                   `json:"removeAllOpportunityTypes,omitempty"`
	RestrictViewByOpportunityStatusFlag                bool                   `json:"restrictViewByOpportunityStatusFlag,omitempty"`
	RestrictViewByOpportunityTypeFlag                  bool                   `json:"restrictViewByOpportunityTypeFlag,omitempty"`
}

type PortalConfigurationPasswordEmailSetup struct {
	Info                                   map[string]interface{} `json:"_info,omitempty"`
	ID                                     int                    `json:"id,omitempty"`
	InvalidPasswordEmailBody               string                 `json:"invalidPasswordEmailBody,omitempty"`
	InvalidPasswordEmailFromEmail          string                 `json:"invalidPasswordEmailFromEmail,omitempty"`
	InvalidPasswordEmailFromFirstName      string                 `json:"invalidPasswordEmailFromFirstName,omitempty"`
	InvalidPasswordEmailFromLastName       string                 `json:"invalidPasswordEmailFromLastName,omitempty"`
	InvalidPasswordEmailSubject            string                 `json:"invalidPasswordEmailSubject,omitempty"`
	InvalidPasswordEmailUseCustomEmailFlag bool                   `json:"invalidPasswordEmailUseCustomEmailFlag,omitempty"`
	ValidPasswordEmailBody                 string                 `json:"validPasswordEmailBody,omitempty"`
	ValidPasswordEmailFromEmail            string                 `json:"validPasswordEmailFromEmail,omitempty"`
	ValidPasswordEmailFromFirstName        string                 `json:"validPasswordEmailFromFirstName,omitempty"`
	ValidPasswordEmailFromLastName         string                 `json:"validPasswordEmailFromLastName,omitempty"`
	ValidPasswordEmailSubject              string                 `json:"validPasswordEmailSubject,omitempty"`
	ValidPasswordEmailUseCustomEmailFlag   bool                   `json:"validPasswordEmailUseCustomEmailFlag,omitempty"`
}

type PortalConfigurationPaymentProcessor struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	TestURL string `json:"testURL,omitempty"`
}

type PortalConfigurationProjectSetup struct {
	Info                            map[string]interface{} `json:"_info,omitempty"`
	BillingMethodFlag               bool                   `json:"billingMethodFlag,omitempty"`
	ContactFlag                     bool                   `json:"contactFlag,omitempty"`
	DescriptionFlag                 bool                   `json:"descriptionFlag,omitempty"`
	EstimatedEndFlag                bool                   `json:"estimatedEndFlag,omitempty"`
	EstimatedStartFlag              bool                   `json:"estimatedStartFlag,omitempty"`
	FixedFeeActualFinishFlag        bool                   `json:"fixedFeeActualFinishFlag,omitempty"`
	FixedFeeActualHrsFlag           bool                   `json:"fixedFeeActualHrsFlag,omitempty"`
	FixedFeeActualStartFlag         bool                   `json:"fixedFeeActualStartFlag,omitempty"`
	FixedFeeAssignedFlag            bool                   `json:"fixedFeeAssignedFlag,omitempty"`
	FixedFeeBillFlag                bool                   `json:"fixedFeeBillFlag,omitempty"`
	FixedFeeBudgetHrsFlag           bool                   `json:"fixedFeeBudgetHrsFlag,omitempty"`
	FixedFeeScheduledFinishFlag     bool                   `json:"fixedFeeScheduledFinishFlag,omitempty"`
	FixedFeeScheduledHrsFlag        bool                   `json:"fixedFeeScheduledHrsFlag,omitempty"`
	FixedFeeScheduledStartFlag      bool                   `json:"fixedFeeScheduledStartFlag,omitempty"`
	FixedFeeStatusFlag              bool                   `json:"fixedFeeStatusFlag,omitempty"`
	ID                              int                    `json:"id,omitempty"`
	LastUpdatedFlag                 bool                   `json:"lastUpdatedFlag,omitempty"`
	OnlyDisplay                     string                 `json:"onlyDisplay,omitempty"`
	PortalConfig                    PortalConfiguration    `json:"portalConfig,omitempty"`
	ProjectDetailTotalHoursFlag     bool                   `json:"projectDetailTotalHoursFlag,omitempty"`
	ProjectIssueActualFinishFlag    bool                   `json:"projectIssueActualFinishFlag,omitempty"`
	ProjectIssueActualHrsFlag       bool                   `json:"projectIssueActualHrsFlag,omitempty"`
	ProjectIssueActualStartFlag     bool                   `json:"projectIssueActualStartFlag,omitempty"`
	ProjectIssueAssignedFlag        bool                   `json:"projectIssueAssignedFlag,omitempty"`
	ProjectIssueBillFlag            bool                   `json:"projectIssueBillFlag,omitempty"`
	ProjectIssueBudgetHrsFlag       bool                   `json:"projectIssueBudgetHrsFlag,omitempty"`
	ProjectIssueScheduledFinishFlag bool                   `json:"projectIssueScheduledFinishFlag,omitempty"`
	ProjectIssueScheduledHrsFlag    bool                   `json:"projectIssueScheduledHrsFlag,omitempty"`
	ProjectIssueScheduledStartFlag  bool                   `json:"projectIssueScheduledStartFlag,omitempty"`
	ProjectIssueStatusFlag          bool                   `json:"projectIssueStatusFlag,omitempty"`
	ProjectManagerFlag              bool                   `json:"projectManagerFlag,omitempty"`
	ProjectNameFlag                 bool                   `json:"projectNameFlag,omitempty"`
	ProjectTypeFlag                 bool                   `json:"projectTypeFlag,omitempty"`
	StatusFlag                      bool                   `json:"statusFlag,omitempty"`
	TimeMaterialActualFinishFlag    bool                   `json:"timeMaterialActualFinishFlag,omitempty"`
	TimeMaterialActualHrsFlag       bool                   `json:"timeMaterialActualHrsFlag,omitempty"`
	TimeMaterialActualStartFlag     bool                   `json:"timeMaterialActualStartFlag,omitempty"`
	TimeMaterialAssignedFlag        bool                   `json:"timeMaterialAssignedFlag,omitempty"`
	TimeMaterialBillFlag            bool                   `json:"timeMaterialBillFlag,omitempty"`
	TimeMaterialBudgetHrsFlag       bool                   `json:"timeMaterialBudgetHrsFlag,omitempty"`
	TimeMaterialScheduledFinishFlag bool                   `json:"timeMaterialScheduledFinishFlag,omitempty"`
	TimeMaterialScheduledHrsFlag    bool                   `json:"timeMaterialScheduledHrsFlag,omitempty"`
	TimeMaterialScheduledStartFlag  bool                   `json:"timeMaterialScheduledStartFlag,omitempty"`
	TimeMaterialStatusFlag          bool                   `json:"timeMaterialStatusFlag,omitempty"`
}

type PortalConfigurationServiceSetup struct {
	Info                        map[string]interface{} `json:"_info,omitempty"`
	ActualHoursFlag             bool                   `json:"actualHoursFlag,omitempty"`
	ApprovalStatusFlag          bool                   `json:"approvalStatusFlag,omitempty"`
	AssignedResourcesFlag       bool                   `json:"assignedResourcesFlag,omitempty"`
	BudgetHoursFlag             bool                   `json:"budgetHoursFlag,omitempty"`
	ClosedTasksFlag             bool                   `json:"closedTasksFlag,omitempty"`
	ContactFlag                 bool                   `json:"contactFlag,omitempty"`
	DisplayClosedTicketsOption  string                 `json:"displayClosedTicketsOption,omitempty"`
	EnableChatAssistFlag        bool                   `json:"enableChatAssistFlag,omitempty"`
	EnteredDateFlag             bool                   `json:"enteredDateFlag,omitempty"`
	FixedFeeTicketTemplate      ServiceSignoff         `json:"fixedFeeTicketTemplate"`
	ID                          int                    `json:"id,omitempty"`
	LastUpdateFlag              bool                   `json:"lastUpdateFlag,omitempty"`
	OpenTasksFlag               bool                   `json:"openTasksFlag,omitempty"`
	RequiredDateFlag            bool                   `json:"requiredDateFlag,omitempty"`
	ServiceBoardFlag            bool                   `json:"serviceBoardFlag,omitempty"`
	ServiceSubTypeFlag          bool                   `json:"serviceSubTypeFlag,omitempty"`
	ServiceSubTypeItemFlag      bool                   `json:"serviceSubTypeItemFlag,omitempty"`
	ServiceTypeFlag             bool                   `json:"serviceTypeFlag,omitempty"`
	SiteNameFlag                bool                   `json:"siteNameFlag,omitempty"`
	SlaInfoFlag                 bool                   `json:"slaInfoFlag,omitempty"`
	StatusFlag                  bool                   `json:"statusFlag,omitempty"`
	TimeMaterialsTicketTemplate ServiceSignoff         `json:"timeMaterialsTicketTemplate"`
}

type PortalSecurity struct {
	Enabled    bool   `json:"enabled,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type PortalSecurityLevel struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	Caption           string                 `json:"caption,omitempty"`
	CaptionIdentifier string                 `json:"captionIdentifier,omitempty"`
	ID                int                    `json:"id,omitempty"`
	IsDefaultFlag     bool                   `json:"isDefaultFlag,omitempty"`
}

type PortalSecuritySetting struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	FunctionDescription string                 `json:"functionDescription,omitempty"`
	FunctionIdentifier  string                 `json:"functionIdentifier,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	LevelFive           bool                   `json:"levelFive,omitempty"`
	LevelFour           bool                   `json:"levelFour,omitempty"`
	LevelOne            bool                   `json:"levelOne,omitempty"`
	LevelSix            bool                   `json:"levelSix,omitempty"`
	LevelThree          bool                   `json:"levelThree,omitempty"`
	LevelTwo            bool                   `json:"levelTwo,omitempty"`
}

type RequestPasswordRequest struct {
	Email string `json:"email"`
}

type State struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Country    Country                `json:"country"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier"`
	Name       string                 `json:"name"`
}

type StateInfo struct {
	Info       map[string]interface{} `json:"_info,omitempty"`
	Country    Country                `json:"country,omitempty"`
	ID         int                    `json:"id,omitempty"`
	Identifier string                 `json:"identifier,omitempty"`
	Name       string                 `json:"name,omitempty"`
}

type TeamRole struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	AccountManagerFlag bool                   `json:"accountManagerFlag,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	Name               string                 `json:"name"`
	SalesFlag          bool                   `json:"salesFlag,omitempty"`
	TechFlag           bool                   `json:"techFlag,omitempty"`
}

type TeamRoleInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type Track struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	ID              int                    `json:"id,omitempty"`
	InactiveFlag    bool                   `json:"inactiveFlag,omitempty"`
	Name            string                 `json:"name"`
	NotifyActionIds []int                  `json:"notifyActionIds,omitempty"`
}

type TrackAction struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	ActivityStatus     ActivityStatus         `json:"activityStatus,omitempty"`
	ActivityType       ActivityType           `json:"activityType,omitempty"`
	AttachedTrack      Track                  `json:"attachedTrack,omitempty"`
	BccContact         Contact                `json:"bccContact,omitempty"`
	CcContact          Contact                `json:"ccContact,omitempty"`
	CompanyStatus      CompanyStatus          `json:"companyStatus,omitempty"`
	DaysToExecute      int                    `json:"daysToExecute,omitempty"`
	EmailFrom          string                 `json:"emailFrom,omitempty"`
	EmailRecipient     string                 `json:"emailRecipient,omitempty"`
	Group              Group                  `json:"group,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	Notes              string                 `json:"notes,omitempty"`
	NotifyFrom         NotificationRecipient  `json:"notifyFrom,omitempty"`
	NotifyType         string                 `json:"notifyType"`
	NotifyWho          NotificationRecipient  `json:"notifyWho,omitempty"`
	ServiceTemplate    ServiceTemplate        `json:"serviceTemplate,omitempty"`
	SpecificMemberFrom Member                 `json:"specificMemberFrom,omitempty"`
	SpecificMemberTo   Member                 `json:"specificMemberTo,omitempty"`
	Subject            string                 `json:"subject,omitempty"`
	Track              Track                  `json:"track,omitempty"`
}

type Usage struct {
	Count       int    `json:"count,omitempty"`
	Description string `json:"description,omitempty"`
	Hyperlink   string `json:"hyperlink,omitempty"`
	ID          int    `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	TypeKey     string `json:"typeKey,omitempty"`
}

type ValidatePortalRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidatePortalResponse struct {
	ContactId int  `json:"contactId,omitempty"`
	Success   bool `json:"success,omitempty"`
}
