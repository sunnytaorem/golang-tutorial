package structs

import (
	"fmt"
	"go-project/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date (DD/MM/YYYY): ")
	// These data can be stored in a struct, which is a custom data type that groups together related data.
	// We can create an instance of the struct and assign values to its fields.

	// Using a constructor function to create an instance of the User struct.
	userData, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	// create an instance of the User struct and assign values to its fields.
	// User{} is a struct literal that initializes a new User struct with the specified field values.
	var userD *user.User
	// userD = &user.User{
	// 	FirstName: firstName,
	// 	LastName:  lastName,
	// 	BirthDate: birthDate,
	// }
	userD = &user.User{
		FirstName: firstName,
	}

	admin, err := user.NewAdmin("sunny@gmail.com", "3333", userD)
	if err != nil {
		fmt.Println("Error creating admin:", err)
		return
	}
	// userData = newUser(firstName, lastName, birthDate)
	// you can also create an instance like this:
	// var userData User
	// userData = User{
	// 	firstName,
	// 	lastName,
	// 	birthDate,
	// 	time.Now(),
	// }
	// you can also instantiate a struct with empty values and then assign values to its fields later.
	// you can also leave fields empty when instantiating a struct, and they will be initialized with their zero values (e.g., empty string for string fields, zero for int fields, etc.).
	// var userData2 User
	// userData2 = User{}
	// userData2.FirstName = firstName
	// userData2.LastName = lastName
	// userData2.BirthDate = birthDate
	// userData2.CreatedAt = time.Now()

	userData.OutputUserDetails()
	userData.CalculateAge()
	// user.ClearUserName()
	// user.OutputUserDetails()
	// userD.OutputUserDetails()

	admin.OutputUserDetails()
	admin.CalculateAge()

	OutputUserDetailsWithPointer(userD)
	// outputUserDetailsWithPointer(&user)
	// fmt.Printf("User %s %s is %d years old\n", user.FirstName, user.LastName, user.CalculateAge())
}

func getUserData(promtText string) string {
	fmt.Print(promtText)
	var input string
	fmt.Scanln(&input)
	return input
}

func OutputUserDetailsWithPointer(u *user.User) {
	// fmt.Print((*user).firstName) // This is the technically correct way to access fields of a struct through a pointer, you can use the dot notation directly without dereferencing the pointer. The Go compiler automatically dereferences the pointer for you when you access its fields.
	fmt.Printf("Hello %s\n", u.FirstName)
}
