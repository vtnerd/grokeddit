package grokeddit

import "encoding/json"

type Thing struct {
	Author      string          // name of the poster
	Created_utc int64           // utc of creation time
	Id          GlobalId        // uniquely indentifies the thing
	LastUpdate  int64           // utc of last update
	ParentId    GlobalId        // parent id of a comment
	replies     json.RawMessage // unparsed replies to a comment
	Subreddit   string          // name of the subreddit associated with the thing
	SubredditId GlobalId        // uniquely identifies the subreddit associated with the thing
	Kind        KindType        // Indicates the type of thing
	Text_html   string          // html from post
	Title       string          // title of the post
	Url         string          // url of the post
}
