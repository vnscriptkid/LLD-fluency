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

type DoorState struct {
	ID     uint `gorm:"primaryKey"`
	DoorID uint
	State  string
}

type Door struct {
	ID    uint
	State string
}

func (d *Door) Apply(states ...DoorState) {
	// Simplified version
	// In the real world, we would have different apply methods for each state
	for _, state := range states {
		d.State = state.State
		d.ID = state.DoorID
	}
}

const (
	Open   = "Open"
	Closed = "Closed"
	Locked = "Locked"
)

type CallbackFunc func(tx *gorm.DB, doorStates []DoorState) error

func processStateStates(db *gorm.DB, DoorID uint, cb CallbackFunc) error {
	// Get the door states using the DoorStateID with a pessimistic lock
	var doorStates []DoorState

	// Start a transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		resp := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&doorStates, "door_id = ?", DoorID)
		if resp.Error != nil {
			return resp.Error
		}

		return cb(tx, doorStates)
	})

	if err != nil {
		return fmt.Errorf("failed to get door states in txn: %w", err)
	}

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

	if err := db.AutoMigrate(&DoorState{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	doorState1 := DoorState{State: Open, DoorID: 1}
	doorState2 := DoorState{State: Closed, DoorID: 1}
	doorStates := []DoorState{doorState1, doorState2}

	err = db.Create(&doorStates).Error
	if err != nil {
		log.Fatalf("Failed to create DoorState: %v", err)
	}

	var wg sync.WaitGroup
	var errCount1 int32

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := processStateStates(db, 1, func(tx *gorm.DB, doorStates []DoorState) error {
				// Replay the door states one by one
				var door Door
				door.Apply(doorStates...)

				if door.State != Closed {
					return fmt.Errorf("door is not closed")
				}

				// insert a new door state
				newDoorState := DoorState{State: Locked, DoorID: 1}
				if err := tx.Create(&newDoorState).Error; err != nil {
					return fmt.Errorf("failed to create new door state: %w", err)
				}

				return nil
			})

			if err != nil {
				atomic.AddInt32(&errCount1, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Failed to get door states of id %d: %d times\n", 1, errCount1)
}
