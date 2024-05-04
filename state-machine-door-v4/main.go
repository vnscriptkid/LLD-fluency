package main

import (
	"context"
	"fmt"

	"github.com/qmuntal/stateless"
)

type Door struct {
	stateMachine *stateless.StateMachine
	keyCode      string
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

func NewDoor(keyCode string) *Door {
	sm := stateless.NewStateMachine(Closed)

	// PermitReentry: you want to perform some action even if the state doesn't change.
	// Ignore: you want to discard the trigger without any action

	door := &Door{stateMachine: sm, keyCode: keyCode}

	sm.Configure(Open).
		Permit(CloseTrigger, Closed).
		Ignore(OpenTrigger).
		PermitReentry(UnlockTrigger)

	sm.Configure(Closed).
		Permit(OpenTrigger, Open).
		Permit(LockTrigger, Locked).
		Ignore(CloseTrigger).
		PermitReentry(UnlockTrigger).
		OnEntryFrom(UnlockTrigger, door.playMusic)

	sm.Configure(Locked).
		Permit(UnlockTrigger, Closed, door.isKeyCodeValid).
		Ignore(LockTrigger).
		Ignore(CloseTrigger).
		Ignore(OpenTrigger)

	return door
}

func (d *Door) isKeyCodeValid(ctx context.Context, args ...any) bool {
	keyCode := args[0].(string)
	isValid := keyCode == d.keyCode

	if !isValid {
		fmt.Println("Invalid key code")
		return false
	}

	return true
}

func (d *Door) playMusic(ctx context.Context, args ...any) error {
	fmt.Println("Playing music when door is unlocked.")
	return nil
}

func (d *Door) Open() {
	err := d.stateMachine.Fire(OpenTrigger)
	if err != nil {
		fmt.Println("Failed to open:", err)
	}
}

func (d *Door) Close() {
	err := d.stateMachine.Fire(CloseTrigger)
	if err != nil {
		fmt.Println("Failed to close:", err)
	}
}

func (d *Door) Lock() {
	err := d.stateMachine.Fire(LockTrigger)
	if err != nil {
		fmt.Println("Failed to lock:", err)
	}
}

func (d *Door) Unlock(keyCode string) {
	err := d.stateMachine.FireCtx(context.Background(), UnlockTrigger, keyCode)
	if err != nil {
		fmt.Println("Failed to unlock:", err)
	}
}

func main() {
	door := NewDoor("123")

	fmt.Println("Current State:", door.stateMachine.MustState()) // Closed
	door.Open()
	fmt.Println("Current State:", door.stateMachine.MustState()) // Open
	door.Open()
	fmt.Println("Current State:", door.stateMachine.MustState()) // Open
	door.Close()
	fmt.Println("Current State:", door.stateMachine.MustState()) // Closed
	door.Lock()
	fmt.Println("Current State:", door.stateMachine.MustState()) // Locked

	door.Unlock("9999")
	fmt.Println("Current State:", door.stateMachine.MustState()) // Locked

	door.Unlock("123")
	fmt.Println("Current State:", door.stateMachine.MustState()) // Closed
}
