package helloworld

import "github.com/gdbu/dbl"

// Controller represents a management layer to facilitate the retrieval and modification of Entries
type Controller struct {
	// Core will manage the data layer and will utilize the underlying back-end
	c *dbl.Core
}
