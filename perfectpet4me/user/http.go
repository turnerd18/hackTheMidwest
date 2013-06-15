package user

import (
	"net/http"
)

func init() {
	http.HandleFunc("/user", handler)
}

var userTmpls = template.Must(template.ParseFiles("hackTheMidwest/perfectpet4me/templates/base.html", "templates/login.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	if err := listTmpl.Execute(w, tc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
