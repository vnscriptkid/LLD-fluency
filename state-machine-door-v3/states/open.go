package states

import (
	"fmt"

	"github.com/vnscriptkid/LLD-fluency/state-machine-door-v3/models"
)

var _ models.State = &OpenState{}

type OpenState struct{}

func (s *OpenState) Open(d *models.Door) {
	fmt.Println("The door is already open.")
}

func (s *OpenState) Close(d *models.Door) {
	fmt.Println("Closing the door.")
	d.SetState(&ClosedState{})
}

func (s *OpenState) Lock(d *models.Door) {
	fmt.Println("Close the door before locking.")
}

func (s *OpenState) Unlock(d *models.Door) {
	fmt.Println("The door is open, cannot unlock.")
}
