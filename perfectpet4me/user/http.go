func init() {
    http.HanldeFunc("/user", start)
}

func start(w http.ResponseWriter, r *http.Request) {
}
