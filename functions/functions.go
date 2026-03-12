package functions

import "fmt"

// We can also use functions as parameters to other functions. This is known as higher-order functions. It allows us to create more flexible and reusable code.
// We can also return functions from other functions, which is known as closures. Closures are a powerful feature of Go that allows us to create functions that can access variables from their enclosing scope, even after the outer function has returned.
func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9, 10}

	doubled := transformNumbers(&nums1, double)
	tripled := transformNumbers(&nums1, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&nums1)
	transformerFn2 := getTransformerFunction(&nums2)
	transformedNumbers1 := transformNumbers(&nums1, transformerFn1)
	transformedNumbers2 := transformNumbers(&nums2, transformerFn2)

	fmt.Println(transformedNumbers1)
	fmt.Println(transformedNumbers2)
}

// Create a custom type for the transform function. It is helpful when the function has a lot of parameters.
type transformFunc func(int) int

// type multiParamFunc func(int, []string, map[string][]int) ([]int, string)

// using functions as first class values
func transformNumbers(nums *[]int, transform transformFunc) []int {
	dNumbers := []int{}
	for _, val := range *nums {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// returning a function from another function
func getTransformerFunction(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
