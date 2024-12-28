package internalsql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		fmt.Printf("dsn dari postgresql = %s", dataSourceName)
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	return db, nil
}
