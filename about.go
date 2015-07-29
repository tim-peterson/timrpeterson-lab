package guestbook

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type aboutInfo struct {
	Posts int
}

func aboutPage(r *http.Request) interface{} {
	c := appengine.NewContext(r)
	count, err := datastore.NewQuery("GuestbookGreeting").Ancestor(guestbookKey(c)).Count(c)
	if err != nil {
		return nil
	}

	info := &aboutInfo{
		Posts: count,
	}
	return info
}
