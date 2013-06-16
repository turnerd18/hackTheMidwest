package perfectpet4me

import (
    "fmt"
    "net/http"
    "perfectpet4me/petfinder"
)

func init() {
    http.HandleFunc("/pftest", start)
}

func start(w http.ResponseWriter, r *http.Request) {
    pf := petfinder.NewPetFinder(w,r)
    if pf != nil {
        fmt.Fprintf(w, "%v\n", pf.Token)
    } else {
        fmt.Fprintf(w, "%v\n", "error")
    }

    testpet := pf.GetPet("dog","66061")
    fmt.Fprintf(w, "%v\n", testpet.Name)
    fmt.Fprintf(w, "%v\n", testpet.Id)
    fmt.Fprintf(w, "%v\n", testpet)
}
