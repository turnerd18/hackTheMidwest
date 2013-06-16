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

    testpet := pf.GetPets("dog","66061", 5)
    for i := 0; i < 5; i++ {
        fmt.Fprintf(w, "%v\n", testpet[i].Name)
        fmt.Fprintf(w, "%v\n", testpet[i])
    }
}
