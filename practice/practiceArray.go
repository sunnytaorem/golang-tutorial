package practice

import "fmt"

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	var myHobbies [3]string = [3]string{"Coding", "Dancing", "Traveling"}
	fmt.Println(myHobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(myHobbies[0])
	funHobbies := myHobbies[1:3]
	fmt.Println(funHobbies)

	// 3) Create a slice based on the first task that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	sliced1Hobbies := myHobbies[:2]
	sliced2Hobbies := myHobbies[0:2]
	fmt.Println(sliced1Hobbies)
	fmt.Println(sliced2Hobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	fmt.Println(cap(sliced1Hobbies))
	reslicedHobbies := sliced1Hobbies[1:3] // here we need to specify the end index as 3, because sliced1Hobbies last element is index 2 although it has a capacity of 3.
	fmt.Println(reslicedHobbies)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Learn Go", "Build a project"}
	fmt.Println(courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Master Go"
	courseGoals = append(courseGoals, "Get a job")
	fmt.Println(courseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type Product struct {
		id    float64
		title string
		price float64
	}
	products := []Product{
		{id: 1, title: "Laptop", price: 999.99},
		{id: 2, title: "Phone", price: 499.99},
	}
	fmt.Println(products)
	products = append(products, Product{id: 3, title: "Tablet", price: 299.99})
	fmt.Println(products)
}
