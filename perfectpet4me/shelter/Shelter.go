package main


import "fmt"


type Shelter struct {
	name string
	phone string
	email string
	ID string
	address1 string
	address2 string
	state string
	zip dtring
}


func NewShelter (newName, newPhone, newEmail, newID, newAddress1, newAddress2, newState, newZip) *Shelter {
	newShelter := new(Shelter)

	newShelter.name = newName
	newShelter.phone = newPhone
	newShelter.email = newEmail
	newShelter.ID = newID
	newShelter.address1 = newAddress1
	newShelter.address2 = newAddress2
	newShelter.state = newState
	newShelter.zip = newZip

	return newShelter
}


