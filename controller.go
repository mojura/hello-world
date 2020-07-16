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
