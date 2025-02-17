package extensivWms

import "time"

type Order struct {
	CustomerIdentifier             *CustomerIdentifier             `json:"customerIdentifier"`
	FacilityIdentifier             *FacilityIdentifier             `json:"facilityIdentifier"`
	WarehouseTransactionSourceEnum *WarehouseTransactionSourceEnum `json:"warehouseTransactionSourceEnum"`
	TransactionEntryTypeEnum       *TransactionEntryTypeEnum       `json:"transactionEntryTypeEnum"`
	DeferNotification              *bool                           `json:"deferNotification"`
}

type OrderItem struct {
	ItemIdentifier        *ItemIdentifier `json:"itemIdentifier"`
	Qualifier             *string         `json:"qualifier"`
	ExternalId            *string         `json:"externalId"`
	Qty                   *float64        `json:"qty"`
	SecondaryQty          *float64        `json:"secondaryQty"`
	LotNumber             *string         `json:"lotNumber"`
	SerialNumber          *string         `json:"serialNumber"`
	ExpirationDate        *time.Time      `json:"expirationDate"`
	WeightImperial        *float64        `json:"weightImperial"`
	WeightMetric          *float64        `json:"weightMetric"`
	Notes                 *string         `json:"notes"`
	FulfillInvSalePrice   *float64        `json:"fulfillInvSalePrice"`
	FulfillInvDiscountPct *float64        `json:"fulfillInvDiscountPct"`
	FulfillInvDiscountAmt *float64        `json:"fulfillInvDiscountAmt"`
	FulfillInvNote        *float64        `json:"fulfillInvNote"`
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

type ItemIdentifier struct {
	Sku *string `json:"sku"`
	Id  *int    `json:"id"`
}
