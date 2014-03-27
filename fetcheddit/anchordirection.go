package fetcheddit

/* Represents an anchor direction. Indicates whether next or previous listings
need to be fetched. */
type AnchorDirection uint8

const (
	Previous AnchorDirection = iota // Indicates reddit before
	Next     AnchorDirection = iota // Indicates reddit after
)

var directionConversion = [...]string{
	"before",
	"after",
}

/* Convert AnchorDirection enumerated kind type to its string representation. If
the value is not one defined by a constant, then this could result in a runtime
error (array out of bounds access). */
func (direction AnchorDirection) String() string {
	/* There isn't a way to force the values of a KindType to be from 0..1
	 without a check. A check would imply a return value of
	 (string, error), which I would like to avoid for simplicity. If
	 someone does AnchorDirection(100).String(), let the runtime crap it
	 out (terrible I know, but this is defined behavior in Go,
	unlike C/C++). */
	return directionConversion[uint8(direction)]
}
