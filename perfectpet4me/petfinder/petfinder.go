package petfinder

import (
    "crypto/md5"
    "fmt"
    "io"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "perfectpet4me/pet"
    "strings"

    "appengine"
    "appengine/urlfetch"
)

const API_URL string    = "http://api.petfinder.com/"
const API_KEY string    = "4b7c674b10cfc0793e935e8abd7346c2"
const SECRET string     = "15e0e406bc3bcc284773a57957492133"

type PetFinder struct {
    Token string
    Expires string
    w http.ResponseWriter
    r *http.Request
}

type TokenFetcher struct {
    Petfinder struct {
        Auth struct {
            Token map[string]string
            Expires map[string]string
            ExpiresString map[string]string
        }
    }
}

func md5hash(text string) string {
    h := md5.New()
    io.WriteString(h, text)
    return fmt.Sprintf("%x", h.Sum(nil))
}

func (pf *PetFinder) SetToken(w http.ResponseWriter, r *http.Request) (string) {
    method := "auth.getToken"
    errstr := ""
    args := map[string]string {
        "key" : API_KEY,
    }
    getstr := pf.RequestBuilder(method, args)
        fmt.Fprintf(pf.w, "%v\n", getstr)

    body, _ := pf.Execute(getstr)
/*
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)

    resp, err := client.Get(getstr)
    if err != nil {
        errstr = "couldn't open client"
        return errstr
    }

    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        errstr = "couldn't read body"
        return errstr
    }
*/
    tf := new(TokenFetcher)
    err := json.Unmarshal(body, &tf)
    if err != nil {
        fmt.Fprintf(w, "%v\n", getstr)
        fmt.Fprintf(w, "%v\n", body)
        fmt.Fprintf(w, "%v\n", err.Error())
        errstr = "couldn't unmarshal"
        return errstr
    }

    pf.Token = tf.Petfinder.Auth.Token["$t"]
    pf.Expires = tf.Petfinder.Auth.Expires["$t"]

    return errstr
}

func NewPetFinder(w http.ResponseWriter, r *http.Request) (*PetFinder) {
    pf := new(PetFinder)
    pf.w = w
    pf.r = r
    err := pf.SetToken(w, r)

    if err != "" {
        return nil
    }

    //apet := pf.GetRandomPet("dog", w, r)
    //fmt.Fprintf(w, "%v\n", apet.Name)

    return pf
}

func (pf *PetFinder) GetRandomPet(animal string) (*pet.Pet) {
    p := new(pet.Pet)

    errstr := ""
    method := "pet.find"
    args := map[string]string {
        "key"       : API_KEY,
        "token"     : pf.Token,
        "animal"    : animal,
    }
    getstr := pf.RequestBuilder(method, args)
        fmt.Fprintf(pf.w, "%v\n", getstr)

    body, errstr := pf.Execute(getstr)
    if errstr != "" {
        return nil
    }

    pfetch := new(TokenFetcher)
    err := json.Unmarshal(body, &pfetch)
    if err != nil {
        fmt.Fprintf(pf.w, "%v\n", getstr)
        fmt.Fprintf(pf.w, "%v\n", body)
        fmt.Fprintf(pf.w, "%v\n", err.Error())
    }

    p.Name = "butthead"

    return p
}

func (pf *PetFinder) RequestBuilder(apicall string, args map[string]string) string {
    split := strings.Split(apicall, ".")
    getstr := ""
    switch objtype := split[0]; objtype {
        case "pet":
            switch method := split[1]; method {
                case "get":
                case "getRandom":
                case "find":
                    getstr = "key=" + args["key"] +/* "&token=" + args["token"] +*/ "&animal=" + args["animal"] + "&format=json"
                    signature := md5hash(SECRET + getstr)
                    getstr = apicall + "?" + getstr + "&sig=" + signature
            }
        case "breed":
            switch method := split[1]; method {
                case "list":
            }
        case "shelter":
            switch method := split[1]; method {
                case "find":
                case "get":
                case "getPets":
                case "listByBreed":
            }
        case "auth":
            switch method := split[1]; method {
                case "getToken":
                    getstr = "key=" + args["key"] + "&format=json"
                    signature := md5hash(SECRET + getstr)
                    getstr = apicall + "?" + getstr + "&sig=" + signature
            }
    }

    return API_URL + getstr
}

func (pf *PetFinder) Execute(getstr string) ([]byte, string) {
    errstr := ""

    c := appengine.NewContext(pf.r)
    client := urlfetch.Client(c)

    resp, err := client.Get(getstr)
    if err != nil {
        errstr = "couldn't open client"
        return nil, errstr
    }

    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        errstr = "couldn't read body"
        return body, errstr
    }

    return body, errstr
}
