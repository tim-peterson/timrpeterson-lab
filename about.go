package guestbook

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type aboutInfo struct {
	Posts int
}

func about(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	count, err := datastore.NewQuery("GuestbookGreeting").Ancestor(guestbookKey(c)).Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	info := &aboutInfo{
		Posts: count,
	}
	if err := aboutTemplate.Execute(w, info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
