package helloworld

import (
	"github.com/Hatch1fy/errors"
	"github.com/gdbu/dbl"
)

const (
	// ErrEmptyUserID is returned when the User ID for an Entry is empty
	ErrEmptyUserID = errors.Error("invalid user ID, cannot be empty")
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

// Validate will ensure an Entry is valid
func (e *Entry) Validate() (err error) {
	// An error list allows us to collect all the errors and return them as a group
	var errs errors.ErrorList
	// Check to see if User ID is set
	if len(e.UserID) == 0 {
		// User ID is empty, append ErrEmptyUserID
		errs.Push(ErrEmptyUserID)
	}

	// Note: If error list is empty, a nil value is returned
	return errs.Err()
}
