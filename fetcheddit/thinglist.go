package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"errors"
)

type thingIterater interface {
	setArray([]grokeddit.Thing)
	hasNext() bool
	getNext() (grokeddit.Thing, error)
}

type thingList struct {
	fetchMoreThings func() ([]grokeddit.Thing, bool, error)
	currentThings   thingIterater // An abstraction for iterating over the current slice of Thing objects. Implementations can change the order of iterating.
	moreThings      bool          // The last value returned by fetchMoreThings
}

func fetchThingList(fetchGroked func(anchor *AnchorPoint) (grokeddit.Groked, error), anchor *AnchorPoint) (thingList, error) {
	var currentThingIterater thingIterater
	var fetchMore func() ([]grokeddit.Thing, bool, error)

	// make a copy of the anchor, since it will be stored in a
	// closure thats invoked asynchronously
	if anchor != nil {
		newAnchor := new(AnchorPoint)
		*newAnchor = *anchor
		anchor = newAnchor
	}

	if anchor == nil || anchor.Direction == Next {
		currentThingIterater = &forwardThingIterate{}
		fetchMore = func() ([]grokeddit.Thing, bool, error) {
			newGroked, error := fetchGroked(anchor)
			if error != nil {
				return nil, false, error
			}

			if anchor == nil {
				anchor = new(AnchorPoint)
				anchor.Direction = Next
			}

			if newGroked.ListingNext != nil {
				anchor.Anchor = *newGroked.ListingNext
			}

			return newGroked.Children, newGroked.ListingNext != nil, nil
		}
	} else {
		currentThingIterater = &reverseThingIterate{}
		fetchMore = func() ([]grokeddit.Thing, bool, error) {
			newGroked, error := fetchGroked(anchor)
			if error != nil {
				return nil, false, error
			}

			// There is a bug on the reddit side, the "before" 
			// field is always NULL with a before request. This is 
			// solved with a hack that will do 1 extra request
			// than necessary, the additional traffic is on you 
			// reddit.
			more := false
			if newGroked.Children != nil && len(newGroked.Children) > 0 {
				anchor.Anchor = newGroked.Children[0].Id
				more = true
			}

			return newGroked.Children, more, nil
		}
	}

	firstThings, moreThings, error := fetchMore()
	if error != nil {
		return thingList{}, errors.New("No things in first request")
	}

	currentThingIterater.setArray(firstThings)
	return thingList{fetchMore, currentThingIterater, moreThings}, nil
}

func (list *thingList) hasNext() bool {
	return list.moreThings || list.currentThings.hasNext()
}

func (list *thingList) getNext() (grokeddit.Thing, error) {
	if !list.hasNext() {
		return grokeddit.Thing{}, errors.New("No more things to iterate")
	}

	if !list.currentThings.hasNext() {
		nextThings, moreThings, error := list.fetchMoreThings()

		if error != nil {
			return grokeddit.Thing{}, errors.New("Unable to fetch more things: " + error.Error())
		}

		list.moreThings = moreThings
		list.currentThings.setArray(nextThings)
	}

	return list.currentThings.getNext()
}