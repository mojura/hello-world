package helloworld

import (
	"github.com/gdbu/dbl"
)

// Entry represents a stored entry within the Controller
type Entry struct {
	// Include dbl.Entry to auto-populate fields/methods needed to match the
	dbl.Entry

	// UserID which Entry is related to, used a relational referenced
	UserID string `json:"userID"`
}

// GetRelationshipIDs will return the relationship IDs associated with the Entry
func (e *Entry) GetRelationshipIDs() (ids []string) {
	// UserID is our only relationship at the moment. Any number of relationships can
	// be set for entries. Just know that each added relationship is an extra reference
	// to be managed by DBL and stored by the underlying back-end.
	ids = append(ids, e.UserID)
	return
}
