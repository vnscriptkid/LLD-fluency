package statemachine

import "github.com/vnscriptkid/LLD-fluency/ticket-transition-v2/models"

var _ models.StateMachine = &StateResolved{}

type StateResolved struct{}

// Inprogress implements models.StateMachine.
func (s *StateResolved) Inprogress(*models.Ticket) error {
	panic("unimplemented")
}

// Open implements models.StateMachine.
func (s *StateResolved) Open(*models.Ticket) error {
	panic("unimplemented")
}

// Resolve implements models.StateMachine.
func (s *StateResolved) Resolve(*models.Ticket) error {
	panic("unimplemented")
}
