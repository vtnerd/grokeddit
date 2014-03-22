package grokeddit

import (
	"errors"
	"strconv"
)

/* Identifier of a thing. The KindType + ThingId 
uniquely identifies the element (see GlobalId). */
type ThingId uint64

/* Parse the thing id as sent by the reddit api. An error will be returned if 
the string is not in this format (base36 number). */
func ParseId(id string) (ThingId, error) {
	idConverted, error := strconv.ParseUint(id, 36, 64)
	if error != nil {
		return ThingId(0), errors.New("Unable to parse id: " + error.Error())
	}

	return ThingId(idConverted), nil
}

/* Return the thing id as a base36 ascii string. This is the format used by 
the reddit api. */
func (id ThingId) String() string {
	return strconv.FormatUint(uint64(id), 36)
}
