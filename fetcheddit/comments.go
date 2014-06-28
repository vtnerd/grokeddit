package fetcheddit

import "errors"

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
