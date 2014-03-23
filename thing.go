package grokeddit

/* Represents any type of thing in reddit. Some fields may be 
empty/invalid for the particular type. */
type Thing struct {
	Author      string   // name of the poster
	Created_utc int64    // utc of creation time
	Id          GlobalId // uniquely indentifies the thing
	LastUpdate  int64    // utc of last update
	ParentId    GlobalId // parent id of a comment
	Replies     Groked   // replies to a comment
	Subreddit   string   // name of the subreddit associated with the thing
	SubredditId GlobalId // uniquely identifies the subreddit associated with the thing
	Text_html   string   // html from post
	Title       string   // title of the post
	Url         string   // url of the post
}
