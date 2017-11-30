package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

// UI interface definition
type UI interface {
	Widget() gtk.IWidget
	Create() error
	Draw() error
}
