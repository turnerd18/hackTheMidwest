package match

import (
	"appengine"
	"appengine/datastore"
	aeuser "appengine/user"
	//"fmt"
	"html"
	"html/template"
	"net/http"
	"perfectpet4me/petfinder"
	"perfectpet4me/user"
)

func init() {
	http.HandleFunc("/search", lucky)
	http.HandleFunc("/search/results", luckyStrike)
    http.HandleFunc("/detail", detail)
}

var srchTmpls = template.Must(template.ParseFiles("templates/base.html", "perfectpet4me/match/templates/search.html"))
var resTmpls = template.Must(template.ParseFiles("templates/base.html", "perfectpet4me/match/templates/results.html"))
var detTmpls = template.Must(template.ParseFiles("templates/base.html", "perfectpet4me/match/templates/detail.html"))

func lucky(w http.ResponseWriter, r *http.Request) {
	if err := srchTmpls.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func luckyStrike(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	aeu := aeuser.Current(c)
	k := datastore.NewKey(c, "user", aeu.Email, 0, nil)

	var u user.User

	if err := datastore.Get(c, k, &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pf := petfinder.NewPetFinder(w, r)
	//fmt.Fprintf(w, "%v %v", r.FormValue("lucky"), u.Zip)
	foundPet := pf.GetPet(r.FormValue("lucky"), u.Zip)
	foundPet.Description = html.UnescapeString(foundPet.Description)

	//fmt.Fprintf(w, "%v", foundPet)
	if err := resTmpls.Execute(w, foundPet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func detail(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	aeu := aeuser.Current(c)
	k := datastore.NewKey(c, "user", aeu.Email, 0, nil)

	var u user.User

	if err := datastore.Get(c, k, &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pf := petfinder.NewPetFinder(w, r)
	//fmt.Fprintf(w, "%v %v", r.FormValue("lucky"), u.Zip)
	foundPets := pf.GetPets("dog", u.Zip, 4)

	//fmt.Fprintf(w, "%v", foundPets)
	if err := detTmpls.Execute(w, foundPets); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
