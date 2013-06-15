package perfectpet4me

import (
    "fmt"
    "net/http"
    "perfectpet4me/petfinder"
//    "perfectpet4me/pet"
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

    testpet := pf.GetRandomPet("dog")
    fmt.Fprintf(w, "%v\n", testpet.Name)
}
