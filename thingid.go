package grokeddit

import (
	"errors"
	"strconv"
)

type ThingId uint64

func ParseId(id string) (ThingId, error) {
	idConverted, error := strconv.ParseUint(id, 36, 64)
	if error != nil {
		return ThingId(0), errors.New("Unable to parse id: " + error.Error())
	}

	return ThingId(idConverted), nil
}

func (id ThingId) String() string {
	return strconv.FormatUint(uint64(id), 36)
}
