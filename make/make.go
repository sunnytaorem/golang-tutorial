package make

import (
	"fmt"
)

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := []string{}
	fmt.Println(userNames)

	// Adding elements to the slice using the append function
	userNames = append(userNames, "Sunny")
	userNames = append(userNames, "John")

	fmt.Println(userNames)

	// If you know the number of elements in advance, you can create a slice with a specific length and capacity.
	userNamesFixed := make([]string, 2, 5) // This creates an empty array with length 2 and capacity 5
	fmt.Println(userNamesFixed)

	// Adding elements to the slice using the append function
	userNamesFixed = append(userNamesFixed, "Tom")
	userNamesFixed = append(userNamesFixed, "Miguel")
	userNamesFixed = append(userNamesFixed, "Sara")
	userNamesFixed = append(userNamesFixed, "Anna")
	fmt.Println(userNamesFixed) // This will add the new elements to the end of the slice, and it will leave the first two elements as empty strings because we initialized the slice with a length of 2.

	// Adding elements to the slice using index
	userNamesFixed[0] = "Sunny"
	// userNamesFixed[1] = "John"
	fmt.Println(userNamesFixed)

	courseRatings := make(map[string]float64, 3) // This creates a map with an initial capacity of 3. By preassigning the capacity, we can optimize the performance of the map by reducing the number of times it needs to reallocate memory as we add elements to it.

	courseRatings["Go"] = 4.5
	courseRatings["Python"] = 4.7
	courseRatings["Java"] = 4.3
	courseRatings["JavaScript"] = 4.6 // only when it exceeds the initial capacity of 3, the map will need to reallocate memory to accommodate the new element.

	fmt.Println(courseRatings)

	// You can also define custom type for the map. It helps to make the code more readable.
	courseRatingsCustom := make(floatMap, 3)
	courseRatingsCustom["Go"] = 4.5
	courseRatingsCustom["Python"] = 4.7
	courseRatingsCustom["Java"] = 4.3

	courseRatingsCustom.output()

	// Iterating over the slice and map using the range keyword
	for index, value := range userNamesFixed {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
