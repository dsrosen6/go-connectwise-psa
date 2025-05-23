package types

import (
	"time"
)

type AdjustmentDetail struct {
	Info             map[string]interface{} `json:"_info,omitempty"`
	Adjustment       Adjustment             `json:"adjustment,omitempty"`
	CatalogItem      CatalogItem            `json:"catalogItem"`
	Description      string                 `json:"description,omitempty"`
	ID               int                    `json:"id,omitempty"`
	QuantityAdjusted int                    `json:"quantityAdjusted,omitempty"`
	QuantityOnHand   float64                `json:"quantityOnHand,omitempty"`
	SerialNumber     string                 `json:"serialNumber,omitempty"`
	UnitCost         float64                `json:"unitCost,omitempty"`
	Warehouse        Warehouse              `json:"warehouse"`
	WarehouseBin     WarehouseBin           `json:"warehouseBin"`
}

type AdjustmentType struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	AuditTrailFlag bool                   `json:"auditTrailFlag,omitempty"`
	CreatedBy      string                 `json:"createdBy,omitempty"`
	DateCreated    time.Time              `json:"dateCreated,omitempty"`
	ID             int                    `json:"id,omitempty"`
	Identifier     string                 `json:"identifier"`
	Name           string                 `json:"name,omitempty"`
}

type AdjustmentTypeInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type BulkResult struct {
	Info    map[string]interface{} `json:"_info,omitempty"`
	Payload []ResultInfo           `json:"payload,omitempty"`
}

type CatalogComponent struct {
	Info                   map[string]interface{} `json:"_info,omitempty"`
	CatalogItem            CatalogItem            `json:"catalogItem"`
	Cost                   float64                `json:"cost,omitempty"`
	HideDescriptionFlag    bool                   `json:"hideDescriptionFlag,omitempty"`
	HideExtendedPriceFlag  bool                   `json:"hideExtendedPriceFlag,omitempty"`
	HideItemIdentifierFlag bool                   `json:"hideItemIdentifierFlag,omitempty"`
	HidePriceFlag          bool                   `json:"hidePriceFlag,omitempty"`
	HideQuantityFlag       bool                   `json:"hideQuantityFlag,omitempty"`
	ID                     int                    `json:"id,omitempty"`
	ParentCatalogItem      CatalogItem            `json:"parentCatalogItem,omitempty"`
	Price                  float64                `json:"price,omitempty"`
	Quantity               float64                `json:"quantity,omitempty"`
	SequenceNumber         int                    `json:"sequenceNumber,omitempty"`
}

type CatalogInventory struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	CatalogItem   CatalogItem            `json:"catalogItem,omitempty"`
	ID            int                    `json:"id,omitempty"`
	OnHand        int                    `json:"onHand,omitempty"`
	SerialNumbers []OnHandSerialNumber   `json:"serialNumbers,omitempty"`
	Warehouse     Warehouse              `json:"warehouse,omitempty"`
	WarehouseBin  WarehouseBin           `json:"warehouseBin,omitempty"`
}

type CatalogItem struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	AgreementType           AgreementType          `json:"agreementType,omitempty"`
	AutoUpdateUnitCostFlag  bool                   `json:"autoUpdateUnitCostFlag,omitempty"`
	AutoUpdateUnitPriceFlag bool                   `json:"autoUpdateUnitPriceFlag,omitempty"`
	BillableOption          string                 `json:"billableOption,omitempty"`
	CalculatedCost          float64                `json:"calculatedCost,omitempty"`
	CalculatedCostFlag      bool                   `json:"calculatedCostFlag,omitempty"`
	CalculatedPrice         float64                `json:"calculatedPrice,omitempty"`
	CalculatedPriceFlag     bool                   `json:"calculatedPriceFlag,omitempty"`
	Category                ProductCategory        `json:"category,omitempty"`
	ConnectWiseID           string                 `json:"connectWiseID,omitempty"`
	Cost                    float64                `json:"cost,omitempty"`
	CustomFields            []CustomFieldValue     `json:"customFields,omitempty"`
	CustomerDescription     string                 `json:"customerDescription"`
	Description             string                 `json:"description"`
	DropShipFlag            bool                   `json:"dropShipFlag,omitempty"`
	EntityType              EntityType             `json:"entityType,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	Identifier              string                 `json:"identifier"`
	InactiveFlag            bool                   `json:"inactiveFlag,omitempty"`
	IntegrationXRef         string                 `json:"integrationXRef,omitempty"`
	Manufacturer            Manufacturer           `json:"manufacturer,omitempty"`
	ManufacturerPartNumber  string                 `json:"manufacturerPartNumber,omitempty"`
	MarkupFlag              bool                   `json:"markupFlag,omitempty"`
	MarkupPercentage        float64                `json:"markupPercentage,omitempty"`
	MinStockLevel           int                    `json:"minStockLevel,omitempty"`
	Notes                   string                 `json:"notes,omitempty"`
	PhaseProductFlag        bool                   `json:"phaseProductFlag,omitempty"`
	Price                   float64                `json:"price,omitempty"`
	PriceAttribute          string                 `json:"priceAttribute,omitempty"`
	ProductClass            string                 `json:"productClass,omitempty"`
	RecurringBillCycle      BillingCycle           `json:"recurringBillCycle,omitempty"`
	RecurringCost           float64                `json:"recurringCost,omitempty"`
	RecurringCycleType      string                 `json:"recurringCycleType,omitempty"`
	RecurringFlag           bool                   `json:"recurringFlag,omitempty"`
	RecurringOneTimeFlag    bool                   `json:"recurringOneTimeFlag,omitempty"`
	RecurringRevenue        float64                `json:"recurringRevenue,omitempty"`
	SerializedCostFlag      bool                   `json:"serializedCostFlag,omitempty"`
	SerializedFlag          bool                   `json:"serializedFlag,omitempty"`
	Sla                     SLA                    `json:"sla,omitempty"`
	SpecialOrderFlag        bool                   `json:"specialOrderFlag,omitempty"`
	Subcategory             ProductSubCategory     `json:"subcategory"`
	TaxableFlag             bool                   `json:"taxableFlag,omitempty"`
	Type                    ProductType            `json:"type"`
	UnitOfMeasure           UnitOfMeasure          `json:"unitOfMeasure,omitempty"`
	Vendor                  Company                `json:"vendor,omitempty"`
	VendorSku               string                 `json:"vendorSku,omitempty"`
}

type CatalogItemInfo struct {
	Info                   map[string]interface{} `json:"_info,omitempty"`
	BillableOption         string                 `json:"billableOption,omitempty"`
	Cost                   float64                `json:"cost,omitempty"`
	CustomerDescription    string                 `json:"customerDescription,omitempty"`
	Description            string                 `json:"description,omitempty"`
	DropShipFlag           bool                   `json:"dropShipFlag,omitempty"`
	ID                     int                    `json:"id,omitempty"`
	Identifier             string                 `json:"identifier,omitempty"`
	InactiveFlag           bool                   `json:"inactiveFlag,omitempty"`
	ManufacturerPartNumber string                 `json:"manufacturerPartNumber,omitempty"`
	Price                  float64                `json:"price,omitempty"`
	ProductClass           string                 `json:"productClass,omitempty"`
	SerializedCostFlag     bool                   `json:"serializedCostFlag,omitempty"`
	SpecialOrderFlag       bool                   `json:"specialOrderFlag,omitempty"`
	TaxableFlag            bool                   `json:"taxableFlag,omitempty"`
	VendorSku              string                 `json:"vendorSku,omitempty"`
}

type CatalogPricing struct {
	CatalogItem CatalogItem    `json:"catalogItem,omitempty"`
	Company     Company        `json:"company,omitempty"`
	Date        string         `json:"date,omitempty"`
	Location    SystemLocation `json:"location,omitempty"`
	Price       float64        `json:"price,omitempty"`
	Quantity    int            `json:"quantity,omitempty"`
}

type CatalogVendors struct {
	CatalogItemId     int    `json:"catalogItemId,omitempty"`
	ID                int    `json:"id,omitempty"`
	IsPreferredVendor bool   `json:"isPreferredVendor,omitempty"`
	VendorId          int    `json:"vendorId,omitempty"`
	VendorName        string `json:"vendorName,omitempty"`
	VendorSku         string `json:"vendorSku,omitempty"`
}

type Category struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	AddAllLocations    bool                   `json:"addAllLocations,omitempty"`
	DefaultFlag        bool                   `json:"defaultFlag,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	InactiveFlag       bool                   `json:"inactiveFlag,omitempty"`
	IntegrationXref    string                 `json:"integrationXref,omitempty"`
	LocationIds        []int                  `json:"locationIds,omitempty"`
	Name               string                 `json:"name"`
	PriceLevelXref     string                 `json:"priceLevelXref,omitempty"`
	RemoveAllLocations bool                   `json:"removeAllLocations,omitempty"`
}

type CategoryInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type ChangeOrder struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	PurchaseHeaderRecId int                    `json:"purchaseHeaderRecId"`
}

type Conversion struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	ID        int                    `json:"id,omitempty"`
	ParentUOM UnitOfMeasure          `json:"parentUOM,omitempty"`
	Quantity  float64                `json:"quantity,omitempty"`
	UomType   UnitOfMeasure          `json:"uomType"`
}

type DirectionalSync struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name"`
}

type HttpResponseMessage struct {
	Content             HttpContent         `json:"content,omitempty"`
	Headers             []interface{}       `json:"headers,omitempty"`
	IsSuccessStatusCode bool                `json:"isSuccessStatusCode,omitempty"`
	ReasonPhrase        string              `json:"reasonPhrase,omitempty"`
	RequestMessage      *HttpRequestMessage `json:"requestMessage,omitempty"`
	StatusCode          string              `json:"statusCode,omitempty"`
	Version             Version             `json:"version,omitempty"`
}

type IdCollection struct {
	Ids []int `json:"ids,omitempty"`
}

type InventoryOnHand struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	CatalogItem   CatalogItem            `json:"catalogItem,omitempty"`
	ID            int                    `json:"id,omitempty"`
	OnHand        int                    `json:"onHand,omitempty"`
	SerialNumbers []OnHandSerialNumber   `json:"serialNumbers,omitempty"`
	Warehouse     Warehouse              `json:"warehouse,omitempty"`
	WarehouseBin  WarehouseBin           `json:"warehouseBin,omitempty"`
}

type InvoiceGrouping struct {
	Info                          map[string]interface{} `json:"_info,omitempty"`
	CustomerDescription           string                 `json:"customerDescription"`
	GroupParentChildAdditionsFlag bool                   `json:"groupParentChildAdditionsFlag,omitempty"`
	ID                            int                    `json:"id,omitempty"`
	InactiveFlag                  bool                   `json:"inactiveFlag,omitempty"`
	Name                          string                 `json:"name"`
	ShowPriceFlag                 bool                   `json:"showPriceFlag,omitempty"`
	ShowSubItemsFlag              bool                   `json:"showSubItemsFlag,omitempty"`
}

type LegacySubCategory struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name"`
}

type LegacySubCategoryInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type Manufacturer struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name"`
}

type ManufacturerInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type MinimumStockByWarehouse struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	ID           int                    `json:"id,omitempty"`
	MinimumStock int                    `json:"minimumStock,omitempty"`
	Warehouse    Warehouse              `json:"warehouse"`
}

type OnHandSerialNumber struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	CatalogItem  CatalogItem            `json:"catalogItem,omitempty"`
	ID           int                    `json:"id,omitempty"`
	Serial       string                 `json:"serial,omitempty"`
	Warehouse    Warehouse              `json:"warehouse,omitempty"`
	WarehouseBin WarehouseBin           `json:"warehouseBin,omitempty"`
}

type PricingBreak struct {
	Info          map[string]interface{} `json:"_info,omitempty"`
	Amount        float64                `json:"amount,omitempty"`
	DetailId      int                    `json:"detailId,omitempty"`
	ID            int                    `json:"id,omitempty"`
	PriceMethod   string                 `json:"priceMethod,omitempty"`
	QuantityEnd   float64                `json:"quantityEnd,omitempty"`
	QuantityStart float64                `json:"quantityStart,omitempty"`
	Unlimited     bool                   `json:"unlimited,omitempty"`
}

type PricingDetail struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	Category    ProductCategory        `json:"category,omitempty"`
	EndDate     time.Time              `json:"endDate,omitempty"`
	ID          int                    `json:"id,omitempty"`
	NoEndDate   bool                   `json:"noEndDate,omitempty"`
	Product     CatalogItem            `json:"product,omitempty"`
	StartDate   time.Time              `json:"startDate"`
	SubCategory ProductSubCategory     `json:"subCategory,omitempty"`
}

type PricingSchedule struct {
	Info                   map[string]interface{} `json:"_info,omitempty"`
	Companies              []int                  `json:"companies,omitempty"`
	Currency               Currency               `json:"currency,omitempty"`
	DefaultFlag            bool                   `json:"defaultFlag,omitempty"`
	ID                     int                    `json:"id,omitempty"`
	InactiveFlag           bool                   `json:"inactiveFlag,omitempty"`
	Name                   string                 `json:"name"`
	RemoveAllCompaniesFlag bool                   `json:"removeAllCompaniesFlag,omitempty"`
	SetAllCompaniesFlag    bool                   `json:"setAllCompaniesFlag,omitempty"`
}

type ProcurementAdjustment struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	AdjustmentDetails []AdjustmentDetail     `json:"adjustmentDetails,omitempty"`
	ClosedBy          string                 `json:"closedBy,omitempty"`
	ClosedDate        time.Time              `json:"closedDate,omitempty"`
	ClosedFlag        bool                   `json:"closedFlag,omitempty"`
	ID                int                    `json:"id,omitempty"`
	Identifier        string                 `json:"identifier"`
	Notes             string                 `json:"notes,omitempty"`
	Reason            string                 `json:"reason,omitempty"`
	Type              AdjustmentType         `json:"type"`
}

type ProcurementSetting struct {
	Info                                     map[string]interface{} `json:"_info,omitempty"`
	AutoApprovePurchaseOrderFlag             bool                   `json:"autoApprovePurchaseOrderFlag,omitempty"`
	AutoClosePurchaseOrderFlag               bool                   `json:"autoClosePurchaseOrderFlag,omitempty"`
	AutoClosePurchaseOrderItemFlag           bool                   `json:"autoClosePurchaseOrderItemFlag,omitempty"`
	CostingMethod                            string                 `json:"costingMethod,omitempty"`
	DefaultProductTaxableFlag                bool                   `json:"defaultProductTaxableFlag,omitempty"`
	DisableAutoPickFlag                      bool                   `json:"disableAutoPickFlag,omitempty"`
	DisableCostUpdatesFlag                   bool                   `json:"disableCostUpdatesFlag,omitempty"`
	DisableNegativeInventoryFlag             bool                   `json:"disableNegativeInventoryFlag,omitempty"`
	EoriNumber                               string                 `json:"eoriNumber,omitempty"`
	ID                                       int                    `json:"id,omitempty"`
	NotificationForChangesInShippingInfoFlag bool                   `json:"notificationForChangesInShippingInfoFlag,omitempty"`
	NumDecimalPlaces                         int                    `json:"numDecimalPlaces,omitempty"`
	PrefixSuffixType                         string                 `json:"prefixSuffixType,omitempty"`
	PurchaseOrderPrefix                      string                 `json:"purchaseOrderPrefix,omitempty"`
	PurchaseOrderSuffix                      string                 `json:"purchaseOrderSuffix,omitempty"`
	ShippingInfoNotificationEmail            string                 `json:"shippingInfoNotificationEmail,omitempty"`
	StartingPurchaseOrderNum                 int                    `json:"startingPurchaseOrderNum"`
	TaxFreightFlag                           bool                   `json:"taxFreightFlag,omitempty"`
	TaxPurchaseOrderFlag                     bool                   `json:"taxPurchaseOrderFlag,omitempty"`
	UseVendorTaxCodeFlag                     bool                   `json:"useVendorTaxCodeFlag,omitempty"`
}

type ProductComponent struct {
	Info                   map[string]interface{} `json:"_info,omitempty"`
	CatalogItem            CatalogItem            `json:"catalogItem"`
	Cost                   float64                `json:"cost,omitempty"`
	HideDescriptionFlag    bool                   `json:"hideDescriptionFlag,omitempty"`
	HideExtendedPriceFlag  bool                   `json:"hideExtendedPriceFlag,omitempty"`
	HideItemIdentifierFlag bool                   `json:"hideItemIdentifierFlag,omitempty"`
	HidePriceFlag          bool                   `json:"hidePriceFlag,omitempty"`
	HideQuantityFlag       bool                   `json:"hideQuantityFlag,omitempty"`
	ID                     int                    `json:"id,omitempty"`
	ParentProductItem      ProductItem            `json:"parentProductItem,omitempty"`
	Price                  float64                `json:"price,omitempty"`
	ProductItem            ProductItem            `json:"productItem,omitempty"`
	Quantity               float64                `json:"quantity,omitempty"`
	SequenceNumber         int                    `json:"sequenceNumber,omitempty"`
	Vendor                 Company                `json:"vendor,omitempty"`
}

type ProductDetach struct {
	RemoveFromInvoice     bool `json:"removeFromInvoice,omitempty"`
	RemoveFromOpportunity bool `json:"removeFromOpportunity,omitempty"`
	RemoveFromProject     bool `json:"removeFromProject,omitempty"`
	RemoveFromSalesOrder  bool `json:"removeFromSalesOrder,omitempty"`
	RemoveFromTicket      bool `json:"removeFromTicket,omitempty"`
}

type ProductItem struct {
	Info                       map[string]interface{} `json:"_info,omitempty"`
	AddComponentsFlag          bool                   `json:"addComponentsFlag,omitempty"`
	Agreement                  Agreement              `json:"agreement,omitempty"`
	AgreementAmount            float64                `json:"agreementAmount,omitempty"`
	AsioSubscriptionsID        string                 `json:"asioSubscriptionsID,omitempty"`
	BillableOption             string                 `json:"billableOption,omitempty"`
	BusinessUnit               BillingUnit            `json:"businessUnit,omitempty"`
	BusinessUnitId             int                    `json:"businessUnitId,omitempty"`
	BypassForecastUpdate       bool                   `json:"bypassForecastUpdate,omitempty"`
	CalculatedCost             float64                `json:"calculatedCost,omitempty"`
	CalculatedCostFlag         bool                   `json:"calculatedCostFlag,omitempty"`
	CalculatedPrice            float64                `json:"calculatedPrice,omitempty"`
	CalculatedPriceFlag        bool                   `json:"calculatedPriceFlag,omitempty"`
	CancelledBy                int                    `json:"cancelledBy,omitempty"`
	CancelledDate              time.Time              `json:"cancelledDate,omitempty"`
	CancelledFlag              bool                   `json:"cancelledFlag,omitempty"`
	CancelledReason            string                 `json:"cancelledReason,omitempty"`
	CatalogItem                CatalogItem            `json:"catalogItem"`
	Company                    Company                `json:"company,omitempty"`
	Cost                       float64                `json:"cost,omitempty"`
	CustomFields               []CustomFieldValue     `json:"customFields,omitempty"`
	CustomerDescription        string                 `json:"customerDescription,omitempty"`
	Description                string                 `json:"description,omitempty"`
	Discount                   float64                `json:"discount,omitempty"`
	DropshipFlag               bool                   `json:"dropshipFlag,omitempty"`
	EntityType                 EntityType             `json:"entityType,omitempty"`
	ExtCost                    float64                `json:"extCost,omitempty"`
	ExtPrice                   float64                `json:"extPrice,omitempty"`
	ForecastDetailId           int                    `json:"forecastDetailId,omitempty"`
	ForecastStatus             OpportunityStatus      `json:"forecastStatus,omitempty"`
	ID                         int                    `json:"id,omitempty"`
	IgnorePricingSchedulesFlag bool                   `json:"ignorePricingSchedulesFlag,omitempty"`
	IntegrationXRef            string                 `json:"integrationXRef,omitempty"`
	InternalNotes              string                 `json:"internalNotes,omitempty"`
	Invoice                    Invoice                `json:"invoice,omitempty"`
	InvoiceGrouping            InvoiceGrouping        `json:"invoiceGrouping,omitempty"`
	ListPrice                  float64                `json:"listPrice,omitempty"`
	Location                   SystemLocation         `json:"location,omitempty"`
	LocationId                 int                    `json:"locationId,omitempty"`
	Margin                     float64                `json:"margin,omitempty"`
	MinimumStockFlag           bool                   `json:"minimumStockFlag,omitempty"`
	NeedToOrderQuantity        int                    `json:"needToOrderQuantity,omitempty"`
	NeedToPurchaseFlag         bool                   `json:"needToPurchaseFlag,omitempty"`
	Opportunity                Opportunity            `json:"opportunity,omitempty"`
	Phase                      ProjectPhase           `json:"phase,omitempty"`
	PhaseProductFlag           bool                   `json:"phaseProductFlag,omitempty"`
	PoApprovedFlag             bool                   `json:"poApprovedFlag,omitempty"`
	Price                      float64                `json:"price,omitempty"`
	PriceMethod                string                 `json:"priceMethod,omitempty"`
	ProductClass               string                 `json:"productClass,omitempty"`
	ProductSuppliedFlag        bool                   `json:"productSuppliedFlag,omitempty"`
	Project                    Project                `json:"project,omitempty"`
	PurchaseDate               time.Time              `json:"purchaseDate,omitempty"`
	Quantity                   float64                `json:"quantity,omitempty"`
	QuantityCancelled          float64                `json:"quantityCancelled,omitempty"`
	Recurring                  ProductRecurring       `json:"recurring,omitempty"`
	SalesOrder                 SalesOrder             `json:"salesOrder,omitempty"`
	SequenceNumber             float64                `json:"sequenceNumber,omitempty"`
	SerialNumberIds            []int                  `json:"serialNumberIds,omitempty"`
	SerialNumbers              []string               `json:"serialNumbers,omitempty"`
	ShipSet                    string                 `json:"shipSet,omitempty"`
	Sla                        SLA                    `json:"sla,omitempty"`
	SpecialOrderFlag           bool                   `json:"specialOrderFlag,omitempty"`
	SubContractorAmountLimit   float64                `json:"subContractorAmountLimit,omitempty"`
	SubContractorShipToId      int                    `json:"subContractorShipToId,omitempty"`
	TaxCode                    TaxCode                `json:"taxCode,omitempty"`
	TaxableFlag                bool                   `json:"taxableFlag,omitempty"`
	Ticket                     Ticket                 `json:"ticket,omitempty"`
	UnitOfMeasure              UnitOfMeasure          `json:"unitOfMeasure,omitempty"`
	Uom                        string                 `json:"uom,omitempty"`
	Vendor                     Company                `json:"vendor,omitempty"`
	VendorSku                  string                 `json:"vendorSku,omitempty"`
	Warehouse                  string                 `json:"warehouse,omitempty"`
	WarehouseBin               string                 `json:"warehouseBin,omitempty"`
	WarehouseBinId             int                    `json:"warehouseBinId,omitempty"`
	WarehouseBinIdObject       WarehouseBin           `json:"warehouseBinIdObject,omitempty"`
	WarehouseId                int                    `json:"warehouseId,omitempty"`
	WarehouseIdObject          Warehouse              `json:"warehouseIdObject,omitempty"`
}

type ProductPickingShippingDetail struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	ExpectedArrivalDate time.Time              `json:"expectedArrivalDate,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	LineNumber          int                    `json:"lineNumber,omitempty"`
	PickedQuantity      int                    `json:"pickedQuantity,omitempty"`
	ProductItem         ProductItem            `json:"productItem,omitempty"`
	Quantity            int                    `json:"quantity,omitempty"`
	SerialNumber        string                 `json:"serialNumber,omitempty"`
	SerialNumberIds     []int                  `json:"serialNumberIds,omitempty"`
	ShipmentDate        time.Time              `json:"shipmentDate,omitempty"`
	ShipmentMethod      ShipmentMethod         `json:"shipmentMethod,omitempty"`
	ShippedQuantity     int                    `json:"shippedQuantity,omitempty"`
	TrackingNumber      string                 `json:"trackingNumber,omitempty"`
	Warehouse           Warehouse              `json:"warehouse"`
	WarehouseBin        WarehouseBin           `json:"warehouseBin"`
}

type ProductType struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name"`
	TypeXref     string                 `json:"typeXref,omitempty"`
}

type ProductTypeInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type PurchaseOrder struct {
	Info                    map[string]interface{} `json:"_info,omitempty"`
	BusinessUnit            BillingUnit            `json:"businessUnit,omitempty"`
	BusinessUnitId          int                    `json:"businessUnitId,omitempty"`
	CancelReason            string                 `json:"cancelReason,omitempty"`
	ClosedBy                string                 `json:"closedBy,omitempty"`
	ClosedFlag              bool                   `json:"closedFlag,omitempty"`
	Currency                Currency               `json:"currency,omitempty"`
	CustomFields            []CustomFieldValue     `json:"customFields,omitempty"`
	CustomerCity            string                 `json:"customerCity,omitempty"`
	CustomerCompany         Company                `json:"customerCompany,omitempty"`
	CustomerContact         Contact                `json:"customerContact,omitempty"`
	CustomerCountry         Country                `json:"customerCountry,omitempty"`
	CustomerExtension       string                 `json:"customerExtension,omitempty"`
	CustomerName            string                 `json:"customerName,omitempty"`
	CustomerPhone           string                 `json:"customerPhone,omitempty"`
	CustomerSite            Site                   `json:"customerSite,omitempty"`
	CustomerSiteName        string                 `json:"customerSiteName,omitempty"`
	CustomerState           string                 `json:"customerState,omitempty"`
	CustomerStreetLine1     string                 `json:"customerStreetLine1,omitempty"`
	CustomerStreetLine2     string                 `json:"customerStreetLine2,omitempty"`
	CustomerZip             string                 `json:"customerZip,omitempty"`
	DateClosed              time.Time              `json:"dateClosed,omitempty"`
	DropShipCustomerFlag    bool                   `json:"dropShipCustomerFlag,omitempty"`
	EnteredBy               string                 `json:"enteredBy,omitempty"`
	FreightCost             float64                `json:"freightCost,omitempty"`
	FreightPackingSlip      string                 `json:"freightPackingSlip,omitempty"`
	FreightTaxTotal         float64                `json:"freightTaxTotal,omitempty"`
	ID                      int                    `json:"id,omitempty"`
	InternalNotes           string                 `json:"internalNotes,omitempty"`
	Location                SystemLocation         `json:"location,omitempty"`
	LocationId              int                    `json:"locationId,omitempty"`
	PoDate                  time.Time              `json:"poDate,omitempty"`
	PoNumber                string                 `json:"poNumber,omitempty"`
	SalesTax                float64                `json:"salesTax,omitempty"`
	ShipmentDate            time.Time              `json:"shipmentDate,omitempty"`
	ShipmentMethod          ShipmentMethod         `json:"shipmentMethod,omitempty"`
	ShippingInstructions    string                 `json:"shippingInstructions,omitempty"`
	Status                  PurchaseOrderStatus    `json:"status"`
	SubTotal                float64                `json:"subTotal,omitempty"`
	TaxCode                 TaxCode                `json:"taxCode,omitempty"`
	TaxFreightFlag          bool                   `json:"taxFreightFlag,omitempty"`
	TaxPoFlag               bool                   `json:"taxPoFlag,omitempty"`
	Terms                   BillingTerms           `json:"terms"`
	Total                   float64                `json:"total,omitempty"`
	TrackingNumber          string                 `json:"trackingNumber,omitempty"`
	UpdateShipmentInfo      bool                   `json:"updateShipmentInfo,omitempty"`
	UpdateVendorOrderNumber bool                   `json:"updateVendorOrderNumber,omitempty"`
	VendorCompany           Company                `json:"vendorCompany"`
	VendorContact           Contact                `json:"vendorContact,omitempty"`
	VendorInvoiceDate       time.Time              `json:"vendorInvoiceDate,omitempty"`
	VendorInvoiceNumber     string                 `json:"vendorInvoiceNumber,omitempty"`
	VendorOrderNumber       string                 `json:"vendorOrderNumber,omitempty"`
	VendorSite              Site                   `json:"vendorSite,omitempty"`
	Warehouse               Warehouse              `json:"warehouse,omitempty"`
	WarehouseContact        Contact                `json:"warehouseContact,omitempty"`
}

type PurchaseOrderInfo struct {
	Info     map[string]interface{} `json:"_info,omitempty"`
	Currency Currency               `json:"currency,omitempty"`
	ID       int                    `json:"id,omitempty"`
}

type PurchaseOrderLineItem struct {
	Info                     map[string]interface{} `json:"_info,omitempty"`
	BackorderedFlag          bool                   `json:"backorderedFlag,omitempty"`
	BatchedFlag              bool                   `json:"batchedFlag,omitempty"`
	CanceledBy               string                 `json:"canceledBy,omitempty"`
	CanceledFlag             bool                   `json:"canceledFlag,omitempty"`
	CanceledReason           string                 `json:"canceledReason,omitempty"`
	ClosedFlag               bool                   `json:"closedFlag,omitempty"`
	CustomFields             []CustomFieldValue     `json:"customFields,omitempty"`
	DateCanceled             time.Time              `json:"dateCanceled,omitempty"`
	DateCanceledUtc          time.Time              `json:"dateCanceledUtc,omitempty"`
	DateReceived             time.Time              `json:"dateReceived,omitempty"`
	Description              string                 `json:"description"`
	DisplayInternalNotesFlag bool                   `json:"displayInternalNotesFlag,omitempty"`
	ExpectedArrivalDate      time.Time              `json:"expectedArrivalDate,omitempty"`
	ExpectedShipDate         time.Time              `json:"expectedShipDate,omitempty"`
	ExtCost                  float64                `json:"extCost,omitempty"`
	ID                       int                    `json:"id,omitempty"`
	InternalNotes            string                 `json:"internalNotes,omitempty"`
	IsDetachAvailable        bool                   `json:"isDetachAvailable,omitempty"`
	LineNumber               int                    `json:"lineNumber,omitempty"`
	PackingSlip              string                 `json:"packingSlip,omitempty"`
	Product                  IvItem                 `json:"product"`
	PurchaseOrderId          int                    `json:"purchaseOrderId,omitempty"`
	PurchaseOrderNumber      string                 `json:"purchaseOrderNumber,omitempty"`
	Quantity                 float64                `json:"quantity,omitempty"`
	ReceivedQuantity         int                    `json:"receivedQuantity,omitempty"`
	ReceivedStatus           string                 `json:"receivedStatus,omitempty"`
	SalesOrder               []SalesOrder           `json:"salesOrder,omitempty"`
	SerialNumbers            string                 `json:"serialNumbers,omitempty"`
	ShipDate                 time.Time              `json:"shipDate,omitempty"`
	ShipSet                  string                 `json:"shipSet,omitempty"`
	ShipmentMethod           ShipmentMethod         `json:"shipmentMethod,omitempty"`
	Tax                      float64                `json:"tax,omitempty"`
	TrackingNumber           string                 `json:"trackingNumber,omitempty"`
	UnbatchedRecId           int                    `json:"unbatchedRecId,omitempty"`
	UnitCost                 float64                `json:"unitCost,omitempty"`
	UnitOfMeasure            UnitOfMeasure          `json:"unitOfMeasure"`
	VendorOrderNumber        string                 `json:"vendorOrderNumber,omitempty"`
	VendorSku                string                 `json:"vendorSku,omitempty"`
	Warehouse                Warehouse              `json:"warehouse,omitempty"`
	WarehouseBin             WarehouseBin           `json:"warehouseBin,omitempty"`
}

type PurchaseOrderNote struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	EnteredBy           string                 `json:"enteredBy,omitempty"`
	Flagged             bool                   `json:"flagged,omitempty"`
	ID                  int                    `json:"id,omitempty"`
	PurchaseHeaderRecID int                    `json:"purchaseHeaderRecID,omitempty"`
	Text                string                 `json:"text,omitempty"`
	Type                NoteType               `json:"type,omitempty"`
}

type PurchaseOrderStatus struct {
	Info              map[string]interface{}            `json:"_info,omitempty"`
	ClosedFlag        bool                              `json:"closedFlag,omitempty"`
	DefaultClosedFlag bool                              `json:"defaultClosedFlag,omitempty"`
	DefaultFlag       bool                              `json:"defaultFlag,omitempty"`
	EmailTemplate     *PurchaseOrderStatusEmailTemplate `json:"emailTemplate,omitempty"`
	ID                int                               `json:"id,omitempty"`
	InactiveFlag      bool                              `json:"inactiveFlag,omitempty"`
	Name              string                            `json:"name"`
	SortOrder         int                               `json:"sortOrder,omitempty"`
}

type PurchaseOrderStatusEmailTemplate struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	Body           string                 `json:"body,omitempty"`
	CopySenderFlag bool                   `json:"copySenderFlag,omitempty"`
	EmailAddress   string                 `json:"emailAddress,omitempty"`
	FirstName      string                 `json:"firstName,omitempty"`
	ID             int                    `json:"id,omitempty"`
	LastName       string                 `json:"lastName,omitempty"`
	Status         PurchaseOrderStatus    `json:"status,omitempty"`
	Subject        string                 `json:"subject"`
	UseSenderFlag  bool                   `json:"useSenderFlag,omitempty"`
}

type PurchaseOrderStatusInfo struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	ClosedFlag        bool                   `json:"closedFlag,omitempty"`
	DefaultClosedFlag bool                   `json:"defaultClosedFlag,omitempty"`
	DefaultFlag       bool                   `json:"defaultFlag,omitempty"`
	ID                int                    `json:"id,omitempty"`
	InactiveFlag      bool                   `json:"inactiveFlag,omitempty"`
	Name              string                 `json:"name,omitempty"`
	SortOrder         int                    `json:"sortOrder,omitempty"`
}

type PurchaseOrderStatusNotification struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	Email        string                 `json:"email,omitempty"`
	ID           int                    `json:"id,omitempty"`
	Member       Member                 `json:"member,omitempty"`
	NotifyWho    NotificationRecipient  `json:"notifyWho"`
	Status       PurchaseOrderStatus    `json:"status,omitempty"`
	WorkflowStep int                    `json:"workflowStep,omitempty"`
}

type PurchasingDemand struct {
	Products      []ProductDemand `json:"products,omitempty"`
	PurchaseOrder PurchaseOrder   `json:"purchaseOrder,omitempty"`
	Vendor        Company         `json:"vendor,omitempty"`
	Warehouse     Warehouse       `json:"warehouse,omitempty"`
}

type RmaAction struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name"`
}

type RmaActionInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type RmaDisposition struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name"`
}

type RmaDispositionInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type RmaStatus struct {
	Info          map[string]interface{}  `json:"_info,omitempty"`
	ClosedFlag    bool                    `json:"closedFlag,omitempty"`
	DefaultFlag   bool                    `json:"defaultFlag,omitempty"`
	EmailTemplate *RmaStatusEmailTemplate `json:"emailTemplate,omitempty"`
	ID            int                     `json:"id,omitempty"`
	Name          string                  `json:"name"`
	SortOrder     int                     `json:"sortOrder,omitempty"`
}

type RmaStatusEmailTemplate struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	Body           string                 `json:"body"`
	CopySenderFlag bool                   `json:"copySenderFlag,omitempty"`
	EmailAddress   string                 `json:"emailAddress,omitempty"`
	FirstName      string                 `json:"firstName,omitempty"`
	ID             int                    `json:"id,omitempty"`
	LastName       string                 `json:"lastName,omitempty"`
	Status         RmaStatus              `json:"status,omitempty"`
	Subject        string                 `json:"subject"`
	UseSenderFlag  bool                   `json:"useSenderFlag,omitempty"`
}

type RmaStatusInfo struct {
	Info      map[string]interface{} `json:"_info,omitempty"`
	ID        int                    `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	SortOrder int                    `json:"sortOrder,omitempty"`
}

type RmaStatusNotification struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	Email        string                 `json:"email,omitempty"`
	ID           int                    `json:"id,omitempty"`
	Member       Member                 `json:"member,omitempty"`
	NotifyWho    NotificationRecipient  `json:"notifyWho"`
	Status       RmaStatus              `json:"status,omitempty"`
	WorkflowStep int                    `json:"workflowStep,omitempty"`
}

type RmaTag struct {
	Info                         map[string]interface{} `json:"_info,omitempty"`
	AccountManager               Member                 `json:"accountManager,omitempty"`
	ClosedBy                     Member                 `json:"closedBy,omitempty"`
	ClosingNotes                 string                 `json:"closingNotes,omitempty"`
	Currency                     Currency               `json:"currency,omitempty"`
	CustomFields                 []CustomFieldValue     `json:"customFields,omitempty"`
	DateClosed                   string                 `json:"dateClosed,omitempty"`
	Department                   SystemDepartment       `json:"department"`
	DropShipFlag                 bool                   `json:"dropShipFlag,omitempty"`
	ID                           int                    `json:"id,omitempty"`
	InternalNotes                string                 `json:"internalNotes,omitempty"`
	Invoice                      Invoice                `json:"invoice,omitempty"`
	IvDescription                string                 `json:"ivDescription,omitempty"`
	ListPrice                    float64                `json:"listPrice,omitempty"`
	Location                     SystemLocation         `json:"location"`
	MfgItemID                    string                 `json:"mfgItemID,omitempty"`
	ProblemDescription           string                 `json:"problemDescription,omitempty"`
	Product                      IvItem                 `json:"product"`
	ProductDescription           string                 `json:"productDescription"`
	Project                      Project                `json:"project,omitempty"`
	PurchasedCompany             Company                `json:"purchasedCompany,omitempty"`
	PurchasedContact             Contact                `json:"purchasedContact,omitempty"`
	PurchasedContactAddressLine1 string                 `json:"purchasedContactAddressLine1,omitempty"`
	PurchasedContactAddressLine2 string                 `json:"purchasedContactAddressLine2,omitempty"`
	PurchasedContactCity         string                 `json:"purchasedContactCity,omitempty"`
	PurchasedContactCountry      Country                `json:"purchasedContactCountry,omitempty"`
	PurchasedContactEmail        string                 `json:"purchasedContactEmail,omitempty"`
	PurchasedContactExtension    string                 `json:"purchasedContactExtension,omitempty"`
	PurchasedContactPhone        string                 `json:"purchasedContactPhone,omitempty"`
	PurchasedContactState        string                 `json:"purchasedContactState,omitempty"`
	PurchasedContactType         string                 `json:"purchasedContactType,omitempty"`
	PurchasedContactZip          string                 `json:"purchasedContactZip,omitempty"`
	PurchasedInvoiceDate         string                 `json:"purchasedInvoiceDate,omitempty"`
	PurchasedInvoiceNumber       string                 `json:"purchasedInvoiceNumber,omitempty"`
	PurchasedNotes               string                 `json:"purchasedNotes,omitempty"`
	PurchasedOrderNumber         string                 `json:"purchasedOrderNumber,omitempty"`
	PurchasedSite                Site                   `json:"purchasedSite,omitempty"`
	PurchasedVendorAction        RmaAction              `json:"purchasedVendorAction,omitempty"`
	PurchasedVendorRmaNumber     string                 `json:"purchasedVendorRmaNumber,omitempty"`
	RepairCompany                Company                `json:"repairCompany,omitempty"`
	RepairContact                Contact                `json:"repairContact,omitempty"`
	RepairContactAddressLine1    string                 `json:"repairContactAddressLine1,omitempty"`
	RepairContactAddressLine2    string                 `json:"repairContactAddressLine2,omitempty"`
	RepairContactCity            string                 `json:"repairContactCity,omitempty"`
	RepairContactCountry         Country                `json:"repairContactCountry,omitempty"`
	RepairContactEmail           string                 `json:"repairContactEmail,omitempty"`
	RepairContactExtension       string                 `json:"repairContactExtension,omitempty"`
	RepairContactPhone           string                 `json:"repairContactPhone,omitempty"`
	RepairContactState           string                 `json:"repairContactState,omitempty"`
	RepairContactType            string                 `json:"repairContactType,omitempty"`
	RepairContactZip             string                 `json:"repairContactZip,omitempty"`
	RepairNotes                  string                 `json:"repairNotes,omitempty"`
	RepairOrderNumber            string                 `json:"repairOrderNumber,omitempty"`
	RepairSite                   Site                   `json:"repairSite,omitempty"`
	ReturnedCompany              Company                `json:"returnedCompany"`
	ReturnedContact              Contact                `json:"returnedContact,omitempty"`
	ReturnedContactAddressLine1  string                 `json:"returnedContactAddressLine1,omitempty"`
	ReturnedContactAddressLine2  string                 `json:"returnedContactAddressLine2,omitempty"`
	ReturnedContactCity          string                 `json:"returnedContactCity,omitempty"`
	ReturnedContactCountry       Country                `json:"returnedContactCountry,omitempty"`
	ReturnedContactEmail         string                 `json:"returnedContactEmail,omitempty"`
	ReturnedContactExtension     string                 `json:"returnedContactExtension,omitempty"`
	ReturnedContactPhone         string                 `json:"returnedContactPhone,omitempty"`
	ReturnedContactState         string                 `json:"returnedContactState,omitempty"`
	ReturnedContactType          string                 `json:"returnedContactType,omitempty"`
	ReturnedContactZip           string                 `json:"returnedContactZip,omitempty"`
	ReturnedSite                 Site                   `json:"returnedSite,omitempty"`
	RmaDisposition               RmaDisposition         `json:"rmaDisposition"`
	SalesOrder                   SalesOrder             `json:"salesOrder,omitempty"`
	SerialNumber                 string                 `json:"serialNumber,omitempty"`
	ServiceTicket                Ticket                 `json:"serviceTicket,omitempty"`
	ShipMethod                   ShipmentMethod         `json:"shipMethod,omitempty"`
	ShippingDate                 string                 `json:"shippingDate,omitempty"`
	ShippingTrackingNumber       string                 `json:"shippingTrackingNumber,omitempty"`
	Status                       RmaStatus              `json:"status"`
	Summary                      string                 `json:"summary,omitempty"`
	TechnicalContact             Member                 `json:"technicalContact,omitempty"`
	UnitPrice                    float64                `json:"unitPrice,omitempty"`
	WarrantyCompany              Company                `json:"warrantyCompany,omitempty"`
	WarrantyContact              Contact                `json:"warrantyContact,omitempty"`
	WarrantyContactAddressLine1  string                 `json:"warrantyContactAddressLine1,omitempty"`
	WarrantyContactAddressLine2  string                 `json:"warrantyContactAddressLine2,omitempty"`
	WarrantyContactCity          string                 `json:"warrantyContactCity,omitempty"`
	WarrantyContactCountry       Country                `json:"warrantyContactCountry,omitempty"`
	WarrantyContactEmail         string                 `json:"warrantyContactEmail,omitempty"`
	WarrantyContactExtension     string                 `json:"warrantyContactExtension,omitempty"`
	WarrantyContactPhone         string                 `json:"warrantyContactPhone,omitempty"`
	WarrantyContactState         string                 `json:"warrantyContactState,omitempty"`
	WarrantyContactType          string                 `json:"warrantyContactType,omitempty"`
	WarrantyContactZip           string                 `json:"warrantyContactZip,omitempty"`
	WarrantyNotes                string                 `json:"warrantyNotes,omitempty"`
	WarrantySite                 Site                   `json:"warrantySite,omitempty"`
}

type ShipmentMethod struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name"`
	TrackingUrl string                 `json:"trackingUrl,omitempty"`
}

type ShipmentMethodInfo struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	TrackingUrl string                 `json:"trackingUrl,omitempty"`
}

type SubCategory struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	Category        ProductCategory        `json:"category"`
	DefaultFlag     bool                   `json:"defaultFlag,omitempty"`
	ID              int                    `json:"id,omitempty"`
	InactiveFlag    bool                   `json:"inactiveFlag,omitempty"`
	IntegrationXref string                 `json:"integrationXref,omitempty"`
	Name            string                 `json:"name"`
}

type SubCategoryInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	Category     ProductCategory        `json:"category,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
}

type UnitOfMeasure struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag     bool                   `json:"defaultFlag,omitempty"`
	ID              int                    `json:"id,omitempty"`
	InactiveFlag    bool                   `json:"inactiveFlag,omitempty"`
	Name            string                 `json:"name"`
	UomScheduleXref string                 `json:"uomScheduleXref,omitempty"`
}

type Warehouse struct {
	Info                map[string]interface{} `json:"_info,omitempty"`
	Company             Company                `json:"company,omitempty"`
	Contact             Contact                `json:"contact,omitempty"`
	Currency            Currency               `json:"currency,omitempty"`
	Department          SystemDepartment       `json:"department"`
	ID                  int                    `json:"id,omitempty"`
	InactiveFlag        bool                   `json:"inactiveFlag,omitempty"`
	Location            SystemLocation         `json:"location"`
	LocationDefaultFlag bool                   `json:"locationDefaultFlag,omitempty"`
	LocationXref        string                 `json:"locationXref,omitempty"`
	LockedFlag          bool                   `json:"lockedFlag,omitempty"`
	Manager             Member                 `json:"manager,omitempty"`
	Name                string                 `json:"name"`
	OverallDefaultFlag  bool                   `json:"overallDefaultFlag,omitempty"`
	Site                Site                   `json:"site,omitempty"`
}

type WarehouseBin struct {
	Info           map[string]interface{} `json:"_info,omitempty"`
	Company        Company                `json:"company,omitempty"`
	DefaultFlag    bool                   `json:"defaultFlag,omitempty"`
	Department     SystemDepartment       `json:"department,omitempty"`
	Height         float64                `json:"height,omitempty"`
	ID             int                    `json:"id,omitempty"`
	InactiveFlag   bool                   `json:"inactiveFlag,omitempty"`
	Length         float64                `json:"length,omitempty"`
	Location       SystemLocation         `json:"location,omitempty"`
	Manager        Member                 `json:"manager,omitempty"`
	MaxQuantity    float64                `json:"maxQuantity,omitempty"`
	MinQuantity    float64                `json:"minQuantity,omitempty"`
	Name           string                 `json:"name"`
	OverflowBin    *WarehouseBin          `json:"overflowBin,omitempty"`
	QuantityOnHand int                    `json:"quantityOnHand,omitempty"`
	TransferBin    *WarehouseBin          `json:"transferBin,omitempty"`
	Warehouse      Warehouse              `json:"warehouse"`
	Weight         float64                `json:"weight,omitempty"`
	Width          float64                `json:"width,omitempty"`
}

type WarehouseBinInfo struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Warehouse    Warehouse              `json:"warehouse,omitempty"`
}

type WarehouseInfo struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	Company            Company                `json:"company,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	InactiveFlag       bool                   `json:"inactiveFlag,omitempty"`
	Name               string                 `json:"name,omitempty"`
	OverallDefaultFlag bool                   `json:"overallDefaultFlag,omitempty"`
}
