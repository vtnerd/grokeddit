package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"errors"
)

type reverseThingIterate struct {
	things    []grokeddit.Thing
	lastIndex int
}

func (iterater *reverseThingIterate) setArray(newThings []grokeddit.Thing) {
	iterater.lastIndex = 0
	iterater.things = newThings
	if newThings != nil {
		iterater.lastIndex = len(newThings)
	}
}

func (iterater *reverseThingIterate) hasNext() bool {
	return iterater.things != nil && iterater.lastIndex != 0
}

func (iterater *reverseThingIterate) getNext() (grokeddit.Thing, error) {
	if !iterater.hasNext() {
		return grokeddit.Thing{}, errors.New("No more things in current array")
	}

	iterater.lastIndex = iterater.lastIndex - 1
	nextThing := iterater.things[iterater.lastIndex]
	return nextThing, nil
}
