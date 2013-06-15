package pet


import "fmt"


type Pet struct {
	Name string 
	Age string 			//Baby, Young, Adult, or Senior
	Sex string 			// M or F
	AnimalType string 	//Dog, Cat, Small&Furry, BarnYard, Bird, Horse, Pig, Rabbit, or Reptile
	Breed string
	Size string			//S, M, L, or XL
	ShelterID int		//ID of shelter the pet is at
}


func NewPet (newName, newAge, newSex, newAnimalType, newBreed, newSize string, newShelterID int) *Pet {
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


/*func main () {
	pet := NewPet("Fluffy", "Young", "M", "Dog", "Doberman", "L", 123456)

	fmt.Println(pet.Name)
	fmt.Println(pet.ShelterID)
}*/
