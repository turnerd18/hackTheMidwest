package main


import "fmt"


type Pet struct {
	name string 
	age string 			//Baby, Young, Adult, or Senior
	sex string 			// M or F
	animalType string 	//Dog, Cat, Small&Furry, BarnYard, Bird, Horse, Pig, Rabbit, or Reptile
	breed string
	size string			//S, M, L, or XL
	shelterID int		//ID of shelter the pet is at
}


func NewPet (newName, newAge, newSex, newAnimalType, newBreed, newSize string, newShelterID int) *Pet {
	newPet := new(Pet)

	newPet.name = newName
	newPet.age = newAge
	newPet.sex = newSex
	newPet.animalType = newAnimalType
	newPet.breed = newBreed
	newPet.size = newSize
	newPet.shelterID = newShelterID

	return newPet
}


/*func main () {
	pet := NewPet("Fluffy", "Young", "M", "Dog", "Doberman", "L", 123456)

	fmt.Println(pet.name)
	fmt.Println(pet.shelterID)
}*/
