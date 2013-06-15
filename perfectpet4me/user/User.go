package user


//import "fmt"



type User struct {
	FirstName string
	LastName string
	Phone string
	Email string
	CityState string
	Zip string
}


func NewUser(newFirstName, newLastName, newPhone, newEmail, newCityState, newZip string) *User {
	newUser := new(User)

	newUser.FirstName = newFirstName
	newUser.LastName = newLastName
	newUser.Phone = newPhone
	newUser.Email = newEmail
	newUser.CityState = newCityState
	newUser.Zip = newZip
	
	return newUser
}


/*func main() {
	user1 := NewUser("John", "Smith", "1234567890", "jsmith@go.com", "lawrence,ks", "66044")

	fmt.Println(user1.FirstName)
	fmt.Println(user1.LastName)
	fmt.Println(user1.Phone)
	fmt.Println(user1.Email)
}*/



