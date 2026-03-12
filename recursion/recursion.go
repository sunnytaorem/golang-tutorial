package recursion

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

// recursion is when a function calls itself.
func factorial(n int) int {
	// result := 1

	// for i := 1; i <= n; i++ {
	// 	result *= i
	// }
	// return result

	// The recursive version of the factorial function would look like this:
	if n == 0 { // You need an exit condition to prevent infinite recursion.
		return 1
	}
	return n * factorial(n-1)
}
