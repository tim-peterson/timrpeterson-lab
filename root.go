package guestbook

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func rootPage(r *http.Request) interface{} {
	c := appengine.NewContext(r)

	//http.FileServer(http.Dir("./public"))

	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	// [START query]
	q := datastore.NewQuery("GuestbookGreeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
	// [END query]

	// [START getall]
	greetings := make([]Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		return nil
	}
	// [END getall]

	return greetings
}
