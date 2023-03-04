package db

import (
	"log"
	"sync"

	"github.com/Roy19/distributed-transaction-2pc/store-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbClients map[string]*gorm.DB
	dbonce    sync.Once
)

func InitDB(dsn string, svcName string) {
	dbonce.Do(func() {
		dbClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("Error connecting to database: %v\n", err)
		}
		dbClients[svcName] = dbClient
		dbClients[svcName].AutoMigrate(&models.StoreItem{})
		dbClients[svcName].AutoMigrate(&models.StoreItemReservation{})
	})
}

func GetDBClient(svcName string) *gorm.DB {
	return dbClients[svcName]
}

func PutDummyData(svcName string) {
	if dbClients[svcName] != nil {
		storeItem := models.StoreItem{
			Name: "iPhone 12",
		}
		dbClients[svcName].Create(&storeItem)
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
		dbClients[svcName].Create(&storeItemReservations)
	}
}
