package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2) // Create a function that multiplies a number by 2
	triple := createTransformer(3) // Create a function that multiplies a number by 3

	// Using an anonymous function as a parameter to another function.
	// Anonymous functions are functions that do not have a name. They are also known as lambda functions in other programming languages.
	// They are useful when you want to create a function that is only used once and does not need to be reused elsewhere in the code.
	transformed := transformNumbers(&numbers, func(num int) int { // NOTE that the function we are passing her is a value parameter for the transformNumbers function. It is not a type.
		return num * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// Return an anonymous function from another function. This is also known as a closure.
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// Why are anonymous functions called closures?
// Anonymous functions are called closures because they can "close over" variables from their surrounding scope.
// This means that they can access and modify variables that are defined outside of the function, even after the outer function has returned.
// Which means the anonymous function returned are locked in by the variables in the outer function's scope. It won't be impacted by a later call with another factor value. Each call to createTransformer creates a new closure that captures the specific factor value for that call.
// In the example above, the anonymous function returned by createTransformer can access the factor variable, which is defined in the outer function's scope.
// This allows us to create a new function that multiplies a number by a specific factor, without having to define a separate named function for each factor we want to use.
