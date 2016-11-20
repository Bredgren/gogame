package fsm

import "fmt"

// State is the name of a state in the finite state machine.
type State string

const (
	// InitialState is the initial state that an FSM is in.
	InitialState State = ""
)

// Transition defines what action to take when transitioning between the states From and To.
type Transition struct {
	From, To State
	Action   func()
}

// FSM is a finite state machine. Transitions is a list of valid state transitions. The
// initial state is the empty string, and transitions must include at least one transition
// from this initial state. If multiple transitions exist for the same From-To pair then
// each of the corresponding actions will be taken in the order the appaer in the list.
type FSM struct {
	Transitions []*Transition
	current     State
	previous    State
}

// CurrentState returns the current state.
func (f *FSM) CurrentState() State {
	return f.current
}

// PreviousState returns the previous state. This will be initally be equal to the current
// state (the empty string), until the state has changed at least once.
func (f *FSM) PreviousState() State {
	return f.previous
}

// GotoState transitions to the new state given and calls the associated action callback
// if there is one. The callback takes place after the current and previous states have
// been updated. If the given state is not a valid transition then an error is returned.
func (f *FSM) GotoState(s State) error {
	err := fmt.Errorf("cannot transition to state '%s' from '%s'", s, f.current)
	current := f.current
	for _, t := range f.Transitions {
		if t.To == s && t.From == current {
			f.previous = f.current
			f.current = t.To
			t.Action()
			err = nil
		}
	}
	return err
}

// ValidTransitions returns a list of valid states that can be transitioned to from the
// given state.
func (f *FSM) ValidTransitions(s State) []State {
	added := map[State]bool{}
	valid := []State{}
	for _, t := range f.Transitions {
		if !added[t.To] && t.From == s {
			valid = append(valid, t.To)
			added[t.To] = true
		}
	}
	return valid
}
