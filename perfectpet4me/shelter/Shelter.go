package main


//import "fmt"


type Shelter struct {
	Name string
	Phone string
	Email string
	ID string
	Address1 string
	Address2 string
	State string
	Zip dtring
}


func NewShelter (newName, newPhone, newEmail, newID, newAddress1, newAddress2, newState, newZip) *Shelter {
	newShelter := new(Shelter)

	newShelter.Name = newName
	newShelter.Phone = newPhone
	newShelter.Email = newEmail
	newShelter.ID = newID
	newShelter.Address1 = newAddress1
	newShelter.Address2 = newAddress2
	newShelter.State = newState
	newShelter.Zip = newZip

	return newShelter
}


