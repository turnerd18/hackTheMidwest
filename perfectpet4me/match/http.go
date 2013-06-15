package match

import (
	"net/http"
)

func init() {
	http.HandleFunc("/match", start)
}

func start(w http.ResponseWriter, r *http.Request) {
}
