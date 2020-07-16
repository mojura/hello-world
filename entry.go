package helloworld

import (
	"github.com/gdbu/dbl"
)

// Entry represents a stored entry within the Controller
type Entry struct {
	// Include dbl.Entry to auto-populate fields/methods needed to match the
	dbl.Entry
}
