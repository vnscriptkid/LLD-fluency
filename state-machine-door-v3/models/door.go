package models

type State interface {
	Open(door *Door)
	Close(door *Door)
	Lock(door *Door)
	Unlock(door *Door)
}

type Door struct {
	state State
}

func NewDoor(state State) *Door {
	return &Door{state: state}
}

func (d *Door) SetState(state State) {
	d.state = state
}

func (d *Door) Open() {
	d.state.Open(d)
}

func (d *Door) Close() {
	d.state.Close(d)
}

func (d *Door) Lock() {
	d.state.Lock(d)
}

func (d *Door) Unlock() {
	d.state.Unlock(d)
}
