package models

import "gorm.io/gorm"

type StoreItemReservation struct {
	gorm.Model
	StoreItemID    int
	StoreItem      StoreItem
	IsReserved     bool
	CurrentOrderId string
}
