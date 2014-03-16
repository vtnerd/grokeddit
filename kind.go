package grokeddit

type ThingType uint8

const (
	CommentType   ThingType = iota
	LinkType      ThingType = iota
	SubredditType ThingType = iota
)
