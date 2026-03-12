package practice

import (
	"bufio"
	"fmt"
	"go-project/note"
	"go-project/todo"
	"os"
	"strings"
)

// Interface is collection of method signatures that a type can implement.
// They are not about defining the logic of the methods, but rather about defining a contract that types can adhere to.
// It simply defines that a certain method exists
type saver interface { // common convention is to name the interface with an -er suffix, indicating that it represents a capability or action (if it has a single method).
	Save() error
}

// type displayer interface {
// 	Display()
// }

// embedding multiple interfaces into a new interface allows us to create a composite interface that combines the behaviors of both interfaces.
// This can be useful when we want to work with types that need to implement multiple behaviors, such as saving and displaying data in this case.
type outputtable interface {
	saver
	// displayer
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInputMultiLine("Todo Text:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
	result := add(1, 2)
	fmt.Println("Result of add:", result)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// interface{} is a special type in Go that can hold any value.
// It is often used when we want to work with values of unknown types or when we want to create functions that can accept any type of argument.
func printData(value interface{}) {
	// type assertion is a way to extract the underlying value of an interface{} type and check if it is of a specific type (in this case, int). The syntax value.(int) attempts to assert that the value stored in the interface{} is of type int.
	// If the assertion is successful, intVal will hold the integer value, and ok will be true. If the assertion fails (i.e., the value is not an int), intVal will be the zero value for int (0), and ok will be false.
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Value is an integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Value is a float:", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("Value is a string:", stringVal)
		return
	}

	switch value.(type) {
	case int:
		fmt.Println("Value is an integer:", value)
	case float64:
		fmt.Println("Value is a float:", value)
	case string:
		fmt.Println("Value is a string:", value)
	default:
		fmt.Println("Value is of unknown type:", value)
	}
}

// The saveData function takes an argument of type saver, which means it can accept any value that implements the saver interface.
// This allows us to pass both the todo and userNote objects to the saveData function, as long as they implement the Save() method defined in the saver interface.
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving data failed")
		return err
	}
	fmt.Println("Data saved successfully")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInputMultiLine("Note Title:")
	content := getUserInputMultiLine("Note Content:")
	return title, content
}

func getUserInput(prompt string) string {
	var input string
	fmt.Printf("%v ", prompt)
	fmt.Scanln(&input) //only reads single word input
	return input
}

func getUserInputMultiLine(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)   // bufio.NewReader creates a new reader that reads from the standard input (os.Stdin). This allows us to read user input from the console, including multi-line input if needed.
	input, err := reader.ReadString('\n') // reads until it encounters a line break \n. Here '\n' is a single byte character also known as a Rune in GO.

	if err != nil {
		fmt.Println(err)
		return ""
	}

	input = strings.TrimSuffix(input, "\n") // removes leading and trailing whitespace characters, including the newline character.
	input = strings.TrimSuffix(input, "\r") // removes leading and trailing whitespace characters, including the carriage return character.

	return input
}

// interface{} is too wide and unspecific.
// func add(a, b interface{}) interface{} {
// 	aInt, aIsInt := a.(int)
// 	bInt, bIsInt := b.(int)

// 	if aIsInt && bIsInt {
// 		return aInt + bInt
// 	}
// 	aFloat, aIsFloat := a.(float64)
// 	bFloat, bIsFloat := b.(float64)
// 	if aIsFloat && bIsFloat {
// 		return aFloat + bFloat
// 	}
// 	aStr, aIsStr := a.(string)
// 	bStr, bIsStr := b.(string)
// 	if aIsStr && bIsStr {
// 		return aStr + bStr
// 	}
// 	return nil
// }

// You can instead use generic functions with the help of Type placeholders [T]. We can pass [T any] which means that T can be any type, but in this case we want to restrict it to only int, float64, and string types.
// This allows us to create a more specific and type-safe add function that can handle only the types we want to support.
func add[T int | float64 | string](a, b T) T {
	return a + b
}
