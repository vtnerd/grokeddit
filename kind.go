package grokeddit

type KindType uint8

const (
	Comment   KindType = iota // Indicates Reddit comment
	Link      KindType = iota // Indicates Reddit link
	Subreddit KindType = iota // Indicates Reddit Subreddit
)
