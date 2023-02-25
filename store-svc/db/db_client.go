package db

import (
	"log"
	"os"
	"sync"

	"github.com/Roy19/distributed-transaction-2pc/store-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	dbonce sync.Once
)

func InitDB() {
	dbonce.Do(func() {
		dsn := os.Getenv("DB_DSN")
		dbClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error connecting to database: %v\n", err)
		}
		DB = dbClient
		DB.AutoMigrate(&models.StoreItem{})
		DB.AutoMigrate(&models.StoreItemReservation{})
	})
}

func PutDummyData() {
	if DB != nil {
		storeItem := models.StoreItem{
			Name: "iPhone 12",
		}
		DB.Create(&storeItem)
		storeItemReservations := []models.StoreItemReservation{
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
			{
				StoreItem:  storeItem,
				IsReserved: false,
			},
		}
		DB.Create(&storeItemReservations)
	}
}
