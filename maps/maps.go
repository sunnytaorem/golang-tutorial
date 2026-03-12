package maps

import "fmt"

func main() {
	// Maps are a collection of key-value pairs. They are similar to a struct, but not entirely.
	// They are the python disctionaries equivalent in Go. They are also known as hash maps or hash tables in other programming languages.
	// The key value pairs in a map are not ordered. They are also mutable, which means that you can change the value of a key after it has been created.
	// The key makes it easy to access the value associated with it, instead of having to remember the index of the value in a list or array.

	// websites := []string{"https://google.com", "https://facebook.com", "https://twitter.com", "https://aws.com"}
	// map[key]value
	websitesMap := map[string]string{
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
		"Twitter":  "https://twitter.com",
		"AWS":      "https://aws.com",
	}
	fmt.Println(websitesMap)
	fmt.Println(websitesMap["Google"])

	// Adding a new key-value pair to the map (mutable)
	websitesMap["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websitesMap)

	// Deleting a key-value pair from the map
	delete(websitesMap, "Twitter")
	fmt.Println(websitesMap)

	// Why use maps? Instead of Structs
	// Maps are more flexible than structs because they can have any number of key-value pairs, while structs have a fixed number of fields.
	// For Maps, the keys can be of any type, while for Structs, the fields must be of a specific type. The key can be a struct, an array, integer, etc.
	// Both Maps and Structs solves different problems.
}
