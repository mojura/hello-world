package helloworld

import (
	"github.com/mojura/mojura"
	"github.com/mojura/mojura/filters"
)

// Relationship key const block
const (
	// relationshipUsers represents the key for the UserID relationship
	relationshipUsers = "users"
)

// relationships is a collection of all the supported relationship keys
// Note: Utilizing this pattern makes it easier to expand the relationship list in the future
var relationships = []string{
	relationshipUsers,
}

// New will return a new instance of the Controller
func New(dir string) (cc *Controller, err error) {
	var c Controller
	// Make options with:
	//	- Database name
	//	- Data directory
	opts := mojura.MakeOpts("helloWorld", dir)

	// Initialize a new instance of mojura.Mojura, we pass it the following values:
	//	- Options
	//	- Appended list of relationship keys
	if c.c, err = mojura.New[*Entry](opts, relationships...); err != nil {
		return
	}

	// Assign pointer reference to our controller
	cc = &c
	return
}

// Controller represents a management layer to facilitate the retrieval and modification of Entries
type Controller struct {
	// Mojura will manage the data layer and will utilize the underlying back-end
	c *mojura.Mojura[*Entry]
}

// New will insert a new Entry to the back-end
// Note: ID, CreatedAt, UpdatedAt will all be auto-set by mojura.Mojura
func (c *Controller) New(e Entry) (created *Entry, err error) {
	// Attempt to validate Entry
	if err = e.Validate(); err != nil {
		// Entry is not valid, return validation error
		return
	}

	// Entry is valid!

	// Insert Entry into mojura.Mojura and return the results
	return c.c.New(&e)
}

// Get will retrieve an Entry which has the same ID as the provided entryID
func (c *Controller) Get(entryID string) (entry *Entry, err error) {
	// Attempt to get Entry with the provided ID, if no entry exists mojura.ErrEntryNotFound will be returned
	return c.c.Get(entryID)
}

// GetByUser will retrieve all Entries related to the provided userID by way of the Users relationship
func (c *Controller) GetByUser(userID string) (entries []*Entry, err error) {
	// Create user relationship match filter for provided user ID
	ownedByUser := filters.Match(relationshipUsers, userID)
	// Initialize opts with the user filter
	opts := mojura.NewFilteringOpts(ownedByUser)

	// Get all matching entries
	if entries, _, err = c.c.GetFiltered(opts); err != nil {
		return
	}

	// No matches will still yield no error
	return
}

// ForEach without any filters will iterate through all Entries
// Note: The error constant mojura.Break can returned by the iterating func to end the iteration early
func (c *Controller) ForEach(fn func(*Entry) error, filters ...mojura.Filter) (err error) {
	// Initialize opts with the provided filters
	opts := mojura.NewFilteringOpts(filters...)
	// Iterate through all entries
	err = c.c.ForEach(func(key string, e *Entry) (err error) {
		// Pass iterating Entry to iterating function
		return fn(e)
	}, opts)
	return
}

// ForEachByUser will iterate through all Entries for a provided userID
// Note: The error constant mojura.Break can returned by the iterating func to end the iteration early
func (c *Controller) ForEachByUser(userID string, fn func(*Entry) error) (err error) {
	// Create user relationship match filter for provided user ID
	ownedByUser := filters.Match(relationshipUsers, userID)
	// Call ForEach with the ownedByUser filter
	return c.ForEach(fn, ownedByUser)
}

// Update will update the Entry for a given entryID
func (c *Controller) Update(entryID string, e Entry) (updated *Entry, err error) {
	// Attempt to validate Entry
	if err = e.Validate(); err != nil {
		// Entry is not valid, return validation error
		return
	}

	// Insert Entry into mojura.Mojura and return the results
	return c.c.Update(entryID, func(orig *Entry) (err error) {
		// It's safer to manually update each field, instead of replacing the entire entry.
		// This helps to ensure that we don't trust the client fully and allow only specific
		// fields to be updated
		orig.FavoriteTimeOfDay = e.FavoriteTimeOfDay
		orig.Greeting = e.Greeting
		return
	})
}

// Delete will remove an Entry for a given entryID
func (c *Controller) Delete(entryID string) (deleted *Entry, err error) {
	// Remove Entry from mojura.Mojura
	return c.c.Delete(entryID)
}

// Close will close the controller and it's underlying dependencies
func (c *Controller) Close() (err error) {
	// Close database
	return c.c.Close()
}
