package ggweb

import (
	"github.com/Bredgren/gogame/key"
	"github.com/gopherjs/gopherjs/js"
)

// EventToKey takes a js event and returns the key.Key described by the event.
func EventToKey(evt *js.Object) key.Key {
	return jsToKey[evt.Get("code").String()]
}

var jsToKey = map[string]key.Key{
	"KeyA":           key.A,
	"KeyB":           key.B,
	"KeyC":           key.C,
	"KeyD":           key.D,
	"KeyE":           key.E,
	"KeyF":           key.F,
	"KeyG":           key.G,
	"KeyH":           key.H,
	"KeyI":           key.I,
	"KeyJ":           key.J,
	"KeyK":           key.K,
	"KeyL":           key.L,
	"KeyM":           key.M,
	"KeyN":           key.N,
	"KeyO":           key.O,
	"KeyP":           key.P,
	"KeyQ":           key.Q,
	"KeyR":           key.R,
	"KeyS":           key.S,
	"KeyT":           key.T,
	"KeyU":           key.U,
	"KeyV":           key.V,
	"KeyW":           key.W,
	"KeyX":           key.X,
	"KeyY":           key.Y,
	"KeyZ":           key.Z,
	"Digit0":         key.D0,
	"Digit1":         key.D1,
	"Digit2":         key.D2,
	"Digit3":         key.D3,
	"Digit4":         key.D4,
	"Digit5":         key.D5,
	"Digit6":         key.D6,
	"Digit7":         key.D7,
	"Digit8":         key.D8,
	"Digit9":         key.D9,
	"Backquote":      key.Backquote,
	"Minus":          key.Minus,
	"Equal":          key.Equal,
	"BracketLeft":    key.LeftBracket,
	"BracketRight":   key.RightBracket,
	"Backslash":      key.Backslash,
	"Semicolon":      key.Semicolon,
	"Quote":          key.Quote,
	"Comma":          key.Comma,
	"Period":         key.Period,
	"Slash":          key.Slash,
	"Space":          key.Space,
	"Backspace":      key.Backspace,
	"Delete":         key.Delete,
	"Tab":            key.Tab,
	"Enter":          key.Enter,
	"Escape":         key.Escape,
	"Numpad0":        key.Np0,
	"Numpad1":        key.Np1,
	"Numpad2":        key.Np2,
	"Numpad3":        key.Np3,
	"Numpad4":        key.Np4,
	"Numpad5":        key.Np5,
	"Numpad6":        key.Np6,
	"Numpad7":        key.Np7,
	"Numpad8":        key.Np8,
	"Numpad9":        key.Np9,
	"NumpadDecimal":  key.NpPeriod,
	"NumpadDivide":   key.NpDivide,
	"NumpadMultiply": key.NpMultiply,
	"NumpadSubtract": key.NpMinus,
	"NumpadAdd":      key.NpPlus,
	"NumpadEnter":    key.NpEnter,
	"ArrowLeft":      key.Left,
	"ArrowUp":        key.Up,
	"ArrowRight":     key.Right,
	"ArrowDown":      key.Down,
	"Home":           key.Home,
	"PageUp":         key.PageUp,
	"End":            key.End,
	"PageDown":       key.PageDown,
	"F1":             key.F1,
	"F2":             key.F2,
	"F3":             key.F3,
	"F4":             key.F4,
	"F5":             key.F5,
	"F6":             key.F6,
	"F7":             key.F7,
	"F8":             key.F8,
	"F9":             key.F9,
	"F10":            key.F10,
	"F11":            key.F11,
	"F12":            key.F12,
	"F13":            key.F13,
	"F14":            key.F14,
	"F15":            key.F15,
	"ShiftLeft":      key.LShift,
	"ShiftRight":     key.RShift,
	"ControlLeft":    key.LCtrl,
	"ControlRight":   key.RCtrl,
	"AltLeft":        key.LAlt,
	"AltRight":       key.RAlt,
	"MetaLeft":       key.LMeta,
	"MetaRight":      key.RMeta,
	"NumLock":        key.NumLock,
	"CapsLock":       key.CapsLock,
}
