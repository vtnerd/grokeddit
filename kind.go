package grokeddit

import (
	"errors"
)

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

func ParseKind(kind string) (KindType, error) {
	for index, stringName := range kindConversion {
		if kind == stringName {
			return KindType(index), nil
		}
	}

	return KindType(0), errors.New("Invalid kubd \"" + kind + "\"")
}

func (kind KindType) String() string {
	return kindConversion[uint8(kind)]
}
