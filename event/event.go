// Package event manages a single queue for all events, including input, window, and
// user-defined events.
package event

import (
	"fmt"

	"github.com/Bredgren/gogame/key"
)

// MaxEventsInQueue is number of events the queue can hold before it starts dropping new
// events.
const MaxEventsInQueue = 256

var q = make(chan Event, MaxEventsInQueue)

// Get returns all events and removes them from the queue.
func Get() []Event {
	var events []Event
	for len(q) > 0 {
		events = append(events, <-q)

	}
	return events
}

// GetType returns all events in the queue of the given type. Note that if you are mostly
// using this function (or GetTypeList) then the queue may eventually fill up with events
// you are not interested in.
func GetType(t Type) []Event {
	var events []Event
	count := len(q)
	for i := 0; i < count; i++ {
		e := <-q
		if e.Type == t {
			events = append(events, e)
			continue
		}
		q <- e
	}
	return events
}

// GetTypeList returns all events in the queue that match one of the types given. Note
// that if you are mostly using this function (or GetType) then the queue may eventually
// fill up with events you are not interested in.
func GetTypeList(ts []Type) []Event {
	var events []Event
	count := len(q)
Outer:
	for i := 0; i < count; i++ {
		e := <-q
		for _, t := range ts {
			if e.Type == t {
				events = append(events, e)
				continue Outer
			}
		}
		q <- e
	}
	return events
}

// Poll removes and returns the next event in the queue. If there are no events it returns
// an Event with type NoEvent.
func Poll() Event {
	select {
	case e := <-q:
		return e
	default:
		return Event{Type: NoEvent}
	}
}

// Peak returns true if there is an event of the given type on the queue.
// func Peak(t Type) bool {}
// PeakList returns true if an event that matches any of the given types are on the queue.
// func PeakList(ts []Type) bool {}

// Clear removes all events form the queue.
func Clear() {
	for len(q) > 0 {
		_ = <-q
	}
}

//ClearType removes all events of the given type form the queue.
func ClearType(t Type) {
	count := len(q)
	for i := 0; i < count; i++ {
		e := <-q
		if e.Type == t {
			continue
		}
		q <- e
	}
}

// ClearTypeList removes all events with the given types form the queue.
func ClearTypeList(ts []Type) {
	count := len(q)
Outer:
	for i := 0; i < count; i++ {
		e := <-q
		for _, t := range ts {
			if e.Type == t {
				continue Outer
			}
		}
		q <- e
	}
}

// Post pushes the given event onto the queue. If the queue is already full then a non-nil
// error is returned.
func Post(e Event) error {
	if len(q) == MaxEventsInQueue {
		return fmt.Errorf("event queue is full")
	}
	q <- e
	return nil
}

// Event has a type and arbitrary data. See the documentation for each Type for a description
// of what data will be returned.
type Event struct {
	Type Type
	Data interface{}
}

// KeyData holds a key that is of interest to the event and also any modifier keys that were
// held down at the time of the event.
type KeyData struct {
	Key key.Key
	Mod map[key.Key]bool
}

// MouseMotionData holds the position of the mouse relative to the uppert left corner of
// the display, the position relative to it's previous position, and which buttons where
// held down.
type MouseMotionData struct {
	Pos     struct{ X, Y float64 }
	Rel     struct{ X, Y float64 }
	Buttons map[int]bool
}

// MouseWheelData holds the delta x, y, and z for the mouse wheel.
type MouseWheelData struct {
	Dx, Dy, Dz float64
}

// MouseData holds the position of the mouse relative to the uppert left corner of the
// display and the button of interest to the event.
type MouseData struct {
	Pos    struct{ X, Y float64 }
	Button int
}

// ResizeData holds the new width and height.
type ResizeData struct {
	W, H int
}

// Type is used to distinguish different kinds of events. Custom types may be defined and
// should start from UserEvent.
type Type int

const (
	// NoEvent is used when an Event is expected to be returned but there is no Event available
	// to return. It has no data assiciated with it.
	NoEvent Type = iota
	// Quit signals the user closing/reloading the page. It has no data assiciated with it.
	Quit
	// KeyDown is when a key on the keyboard is pressed down. Its data will be of type KeyData.
	KeyDown
	// KeyUp is when a key on the keyboard is released. Its data will be of type KeyData.
	KeyUp
	// MouseMotion is when the mouse moves. Its data will be of type MouseMotionData.
	MouseMotion
	// MouseButtonDown is when a button on the mouse is pressed down. Its data will be of
	// type MouseData.
	MouseButtonDown
	// MouseButtonUp is when a button on the mouse is released. Its data will be of type MouseData.
	MouseButtonUp
	// MouseWheel is when the mouse wheel is moved.
	MouseWheel
	// VideoResize is when the window's dimensions change. Its data will be of type ResizeData.
	VideoResize
	// UserEvent is the base for user events. Users may define custom Types but they're value
	// shuold be at least equal to UserEvent. Their data is whatever the user defines it to be.
	UserEvent
)
