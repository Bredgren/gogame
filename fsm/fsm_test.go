package fsm

import "testing"

func TestFSM(t *testing.T) {
	sm := FSM{}
	if sm.Current() != InitialState {
		t.Errorf("Current is %s, wanted %s", sm.Current(), InitialState)
	}
	if sm.Previous() != InitialState {
		t.Errorf("Previous is %s, wanted %s", sm.Previous(), InitialState)
	}

	EmptyToState1 := 0
	State1ToState2 := 0
	State2ToState1 := 0
	State1ToState3 := 0
	sm.Transitions = []*Transition{
		{"", "State1", func() {
			EmptyToState1++
		}},
		{"State1", "State2", func() {
			State1ToState2++
		}},
		{"State2", "State1", func() {
			State2ToState1++
		}},
		{"State1", "State3", func() {
			t.Log("1 -> 3")
			State1ToState3++
		}},
		{"State1", "State3", func() {
			t.Log("1 -> 3 other")
			State1ToState3 *= 3 // so we can check that it is (0 + 1) * 3 and not (0 * 3) + 1
		}},
	}

	s1T := sm.ValidTransitions("State1")
	expectedS1T := []State{"State2", "State3"}
	if len(s1T) != len(expectedS1T) || expectedS1T[0] != s1T[0] || expectedS1T[1] != s1T[1] {
		t.Errorf("Expected transitions %#v, got %#v", expectedS1T, s1T)
	}

	err := sm.Goto("State2")
	if err == nil {
		t.Errorf("Transition to State2 from initial succeeded: %s -> %s", sm.Previous(), sm.Current())
	}

	err = sm.Goto("State1")
	if err != nil {
		t.Error(err)
	}
	if EmptyToState1 != 1 {
		t.Errorf("EmptyToState1 is %d, expected 1", EmptyToState1)
	}

	err = sm.Goto("State2")
	if err != nil {
		t.Error(err)
	}
	if State1ToState2 != 1 {
		t.Errorf("State1ToState2 is %d, expected 1", State1ToState2)
	}

	err = sm.Goto("State1")
	if err != nil {
		t.Error(err)
	}
	if State2ToState1 != 1 {
		t.Errorf("State2ToState1 is %d, expected 1", State2ToState1)
	}

	err = sm.Goto("State3")
	if err != nil {
		t.Error(err)
	}
	// Check goto current state
	err = sm.Goto("State3")
	if err != nil {
		t.Error(err)
	}
	if State1ToState3 != 3 {
		t.Errorf("State1ToState3 is %d, expected 3", State1ToState3)
	}
}
