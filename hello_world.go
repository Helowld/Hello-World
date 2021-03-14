// Official Website: https://golang.org/

/* package main - Every go file must start with the package name statement.
 * Packages are used to provide code compartmentalisation and reusability.
 * 'main' is the entry point for every go file.
 */
package main

// import "fmt" - The fmt package is imported and it will be used inside the main function to print text to the standard output.
import "fmt"

// func main() - The main is a special function. The program execution starts from the main function.
// The main function should always reside in the main package. The { and } indicate the start and end of the main function.
func main() {
	// fmt.Println("Hello World") - The Println function of the fmt package is used to write text to the standard output.
	fmt.Println("Hello, World!")
}

// ~$ go run hello_world.go
// Output: hello world

/* Learn more:
 * https://golangbot.com/learn-golang-series/
 */
