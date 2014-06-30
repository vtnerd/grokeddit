package fetcheddit

import (
	"errors"
	"code.leeclagett.com/grokeddit"
)

type Comments struct {
	things thingList
}

func fetchComments(things thingList) Comments {
	return Comments{things}
}

func (links *Comments) HasNext() bool {
	return links.things.hasNext()
}

func (comments *Comments) GetNext() (Comment, error) {

	nextThing, error := comments.things.getNext()
	if error != nil {
		return Comment{}, errors.New("Unable to retrieve next link: " + error.Error())
	}

	if nextThing.Id.Kind != grokeddit.Comment {
		return Comment{}, errors.New(
			"Expected type \"Comment\" but got \"" + nextThing.Id.Kind.String() + "\"")
	}

	return Comment{
		nextThing.Author, 
		nextThing.CreatedUtc, 
		nextThing.Id, 
		nextThing.LastUpdateUtc, 
		nextThing.ParentId, 
		nextThing.Replies, 
		nextThing.Subreddit, 
		nextThing.SubredditId, 
		nextThing.Text_html,
	}, nil
}
