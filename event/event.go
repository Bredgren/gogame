package event

// MaxEventsInQueue is number of events the queue can hold before it starts dropping new
// events.
const MaxEventsInQueue = 256

var q = make(chan Event, MaxEventsInQueue)

// Get returns all events and removes them from the queue.
func Get() []Event {
	return nil
}

// GetType returns all events in the queue of the given type. Note that if you are mostly
// using this function (or GetTypeList) then the queue may eventually fill up with events
// you are not interested in.
func GetType(t Type) []Event {
	return nil
}

// GetTypeList returns all events in the queue that match one of the types given. Note
// that if you are mostly using this function (or GetType) then the queue may eventually
// fill up with events you are not interested in.
func GetTypeList(ts []Type) []Event {
	return nil
}

// Poll removes and returns the next event in the queue. If there are no events it returns
// NoEvent.
func Poll() Event {
	return nil
}

// func Wait() {
// }
//
// func Clear() {
// }
// ...

// Event is a single event.
type Event interface {
	Type() Type
}

// Type is used to distinguish different kinds of events.
type Type int
