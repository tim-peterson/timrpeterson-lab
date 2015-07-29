package guestbook

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

var (
	box              *rice.Box
	indexTemplate    *template.Template
	aboutTemplate    *template.Template
	buttonTemplate   *template.Template
	ideaTemplate     *template.Template
	commentsTemplate *template.Template
)

func init() {
	// setup templates
	box = rice.MustFindBox("templates")
	templateTree := template.Must(template.New("guestbook").Parse(""))
	templateTree.Parse(box.MustString("header.html.tmpl"))
	templateTree.Parse(box.MustString("footer.html.tmpl"))
	indexTemplate = template.Must(templateTree.New("index").Parse(box.MustString("index.html.tmpl")))
	aboutTemplate = template.Must(templateTree.New("about").Parse(box.MustString("about.html.tmpl")))
	buttonTemplate = template.Must(templateTree.New("button").Delims("{[", "]}").Parse(box.MustString("button.html.tmpl")))
	ideaTemplate = template.Must(templateTree.New("idea").Parse(box.MustString("idea.html.tmpl")))
	commentsTemplate = template.Must(templateTree.New("comments").Parse(box.MustString("comments.html.tmpl")))

	// setup http routing
	router := mux.NewRouter()
	router.HandleFunc("/", newGenericPage(rootPage, indexTemplate))
	router.HandleFunc("/sign", signPage)
	router.HandleFunc("/about", newGenericPage(aboutPage, aboutTemplate))
	router.HandleFunc("/button", buttonPage)
	router.HandleFunc("/buttonClicked", buttonClickedPage)
	router.HandleFunc("/idea", ideaPage)
	router.HandleFunc("/submittedIdea", submittedIdeaPage)
	router.HandleFunc("/comments", commentsPage)
	router.HandleFunc("/comments.json", handleCommentsPage).Methods("POST")
	router.HandleFunc("/comments.json", newGenericJSON(handleCommentsGet)).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", userPage)

	http.Handle("/", router)
	//http.Handle("/", http.FileServer(http.Dir("./public")))
	//	log.Println("Server started: http://localhost:" + port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c context.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

type genericHandlerFunc func(*http.Request) interface{}

func newGenericPage(handler genericHandlerFunc, tmpl *template.Template) http.HandlerFunc {
	type genericPageData struct {
		Username string
		Data     interface{}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		data := &genericPageData{
			Username: "foobar",
		}

		data.Data = handler(r)
		if data.Data == nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func newGenericJSON(handler genericHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := handler(r)
		if data == nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		// stream the contents of the file to the response
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
