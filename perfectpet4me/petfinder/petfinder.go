package petfinder

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	petp "perfectpet4me/pet"
	"strconv"
	"strings"

	"appengine"
	"appengine/urlfetch"
	simplejson "github.com/bitly/go-simplejson"
)

/**********************************************************
                       Constants
*********************************************************/

const API_URL string = "http://api.petfinder.com/"
const API_KEY string = "4b7c674b10cfc0793e935e8abd7346c2"
const SECRET string = "15e0e406bc3bcc284773a57957492133"

/**********************************************************
                   Data Structures
*********************************************************/

type PetFinder struct {
	Token      string
	Expires    string
	w          http.ResponseWriter
	r          *http.Request
	LastOffset int
}

type TokenFetcher struct {
	Petfinder struct {
		Auth struct {
			Token         map[string]string
			Expires       map[string]string
			ExpiresString map[string]string
		}
	}
}

type PetType struct {
	Animal  map[string]string
	Options struct {
		Option [](map[string]string) `json:"option"`
	}
	ShelterPetId map[string]string
	Name         map[string]string
	Contact      struct {
		Email    map[string]string
		Zip      map[string]string
		City     map[string]string
		Fax      map[string]string
		Address1 map[string]string
		Phone    map[string]string
		State    map[string]string
		Address2 map[string]string
	}
	Description map[string]string
	Sex         map[string]string
	Age         map[string]string
	ShelterId   map[string]string
	LastUpdate  map[string]string
	Media       struct {
		Photos struct {
			Photo [](map[string]string) `json:"photo"`
		}
	}
	Size map[string]string
	Id   map[string]string
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

func (pf *PetFinder) SetToken(w http.ResponseWriter, r *http.Request) string {
	method := "auth.getToken"
	errstr := ""
	args := map[string]string{
		"key": API_KEY,
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

func NewPetFinder(w http.ResponseWriter, r *http.Request) *PetFinder {
	pf := new(PetFinder)
	pf.w = w
	pf.r = r
	err := pf.SetToken(w, r)

	if err != "" {
		return nil
	}

	return pf
}

func (pf *PetFinder) GetPet(animal string, location string) *petp.Pet {
	errstr := ""
	method := "pet.find"
	args := map[string]string{
		"key":      API_KEY,
		"token":    pf.Token,
		"animal":   animal,
		"location": location,
		"count":    "1",
	}
	getstr := pf.RequestBuilder(method, args)

	body, errstr := pf.Execute(getstr)
	if errstr != "" {
		return nil
	}

	j, _ := simplejson.NewJson(body)
	p := new(petp.Pet)
	jo := j.Get("petfinder").Get("pets").Get("pet")
	p.Id, _ = jo.Get("id").Get("$t").String()
	p.Name, _ = jo.Get("name").Get("$t").String()
	p.Age, _ = jo.Get("age").Get("$t").String()
	p.AnimalType, _ = jo.Get("animal").Get("$t").String()
	p.Breed, _ = jo.Get("breeds").Get("breed").Get("$t").String()
	p.Sex, _ = jo.Get("sex").Get("$t").String()
	p.Description, _ = jo.Get("description").Get("$t").String()
	x, _ := jo.Get("media").Get("photos").Get("photo").GetIndex(0).Get("$t").String()
	t, _ := jo.Get("media").Get("photos").Get("photo").GetIndex(4).Get("$t").String()
	p.PictureURLs[0] = map[string]string{
		"x": x,
		"t": t,
	}
	//temp, _ := jo.Get("media").Get("photos").Get("photo").GetIndex(0).Get("$t").String()
	//fmt.Fprintf(pf.w, "%v\n", temp)

	return p
	/*
		    fmt.Fprintf(pf.w, "%v\n", p.Breed)

		    pfetch := new(PetFetcher)
		    err := json.Unmarshal(body, &pfetch)
		    if err != nil {
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
}

func (pf *PetFinder) GetPets(animal string, location string, numResults int) []*petp.Pet {
	errstr := ""
	method := "pet.find"
	args := map[string]string{
		"key":      API_KEY,
		"token":    pf.Token,
		"animal":   animal,
		"location": location,
		"count":    strconv.Itoa(numResults),
	}
	getstr := pf.RequestBuilder(method, args)

	body, errstr := pf.Execute(getstr)
	if errstr != "" {
		return nil
	}

	j, _ := simplejson.NewJson(body)
	var pets []*petp.Pet
	for i := 0; i < numResults; i++ {
		p := new(petp.Pet)
		jo := j.Get("petfinder").Get("pets").Get("pet").GetIndex(i)
		p.Id, _ = jo.Get("id").Get("$t").String()
		p.Name, _ = jo.Get("name").Get("$t").String()
		p.Age, _ = jo.Get("age").Get("$t").String()
		p.AnimalType, _ = jo.Get("animal").Get("$t").String()
		p.Breed, _ = jo.Get("breeds").Get("breed").Get("$t").String()
		p.Sex, _ = jo.Get("sex").Get("$t").String()
		x, _ := jo.Get("media").Get("photos").Get("photo").GetIndex(0).Get("$t").String()
		t, _ := jo.Get("media").Get("photos").Get("photo").GetIndex(4).Get("$t").String()
		p.PictureURLs[0] = map[string]string{
			"x": x,
			"t": t,
		}
		pets = append(pets, p)
	}
	fmt.Fprintf(pf.w, "%v\n", pets)
	//temp, _ := jo.Get("media").Get("photos").Get("photo").GetIndex(0).Get("$t").String()
	//fmt.Fprintf(pf.w, "%v\n", temp)

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
			getstr = "key=" + args["key"] + "&token=" + args["token"] + "&animal=" + args["animal"] + "&location=" + args["location"] + "&count=" + args["count"] + "&format=json"
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
