/**
 * This file provided by Facebook is for non-commercial testing and evaluation purposes only.
 * Facebook reserves all rights not expressly granted.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 * FACEBOOK BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package guestbook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	//"log"
	"net/http"
	"os"
	"sync"
	"time"

  "golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)


type Comment struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Date    time.Time
}

/*type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}*/

const dataFile = "./comments.json"

var commentMutex = new(sync.Mutex)

func comments(w http.ResponseWriter, r *http.Request) {
	if err := commentsTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func commentsKey(c context.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Comments", "default_comment", 0, nil)
}

// Handle comments
func handleComments(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	comm := Comment{
		Text: r.FormValue("text"),
		Date:    time.Now(),
	}
	if u := user.Current(c); u != nil {
		comm.Author = u.String()
	}
	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Comments", commentsKey(c))
	_, err := datastore.Put(c, key, &comm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Since multiple requests could come in at once, ensure we have a lock
	// around all file operations
	commentMutex.Lock()
	defer commentMutex.Unlock()

	// Stat the file, so we can find its current permissions
	fi, err := os.Stat(dataFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to stat the data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	// Read the comments from the file.
	commentData, err := ioutil.ReadFile(dataFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read the data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":
		// Decode the JSON data
		comments := make([]Comment, 0)
		if err := json.Unmarshal(commentData, &comments); err != nil {
			http.Error(w, fmt.Sprintf("Unable to Unmarshal comments from data file (%s): %s", dataFile, err), http.StatusInternalServerError)
			return
		}

		// Add a new comment to the in memory slice of comments
		comments = append(comments, Comment{Author: r.FormValue("author"), Text: r.FormValue("text")})

		// Marshal the comments to indented json.
		commentData, err = json.MarshalIndent(comments, "", "    ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to marshal comments to json: %s", err), http.StatusInternalServerError)
			return
		}

		// Write out the comments to the file, preserving permissions
		err := ioutil.WriteFile(dataFile, commentData, fi.Mode())
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to write comments to data file (%s): %s", dataFile, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		io.Copy(w, bytes.NewReader(commentData))

	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		// stream the contents of the file to the response
		io.Copy(w, bytes.NewReader(commentData))

	default:
		// Don't know the method, so error
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method), http.StatusMethodNotAllowed)
	}
}

// Handle comments
/*func handleComments(w http.ResponseWriter, r *http.Request) {
	// Since multiple requests could come in at once, ensure we have a lock
	// around all file operations
	commentMutex.Lock()
	defer commentMutex.Unlock()

	// Stat the file, so we can find its current permissions
	fi, err := os.Stat(dataFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to stat the data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	// Read the comments from the file.
	commentData, err := ioutil.ReadFile(dataFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read the data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":
		// Decode the JSON data
		comments := make([]comment, 0)
		if err := json.Unmarshal(commentData, &comments); err != nil {
			http.Error(w, fmt.Sprintf("Unable to Unmarshal comments from data file (%s): %s", dataFile, err), http.StatusInternalServerError)
			return
		}

		// Add a new comment to the in memory slice of comments
		comments = append(comments, comment{Author: r.FormValue("author"), Text: r.FormValue("text")})

		// Marshal the comments to indented json.
		commentData, err = json.MarshalIndent(comments, "", "    ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to marshal comments to json: %s", err), http.StatusInternalServerError)
			return
		}

		// Write out the comments to the file, preserving permissions
		err := ioutil.WriteFile(dataFile, commentData, fi.Mode())
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to write comments to data file (%s): %s", dataFile, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		io.Copy(w, bytes.NewReader(commentData))

	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		// stream the contents of the file to the response
		io.Copy(w, bytes.NewReader(commentData))

	default:
		// Don't know the method, so error
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method), http.StatusMethodNotAllowed)
	}
}
*/
