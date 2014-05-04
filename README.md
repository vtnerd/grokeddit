grokeddit
=========

Library for interacting with Reddit API in Go! The library provides an easy interface for retrieving links and comments from Reddit. Example:
```go
subredditsHandle, error := fetcheddit.FetchSubreddits({"music", "movies"}, fetcheddit.DefaultFetch)
if error != nil {
  return error
}

links, error := subredditsHandle.FetchNewLinks(nil)
if error != nil {
  return error
}
	
for links.HasNext() {
  link, error := links.GetNext()
  if error != nil {
    return error
  }
	
  log.Printf("Link Title: %s", link.Title)
} 
```
The library automatically fetches a block of elements on a separate goroutine. This allows for a range style interface (the user is not exposed to the individual blocks of elements retrieved), and interleaving of processing and retrieval (elements are being traversed while the OS is waiting for Reddit to respond with the next block of elements).

_This library is a work-in-progress and should be used with caution. The API is unlikely to change, but more tests and features are likely needed to make it usuable for your project. The library does not support writing to Reddit, which may be implemented at a later time._

Design
------
The library is split into two parts:

1. The **grokeddit** package parses the JSON and massages the data from the server to more relevant types (less strings).
2. The **fetcheddit** package intelligently retrieves JSON documents, passes the information to grokeddit, then strips unused fields (based on context). Most users will want to use this package, instead of directly using grokeddit.
