package statemachine

import (
	"fmt"

	"github.com/vnscriptkid/LLD-fluency/ticket-transition-v2/models"
)

var _ models.StateMachine = &StateOpen{}

type StateOpen struct {
}

// Inprogress implements models.StateMachine.
func (s *StateOpen) Inprogress(t *models.Ticket) error {
	t.State = &StateInprogress{}
	return nil
}

// Open implements models.StateMachine.
func (s *StateOpen) Open(*models.Ticket) error {
	return fmt.Errorf("failed to move from OPEN to OPEN")
}

// Resolve implements models.StateMachine.
func (s *StateOpen) Resolve(*models.Ticket) error {
	panic("unimplemented")
}
