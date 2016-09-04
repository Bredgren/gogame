package key

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

// Key represents a key on a keyboard. It does not differentiate between capitalization,
// (e.g. 'a' and 'A' are the same Key), but it does differentiate between other keys that
// that have dual characters (e.g. '1' and '!', are different Keys). A user of this package
// will usueally find the predefined Keys below enough, and wont need to use the Key struct
// directly.
type Key struct {
	Code int
	// Right is only used for the mod keys: false means the left one and true means right.
	// For non-mod keys this should be false.
	Right bool
	// Upper is used to indicate whether it is the upper or lower version of the key, e.g.
	// '/' is lower (so Upper will be false) and '?' is upper (so it will be true). This
	// field is necessary because the upper and lower versions share the same Code. Upper
	// will always be false for alphebetic keys. Keys with a single use (e.g. space or enter)
	// will also have Upper as false.
	Upper bool
}

var (
	// A - Z are the alphebetic characters. These do not differentiate between upper and
	// lower case, and their string representations are always uppercase.
	A = Key{Code: 65}
	B = Key{Code: 66}
	C = Key{Code: 67}
	D = Key{Code: 68}
	E = Key{Code: 69}
	F = Key{Code: 70}
	G = Key{Code: 71}
	H = Key{Code: 72}
	I = Key{Code: 73}
	J = Key{Code: 74}
	K = Key{Code: 75}
	L = Key{Code: 76}
	M = Key{Code: 77}
	N = Key{Code: 78}
	O = Key{Code: 79}
	P = Key{Code: 80}
	Q = Key{Code: 81}
	R = Key{Code: 82}
	S = Key{Code: 83}
	T = Key{Code: 84}
	U = Key{Code: 85}
	V = Key{Code: 86}
	W = Key{Code: 87}
	X = Key{Code: 88}
	Y = Key{Code: 89}
	Z = Key{Code: 90}

	// D1 - D0 are the digits on the number row at the top of the keyboard.
	D1 = Key{Code: 49}
	D2 = Key{Code: 50}
	D3 = Key{Code: 51}
	D4 = Key{Code: 52}
	D5 = Key{Code: 53}
	D6 = Key{Code: 54}
	D7 = Key{Code: 55}
	D8 = Key{Code: 56}
	D9 = Key{Code: 57}
	D0 = Key{Code: 48}

	// Exclaim - RightParent are the symbols that share the number row keys.
	Exclaim    = Key{Code: 49, Upper: true}
	At         = Key{Code: 50, Upper: true}
	Hash       = Key{Code: 51, Upper: true}
	Dollar     = Key{Code: 52, Upper: true}
	Percent    = Key{Code: 53, Upper: true}
	Caret      = Key{Code: 54, Upper: true}
	Ampersand  = Key{Code: 55, Upper: true}
	Asterisk   = Key{Code: 56, Upper: true}
	LeftParen  = Key{Code: 57, Upper: true}
	RightParen = Key{Code: 48, Upper: true}

	// Backquote - Qusteion are the other punctuation characters.
	Backquote    = Key{Code: 192}
	Tilde        = Key{Code: 192, Upper: true}
	Minus        = Key{Code: 189}
	Underscore   = Key{Code: 189, Upper: true}
	Equals       = Key{Code: 187}
	Plus         = Key{Code: 187, Upper: true}
	LeftBracket  = Key{Code: 219}
	LeftBrace    = Key{Code: 219, Upper: true}
	RightBracket = Key{Code: 221}
	RightBrace   = Key{Code: 221, Upper: true}
	Backslash    = Key{Code: 220}
	Pipe         = Key{Code: 220, Upper: true}
	Semicolon    = Key{Code: 186}
	Colon        = Key{Code: 186, Upper: true}
	Quote        = Key{Code: 222}
	QuoteDbl     = Key{Code: 222, Upper: true}
	Comma        = Key{Code: 188}
	Less         = Key{Code: 188, Upper: true}
	Period       = Key{Code: 190}
	Greater      = Key{Code: 190, Upper: true}
	Slash        = Key{Code: 191}
	Question     = Key{Code: 191, Upper: true}

	// Space - Escape are the spacing related keys.
	Space     = Key{Code: 32}
	Backspace = Key{Code: 8}
	Delete    = Key{Code: 46}
	Tab       = Key{Code: 9}
	Enter     = Key{Code: 13}
	Escape    = Key{Code: 27}

	// Np0 - NpEnter are the number pad keys. We need to use "Right" for some.
	Np0        = Key{Code: 48, Right: true}
	Np1        = Key{Code: 49, Right: true}
	Np2        = Key{Code: 50, Right: true}
	Np3        = Key{Code: 51, Right: true}
	Np4        = Key{Code: 52, Right: true}
	Np5        = Key{Code: 53, Right: true}
	Np6        = Key{Code: 54, Right: true}
	Np7        = Key{Code: 55, Right: true}
	Np8        = Key{Code: 56, Right: true}
	Np9        = Key{Code: 57, Right: true}
	NpPeriod   = Key{Code: 110}
	NpDivide   = Key{Code: 111}
	NpMultiply = Key{Code: 106}
	NpMinus    = Key{Code: 109}
	NpPlus     = Key{Code: 107}
	NpEnter    = Key{Code: 13, Right: true}

	// Up - PageDown are the navigation keys.
	Up       = Key{Code: 38}
	Down     = Key{Code: 40}
	Right    = Key{Code: 39}
	Left     = Key{Code: 37}
	Home     = Key{Code: 36}
	End      = Key{Code: 35}
	PageUp   = Key{Code: 33}
	PageDown = Key{Code: 34}
	// Insert   = "insert"

	// F1 - F15 are the function keys.
	F1  = Key{Code: 112}
	F2  = Key{Code: 113}
	F3  = Key{Code: 114}
	F4  = Key{Code: 115}
	F5  = Key{Code: 116}
	F6  = Key{Code: 117}
	F7  = Key{Code: 118}
	F8  = Key{Code: 119}
	F9  = Key{Code: 120}
	F10 = Key{Code: 121}
	F11 = Key{Code: 122}
	F12 = Key{Code: 123}
	F13 = Key{Code: 124}
	F14 = Key{Code: 125}
	F15 = Key{Code: 126}

	// NumLock - Print are the lock and print screen keys.
	NumLock  = Key{Code: 12}
	CapsLock = Key{Code: 20}
	// ScrolLock = Key{Code: }
	// Print     = Key{Code: }

	// LShift - RMeta are the modifier keys.
	LShift = Key{Code: 16}
	RShift = Key{Code: 16, Right: true}
	LCtrl  = Key{Code: 17}
	RCtrl  = Key{Code: 17, Right: true}
	LAlt   = Key{Code: 18}
	RAlt   = Key{Code: 18, Right: true}
	LMeta  = Key{Code: 91}
	RMeta  = Key{Code: 93, Right: true}
)

var keyName = map[Key]string{
	A:            "a",
	B:            "b",
	C:            "c",
	D:            "d",
	E:            "e",
	F:            "f",
	G:            "g",
	H:            "h",
	I:            "i",
	J:            "j",
	K:            "k",
	L:            "l",
	M:            "m",
	N:            "n",
	O:            "o",
	P:            "p",
	Q:            "q",
	R:            "r",
	S:            "s",
	T:            "t",
	U:            "u",
	V:            "v",
	W:            "w",
	X:            "x",
	Y:            "y",
	Z:            "z",
	D1:           "1",
	D2:           "2",
	D3:           "3",
	D4:           "4",
	D5:           "5",
	D6:           "6",
	D7:           "7",
	D8:           "8",
	D9:           "9",
	D0:           "0",
	Exclaim:      "exclaim",
	At:           "at",
	Hash:         "hash",
	Dollar:       "dollar",
	Percent:      "percent",
	Caret:        "caret",
	Ampersand:    "ampersand",
	Asterisk:     "asterisk",
	LeftParen:    "left parenthesis",
	RightParen:   "right parenthesis",
	Backquote:    "backquote",
	Tilde:        "tilde",
	Minus:        "minus sign",
	Underscore:   "underscore",
	Equals:       "equals sign",
	Plus:         "plus sign",
	LeftBracket:  "left bracket",
	LeftBrace:    "left brace",
	RightBracket: "right bracket",
	RightBrace:   "right brace",
	Backslash:    "backslash",
	Pipe:         "pipe",
	Semicolon:    "semicolon",
	Colon:        "colon",
	Quote:        "quote",
	QuoteDbl:     "quotedbl",
	Comma:        "comma",
	Less:         "less-than sign",
	Period:       "period",
	Greater:      "greater-than sign",
	Slash:        "slash",
	Question:     "question",
	Space:        "Space",
	Backspace:    "Backspace",
	Delete:       "Delete",
	Tab:          "Tab",
	Enter:        "Enter",
	Escape:       "Escape",
	Np0:          "numpad 0",
	Np1:          "numpad 1",
	Np2:          "numpad 2",
	Np3:          "numpad 3",
	Np4:          "numpad 4",
	Np5:          "numpad 5",
	Np6:          "numpad 6",
	Np7:          "numpad 7",
	Np8:          "numpad 8",
	Np9:          "numpad 9",
	NpPeriod:     "numpad period",
	NpDivide:     "numpad divide",
	NpMultiply:   "numpad multiply",
	NpMinus:      "numpad minus",
	NpPlus:       "numpad plus",
	NpEnter:      "numpad enter",
	Up:           "up arrow",
	Down:         "down arrow",
	Right:        "right arrow",
	Left:         "left arrow",
	Home:         "home",
	End:          "end",
	PageUp:       "page up",
	PageDown:     "page down",
	F1:           "f1",
	F2:           "f2",
	F3:           "f3",
	F4:           "f4",
	F5:           "f5",
	F6:           "f6",
	F7:           "f7",
	F8:           "f8",
	F9:           "f9",
	F10:          "f10",
	F11:          "f11",
	F12:          "f12",
	F13:          "f13",
	F14:          "f14",
	F15:          "f15",
	NumLock:      "numlock",
	CapsLock:     "capslock",
	LShift:       "left shift",
	RShift:       "right shift",
	LCtrl:        "left ctrl",
	RCtrl:        "right ctrl",
	LAlt:         "left alt",
	RAlt:         "right alt",
	LMeta:        "left meta",
	RMeta:        "right meta",
	// LSuper:       "left windows key",
	// RSuper:       "right windows key",
}

// IsMod returns true if the key is a modifier key.
func (k Key) IsMod() bool {
	return k == LShift || k == RShift || k == LCtrl || k == RCtrl || k == LAlt || k == RAlt ||
		k == LMeta || k == RMeta
}

var keyRune = map[Key]rune{
	A:            'a',
	B:            'b',
	C:            'c',
	D:            'd',
	E:            'e',
	F:            'f',
	G:            'g',
	H:            'h',
	I:            'i',
	J:            'j',
	K:            'k',
	L:            'l',
	M:            'm',
	N:            'n',
	O:            'o',
	P:            'p',
	Q:            'q',
	R:            'r',
	S:            's',
	T:            't',
	U:            'u',
	V:            'v',
	W:            'w',
	X:            'x',
	Y:            'y',
	Z:            'z',
	D1:           '1',
	D2:           '2',
	D3:           '3',
	D4:           '4',
	D5:           '5',
	D6:           '6',
	D7:           '7',
	D8:           '8',
	D9:           '9',
	D0:           '0',
	Exclaim:      '!',
	At:           '@',
	Hash:         '#',
	Dollar:       '$',
	Percent:      '%',
	Caret:        '^',
	Ampersand:    '&',
	Asterisk:     '*',
	LeftParen:    '(',
	RightParen:   ')',
	Backquote:    '`',
	Tilde:        '~',
	Minus:        '-',
	Underscore:   '_',
	Equals:       '=',
	Plus:         '+',
	LeftBracket:  '[',
	LeftBrace:    '{',
	RightBracket: ']',
	RightBrace:   '}',
	Backslash:    '\\',
	Pipe:         '|',
	Semicolon:    ';',
	Colon:        ':',
	Quote:        '\'',
	QuoteDbl:     '"',
	Comma:        ',',
	Less:         '<',
	Period:       '.',
	Greater:      '>',
	Slash:        '/',
	Question:     '?',
	Space:        ' ',
	Backspace:    '\b',
	Delete:       '\u007f',
	Tab:          '\t',
	Enter:        '\n',
	Escape:       '\u001b',
	Np0:          '0',
	Np1:          '1',
	Np2:          '2',
	Np3:          '3',
	Np4:          '4',
	Np5:          '5',
	Np6:          '6',
	Np7:          '7',
	Np8:          '8',
	Np9:          '9',
	NpPeriod:     '.',
	NpDivide:     '/',
	NpMultiply:   '*',
	NpMinus:      '-',
	NpPlus:       '+',
	NpEnter:      '\n',
	Up:           '\u2191',
	Down:         '\u2193',
	Right:        '\u2192',
	Left:         '\u2190',
}

// String returns a human readable name for the key if one is known, otherwise it returns
// a string representation of the Key struct and its fields.
func (k Key) String() string {
	if s, ok := keyName[k]; ok {
		return s
	}
	return fmt.Sprintf("%#v", k)
}

// Rune returns the printable rune equivalent of the Key if one exists, otherwise it returns
// the unicode null character.
func (k Key) Rune() rune {
	if s, ok := keyRune[k]; ok {
		return s
	}
	return '\u0000'
}

var upperMap = map[rune]rune{
	'1':  '!',
	'2':  '@',
	'3':  '#',
	'4':  '$',
	'5':  '%',
	'6':  '^',
	'7':  '&',
	'8':  '*',
	'9':  '(',
	'0':  ')',
	'`':  '~',
	'-':  '_',
	'=':  '+',
	'[':  '{',
	']':  '}',
	'\\': '|',
	';':  ':',
	'\'': '"',
	',':  '<',
	'.':  '>',
	'/':  '?',
}

var numpadRe = regexp.MustCompile(`Numpad((Enter)|\d+)`)

// FromJsEvent returns the Key for the corresponding js event. It is not necessarily one of
// the predefined Keys.
func FromJsEvent(event *js.Object) Key {
	name := event.Get("code").String()
	k := Key{
		Code: event.Get("keyCode").Int(),
		Right: (strings.Contains(name, "Right") && (strings.Contains(name, "Shift") || strings.Contains(name, "Control") ||
			strings.Contains(name, "Alt") || strings.Contains(name, "Meta"))) ||
			numpadRe.MatchString(name),
	}
	k.Upper = event.Get("key").String() == string(upperMap[k.Rune()])
	return k
}
