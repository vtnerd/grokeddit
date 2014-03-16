package grokeddit

import (
	"encoding/json"
	"errors"
	"reflect"
)

type Thing struct {
	Type ThingType
	data json.RawMessage
}

type internalThing struct {
	Author        string          // name of the poster
	Body_html     string          // html from a comment
	Created_utc   int64           // utc of
	Edited        interface{}     // false or utc of last edit
	Id            string          // unique indentifier for the thing
	lastUpdate    int64           // after parsing, this will be the last modification time (usually created_utc)
	Parent_id     string          // parent id of a comment
	Replies       json.RawMessage // unparsed replies to a comment
	Selftext_html string          // html from a new post
	Subreddit     string          // name of the subreddit associated with the thing
	Subreddit_id  string          // id of the subreddit associated with the thing
	Title         string          // title of the post
	Url           string          // url of the post
}

func parseInternal(data json.RawMessage) (*internalThing, error) {
	var parsedData internalThing
	if error := json.Unmarshal(data, &parsedData); error != nil {
		return nil, errors.New("Unable to decode comment: " + error.Error())
	}

	// Default last update time to creation time
	parsedData.lastUpdate = parsedData.Created_utc

	if parsedData.Edited != nil {

		editValue := reflect.ValueOf(parsedData.Edited)

		if editValue.Kind() == reflect.Int64 {
			parsedData.lastUpdate = editValue.Int()
		} else {
			return nil, errors.New("Unexpected type \"" + editValue.Type().String() + "\" for edited field")
		}
	}

	return &parsedData, nil
}

func (thing *Thing) AsComment() (*Comment, error) {
	if thing.Type != CommentType {
		return nil, errors.New("Expected comment type")
	}

	parsed, error := parseInternal(thing.data)
	if error != nil {
		return nil, error
	}

	return &Comment{BaseThing{parsed.Id, parsed.Created_utc, parsed.Subreddit, parsed.Subreddit_id}, ContentThing{parsed.Author, parsed.Body_html, parsed.lastUpdate}, parsed.Parent_id, parsed.Replies}, nil
}

func (thing *Thing) AsLink() (*Link, error) {
	if thing.Type != LinkType {
		return nil, errors.New("Expected link type")
	}

	parsed, error := parseInternal(thing.data)
	if error != nil {
		return nil, error
	}

	return &Link{BaseThing{parsed.Id, parsed.Created_utc, parsed.Subreddit, parsed.Subreddit_id}, ContentThing{parsed.Author, parsed.Body_html, parsed.lastUpdate}, parsed.Title, parsed.Url}, nil
}

func (thing *Thing) AsSubreddit() (*Subreddit, error) {
	if thing.Type != SubredditType {
		return nil, errors.New("Expected subreddit type")
	}

	parsed, error := parseInternal(thing.data)
	if error != nil {
		return nil, error
	}

	return &Subreddit{BaseThing{parsed.Id, parsed.Created_utc, parsed.Subreddit, parsed.Subreddit_id}}, nil
}
