package main

import (
	"fmt"

	"github.com/vnscriptkid/LLD-fluency/ticket-transition-v2/models"
	"github.com/vnscriptkid/LLD-fluency/ticket-transition-v2/service"
)

func main() {
	fmt.Println("Hello, World!")

	ts := service.TicketService{}
	user := models.User{Name: "John Doe"}
	ticket := ts.Create("Fix bug", user)

	fmt.Printf("Ticket: %#v\n", ticket)

	ticket.Open()

	fmt.Printf("Ticket: %#v\n", ticket)

	ticket.Inprogress()

	fmt.Printf("Ticket: %#v\n", ticket)
}
