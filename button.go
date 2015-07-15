package guestbook

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func button(w http.ResponseWriter, r *http.Request) {
	if err := buttonTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type clickData struct {
	IPAddress string
}

type buttonStats struct {
	Clicks int `json:"clicks"`
}

func clickKey(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "Clicks", "default_click", 0, nil)
}

func buttonClicked(w http.ResponseWriter, r *http.Request) {
	click := &clickData{
		IPAddress: r.RemoteAddr,
	}

	c := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(c, "Click", clickKey(c))
	_, err := datastore.Put(c, key, click)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clickCount, err := datastore.NewQuery("Click").Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stats := &buttonStats{
		Clicks: clickCount,
	}

	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
