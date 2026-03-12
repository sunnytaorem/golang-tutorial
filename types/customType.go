package types

import "fmt"

type customeString string // defining a new type customeString that is based on the built-in string type. This allows us to create a new type that has the same underlying representation as a string, but can have its own methods and behaviors.

func (s customeString) log() {
	fmt.Println(s)
}

func main() {
	var myString customeString // explicitly declaring a variable of type customeString
	myString = "Hello, World!"
	myString.log()
}
