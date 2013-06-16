package user

type User struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
	CityState string
	Zip       string
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
