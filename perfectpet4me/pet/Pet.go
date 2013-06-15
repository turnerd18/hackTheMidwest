package pet


//import "fmt"


type Pet struct {
	Name string 
	Age string 				//Baby, Young, Adult, or Senior
	Sex string 				// M or F
	AnimalType string 		//Dog, Cat, Small&Furry, BarnYard, Bird, Horse, Pig, Rabbit, or Reptile
	Breed string
	Size string				//S, M, L, or XL
	ShelterID string		//ID of shelter the pet is at
	ContactInfo map[string]string
}


func NewPet (newName, newAge, newSex, newAnimalType, newBreed, newSize string, newShelterID string) *Pet {
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


func NewPet1(newName, newAge, newSex, newAnimalType, newBreed, newSize, newShelterID, newContactName, newPhone, newEmail, newAddress1, newAddress2, newCityState, newZip string) *Pet {
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

	return newPet
}

/*func main () {
	pet1 := NewPet("Fluffy", "Young", "M", "Dog", "Doberman", "L", "123456")

	fmt.Println(pet1.Name)
	fmt.Println(pet1.ShelterID)

	fmt.Println()

	pet2 := NewPet1("Reptar", "Adult", "M", "Reptile", "T-rex", "XL", "654321", "Tommy", "0987654321", "tommy@go.com", "1234 Street Ave", "Apt A1", "Lawrence, KS", "66044")

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
}*/
