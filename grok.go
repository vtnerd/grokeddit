package grokeddit

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
)

type Groked struct {
	ListingPrev string  // Indicates thing before this listing (if listing was groked)
	ListingNext string  // Indicates thing after this listing (if listing was groked)
	Children    []Thing // Children things to the listing, or the single thing groked
}

type internalGroked struct {
	Kind string
	Data json.RawMessage
}

type listing struct {
	Children []internalGroked // the children of the listing
	Before   string           // indicates value for prev filter
	After    string           // indicates value for next filter
}

type internalThing struct {
	Author        string          // name of the poster
	Body_html     string          // html from a comment
	Created_utc   float64         // utc of creation time
	Edited        interface{}     // false or utc of last edit
	Id            string          // unique indentifier for the thing
	lastUpdate    float64         // after parsing, this will be the last modification time (usually created_utc)
	Parent_id     string          // parent id of a comment
	Replies       json.RawMessage // unparsed replies to a comment
	Selftext_html string          // html from a new post
	Subreddit     string          // name of the subreddit associated with the thing
	Subreddit_id  string          // id of the subreddit associated with the thing
	Title         string          // title of the post
	Url           string          // url of the post
}

func addNewThing(in internalGroked, out *Groked) error {
	// 
	// Normalize Kind. Do before parsing "data", in case
	// we cannot handle the kind
	//
	currentKind, error := ParseKind(in.Kind)
	if error != nil {
		return errors.New("Unable to grok: " + error.Error())
	}

	// 
	// Parse data portion of thing
	//
	var parsedIn internalThing
	if error := json.Unmarshal(in.Data, &parsedIn); error != nil {
		return errors.New("Unable to grok: Unable to decode thing of type \"" + in.Kind + "\"")
	}

	//
	// Normalize body text
	//
	var bodyHtml string
	// sucks that the switch happens twice, but I'd rather do
	// the unmarshal after we know it can be properly decoded
	switch currentKind {
	case Comment:
		bodyHtml = parsedIn.Body_html
	case Link:
		bodyHtml = parsedIn.Selftext_html
	default:
	}

	//
	// Normalize creation and last modification timestamp
	//
	creationTime := int64(parsedIn.Created_utc)
	lastModificationTime := creationTime
	if parsedIn.Edited != nil {

		editValue := reflect.ValueOf(parsedIn.Edited)

		if editValue.Kind() == reflect.Float64 {
			lastModificationTime = int64(editValue.Float())
		} else if editValue.Kind() != reflect.Bool {
			return errors.New("Unable to grok: Unexpected type \"" + editValue.Type().String() + "\" for edited field")
		}
	}

	//
	// Normalize Id fields
	//
	thingId, error := ParseId(parsedIn.Id) // thing id should always be present
	if error != nil {
		return errors.New("Unable to grok: " + error.Error())
	}

	subredditId := GlobalId{thingId, Subreddit}
	parentId := GlobalId{}
	if currentKind != Subreddit {
		subredditId.Id, error = ParseId(parsedIn.Subreddit_id)
		if error != nil {
			return errors.New("Unable to grok subreddit id: " + error.Error())
		}

		parentId, error = ParseGlobalId(parsedIn.Parent_id)
		if error != nil {
			return errors.New("Unable to grok parent id: " + error.Error())
		}
	}
	

	out.Children = append(out.Children, Thing{parsedIn.Author, creationTime, GlobalId{thingId, currentKind}, lastModificationTime, parentId, parsedIn.Replies, parsedIn.Subreddit, subredditId, currentKind, bodyHtml, parsedIn.Title, parsedIn.Url})
	return nil
}

func internalGrok(parsedObject internalGroked) (*Groked, error) {
	groked := Groked{}

	if parsedObject.Kind == "Listing" {
		var parsedListing listing
		if error := json.Unmarshal(parsedObject.Data, &parsedListing); error != nil {
			return nil, errors.New("Unable to grok listing: " + error.Error())
		}

		groked.ListingPrev = parsedListing.Before
		groked.ListingNext = parsedListing.After
		groked.Children = make([]Thing, 0, len(parsedListing.Children))

		for _, element := range parsedListing.Children {

			// There should be only "Things" in a child listing. So
			// this will return an error if a non-thing is found.
			if error := addNewThing(element, &groked); error != nil {
				return nil, error
			}
		}
	} else {
		groked.Children = make([]Thing, 0, 1)
		if error := addNewThing(parsedObject, &groked); error != nil {
			return nil, error
		}
	}

	return &groked, nil
}

func GrokObject(dataSource io.Reader) (*Groked, error) {
	var parsedObject internalGroked
	if error := json.NewDecoder(dataSource).Decode(&parsedObject); error != nil && error != io.EOF {
		return nil, errors.New("Unable to grok reddit JSON object: " + error.Error())
	}

	return internalGrok(parsedObject)
}

func GrokArray(dataSource io.Reader) ([]Groked, error) {
	var parsedArray []internalGroked

	if error := json.NewDecoder(dataSource).Decode(&parsedArray); error != nil && error != io.EOF {
		return nil, errors.New("Unable to grok reddit JSON array: " + error.Error())
	}

	grokList := make([]Groked, 0, len(parsedArray))

	for _, parsedObject := range parsedArray {
		newGrok, error := internalGrok(parsedObject)

		if error != nil {
			return nil, error
		}

		grokList = append(grokList, *newGrok)
	}

	return grokList, nil
}
