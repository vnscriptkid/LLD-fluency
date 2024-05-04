package main

import (
	"fmt"
)

const (
	Open = iota
	Closed
	Locked
)

type Door struct {
	state int
}

func changeState(d *Door, state int) {
	currentState := d.state

	switch currentState {
	case Open:
		switch state {
		case Open:
			fmt.Println("[x] The door is already open.")
		case Closed:
			fmt.Println("Closing the door.")
			d.state = Closed
		case Locked:
			fmt.Println("[x] Close the door before locking.")
		}
	case Closed:
		switch state {
		case Open:
			fmt.Println("Opening the door.")
			d.state = Open
		case Closed:
			fmt.Println("[x] The door is already closed.")
		case Locked:
			fmt.Println("Locking the door.")
			d.state = Locked
		}
	case Locked:
		switch state {
		case Open:
			fmt.Println("[x] Need to unlock the door first.")
		case Closed:
			fmt.Println("Unlocking the door.")
			d.state = Closed
		case Locked:
			fmt.Println("[x] The door is already locked.")
		}
	}
}

func main() {
	door := Door{state: Open}

	changeState(&door, Closed)
	changeState(&door, Locked)
	changeState(&door, Open)

}
