package key

// Key represents a key on a keyboard. For alphabetic characters the string representation
// will always be capitalized.
type Key string

const (
	A = "A"
	B = "B"
)

func (k Key) String() string {
	return string(k)
}
