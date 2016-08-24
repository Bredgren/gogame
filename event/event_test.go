package event

import "testing"

func TestPost(t *testing.T) {
	Clear()
	for i := 0; i < MaxEventsInQueue; i++ {
		err := Post(Event{Type: UserEvent})
		if err != nil {
			t.Fatalf("Post failed after %d calls", i+1)
		}
	}
	err := Post(Event{Type: UserEvent})
	if err == nil {
		t.Fatalf("Post did not fail after MaxEventsInQueue+1=%d calls", MaxEventsInQueue)
	}
}

func TestGet(t *testing.T) {
	Clear()
	count := 3
	for i := 0; i < count; i++ {
		err := Post(Event{Type: UserEvent, Data: i})
		if err != nil {
			t.Fatalf("Post %d failed", i)
		}
	}

	events := Get()
	if len(events) != count {
		t.Errorf("Get returned %d events, want %d", len(events), count)
	}
	for want, e := range events {
		got := e.Data.(int)
		if got != want {
			t.Errorf("Event %d had data %d, want %d", want, got, want)
		}
	}
}

func TestGetType(t *testing.T) {
	Clear()
	count := 10
	for i := 0; i < count; i++ {
		err := Post(Event{Type: Type(int(UserEvent) + i%2), Data: i})
		if err != nil {
			t.Fatalf("Post %d failed", i)
		}
	}

	usr := GetType(UserEvent)
	if len(usr) != count/2 {
		t.Errorf("Got %d UserEvents, want %d", len(usr), count/2)
	}

	if len(q) != count/2 {
		t.Errorf("Got %d events left in q, want %d", len(q), count/2)
	}

	for i, e := range usr {
		if e.Type != UserEvent {
			t.Errorf("UserEvent %d had Type %d, want %d", i, e.Type, UserEvent)
		}
	}

	usr1 := GetType(UserEvent + 1)
	if len(usr1) != count/2 {
		t.Errorf("Got %d UserEvent+1, want %d", len(usr1), count/2)
	}

	if len(q) != 0 {
		t.Errorf("Got %d events left in q, want %d", len(q), 0)
	}

	for i, e := range usr1 {
		if e.Type != UserEvent+1 {
			t.Errorf("UserEvent+1 %d had Type %d, want %d", i, e.Type, UserEvent+1)
		}
	}
}

func TestGetTypeList(t *testing.T) {
	Clear()
	count := 12
	for i := 0; i < count; i++ {
		err := Post(Event{Type: Type(int(UserEvent) + i%3), Data: i})
		if err != nil {
			t.Fatalf("Post %d failed", i)
		}
	}

	usr01 := GetTypeList([]Type{UserEvent, UserEvent + 1})
	if len(q) != count/3 {
		t.Errorf("Got %d events left in q, want %d", len(q), count/3)
	}

	for i, e := range usr01 {
		if e.Type != UserEvent && e.Type != UserEvent+1 {
			t.Errorf("usr01[%d] had Type %d, want %d or %d", i, e.Type, UserEvent, UserEvent+1)
		}
	}
}

func TestPoll(t *testing.T) {
	Clear()
	count := 2
	for i := 0; i < count; i++ {
		err := Post(Event{Type: UserEvent, Data: i})
		if err != nil {
			t.Fatalf("Post %d failed", i)
		}
	}

	for i := 0; i < count; i++ {
		e := Poll()
		if e.Type != UserEvent && e.Data.(int) != 0 {
			t.Errorf("Poll %d: got %#v, want, UserEvent/%d", i, e, i)
		}
	}
	e := Poll()
	if e.Type != NoEvent {
		t.Errorf("Didn't get NoEvent when queue was supposed to be empty, got %#v", e)
	}
}

func TestClearType(t *testing.T) {
	Clear()
	count := 12
	for i := 0; i < count; i++ {
		err := Post(Event{Type: Type(int(UserEvent) + i%3), Data: i})
		if err != nil {
			t.Fatalf("Post %d failed", i)
		}
	}

	ClearType(UserEvent)

	events := Get()
	if len(events) != count/3*2 {
		t.Errorf("%d events left in q, want %d", len(events), count/3*2)
	}
	for _, e := range events {
		if e.Type == UserEvent {
			t.Errorf("Event %d with type UserEvent was left in q", e.Data.(int))
		}
	}
}

func TestClearTypeList(t *testing.T) {
	Clear()
	count := 12
	for i := 0; i < count; i++ {
		err := Post(Event{Type: Type(int(UserEvent) + i%3), Data: i})
		if err != nil {
			t.Fatalf("Post %d failed", i)
		}
	}

	ClearTypeList([]Type{UserEvent, UserEvent + 2})

	events := Get()
	if len(events) != count/3 {
		t.Errorf("%d events left in q, want %d", len(events), count/3)
	}
	for _, e := range events {
		if e.Type != UserEvent+1 {
			t.Errorf("Event %d with type %d was left in q", e.Type, e.Data.(int))
		}
	}
}
