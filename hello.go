package guestbook

import (
	"html/template"
	"net/http"
	"time"

	"github.com/GeertJohan/go.rice"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

var (
	box               *rice.Box
	guestbookTemplate *template.Template
	aboutTemplate     *template.Template
	buttonTemplate    *template.Template
	ideaTemplate    	*template.Template
	commentsTemplate  *template.Template
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
	http.HandleFunc("/about", about)
	http.HandleFunc("/button", button)
	http.HandleFunc("/buttonClicked", buttonClicked)
	http.HandleFunc("/idea", idea)
	http.HandleFunc("/submittedIdea", submittedIdea)
	http.HandleFunc("/comments", comments)

	http.HandleFunc("/comments.json", handleComments)
		//http.Handle("/", http.FileServer(http.Dir("./public")))
	//	log.Println("Server started: http://localhost:" + port)
		//log.Fatal(http.ListenAndServe(":"+port, nil))

	box = rice.MustFindBox("templates")

	templateTree := template.Must(template.New("guestbook").Parse(""))
	templateTree.Parse(box.MustString("header.html.tmpl"))
	templateTree.Parse(box.MustString("footer.html.tmpl"))
	guestbookTemplate = template.Must(templateTree.New("index").Parse(box.MustString("index.html.tmpl")))
	aboutTemplate = template.Must(templateTree.New("about").Parse(box.MustString("about.html.tmpl")))

	buttonTemplate = template.Must(templateTree.New("button").Delims("{[", "]}").Parse(box.MustString("button.html.tmpl")))

	ideaTemplate = template.Must(templateTree.New("idea").Parse(box.MustString("idea.html.tmpl")))
	commentsTemplate = template.Must(templateTree.New("comments").Parse(box.MustString("comments.html.tmpl")))

}



// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c context.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// [END getall]

	if err := guestbookTemplate.Execute(w, greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
