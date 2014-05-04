package fetcheddit

import "code.leeclagett.com/grokeddit"

// An interface for iterating over a slice of grokeddit.Thing objects.
type thingIterater interface {
	setArray([]grokeddit.Thing)        // Set the slice to iterate
	hasNext() bool                     // Return true if there is another Thing to return
	getNext() (grokeddit.Thing, error) // Get the next thing
}
