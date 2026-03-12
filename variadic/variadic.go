package variadic

import "fmt"

func main() {
	numbers := []int{10, 15, 10, 5}
	result := sumup(numbers)
	fmt.Println(result)

	// using variadic function
	resultVariadic := sumupVariadic(10, 15, 10, 5)
	fmt.Println(resultVariadic)

	// you can also pass a slice to a variadic function using the ... syntax after the slice variable.
	resultVariadicFromSlice := sumupVariadic(0, numbers...)
	fmt.Println(resultVariadicFromSlice)
}

func sumup(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Variadic functions are functions that can take a variable number of arguments. They are defined using the ... syntax in the function signature.
// The variadic parameter must be the last parameter in the function signature, and it is treated as a slice of the specified type within the function body.
func sumupVariadic(startValue int, numbers ...int) int {
	sum := startValue

	for _, num := range numbers {
		sum += num
	}
	return sum
}
