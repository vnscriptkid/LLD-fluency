package service

import (
	"github.com/vnscriptkid/LLD-fluency/ticket-transition-v2/models"
	"github.com/vnscriptkid/LLD-fluency/ticket-transition-v2/statemachine"
)

type TicketService struct{}

func (ts *TicketService) Create(title string, user models.User) models.Ticket {
	return models.Ticket{
		Title:    title,
		Assignee: user,
		State:    &statemachine.StateOpen{},
	}
}
