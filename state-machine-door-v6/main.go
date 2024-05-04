package main

import (
	"context"
	"fmt"
	"log"

	"github.com/qmuntal/stateless"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/optimisticlock"
)

type Door struct {
	ID      uint `gorm:"primaryKey"`
	State   string
	KeyCode string
	Version optimisticlock.Version
}

type DoorController struct {
	door         Door
	stateMachine *stateless.StateMachine
	db           *gorm.DB
}

const (
	Open   = "Open"
	Closed = "Closed"
	Locked = "Locked"
)

const (
	OpenTrigger   = "Open"
	CloseTrigger  = "Close"
	LockTrigger   = "Lock"
	UnlockTrigger = "Unlock"
)

func NewDoorController(keyCode string, db *gorm.DB) (*DoorController, error) {
	door := Door{State: Closed, KeyCode: keyCode}
	if err := db.Create(&door).Error; err != nil {
		return nil, err
	}

	dc := &DoorController{door: door, db: db}
	dc.stateMachine = stateless.NewStateMachine(Closed)

	dc.stateMachine.Configure(Open).
		Permit(CloseTrigger, Closed).
		Ignore(OpenTrigger).
		PermitReentry(UnlockTrigger).
		OnEntry(dc.saveState)

	dc.stateMachine.Configure(Closed).
		Permit(OpenTrigger, Open).
		Permit(LockTrigger, Locked).
		Ignore(CloseTrigger).
		PermitReentry(UnlockTrigger).
		OnEntryFrom(UnlockTrigger, dc.saveState).
		OnEntryFrom(LockTrigger, dc.saveState)

	dc.stateMachine.Configure(Locked).
		PermitDynamic(UnlockTrigger, func(ctx context.Context, args ...any) (stateless.State, error) {
			if len(args) == 0 {
				return nil, fmt.Errorf("key code is required")
			}

			inputKeyCode, ok := args[0].(string)

			if !ok {
				return nil, fmt.Errorf("failed to convert key code to string")
			}

			if !dc.isKeyCodeValid(inputKeyCode) {
				return nil, fmt.Errorf("invalid key code")
			}
			return Closed, nil
		}).
		Ignore(LockTrigger).
		Ignore(CloseTrigger).
		Ignore(OpenTrigger).
		OnEntry(dc.saveState)

	return dc, nil
}

func (dc *DoorController) isKeyCodeValid(inputKeyCode string) bool {
	return inputKeyCode == dc.door.KeyCode
}

func (dc *DoorController) saveState(ctx context.Context, args ...any) error {
	state, ok := dc.stateMachine.MustState().(string)

	if !ok {
		return fmt.Errorf("failed to convert state to string")
	}

	resp := dc.db.Model(&dc.door).Select("*").Updates(map[string]interface{}{
		"state": state,
	})

	if resp.Error != nil {
		return fmt.Errorf("failed to update door state: %w", resp.Error)
	}

	if resp.RowsAffected == 0 {
		return fmt.Errorf("door state was updated by another transaction")
	}

	// reload door
	resp = dc.db.Model(&dc.door).Where("id = ?", dc.door.ID).Select("*").First(&dc.door)

	if resp.Error != nil {
		return fmt.Errorf("failed to reload door: %w", resp.Error)
	}

	fmt.Printf("Reloaded door: %#v\n", dc.door)

	return nil
}

func (dc *DoorController) Open() {
	err := dc.stateMachine.Fire(OpenTrigger)
	if err != nil {
		fmt.Println("Failed to open:", err)
	}
}

func (dc *DoorController) Close() {
	err := dc.stateMachine.Fire(CloseTrigger)
	if err != nil {
		fmt.Println("Failed to close:", err)
	}
}

func (dc *DoorController) Lock() {
	err := dc.stateMachine.Fire(LockTrigger)
	if err != nil {
		fmt.Println("Failed to lock:", err)
	}
}

func (dc *DoorController) Unlock(keyCodeInput string) {
	err := dc.stateMachine.FireCtx(context.Background(), UnlockTrigger, keyCodeInput)
	if err != nil {
		fmt.Println("Failed to unlock:", err)
	}
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=door port=5433 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Door{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	doorController, err := NewDoorController("1234", db)
	if err != nil {
		log.Fatalf("Failed to create door controller: %v", err)
	}

	fmt.Println("Current State:", doorController.stateMachine.MustState()) // Closed

	doorController.Lock()
	fmt.Println("Current State:", doorController.stateMachine.MustState()) // Locked

	doorController.Unlock("1")
	fmt.Println("Current State:", doorController.stateMachine.MustState()) // Locked

	doorController.Unlock("1234")
	fmt.Println("Current State:", doorController.stateMachine.MustState()) // Closed
}
