package user

import (
	"errors"
	"fmt"
	"time"
)

// Define a struct to hold user data (UserData is a custom data type that groups together related data)
type User struct {
	FirstName string // field names should start with a capital letter if we want to access them from other packages.
	lastName  string
	birthDate string
	createdAt time.Time
}

// You can also define another struct that embeds the User struct, which allows you to reuse the fields and methods of the User struct in the Admin struct.
// This is a common way to create a hierarchy of related types in Go.
// In other languages, this might be achieved through inheritance, but in Go, we use composition (embedding) to achieve similar functionality.
type Admin struct {
	User // annonymous embedding
	// User     User
	email    string
	password string
}

// You can add methods to a struct by defining functions with a receiver of the struct type. This allows you to associate behavior with the data stored in the struct. For example, you could add a method to the User struct that calculates the user's age based on their birth date.
func (u User) CalculateAge() int { // (u User) is the receiver of the method, which allows us to access the fields of the User struct within the method.
	birthDate, err := time.Parse("02/01/2006", u.birthDate)
	if err != nil {
		fmt.Println("Error parsing birth date:", err)
		return 0 // return a default age if there's an error
	}
	age := time.Now().Year() - birthDate.Year()

	return age // return a dummy age for demonstration purposes
}

func (u User) OutputUserDetails() {
	fmt.Printf("Hello %s %s, born on %s\n", u.FirstName, u.lastName, u.birthDate)
}

// (u *User) is a pointer receiver, which allows us to modify the fields of the User struct within the method.
// When you use a pointer receiver, you can change the values of the struct's fields, and those changes will be reflected outside the method.
func (u *User) ClearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

// Constructors are not a built-in feature in Go, but you can create a constructor function that initializes and returns an instance of a struct.
// This is a common pattern in Go to provide a convenient way to create and initialize structs with specific values.
// You can return a pointer to the struct from the constructor function, which allows you to work with the struct instance without having to worry about copying it around.
// This is especially useful for larger structs or when you want to modify the struct's fields after creation.
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields must be provided")
	}
	return &User{
		FirstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string, user *User) (*Admin, error) {
	if email == "" || password == "" || user == nil {
		return nil, errors.New("all fields must be provided")
	}
	return &Admin{
		User:     *user,
		email:    email,
		password: password,
	}, nil
}
