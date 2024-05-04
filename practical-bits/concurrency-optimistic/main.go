package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Door struct {
	ID      uint `gorm:"primaryKey"`
	State   string
	Version uint
}

const (
	Open   = "Open"
	Closed = "Closed"
	Locked = "Locked"
)

func updateDoorState(db *gorm.DB, door *Door, state string) error {
	oldVersion := door.Version

	result := db.Model(door).
		Clauses(clause.Returning{}).
		Where("id = ? AND version = ?", door.ID, door.Version).
		Updates(map[string]interface{}{"state": state, "version": gorm.Expr("version + 1")})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("outdated version")
	}

	newVersion := door.Version

	fmt.Printf("Old version: %d, New version: %d\n", oldVersion, newVersion)

	return nil
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=test port=5433 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Door{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	door := Door{State: Closed}

	err = db.Create(&door).Error
	if err != nil {
		log.Fatalf("Failed to create door: %v", err)
	}

	var wg sync.WaitGroup
	var errCount int32

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := updateDoorState(db, &door, Open); err != nil {
				atomic.AddInt32(&errCount, 1)
				log.Printf("Failed to update door state: %v", err)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Failed to update door state %d times\n", errCount)
	// Failed to update door state 9 times
}
