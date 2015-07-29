package guestbook

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func userPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["id"]
	fmt.Fprintf(w, "user %s", userid)
}

type jsonUser struct {
	ID       uint64
	Username string
}

func handleGetUser(r *http.Request) interface{} {
	useridString := r.FormValue("ID")
	userid, err := strconv.ParseUint(useridString, 10, 64)
	if err != nil {
		return nil
	}

	response := &jsonUser{
		ID:       userid,
		Username: "foobar",
	}

	return response
}
