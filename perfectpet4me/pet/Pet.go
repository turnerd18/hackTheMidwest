package pet

type Pet struct {
	Name        string
	Age         string //Baby, Young, Adult, or Senior
	Sex         string // M or F
    Id          string
	AnimalType  string //Dog, Cat, Small&Furry, BarnYard, Bird, Horse, Pig, Rabbit, or Reptile
	Breed       string
	Size        string //S, M, L, or XL
	ShelterID   string //ID of shelter the pet is at
	ContactInfo map[string]string
	PictureURLs [5]map[string]string //map keys = {"x", "fpm", "pn", "pnt", "t"} in that order
}

func NewPet1(newName, newAge, newSex, newAnimalType, newBreed, newSize,
	newShelterID string) *Pet {

	newPet := new(Pet)

	newPet.Name = newName
	newPet.Age = newAge
	newPet.Sex = newSex
	newPet.AnimalType = newAnimalType
	newPet.Breed = newBreed
	newPet.Size = newSize
	newPet.ShelterID = newShelterID

	return newPet
}

func NewPet(newName, newAge, newSex, newAnimalType, newBreed, newSize, newShelterID,
	newContactName, newPhone, newEmail, newAddress1, newAddress2, newCityState, newZip string,
	newPic1, newPic2, newPic3, newPic4, newPic5 map[string]string) *Pet {

	newPet := new(Pet)

	newPet.Name = newName
	newPet.Age = newAge
	newPet.Sex = newSex
	newPet.AnimalType = newAnimalType
	newPet.Breed = newBreed
	newPet.Size = newSize
	newPet.ShelterID = newShelterID

	newPet.ContactInfo = make(map[string]string)
	newPet.ContactInfo["ContactName"] = newContactName
	newPet.ContactInfo["Phone"] = newPhone
	newPet.ContactInfo["Email"] = newEmail
	newPet.ContactInfo["Address1"] = newAddress1
	newPet.ContactInfo["Address2"] = newAddress2
	newPet.ContactInfo["CityState"] = newCityState
	newPet.ContactInfo["Zip"] = newZip

	newPet.PictureURLs[0] = make(map[string]string)
	newPet.PictureURLs[1] = make(map[string]string)
	newPet.PictureURLs[2] = make(map[string]string)
	newPet.PictureURLs[3] = make(map[string]string)
	newPet.PictureURLs[4] = make(map[string]string)

	newPet.PictureURLs[0] = newPic1
	newPet.PictureURLs[1] = newPic2
	newPet.PictureURLs[2] = newPic3
	newPet.PictureURLs[3] = newPic4
	newPet.PictureURLs[4] = newPic5

	return newPet
}

/*func main () {
	pet1 := NewPet1("Fluffy", "Young", "M", "Dog", "Doberman", "L", "123456")

	fmt.Println(pet1.Name)
	fmt.Println(pet1.ShelterID)

	fmt.Println()

	urlmap1 := map[string]string {
		"x"		:	"url1a.com",
		"fpm"	:	"url1b.com",
		"pn"	:	"url1c.com",
		"pnt"	:	"url1d.com",
		"t"		:	"url1e.com"	}

	urlmap2 := map[string]string {
		"x"		:	"url2a.com",
		"fpm"	:	"url2b.com",
		"pn"	:	"url2c.com",
		"pnt"	:	"url2d.com",
		"t"		:	"url2e.com"	}

	urlmap3 := map[string]string {
		"x"		:	"url3a.com",
		"fpm"	:	"url3b.com",
		"pn"	:	"url3c.com",
		"pnt"	:	"url3d.com",
		"t"		:	"url3e.com"	}

	urlmap4 := map[string]string {
		"x"		:	"url4a.com",
		"fpm"	:	"url4b.com",
		"pn"	:	"url4c.com",
		"pnt"	:	"url4d.com",
		"t"		:	"url4e.com"	}

	urlmap5 := map[string]string {
		"x"		:	"url5a.com",
		"fpm"	:	"url5b.com",
		"pn"	:	"url5c.com",
		"pnt"	:	"url5d.com",
		"t"		:	"url5e.com"	}



	pet2 := NewPet("Reptar", "Adult", "M", "Reptile", "T-rex", "XL", "654321", "Tommy",
		"0987654321", "tommy@go.com", "1234 Street Ave", "Apt A1", "Lawrence, KS",
		"66044", urlmap1, urlmap2, urlmap3, urlmap4, urlmap5)

	fmt.Println(pet2.Name)
	fmt.Println(pet2.Age)
	fmt.Println(pet2.Sex)
	fmt.Println(pet2.AnimalType)
	fmt.Println(pet2.Breed)
	fmt.Println(pet2.Size)
	fmt.Println(pet2.ShelterID)
	fmt.Println(pet2.ContactInfo["ContactName"])
	fmt.Println(pet2.ContactInfo["Phone"])
	fmt.Println(pet2.ContactInfo["Email"])
	fmt.Println(pet2.ContactInfo["Address1"])
	fmt.Println(pet2.ContactInfo["Address2"])
	fmt.Println(pet2.ContactInfo["CityState"])
	fmt.Println(pet2.ContactInfo["Zip"])

	fmt.Println(pet2.PictureURLs[0]["x"])
	fmt.Println(pet2.PictureURLs[0]["fpm"])
	fmt.Println(pet2.PictureURLs[0]["pn"])
	fmt.Println(pet2.PictureURLs[0]["pnt"])
	fmt.Println(pet2.PictureURLs[0]["t"])

	fmt.Println(pet2.PictureURLs[1]["x"])
	fmt.Println(pet2.PictureURLs[1]["fpm"])
	fmt.Println(pet2.PictureURLs[1]["pn"])
	fmt.Println(pet2.PictureURLs[1]["pnt"])
	fmt.Println(pet2.PictureURLs[1]["t"])

	fmt.Println(pet2.PictureURLs[2]["x"])
	fmt.Println(pet2.PictureURLs[2]["fpm"])
	fmt.Println(pet2.PictureURLs[2]["pn"])
	fmt.Println(pet2.PictureURLs[2]["pnt"])
	fmt.Println(pet2.PictureURLs[2]["t"])

	fmt.Println(pet2.PictureURLs[3]["x"])
	fmt.Println(pet2.PictureURLs[3]["fpm"])
	fmt.Println(pet2.PictureURLs[3]["pn"])
	fmt.Println(pet2.PictureURLs[3]["pnt"])
	fmt.Println(pet2.PictureURLs[3]["t"])

	fmt.Println(pet2.PictureURLs[4]["x"])
	fmt.Println(pet2.PictureURLs[4]["fpm"])
	fmt.Println(pet2.PictureURLs[4]["pn"])
	fmt.Println(pet2.PictureURLs[4]["pnt"])
	fmt.Println(pet2.PictureURLs[4]["t"])
}*/
