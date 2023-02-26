package controllers

import (
	"log"

	"github.com/Roy19/distributed-transaction-2pc/store-svc/repository"
)

type StoreController struct {
	StoreRepository *repository.StoreRepository
}

func (c *StoreController) GetItem(itemID int64) error {
	_, err := c.StoreRepository.GetItem(itemID)
	return err
}

func (c *StoreController) ReserveItem(itemID int64) error {
	err := c.StoreRepository.CreateReservation(itemID)
	if err != nil {
		log.Printf("[ERROR] Failed to create a reservation on that item")
	}
	return err
}

func (c *StoreController) BookItem(itemID int64) {
	// check if item exists in db
	// check if item is already reserved
	// book item
}
