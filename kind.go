package grokeddit

import 	"errors"

/* Elements in Reddit are broken into different types whicht they call "kind".
This type is an enumerated type representing each "kind" that can be handled
by grokkeddit. */
type KindType uint8

const (
	Comment   KindType = iota // Indicates Reddit comment
	Link      KindType = iota // Indicates Reddit link
	Subreddit KindType = iota // Indicates Reddit Subreddit
)

var kindConversion = [...]string{
	"t1",
	"t3",
	"t5",
}

/* Parse a value for "kind" as received by reddit, and return the enumerated
type used by grokeddit for that kind. If the string is not a valid kind (or one
not currently handled by grokeddit), an error is returned. */
func ParseKind(kind string) (KindType, error) {
	for index, stringName := range kindConversion {
		if kind == stringName {
			return KindType(index), nil
		}
	}

	return KindType(0), errors.New("Invalid kind \"" + kind + "\"")
}

/* Convert grokeddits enumerated kind type to its string representation. If the
value is not one defined by a constant, or returned by ParseKind, then this 
could result in a runtime error (array out of bounds access). */
func (kind KindType) String() string {
	/* There isn't a way to force the values of a KindType to be from 0..2
	 without a check. A check would imply a return value of
	 (string, error), which I would like to avoid for simplicity. If
	 someone does KindType(100).String(), let the runtime crap it
	 out (terrible I know, but this is defined behavior in Go,
	unlike C/C++). */
	return kindConversion[uint8(kind)]
}
