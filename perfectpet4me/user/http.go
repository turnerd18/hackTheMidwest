package user

import (
	//"appengine"
	//"appengine/datastore"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/user", handler)
}

var userTmpls = template.Must(template.ParseFiles("templates/base.html", "perfectpet4me/user/templates/user.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	//User u :=

	if err := userTmpls.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
