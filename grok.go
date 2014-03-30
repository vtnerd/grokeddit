package grokeddit

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"
)

type Groked struct {
	ListingPrev *GlobalId // Indicates thing before this listing (if listing was groked)
	ListingNext *GlobalId // Indicates thing after this listing (if listing was groked)
	Children    []Thing   // Children things to the listing, or the single thing groked
}

type thing struct {
	Kind string
	Data struct {
		Author        string          // name of the poster
		Body_html     string          // html from a comment
		Created_utc   float64         // utc of creation time
		Display_name  string          // Name of subreddit (only used with subreddit thing type)
		Edited        interface{}     // false or utc of last edit
		Id            string          // unique indentifier for the thing
		Parent_id     string          // parent id of a comment
		Replies       json.RawMessage // replies to a comment
		Selftext_html string          // html from a new post
		Subreddit     string          // name of the subreddit associated with the thing
		Subreddit_id  string          // id of the subreddit associated with the thing
		Title         string          // title of the post
		Url           string          // url of the post
	}
}

type listing struct {
	Data struct {
		Children []thing // the children of the listing
		Before   string  // indicates value for prev filter
		After    string  // indicates value for next filter
	}
}

func createNewThing(in thing) (Thing, error) {

	//
	// Normalize Kind. Do first, since nothing else will work with
	// unsupported kind being parsed.
	//
	currentKind, error := ParseKind(in.Kind)
	if error != nil {
		return Thing{}, errors.New("Unable to grok: " + error.Error())
	}

	//
	// Normalize body text
	//
	var bodyHtml string
	// sucks that the switch happens twice, but I'd rather do
	// the unmarshal after we know it can be properly decoded
	switch currentKind {
	case Comment:
		bodyHtml = in.Data.Body_html
	case Link:
		bodyHtml = in.Data.Selftext_html
	default:
	}

	//
	// Normalize creation and last modification timestamp
	//
	creationTime := int64(in.Data.Created_utc)
	lastModificationTime := creationTime
	if in.Data.Edited != nil {

		editValue := reflect.ValueOf(in.Data.Edited)

		if editValue.Kind() == reflect.Float64 {
			lastModificationTime = int64(editValue.Float())
		} else if editValue.Kind() != reflect.Bool {
			return Thing{}, errors.New("Unable to grok: Unexpected type \"" + editValue.Type().String() + "\" for edited field")
		}
	}

	//
	// Normalize Id fields
	//
	thingId, error := ParseId(in.Data.Id) // thing id should always be present
	if error != nil {
		return Thing{}, errors.New("Unable to grok: " + error.Error())
	}

	subredditName := in.Data.Subreddit
	subredditId := GlobalId{thingId, Subreddit}
	parentId := GlobalId{}
	if currentKind != Subreddit {
		subredditId, error = ParseGlobalId(in.Data.Subreddit_id)
		if error != nil {
			return Thing{}, errors.New("Unable to grok subreddit id: " + error.Error())
		}

		if len(in.Data.Parent_id) != 0 {

			parentId, error = ParseGlobalId(in.Data.Parent_id)
			if error != nil {
				return Thing{}, errors.New("Unable to grok parent id: " + error.Error())
			}
		}
	} else { // type subreddit
		subredditName = in.Data.Display_name
	}

	//
	// Cull through replies too
	//
	var newReplies Groked

	// If reply field was present
	if in.Data.Replies != nil && len(in.Data.Replies) > 0 {

		// if reply field wasn't empty string
		if in.Data.Replies[0] != '"' {
			// gross! indirect recursion! easiest way for now
			newReplies, error = GrokListing(bytes.NewReader(in.Data.Replies))

			if error != nil {
				return Thing{}, error
			}
		}
	}

	return Thing{in.Data.Author, creationTime, GlobalId{thingId, currentKind}, lastModificationTime, parentId, newReplies, subredditName, subredditId, bodyHtml, in.Data.Title, in.Data.Url}, nil
}

func internalGrok(parsedListing listing) (Groked, error) {
	groked := Groked{}

	groked.Children = make([]Thing, 0, len(parsedListing.Data.Children))

	if len(parsedListing.Data.Before) > 0 {
		previousListing, error := ParseGlobalId(parsedListing.Data.Before)
		if error != nil {
			return groked, errors.New("Invalid previous listing value: " + error.Error())
		}
		
		groked.ListingPrev = &previousListing
	}

	if len(parsedListing.Data.After) > 0 {
		nextListing, error := ParseGlobalId(parsedListing.Data.After)
		if error != nil {
			return groked, errors.New("Invalid next listing value: " + error.Error())
		}

		groked.ListingNext = &nextListing
	}

	for _, element := range parsedListing.Data.Children {

		thing, error := createNewThing(element)

		// There should be only "Things" in a child listing. So
		// this will return an error if a non-thing is found.
		if error != nil {
			return groked, error
		}

		groked.Children = append(groked.Children, thing)
	}

	return groked, nil
}

func GrokListing(dataSource io.Reader) (Groked, error) {
	var parsedData listing
	if error := json.NewDecoder(dataSource).Decode(&parsedData); error != nil && error != io.EOF {
		return Groked{}, errors.New("Unable to grok reddit JSON object: " + error.Error())
	}

	return internalGrok(parsedData)
}

func GrokListingArray(dataSource io.Reader) ([]Groked, error) {
	var grokList []Groked
	var parsedData []listing
	if error := json.NewDecoder(dataSource).Decode(&parsedData); error != nil && error != io.EOF {
		return grokList, errors.New("Unable to grok reddit JSON array: " + error.Error())
	}

	grokList = make([]Groked, len(parsedData))

	for index, parsedListing := range parsedData {
		newGrok, error := internalGrok(parsedListing)

		if error != nil {
			return grokList, error
		}

		grokList[index] = newGrok
	}

	return grokList, nil
}
