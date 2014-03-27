package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"errors"
)

type forwardThingIterate struct {
	things []grokeddit.Thing
}

func (iterater *forwardThingIterate) setArray(newThings []grokeddit.Thing) {
	iterater.things = newThings
}

func (iterater *forwardThingIterate) hasNext() bool {
	return iterater.things != nil && len(iterater.things) != 0
}

func (iterater *forwardThingIterate) getNext() (grokeddit.Thing, error) {
	if !iterater.hasNext() {
		return grokeddit.Thing{}, errors.New("No more things in current array")
	}

	nextThing := iterater.things[0]
	iterater.things = iterater.things[1:len(iterater.things)]
	return nextThing, nil
}
