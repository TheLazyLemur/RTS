package main

type Game struct {
	ui    *UI
	state State
}

type UI struct {
	previousState State
	rectWidth     int32
	rectHeight    int32
}

type State int

const (
	Selection State = iota
	Placing
	Selected
	Interface
)

func (s State) String() string {
	switch s {
	case Selection:
		return "Selection"
	case Placing:
		return "Placing"
	case Selected:
		return "Selected"
	case Interface:
		return "Interface"
	default:
		return "Unknown"
	}
}
