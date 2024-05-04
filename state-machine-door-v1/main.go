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

func (d *Door) Open() {
	switch d.state {
	case Open:
		fmt.Println("The door is already open.")
	case Closed:
		fmt.Println("Opening the door.")
		d.state = Open
	case Locked:
		fmt.Println("The door is locked, cannot open.")
	}
}

func (d *Door) Close() {
	switch d.state {
	case Open:
		fmt.Println("Closing the door.")
		d.state = Closed
	case Closed:
		fmt.Println("The door is already closed.")
	case Locked:
		fmt.Println("The door is locked, cannot close.")
	}
}

func (d *Door) Lock() {
	switch d.state {
	case Open:
		fmt.Println("Close the door before locking.")
	case Closed:
		fmt.Println("Locking the door.")
		d.state = Locked
	case Locked:
		fmt.Println("The door is already locked.")
	}
}

func (d *Door) Unlock() {
	switch d.state {
	case Open:
		fmt.Println("The door is open, cannot unlock.")
	case Closed:
		fmt.Println("The door is closed, cannot unlock.")
	case Locked:
		fmt.Println("Unlocking the door.")
		d.state = Closed
	}
}

func main() {
	door := Door{state: Closed}
	door.Open()
	door.Close()
	door.Lock()
	door.Unlock()
	door.Open()
	door.Lock()
}
