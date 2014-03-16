package grokeddit

import (
	"encoding/json"
)

type Thing struct {
	Type ThingType
	data json.RawMessage
}
