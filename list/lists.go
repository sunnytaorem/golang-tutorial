package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

// Array is useful to store data of the same type and a lot of it. In python, its called list. They come with fixed size, which means that once we declare an array with a certain size, we cannot change its size.
// The [] syntax indicates that we are declaring a slice, which is a dynamic array that can grow and shrink in size. The type of the elements in the slice is specified after the [] syntax (in this case, float64).
func main() {
	prices := []float64{10.99, 9.99}
	fmt.Println(prices[1])
	prices[1] = 8.99
	updatedPrices := append(prices, 11.99, 5.99) // creates a new slice with the new values added to the end of the original slice. The original slice remains unchanged, and the new slice contains all the elements of the original slice plus the new elements.
	fmt.Println(updatedPrices, prices)

	prices = append(prices, 11.99, 5.99) // this will update the original slice with the new values added to the end of it.
	fmt.Println(prices)

	discountPrices := []float64{5.99, 6.99, 8.99, 10.99, 2.99}
	prices = append(prices, discountPrices...) // the ... syntax is used to unpack the elements of the discountPrices slice and add them individually to the prices slice. Because prices is a list of floats not a list of lists.
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"Laptop"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0} // this is stored in memory
// 	fmt.Println(prices)

// 	productNames[2] = "Phone"
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	// Slices can be a subset of an array or another slice.
// 	// When we create a slice from an array, it does not create a new copy of the array. Instead, it creates a reference to the original array.
// 	// This means that if we modify the slice, it will also modify the original array.
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 18.99 // this will change the value in the original array as well, because slices are references to the underlying array.
// 	highlightedPrices := featuredPrices[:1]

// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)

// 	// For every slice, there is a length and a capacity meta data
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))

// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }

// // Important note: In using slices, we can't go on the left of the starting index of a first slice it is based on. You can always access the values on the right of the index.
// // HighlightedPrices which is based on featuredPrices, which is bases on the starting slice prices[1:], has a starting index of 1 in the original array.
// // So, any slices based on highlightedPrices can't go on the left of index 1 in the original array. Therefore, highlightedPrices has a capacity of 3 and not 4. Its length is 1.
