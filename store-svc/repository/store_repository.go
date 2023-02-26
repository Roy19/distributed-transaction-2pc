package repository

import (
	"fmt"

	"github.com/Roy19/distributed-transaction-2pc/store-svc/db"
	"github.com/Roy19/distributed-transaction-2pc/store-svc/models"
	"gorm.io/gorm"
)

type StoreRepository struct {
}

func (s *StoreRepository) GetItem(itemID int64) (int64, error) {
	var item models.StoreItem
	txOut := db.DB.First(&item, itemID)
	if txOut.Error == gorm.ErrRecordNotFound {
		return 0, fmt.Errorf("item not found")
	}
	return int64(item.ID), nil
}

func (s *StoreRepository) CreateReservation(itemID int64) error {
	txn := db.DB.Model(&models.StoreItemReservation{}).Begin()
	var storeReservation models.StoreItemReservation
	txn = txn.Raw(`select * from store_item_reservations 
		where is_reserved = false and store_item_id = ?
		for update`, int(itemID)).Scan(&storeReservation)
	if txn.Error != nil || txn.RowsAffected == 0 {
		txn.Rollback()
		return fmt.Errorf("no more reservations can be done on item")
	}
	txn = txn.Exec(`update store_item_reservations
			set is_reserved = true
			where id = ?`, storeReservation.ID)
	if txn.Error != nil {
		txn.Rollback()
		return fmt.Errorf("failed to set lock on store item")
	}
	txn.Commit()
	return nil
}
