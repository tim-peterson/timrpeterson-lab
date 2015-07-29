package guestbook

import (
	//"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

func ideaPage(w http.ResponseWriter, r *http.Request) {
	if err := ideaTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Idea struct {
	Author string
	Idea   string
	Date   time.Time
}

type ideaData struct {
	IPAddress string
}

type ideaStats struct {
	Ideas int `json:"ideas"`
}

func ideaKey(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "Ideas", "default_ideas", 0, nil)
}

func submittedIdeaPage(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Access-Control-Allow-Origin", "*")

	c := appengine.NewContext(r)
	i := Idea{
		Idea: r.FormValue("idea"),
		Date: time.Now(),
	}
	if u := user.Current(c); u != nil {
		i.Author = u.String()
	}
	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "SubmittedIdea", ideaKey(c))
	_, err := datastore.Put(c, key, &i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

/*func buttonClicked(w http.ResponseWriter, r *http.Request) {
	idea := &ideaData{
		IPAddress: r.RemoteAddr,
	}

	c := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(c, "Idea", ideaKey(c))
	_, err := datastore.Put(c, key, idea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ideaCount, err := datastore.NewQuery("Ideas").Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stats := &ideaStats{
		Ideas: ideaCount,
	}

	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}*/
