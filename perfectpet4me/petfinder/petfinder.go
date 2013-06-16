package petfinder

import (
    "crypto/md5"
    "fmt"
    "io"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "strings"

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
    LastOffset int
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

type PetType struct {
    Animal map[string]string
    Options struct {
            Option [](map[string]string) `json:"option"`
        }
    ShelterPetId map[string]string
    Name map[string]string
    Contact struct {
        Email map[string]string
        Zip map[string]string
        City map[string]string
        Fax map[string]string
        Address1 map[string]string
        Phone map[string]string
        State map[string]string
        Address2 map[string]string
    }
    Description map[string]string
    Sex map[string]string
    Age map[string]string
    ShelterId map[string]string
    LastUpdate map[string]string
    Media struct {
        Photos struct {
            Photo [](map[string]string) `json:"photo"`
        }
    }
    Size map[string]string
    Id map[string]string
}

type PetFetcher struct {
    Petfinder struct {
        Pets struct {
            Pet PetType
        }
    }
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

    body, _ := pf.Execute(getstr)

    tf := new(TokenFetcher)
    err := json.Unmarshal(body, &tf)
    if err != nil {
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

func (pf *PetFinder) GetPet(animal string, location string) (*PetType) {
    errstr := ""
    method := "pet.find"
    args := map[string]string {
        "key"       : API_KEY,
        "token"     : pf.Token,
        "animal"    : animal,
        "location"  : location,
        "count"     : "1",
    }
    getstr := pf.RequestBuilder(method, args)

    body, errstr := pf.Execute(getstr)
    if errstr != "" {
        return nil
    }

    pfetch := new(PetFetcher)
    err := json.Unmarshal(body, &pfetch)
    if err != nil {
        fmt.Fprintf(pf.w, "%v\n", getstr)
        fmt.Fprintf(pf.w, "%v\n", body)
        fmt.Fprintf(pf.w, "%v\n", err.Error())
    }

    return &pfetch.Petfinder.Pets.Pet
/*
    p = NewPet(
        pfetch.Petfinder.Pets.Pet.Name["$t"],
        pfetch.Petfinder.Pets.Pet.Age["$t"],
        pfetch.Petfinder.Pets.Pet.Sex["$t"],
"border collie",//        pfetch.Petfinder.Pets.Pet.Breed,
        pfetch.Petfinder.Pets.Pet.Size["$t"],
        pfetch.Petfinder.Pets.Pet.ShelterId["$t"],
        pfetch.Petfinder.Pets.Pet.Contact.City["$t"],
        pfetch.Petfinder.Pets.Pet.Contact.Phone["$t"],
        pfetch.Petfinder.Pets.Pet.Contact.Email["$t"],
        pfetch.Petfinder.Pets.Pet.Contact.Address1["$t"],
        pfetch.Petfinder.Pets.Pet.Contact.Address2["$t"],
        pfetch.Petfinder.Pets.Pet.Contact.State["$t"],
        pfetch.Petfinder.Pets.Pet.Media.Photos.Photo
        )
*/


/*
    p.Name = "butthead"

    pets := []*pet.Pet{p}
    return pets
    */
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
