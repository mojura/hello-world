package helloworld

import (
	"fmt"

	"github.com/gdbu/dbl"
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
	// Initialize a new instance of dbl.Core, we pass it the following values:
	//	- Name of the controller (Must be unique, used to create the database files)
	//	- Directory location to persist data
	//	- Example Entry value (Used to populate reflected values)
	//	- Appended list of relationship keys
	if c.c, err = dbl.New("helloWorld", dir, &Entry{}, relationships...); err != nil {
		return
	}

	// Assign pointer reference to our controller
	cc = &c
	return
}

// Controller represents a management layer to facilitate the retrieval and modification of Entries
type Controller struct {
	// Core will manage the data layer and will utilize the underlying back-end
	c *dbl.Core
}

// New will insert a new Entry to the back-end
// Note: ID, CreatedAt, UpdatedAt will all be auto-set by dbl.Core
func (c *Controller) New(e Entry) (createdID string, err error) {
	// Attempt to validate Entry
	if err = e.Validate(); err != nil {
		// Entry is not valid, return validation error
		return
	}

	// Entry is valid!

	// Insert Entry into dbl.Core and return the results
	return c.c.New(&e)
}

// Get will retrieve an Entry which has the same ID as the provided entryID
func (c *Controller) Get(entryID string) (entry *Entry, err error) {
	var e Entry
	// Attempt to get Entry with the provided ID, pass reference to entry for which values to be applied
	if err = c.c.Get(entryID, &e); err != nil {
		// No entry with the provided ID was found, return error
		return
	}

	// Assign reference to retrieved Entry
	entry = &e
	return
}

// GetByUser will retrieve all Entries related to the provided userID by way of the Users relationship
func (c *Controller) GetByUser(userID string) (entries []*Entry, err error) {
	var e Entry
	// Attempt to get the Entries related to the userID, passing the:
	//	- Relationship type (Users)
	//	- Relationship ID (userID)
	//	- Entries slice to be appended to
	if err = c.c.GetByRelationship(relationshipUsers, userID, &e); err != nil {
		// No entry with the provided ID was found, return error
		// Note: Utilizing this err/return pattern will yield in less mistakes if/when logic is expanded below
		return
	}

	return
}

// ForEach will iterate through all Entries
// Note: The error constant dbl.Break can returned by the iterating func to end the iteration early
func (c *Controller) ForEach(fn func(*Entry) error) (err error) {
	// Iterate through all entries
	c.c.ForEach(func(key string, val dbl.Value) (err error) {
		// Attempt to assert the value as an *Entry
		e, ok := val.(*Entry)
		// Ensure assertion was successful
		if !ok {
			// Invalid type provided, return error
			err = fmt.Errorf("invalid entry type, expected %T and received %T", e, val)
			return
		}

		// Pass iterating Entry to iterating function
		return fn(e)
	})

	return
}

// Update will update the Entry for a given entryID
func (c *Controller) Update(entryID string, e Entry) (err error) {
	// Attempt to validate Entry
	if err = e.Validate(); err != nil {
		// Entry is not valid, return validation error
		return
	}

	// Entry is valid!

	// Insert Entry into dbl.Core and return the results
	return c.c.Edit(entryID, &e)
}

// Delete will remove an Entry for a given entryID
func (c *Controller) Delete(entryID string) (err error) {
	// Remove Entry from dbl.Core
	return c.c.Remove(entryID)
}

// Close will close the controller and it's underlying dependencies
func (c *Controller) Close() (err error) {
	// Since we only have one dependency, we can just call this func directly
	return c.c.Close()
}
