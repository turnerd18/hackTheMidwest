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

    testpet := pf.GetPets("dog","66067", 5)
    fmt.Fprintf(w, "%v\n", testpet[0].Name)
}
