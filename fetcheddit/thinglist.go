package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"errors"
)

type nextChunk struct {
	nextThings     []grokeddit.Thing
	moreThings     bool
	retrievalError error
}

type thingList struct {
	fetchMoreThings func() ([]grokeddit.Thing, bool, error)
	currentThings   thingIterater // Dictates how a slice returned by fetchMoreThings is iterated
	moreThings      bool          // The last value returned by fetchMoreThings
	chunkChannel    chan nextChunk
}

func fetchThingList(
	fetchGroked func(Fetcher, *AnchorPoint) (grokeddit.Groked, error),
	contentFetcher Fetcher,
	anchor *AnchorPoint) (thingList, error) {

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
			newGroked, error := fetchGroked(contentFetcher, anchor)
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
			newGroked, error := fetchGroked(contentFetcher, anchor)
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
	newThingList := thingList{fetchMore, currentThingIterater, moreThings, make(chan nextChunk)}

	if moreThings {
		newThingList.fetchNextBlockAsync()
	}
	return newThingList, nil
}

func (list *thingList) fetchNextBlockAsync() {
	go func() {
		nextThings, moreThings, error := list.fetchMoreThings()
		list.chunkChannel <- nextChunk{nextThings, moreThings, error}
	}()
}

func (list *thingList) hasNext() bool {
	return list.moreThings || list.currentThings.hasNext()
}

func (list *thingList) getNext() (grokeddit.Thing, error) {
	if !list.hasNext() {
		return grokeddit.Thing{}, errors.New("No more things to iterate")
	}

	if !list.currentThings.hasNext() {
		nextChunk := <-list.chunkChannel

		if nextChunk.moreThings {
			list.fetchNextBlockAsync()
		}

		if nextChunk.retrievalError != nil {
			return grokeddit.Thing{}, errors.New("Unable to fetch more things: " + nextChunk.retrievalError.Error())
		}

		list.moreThings = nextChunk.moreThings
		list.currentThings.setArray(nextChunk.nextThings)
	}

	return list.currentThings.getNext()
}
