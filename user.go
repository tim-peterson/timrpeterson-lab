package guestbook

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func userPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["id"]
	fmt.Fprintf(w, "user %s", userid)
}
