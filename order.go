package extensivWms

import (
	"context"
	"time"
)

type OrderService interface {
	Get(context.Context, uint64, interface{}) (*GetOrder, error)
}

type OrderSerivceOp struct {
	client *Client
}

type CreateOrder struct {
	CustomerIdentifier             *CustomerIdentifier             `json:"customerIdentifier,omitempty"`
	FacilityIdentifier             *FacilityIdentifier             `json:"facilityIdentifier,omitempty"`
	WarehouseTransactionSourceEnum *WarehouseTransactionSourceEnum `json:"warehouseTransactionSourceEnum,omitempty"`
	TransactionEntryTypeEnum       *TransactionEntryTypeEnum       `json:"transactionEntryTypeEnum,omitempty"`
	DeferNotification              *bool                           `json:"deferNotification,omitempty"`
	OrderItems                     []OrderItem                     `json:"orderItems,omitempty"`
	ReferenceNum                   *string                         `json:"referenceNum,omitempty"`
	PoNum                          *string                         `json:"poNum,omitempty"`
	ExternalId                     *string                         `json:"externalId,omitempty"`
	EarliestShipDate               *time.Time                      `json:"earliestShipDate,omitempty"`
	ShipCancelDate                 *time.Time                      `json:"shipCancelDate,omitempty"`
	Notes                          *string                         `json:"notes,omitempty"`
	NumUnits1                      *float64                        `json:"numUnits1,omitempty"`
	Unit1Identifier                *UnitIdentifier                 `json:"unit1Identifier,omitempty"`
	NumUnits2                      *float64                        `json:"numUnits2,omitempty"`
	Unit2Identifier                *UnitIdentifier                 `json:"unit2Identifier,omitempty"`
	TotalWeight                    *float64                        `json:"totalWeight,omitempty"`
	TotalVolumet                   *float64                        `json:"totalVolumet,omitempty"`
	BillingCode                    *string                         `json:"billingCode,omitempty"`
	AsnNumber                      *string                         `json:"asnNumber,omitempty"`
	ExportChannelIdentifier        *ExportChannelIdentifier        `json:"exportChannelIdentifier,omitempty"`
	RoutePickUpDate                *time.Time                      `json:"routePickupDate,omitempty"`
	ShippingNotes                  *string                         `json:"shippingNotes,omitempty"`
	MasterBillOfLadingId           *string                         `json:"masterBillOfLadingId,omitempty"`
	InvoiceNumber                  *string                         `json:"invoiceNumber,omitempty"`
	FulfillInvInfo                 *FulfillInvInfo                 `json:"fulfillInvInfo,omitempty"`
	RoutingInfo                    *RoutingInfo                    `json:"routingInfo,omitempty"`
	LoadNumber                     *string                         `json:"loadNumber,omitempty"`
	ParcelOption                   *ParcelOption                   `json:"parcelOption,omitempty"`
	BillOfLading                   *string                         `json:"billOfLading,omitempty"`
	TrackingNumber                 *string                         `json:"trackingNumber,omitempty"`
	TrailerNumber                  *string                         `json:"trailerNumber,omitempty"`
	SealNumber                     *string                         `json:"sealNumber,omitempty"`
	DoorNumber                     *string                         `json:"doorNumber,omitempty"`
	PickupDate                     *time.Time                      `json:"pickupDate,omitempty"`
	ShipTo                         *ShipSoldBillTo                 `json:"shipTo,omitempty"`
	SoldTo                         *ShipSoldBillTo                 `json:"soldTo,omitempty"`
	BillTo                         *ShipSoldBillTo                 `json:"billTo,omitempty"`
	SavedElements                  []SavedElement                  `json:"savedElements,omitempty"`
}

type OrderItem struct {
	ItemIdentifier        *ItemIdentifier      `json:"itemIdentifier,omitempty"`
	Qualifier             *string              `json:"qualifier,omitempty"`
	ExternalId            *string              `json:"externalId,omitempty"`
	Qty                   *float64             `json:"qty,omitempty"`
	SecondaryQty          *float64             `json:"secondaryQty,omitempty"`
	LotNumber             *string              `json:"lotNumber,omitempty"`
	SerialNumber          *string              `json:"serialNumber,omitempty"`
	ExpirationDate        *time.Time           `json:"expirationDate,omitempty"`
	WeightImperial        *float64             `json:"weightImperial,omitempty"`
	WeightMetric          *float64             `json:"weightMetric,omitempty"`
	Notes                 *string              `json:"notes,omitempty"`
	FulfillInvSalePrice   *float64             `json:"fulfillInvSalePrice,omitempty"`
	FulfillInvDiscountPct *float64             `json:"fulfillInvDiscountPct,omitempty"`
	FulfillInvDiscountAmt *float64             `json:"fulfillInvDiscountAmt,omitempty"`
	FulfillInvNote        *float64             `json:"fulfillInvNote,omitempty"`
	SavedElements         []SavedElement       `json:"savedElements,omitempty"`
	ProposedAloocations   []ProposedAllocation `json:"proposedAllocations,omitempty"`
	OrderItemInPackages   []OrderItemInPackage `json:"orderItemInPackages,omitempty"`
}

type CustomerIdentifier struct {
	Name       *string `json:"name,omitempty"`
	Id         *int    `json:"id,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
}

type FacilityIdentifier struct {
	Name *string `json:"name,omitempty"`
	Id   *int    `json:"id,omitempty"`
}

type WarehouseTransactionSourceEnum int

const (
	UiManual WarehouseTransactionSourceEnum = iota
	UiImport
	AutomatedImport
	ExternalSoapApi
	QuickBooksSoapApi
	AutomatedSystemCharge
	RestApi
)

type TransactionEntryTypeEnum int

const (
	ManualWarehouseUser TransactionEntryTypeEnum = iota
	ManualCustomerUser
	Ftp
	Api
)

type UnitIdentifier struct {
	Name *string `json:"name,omitempty"`
}

type ExportChannelIdentifier struct {
	Name *string `json:"name,omitempty"`
}

type FulfillInvInfo struct {
	FulfillInvShippingAndHandling *float64 `json:"fulfillInvShippingAndHandling,omitempty"`
	FulfillInvTax                 *float64 `json:"fulfillInvTax,omitempty"`
	FulfillInvDiscountCode        *string  `json:"fulfillInvDiscountCode,omitempty"`
	FulfillInvDiscountAmount      *float64 `json:"fulfillInvDiscountAmount,omitempty"`
	FulfillInvGiftMessage         *string  `json:"fulfillInvGiftMessage,omitempty"`
}

type RoutingInfo struct {
	IsCOD                  *bool                   `json:"isCOD,omitempty"`
	IsInsurance            *bool                   `json:"isInsurance,omitempty"`
	RequiresDeliveryConf   *bool                   `json:"requiresDeliveryConf,omitempty"`
	RequiresReturnReceipt  *bool                   `json:"requiresReturnReceipt,omitempty"`
	ScacCode               *string                 `json:"scacCode,omitempty"`
	Carrier                *string                 `json:"carrier,omitempty"`
	Mode                   *string                 `json:"mode,omitempty"`
	Account                *string                 `json:"account,omitempty"`
	ShipPointZip           *string                 `json:"shipPointZip,omitempty"`
	CapacityTypeIdentifier *CapacityTypeIdentifier `json:"capacityTypeIdentifier,omitempty"`
	LoadNumber             *string                 `json:"loadNumber,omitempty"`
	ParcelOption           *ParcelOption           `json:"parcelOption,omitempty"`
	BillOfLading           *string                 `json:"billOfLading,omitempty"`
	TrackingNumber         *string                 `json:"trackingNumber,omitempty"`
	TrailerNumber          *string                 `json:"trailerNumber,omitempty"`
	SealNumber             *string                 `json:"sealNumber,omitempty"`
	DoorNumber             *string                 `json:"doorNumber,omitempty"`
	PickupDate             *time.Time              `json:"pickupDate,omitempty"`
}

type CapacityTypeIdentifier struct {
	Name *string `json:"name,omitempty"`
}

type ParcelOption struct {
	OrderId                   *int                       `json:"orderId,omitempty"`
	DeliveryConfirmationType  *string                    `json:"deliveryConfirmationType,omitempty"`
	DeliveryDutyPaid          *int                       `json:"deliveryDutyPaid,omitempty"`
	DryIceWeight              *float64                   `json:"dryIceWeight,omitempty"`
	InsuranceAmount           *float64                   `json:"insuranceAmount,omitempty"`
	InsuranceType             *InsuranceType             `json:"insuranceType,omitempty"`
	InternationalContentsType *InternationalContentsType `json:"internationalContentsType,omitempty"`
	InternationalNonDelivery  *InternationalNonDelivery  `json:"internationalNonDelivery,omitempty"`
	ResidentialFlag           bool                       `json:"residentialFlag,omitempty"`
	SaturdayDeliveryFlag      bool                       `json:"saturdayDeliveryFlag,omitempty"`
}

type InsuranceType int

const (
	InsuranceTypeNone       InsuranceType = 0
	InsuranceTypeCarrier    InsuranceType = 106
	InsuranceTypeThirdParty InsuranceType = 107
)

type InternationalContentsType string

const (
	InternationalContentsTypeNone          InternationalContentsType = "none"
	InternationalContentsTypeDocuments     InternationalContentsType = "documents"
	InternationalContentsTypeGift          InternationalContentsType = "gift"
	InternationalContentsTypeMerchandise   InternationalContentsType = "merchandise"
	InternationalContentsTypeReturnedGoods InternationalContentsType = "returned_goods"
	InternationalContentsTypeSample        InternationalContentsType = "sample"
)

type InternationalNonDelivery string

const (
	InternationalNonDeliveryNone             InternationalNonDelivery = "none"
	InternationalNonDeliveryTreatAsAbandoned InternationalNonDelivery = "treat_as_abandoned"
	InternationalNonDeliveryReturnToSender   InternationalNonDelivery = "return_to_sender"
)

type ShipSoldBillTo struct {
	RetailerInfo  *RetailerInfo `json:"retailerInfo,omitempty"`
	SameAs        *SameAs       `json:"sameAs,omitempty"`
	RetailerId    *int          `json:"retailerId,omitempty"`
	IsQuickLookup *bool         `json:"isQuickLookup,omitempty"`
	ContactId     *string       `json:"contactId,omitempty"`
	CompanyName   *string       `json:"companyName,omitempty"`
	Name          *string       `json:"name,omitempty"`
	Address1      *string       `json:"address1,omitempty"`
	Address2      *string       `json:"address2,omitempty"`
	City          *string       `json:"city,omitempty"`
	State         *string       `json:"state,omitempty"`
	Zip           *string       `json:"zip,omitempty"`
	Country       *string       `json:"country,omitempty"`
	PhoneNumber   *string       `json:"phoneNumber,omitempty"`
	Fax           *string       `json:"fax,omitempty"`
	EmailAddress  *string       `json:"emailAddress,omitempty"`
	Dept          *string       `json:"dept,omitempty"`
	Code          *string       `json:"code,omitempty"`
	AddressStatus *int          `json:"addressStatus,omitempty"`
}

type SameAs int

const (
	SameAsShipTo SameAs = iota
	SameAsSoldTo
	SameAsBillTo
)

type RetailerInfo struct {
	NameKey *NameKey `json:"nameKey,omitempty"`
	Id      *int     `json:"id,omitempty"`
}

type NameKey struct {
	CustomerIdentifier *NameKeyCustomerIdentifier `json:"customerIdentifier,omitempty"`
	Name               *string                    `json:"name,omitempty"`
	Number             *string                    `json:"number,omitempty"`
}

type NameKeyCustomerIdentifier struct {
	ExternalId *string `json:"externalId,omitempty"`
	Name       *string `json:"name,omitempty"`
	Id         *int    `json:"id,omitempty"`
}

type ItemIdentifier struct {
	Sku *string `json:"sku,omitempty"`
	Id  *int    `json:"id,omitempty"`
}

type SavedElement struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ProposedAllocation struct {
	ReceivedItemId *int     `json:"receivedItemId,omitempty"`
	Qty            *float64 `json:"qty,omitempty"`
}

type OrderItemInPackage struct {
	PackageNumber     *int               `json:"packageNumber,omitempty"`
	Quantity          *float64           `json:"quantity,omitempty"`
	PackageIdentifier *PackageIdentifier `json:"packageIdentifier,omitempty"`
}

type PackageIdentifier struct {
	Name *string `json:"name,omitempty"`
	Id   *int    `json:"id,omitempty"`
}

type GetOrder struct {
	ReadOnly                *ReadOnly                `json:"readOnly,omitempty"`
	ReferenceNum            *string                  `json:"referenceNum,omitempty"`
	Description             *string                  `json:"description,omitempty"`
	PoNum                   *string                  `json:"poNum,omitempty"`
	ExternalId              *string                  `json:"externalId,omitempty"`
	EarliestShipDate        *time.Time               `json:"earliestShipDate,omitempty"`
	ShipCancelDate          *time.Time               `json:"shipCancelDate,omitempty"`
	Notes                   *string                  `json:"notes,omitempty"`
	NumUnits1               *float64                 `json:"numUnits1,omitempty"`
	Unit1Identifier         *UnitIdentifier          `json:"unit1Identifier,omitempty"`
	NumUnits2               *float64                 `json:"numUnits2,omitempty"`
	Unit2Identifier         *UnitIdentifier          `json:"unit2Identifier,omitempty"`
	TotalWeight             *float64                 `json:"totalWeight,omitempty"`
	TotalVolume             *float64                 `json:"totalVolume,omitempty"`
	BillingCode             *string                  `json:"billingCode,omitempty"`
	AsnNumber               *string                  `json:"asnNumber,omitempty"`
	UpsServiceOptionCharge  *float64                 `json:"upsServiceOptionCharge,omitempty"`
	UpsTransportationCharge *float64                 `json:"upsTransportationCharge,omitempty"`
	UpsIsResidential        *bool                    `json:"upsIsResidential,omitempty"`
	ExportChannelIdentifier *ExportChannelIdentifier `json:"exportChannelIdentifier,omitempty"`
	RoutePickupDate         *time.Time               `json:"routePickupDate,omitempty"`
	ShippingNotes           *string                  `json:"shippingNotes,omitempty"`
	MasterBillOfLadingId    *string                  `json:"masterBillOfLadingId,omitempty"`
	InvoiceNumber           *string                  `json:"invoiceNumber,omitempty"`
	FulfillInvInfo          *FulfillInvInfo          `json:"fulfillInvInfo,omitempty"`
	RoutingInfo             *RoutingInfo             `json:"routingInfo,omitempty"`
	Billing                 *Billing                 `json:"billing,omitempty"`
	ShipTo                  *ShipSoldBillTo          `json:"shipTo,omitempty"`
	SoldTo                  *ShipSoldBillTo          `json:"soldTo,omitempty"`
	BillTo                  *ShipSoldBillTo          `json:"billTo,omitempty"`
	SavedElements           []SavedElement           `json:"savedElements,omitempty"`
	ParcelOption            *ParcelOption            `json:"parcelOption,omitempty"`
}

type ReadOnly struct {
	OrderId                        *int                           `json:"orderId,omitempty"`
	AsnCandidate                   *int                           `json:"asnCandidate,omitempty"`
	RouteCandidate                 *int                           `json:"routeCandidate,omitempty"`
	FullyAllocated                 *bool                          `json:"fullyAllocated,omitempty"`
	ImportModuleId                 *int                           `json:"importModuleId,omitempty"`
	ExportModuleId                 *string                        `json:"exportModuleId,omitempty"`
	DeferNotification              *bool                          `json:"deferNotification,omitempty"`
	IsClosed                       *bool                          `json:"isClosed,omitempty"`
	ProcessDate                    *time.Time                     `json:"processDate,omitempty"`
	PickDoneDate                   *time.Time                     `json:"pickDoneDate,omitempty"`
	PickTicketPrintDate            *time.Time                     `json:"pickTicketPrintDate,omitempty"`
	PackDoneDate                   *time.Time                     `json:"packDoneDate,omitempty"`
	LabelIsExported                *bool                          `json:"labelIsExported,omitempty"`
	InvoiceExportedDate            *time.Time                     `json:"invoiceExportedDate,omitempty"`
	InvoiceDeliveredDate           *time.Time                     `json:"invoiceDeliveredDate,omitempty"`
	LoadedState                    *int                           `json:"loadedState,omitempty"`
	LoadOutDoneDate                *time.Time                     `json:"loadOutDoneDate,omitempty"`
	ReallocatedAfterPickTicketDate *time.Time                     `json:"reallocatedAfterPickTicketDate,omitempty"`
	RouteSent                      *bool                          `json:"routeSent,omitempty"`
	AsnSentDate                    *time.Time                     `json:"asnSentDate,omitempty"`
	PkgsExported                   *bool                          `json:"pkgsExported,omitempty"`
	BatchIdentifier                *BatchIdentifier               `json:"batchIdentifier,omitempty"`
	Packages                       []Package                      `json:"packages,omitempty"`
	OutboundSerialNumbers          []OutboundSerialNumber         `json:"outboundSerialNumbers,omitempty"`
	ParcelLabelType                *int                           `json:"parcelLabelType,omitempty"`
	CustomerIdentifier             *CustomerIdentifier            `json:"customerIdentifier,omitempty"`
	FacilityIdentifier             *FacilityIdentifier            `json:"facilityIdentifier,omitempty"`
	WarehouseTransactionSourceType *int                           `json:"warehouseTransactionSourceType,omitempty"`
	TransactionEntryType           *int                           `json:"transactionEntryType,omitempty"`
	ImportChannelIdentifier        *ImportChannelIdentifier       `json:"importChannelIdentifier,omitempty"`
	CreationDate                   *time.Time                     `json:"creationDate,omitempty"`
	CreatedByIdentifier            *CreatedOrModifiedByIdentifier `json:"createdByIdentifier,omitempty"`
	LastModifiedDate               *time.Time                     `json:"lastModifiedDate,omitempty"`
	LastModifiedByIdentifier       *CreatedOrModifiedByIdentifier `json:"lastModifiedByIdentifier,omitempty"`
	Status                         *int                           `json:"status,omitempty"`
}

type BatchIdentifier struct {
	NameKey *NameKey `json:"nameKey,omitempty"`
	Id      *int     `json:"id,omitempty"`
}

type Package struct {
	PackageId            *int                  `json:"packageId,omitempty"`
	PackageTypeId        *int                  `json:"packageTypeId,omitempty"`
	PackageDefIdentifier *PackageDefIdentifier `json:"packageDefIdentifier,omitempty"`
	Length               *float64              `json:"length,omitempty"`
	Width                *float64              `json:"width,omitempty"`
	Height               *float64              `json:"height,omitempty"`
	CodAmount            *float64              `json:"codAmount,omitempty"`
	InsuredAmount        *float64              `json:"insuredAmount,omitempty"`
	TrackingNumber       *string               `json:"trackingNumber,omitempty"`
	Description          *string               `json:"description,omitempty"`
	CreateDate           *time.Time            `json:"createDate,omitempty"`
	Oversize             *bool                 `json:"oversize,omitempty"`
	Cod                  *bool                 `json:"cod,omitempty"`
	Ucc128               *int                  `json:"ucc128,omitempty"`
	CartonId             *string               `json:"cartonId,omitempty"`
	Label                *string               `json:"label,omitempty"`
	PackageContents      []PackageContent      `json:"packageContents,omitempty"`
}

type PackageContent struct {
	PackageContentId         *int       `json:"packageContentId,omitempty"`
	PackageId                *int       `json:"packageId,omitempty"`
	OrderItemId              *int       `json:"orderItemId,omitempty"`
	ReceiveItemId            *int       `json:"receiveItemId,omitempty"`
	OrderItemPickExceptionId *int       `json:"orderItemPickExceptionId,omitempty"`
	Qty                      *float64   `json:"qty,omitempty"`
	LotNumber                *string    `json:"lotNumber,omitempty"`
	SerialNumber             *string    `json:"serialNumber,omitempty"`
	ExpirationDate           *time.Time `json:"expirationDate,omitempty"`
	CreateDate               *time.Time `json:"createDate,omitempty"`
	SerialNumbers            []string   `json:"serialNumbers,omitempty"`
}

type PackageDefIdentifier struct {
	Name *string `json:"name,omitempty"`
	Id   *int    `json:"id,omitempty"`
}

type OutboundSerialNumber struct {
	ItemIdentifier *ItemIdentifier `json:"itemIdentifier,omitempty"`
	Qualifier      *string         `json:"qualifier,omitempty"`
	SerialNumbers  []string        `json:"serialNumbers,omitempty"`
}

type ImportChannelIdentifier struct {
	Name *string `json:"name,omitempty"`
	Id   *int    `json:"id,omitempty"`
}

type CreatedOrModifiedByIdentifier struct {
	Name *string `json:"name,omitempty"`
	Id   *int    `json:"int,omitempty"`
}

type Billing struct {
	BillingCharges []BillingCharge `json:"billingCharges,omitempty"`
}

type BillingCharge struct {
	ChargeType *int     `json:"chargeType,omitempty"`
	Subtotal   *float64 `json:"subtotal,omitempty"`
	Details    []Detail `json:"details,omitempty"`
}

type Detail struct {
	WarehouseTransactionPriceCalcId *int     `json:"warehouseTransactionPriceCalcId,omitempty"`
	NumUnits                        *float64 `json:"numUnits,omitempty"`
	ChargeLabel                     *string  `json:"chargeLabel,omitempty"`
	UnitDescription                 *string  `json:"unitDescription,omitempty"`
	ChargePerUnit                   *float64 `json:"chargePerUnit,omitempty"`
	GlAcctNum                       *string  `json:"glAcctNum,omitempty"`
	Sku                             *string  `json:"sku,omitempty"`
	PtItem                          *string  `json:"ptItem,omitempty"`
	PtItemDescription               *string  `json:"ptItemDescription,omitempty"`
	PtArAcct                        *string  `json:"ptArAcct,omitempty"`
	SystemGenerated                 *bool    `json:"systemGenerated,omitempty"`
	TaxCode                         *string  `json:"taxCode,omitempty"`
}
