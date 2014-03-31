package fetcheddit

import "code.leeclagett.com/grokeddit"

type thingIterater interface {
	setArray([]grokeddit.Thing)
	hasNext() bool
	getNext() (grokeddit.Thing, error)
}
