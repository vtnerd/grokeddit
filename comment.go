package grokeddit

import "encoding/json"

type Comment struct {
	BaseThing
	ContentThing
	ParentId string
	replies  json.RawMessage
}
