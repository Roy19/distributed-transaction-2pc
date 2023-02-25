package repository

import (
	"fmt"

	"github.com/Roy19/distributed-transaction-2pc/store-svc/db"
	"github.com/Roy19/distributed-transaction-2pc/store-svc/models"
	"gorm.io/gorm"
)

type StoreRepository struct {
}

func (s *StoreRepository) GetItem(itemID int) (int64, error) {
	var item models.StoreItem
	txOut := db.DB.First(&item, itemID)
	if txOut.Error == gorm.ErrRecordNotFound {
		return 0, fmt.Errorf("item not found")
	}
	return int64(item.ID), nil
}
