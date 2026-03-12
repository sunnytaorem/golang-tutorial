package pointers

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int // pointer variable that can hold the address of an int variable
	agePointer = &age

	fmt.Println("Age: ", *agePointer)
	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult Years: ", adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18 // this will change the original value of age (overrides the value of age with the new value)
}
