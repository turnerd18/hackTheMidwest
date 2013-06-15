package main


import "fmt"


type User struct {
	firstName string
	lastName string
	phone string
	email string
	cityState string
	zip string
}


func NewUser(newFirstName, newLastName, newPhone, newEmail, newCityState, newZip string) *User {
	newUser := new(User)

	newUser.firstName = newFirstName
	newUser.lastName = newLastName
	newUser.phone = newPhone
	newUser.email = newEmail
	newUser.cityState = newCityState
	newUser.zip = newZip
	
	return newUser
}


/*func main() {
	user1 := NewUser("John", "Smith", "1234567890", "jsmith@go.com", "lawrence,ks", "66044")

	fmt.Println(user1.firstName)
	fmt.Println(user1.lastName)
	fmt.Println(user1.phone)
	fmt.Println(user1.email)
}*/



