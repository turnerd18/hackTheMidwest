package perfectpet4me

import (
	"appengine"
	//"appengine/datastore"
	"appengine/user"
	//"html/template"
	"fmt"
	"net/http"
)

type PetLooker struct {
	Email string
	Name  string
	City  string
	State string
	Zip   string
	Phone string
}

func init() {
	http.HanldeFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	pl := PetLooker{}

	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Hello, %v!", u)
}
