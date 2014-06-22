package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
)

/* Abstraction for specifying the next listing desired. Specifies an 
anchor and direction */
type AnchorPoint struct {
	Anchor    grokeddit.GlobalId // Specifies the next/prev item in the anchor
	Direction AnchorDirection    // Specifies the anchor direction (next/prev)
}

/* Return the anchor point as a string. Example: "before=t3_blah" */
func (anchor AnchorPoint) String() string {
	return anchor.Direction.String() + "=" + anchor.Anchor.String()
}
