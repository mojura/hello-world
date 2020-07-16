package helloworld

import "github.com/gdbu/dbl"

// New will return a new instance of the Controller
func New(dir string) (cc *Controller, err error) {
	var c Controller
	// Initialize a new instance of dbl.Core, we pass it the following values:
	//	- Name of the controller (Must be unique, used to create the database files)
	//	- Directory location to persist data
	//	- Example Entry value (Used to populate reflected values)
	if c.c, err = dbl.New("helloWorld", dir, &Entry{}); err != nil {
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

// Close will close the controller and it's underlying dependencies
func (c *Controller) Close() (err error) {
	// Since we only have one dependency, we can just call this func directly
	return c.c.Close()
}
