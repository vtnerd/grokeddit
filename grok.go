package grokeddit

import (
	"encoding/json"
	"errors"
	"io"
)

type Groked struct {
	ListingPrev string
	ListingNext string
	Children    []Thing
}

type internalGroked struct {
	Kind string
	Data json.RawMessage
}

type listing struct {
	Children []internalGroked
	Before   string
	After    string
}

func addNewThing(in *internalGroked, out *Groked) error {

	switch in.Kind {
	case "t1":
		out.Children = append(out.Children, Thing{CommentType, in.Data})

	case "t3":
		out.Children = append(out.Children, Thing{LinkType, in.Data})

	case "t5":
		out.Children = append(out.Children, Thing{SubredditType, in.Data})

	default:
		return errors.New("Unable to grok: Unsupported kind \"" + in.Kind + "\"")
	}

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
			if error := addNewThing(&element, &groked); error != nil {
				return nil, error
			}
		}
	} else {
		groked.Children = make([]Thing, 0, 1)
		if error := addNewThing(&parsedObject, &groked); error != nil {
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
