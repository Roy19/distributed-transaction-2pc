package controllers

import "github.com/Roy19/distributed-transaction-2pc/store-svc/repository"

type StoreController struct {
	StoreRepository *repository.StoreRepository
}

func (c *StoreController) GetItem(itemID int64) {
	// check if item exists in db
}

func (c *StoreController) ReserveItem(itemID int64) {
	// check if item exists in db
	// check if item is already reserved
	// reserve item
}

func (c *StoreController) BookItem(itemID int64) {
	// check if item exists in db
	// check if item is already reserved
	// book item
}
