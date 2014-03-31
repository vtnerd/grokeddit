package grokeddit

import (
	"errors"
	"strings"
)

/* Uniquely identifies any "thing" in reddit. Things processed by grokeddit are:
subreddits, links, and comments. */
type GlobalId struct {
	Id   ThingId  // unique identifier
	Kind KindType // the type of thing
}

/* Parses the global id format, as used by reddit. The format is "type_id". 
Type must be t1, t2, t3, t4, or t5. id must be a base36 number. Error is 
returned if the format is invalid. */
func ParseGlobalId(globalId string) (GlobalId, error) {

	splitId := strings.Split(globalId, "_")
	if len(splitId) != 2 {
		return GlobalId{}, errors.New("Invalid global id \"" + globalId + "\"")
	}

	kind, error := ParseKind(splitId[0])
	if error != nil {
		return GlobalId{}, errors.New("Invalid global id : " + error.Error())
	}

	id, error := ParseId(splitId[1])
	if error != nil {
		return GlobalId{}, errors.New("Invalid global id : " + error.Error())
	}

	return GlobalId{id, kind}, nil
}

// Returns a string representation for the global ID, as used by reddit.
func (globalId GlobalId) String() string {
	return globalId.Kind.String() + "_" + globalId.Id.String()
}
