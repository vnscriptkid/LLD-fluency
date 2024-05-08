package models

type User struct {
	Name string
}

type StateMachine interface {
	Open(*Ticket) error
	Inprogress(*Ticket) error
	Resolve(*Ticket) error
}

type Ticket struct {
	Title    string
	Assignee User
	State    StateMachine
}

// Inprogress implements StateMachine.
func (t *Ticket) Inprogress() error {
	return t.State.Inprogress(t)
}

// Open implements StateMachine.
func (t *Ticket) Open() error {
	return t.State.Open(t)
}

// Resolve implements StateMachine.
func (t *Ticket) Resolve() error {
	return t.State.Resolve(t)
}
