package fetcheddit

import "errors"

type Links struct {
	things thingList
}

func fetchLinks(things thingList) Links {
	return Links{things}
}

func (links *Links) HasNext() bool {
	return links.things.hasNext()
}

func (links *Links) GetNext() (Link, error) {

	nextThing, error := links.things.getNext()
	if error != nil {
		return Link{}, errors.New("Unable to retrieve next link: " + error.Error())
	}

	return Link{nextThing.Author, nextThing.CreatedUtc, nextThing.Id, nextThing.LastUpdateUtc, nextThing.Replies, nextThing.Subreddit, nextThing.SubredditId, nextThing.Text_html, nextThing.Title, nextThing.Url}, nil
}
