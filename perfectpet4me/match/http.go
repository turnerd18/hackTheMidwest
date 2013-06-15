package match

import (
	"net/http"
)

func init() {
	http.HanldeFunc("/match", start)
}

func start(w http.ResponseWriter, r *http.Request) {
}
