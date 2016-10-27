// Package key provides a type that represents keys on a keyboard, and predefines most
// keys.
package key

// Key represents a key on a keyboard.
type Key struct {
	// Name is a unique string that identifies the key.
	Name string

	// Rune is the printable character for this key if there is one, otherwise it's the
	// null character.
	Rune rune

	// ShifRune is the printable character for this key when combined with the shift key.
	// If there isn't one then this is the null character.
	ShiftRune rune
}

var (
	// Alphabetic.
	A = Key{Name: "A", Rune: 'a', ShiftRune: 'A'}
	B = Key{Name: "B", Rune: 'b', ShiftRune: 'B'}
	C = Key{Name: "C", Rune: 'c', ShiftRune: 'C'}
	D = Key{Name: "D", Rune: 'd', ShiftRune: 'D'}
	E = Key{Name: "E", Rune: 'e', ShiftRune: 'E'}
	F = Key{Name: "F", Rune: 'f', ShiftRune: 'F'}
	G = Key{Name: "G", Rune: 'g', ShiftRune: 'G'}
	H = Key{Name: "H", Rune: 'h', ShiftRune: 'H'}
	I = Key{Name: "I", Rune: 'i', ShiftRune: 'I'}
	J = Key{Name: "J", Rune: 'j', ShiftRune: 'J'}
	K = Key{Name: "K", Rune: 'k', ShiftRune: 'K'}
	L = Key{Name: "L", Rune: 'l', ShiftRune: 'L'}
	M = Key{Name: "M", Rune: 'm', ShiftRune: 'M'}
	N = Key{Name: "N", Rune: 'n', ShiftRune: 'N'}
	O = Key{Name: "O", Rune: 'o', ShiftRune: 'O'}
	P = Key{Name: "P", Rune: 'p', ShiftRune: 'P'}
	Q = Key{Name: "Q", Rune: 'q', ShiftRune: 'Q'}
	R = Key{Name: "R", Rune: 'r', ShiftRune: 'R'}
	S = Key{Name: "S", Rune: 's', ShiftRune: 'S'}
	T = Key{Name: "T", Rune: 't', ShiftRune: 'T'}
	U = Key{Name: "U", Rune: 'u', ShiftRune: 'U'}
	V = Key{Name: "V", Rune: 'v', ShiftRune: 'V'}
	W = Key{Name: "W", Rune: 'w', ShiftRune: 'W'}
	X = Key{Name: "X", Rune: 'x', ShiftRune: 'X'}
	Y = Key{Name: "Y", Rune: 'y', ShiftRune: 'Y'}
	Z = Key{Name: "Z", Rune: 'z', ShiftRune: 'Z'}

	// Number row.
	D0 = Key{Name: "0/left paren", Rune: '0', ShiftRune: ')'}
	D1 = Key{Name: "1/exclamation", Rune: '1', ShiftRune: '!'}
	D2 = Key{Name: "2/at", Rune: '2', ShiftRune: '@'}
	D3 = Key{Name: "3/hash", Rune: '3', ShiftRune: '#'}
	D4 = Key{Name: "4/dollar", Rune: '4', ShiftRune: '$'}
	D5 = Key{Name: "5/percent", Rune: '5', ShiftRune: '%'}
	D6 = Key{Name: "6/caret", Rune: '6', ShiftRune: '^'}
	D7 = Key{Name: "7/ampersand", Rune: '7', ShiftRune: '&'}
	D8 = Key{Name: "8/asterisk", Rune: '8', ShiftRune: '*'}
	D9 = Key{Name: "9/right paren", Rune: '9', ShiftRune: '('}

	// Punctuation.
	Backquote    = Key{Name: "backquote/tilde", Rune: '`', ShiftRune: '~'}
	Minus        = Key{Name: "minus/underscore", Rune: '-', ShiftRune: '_'}
	Equal        = Key{Name: "equal/plus", Rune: '=', ShiftRune: '+'}
	LeftBracket  = Key{Name: "left bracket/brace", Rune: '[', ShiftRune: '{'}
	RightBracket = Key{Name: "right bracket/brace", Rune: ']', ShiftRune: '}'}
	Backslash    = Key{Name: "backslash/pipe", Rune: '\\', ShiftRune: '|'}
	Semicolon    = Key{Name: "semicolon/colon", Rune: ';', ShiftRune: ':'}
	Quote        = Key{Name: "quote/double quote", Rune: '\'', ShiftRune: '"'}
	Comma        = Key{Name: "comma/less than", Rune: ',', ShiftRune: '<'}
	Period       = Key{Name: "period/greater than", Rune: '.', ShiftRune: '>'}
	Slash        = Key{Name: "slash/question", Rune: '/', ShiftRune: '?'}

	// Spacing.
	Space     = Key{Name: "space", Rune: ' '}
	Backspace = Key{Name: "backspace", Rune: '\u0008'}
	Delete    = Key{Name: "delete", Rune: '\u007f'}
	Tab       = Key{Name: "tab", Rune: '\u0009'}
	Enter     = Key{Name: "enter", Rune: '\n'}
	Escape    = Key{Name: "escape", Rune: '\u001b'}

	// Numberpad. Names are the same as number row, but there's no ShiftRune.
	Np0        = Key{Name: "numpad 0", Rune: '0'}
	Np1        = Key{Name: "numpad 1", Rune: '1'}
	Np2        = Key{Name: "numpad 2", Rune: '2'}
	Np3        = Key{Name: "numpad 3", Rune: '3'}
	Np4        = Key{Name: "numpad 4", Rune: '4'}
	Np5        = Key{Name: "numpad 5", Rune: '5'}
	Np6        = Key{Name: "numpad 6", Rune: '6'}
	Np7        = Key{Name: "numpad 7", Rune: '7'}
	Np8        = Key{Name: "numpad 8", Rune: '8'}
	Np9        = Key{Name: "numpad 9", Rune: '9'}
	NpPeriod   = Key{Name: "numpad period", Rune: '.'}
	NpDivide   = Key{Name: "numpad divide", Rune: '/'}
	NpMultiply = Key{Name: "numpad multiply", Rune: '*'}
	NpMinus    = Key{Name: "numpad minus", Rune: '-'}
	NpPlus     = Key{Name: "numpad plus", Rune: '+'}
	NpEnter    = Key{Name: "numpad enter", Rune: '\r'}

	// Navigation.
	Left     = Key{Name: "left", Rune: '\u2190'}
	Up       = Key{Name: "up", Rune: '\u2191'}
	Right    = Key{Name: "right", Rune: '\u2192'}
	Down     = Key{Name: "down", Rune: '\u2193'}
	Home     = Key{Name: "home", Rune: '\u21e6'}
	PageUp   = Key{Name: "page up", Rune: '\u21e7'}
	End      = Key{Name: "end", Rune: '\u21e8'}
	PageDown = Key{Name: "page down", Rune: '\u21e9'}

	// Function.
	F1  = Key{Name: "F1"}
	F2  = Key{Name: "F2"}
	F3  = Key{Name: "F3"}
	F4  = Key{Name: "F4"}
	F5  = Key{Name: "F5"}
	F6  = Key{Name: "F6"}
	F7  = Key{Name: "F7"}
	F8  = Key{Name: "F8"}
	F9  = Key{Name: "F9"}
	F10 = Key{Name: "F10"}
	F11 = Key{Name: "F11"}
	F12 = Key{Name: "F12"}
	F13 = Key{Name: "F13"}
	F14 = Key{Name: "F14"}
	F15 = Key{Name: "F15"}

	// Modifiers.
	LShift = Key{Name: "left shift"}
	RShift = Key{Name: "right shift"}
	LCtrl  = Key{Name: "left ctrl"}
	RCtrl  = Key{Name: "right ctrl"}
	LAlt   = Key{Name: "left alt"}
	RAlt   = Key{Name: "right alt"}
	LMeta  = Key{Name: "left meta"}
	RMeta  = Key{Name: "right meta"}

	// Misc.
	NumLock  = Key{Name: "num lock"}
	CapsLock = Key{Name: "caps lock"}
)

// IsMod returns true if the key is a modifier key.
func (k Key) IsMod() bool {
	return k == LShift || k == RShift || k == LCtrl || k == RCtrl || k == LAlt || k == RAlt ||
		k == LMeta || k == RMeta
}

func (k Key) String() string {
	return k.Name
}
