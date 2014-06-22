package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"errors"
)

// Aggregate returned in the channel that fetches reddit thing lists
type nextChunk struct {
	nextThings     []grokeddit.Thing // The list of things retrieved
	moreThings     bool              // True if there is another list of things
	retrievalError error             // Non-nil if there was an error trying to retrieve the list of things
}

/* An abstraction for a list of reddit things. The paginated lists are
"flattened" to appear as giant single list. The pages are retreived
asynchronously, so there is less blocking time waiting for the next page. */
type thingList struct {
	fetchMoreThings func() ([]grokeddit.Thing, bool, error) // Function that can be called in a separate goroutine to retrieve the next page of things
	currentThings   thingIterater                           // Dictates how a slice returned by fetchMoreThings is iterated
	moreThings      bool                                    // The last value returned by fetchMoreThings
	chunkChannel    chan nextChunk                          // Receives the next page of reddit things
}

/* Creates a new thing list.

The closure arugment (fetchedGroked) allows the caller to handle differences in
the reddit API. The lists can be sent in two different ways depending on the
context, and the caller must massage those differences into a single Groked
object.

The contentFetcher allows the caller to modify how the reddit things are being
retrieved, mainly used for testing (possibly different URLs too).

The anchor specifies a starting point, and an iteration direction. NULL
indicates to start from the beginning, and iterate forward. */
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

			if newGroked.ListingPrev != nil {
				anchor.Anchor = *newGroked.ListingPrev
			}

			return newGroked.Children, newGroked.ListingPrev != nil, nil
		}
	}

	firstThings, moreThings, error := fetchMore()
	if error != nil {
		return thingList{}, errors.New("No things in first request")
	}

	currentThingIterater.setArray(firstThings)
	newThingList := thingList{
		fetchMore,
		currentThingIterater,
		moreThings,
		make(chan nextChunk),
	}

	if moreThings {
		newThingList.fetchNextBlockAsync()
	}
	return newThingList, nil
}

// Asynchronously fetches the next page of reddit things.
func (list *thingList) fetchNextBlockAsync() {
	go func() {
		nextThings, moreThings, error := list.fetchMoreThings()
		list.chunkChannel <- nextChunk{nextThings, moreThings, error}
	}()
}

// Return true if the thingList has another thing available.
func (list *thingList) hasNext() bool {
	return list.moreThings || list.currentThings.hasNext()
}

// Return the next thing in the list, or error.
func (list *thingList) getNext() (grokeddit.Thing, error) {
	if !list.hasNext() {
		return grokeddit.Thing{}, errors.New("No more things to iterate")
	}

	if !list.currentThings.hasNext() {
		nextChunk := <-list.chunkChannel

		// set this variable now, a channel check can only be
		// done if async request is active.
		list.moreThings = nextChunk.moreThings
		if nextChunk.moreThings {
			list.fetchNextBlockAsync()
		}

		if nextChunk.retrievalError != nil {
			return grokeddit.Thing{}, errors.New("Unable to fetch more things: " + nextChunk.retrievalError.Error())
		}

		list.currentThings.setArray(nextChunk.nextThings)
	}

	return list.currentThings.getNext()
}
