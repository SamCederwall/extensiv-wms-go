package extensivWms

import (
	"context"
	"time"
)

type OrderService interface {
	Get(context.Context, uint64)
}

type CreateOrder struct {
	CustomerIdentifier             *CustomerIdentifier             `json:"customerIdentifier"`
	FacilityIdentifier             *FacilityIdentifier             `json:"facilityIdentifier"`
	WarehouseTransactionSourceEnum *WarehouseTransactionSourceEnum `json:"warehouseTransactionSourceEnum"`
	TransactionEntryTypeEnum       *TransactionEntryTypeEnum       `json:"transactionEntryTypeEnum"`
	DeferNotification              *bool                           `json:"deferNotification"`
	OrderItems                     []OrderItem                     `json:"orderItems"`
	ReferenceNum                   *string                         `json:"referenceNum"`
	PoNum                          *string                         `json:"poNum"`
	ExternalId                     *string                         `json:"externalId"`
	EarliestShipDate               *time.Time                      `json:"earliestShipDate"`
	ShipCancelDate                 *time.Time                      `json:"shipCancelDate"`
	Notes                          *string                         `json:"notes"`
	NumUnits1                      *float64                        `json:"numUnits1"`
	Unit1Identifier                *UnitIdentifier                 `json:"unit1Identifier"`
	NumUnits2                      *float64                        `json:"numUnits2"`
	Unit2Identifier                *UnitIdentifier                 `json:"unit2Identifier"`
	TotalWeight                    *float64                        `json:"totalWeight"`
	TotalVolumet                   *float64                        `json:"totalVolumet"`
	BillingCode                    *string                         `json:"billingCode"`
	AsnNumber                      *string                         `json:"asnNumber"`
	ExportChannelIdentifier        *ExportChannelIdentifier        `json:"exportChannelIdentifier"`
	RoutePickUpDate                *time.Time                      `json:"routePickupDate"`
	ShippingNotes                  *string                         `json:"shippingNotes"`
	MasterBillOfLadingId           *string                         `json:"masterBillOfLadingId"`
	InvoiceNumber                  *string                         `json:"invoiceNumber"`
	FulfillInvInfo                 *FulfillInvInfo                 `json:"fulfillInvInfo"`
	RoutingInfo                    *RoutingInfo                    `json:"routingInfo"`
	LoadNumber                     *string                         `json:"loadNumber"`
	ParcelOption                   *ParcelOption                   `json:"parcelOption"`
	BillOfLading                   *string                         `json:"billOfLading"`
	TrackingNumber                 *string                         `json:"trackingNumber"`
	TrailerNumber                  *string                         `json:"trailerNumber"`
	SealNumber                     *string                         `json:"sealNumber"`
	DoorNumber                     *string                         `json:"doorNumber"`
	PickupDate                     *time.Time                      `json:"pickupDate"`
	ShipTo                         *ShipSoldBillTo                 `json:"shipTo"`
	SoldTo                         *ShipSoldBillTo                 `json:"soldTo"`
	BillTo                         *ShipSoldBillTo                 `json:"billTo"`
	SavedElements                  []SavedElement                  `json:"savedElements"`
}

type OrderItem struct {
	ItemIdentifier        *ItemIdentifier      `json:"itemIdentifier"`
	Qualifier             *string              `json:"qualifier"`
	ExternalId            *string              `json:"externalId"`
	Qty                   *float64             `json:"qty"`
	SecondaryQty          *float64             `json:"secondaryQty"`
	LotNumber             *string              `json:"lotNumber"`
	SerialNumber          *string              `json:"serialNumber"`
	ExpirationDate        *time.Time           `json:"expirationDate"`
	WeightImperial        *float64             `json:"weightImperial"`
	WeightMetric          *float64             `json:"weightMetric"`
	Notes                 *string              `json:"notes"`
	FulfillInvSalePrice   *float64             `json:"fulfillInvSalePrice"`
	FulfillInvDiscountPct *float64             `json:"fulfillInvDiscountPct"`
	FulfillInvDiscountAmt *float64             `json:"fulfillInvDiscountAmt"`
	FulfillInvNote        *float64             `json:"fulfillInvNote"`
	SavedElements         []SavedElement       `json:"savedElements"`
	ProposedAloocations   []ProposedAllocation `json:"proposedAllocations"`
	OrderItemInPackages   []OrderItemInPackage `json:"orderItemInPackages"`
}

type CustomerIdentifier struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type FacilityIdentifier struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
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
	Name *string `json:"name"`
}

type ExportChannelIdentifier struct {
	Name *string `json:"name"`
}

type FulfillInvInfo struct {
	FulfillInvShippingAndHandling *float64 `json:"fulfillInvShippingAndHandling"`
	FulfillInvTax                 *float64 `json:"fulfillInvTax"`
	FulfillInvDiscountCode        *string  `json:"fulfillInvDiscountCode"`
	FulfillInvDiscountAmount      *float64 `json:"fulfillInvDiscountAmount"`
	FulfillInvGiftMessage         *string  `json:"fulfillInvGiftMessage"`
}

type RoutingInfo struct {
	IsCOD                  *bool                   `json:"isCOD"`
	IsInsurance            *bool                   `json:"isInsurance"`
	RequiresDeliveryConf   *bool                   `json:"requiresDeliveryConf"`
	RequiresReturnReceipt  *bool                   `json:"requiresReturnReceipt"`
	ScacCode               *string                 `json:"scacCode"`
	Carrier                *string                 `json:"carrier"`
	Mode                   *string                 `json:"mode"`
	Account                *string                 `json:"account"`
	ShipPointZip           *string                 `json:"shipPointZip"`
	CapacityTypeIdentifier *CapacityTypeIdentifier `json:"capacityTypeIdentifier"`
	LoadNumber             *string                 `json:"loadNumber"`
	ParcelOption           *ParcelOption           `json:"parcelOption"`
}

type CapacityTypeIdentifier struct {
	Name *string `json:"name"`
}

type ParcelOption struct {
	DeliveryConfirmationType  *string                    `json:"deliveryConfirmationType"`
	DryIceWeight              *float64                   `json:"dryIceWeight"`
	InsuranceAmount           *float64                   `json:"insuranceAmount"`
	InsuranceType             *InsuranceType             `json:"insuranceType"`
	InternationalContentsType *InternationalContentsType `json:"internationalContentsType"`
	InternationalNonDelivery  *InternationalNonDelivery  `json:"internationalNonDelivery"`
	ResidentialFlag           bool                       `json:"residentialFlag"`
	SaturdayDeliveryFlag      bool                       `json:"saturdayDeliveryFlag"`
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
	RetailerInfo  *RetailerInfo `json:"retailerInfo"`
	SameAs        *SameAs       `json:"sameAs"`
	IsQuickLookup *bool         `json:"isQuickLookup"`
	ContactId     *string       `json:"contactId"`
	CompanyName   *string       `json:"companyName"`
	Name          *string       `json:"name"`
	Address1      *string       `json:"address1"`
	Address2      *string       `json:"address2"`
	City          *string       `json:"city"`
	State         *string       `json:"state"`
	Zip           *string       `json:"zip"`
	Country       *string       `json:"country"`
	PhoneNumber   *string       `json:"phoneNumber"`
	Fax           *string       `json:"fax"`
	EmailAddress  *string       `json:"emailAddress"`
	Dept          *string       `json:"dept"`
	Code          *string       `json:"code"`
}

type SameAs int

const (
	SameAsShipTo SameAs = iota
	SameAsSoldTo
	SameAsBillTo
)

type RetailerInfo struct {
	NameKey *NameKey `json:"nameKey"`
}

type NameKey struct {
	CustomerIdentifier *NameKeyCustomerIdentifier `json:"customerIdentifier"`
	Name               *string                    `json:"name"`
	Number             *string                    `json:"number"`
}

type NameKeyCustomerIdentifier struct {
	ExternalId *string `json:"externalId"`
	Name       *string `json:"name"`
	Id         *int    `json:"id"`
}

type ItemIdentifier struct {
	Sku *string `json:"sku"`
	Id  *int    `json:"id"`
}

type SavedElement struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProposedAllocation struct {
	ReceivedItemId *int     `json:"receivedItemId"`
	Qty            *float64 `json:"qty"`
}

type OrderItemInPackage struct {
	PackageNumber     *int               `json:"packageNumber"`
	Quantity          *float64           `json:"quantity"`
	PackageIdentifier *PackageIdentifier `json:"packageIdentifier"`
}

type PackageIdentifier struct {
	Name *string `json:"name"`
	Id   *int    `json:"id"`
}

type GetOrder struct {
	ReadOnly *ReadOnly `json:"readOnly"`
}

type ReadOnly struct {
	OrderId                        *int             `json:"orderId"`
	AsnCandidate                   *int             `json:"asnCandidate"`
	RouteCandidate                 *int             `json:"routeCandidate"`
	FullyAllocated                 *bool            `json:"fullyAllocated"`
	ImportModuleId                 *int             `json:"importModuleId"`
	ExportModuleId                 *string          `json:"exportModuleId"`
	DeferNotification              *bool            `json:"deferNotification"`
	IsClosed                       *bool            `json:"isClosed"`
	ProcessDate                    *time.Time       `json:"processDate"`
	PickDoneDate                   *time.Time       `json:"pickDoneDate"`
	PickTicketPrintDate            *time.Time       `json:"pickTicketPrintDate"`
	PackDoneDate                   *time.Time       `json:"packDoneDate"`
	LabelIsExported                *bool            `json:"labelIsExported"`
	InvoiceExportedDate            *time.Time       `json:"invoiceExportedDate"`
	InvoiceDeliveredDate           *time.Time       `json:"invoiceDeliveredDate"`
	LoadedState                    *int             `json:"loadedState"`
	LoadOutDoneDate                *time.Time       `json:"loadOutDoneDate"`
	ReallocatedAfterPickTicketDate *time.Time       `json:"reallocatedAfterPickTicketDate"`
	RouteSent                      *bool            `json:"routeSent"`
	AsnSentDate                    *time.Time       `json:"asnSentDate"`
	PkgsExported                   *bool            `json:"pkgsExported"`
	BatchIdentifier                *BatchIdentifier `json:"batchIdentifier"`
}

type BatchIdentifier struct {
	NameKey *NameKey `json:"nameKey"`
	Id      *int     `json:"id"`
}
