package states

import (
	"fmt"

	"github.com/vnscriptkid/LLD-fluency/state-machine-door-v3/models"
)

var _ models.State = &ClosedState{}

type ClosedState struct{}

func (s *ClosedState) Open(d *models.Door) {
	fmt.Println("Opening the door.")
	d.SetState(&OpenState{})
}

func (s *ClosedState) Close(d *models.Door) {
	fmt.Println("The door is already closed.")
}

func (s *ClosedState) Lock(d *models.Door) {
	fmt.Println("Locking the door.")
	d.SetState(&LockedState{})
}

func (s *ClosedState) Unlock(d *models.Door) {
	fmt.Println("The door is closed, cannot unlock.")
}
