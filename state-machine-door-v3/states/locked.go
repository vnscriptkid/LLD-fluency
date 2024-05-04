package states

import (
	"fmt"

	"github.com/vnscriptkid/LLD-fluency/state-machine-door-v3/models"
)

var _ models.State = &LockedState{}

type LockedState struct{}

func (s *LockedState) Open(d *models.Door) {
	fmt.Println("The door is locked, cannot open.")
}

func (s *LockedState) Close(d *models.Door) {
	fmt.Println("The door is locked, cannot close.")
}

func (s *LockedState) Lock(d *models.Door) {
	fmt.Println("The door is already locked.")
}

func (s *LockedState) Unlock(d *models.Door) {
	fmt.Println("Unlocking the door.")
	d.SetState(&ClosedState{})
}
