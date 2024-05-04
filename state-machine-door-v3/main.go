package main

import (
	"github.com/vnscriptkid/LLD-fluency/state-machine-door-v3/models"
	"github.com/vnscriptkid/LLD-fluency/state-machine-door-v3/states"
)

func main() {
	door := models.NewDoor(&states.ClosedState{})
	door.Open()
	door.Close()
	door.Lock()
	door.Unlock()
}
