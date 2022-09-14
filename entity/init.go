package entity

import (
	"Cocome/entityRepo"
	"log"
)

const (
	CardPaymentInheritType entityRepo.GoenInheritType = iota + 1
	CashPaymentInheritType
)

func init() {
	tmpPaymentRepo, err := entityRepo.NewRepo[PaymentEntity, Payment]("payment")
	if err != nil {
		log.Fatal(err)
	}
	paymentRepo = tmpPaymentRepo
	PaymentRepo = tmpPaymentRepo
	tmpCardPaymentRepo, err := entityRepo.NewInheritRepo[CardPaymentEntity, CardPayment]("card_payment", tmpPaymentRepo, CardPaymentInheritType)
	if err != nil {
		log.Fatal(err)
	}
	cardPaymentRepo = tmpCardPaymentRepo
	CardPaymentRepo = tmpCardPaymentRepo
	tmpCashDeskRepo, err := entityRepo.NewRepo[CashDeskEntity, CashDesk]("cash_desk")
	if err != nil {
		log.Fatal(err)
	}
	cashDeskRepo = tmpCashDeskRepo
	CashDeskRepo = tmpCashDeskRepo
	tmpCashPaymentRepo, err := entityRepo.NewInheritRepo[CashPaymentEntity, CashPayment]("cash_payment", tmpPaymentRepo, CashPaymentInheritType)
	if err != nil {
		log.Fatal(err)
	}
	cashPaymentRepo = tmpCashPaymentRepo
	CashPaymentRepo = tmpCashPaymentRepo
	tmpCashierRepo, err := entityRepo.NewRepo[CashierEntity, Cashier]("cashier")
	if err != nil {
		log.Fatal(err)
	}
	cashierRepo = tmpCashierRepo
	CashierRepo = tmpCashierRepo
	tmpItemRepo, err := entityRepo.NewRepo[ItemEntity, Item]("item")
	if err != nil {
		log.Fatal(err)
	}
	itemRepo = tmpItemRepo
	ItemRepo = tmpItemRepo
	tmpOrderEntryRepo, err := entityRepo.NewRepo[OrderEntryEntity, OrderEntry]("order_entry")
	if err != nil {
		log.Fatal(err)
	}
	orderEntryRepo = tmpOrderEntryRepo
	OrderEntryRepo = tmpOrderEntryRepo
	tmpOrderProductRepo, err := entityRepo.NewRepo[OrderProductEntity, OrderProduct]("order_product")
	if err != nil {
		log.Fatal(err)
	}
	orderProductRepo = tmpOrderProductRepo
	OrderProductRepo = tmpOrderProductRepo
	tmpProductCatalogRepo, err := entityRepo.NewRepo[ProductCatalogEntity, ProductCatalog]("product_catalog")
	if err != nil {
		log.Fatal(err)
	}
	productCatalogRepo = tmpProductCatalogRepo
	ProductCatalogRepo = tmpProductCatalogRepo
	tmpSaleRepo, err := entityRepo.NewRepo[SaleEntity, Sale]("sale")
	if err != nil {
		log.Fatal(err)
	}
	saleRepo = tmpSaleRepo
	SaleRepo = tmpSaleRepo
	tmpSalesLineItemRepo, err := entityRepo.NewRepo[SalesLineItemEntity, SalesLineItem]("sales_line_item")
	if err != nil {
		log.Fatal(err)
	}
	salesLineItemRepo = tmpSalesLineItemRepo
	SalesLineItemRepo = tmpSalesLineItemRepo
	tmpStoreRepo, err := entityRepo.NewRepo[StoreEntity, Store]("store")
	if err != nil {
		log.Fatal(err)
	}
	storeRepo = tmpStoreRepo
	StoreRepo = tmpStoreRepo
	tmpSupplierRepo, err := entityRepo.NewRepo[SupplierEntity, Supplier]("supplier")
	if err != nil {
		log.Fatal(err)
	}
	supplierRepo = tmpSupplierRepo
	SupplierRepo = tmpSupplierRepo

}
