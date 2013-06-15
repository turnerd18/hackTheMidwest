package user

import (
	"net/http"
)

func init() {
	http.HandleFunc("/user", start)
}

func start(w http.ResponseWriter, r *http.Request) {
}
