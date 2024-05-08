package statemachine

import (
	"fmt"

	"github.com/vnscriptkid/LLD-fluency/ticket-transition-v2/models"
)

var _ models.StateMachine = &StateInprogress{}

type StateInprogress struct{}

// Inprogress implements models.StateMachine.
func (s *StateInprogress) Inprogress(t *models.Ticket) error {
	return fmt.Errorf("failed to move from INPROGRESS to INPROGRESS")
}

// Open implements models.StateMachine.
func (s *StateInprogress) Open(t *models.Ticket) error {
	t.State = &StateOpen{}
	return nil
}

// Resolve implements models.StateMachine.
func (s *StateInprogress) Resolve(t *models.Ticket) error {
	t.State = &StateResolved{}
	return nil
}
