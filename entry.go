package helloworld

import (
	"github.com/hatchify/errors"
	"github.com/mojura/mojura"
)

const (
	// ErrEmptyUserID is returned when the User ID for an Entry is empty
	ErrEmptyUserID = errors.Error("invalid user ID, cannot be empty")
	// ErrEmptyGreeting is returned when the Greeting for an Entry is empty
	ErrEmptyGreeting = errors.Error("invalid greeting, cannot be empty")
	// ErrEmptyFavoriteTimeOfDay is returned when the Favorite time of day for an Entry is empty
	ErrEmptyFavoriteTimeOfDay = errors.Error("invalid favorite time of day, cannot be empty")
)

// Entry represents a stored entry within the Controller
type Entry struct {
	// Include mojura.Entry to auto-populate fields/methods needed to match the
	mojura.Entry

	// UserID which Entry is related to, used a relational referenced
	UserID string `json:"userID"`

	// Greeting represents the favorite greeting of the related user
	Greeting string `json:"greeting"`
	// FavoriteTimeOfDay represents the favorite time of day of the related user
	FavoriteTimeOfDay string `json:"favoriteTimeOfDay"`
}

// GetRelationshipIDs will return the relationship IDs associated with the Entry
func (e *Entry) GetRelationships() (r mojura.Relationships) {
	// UserID is our only relationship at the moment. Any number of relationships can
	// be set for entries. Just know that each added relationship is an extra reference
	// to be managed by Mojua and stored by the underlying back-end.
	r.Append(e.UserID)
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

	if len(e.Greeting) == 0 {
		// Greeting is empty, append ErrEmptyGreeting
		errs.Push(ErrEmptyGreeting)
	}

	if len(e.FavoriteTimeOfDay) == 0 {
		// Favorite time of day is empty, append ErrEmptyFavoriteTimeOfDay
		errs.Push(ErrEmptyFavoriteTimeOfDay)
	}

	// Note: If error list is empty, a nil value is returned
	return errs.Err()
}
