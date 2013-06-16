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
    "strconv"

    "appengine"
    "appengine/urlfetch"
)

/**********************************************************
                        Constants
 *********************************************************/

const API_URL string    = "http://api.petfinder.com/"
const API_KEY string    = "4b7c674b10cfc0793e935e8abd7346c2"
const SECRET string     = "15e0e406bc3bcc284773a57957492133"

/**********************************************************
                    Data Structures
 *********************************************************/

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

type PetsFetcher struct {
    /*
    Petfinder struct {
        Pets struct {
            Pet struct {
                Animal map[string]string
            }
        }
    }
    */
    Lo string `json:"petfinder>lastOffset>$t`
}

/**********************************************************
                        Methods
 *********************************************************/
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

    tf := new(TokenFetcher)
    err := json.Unmarshal(body, &tf)
    if err != nil {
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

    return pf
}

func (pf *PetFinder) GetPets(animal string, location string, numResults int) ([]*pet.Pet) {
    p := new(pet.Pet)

    errstr := ""
    method := "pet.find"
    args := map[string]string {
        "key"       : API_KEY,
        "token"     : pf.Token,
        "animal"    : animal,
        "location"  : location,
        "count"     : strconv.Itoa(numResults),
    }
    getstr := pf.RequestBuilder(method, args)
        fmt.Fprintf(pf.w, "%v\n", getstr)

    body, errstr := pf.Execute(getstr)
    if errstr != "" {
        return nil
    }
/*
    fmt.Fprintf(pf.w, "BODY: %v\n", string(body))
    str := string(body)
    if i := strings.Index(str, "\"pets\" : ["); i != -1 {
        str = str[i:]
    }
    if i := strings.Index(str, "]},\"header\""); i != -1 {
        str = str[:i]
    }
    fmt.Fprintf(pf.w, "BODY: %v\n", str)
*/


    var f interface{}
    //pfetch := new(PetsFetcher)
    err := json.Unmarshal(body, &f)
    //err := json.Unmarshal(body, &pfetch)
    if err != nil {
        fmt.Fprintf(pf.w, "%v\n", getstr)
        fmt.Fprintf(pf.w, "%v\n", body)
        fmt.Fprintf(pf.w, "%v\n", err.Error())
    }

    fmt.Fprintf(pf.w, "BODY: %v\n", string(body))
    a := f.(map[string]interface{})
    b := a["petfinder"].(map[string]interface{})
    c := b["pets"].(map[string]interface{})
    fmt.Fprintf(pf.w, "PETSFETCHER: %v\n", c)

    p.Name = "butthead"

    pets := []*pet.Pet{p}
    return pets
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
                    getstr = "key=" + args["key"] + "&token=" + args["token"] + "&animal=" + args["animal"] + "&location=" + args["location"] + "&count=" + args["count"] +"&format=json"
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
