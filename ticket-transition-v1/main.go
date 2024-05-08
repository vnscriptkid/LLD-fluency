package main

import "fmt"

type User struct {
	Name string
}

type TicketState int

const (
	Open TicketState = iota
	InProgress
	Resolved
)

type Ticket struct {
	Title    string
	Assignee User
	State    TicketState
}

type TicketService struct{}

func (ts *TicketService) Transition(ticket *Ticket, destState TicketState) error {

	srcState := ticket.State

	switch srcState {
	case Open:
		switch destState {
		case Open:
			return fmt.Errorf("invalid transition from Open to Open")
		case InProgress:
			ticket.State = InProgress
		case Resolved:
			ticket.State = Resolved
		}
	case InProgress:
		switch destState {
		case Open:
			ticket.State = Open
		case InProgress:
			return fmt.Errorf("invalid transition from InProgress to InProgress")
		case Resolved:
			ticket.State = Resolved
		}
	case Resolved:
		switch destState {
		case Open:
			ticket.State = Open
		case InProgress:
			ticket.State = InProgress
		case Resolved:
			return fmt.Errorf("invalid transition from Resolved to Resolved")
		}
	}
	return nil
}

func (ts *TicketService) Create(title string, user User) Ticket {
	return Ticket{
		Title:    title,
		Assignee: user,
		State:    Open,
	}
}

func main() {
	fmt.Println("Hello, World!")

	ts := TicketService{}
	user := User{Name: "John Doe"}
	ticket := ts.Create("Fix bug", user)

	fmt.Printf("Ticket: %#v\n", ticket)

	err := ts.Transition(&ticket, Open)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Ticket: %#v\n", ticket)

	err = ts.Transition(&ticket, InProgress)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Ticket: %+v\n", ticket)
}
